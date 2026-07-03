package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CustomAudienceSessionFields = struct {
	EndTime           string
	NumInvalidEntries string
	NumMatched        string
	NumReceived       string
	Progress          string
	SessionID         string
	Stage             string
	StartTime         string
}{
	EndTime:           "end_time",
	NumInvalidEntries: "num_invalid_entries",
	NumMatched:        "num_matched",
	NumReceived:       "num_received",
	Progress:          "progress",
	SessionID:         "session_id",
	Stage:             "stage",
	StartTime:         "start_time",
}
