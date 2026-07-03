package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AudioSubLabel struct {
	ExpirationTimestamp *time.Time `json:"expiration_timestamp,omitempty"`
	FlaggedTimestamp    *time.Time `json:"flagged_timestamp,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	LabelName           *string    `json:"label_name,omitempty"`
	LastUpdateTimestamp *time.Time `json:"last_update_timestamp,omitempty"`
	NumAudioTracks      *int       `json:"num_audio_tracks,omitempty"`
	State               *string    `json:"state,omitempty"`
}
