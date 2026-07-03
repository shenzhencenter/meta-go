package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenDraftQuestion struct {
	ConditionalQuestionsChoices   *[]LeadGenConditionalQuestionsGroupChoices   `json:"conditional_questions_choices,omitempty"`
	ConditionalQuestionsGroupID   *core.ID                                     `json:"conditional_questions_group_id,omitempty"`
	DependentConditionalQuestions *[]LeadGenConditionalQuestionsGroupQuestions `json:"dependent_conditional_questions,omitempty"`
	InlineContext                 *string                                      `json:"inline_context,omitempty"`
	Key                           *string                                      `json:"key,omitempty"`
	Label                         *string                                      `json:"label,omitempty"`
	Options                       *[]LeadGenQuestionOption                     `json:"options,omitempty"`
	Type                          *string                                      `json:"type,omitempty"`
}

var LeadGenDraftQuestionFields = struct {
	ConditionalQuestionsChoices   string
	ConditionalQuestionsGroupID   string
	DependentConditionalQuestions string
	InlineContext                 string
	Key                           string
	Label                         string
	Options                       string
	Type                          string
}{
	ConditionalQuestionsChoices:   "conditional_questions_choices",
	ConditionalQuestionsGroupID:   "conditional_questions_group_id",
	DependentConditionalQuestions: "dependent_conditional_questions",
	InlineContext:                 "inline_context",
	Key:                           "key",
	Label:                         "label",
	Options:                       "options",
	Type:                          "type",
}
