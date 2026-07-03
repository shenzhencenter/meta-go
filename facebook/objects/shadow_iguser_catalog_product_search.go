package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGUserCatalogProductSearch struct {
	ImageURL        *string                              `json:"image_url,omitempty"`
	IsCheckoutFlow  *bool                                `json:"is_checkout_flow,omitempty"`
	MerchantID      *core.ID                             `json:"merchant_id,omitempty"`
	ProductID       *core.ID                             `json:"product_id,omitempty"`
	ProductName     *string                              `json:"product_name,omitempty"`
	ProductVariants *[]ShadowIGUserCatalogProductVariant `json:"product_variants,omitempty"`
	RetailerID      *core.ID                             `json:"retailer_id,omitempty"`
	ReviewStatus    *string                              `json:"review_status,omitempty"`
}
