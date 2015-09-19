package campbx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	prefix     = "/api"
	timeout    = 5 * time.Second
	tickerPath = "xticker"
	depthPath  = "xdepth"
)

// Client communicates with the API
type Client struct {
	Host       string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		Host:       "https://campbx.com",
		HTTPClient: &http.Client{Timeout: timeout},
	}
}

// url returns a full URL for a resource
func (c *Client) url(resource string) string {
	return fmt.Sprintf("%v%v/%v.php", c.Host, prefix, resource)
}

// get a response from a URL
func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.HTTPClient.Get(url)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetTicker returns a Ticker from the API
func (c *Client) GetTicker() (*Ticker, error) {
	var body []byte
	var err error

	if body, err = c.get(c.url(tickerPath)); err != nil {
		return nil, err
	}

	var ticker Ticker

	if err = json.Unmarshal(body, &ticker); err != nil {
		return nil, err
	}

	return &ticker, nil
}

// GetOrderBook returns the OrderBook from the API
func (c *Client) GetOrderBook() (*OrderBook, error) {
	var body []byte
	var err error

	if body, err = c.get(c.url(depthPath)); err != nil {
		return nil, err
	}

	// this is ugly, I want to Unmarshal this in one shot.
	// get the float array data first
	holder := struct {
		Asks [][]float32 `json:"Asks"`
		Bids [][]float32 `json:"Bids"`
	}{}

	if err := json.Unmarshal(body, &holder); err != nil {
		return nil, err
	}

	// now convert those []float32 arrays of price/quantity values to Order
	// structs

	asks := make([]Order, len(holder.Asks))

	for i, order := range holder.Asks {
		asks[i] = NewOrder(order)
	}

	bids := make([]Order, len(holder.Bids))

	for i, order := range holder.Bids {
		bids[i] = NewOrder(order)
	}

	orderBook := OrderBook{
		Asks: asks,
		Bids: bids,
	}

	return &orderBook, nil
}

func NewOrder(order []float32) Order {
	return Order{
		Price:    order[0],
		Quantity: order[1],
	}

}
