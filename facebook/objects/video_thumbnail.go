package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoThumbnail struct {
	Height      *uint64  `json:"height,omitempty"`
	ID          *core.ID `json:"id,omitempty"`
	IsPreferred *bool    `json:"is_preferred,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Scale       *float64 `json:"scale,omitempty"`
	URI         *string  `json:"uri,omitempty"`
	Width       *uint64  `json:"width,omitempty"`
}

var VideoThumbnailFields = struct {
	Height      string
	ID          string
	IsPreferred string
	Name        string
	Scale       string
	URI         string
	Width       string
}{
	Height:      "height",
	ID:          "id",
	IsPreferred: "is_preferred",
	Name:        "name",
	Scale:       "scale",
	URI:         "uri",
	Width:       "width",
}
