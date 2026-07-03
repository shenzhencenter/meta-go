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
