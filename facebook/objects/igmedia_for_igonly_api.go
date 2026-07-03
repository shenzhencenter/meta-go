package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type IGMediaForIGOnlyAPI struct {
	AltText                *string    `json:"alt_text,omitempty"`
	Caption                *string    `json:"caption,omitempty"`
	CommentsCount          *int       `json:"comments_count,omitempty"`
	CurrentLiveViewerCount *int       `json:"current_live_viewer_count,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	IsCommentEnabled       *bool      `json:"is_comment_enabled,omitempty"`
	IsSharedToFeed         *bool      `json:"is_shared_to_feed,omitempty"`
	LikeCount              *int       `json:"like_count,omitempty"`
	MediaProductType       *string    `json:"media_product_type,omitempty"`
	MediaType              *string    `json:"media_type,omitempty"`
	MediaURL               *string    `json:"media_url,omitempty"`
	Owner                  *User      `json:"owner,omitempty"`
	Permalink              *string    `json:"permalink,omitempty"`
	Shortcode              *string    `json:"shortcode,omitempty"`
	ThumbnailURL           *string    `json:"thumbnail_url,omitempty"`
	Timestamp              *time.Time `json:"timestamp,omitempty"`
	Username               *string    `json:"username,omitempty"`
}
