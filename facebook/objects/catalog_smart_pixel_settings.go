package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CatalogSmartPixelSettings struct {
	AllowedDomains            *[]string                 `json:"allowed_domains,omitempty"`
	AvailablePropertyFilters  *[]string                 `json:"available_property_filters,omitempty"`
	Catalog                   *ProductCatalog           `json:"catalog,omitempty"`
	CbbCustomOverrideFilters  *[]map[string]interface{} `json:"cbb_custom_override_filters,omitempty"`
	CbbDefaultFilter          *[]map[string][]string    `json:"cbb_default_filter,omitempty"`
	Defaults                  *[]map[string]string      `json:"defaults,omitempty"`
	Filters                   *[]map[string][]string    `json:"filters,omitempty"`
	ID                        *core.ID                  `json:"id,omitempty"`
	IsCbbEnabled              *bool                     `json:"is_cbb_enabled,omitempty"`
	IsCreateEnabled           *bool                     `json:"is_create_enabled,omitempty"`
	IsDeleteEnabled           *bool                     `json:"is_delete_enabled,omitempty"`
	IsUpdateEnabled           *bool                     `json:"is_update_enabled,omitempty"`
	MicrodataFormatPrecedence *[]string                 `json:"microdata_format_precedence,omitempty"`
	Pixel                     *AdsPixel                 `json:"pixel,omitempty"`
	PropertyFilter            *[]string                 `json:"property_filter,omitempty"`
	TrustedDomains            *[]string                 `json:"trusted_domains,omitempty"`
}
