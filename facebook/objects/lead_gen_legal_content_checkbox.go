package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenLegalContentCheckbox struct {
	ID                 *core.ID `json:"id,omitempty"`
	IsCheckedByDefault *bool    `json:"is_checked_by_default,omitempty"`
	IsRequired         *bool    `json:"is_required,omitempty"`
	Key                *string  `json:"key,omitempty"`
	Text               *string  `json:"text,omitempty"`
}

var LeadGenLegalContentCheckboxFields = struct {
	ID                 string
	IsCheckedByDefault string
	IsRequired         string
	Key                string
	Text               string
}{
	ID:                 "id",
	IsCheckedByDefault: "is_checked_by_default",
	IsRequired:         "is_required",
	Key:                "key",
	Text:               "text",
}
