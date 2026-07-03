package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdPlacePageSet struct {
	AccountID     *core.ID  `json:"account_id,omitempty"`
	ID            *core.ID  `json:"id,omitempty"`
	LocationTypes *[]string `json:"location_types,omitempty"`
	Name          *string   `json:"name,omitempty"`
	PagesCount    *int      `json:"pages_count,omitempty"`
	ParentPage    *Page     `json:"parent_page,omitempty"`
}

var AdPlacePageSetFields = struct {
	AccountID     string
	ID            string
	LocationTypes string
	Name          string
	PagesCount    string
	ParentPage    string
}{
	AccountID:     "account_id",
	ID:            "id",
	LocationTypes: "location_types",
	Name:          "name",
	PagesCount:    "pages_count",
	ParentPage:    "parent_page",
}
