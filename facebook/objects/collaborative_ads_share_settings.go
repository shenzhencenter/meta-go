package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CollaborativeAdsShareSettings struct {
	AgencyBusiness        *Business `json:"agency_business,omitempty"`
	ID                    *core.ID  `json:"id,omitempty"`
	ProductCatalogProxyID *core.ID  `json:"product_catalog_proxy_id,omitempty"`
	UtmCampaign           *string   `json:"utm_campaign,omitempty"`
	UtmMedium             *string   `json:"utm_medium,omitempty"`
	UtmSource             *string   `json:"utm_source,omitempty"`
}
