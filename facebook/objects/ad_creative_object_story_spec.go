package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCreativeObjectStorySpec struct {
	InstagramUserID *core.ID                 `json:"instagram_user_id,omitempty"`
	LinkData        *AdCreativeLinkData      `json:"link_data,omitempty"`
	PageID          *core.ID                 `json:"page_id,omitempty"`
	PhotoData       *AdCreativePhotoData     `json:"photo_data,omitempty"`
	ProductData     *[]AdCreativeProductData `json:"product_data,omitempty"`
	TemplateData    *AdCreativeLinkData      `json:"template_data,omitempty"`
	TextData        *AdCreativeTextData      `json:"text_data,omitempty"`
	VideoData       *AdCreativeVideoData     `json:"video_data,omitempty"`
}
