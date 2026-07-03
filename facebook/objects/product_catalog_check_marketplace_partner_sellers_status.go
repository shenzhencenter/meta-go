package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalogCheckMarketplacePartnerSellersStatus struct {
	SampleErrors *[]map[string]interface{} `json:"sample_errors,omitempty"`
	SessionID    *core.ID                  `json:"session_id,omitempty"`
	Status       *string                   `json:"status,omitempty"`
}

var ProductCatalogCheckMarketplacePartnerSellersStatusFields = struct {
	SampleErrors string
	SessionID    string
	Status       string
}{
	SampleErrors: "sample_errors",
	SessionID:    "session_id",
	Status:       "status",
}
