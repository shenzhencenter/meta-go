package objects

type CampaignGroupBrandConfiguration struct {
	BrandProductName *string `json:"brand_product_name,omitempty"`
	Locale           *string `json:"locale,omitempty"`
	Vertical         *string `json:"vertical,omitempty"`
}

var CampaignGroupBrandConfigurationFields = struct {
	BrandProductName string
	Locale           string
	Vertical         string
}{
	BrandProductName: "brand_product_name",
	Locale:           "locale",
	Vertical:         "vertical",
}
