package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MessengerBusinessTemplate struct {
	Category                 *string                   `json:"category,omitempty"`
	Components               *[]map[string]interface{} `json:"components,omitempty"`
	CreationTime             *int                      `json:"creation_time,omitempty"`
	ID                       *core.ID                  `json:"id,omitempty"`
	Language                 *string                   `json:"language,omitempty"`
	LanguageCount            *int                      `json:"language_count,omitempty"`
	LastUpdatedTime          *core.Time                `json:"last_updated_time,omitempty"`
	LibraryTemplateName      *string                   `json:"library_template_name,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	ParameterFormat          *string                   `json:"parameter_format,omitempty"`
	RejectedReason           *string                   `json:"rejected_reason,omitempty"`
	RejectionReasons         *map[string]interface{}   `json:"rejection_reasons,omitempty"`
	SpecificRejectionReasons *map[string]interface{}   `json:"specific_rejection_reasons,omitempty"`
	Status                   *string                   `json:"status,omitempty"`
}

var MessengerBusinessTemplateFields = struct {
	Category                 string
	Components               string
	CreationTime             string
	ID                       string
	Language                 string
	LanguageCount            string
	LastUpdatedTime          string
	LibraryTemplateName      string
	Name                     string
	ParameterFormat          string
	RejectedReason           string
	RejectionReasons         string
	SpecificRejectionReasons string
	Status                   string
}{
	Category:                 "category",
	Components:               "components",
	CreationTime:             "creation_time",
	ID:                       "id",
	Language:                 "language",
	LanguageCount:            "language_count",
	LastUpdatedTime:          "last_updated_time",
	LibraryTemplateName:      "library_template_name",
	Name:                     "name",
	ParameterFormat:          "parameter_format",
	RejectedReason:           "rejected_reason",
	RejectionReasons:         "rejection_reasons",
	SpecificRejectionReasons: "specific_rejection_reasons",
	Status:                   "status",
}
