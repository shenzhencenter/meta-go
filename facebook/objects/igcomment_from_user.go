package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGCommentFromUser struct {
	ID             *core.ID `json:"id,omitempty"`
	SelfIgScopedID *core.ID `json:"self_ig_scoped_id,omitempty"`
	Username       *string  `json:"username,omitempty"`
}

var IGCommentFromUserFields = struct {
	ID             string
	SelfIgScopedID string
	Username       string
}{
	ID:             "id",
	SelfIgScopedID: "self_ig_scoped_id",
	Username:       "username",
}
