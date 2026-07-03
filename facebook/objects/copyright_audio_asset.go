package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CopyrightAudioAsset struct {
	AudioAvailabilityStatus *string                                           `json:"audio_availability_status,omitempty"`
	AudioLibraryPolicy      *[]map[string][]map[string]map[string]interface{} `json:"audio_library_policy,omitempty"`
	CreationTime            *core.Time                                        `json:"creation_time,omitempty"`
	ID                      *core.ID                                          `json:"id,omitempty"`
	ReferenceFileIsrcs      *[]string                                         `json:"reference_file_isrcs,omitempty"`
	ReferenceFiles          *[]map[string]interface{}                         `json:"reference_files,omitempty"`
	Title                   *string                                           `json:"title,omitempty"`
	UpdateTime              *core.Time                                        `json:"update_time,omitempty"`
}

var CopyrightAudioAssetFields = struct {
	AudioAvailabilityStatus string
	AudioLibraryPolicy      string
	CreationTime            string
	ID                      string
	ReferenceFileIsrcs      string
	ReferenceFiles          string
	Title                   string
	UpdateTime              string
}{
	AudioAvailabilityStatus: "audio_availability_status",
	AudioLibraryPolicy:      "audio_library_policy",
	CreationTime:            "creation_time",
	ID:                      "id",
	ReferenceFileIsrcs:      "reference_file_isrcs",
	ReferenceFiles:          "reference_files",
	Title:                   "title",
	UpdateTime:              "update_time",
}
