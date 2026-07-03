package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalogCheckMarketplacePartnerDealsStatus struct {
	Errors    *[]map[string]interface{} `json:"errors,omitempty"`
	SessionID *core.ID                  `json:"session_id,omitempty"`
	Status    *string                   `json:"status,omitempty"`
}

var ProductCatalogCheckMarketplacePartnerDealsStatusFields = struct {
	Errors    string
	SessionID string
	Status    string
}{
	Errors:    "errors",
	SessionID: "session_id",
	Status:    "status",
}
