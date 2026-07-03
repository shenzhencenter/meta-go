package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CopyrightMediaMisuse struct {
	AudioSegments          *[]map[string]interface{} `json:"audio_segments,omitempty"`
	CreationTime           *core.Time                `json:"creation_time,omitempty"`
	DisabledAudioSegments  *[]map[string]interface{} `json:"disabled_audio_segments,omitempty"`
	DisabledVideoSegments  *[]map[string]interface{} `json:"disabled_video_segments,omitempty"`
	EntireFileIssue        *bool                     `json:"entire_file_issue,omitempty"`
	EntireFileIssueReasons *[]string                 `json:"entire_file_issue_reasons,omitempty"`
	ExpirationTime         *core.Time                `json:"expiration_time,omitempty"`
	ID                     *core.ID                  `json:"id,omitempty"`
	MediaAssetID           *core.ID                  `json:"media_asset_id,omitempty"`
	Reasons                *[]string                 `json:"reasons,omitempty"`
	RequestedAudioSegments *[]map[string]interface{} `json:"requested_audio_segments,omitempty"`
	RequestedVideoSegments *[]map[string]interface{} `json:"requested_video_segments,omitempty"`
	ResolutionType         *string                   `json:"resolution_type,omitempty"`
	Status                 *string                   `json:"status,omitempty"`
	UpdateTime             *core.Time                `json:"update_time,omitempty"`
	VideoCopyright         *VideoCopyright           `json:"video_copyright,omitempty"`
	VideoSegments          *[]map[string]interface{} `json:"video_segments,omitempty"`
}

var CopyrightMediaMisuseFields = struct {
	AudioSegments          string
	CreationTime           string
	DisabledAudioSegments  string
	DisabledVideoSegments  string
	EntireFileIssue        string
	EntireFileIssueReasons string
	ExpirationTime         string
	ID                     string
	MediaAssetID           string
	Reasons                string
	RequestedAudioSegments string
	RequestedVideoSegments string
	ResolutionType         string
	Status                 string
	UpdateTime             string
	VideoCopyright         string
	VideoSegments          string
}{
	AudioSegments:          "audio_segments",
	CreationTime:           "creation_time",
	DisabledAudioSegments:  "disabled_audio_segments",
	DisabledVideoSegments:  "disabled_video_segments",
	EntireFileIssue:        "entire_file_issue",
	EntireFileIssueReasons: "entire_file_issue_reasons",
	ExpirationTime:         "expiration_time",
	ID:                     "id",
	MediaAssetID:           "media_asset_id",
	Reasons:                "reasons",
	RequestedAudioSegments: "requested_audio_segments",
	RequestedVideoSegments: "requested_video_segments",
	ResolutionType:         "resolution_type",
	Status:                 "status",
	UpdateTime:             "update_time",
	VideoCopyright:         "video_copyright",
	VideoSegments:          "video_segments",
}
