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
