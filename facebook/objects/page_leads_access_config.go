package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageLeadsAccessConfig struct {
	ID   *core.ID `json:"id,omitempty"`
	Page *Page    `json:"page,omitempty"`
}

var PageLeadsAccessConfigFields = struct {
	ID   string
	Page string
}{
	ID:   "id",
	Page: "page",
}
