package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PagePost struct {
	Actions                      *[]interface{}          `json:"actions,omitempty"`
	AdminCreator                 *map[string]interface{} `json:"admin_creator,omitempty"`
	AllowedAdvertisingObjectives *[]string               `json:"allowed_advertising_objectives,omitempty"`
	Application                  *Application            `json:"application,omitempty"`
	BackdatedTime                *core.Time              `json:"backdated_time,omitempty"`
	CallToAction                 *map[string]interface{} `json:"call_to_action,omitempty"`
	CanReplyPrivately            *bool                   `json:"can_reply_privately,omitempty"`
	ChildAttachments             *[]interface{}          `json:"child_attachments,omitempty"`
	CommentsMirroringDomain      *string                 `json:"comments_mirroring_domain,omitempty"`
	Coordinates                  *map[string]interface{} `json:"coordinates,omitempty"`
	CreatedTime                  *core.Time              `json:"created_time,omitempty"`
	Event                        *Event                  `json:"event,omitempty"`
	ExpandedHeight               *uint64                 `json:"expanded_height,omitempty"`
	ExpandedWidth                *uint64                 `json:"expanded_width,omitempty"`
	FeedTargeting                *map[string]interface{} `json:"feed_targeting,omitempty"`
	From                         *map[string]interface{} `json:"from,omitempty"`
	FullPicture                  *string                 `json:"full_picture,omitempty"`
	Height                       *uint64                 `json:"height,omitempty"`
	Icon                         *string                 `json:"icon,omitempty"`
	ID                           *core.ID                `json:"id,omitempty"`
	InstagramEligibility         *string                 `json:"instagram_eligibility,omitempty"`
	IsAppShare                   *bool                   `json:"is_app_share,omitempty"`
	IsEligibleForDarkPost        *bool                   `json:"is_eligible_for_dark_post,omitempty"`
	IsEligibleForPromotion       *bool                   `json:"is_eligible_for_promotion,omitempty"`
	IsExpired                    *bool                   `json:"is_expired,omitempty"`
	IsFbLiveVideos               *bool                   `json:"is_fb_live_videos,omitempty"`
	IsHidden                     *bool                   `json:"is_hidden,omitempty"`
	IsInlineCreated              *bool                   `json:"is_inline_created,omitempty"`
	IsInstagramEligible          *bool                   `json:"is_instagram_eligible,omitempty"`
	IsLiveClip                   *bool                   `json:"is_live_clip,omitempty"`
	IsPopular                    *bool                   `json:"is_popular,omitempty"`
	IsPublished                  *bool                   `json:"is_published,omitempty"`
	IsSpherical                  *bool                   `json:"is_spherical,omitempty"`
	Message                      *string                 `json:"message,omitempty"`
	MessageTags                  *[]interface{}          `json:"message_tags,omitempty"`
	MultiShareEndCard            *bool                   `json:"multi_share_end_card,omitempty"`
	MultiShareOptimized          *bool                   `json:"multi_share_optimized,omitempty"`
	ParentID                     *core.ID                `json:"parent_id,omitempty"`
	PermalinkURL                 *string                 `json:"permalink_url,omitempty"`
	Picture                      *string                 `json:"picture,omitempty"`
	Place                        *Place                  `json:"place,omitempty"`
	Privacy                      *Privacy                `json:"privacy,omitempty"`
	PromotableID                 *core.ID                `json:"promotable_id,omitempty"`
	PromotionStatus              *string                 `json:"promotion_status,omitempty"`
	Properties                   *[]interface{}          `json:"properties,omitempty"`
	ScheduledPublishTime         *float64                `json:"scheduled_publish_time,omitempty"`
	Shares                       *map[string]interface{} `json:"shares,omitempty"`
	StatusType                   *string                 `json:"status_type,omitempty"`
	Story                        *string                 `json:"story,omitempty"`
	StoryTags                    *[]interface{}          `json:"story_tags,omitempty"`
	Subscribed                   *bool                   `json:"subscribed,omitempty"`
	Target                       *Profile                `json:"target,omitempty"`
	Targeting                    *map[string]interface{} `json:"targeting,omitempty"`
	TimelineVisibility           *string                 `json:"timeline_visibility,omitempty"`
	UpdatedTime                  *core.Time              `json:"updated_time,omitempty"`
	Via                          *map[string]interface{} `json:"via,omitempty"`
	VideoBuyingEligibility       *[]string               `json:"video_buying_eligibility,omitempty"`
	WasFbLiveVideos              *bool                   `json:"was_fb_live_videos,omitempty"`
	Width                        *uint64                 `json:"width,omitempty"`
}

