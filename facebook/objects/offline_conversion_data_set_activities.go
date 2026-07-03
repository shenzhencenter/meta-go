package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OfflineConversionDataSetActivities struct {
	ActorID       *core.ID   `json:"actor_id,omitempty"`
	ActorName     *string    `json:"actor_name,omitempty"`
	AdaccountID   *core.ID   `json:"adaccount_id,omitempty"`
	AdaccountName *string    `json:"adaccount_name,omitempty"`
	EventTime     *core.Time `json:"event_time,omitempty"`
	EventType     *string    `json:"event_type,omitempty"`
	ExtraData     *string    `json:"extra_data,omitempty"`
	ObjectID      *core.ID   `json:"object_id,omitempty"`
	ObjectName    *string    `json:"object_name,omitempty"`
}

var OfflineConversionDataSetActivitiesFields = struct {
	ActorID       string
	ActorName     string
	AdaccountID   string
	AdaccountName string
	EventTime     string
	EventType     string
	ExtraData     string
	ObjectID      string
	ObjectName    string
}{
	ActorID:       "actor_id",
	ActorName:     "actor_name",
	AdaccountID:   "adaccount_id",
	AdaccountName: "adaccount_name",
	EventTime:     "event_time",
	EventType:     "event_type",
	ExtraData:     "extra_data",
	ObjectID:      "object_id",
	ObjectName:    "object_name",
}
