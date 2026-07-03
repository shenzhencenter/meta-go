package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdPlacePageSet struct {
	AccountID     *core.ID  `json:"account_id,omitempty"`
	ID            *core.ID  `json:"id,omitempty"`
	LocationTypes *[]string `json:"location_types,omitempty"`
	Name          *string   `json:"name,omitempty"`
	PagesCount    *int      `json:"pages_count,omitempty"`
	ParentPage    *Page     `json:"parent_page,omitempty"`
}
