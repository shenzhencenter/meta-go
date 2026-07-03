package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AudioSubLabel struct {
	ExpirationTimestamp *core.Time `json:"expiration_timestamp,omitempty"`
	FlaggedTimestamp    *core.Time `json:"flagged_timestamp,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	LabelName           *string    `json:"label_name,omitempty"`
	LastUpdateTimestamp *core.Time `json:"last_update_timestamp,omitempty"`
	NumAudioTracks      *int       `json:"num_audio_tracks,omitempty"`
	State               *string    `json:"state,omitempty"`
}

var AudioSubLabelFields = struct {
	ExpirationTimestamp string
	FlaggedTimestamp    string
	ID                  string
	LabelName           string
	LastUpdateTimestamp string
	NumAudioTracks      string
	State               string
}{
	ExpirationTimestamp: "expiration_timestamp",
	FlaggedTimestamp:    "flagged_timestamp",
	ID:                  "id",
	LabelName:           "label_name",
	LastUpdateTimestamp: "last_update_timestamp",
	NumAudioTracks:      "num_audio_tracks",
	State:               "state",
}
