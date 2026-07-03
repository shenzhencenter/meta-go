package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type ScheduledPost struct {
	AdminCreator            *User      `json:"admin_creator,omitempty"`
	CreationTime            *time.Time `json:"creation_time,omitempty"`
	FeedAudienceDescription *string    `json:"feed_audience_description,omitempty"`
	FeedTargeting           *Targeting `json:"feed_targeting,omitempty"`
	ID                      *core.ID   `json:"id,omitempty"`
	IsPostInGoodState       *bool      `json:"is_post_in_good_state,omitempty"`
	Message                 *string    `json:"message,omitempty"`
	ModifiedTime            *time.Time `json:"modified_time,omitempty"`
	OgActionSummary         *string    `json:"og_action_summary,omitempty"`
	PermalinkURL            *string    `json:"permalink_url,omitempty"`
	Place                   *Place     `json:"place,omitempty"`
	PrivacyDescription      *string    `json:"privacy_description,omitempty"`
	ScheduledFailureNotice  *string    `json:"scheduled_failure_notice,omitempty"`
	ScheduledPublishTime    *time.Time `json:"scheduled_publish_time,omitempty"`
	StoryToken              *string    `json:"story_token,omitempty"`
	Thumbnail               *string    `json:"thumbnail,omitempty"`
	VideoID                 *core.ID   `json:"video_id,omitempty"`
}
