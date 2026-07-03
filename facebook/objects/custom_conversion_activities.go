package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CustomConversionActivities struct {
	AppID     *core.ID   `json:"app_id,omitempty"`
	Data      *string    `json:"data,omitempty"`
	EventType *string    `json:"event_type,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}
