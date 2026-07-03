package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeLinkDataImageOverlaySpec struct {
	CustomTextType   *enums.AdcreativelinkdataimageoverlayspecCustomTextType  `json:"custom_text_type,omitempty"`
	FloatWithMargin  *bool                                                    `json:"float_with_margin,omitempty"`
	OverlayTemplate  *enums.AdcreativelinkdataimageoverlayspecOverlayTemplate `json:"overlay_template,omitempty"`
	Position         *enums.AdcreativelinkdataimageoverlayspecPosition        `json:"position,omitempty"`
	TextFont         *enums.AdcreativelinkdataimageoverlayspecTextFont        `json:"text_font,omitempty"`
	TextTemplateTags *[]string                                                `json:"text_template_tags,omitempty"`
	TextType         *enums.AdcreativelinkdataimageoverlayspecTextType        `json:"text_type,omitempty"`
	ThemeColor       *enums.AdcreativelinkdataimageoverlayspecThemeColor      `json:"theme_color,omitempty"`
}

var AdCreativeLinkDataImageOverlaySpecFields = struct {
	CustomTextType   string
	FloatWithMargin  string
	OverlayTemplate  string
	Position         string
	TextFont         string
	TextTemplateTags string
	TextType         string
	ThemeColor       string
}{
	CustomTextType:   "custom_text_type",
	FloatWithMargin:  "float_with_margin",
	OverlayTemplate:  "overlay_template",
	Position:         "position",
	TextFont:         "text_font",
	TextTemplateTags: "text_template_tags",
	TextType:         "text_type",
	ThemeColor:       "theme_color",
}
