package campbx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	host       = "https://campbx.com/api"
	timeout    = 5
	tickerPath = "xticker"
	depthPath  = "xdepth"
)

// Client communicates with the API
type Client struct {
	Timeout int
}

func NewClient() *Client {
	return &Client{
		Timeout: timeout,
	}
}

// url returns a full URL for a resource
func (c *Client) url(resource string) string {
	return fmt.Sprintf("%v/%v.php", host, resource)
}

// get a response from a URL
func (c *Client) get(url string) ([]byte, error) {
	resp, err := http.Get(url)

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

// GetDepth returns the OrderBook from the API
func (c *Client) GetDepth() (*OrderBook, error) {
	var body []byte
	var err error

	if body, err = c.get(c.url(depthPath)); err != nil {
		return nil, err
	}

	var orderBook OrderBook

	if err := json.Unmarshal(body, &orderBook); err != nil {
		return nil, err
	}

	return &orderBook, nil
}
