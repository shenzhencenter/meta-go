package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type McomOnboardingStatus struct {
	OnboardingStatus *string  `json:"onboarding_status,omitempty"`
	PageID           *core.ID `json:"page_id,omitempty"`
}
