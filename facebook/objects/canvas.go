package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CanvasFields = struct {
	BackgroundColor           string
	BodyElements              string
	BusinessID                string
	CanvasLink                string
	CollectionHeroImage       string
	CollectionHeroVideo       string
	CollectionThumbnails      string
	DynamicSetting            string
	ElementPayload            string
	FbBodyElements            string
	HeroAssetFacebookPostID   string
	HeroAssetInstagramMediaID string
	ID                        string
	IsHidden                  string
	IsPublished               string
	LastEditor                string
	LinkedDocuments           string
	Name                      string
	Owner                     string
	PropertyList              string
	SourceTemplate            string
	StoreURL                  string
	StyleList                 string
	Tags                      string
	UiPropertyList            string
	UnusedBodyElements        string
	UpdateTime                string
	UseRetailerItemIds        string
}{
	BackgroundColor:           "background_color",
	BodyElements:              "body_elements",
	BusinessID:                "business_id",
	CanvasLink:                "canvas_link",
	CollectionHeroImage:       "collection_hero_image",
	CollectionHeroVideo:       "collection_hero_video",
	CollectionThumbnails:      "collection_thumbnails",
	DynamicSetting:            "dynamic_setting",
	ElementPayload:            "element_payload",
	FbBodyElements:            "fb_body_elements",
	HeroAssetFacebookPostID:   "hero_asset_facebook_post_id",
	HeroAssetInstagramMediaID: "hero_asset_instagram_media_id",
	ID:                        "id",
	IsHidden:                  "is_hidden",
	IsPublished:               "is_published",
	LastEditor:                "last_editor",
	LinkedDocuments:           "linked_documents",
	Name:                      "name",
	Owner:                     "owner",
	PropertyList:              "property_list",
	SourceTemplate:            "source_template",
	StoreURL:                  "store_url",
	StyleList:                 "style_list",
	Tags:                      "tags",
	UiPropertyList:            "ui_property_list",
	UnusedBodyElements:        "unused_body_elements",
	UpdateTime:                "update_time",
	UseRetailerItemIds:        "use_retailer_item_ids",
}
