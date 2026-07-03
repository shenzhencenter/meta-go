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

var PageCTSTopicFields = struct {
	AppID      string
	Frequency  string
	ImageHash  string
	ImageURL   string
	Subscriber string
	Title      string
}{
	AppID:      "app_id",
	Frequency:  "frequency",
	ImageHash:  "image_hash",
	ImageURL:   "image_url",
	Subscriber: "subscriber",
	Title:      "title",
}
