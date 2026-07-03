package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PageCategory struct {
	APIEnum          *string         `json:"api_enum,omitempty"`
	FbPageCategories *[]PageCategory `json:"fb_page_categories,omitempty"`
	ID               *core.ID        `json:"id,omitempty"`
	Name             *string         `json:"name,omitempty"`
}
