package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
