package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserAvailableCatalogs struct {
	CatalogID    *core.ID `json:"catalog_id,omitempty"`
	CatalogName  *string  `json:"catalog_name,omitempty"`
	ProductCount *int     `json:"product_count,omitempty"`
	ShopName     *string  `json:"shop_name,omitempty"`
}
