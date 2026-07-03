package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type UserIDForPage struct {
	ID   *core.ID `json:"id,omitempty"`
	Page *Page    `json:"page,omitempty"`
}
