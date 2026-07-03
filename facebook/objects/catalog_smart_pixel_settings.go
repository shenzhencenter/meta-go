package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CatalogSmartPixelSettingsFields = struct {
	AllowedDomains            string
	AvailablePropertyFilters  string
	Catalog                   string
	CbbCustomOverrideFilters  string
	CbbDefaultFilter          string
	Defaults                  string
	Filters                   string
	ID                        string
	IsCbbEnabled              string
	IsCreateEnabled           string
	IsDeleteEnabled           string
	IsUpdateEnabled           string
	MicrodataFormatPrecedence string
	Pixel                     string
	PropertyFilter            string
	TrustedDomains            string
}{
	AllowedDomains:            "allowed_domains",
	AvailablePropertyFilters:  "available_property_filters",
	Catalog:                   "catalog",
	CbbCustomOverrideFilters:  "cbb_custom_override_filters",
	CbbDefaultFilter:          "cbb_default_filter",
	Defaults:                  "defaults",
	Filters:                   "filters",
	ID:                        "id",
	IsCbbEnabled:              "is_cbb_enabled",
	IsCreateEnabled:           "is_create_enabled",
	IsDeleteEnabled:           "is_delete_enabled",
	IsUpdateEnabled:           "is_update_enabled",
	MicrodataFormatPrecedence: "microdata_format_precedence",
	Pixel:                     "pixel",
	PropertyFilter:            "property_filter",
	TrustedDomains:            "trusted_domains",
}
