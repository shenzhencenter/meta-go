package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAssetOnsiteDestinations struct {
	AutoOptimization           *string  `json:"auto_optimization,omitempty"`
	DetailsPageProductID       *core.ID `json:"details_page_product_id,omitempty"`
	ShopCollectionProductSetID *core.ID `json:"shop_collection_product_set_id,omitempty"`
	Source                     *string  `json:"source,omitempty"`
	StorefrontShopID           *core.ID `json:"storefront_shop_id,omitempty"`
}

var AdAssetOnsiteDestinationsFields = struct {
	AutoOptimization           string
	DetailsPageProductID       string
	ShopCollectionProductSetID string
	Source                     string
	StorefrontShopID           string
}{
	AutoOptimization:           "auto_optimization",
	DetailsPageProductID:       "details_page_product_id",
	ShopCollectionProductSetID: "shop_collection_product_set_id",
	Source:                     "source",
	StorefrontShopID:           "storefront_shop_id",
}
