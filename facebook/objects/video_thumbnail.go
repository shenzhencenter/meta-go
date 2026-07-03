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
