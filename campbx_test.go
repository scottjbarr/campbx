package campbx

import (
	"fmt"
	"testing"
)

func TestHandlesNonExistingItems(t *testing.T) {
	c := NewClient()
	ticker, err := c.GetTicker()

	if err != nil {
		panic(fmt.Errorf("GetTicker : %v", err))
	}
	fmt.Printf("%+v\n", ticker)
}
