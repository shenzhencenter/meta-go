package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	Timestamp        *core.Time             `json:"timestamp,omitempty"`
}

var AdAccountAdRulesHistoryFields = struct {
	EvaluationSpec   string
	ExceptionCode    string
	ExceptionMessage string
	ExecutionSpec    string
	IsManual         string
	Results          string
	RuleID           string
	ScheduleSpec     string
	Timestamp        string
}{
	EvaluationSpec:   "evaluation_spec",
	ExceptionCode:    "exception_code",
	ExceptionMessage: "exception_message",
	ExecutionSpec:    "execution_spec",
	IsManual:         "is_manual",
	Results:          "results",
	RuleID:           "rule_id",
	ScheduleSpec:     "schedule_spec",
	Timestamp:        "timestamp",
}
