package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	Timestamp                *core.Time         `json:"timestamp,omitempty"`
	User                     *IGUser            `json:"user,omitempty"`
	Username                 *string            `json:"username,omitempty"`
}

var IGCommentFields = struct {
	From                     string
	Hidden                   string
	ID                       string
	LegacyInstagramCommentID string
	LikeCount                string
	Media                    string
	ParentID                 string
	Text                     string
	Timestamp                string
	User                     string
	Username                 string
}{
	From:                     "from",
	Hidden:                   "hidden",
	ID:                       "id",
	LegacyInstagramCommentID: "legacy_instagram_comment_id",
	LikeCount:                "like_count",
	Media:                    "media",
	ParentID:                 "parent_id",
	Text:                     "text",
	Timestamp:                "timestamp",
	User:                     "user",
	Username:                 "username",
}