var PagePostFields = struct {
	Actions                      string
	AdminCreator                 string
	AllowedAdvertisingObjectives string
	Application                  string
	BackdatedTime                string
	CallToAction                 string
	CanReplyPrivately            string
	ChildAttachments             string
	CommentsMirroringDomain      string
	Coordinates                  string
	CreatedTime                  string
	Event                        string
	ExpandedHeight               string
	ExpandedWidth                string
	FeedTargeting                string
	From                         string
	FullPicture                  string
	Height                       string
	Icon                         string
	ID                           string
	InstagramEligibility         string
	IsAppShare                   string
	IsEligibleForDarkPost        string
	IsEligibleForPromotion       string
	IsExpired                    string
	IsFbLiveVideos               string
	IsHidden                     string
	IsInlineCreated              string
	IsInstagramEligible          string
	IsLiveClip                   string
	IsPopular                    string
	IsPublished                  string
	IsSpherical                  string
	Message                      string
	MessageTags                  string
	MultiShareEndCard            string
	MultiShareOptimized          string
	ParentID                     string
	PermalinkURL                 string
	Picture                      string
	Place                        string
	Privacy                      string
	PromotableID                 string
	PromotionStatus              string
	Properties                   string
	ScheduledPublishTime         string
	Shares                       string
	StatusType                   string
	Story                        string
	StoryTags                    string
	Subscribed                   string
	Target                       string
	Targeting                    string
	TimelineVisibility           string
	UpdatedTime                  string
	Via                          string
	VideoBuyingEligibility       string
	WasFbLiveVideos              string
	Width                        string
}{
	Actions:                      "actions",
	AdminCreator:                 "admin_creator",
	AllowedAdvertisingObjectives: "allowed_advertising_objectives",
	Application:                  "application",
	BackdatedTime:                "backdated_time",
	CallToAction:                 "call_to_action",
	CanReplyPrivately:            "can_reply_privately",
	ChildAttachments:             "child_attachments",
	CommentsMirroringDomain:      "comments_mirroring_domain",
	Coordinates:                  "coordinates",
	CreatedTime:                  "created_time",
	Event:                        "event",
	ExpandedHeight:               "expanded_height",
	ExpandedWidth:                "expanded_width",
	FeedTargeting:                "feed_targeting",
	From:                         "from",
	FullPicture:                  "full_picture",
	Height:                       "height",
	Icon:                         "icon",
	ID:                           "id",
	InstagramEligibility:         "instagram_eligibility",
	IsAppShare:                   "is_app_share",
	IsEligibleForDarkPost:        "is_eligible_for_dark_post",
	IsEligibleForPromotion:       "is_eligible_for_promotion",
	IsExpired:                    "is_expired",
	IsFbLiveVideos:               "is_fb_live_videos",
	IsHidden:                     "is_hidden",
	IsInlineCreated:              "is_inline_created",
	IsInstagramEligible:          "is_instagram_eligible",
	IsLiveClip:                   "is_live_clip",
	IsPopular:                    "is_popular",
	IsPublished:                  "is_published",
	IsSpherical:                  "is_spherical",
	Message:                      "message",
	MessageTags:                  "message_tags",
	MultiShareEndCard:            "multi_share_end_card",
	MultiShareOptimized:          "multi_share_optimized",
	ParentID:                     "parent_id",
	PermalinkURL:                 "permalink_url",
	Picture:                      "picture",
	Place:                        "place",
	Privacy:                      "privacy",
	PromotableID:                 "promotable_id",
	PromotionStatus:              "promotion_status",
	Properties:                   "properties",
	ScheduledPublishTime:         "scheduled_publish_time",
	Shares:                       "shares",
	StatusType:                   "status_type",
	Story:                        "story",
	StoryTags:                    "story_tags",
	Subscribed:                   "subscribed",
	Target:                       "target",
	Targeting:                    "targeting",
	TimelineVisibility:           "timeline_visibility",
	UpdatedTime:                  "updated_time",
	Via:                          "via",
	VideoBuyingEligibility:       "video_buying_eligibility",
	WasFbLiveVideos:              "was_fb_live_videos",
	Width:                        "width",
}
