package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGUserCollaborationInvites struct {
	Caption            *string  `json:"caption,omitempty"`
	MediaID            *core.ID `json:"media_id,omitempty"`
	MediaOwnerUsername *string  `json:"media_owner_username,omitempty"`
	MediaURL           *string  `json:"media_url,omitempty"`
}

var ShadowIGUserCollaborationInvitesFields = struct {
	Caption            string
	MediaID            string
	MediaOwnerUsername string
	MediaURL           string
}{
	Caption:            "caption",
	MediaID:            "media_id",
	MediaOwnerUsername: "media_owner_username",
	MediaURL:           "media_url",
}
