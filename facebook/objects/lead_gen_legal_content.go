package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type LeadGenLegalContent struct {
	CustomDisclaimer *LeadGenCustomDisclaimer `json:"custom_disclaimer,omitempty"`
	ID               *core.ID                 `json:"id,omitempty"`
	PrivacyPolicy    *LeadGenPrivacyPolicy    `json:"privacy_policy,omitempty"`
}
