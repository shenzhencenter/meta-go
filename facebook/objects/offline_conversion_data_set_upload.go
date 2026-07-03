package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OfflineConversionDataSetUpload struct {
	APICalls          *int     `json:"api_calls,omitempty"`
	CreationTime      *int     `json:"creation_time,omitempty"`
	DuplicateEntries  *int     `json:"duplicate_entries,omitempty"`
	EventStats        *string  `json:"event_stats,omitempty"`
	EventTimeMax      *int     `json:"event_time_max,omitempty"`
	EventTimeMin      *int     `json:"event_time_min,omitempty"`
	FirstUploadTime   *int     `json:"first_upload_time,omitempty"`
	ID                *core.ID `json:"id,omitempty"`
	IsExcludedForLift *bool    `json:"is_excluded_for_lift,omitempty"`
	LastUploadTime    *int     `json:"last_upload_time,omitempty"`
	MatchRateApprox   *int     `json:"match_rate_approx,omitempty"`
	MatchedEntries    *int     `json:"matched_entries,omitempty"`
	UploadTag         *string  `json:"upload_tag,omitempty"`
	ValidEntries      *int     `json:"valid_entries,omitempty"`
}

var OfflineConversionDataSetUploadFields = struct {
	APICalls          string
	CreationTime      string
	DuplicateEntries  string
	EventStats        string
	EventTimeMax      string
	EventTimeMin      string
	FirstUploadTime   string
	ID                string
	IsExcludedForLift string
	LastUploadTime    string
	MatchRateApprox   string
	MatchedEntries    string
	UploadTag         string
	ValidEntries      string
}{
	APICalls:          "api_calls",
	CreationTime:      "creation_time",
	DuplicateEntries:  "duplicate_entries",
	EventStats:        "event_stats",
	EventTimeMax:      "event_time_max",
	EventTimeMin:      "event_time_min",
	FirstUploadTime:   "first_upload_time",
	ID:                "id",
	IsExcludedForLift: "is_excluded_for_lift",
	LastUploadTime:    "last_upload_time",
	MatchRateApprox:   "match_rate_approx",
	MatchedEntries:    "matched_entries",
	UploadTag:         "upload_tag",
	ValidEntries:      "valid_entries",
}
