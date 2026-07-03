package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MessengerAdsPartialAutomatedStepList struct {
	FbleadForm          *LeadgenForm `json:"fblead_form,omitempty"`
	FirstStepID         *core.ID     `json:"first_step_id,omitempty"`
	ID                  *core.ID     `json:"id,omitempty"`
	Page                *Page        `json:"page,omitempty"`
	PrivacyURL          *string      `json:"privacy_url,omitempty"`
	ReminderText        *string      `json:"reminder_text,omitempty"`
	StopQuestionMessage *string      `json:"stop_question_message,omitempty"`
}

var MessengerAdsPartialAutomatedStepListFields = struct {
	FbleadForm          string
	FirstStepID         string
	ID                  string
	Page                string
	PrivacyURL          string
	ReminderText        string
	StopQuestionMessage string
}{
	FbleadForm:          "fblead_form",
	FirstStepID:         "first_step_id",
	ID:                  "id",
	Page:                "page",
	PrivacyURL:          "privacy_url",
	ReminderText:        "reminder_text",
	StopQuestionMessage: "stop_question_message",
}
