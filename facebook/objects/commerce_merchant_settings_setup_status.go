package objects

type CommerceMerchantSettingsSetupStatus struct {
	DealsSetup                       *string                 `json:"deals_setup,omitempty"`
	MarketplaceApprovalStatus        *string                 `json:"marketplace_approval_status,omitempty"`
	MarketplaceApprovalStatusDetails *map[string]interface{} `json:"marketplace_approval_status_details,omitempty"`
	PaymentSetup                     *string                 `json:"payment_setup,omitempty"`
	ReviewStatus                     *map[string]interface{} `json:"review_status,omitempty"`
	ShopSetup                        *string                 `json:"shop_setup,omitempty"`
}
