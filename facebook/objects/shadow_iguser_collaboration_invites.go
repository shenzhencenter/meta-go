package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ShadowIGUserCollaborationInvites struct {
	Caption            *string  `json:"caption,omitempty"`
	MediaID            *core.ID `json:"media_id,omitempty"`
	MediaOwnerUsername *string  `json:"media_owner_username,omitempty"`
	MediaURL           *string  `json:"media_url,omitempty"`
}
