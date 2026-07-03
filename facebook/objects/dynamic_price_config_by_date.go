package objects

type DynamicPriceConfigByDate struct {
	CheckinDate  *string                   `json:"checkin_date,omitempty"`
	Prices       *string                   `json:"prices,omitempty"`
	PricesPretty *[]map[string]interface{} `json:"prices_pretty,omitempty"`
}
