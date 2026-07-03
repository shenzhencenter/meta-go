package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdRule struct {
	AccountID        *core.ID              `json:"account_id,omitempty"`
	CreatedBy        *User                 `json:"created_by,omitempty"`
	CreatedTime      *core.Time            `json:"created_time,omitempty"`
	DisableErrorCode *int                  `json:"disable_error_code,omitempty"`
	EvaluationSpec   *AdRuleEvaluationSpec `json:"evaluation_spec,omitempty"`
	ExecutionSpec    *AdRuleExecutionSpec  `json:"execution_spec,omitempty"`
	ID               *core.ID              `json:"id,omitempty"`
	Name             *string               `json:"name,omitempty"`
	ScheduleSpec     *AdRuleScheduleSpec   `json:"schedule_spec,omitempty"`
	Status           *string               `json:"status,omitempty"`
	UpdatedTime      *core.Time            `json:"updated_time,omitempty"`
}

var AdRuleFields = struct {
	AccountID        string
	CreatedBy        string
	CreatedTime      string
	DisableErrorCode string
	EvaluationSpec   string
	ExecutionSpec    string
	ID               string
	Name             string
	ScheduleSpec     string
	Status           string
	UpdatedTime      string
}{
	AccountID:        "account_id",
	CreatedBy:        "created_by",
	CreatedTime:      "created_time",
	DisableErrorCode: "disable_error_code",
	EvaluationSpec:   "evaluation_spec",
	ExecutionSpec:    "execution_spec",
	ID:               "id",
	Name:             "name",
	ScheduleSpec:     "schedule_spec",
	Status:           "status",
	UpdatedTime:      "updated_time",
}
