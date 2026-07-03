package objects

type BusinessPartnerPremiumOptions struct {
	EnableBasketInsight               *bool                   `json:"enable_basket_insight,omitempty"`
	EnableExtendedAudienceRetargeting *bool                   `json:"enable_extended_audience_retargeting,omitempty"`
	RetailerCustomAudienceConfig      *map[string]interface{} `json:"retailer_custom_audience_config,omitempty"`
}
