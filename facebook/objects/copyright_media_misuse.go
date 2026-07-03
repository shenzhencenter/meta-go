package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type CopyrightMediaMisuse struct {
	AudioSegments          *[]map[string]interface{} `json:"audio_segments,omitempty"`
	CreationTime           *time.Time                `json:"creation_time,omitempty"`
	DisabledAudioSegments  *[]map[string]interface{} `json:"disabled_audio_segments,omitempty"`
	DisabledVideoSegments  *[]map[string]interface{} `json:"disabled_video_segments,omitempty"`
	EntireFileIssue        *bool                     `json:"entire_file_issue,omitempty"`
	EntireFileIssueReasons *[]string                 `json:"entire_file_issue_reasons,omitempty"`
	ExpirationTime         *time.Time                `json:"expiration_time,omitempty"`
	ID                     *core.ID                  `json:"id,omitempty"`
	MediaAssetID           *core.ID                  `json:"media_asset_id,omitempty"`
	Reasons                *[]string                 `json:"reasons,omitempty"`
	RequestedAudioSegments *[]map[string]interface{} `json:"requested_audio_segments,omitempty"`
	RequestedVideoSegments *[]map[string]interface{} `json:"requested_video_segments,omitempty"`
	ResolutionType         *string                   `json:"resolution_type,omitempty"`
	Status                 *string                   `json:"status,omitempty"`
	UpdateTime             *time.Time                `json:"update_time,omitempty"`
	VideoCopyright         *VideoCopyright           `json:"video_copyright,omitempty"`
	VideoSegments          *[]map[string]interface{} `json:"video_segments,omitempty"`
}
