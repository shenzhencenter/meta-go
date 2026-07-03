package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DynamicARMetadata struct {
	AnchorPoint         *[]float64 `json:"anchor_point,omitempty"`
	ContainerEffectEnum *int       `json:"container_effect_enum,omitempty"`
	EffectIconURL       *string    `json:"effect_icon_url,omitempty"`
	EffectID            *core.ID   `json:"effect_id,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	Platforms           *[]string  `json:"platforms,omitempty"`
	ScaleFactor         *[]float64 `json:"scale_factor,omitempty"`
	ShadowTextureURL    *string    `json:"shadow_texture_url,omitempty"`
	SourceURL           *string    `json:"source_url,omitempty"`
	State               *string    `json:"state,omitempty"`
	Tags                *[]string  `json:"tags,omitempty"`
	VariantPickerURL    *string    `json:"variant_picker_url,omitempty"`
}

var DynamicARMetadataFields = struct {
	AnchorPoint         string
	ContainerEffectEnum string
	EffectIconURL       string
	EffectID            string
	ID                  string
	Platforms           string
	ScaleFactor         string
	ShadowTextureURL    string
	SourceURL           string
	State               string
	Tags                string
	VariantPickerURL    string
}{
	AnchorPoint:         "anchor_point",
	ContainerEffectEnum: "container_effect_enum",
	EffectIconURL:       "effect_icon_url",
	EffectID:            "effect_id",
	ID:                  "id",
	Platforms:           "platforms",
	ScaleFactor:         "scale_factor",
	ShadowTextureURL:    "shadow_texture_url",
	SourceURL:           "source_url",
	State:               "state",
	Tags:                "tags",
	VariantPickerURL:    "variant_picker_url",
}
