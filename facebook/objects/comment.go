package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Comment struct {
	AdminCreator             *User                   `json:"admin_creator,omitempty"`
	Application              *Application            `json:"application,omitempty"`
	Attachment               *map[string]interface{} `json:"attachment,omitempty"`
	CanComment               *bool                   `json:"can_comment,omitempty"`
	CanHide                  *bool                   `json:"can_hide,omitempty"`
	CanLike                  *bool                   `json:"can_like,omitempty"`
	CanRemove                *bool                   `json:"can_remove,omitempty"`
	CanReplyPrivately        *bool                   `json:"can_reply_privately,omitempty"`
	CommentCount             *uint64                 `json:"comment_count,omitempty"`
	CreatedTime              *core.Time              `json:"created_time,omitempty"`
	From                     *map[string]interface{} `json:"from,omitempty"`
	ID                       *core.ID                `json:"id,omitempty"`
	IsHidden                 *bool                   `json:"is_hidden,omitempty"`
	IsPrivate                *bool                   `json:"is_private,omitempty"`
	LikeCount                *uint64                 `json:"like_count,omitempty"`
	LiveBroadcastTimestamp   *uint64                 `json:"live_broadcast_timestamp,omitempty"`
	Message                  *string                 `json:"message,omitempty"`
	MessageTags              *[]EntityAtTextRange    `json:"message_tags,omitempty"`
	Object                   *map[string]interface{} `json:"object,omitempty"`
	Parent                   *Comment                `json:"parent,omitempty"`
	PermalinkURL             *string                 `json:"permalink_url,omitempty"`
	PrivateReplyConversation *map[string]interface{} `json:"private_reply_conversation,omitempty"`
	UserLikes                *bool                   `json:"user_likes,omitempty"`
}

var CommentFields = struct {
	AdminCreator             string
	Application              string
	Attachment               string
	CanComment               string
	CanHide                  string
	CanLike                  string
	CanRemove                string
	CanReplyPrivately        string
	CommentCount             string
	CreatedTime              string
	From                     string
	ID                       string
	IsHidden                 string
	IsPrivate                string
	LikeCount                string
	LiveBroadcastTimestamp   string
	Message                  string
	MessageTags              string
	Object                   string
	Parent                   string
	PermalinkURL             string
	PrivateReplyConversation string
	UserLikes                string
}{
	AdminCreator:             "admin_creator",
	Application:              "application",
	Attachment:               "attachment",
	CanComment:               "can_comment",
	CanHide:                  "can_hide",
	CanLike:                  "can_like",
	CanRemove:                "can_remove",
	CanReplyPrivately:        "can_reply_privately",
	CommentCount:             "comment_count",
	CreatedTime:              "created_time",
	From:                     "from",
	ID:                       "id",
	IsHidden:                 "is_hidden",
	IsPrivate:                "is_private",
	LikeCount:                "like_count",
	LiveBroadcastTimestamp:   "live_broadcast_timestamp",
	Message:                  "message",
	MessageTags:              "message_tags",
	Object:                   "object",
	Parent:                   "parent",
	PermalinkURL:             "permalink_url",
	PrivateReplyConversation: "private_reply_conversation",
	UserLikes:                "user_likes",
}
