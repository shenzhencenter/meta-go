package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeShopSpec struct {
	CollectionID *core.ID `json:"collection_id,omitempty"`
	LandingView  *string  `json:"landing_view,omitempty"`
	ShopID       *core.ID `json:"shop_id,omitempty"`
}
