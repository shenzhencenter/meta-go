package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Canvas struct {
	BackgroundColor           *string                      `json:"background_color,omitempty"`
	BodyElements              *[]map[string]interface{}    `json:"body_elements,omitempty"`
	BusinessID                *core.ID                     `json:"business_id,omitempty"`
	CanvasLink                *string                      `json:"canvas_link,omitempty"`
	CollectionHeroImage       *Photo                       `json:"collection_hero_image,omitempty"`
	CollectionHeroVideo       *AdVideo                     `json:"collection_hero_video,omitempty"`
	CollectionThumbnails      *[]CanvasCollectionThumbnail `json:"collection_thumbnails,omitempty"`
	DynamicSetting            *CanvasDynamicSetting        `json:"dynamic_setting,omitempty"`
	ElementPayload            *string                      `json:"element_payload,omitempty"`
	FbBodyElements            *[]map[string]interface{}    `json:"fb_body_elements,omitempty"`
	HeroAssetFacebookPostID   *core.ID                     `json:"hero_asset_facebook_post_id,omitempty"`
	HeroAssetInstagramMediaID *core.ID                     `json:"hero_asset_instagram_media_id,omitempty"`
	ID                        *core.ID                     `json:"id,omitempty"`
	IsHidden                  *bool                        `json:"is_hidden,omitempty"`
	IsPublished               *bool                        `json:"is_published,omitempty"`
	LastEditor                *User                        `json:"last_editor,omitempty"`
	LinkedDocuments           *[]Canvas                    `json:"linked_documents,omitempty"`
	Name                      *string                      `json:"name,omitempty"`
	Owner                     *Page                        `json:"owner,omitempty"`
	PropertyList              *[]string                    `json:"property_list,omitempty"`
	SourceTemplate            *CanvasTemplate              `json:"source_template,omitempty"`
	StoreURL                  *string                      `json:"store_url,omitempty"`
	StyleList                 *[]string                    `json:"style_list,omitempty"`
	Tags                      *[]string                    `json:"tags,omitempty"`
	UiPropertyList            *[]string                    `json:"ui_property_list,omitempty"`
	UnusedBodyElements        *[]map[string]interface{}    `json:"unused_body_elements,omitempty"`
	UpdateTime                *int                         `json:"update_time,omitempty"`
	UseRetailerItemIds        *bool                        `json:"use_retailer_item_ids,omitempty"`
}
