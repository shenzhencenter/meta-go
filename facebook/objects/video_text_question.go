package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoTextQuestion struct {
	ID               *core.ID `json:"id,omitempty"`
	QuestionTargetID *core.ID `json:"question_target_id,omitempty"`
	QuestionText     *string  `json:"question_text,omitempty"`
	Status           *string  `json:"status,omitempty"`
}

var VideoTextQuestionFields = struct {
	ID               string
	QuestionTargetID string
	QuestionText     string
	Status           string
}{
	ID:               "id",
	QuestionTargetID: "question_target_id",
	QuestionText:     "question_text",
	Status:           "status",
}
