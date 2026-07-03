package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SlicedEventSourceGroup struct {
	EventSourceGroup *EventSourceGroup `json:"event_source_group,omitempty"`
	Filter           *string           `json:"filter,omitempty"`
	ID               *core.ID          `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
}
