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

var ShadowIGUserCatalogProductSearchFields = struct {
	ImageURL        string
	IsCheckoutFlow  string
	MerchantID      string
	ProductID       string
	ProductName     string
	ProductVariants string
	RetailerID      string
	ReviewStatus    string
}{
	ImageURL:        "image_url",
	IsCheckoutFlow:  "is_checkout_flow",
	MerchantID:      "merchant_id",
	ProductID:       "product_id",
	ProductName:     "product_name",
	ProductVariants: "product_variants",
	RetailerID:      "retailer_id",
	ReviewStatus:    "review_status",
}
