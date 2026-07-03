package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGMediaCollaborators struct {
	ID           *core.ID `json:"id,omitempty"`
	InviteStatus *string  `json:"invite_status,omitempty"`
	Username     *string  `json:"username,omitempty"`
}

var ShadowIGMediaCollaboratorsFields = struct {
	ID           string
	InviteStatus string
	Username     string
}{
	ID:           "id",
	InviteStatus: "invite_status",
	Username:     "username",
}
