package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type AdActivity struct {
	ActorID             *core.ID                   `json:"actor_id,omitempty"`
	ActorName           *string                    `json:"actor_name,omitempty"`
	ApplicationID       *core.ID                   `json:"application_id,omitempty"`
	ApplicationName     *string                    `json:"application_name,omitempty"`
	DateTimeInTimezone  *string                    `json:"date_time_in_timezone,omitempty"`
	EventTime           *time.Time                 `json:"event_time,omitempty"`
	EventType           *enums.AdactivityEventType `json:"event_type,omitempty"`
	ExtraData           *string                    `json:"extra_data,omitempty"`
	ObjectID            *core.ID                   `json:"object_id,omitempty"`
	ObjectName          *string                    `json:"object_name,omitempty"`
	ObjectType          *string                    `json:"object_type,omitempty"`
	TranslatedEventType *string                    `json:"translated_event_type,omitempty"`
}
