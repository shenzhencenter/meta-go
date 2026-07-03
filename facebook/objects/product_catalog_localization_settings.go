package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ProductCatalogLocalizationSettings struct {
	DefaultCountry  *string  `json:"default_country,omitempty"`
	DefaultLanguage *string  `json:"default_language,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
}
