package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PhotoMetadata struct {
	CameraMake       *string    `json:"camera_make,omitempty"`
	CameraModel      *string    `json:"camera_model,omitempty"`
	DatetimeModified *core.Time `json:"datetime_modified,omitempty"`
	DatetimeTaken    *core.Time `json:"datetime_taken,omitempty"`
	Exposure         *string    `json:"exposure,omitempty"`
	FocalLength      *string    `json:"focal_length,omitempty"`
	Fstop            *string    `json:"fstop,omitempty"`
	IsoSpeed         *int       `json:"iso_speed,omitempty"`
	OfflineID        *core.ID   `json:"offline_id,omitempty"`
	Orientation      *string    `json:"orientation,omitempty"`
	OriginalHeight   *string    `json:"original_height,omitempty"`
	OriginalWidth    *string    `json:"original_width,omitempty"`
}

var PhotoMetadataFields = struct {
	CameraMake       string
	CameraModel      string
	DatetimeModified string
	DatetimeTaken    string
	Exposure         string
	FocalLength      string
	Fstop            string
	IsoSpeed         string
	OfflineID        string
	Orientation      string
	OriginalHeight   string
	OriginalWidth    string
}{
	CameraMake:       "camera_make",
	CameraModel:      "camera_model",
	DatetimeModified: "datetime_modified",
	DatetimeTaken:    "datetime_taken",
	Exposure:         "exposure",
	FocalLength:      "focal_length",
	Fstop:            "fstop",
	IsoSpeed:         "iso_speed",
	OfflineID:        "offline_id",
	Orientation:      "orientation",
	OriginalHeight:   "original_height",
	OriginalWidth:    "original_width",
}
