package objects

type LeadGenConditionalQuestionsGroupChoices struct {
	CustomizedToken     *string                                    `json:"customized_token,omitempty"`
	NextQuestionChoices *[]LeadGenConditionalQuestionsGroupChoices `json:"next_question_choices,omitempty"`
	Value               *string                                    `json:"value,omitempty"`
}
