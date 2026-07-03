package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExternalEventSource struct {
	ID         *core.ID `json:"id,omitempty"`
	Name       *string  `json:"name,omitempty"`
	SourceType *string  `json:"source_type,omitempty"`
}
