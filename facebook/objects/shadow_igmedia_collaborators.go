package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ShadowIGMediaCollaborators struct {
	ID           *core.ID `json:"id,omitempty"`
	InviteStatus *string  `json:"invite_status,omitempty"`
	Username     *string  `json:"username,omitempty"`
}
