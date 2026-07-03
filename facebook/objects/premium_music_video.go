package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PremiumMusicVideo struct {
	CreationTime                    *string                   `json:"creation_time,omitempty"`
	CrossPostVideos                 *[]map[string]interface{} `json:"cross_post_videos,omitempty"`
	EligibleCrossPostPages          *[]map[string]interface{} `json:"eligible_cross_post_pages,omitempty"`
	ID                              *core.ID                  `json:"id,omitempty"`
	PreferredVideoThumbnailImageURI *string                   `json:"preferred_video_thumbnail_image_uri,omitempty"`
	PremiumMusicVideoMetadata       *map[string]interface{}   `json:"premium_music_video_metadata,omitempty"`
	ScheduledPublishTime            *int                      `json:"scheduled_publish_time,omitempty"`
	Title                           *string                   `json:"title,omitempty"`
}
