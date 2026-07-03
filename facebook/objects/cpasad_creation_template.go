package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASAdCreationTemplate struct {
	Description      *string  `json:"description,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
	IsUnusedTemplate *bool    `json:"is_unused_template,omitempty"`
	Name             *string  `json:"name,omitempty"`
	OptimizationGoal *string  `json:"optimization_goal,omitempty"`
	TargetingType    *string  `json:"targeting_type,omitempty"`
	TemplateType     *string  `json:"template_type,omitempty"`
}

var CPASAdCreationTemplateFields = struct {
	Description      string
	ID               string
	IsUnusedTemplate string
	Name             string
	OptimizationGoal string
	TargetingType    string
	TemplateType     string
}{
	Description:      "description",
	ID:               "id",
	IsUnusedTemplate: "is_unused_template",
	Name:             "name",
	OptimizationGoal: "optimization_goal",
	TargetingType:    "targeting_type",
	TemplateType:     "template_type",
}
