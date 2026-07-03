package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type LeadGenLegalContentCheckbox struct {
	ID                 *core.ID `json:"id,omitempty"`
	IsCheckedByDefault *bool    `json:"is_checked_by_default,omitempty"`
	IsRequired         *bool    `json:"is_required,omitempty"`
	Key                *string  `json:"key,omitempty"`
	Text               *string  `json:"text,omitempty"`
}
