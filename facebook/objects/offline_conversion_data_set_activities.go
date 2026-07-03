package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type OfflineConversionDataSetActivities struct {
	ActorID       *core.ID   `json:"actor_id,omitempty"`
	ActorName     *string    `json:"actor_name,omitempty"`
	AdaccountID   *core.ID   `json:"adaccount_id,omitempty"`
	AdaccountName *string    `json:"adaccount_name,omitempty"`
	EventTime     *time.Time `json:"event_time,omitempty"`
	EventType     *string    `json:"event_type,omitempty"`
	ExtraData     *string    `json:"extra_data,omitempty"`
	ObjectID      *core.ID   `json:"object_id,omitempty"`
	ObjectName    *string    `json:"object_name,omitempty"`
}
