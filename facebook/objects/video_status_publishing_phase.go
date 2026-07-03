package objects

import (
	"time"
)

type VideoStatusPublishingPhase struct {
	Errors        *[]VideoStatusError `json:"errors,omitempty"`
	PublishStatus *string             `json:"publish_status,omitempty"`
	PublishTime   *time.Time          `json:"publish_time,omitempty"`
	Status        *string             `json:"status,omitempty"`
}
