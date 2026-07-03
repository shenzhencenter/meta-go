package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleExecutionSpec struct {
	ExecutionOptions *[]AdRuleExecutionOptions               `json:"execution_options,omitempty"`
	ExecutionType    *enums.AdruleexecutionspecExecutionType `json:"execution_type,omitempty"`
	IsOnceOff        *bool                                   `json:"is_once_off,omitempty"`
}

var AdRuleExecutionSpecFields = struct {
	ExecutionOptions string
	ExecutionType    string
	IsOnceOff        string
}{
	ExecutionOptions: "execution_options",
	ExecutionType:    "execution_type",
	IsOnceOff:        "is_once_off",
}
