package objects

type LeadGenConditionalQuestionsGroupQuestions struct {
	FieldKey  *string `json:"field_key,omitempty"`
	InputType *string `json:"input_type,omitempty"`
	Name      *string `json:"name,omitempty"`
}

var LeadGenConditionalQuestionsGroupQuestionsFields = struct {
	FieldKey  string
	InputType string
	Name      string
}{
	FieldKey:  "field_key",
	InputType: "input_type",
	Name:      "name",
}
