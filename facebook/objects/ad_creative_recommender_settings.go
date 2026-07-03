package objects

type AdCreativeRecommenderSettings struct {
	PreferredEvents     *[]string `json:"preferred_events,omitempty"`
	ProductSalesChannel *string   `json:"product_sales_channel,omitempty"`
}

var AdCreativeRecommenderSettingsFields = struct {
	PreferredEvents     string
	ProductSalesChannel string
}{
	PreferredEvents:     "preferred_events",
	ProductSalesChannel: "product_sales_channel",
}
