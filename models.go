package campbx

// Ticker
//
// Sample response
//
//     {"Last Trade":"244.99","Best Bid":"236.38","Best Ask":"244.99"}
type Ticker struct {
	LastTrade float32 `json:"Last Trade,string"`
	Bid       float32 `json:"Best Bid,string"`
	Ask       float32 `json:"Best Ask,string"`
}

// OrderBook represents the full order book returned by the API.
//
// Sample response/structure
//
//     { "Asks":[ [ 244.99, 0.990 ], ... ], "Bids":[ [ 236.38, 0.020 ], ... ] }
type OrderBook struct {
	Asks []Order `json:"Asks"`
	Bids []Order `json:"Bids"`
}

// Order represents the price and quanty of an individual Order, or the summary
// of multiple Orders (as in the case of an Order Book)
type Order struct {
	Price    float32 `json:"price"`
	Quantity float32 `json:"amount"`
}
