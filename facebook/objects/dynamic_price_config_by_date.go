package objects

type DynamicPriceConfigByDate struct {
	CheckinDate  *string                   `json:"checkin_date,omitempty"`
	Prices       *string                   `json:"prices,omitempty"`
	PricesPretty *[]map[string]interface{} `json:"prices_pretty,omitempty"`
}

var DynamicPriceConfigByDateFields = struct {
	CheckinDate  string
	Prices       string
	PricesPretty string
}{
	CheckinDate:  "checkin_date",
	Prices:       "prices",
	PricesPretty: "prices_pretty",
}
