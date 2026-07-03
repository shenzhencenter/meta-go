package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductGroupGet struct {
	ID                        *core.ID                  `json:"id,omitempty"`
	MiniShopsProductSetsCount *int                      `json:"mini_shops_product_sets_count,omitempty"`
	ProductCatalog            *map[string]interface{}   `json:"product_catalog,omitempty"`
	Products                  *map[string]interface{}   `json:"products,omitempty"`
	RepresentativeItemID      *core.ID                  `json:"representative_item_id,omitempty"`
	RetailerID                *core.ID                  `json:"retailer_id,omitempty"`
	Variants                  *[]map[string]interface{} `json:"variants,omitempty"`
}

var ProductGroupGetFields = struct {
	ID                        string
	MiniShopsProductSetsCount string
	ProductCatalog            string
	Products                  string
	RepresentativeItemID      string
	RetailerID                string
	Variants                  string
}{
	ID:                        "id",
	MiniShopsProductSetsCount: "mini_shops_product_sets_count",
	ProductCatalog:            "product_catalog",
	Products:                  "products",
	RepresentativeItemID:      "representative_item_id",
	RetailerID:                "retailer_id",
	Variants:                  "variants",
}
