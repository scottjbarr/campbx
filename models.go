package campbx

// Ticker
// Sample response
//
//     {"Last Trade":"244.99","Best Bid":"236.38","Best Ask":"244.99"}
type Ticker struct {
	LastTrade float32 `json:"Last Trade,string"`
	Bid       float32 `json:"Best Bid,string"`
	Ask       float32 `json:"Best Ask,string"`
}
