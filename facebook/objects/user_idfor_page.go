package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserIDForPage struct {
	ID   *core.ID `json:"id,omitempty"`
	Page *Page    `json:"page,omitempty"`
}

var UserIDForPageFields = struct {
	ID   string
	Page string
}{
	ID:   "id",
	Page: "page",
}
