package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductGroup struct {
	ID             *core.ID          `json:"id,omitempty"`
	ProductCatalog *ProductCatalog   `json:"product_catalog,omitempty"`
	RetailerID     *core.ID          `json:"retailer_id,omitempty"`
	Variants       *[]ProductVariant `json:"variants,omitempty"`
}
