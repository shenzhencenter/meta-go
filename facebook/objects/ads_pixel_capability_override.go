package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelCapabilityOverride struct {
	Capability    *string  `json:"capability,omitempty"`
	ID            *core.ID `json:"id,omitempty"`
	OverrideValue *string  `json:"override_value,omitempty"`
	Reason        *string  `json:"reason,omitempty"`
}

var AdsPixelCapabilityOverrideFields = struct {
	Capability    string
	ID            string
	OverrideValue string
	Reason        string
}{
	Capability:    "capability",
	ID:            "id",
	OverrideValue: "override_value",
	Reason:        "reason",
}
