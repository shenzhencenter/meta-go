package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessPageRequest struct {
	ID   *core.ID `json:"id,omitempty"`
	Page *Page    `json:"page,omitempty"`
}
