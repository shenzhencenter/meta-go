package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CPASParentCatalogSettings struct {
	AttributionWindows        *[]string `json:"attribution_windows,omitempty"`
	DefaultCurrency           *string   `json:"default_currency,omitempty"`
	DisableUseAsParentCatalog *bool     `json:"disable_use_as_parent_catalog,omitempty"`
	ID                        *core.ID  `json:"id,omitempty"`
}
