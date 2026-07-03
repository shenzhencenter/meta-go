package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PrivacyOption struct {
	Description         *string  `json:"description,omitempty"`
	IconSrc             *string  `json:"icon_src,omitempty"`
	ID                  *core.ID `json:"id,omitempty"`
	IsCurrentlySelected *bool    `json:"is_currently_selected,omitempty"`
	Type                *string  `json:"type,omitempty"`
	UserID              *core.ID `json:"user_id,omitempty"`
}

var PrivacyOptionFields = struct {
	Description         string
	IconSrc             string
	ID                  string
	IsCurrentlySelected string
	Type                string
	UserID              string
}{
	Description:         "description",
	IconSrc:             "icon_src",
	ID:                  "id",
	IsCurrentlySelected: "is_currently_selected",
	Type:                "type",
	UserID:              "user_id",
}
