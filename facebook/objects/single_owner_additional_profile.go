package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SingleOwnerAdditionalProfile struct {
	ID       *core.ID `json:"id,omitempty"`
	Name     *string  `json:"name,omitempty"`
	UserName *string  `json:"user_name,omitempty"`
}

var SingleOwnerAdditionalProfileFields = struct {
	ID       string
	Name     string
	UserName string
}{
	ID:       "id",
	Name:     "name",
	UserName: "user_name",
}
