package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type MessengerBusinessTemplate struct {
	Category                 *string                   `json:"category,omitempty"`
	Components               *[]map[string]interface{} `json:"components,omitempty"`
	CreationTime             *int                      `json:"creation_time,omitempty"`
	ID                       *core.ID                  `json:"id,omitempty"`
	Language                 *string                   `json:"language,omitempty"`
	LanguageCount            *int                      `json:"language_count,omitempty"`
	LastUpdatedTime          *time.Time                `json:"last_updated_time,omitempty"`
	LibraryTemplateName      *string                   `json:"library_template_name,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	ParameterFormat          *string                   `json:"parameter_format,omitempty"`
	RejectedReason           *string                   `json:"rejected_reason,omitempty"`
	RejectionReasons         *map[string]interface{}   `json:"rejection_reasons,omitempty"`
	SpecificRejectionReasons *map[string]interface{}   `json:"specific_rejection_reasons,omitempty"`
	Status                   *string                   `json:"status,omitempty"`
}
