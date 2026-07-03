package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdAccountAdRulesHistory struct {
	EvaluationSpec   *AdRuleEvaluationSpec  `json:"evaluation_spec,omitempty"`
	ExceptionCode    *int                   `json:"exception_code,omitempty"`
	ExceptionMessage *string                `json:"exception_message,omitempty"`
	ExecutionSpec    *AdRuleExecutionSpec   `json:"execution_spec,omitempty"`
	IsManual         *bool                  `json:"is_manual,omitempty"`
	Results          *[]AdRuleHistoryResult `json:"results,omitempty"`
	RuleID           *core.ID               `json:"rule_id,omitempty"`
	ScheduleSpec     *AdRuleScheduleSpec    `json:"schedule_spec,omitempty"`
	Timestamp        *time.Time             `json:"timestamp,omitempty"`
}
