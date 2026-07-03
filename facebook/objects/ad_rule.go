package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdRule struct {
	AccountID        *core.ID              `json:"account_id,omitempty"`
	CreatedBy        *User                 `json:"created_by,omitempty"`
	CreatedTime      *time.Time            `json:"created_time,omitempty"`
	DisableErrorCode *int                  `json:"disable_error_code,omitempty"`
	EvaluationSpec   *AdRuleEvaluationSpec `json:"evaluation_spec,omitempty"`
	ExecutionSpec    *AdRuleExecutionSpec  `json:"execution_spec,omitempty"`
	ID               *core.ID              `json:"id,omitempty"`
	Name             *string               `json:"name,omitempty"`
	ScheduleSpec     *AdRuleScheduleSpec   `json:"schedule_spec,omitempty"`
	Status           *string               `json:"status,omitempty"`
	UpdatedTime      *time.Time            `json:"updated_time,omitempty"`
}
