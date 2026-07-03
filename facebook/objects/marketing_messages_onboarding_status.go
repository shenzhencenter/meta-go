package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MarketingMessagesOnboardingStatus struct {
	Status *string    `json:"status,omitempty"`
	Time   *core.Time `json:"time,omitempty"`
}

var MarketingMessagesOnboardingStatusFields = struct {
	Status string
	Time   string
}{
	Status: "status",
	Time:   "time",
}
