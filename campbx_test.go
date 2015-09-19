package campbx

import (
	"fmt"
	"testing"
)

func TestGetTicker(t *testing.T) {
	c := NewClient()
	ticker, err := c.GetTicker()

	if err != nil {
		panic(fmt.Errorf("GetTicker : %v", err))
	}
	fmt.Printf("%+v\n", ticker)
}

func TestGetDepth(t *testing.T) {
	c := NewClient()
	orderBook, err := c.GetDepth()

	if err != nil {
		panic(fmt.Errorf("GetDepth : %v", err))
	}
	fmt.Printf("%+v\n", orderBook)
}
