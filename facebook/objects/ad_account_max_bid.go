package objects

type AdAccountMaxBid struct {
	MaxBid *int `json:"max_bid,omitempty"`
}

var AdAccountMaxBidFields = struct {
	MaxBid string
}{
	MaxBid: "max_bid",
}
