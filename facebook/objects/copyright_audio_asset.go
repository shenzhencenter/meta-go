package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CopyrightAudioAsset struct {
	AudioAvailabilityStatus *string                                           `json:"audio_availability_status,omitempty"`
	AudioLibraryPolicy      *[]map[string][]map[string]map[string]interface{} `json:"audio_library_policy,omitempty"`
	CreationTime            *time.Time                                        `json:"creation_time,omitempty"`
	ID                      *core.ID                                          `json:"id,omitempty"`
	ReferenceFileIsrcs      *[]string                                         `json:"reference_file_isrcs,omitempty"`
	ReferenceFiles          *[]map[string]interface{}                         `json:"reference_files,omitempty"`
	Title                   *string                                           `json:"title,omitempty"`
	UpdateTime              *time.Time                                        `json:"update_time,omitempty"`
}
