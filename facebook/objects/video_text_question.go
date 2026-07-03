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
