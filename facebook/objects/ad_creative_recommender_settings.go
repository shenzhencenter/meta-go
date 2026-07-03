package objects

type AdCreativeRecommenderSettings struct {
	PreferredEvents     *[]string `json:"preferred_events,omitempty"`
	ProductSalesChannel *string   `json:"product_sales_channel,omitempty"`
}
