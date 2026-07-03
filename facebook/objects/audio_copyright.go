package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AudioCopyright struct {
	CreationTime          *time.Time                           `json:"creation_time,omitempty"`
	DisplayedMatchesCount *int                                 `json:"displayed_matches_count,omitempty"`
	ID                    *core.ID                             `json:"id,omitempty"`
	InConflict            *bool                                `json:"in_conflict,omitempty"`
	Isrc                  *string                              `json:"isrc,omitempty"`
	MatchRule             *VideoCopyrightRule                  `json:"match_rule,omitempty"`
	OwnershipCountries    *[]string                            `json:"ownership_countries,omitempty"`
	OwnershipDetails      *[]map[string]map[string]interface{} `json:"ownership_details,omitempty"`
	ReferenceFileStatus   *string                              `json:"reference_file_status,omitempty"`
	RidgeMonitoringStatus *string                              `json:"ridge_monitoring_status,omitempty"`
	Tags                  *[]string                            `json:"tags,omitempty"`
	UpdateTime            *time.Time                           `json:"update_time,omitempty"`
	WhitelistedFbUsers    *[]map[string]interface{}            `json:"whitelisted_fb_users,omitempty"`
	WhitelistedIgUsers    *[]string                            `json:"whitelisted_ig_users,omitempty"`
}
