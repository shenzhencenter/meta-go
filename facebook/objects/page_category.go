package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageCategory struct {
	APIEnum          *string         `json:"api_enum,omitempty"`
	FbPageCategories *[]PageCategory `json:"fb_page_categories,omitempty"`
	ID               *core.ID        `json:"id,omitempty"`
	Name             *string         `json:"name,omitempty"`
}

var PageCategoryFields = struct {
	APIEnum          string
	FbPageCategories string
	ID               string
	Name             string
}{
	APIEnum:          "api_enum",
	FbPageCategories: "fb_page_categories",
	ID:               "id",
	Name:             "name",
}
