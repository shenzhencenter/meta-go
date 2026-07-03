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

var CreatorAssetCreativeFields = struct {
	ID                    string
	ImageURL              string
	ModerationStatus      string
	ProductItemRetailerID string
	ProductURL            string
	RetailerID            string
	VideoURL              string
}{
	ID:                    "id",
	ImageURL:              "image_url",
	ModerationStatus:      "moderation_status",
	ProductItemRetailerID: "product_item_retailer_id",
	ProductURL:            "product_url",
	RetailerID:            "retailer_id",
	VideoURL:              "video_url",
}
