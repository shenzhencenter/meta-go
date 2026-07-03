package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EventSourceGroup struct {
	Business      *Business              `json:"business,omitempty"`
	EventSources  *[]ExternalEventSource `json:"event_sources,omitempty"`
	ID            *core.ID               `json:"id,omitempty"`
	Name          *string                `json:"name,omitempty"`
	OwnerBusiness *Business              `json:"owner_business,omitempty"`
}
