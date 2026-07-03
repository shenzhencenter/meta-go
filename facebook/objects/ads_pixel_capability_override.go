package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdsPixelCapabilityOverride struct {
	Capability    *string  `json:"capability,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
	OverrideValue *string  `json:"override_value,omitempty"`
	Reason        *string  `json:"reason,omitempty"`
}
