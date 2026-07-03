package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGUserCTXPartnerAppWelcomeMessageFlow struct {
	CompatiblePlatforms *[]string  `json:"compatible_platforms,omitempty"`
	EligiblePlatforms   *[]string  `json:"eligible_platforms,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	IsIgOnlyFlow        *bool      `json:"is_ig_only_flow,omitempty"`
	IsUsedInAd          *bool      `json:"is_used_in_ad,omitempty"`
	LastUpdateTime      *core.Time `json:"last_update_time,omitempty"`
	Name                *string    `json:"name,omitempty"`
	WelcomeMessageFlow  *string    `json:"welcome_message_flow,omitempty"`
}

var ShadowIGUserCTXPartnerAppWelcomeMessageFlowFields = struct {
	CompatiblePlatforms string
	EligiblePlatforms   string
	ID                  string
	IsIgOnlyFlow        string
	IsUsedInAd          string
	LastUpdateTime      string
	Name                string
	WelcomeMessageFlow  string
}{
	CompatiblePlatforms: "compatible_platforms",
	EligiblePlatforms:   "eligible_platforms",
	ID:                  "id",
	IsIgOnlyFlow:        "is_ig_only_flow",
	IsUsedInAd:          "is_used_in_ad",
	LastUpdateTime:      "last_update_time",
	Name:                "name",
	WelcomeMessageFlow:  "welcome_message_flow",
}
