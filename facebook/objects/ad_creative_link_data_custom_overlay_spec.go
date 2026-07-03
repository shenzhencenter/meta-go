package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeLinkDataCustomOverlaySpec struct {
	BackgroundColor *enums.AdcreativelinkdatacustomoverlayspecBackgroundColor `json:"background_color,omitempty"`
	FloatWithMargin *bool                                                     `json:"float_with_margin,omitempty"`
	Font            *enums.AdcreativelinkdatacustomoverlayspecFont            `json:"font,omitempty"`
	Option          *enums.AdcreativelinkdatacustomoverlayspecOption          `json:"option,omitempty"`
	Position        *enums.AdcreativelinkdatacustomoverlayspecPosition        `json:"position,omitempty"`
	RenderWithIcon  *bool                                                     `json:"render_with_icon,omitempty"`
	Template        *enums.AdcreativelinkdatacustomoverlayspecTemplate        `json:"template,omitempty"`
	TextColor       *enums.AdcreativelinkdatacustomoverlayspecTextColor       `json:"text_color,omitempty"`
}

var AdCreativeLinkDataCustomOverlaySpecFields = struct {
	BackgroundColor string
	FloatWithMargin string
	Font            string
	Option          string
	Position        string
	RenderWithIcon  string
	Template        string
	TextColor       string
}{
	BackgroundColor: "background_color",
	FloatWithMargin: "float_with_margin",
	Font:            "font",
	Option:          "option",
	Position:        "position",
	RenderWithIcon:  "render_with_icon",
	Template:        "template",
	TextColor:       "text_color",
}
