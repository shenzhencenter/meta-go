package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalogCheckMarketplacePartnerSellersStatus struct {
	SampleErrors *[]map[string]interface{} `json:"sample_errors,omitempty"`
	SessionID    *core.ID                  `json:"session_id,omitempty"`
	Status       *string                   `json:"status,omitempty"`
}
