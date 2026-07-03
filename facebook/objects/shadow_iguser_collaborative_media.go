package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	Timestamp          *time.Time `json:"timestamp,omitempty"`
	TotalCommentsCount *int       `json:"total_comments_count,omitempty"`
	TotalLikeCount     *int       `json:"total_like_count,omitempty"`
	TotalViewsCount    *int       `json:"total_views_count,omitempty"`
	Username           *string    `json:"username,omitempty"`
}
