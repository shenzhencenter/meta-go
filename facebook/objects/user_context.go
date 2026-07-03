package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type UserContext struct {
	ID *core.ID `json:"id,omitempty"`
}
