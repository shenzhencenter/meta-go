package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AssignedUser struct {
	Business *Business `json:"business,omitempty"`
	ID       *core.ID  `json:"id,omitempty"`
	Name     *string   `json:"name,omitempty"`
	UserType *string   `json:"user_type,omitempty"`
}
