package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PrivacyOption struct {
	Description         *string  `json:"description,omitempty"`
	IconSrc             *string  `json:"icon_src,omitempty"`
	ID                  *core.ID `json:"id,omitempty"`
	IsCurrentlySelected *bool    `json:"is_currently_selected,omitempty"`
	Type                *string  `json:"type,omitempty"`
	UserID              *core.ID `json:"user_id,omitempty"`
}
