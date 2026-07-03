package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BroadTargetingCategories struct {
	CategoryDescription    *string   `json:"category_description,omitempty"`
	ID                     *core.ID  `json:"id,omitempty"`
	Name                   *string   `json:"name,omitempty"`
	ParentCategory         *string   `json:"parent_category,omitempty"`
	Path                   *[]string `json:"path,omitempty"`
	SizeLowerBound         *int      `json:"size_lower_bound,omitempty"`
	SizeUpperBound         *int      `json:"size_upper_bound,omitempty"`
	Source                 *string   `json:"source,omitempty"`
	Type                   *int      `json:"type,omitempty"`
	TypeName               *string   `json:"type_name,omitempty"`
	UntranslatedName       *string   `json:"untranslated_name,omitempty"`
	UntranslatedParentName *string   `json:"untranslated_parent_name,omitempty"`
}
