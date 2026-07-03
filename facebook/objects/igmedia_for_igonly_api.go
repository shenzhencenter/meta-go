package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	Timestamp              *core.Time `json:"timestamp,omitempty"`
	Username               *string    `json:"username,omitempty"`
}

var IGMediaForIGOnlyAPIFields = struct {
	AltText                string
	Caption                string
	CommentsCount          string
	CurrentLiveViewerCount string
	ID                     string
	IsCommentEnabled       string
	IsSharedToFeed         string
	LikeCount              string
	MediaProductType       string
	MediaType              string
	MediaURL               string
	Owner                  string
	Permalink              string
	Shortcode              string
	ThumbnailURL           string
	Timestamp              string
	Username               string
}{
	AltText:                "alt_text",
	Caption:                "caption",
	CommentsCount:          "comments_count",
	CurrentLiveViewerCount: "current_live_viewer_count",
	ID:                     "id",
	IsCommentEnabled:       "is_comment_enabled",
	IsSharedToFeed:         "is_shared_to_feed",
	LikeCount:              "like_count",
	MediaProductType:       "media_product_type",
	MediaType:              "media_type",
	MediaURL:               "media_url",
	Owner:                  "owner",
	Permalink:              "permalink",
	Shortcode:              "shortcode",
	ThumbnailURL:           "thumbnail_url",
	Timestamp:              "timestamp",
	Username:               "username",
}
