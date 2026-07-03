package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CreatorAssetCreative struct {
	ID                    *core.ID `json:"id,omitempty"`
	ImageURL              *string  `json:"image_url,omitempty"`
	ModerationStatus      *string  `json:"moderation_status,omitempty"`
	ProductItemRetailerID *core.ID `json:"product_item_retailer_id,omitempty"`
	ProductURL            *string  `json:"product_url,omitempty"`
	RetailerID            *core.ID `json:"retailer_id,omitempty"`
	VideoURL              *string  `json:"video_url,omitempty"`
}
