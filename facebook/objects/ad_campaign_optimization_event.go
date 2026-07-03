package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdCampaignOptimizationEvent struct {
	CustomConversionID *core.ID `json:"custom_conversion_id,omitempty"`
	EventSequence      *uint64  `json:"event_sequence,omitempty"`
	EventType          *string  `json:"event_type,omitempty"`
}
