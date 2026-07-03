package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Stories struct {
	CreationTime *string  `json:"creation_time,omitempty"`
	MediaID      *core.ID `json:"media_id,omitempty"`
	MediaType    *string  `json:"media_type,omitempty"`
	PostID       *core.ID `json:"post_id,omitempty"`
	Status       *string  `json:"status,omitempty"`
	URL          *string  `json:"url,omitempty"`
}
