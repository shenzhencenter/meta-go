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
