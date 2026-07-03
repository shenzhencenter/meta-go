package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type StoreCatalogSettings struct {
	ID   *core.ID `json:"id,omitempty"`
	Page *Page    `json:"page,omitempty"`
}
