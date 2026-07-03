package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductGroup struct {
	ID             *core.ID          `json:"id,omitempty"`
	ProductCatalog *ProductCatalog   `json:"product_catalog,omitempty"`
	RetailerID     *core.ID          `json:"retailer_id,omitempty"`
	Variants       *[]ProductVariant `json:"variants,omitempty"`
}

var ProductGroupFields = struct {
	ID             string
	ProductCatalog string
	RetailerID     string
	Variants       string
}{
	ID:             "id",
	ProductCatalog: "product_catalog",
	RetailerID:     "retailer_id",
	Variants:       "variants",
}
