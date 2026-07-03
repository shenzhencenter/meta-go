package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductCatalogCheckMarketplacePartnerDealsStatus struct {
	Errors    *[]map[string]interface{} `json:"errors,omitempty"`
	SessionID *core.ID                  `json:"session_id,omitempty"`
	Status    *string                   `json:"status,omitempty"`
}
