package objects

type TargetingGeoLocationMarket struct {
	Country    *string `json:"country,omitempty"`
	Key        *string `json:"key,omitempty"`
	MarketType *string `json:"market_type,omitempty"`
	Name       *string `json:"name,omitempty"`
}
