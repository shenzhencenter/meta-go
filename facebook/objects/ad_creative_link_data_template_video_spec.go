package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeLinkDataTemplateVideoSpec struct {
	CategorizationCriteria *string              `json:"categorization_criteria,omitempty"`
	Customization          *[]map[string]string `json:"customization,omitempty"`
	TemplateID             *core.ID             `json:"template_id,omitempty"`
}
