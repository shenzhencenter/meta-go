package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	UpdatedTime        *core.Time              `json:"updated_time,omitempty"`
	Wallpaper          *string                 `json:"wallpaper,omitempty"`
}

var UnifiedThreadFields = struct {
	CanReply           string
	Folder             string
	FormerParticipants string
	ID                 string
	IsOwner            string
	IsSubscribed       string
	Link               string
	LinkedGroup        string
	MessageCount       string
	Name               string
	Participants       string
	ScopedThreadKey    string
	Senders            string
	Snippet            string
	UnreadCount        string
	UpdatedTime        string
	Wallpaper          string
}{
	CanReply:           "can_reply",
	Folder:             "folder",
	FormerParticipants: "former_participants",
	ID:                 "id",
	IsOwner:            "is_owner",
	IsSubscribed:       "is_subscribed",
	Link:               "link",
	LinkedGroup:        "linked_group",
	MessageCount:       "message_count",
	Name:               "name",
	Participants:       "participants",
	ScopedThreadKey:    "scoped_thread_key",
	Senders:            "senders",
	Snippet:            "snippet",
	UnreadCount:        "unread_count",
	UpdatedTime:        "updated_time",
	Wallpaper:          "wallpaper",
}
