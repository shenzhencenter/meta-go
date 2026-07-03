package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeLinkDataChildAttachment struct {
	AndroidURL              *string                         `json:"android_url,omitempty"`
	CallToAction            *AdCreativeLinkDataCallToAction `json:"call_to_action,omitempty"`
	Caption                 *string                         `json:"caption,omitempty"`
	Description             *string                         `json:"description,omitempty"`
	IgMediaID               *core.ID                        `json:"ig_media_id,omitempty"`
	ImageCrops              *AdsImageCrops                  `json:"image_crops,omitempty"`
	ImageHash               *string                         `json:"image_hash,omitempty"`
	IosURL                  *string                         `json:"ios_url,omitempty"`
	Link                    *string                         `json:"link,omitempty"`
	MarketingMessageButtons *[]map[string]interface{}       `json:"marketing_message_buttons,omitempty"`
	Name                    *string                         `json:"name,omitempty"`
	Picture                 *string                         `json:"picture,omitempty"`
	PlaceData               *AdCreativePlaceData            `json:"place_data,omitempty"`
	StaticCard              *bool                           `json:"static_card,omitempty"`
	VideoID                 *core.ID                        `json:"video_id,omitempty"`
}
