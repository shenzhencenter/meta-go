package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageCTSTopic struct {
	AppID      *core.ID `json:"app_id,omitempty"`
	Frequency  *string  `json:"frequency,omitempty"`
	ImageHash  *string  `json:"image_hash,omitempty"`
	ImageURL   *string  `json:"image_url,omitempty"`
	Subscriber *int     `json:"subscriber,omitempty"`
	Title      *string  `json:"title,omitempty"`
}
