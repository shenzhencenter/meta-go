package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type ShadowIGUserCTXPartnerAppWelcomeMessageFlow struct {
	CompatiblePlatforms *[]string  `json:"compatible_platforms,omitempty"`
	EligiblePlatforms   *[]string  `json:"eligible_platforms,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	IsIgOnlyFlow        *bool      `json:"is_ig_only_flow,omitempty"`
	IsUsedInAd          *bool      `json:"is_used_in_ad,omitempty"`
	LastUpdateTime      *time.Time `json:"last_update_time,omitempty"`
	Name                *string    `json:"name,omitempty"`
	WelcomeMessageFlow  *string    `json:"welcome_message_flow,omitempty"`
}
