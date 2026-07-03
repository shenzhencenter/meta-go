package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdsNamingTemplate struct {
	APIFields         *[][]map[string][]map[string]string `json:"api_fields,omitempty"`
	APIVersion        *string                             `json:"api_version,omitempty"`
	FieldOrder        *[]string                           `json:"field_order,omitempty"`
	ID                *core.ID                            `json:"id,omitempty"`
	Level             *enums.AdsnamingtemplateLevel       `json:"level,omitempty"`
	Separator         *string                             `json:"separator,omitempty"`
	TemplateVersion   *string                             `json:"template_version,omitempty"`
	UserDefinedFields *[][]map[string][]string            `json:"user_defined_fields,omitempty"`
	ValueSeparator    *string                             `json:"value_separator,omitempty"`
}
