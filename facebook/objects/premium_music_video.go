package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var PremiumMusicVideoFields = struct {
	CreationTime                    string
	CrossPostVideos                 string
	EligibleCrossPostPages          string
	ID                              string
	PreferredVideoThumbnailImageURI string
	PremiumMusicVideoMetadata       string
	ScheduledPublishTime            string
	Title                           string
}{
	CreationTime:                    "creation_time",
	CrossPostVideos:                 "cross_post_videos",
	EligibleCrossPostPages:          "eligible_cross_post_pages",
	ID:                              "id",
	PreferredVideoThumbnailImageURI: "preferred_video_thumbnail_image_uri",
	PremiumMusicVideoMetadata:       "premium_music_video_metadata",
	ScheduledPublishTime:            "scheduled_publish_time",
	Title:                           "title",
}
