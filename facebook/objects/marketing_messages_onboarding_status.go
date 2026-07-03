package objects

import (
	"time"
)

type MarketingMessagesOnboardingStatus struct {
	Status *string    `json:"status,omitempty"`
	Time   *time.Time `json:"time,omitempty"`
}
