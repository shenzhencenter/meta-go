package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenQuestion struct {
	ConditionalQuestionsChoices   *[]LeadGenConditionalQuestionsGroupChoices   `json:"conditional_questions_choices,omitempty"`
	ConditionalQuestionsGroupID   *core.ID                                     `json:"conditional_questions_group_id,omitempty"`
	DependentConditionalQuestions *[]LeadGenConditionalQuestionsGroupQuestions `json:"dependent_conditional_questions,omitempty"`
	ID                            *core.ID                                     `json:"id,omitempty"`
	InlineContext                 *string                                      `json:"inline_context,omitempty"`
	Key                           *string                                      `json:"key,omitempty"`
	Label                         *string                                      `json:"label,omitempty"`
	Options                       *[]LeadGenQuestionOption                     `json:"options,omitempty"`
	Type                          *string                                      `json:"type,omitempty"`
}
