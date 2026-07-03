package objects

type TargetingGeoLocationMarket struct {
	Country    *string `json:"country,omitempty"`
	Key        *string `json:"key,omitempty"`
	MarketType *string `json:"market_type,omitempty"`
	Name       *string `json:"name,omitempty"`
}

var TargetingGeoLocationMarketFields = struct {
	Country    string
	Key        string
	MarketType string
	Name       string
}{
	Country:    "country",
	Key:        "key",
	MarketType: "market_type",
	Name:       "name",
}
