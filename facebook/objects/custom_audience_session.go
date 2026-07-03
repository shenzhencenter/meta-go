package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudienceSession struct {
	EndTime           *string  `json:"end_time,omitempty"`
	NumInvalidEntries *string  `json:"num_invalid_entries,omitempty"`
	NumMatched        *string  `json:"num_matched,omitempty"`
	NumReceived       *string  `json:"num_received,omitempty"`
	Progress          *string  `json:"progress,omitempty"`
	SessionID         *core.ID `json:"session_id,omitempty"`
	Stage             *string  `json:"stage,omitempty"`
	StartTime         *string  `json:"start_time,omitempty"`
}
