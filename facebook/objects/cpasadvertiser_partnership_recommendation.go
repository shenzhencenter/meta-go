package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CPASAdvertiserPartnershipRecommendation struct {
	AdvertiserBusinessID *core.ID  `json:"advertiser_business_id,omitempty"`
	BrandBusinessID      *core.ID  `json:"brand_business_id,omitempty"`
	Brands               *[]string `json:"brands,omitempty"`
	Countries            *[]string `json:"countries,omitempty"`
	ID                   *core.ID  `json:"id,omitempty"`
	MerchantBusinessID   *core.ID  `json:"merchant_business_id,omitempty"`
	MerchantCategories   *[]string `json:"merchant_categories,omitempty"`
	Status               *string   `json:"status,omitempty"`
	StatusReason         *string   `json:"status_reason,omitempty"`
}
