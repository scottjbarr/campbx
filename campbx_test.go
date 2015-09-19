package campbx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
)

// Mock out HTTP requests.
//
// Pinched from http://keighl.com/post/mocking-http-responses-in-golang/
// Thanks, Kyle Truscott (@keighl)!
func httpMock(code int,
	body string) (*httptest.Server, *Client) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport, Timeout: 5 * time.Second}

	// setting the host in the Client so I don't need to totally fake out
	// the TLS config
	client := &Client{
		Host:       "http://campbx.com",
		HTTPClient: httpClient,
	}

	return server, client
}

// Test helper. Thanks again, @keighl
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a))
	}
}

func TestGetTicker(t *testing.T) {
	body := `{"Last Trade":"244.99","Best Bid":"236.38","Best Ask":"244.99"}`
	server, client := httpMock(200, body)
	defer server.Close()

	ticker, err := client.GetTicker()

	if err != nil {
		t.Errorf("GetTicker : %v", err)
	}

	expect(t, ticker.LastTrade, float32(244.99))
	expect(t, ticker.Bid, float32(236.38))
	expect(t, ticker.Ask, float32(244.99))
}

func TestGetDepth(t *testing.T) {
	body := `{"Asks":[[244.99, 0.99]], "Bids":[[236.38, 0.02], [234.01, 1.8]]}`
	server, client := httpMock(200, body)
	defer server.Close()

	orderBook, err := client.GetDepth()

	if err != nil {
		t.Errorf("GetDepth : %v", err)
	}

	expect(t, len(orderBook.Asks), 1)
	expect(t, orderBook.Asks[0].Price, float32(244.99))
	expect(t, orderBook.Asks[0].Quantity, float32(0.99))
	expect(t, len(orderBook.Bids), 2)
	expect(t, orderBook.Bids[0].Price, float32(236.38))
	expect(t, orderBook.Bids[0].Quantity, float32(0.02))

}
