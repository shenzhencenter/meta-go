package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeShopSpec struct {
	CollectionID *core.ID `json:"collection_id,omitempty"`
	LandingView  *string  `json:"landing_view,omitempty"`
	ShopID       *core.ID `json:"shop_id,omitempty"`
}

var AdCreativeShopSpecFields = struct {
	CollectionID string
	LandingView  string
	ShopID       string
}{
	CollectionID: "collection_id",
	LandingView:  "landing_view",
	ShopID:       "shop_id",
}
