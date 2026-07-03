package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountDefaultDestination struct {
	DestinationID  *core.ID `json:"destination_id,omitempty"`
	DestinationURL *string  `json:"destination_url,omitempty"`
}
