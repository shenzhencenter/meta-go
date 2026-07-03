package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsConversionGoal struct {
	AdAccountID                *core.ID `json:"ad_account_id,omitempty"`
	ConversionEventValueSource *string  `json:"conversion_event_value_source,omitempty"`
	Description                *string  `json:"description,omitempty"`
	GoalCreationMethod         *string  `json:"goal_creation_method,omitempty"`
	ID                         *core.ID `json:"id,omitempty"`
	Name                       *string  `json:"name,omitempty"`
	PerformanceGoal            *string  `json:"performance_goal,omitempty"`
	UpdateStatus               *string  `json:"update_status,omitempty"`
}
