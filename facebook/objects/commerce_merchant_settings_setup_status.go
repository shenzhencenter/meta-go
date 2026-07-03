package objects

type CommerceMerchantSettingsSetupStatus struct {
	DealsSetup                       *string                 `json:"deals_setup,omitempty"`
	MarketplaceApprovalStatus        *string                 `json:"marketplace_approval_status,omitempty"`
	MarketplaceApprovalStatusDetails *map[string]interface{} `json:"marketplace_approval_status_details,omitempty"`
	PaymentSetup                     *string                 `json:"payment_setup,omitempty"`
	ReviewStatus                     *map[string]interface{} `json:"review_status,omitempty"`
	ShopSetup                        *string                 `json:"shop_setup,omitempty"`
}

var CommerceMerchantSettingsSetupStatusFields = struct {
	DealsSetup                       string
	MarketplaceApprovalStatus        string
	MarketplaceApprovalStatusDetails string
	PaymentSetup                     string
	ReviewStatus                     string
	ShopSetup                        string
}{
	DealsSetup:                       "deals_setup",
	MarketplaceApprovalStatus:        "marketplace_approval_status",
	MarketplaceApprovalStatusDetails: "marketplace_approval_status_details",
	PaymentSetup:                     "payment_setup",
	ReviewStatus:                     "review_status",
	ShopSetup:                        "shop_setup",
}
