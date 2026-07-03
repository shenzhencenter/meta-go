package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type McomOnboardingStatus struct {
	OnboardingStatus *string  `json:"onboarding_status,omitempty"`
	PageID           *core.ID `json:"page_id,omitempty"`
}
