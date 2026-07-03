package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
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
