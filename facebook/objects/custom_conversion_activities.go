package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomConversionActivities struct {
	AppID     *core.ID   `json:"app_id,omitempty"`
	Data      *string    `json:"data,omitempty"`
	EventType *string    `json:"event_type,omitempty"`
	Timestamp *core.Time `json:"timestamp,omitempty"`
}

var CustomConversionActivitiesFields = struct {
	AppID     string
	Data      string
	EventType string
	Timestamp string
}{
	AppID:     "app_id",
	Data:      "data",
	EventType: "event_type",
	Timestamp: "timestamp",
}
