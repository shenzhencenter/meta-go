package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SavedMessageResponse struct {
	ID        *core.ID `json:"id,omitempty"`
	Image     *string  `json:"image,omitempty"`
	IsEnabled *bool    `json:"is_enabled,omitempty"`
	Message   *string  `json:"message,omitempty"`
	Title     *string  `json:"title,omitempty"`
}
