package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
