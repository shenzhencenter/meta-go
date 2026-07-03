package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Stories struct {
	CreationTime *string  `json:"creation_time,omitempty"`
	MediaID      *core.ID `json:"media_id,omitempty"`
	MediaType    *string  `json:"media_type,omitempty"`
	PostID       *core.ID `json:"post_id,omitempty"`
	Status       *string  `json:"status,omitempty"`
	URL          *string  `json:"url,omitempty"`
}

var StoriesFields = struct {
	CreationTime string
	MediaID      string
	MediaType    string
	PostID       string
	Status       string
	URL          string
}{
	CreationTime: "creation_time",
	MediaID:      "media_id",
	MediaType:    "media_type",
	PostID:       "post_id",
	Status:       "status",
	URL:          "url",
}
