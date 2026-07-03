package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
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
