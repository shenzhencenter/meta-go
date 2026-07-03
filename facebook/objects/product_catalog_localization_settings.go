package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalogLocalizationSettings struct {
	DefaultCountry  *string  `json:"default_country,omitempty"`
	DefaultLanguage *string  `json:"default_language,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
}

var ProductCatalogLocalizationSettingsFields = struct {
	DefaultCountry  string
	DefaultLanguage string
	ID              string
}{
	DefaultCountry:  "default_country",
	DefaultLanguage: "default_language",
	ID:              "id",
}
