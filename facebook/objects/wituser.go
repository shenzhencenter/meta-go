package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WITUser struct {
	AccessToken *string  `json:"access_token,omitempty"`
	ID          *core.ID `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
}
