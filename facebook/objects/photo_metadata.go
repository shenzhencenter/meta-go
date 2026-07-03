package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PhotoMetadata struct {
	CameraMake       *string    `json:"camera_make,omitempty"`
	CameraModel      *string    `json:"camera_model,omitempty"`
	DatetimeModified *time.Time `json:"datetime_modified,omitempty"`
	DatetimeTaken    *time.Time `json:"datetime_taken,omitempty"`
	Exposure         *string    `json:"exposure,omitempty"`
	FocalLength      *string    `json:"focal_length,omitempty"`
	Fstop            *string    `json:"fstop,omitempty"`
	IsoSpeed         *int       `json:"iso_speed,omitempty"`
	OfflineID        *core.ID   `json:"offline_id,omitempty"`
	Orientation      *string    `json:"orientation,omitempty"`
	OriginalHeight   *string    `json:"original_height,omitempty"`
	OriginalWidth    *string    `json:"original_width,omitempty"`
}
