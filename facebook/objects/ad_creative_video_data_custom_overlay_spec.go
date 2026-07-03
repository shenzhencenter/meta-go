package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeVideoDataCustomOverlaySpec struct {
	BackgroundColor   *string                                                      `json:"background_color,omitempty"`
	BackgroundOpacity *enums.AdcreativevideodatacustomoverlayspecBackgroundOpacity `json:"background_opacity,omitempty"`
	Duration          *int                                                         `json:"duration,omitempty"`
	FloatWithMargin   *bool                                                        `json:"float_with_margin,omitempty"`
	FullWidth         *bool                                                        `json:"full_width,omitempty"`
	Option            *enums.AdcreativevideodatacustomoverlayspecOption            `json:"option,omitempty"`
	Position          *enums.AdcreativevideodatacustomoverlayspecPosition          `json:"position,omitempty"`
	Start             *int                                                         `json:"start,omitempty"`
	Template          *enums.AdcreativevideodatacustomoverlayspecTemplate          `json:"template,omitempty"`
	TextColor         *string                                                      `json:"text_color,omitempty"`
}

var AdCreativeVideoDataCustomOverlaySpecFields = struct {
	BackgroundColor   string
	BackgroundOpacity string
	Duration          string
	FloatWithMargin   string
	FullWidth         string
	Option            string
	Position          string
	Start             string
	Template          string
	TextColor         string
}{
	BackgroundColor:   "background_color",
	BackgroundOpacity: "background_opacity",
	Duration:          "duration",
	FloatWithMargin:   "float_with_margin",
	FullWidth:         "full_width",
	Option:            "option",
	Position:          "position",
	Start:             "start",
	Template:          "template",
	TextColor:         "text_color",
}
