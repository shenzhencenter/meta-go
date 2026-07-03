package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var BroadTargetingCategoriesFields = struct {
	CategoryDescription    string
	ID                     string
	Name                   string
	ParentCategory         string
	Path                   string
	SizeLowerBound         string
	SizeUpperBound         string
	Source                 string
	Type                   string
	TypeName               string
	UntranslatedName       string
	UntranslatedParentName string
}{
	CategoryDescription:    "category_description",
	ID:                     "id",
	Name:                   "name",
	ParentCategory:         "parent_category",
	Path:                   "path",
	SizeLowerBound:         "size_lower_bound",
	SizeUpperBound:         "size_upper_bound",
	Source:                 "source",
	Type:                   "type",
	TypeName:               "type_name",
	UntranslatedName:       "untranslated_name",
	UntranslatedParentName: "untranslated_parent_name",
}
