package objects

type AdRuleHistoryResultAction struct {
	Action   *string `json:"action,omitempty"`
	Field    *string `json:"field,omitempty"`
	NewValue *string `json:"new_value,omitempty"`
	OldValue *string `json:"old_value,omitempty"`
}
