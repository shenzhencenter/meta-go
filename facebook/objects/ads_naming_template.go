package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var AdsNamingTemplateFields = struct {
	APIFields         string
	APIVersion        string
	FieldOrder        string
	ID                string
	Level             string
	Separator         string
	TemplateVersion   string
	UserDefinedFields string
	ValueSeparator    string
}{
	APIFields:         "api_fields",
	APIVersion:        "api_version",
	FieldOrder:        "field_order",
	ID:                "id",
	Level:             "level",
	Separator:         "separator",
	TemplateVersion:   "template_version",
	UserDefinedFields: "user_defined_fields",
	ValueSeparator:    "value_separator",
}
