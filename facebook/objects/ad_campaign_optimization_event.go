package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCampaignOptimizationEvent struct {
	CustomConversionID *core.ID `json:"custom_conversion_id,omitempty"`
	EventSequence      *uint64  `json:"event_sequence,omitempty"`
	EventType          *string  `json:"event_type,omitempty"`
}

var AdCampaignOptimizationEventFields = struct {
	CustomConversionID string
	EventSequence      string
	EventType          string
}{
	CustomConversionID: "custom_conversion_id",
	EventSequence:      "event_sequence",
	EventType:          "event_type",
}
