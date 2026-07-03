package objects

type ProductItemInsights struct {
	AdClickCount      *int `json:"ad_click_count,omitempty"`
	AdImpressionCount *int `json:"ad_impression_count,omitempty"`
	AddToCartCount    *int `json:"add_to_cart_count,omitempty"`
	PurchaseCount     *int `json:"purchase_count,omitempty"`
	ViewContentCount  *int `json:"view_content_count,omitempty"`
}

var ProductItemInsightsFields = struct {
	AdClickCount      string
	AdImpressionCount string
	AddToCartCount    string
	PurchaseCount     string
	ViewContentCount  string
}{
	AdClickCount:      "ad_click_count",
	AdImpressionCount: "ad_impression_count",
	AddToCartCount:    "add_to_cart_count",
	PurchaseCount:     "purchase_count",
	ViewContentCount:  "view_content_count",
}
