package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Shop struct {
	CommerceMerchantSettings *CommerceMerchantSettings `json:"commerce_merchant_settings,omitempty"`
	FbSalesChannel           *map[string]interface{}   `json:"fb_sales_channel,omitempty"`
	ID                       *core.ID                  `json:"id,omitempty"`
	IgSalesChannel           *map[string]interface{}   `json:"ig_sales_channel,omitempty"`
	ShopStatus               *string                   `json:"shop_status,omitempty"`
	Workspace                *map[string]interface{}   `json:"workspace,omitempty"`
}
