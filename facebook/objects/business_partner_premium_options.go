package objects

type BusinessPartnerPremiumOptions struct {
	EnableBasketInsight               *bool                   `json:"enable_basket_insight,omitempty"`
	EnableExtendedAudienceRetargeting *bool                   `json:"enable_extended_audience_retargeting,omitempty"`
	RetailerCustomAudienceConfig      *map[string]interface{} `json:"retailer_custom_audience_config,omitempty"`
}

var BusinessPartnerPremiumOptionsFields = struct {
	EnableBasketInsight               string
	EnableExtendedAudienceRetargeting string
	RetailerCustomAudienceConfig      string
}{
	EnableBasketInsight:               "enable_basket_insight",
	EnableExtendedAudienceRetargeting: "enable_extended_audience_retargeting",
	RetailerCustomAudienceConfig:      "retailer_custom_audience_config",
}
