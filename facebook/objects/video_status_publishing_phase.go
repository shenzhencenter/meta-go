package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoStatusPublishingPhase struct {
	Errors        *[]VideoStatusError `json:"errors,omitempty"`
	PublishStatus *string             `json:"publish_status,omitempty"`
	PublishTime   *core.Time          `json:"publish_time,omitempty"`
	Status        *string             `json:"status,omitempty"`
}

var VideoStatusPublishingPhaseFields = struct {
	Errors        string
	PublishStatus string
	PublishTime   string
	Status        string
}{
	Errors:        "errors",
	PublishStatus: "publish_status",
	PublishTime:   "publish_time",
	Status:        "status",
}
