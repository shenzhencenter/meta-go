package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SingleOwnerAdditionalProfile struct {
	ID       *core.ID `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`
	UserName *string  `json:"user_name,omitempty"`
}
