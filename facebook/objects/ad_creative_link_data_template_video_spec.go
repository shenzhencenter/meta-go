package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeLinkDataTemplateVideoSpec struct {
	CategorizationCriteria *string              `json:"categorization_criteria,omitempty"`
	Customization          *[]map[string]string `json:"customization,omitempty"`
	TemplateID             *core.ID             `json:"template_id,omitempty"`
}

var AdCreativeLinkDataTemplateVideoSpecFields = struct {
	CategorizationCriteria string
	Customization          string
	TemplateID             string
}{
	CategorizationCriteria: "categorization_criteria",
	Customization:          "customization",
	TemplateID:             "template_id",
}
