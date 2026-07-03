package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var AdCreativeObjectStorySpecFields = struct {
	InstagramUserID string
	LinkData        string
	PageID          string
	PhotoData       string
	ProductData     string
	TemplateData    string
	TextData        string
	VideoData       string
}{
	InstagramUserID: "instagram_user_id",
	LinkData:        "link_data",
	PageID:          "page_id",
	PhotoData:       "photo_data",
	ProductData:     "product_data",
	TemplateData:    "template_data",
	TextData:        "text_data",
	VideoData:       "video_data",
}
