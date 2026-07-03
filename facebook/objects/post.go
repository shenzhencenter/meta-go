package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type Post struct {
	Actions                      *[]interface{}          `json:"actions,omitempty"`
	AdminCreator                 *map[string]interface{} `json:"admin_creator,omitempty"`
	AllowedAdvertisingObjectives *[]string               `json:"allowed_advertising_objectives,omitempty"`
	Application                  *Application            `json:"application,omitempty"`
	BackdatedTime                *time.Time              `json:"backdated_time,omitempty"`
	CallToAction                 *map[string]interface{} `json:"call_to_action,omitempty"`
	CanReplyPrivately            *bool                   `json:"can_reply_privately,omitempty"`
	Caption                      *string                 `json:"caption,omitempty"`
	ChildAttachments             *[]interface{}          `json:"child_attachments,omitempty"`
	CommentsMirroringDomain      *string                 `json:"comments_mirroring_domain,omitempty"`
	Coordinates                  *map[string]interface{} `json:"coordinates,omitempty"`
	CreatedTime                  *time.Time              `json:"created_time,omitempty"`
	Description                  *string                 `json:"description,omitempty"`
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
	Link                         *string                 `json:"link,omitempty"`
	Message                      *string                 `json:"message,omitempty"`
	MessageTags                  *[]interface{}          `json:"message_tags,omitempty"`
	MultiShareEndCard            *bool                   `json:"multi_share_end_card,omitempty"`
	MultiShareOptimized          *bool                   `json:"multi_share_optimized,omitempty"`
	Name                         *string                 `json:"name,omitempty"`
	ObjectID                     *core.ID                `json:"object_id,omitempty"`
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
	Source                       *string                 `json:"source,omitempty"`
	StatusType                   *string                 `json:"status_type,omitempty"`
	Story                        *string                 `json:"story,omitempty"`
	StoryTags                    *[]interface{}          `json:"story_tags,omitempty"`
	Subscribed                   *bool                   `json:"subscribed,omitempty"`
	Target                       *Profile                `json:"target,omitempty"`
	Targeting                    *map[string]interface{} `json:"targeting,omitempty"`
	TimelineVisibility           *string                 `json:"timeline_visibility,omitempty"`
	Type                         *string                 `json:"type,omitempty"`
	UpdatedTime                  *time.Time              `json:"updated_time,omitempty"`
	Via                          *map[string]interface{} `json:"via,omitempty"`
	VideoBuyingEligibility       *[]string               `json:"video_buying_eligibility,omitempty"`
	WasFbLiveVideos              *bool                   `json:"was_fb_live_videos,omitempty"`
	Width                        *uint64                 `json:"width,omitempty"`
}
