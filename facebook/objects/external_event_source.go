package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ExternalEventSource struct {
	ID         *core.ID `json:"id,omitempty"`
	Name       *string  `json:"name,omitempty"`
	SourceType *string  `json:"source_type,omitempty"`
}
