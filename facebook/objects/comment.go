package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	CreatedTime              *time.Time              `json:"created_time,omitempty"`
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
