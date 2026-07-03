package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type UnifiedThread struct {
	CanReply           *bool                   `json:"can_reply,omitempty"`
	Folder             *string                 `json:"folder,omitempty"`
	FormerParticipants *map[string]interface{} `json:"former_participants,omitempty"`
	ID                 *core.ID                `json:"id,omitempty"`
	IsOwner            *bool                   `json:"is_owner,omitempty"`
	IsSubscribed       *bool                   `json:"is_subscribed,omitempty"`
	Link               *string                 `json:"link,omitempty"`
	LinkedGroup        *Group                  `json:"linked_group,omitempty"`
	MessageCount       *int                    `json:"message_count,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Participants       *map[string]interface{} `json:"participants,omitempty"`
	ScopedThreadKey    *string                 `json:"scoped_thread_key,omitempty"`
	Senders            *map[string]interface{} `json:"senders,omitempty"`
	Snippet            *string                 `json:"snippet,omitempty"`
	UnreadCount        *int                    `json:"unread_count,omitempty"`
	UpdatedTime        *time.Time              `json:"updated_time,omitempty"`
	Wallpaper          *string                 `json:"wallpaper,omitempty"`
}
