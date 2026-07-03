package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeLinkDataImageLayerSpec struct {
	BlendingMode              *enums.AdcreativelinkdataimagelayerspecBlendingMode    `json:"blending_mode,omitempty"`
	Content                   *map[string]interface{}                                `json:"content,omitempty"`
	FrameAutoShowEnrollStatus *string                                                `json:"frame_auto_show_enroll_status,omitempty"`
	FrameImageHash            *string                                                `json:"frame_image_hash,omitempty"`
	FrameSource               *enums.AdcreativelinkdataimagelayerspecFrameSource     `json:"frame_source,omitempty"`
	ImageSource               *enums.AdcreativelinkdataimagelayerspecImageSource     `json:"image_source,omitempty"`
	LayerType                 *enums.AdcreativelinkdataimagelayerspecLayerType       `json:"layer_type,omitempty"`
	Opacity                   *int                                                   `json:"opacity,omitempty"`
	OverlayPosition           *enums.AdcreativelinkdataimagelayerspecOverlayPosition `json:"overlay_position,omitempty"`
	OverlayShape              *enums.AdcreativelinkdataimagelayerspecOverlayShape    `json:"overlay_shape,omitempty"`
	Scale                     *int                                                   `json:"scale,omitempty"`
	ShapeColor                *string                                                `json:"shape_color,omitempty"`
	TextColor                 *string                                                `json:"text_color,omitempty"`
	TextFont                  *enums.AdcreativelinkdataimagelayerspecTextFont        `json:"text_font,omitempty"`
}

var AdCreativeLinkDataImageLayerSpecFields = struct {
	BlendingMode              string
	Content                   string
	FrameAutoShowEnrollStatus string
	FrameImageHash            string
	FrameSource               string
	ImageSource               string
	LayerType                 string
	Opacity                   string
	OverlayPosition           string
	OverlayShape              string
	Scale                     string
	ShapeColor                string
	TextColor                 string
	TextFont                  string
}{
	BlendingMode:              "blending_mode",
	Content:                   "content",
	FrameAutoShowEnrollStatus: "frame_auto_show_enroll_status",
	FrameImageHash:            "frame_image_hash",
	FrameSource:               "frame_source",
	ImageSource:               "image_source",
	LayerType:                 "layer_type",
	Opacity:                   "opacity",
	OverlayPosition:           "overlay_position",
	OverlayShape:              "overlay_shape",
	Scale:                     "scale",
	ShapeColor:                "shape_color",
	TextColor:                 "text_color",
	TextFont:                  "text_font",
}
