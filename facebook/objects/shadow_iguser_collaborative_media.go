package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGUserCollaborativeMedia struct {
	Caption            *string    `json:"caption,omitempty"`
	CommentsCount      *int       `json:"comments_count,omitempty"`
	ID                 *core.ID   `json:"id,omitempty"`
	LikeCount          *int       `json:"like_count,omitempty"`
	MediaProductType   *string    `json:"media_product_type,omitempty"`
	MediaType          *string    `json:"media_type,omitempty"`
	MediaURL           *string    `json:"media_url,omitempty"`
	Permalink          *string    `json:"permalink,omitempty"`
	RepostsCount       *int       `json:"reposts_count,omitempty"`
	SavedCount         *int       `json:"saved_count,omitempty"`
	SharesCount        *int       `json:"shares_count,omitempty"`
	ThumbnailURL       *string    `json:"thumbnail_url,omitempty"`
	Timestamp          *core.Time `json:"timestamp,omitempty"`
	TotalCommentsCount *int       `json:"total_comments_count,omitempty"`
	TotalLikeCount     *int       `json:"total_like_count,omitempty"`
	TotalViewsCount    *int       `json:"total_views_count,omitempty"`
	Username           *string    `json:"username,omitempty"`
}

var ShadowIGUserCollaborativeMediaFields = struct {
	Caption            string
	CommentsCount      string
	ID                 string
	LikeCount          string
	MediaProductType   string
	MediaType          string
	MediaURL           string
	Permalink          string
	RepostsCount       string
	SavedCount         string
	SharesCount        string
	ThumbnailURL       string
	Timestamp          string
	TotalCommentsCount string
	TotalLikeCount     string
	TotalViewsCount    string
	Username           string
}{
	Caption:            "caption",
	CommentsCount:      "comments_count",
	ID:                 "id",
	LikeCount:          "like_count",
	MediaProductType:   "media_product_type",
	MediaType:          "media_type",
	MediaURL:           "media_url",
	Permalink:          "permalink",
	RepostsCount:       "reposts_count",
	SavedCount:         "saved_count",
	SharesCount:        "shares_count",
	ThumbnailURL:       "thumbnail_url",
	Timestamp:          "timestamp",
	TotalCommentsCount: "total_comments_count",
	TotalLikeCount:     "total_like_count",
	TotalViewsCount:    "total_views_count",
	Username:           "username",
}
