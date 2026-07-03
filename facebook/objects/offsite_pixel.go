package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OffsitePixel struct {
	Creator        *string    `json:"creator,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	JsPixel        *string    `json:"js_pixel,omitempty"`
	LastFiringTime *core.Time `json:"last_firing_time,omitempty"`
	Name           *string    `json:"name,omitempty"`
	Tag            *string    `json:"tag,omitempty"`
}

var OffsitePixelFields = struct {
	Creator        string
	ID             string
	JsPixel        string
	LastFiringTime string
	Name           string
	Tag            string
}{
	Creator:        "creator",
	ID:             "id",
	JsPixel:        "js_pixel",
	LastFiringTime: "last_firing_time",
	Name:           "name",
	Tag:            "tag",
}
