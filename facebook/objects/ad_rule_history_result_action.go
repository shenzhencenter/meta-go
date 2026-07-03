package objects

type AdRuleHistoryResultAction struct {
	Action   *string `json:"action,omitempty"`
	Field    *string `json:"field,omitempty"`
	NewValue *string `json:"new_value,omitempty"`
	OldValue *string `json:"old_value,omitempty"`
}

var AdRuleHistoryResultActionFields = struct {
	Action   string
	Field    string
	NewValue string
	OldValue string
}{
	Action:   "action",
	Field:    "field",
	NewValue: "new_value",
	OldValue: "old_value",
}
