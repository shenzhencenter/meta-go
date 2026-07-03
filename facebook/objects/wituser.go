package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WITUser struct {
	AccessToken *string  `json:"access_token,omitempty"`
	ID          *core.ID `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
}

var WITUserFields = struct {
	AccessToken string
	ID          string
	Name        string
}{
	AccessToken: "access_token",
	ID:          "id",
	Name:        "name",
}
