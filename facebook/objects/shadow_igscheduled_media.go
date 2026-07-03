package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGScheduledMedia struct {
	Caption          *string  `json:"caption,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
	MediaType        *string  `json:"media_type,omitempty"`
	MediaURL         *string  `json:"media_url,omitempty"`
	PublishTimestamp *int     `json:"publish_timestamp,omitempty"`
	ThumbnailURL     *string  `json:"thumbnail_url,omitempty"`
}
