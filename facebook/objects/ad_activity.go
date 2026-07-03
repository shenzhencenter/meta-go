package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdActivity struct {
	ActorID             *core.ID                   `json:"actor_id,omitempty"`
	ActorName           *string                    `json:"actor_name,omitempty"`
	ApplicationID       *core.ID                   `json:"application_id,omitempty"`
	ApplicationName     *string                    `json:"application_name,omitempty"`
	DateTimeInTimezone  *string                    `json:"date_time_in_timezone,omitempty"`
	EventTime           *core.Time                 `json:"event_time,omitempty"`
	EventType           *enums.AdactivityEventType `json:"event_type,omitempty"`
	ExtraData           *string                    `json:"extra_data,omitempty"`
	ObjectID            *core.ID                   `json:"object_id,omitempty"`
	ObjectName          *string                    `json:"object_name,omitempty"`
	ObjectType          *string                    `json:"object_type,omitempty"`
	TranslatedEventType *string                    `json:"translated_event_type,omitempty"`
}

var AdActivityFields = struct {
	ActorID             string
	ActorName           string
	ApplicationID       string
	ApplicationName     string
	DateTimeInTimezone  string
	EventTime           string
	EventType           string
	ExtraData           string
	ObjectID            string
	ObjectName          string
	ObjectType          string
	TranslatedEventType string
}{
	ActorID:             "actor_id",
	ActorName:           "actor_name",
	ApplicationID:       "application_id",
	ApplicationName:     "application_name",
	DateTimeInTimezone:  "date_time_in_timezone",
	EventTime:           "event_time",
	EventType:           "event_type",
	ExtraData:           "extra_data",
	ObjectID:            "object_id",
	ObjectName:          "object_name",
	ObjectType:          "object_type",
	TranslatedEventType: "translated_event_type",
}
