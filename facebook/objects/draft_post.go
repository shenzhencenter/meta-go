package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DraftPost struct {
	AdminCreator            *User      `json:"admin_creator,omitempty"`
	CreationTime            *core.Time `json:"creation_time,omitempty"`
	FeedAudienceDescription *string    `json:"feed_audience_description,omitempty"`
	FeedTargeting           *Targeting `json:"feed_targeting,omitempty"`
	ID                      *core.ID   `json:"id,omitempty"`
	IsPostInGoodState       *bool      `json:"is_post_in_good_state,omitempty"`
	Message                 *string    `json:"message,omitempty"`
	ModifiedTime            *core.Time `json:"modified_time,omitempty"`
	OgActionSummary         *string    `json:"og_action_summary,omitempty"`
	PermalinkURL            *string    `json:"permalink_url,omitempty"`
	Place                   *Place     `json:"place,omitempty"`
	PrivacyDescription      *string    `json:"privacy_description,omitempty"`
	ScheduledFailureNotice  *string    `json:"scheduled_failure_notice,omitempty"`
	ScheduledPublishTime    *core.Time `json:"scheduled_publish_time,omitempty"`
	StoryToken              *string    `json:"story_token,omitempty"`
	Thumbnail               *string    `json:"thumbnail,omitempty"`
	VideoID                 *core.ID   `json:"video_id,omitempty"`
}

var DraftPostFields = struct {
	AdminCreator            string
	CreationTime            string
	FeedAudienceDescription string
	FeedTargeting           string
	ID                      string
	IsPostInGoodState       string
	Message                 string
	ModifiedTime            string
	OgActionSummary         string
	PermalinkURL            string
	Place                   string
	PrivacyDescription      string
	ScheduledFailureNotice  string
	ScheduledPublishTime    string
	StoryToken              string
	Thumbnail               string
	VideoID                 string
}{
	AdminCreator:            "admin_creator",
	CreationTime:            "creation_time",
	FeedAudienceDescription: "feed_audience_description",
	FeedTargeting:           "feed_targeting",
	ID:                      "id",
	IsPostInGoodState:       "is_post_in_good_state",
	Message:                 "message",
	ModifiedTime:            "modified_time",
	OgActionSummary:         "og_action_summary",
	PermalinkURL:            "permalink_url",
	Place:                   "place",
	PrivacyDescription:      "privacy_description",
	ScheduledFailureNotice:  "scheduled_failure_notice",
	ScheduledPublishTime:    "scheduled_publish_time",
	StoryToken:              "story_token",
	Thumbnail:               "thumbnail",
	VideoID:                 "video_id",
}
