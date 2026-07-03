package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASParentCatalogSettings struct {
	AttributionWindows        *[]string `json:"attribution_windows,omitempty"`
	DefaultCurrency           *string   `json:"default_currency,omitempty"`
	DisableUseAsParentCatalog *bool     `json:"disable_use_as_parent_catalog,omitempty"`
	ID                        *core.ID  `json:"id,omitempty"`
}

var CPASParentCatalogSettingsFields = struct {
	AttributionWindows        string
	DefaultCurrency           string
	DisableUseAsParentCatalog string
	ID                        string
}{
	AttributionWindows:        "attribution_windows",
	DefaultCurrency:           "default_currency",
	DisableUseAsParentCatalog: "disable_use_as_parent_catalog",
	ID:                        "id",
}
