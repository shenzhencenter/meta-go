package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdRuleExecutionSpec struct {
	ExecutionOptions *[]AdRuleExecutionOptions               `json:"execution_options,omitempty"`
	ExecutionType    *enums.AdruleexecutionspecExecutionType `json:"execution_type,omitempty"`
	IsOnceOff        *bool                                   `json:"is_once_off,omitempty"`
}
