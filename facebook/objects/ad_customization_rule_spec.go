package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCustomizationRuleSpec struct {
	Caption           *string                    `json:"caption,omitempty"`
	CustomizationSpec *map[string]interface{}    `json:"customization_spec,omitempty"`
	Description       *string                    `json:"description,omitempty"`
	ImageHash         *string                    `json:"image_hash,omitempty"`
	Link              *string                    `json:"link,omitempty"`
	Message           *string                    `json:"message,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	Priority          *int                       `json:"priority,omitempty"`
	TemplateURLSpec   *AdCreativeTemplateURLSpec `json:"template_url_spec,omitempty"`
	VideoID           *core.ID                   `json:"video_id,omitempty"`
}

var AdCustomizationRuleSpecFields = struct {
	Caption           string
	CustomizationSpec string
	Description       string
	ImageHash         string
	Link              string
	Message           string
	Name              string
	Priority          string
	TemplateURLSpec   string
	VideoID           string
}{
	Caption:           "caption",
	CustomizationSpec: "customization_spec",
	Description:       "description",
	ImageHash:         "image_hash",
	Link:              "link",
	Message:           "message",
	Name:              "name",
	Priority:          "priority",
	TemplateURLSpec:   "template_url_spec",
	VideoID:           "video_id",
}
