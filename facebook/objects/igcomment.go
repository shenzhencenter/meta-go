package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type IGComment struct {
	From                     *IGCommentFromUser `json:"from,omitempty"`
	Hidden                   *bool              `json:"hidden,omitempty"`
	ID                       *core.ID           `json:"id,omitempty"`
	LegacyInstagramCommentID *core.ID           `json:"legacy_instagram_comment_id,omitempty"`
	LikeCount                *int               `json:"like_count,omitempty"`
	Media                    *IGMedia           `json:"media,omitempty"`
	ParentID                 *core.ID           `json:"parent_id,omitempty"`
	Text                     *string            `json:"text,omitempty"`
	Timestamp                *time.Time         `json:"timestamp,omitempty"`
	User                     *IGUser            `json:"user,omitempty"`
	Username                 *string            `json:"username,omitempty"`
}
