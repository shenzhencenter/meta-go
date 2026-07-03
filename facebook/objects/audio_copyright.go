package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AudioCopyright struct {
	CreationTime          *core.Time                           `json:"creation_time,omitempty"`
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
	UpdateTime            *core.Time                           `json:"update_time,omitempty"`
	WhitelistedFbUsers    *[]map[string]interface{}            `json:"whitelisted_fb_users,omitempty"`
	WhitelistedIgUsers    *[]string                            `json:"whitelisted_ig_users,omitempty"`
}

var AudioCopyrightFields = struct {
	CreationTime          string
	DisplayedMatchesCount string
	ID                    string
	InConflict            string
	Isrc                  string
	MatchRule             string
	OwnershipCountries    string
	OwnershipDetails      string
	ReferenceFileStatus   string
	RidgeMonitoringStatus string
	Tags                  string
	UpdateTime            string
	WhitelistedFbUsers    string
	WhitelistedIgUsers    string
}{
	CreationTime:          "creation_time",
	DisplayedMatchesCount: "displayed_matches_count",
	ID:                    "id",
	InConflict:            "in_conflict",
	Isrc:                  "isrc",
	MatchRule:             "match_rule",
	OwnershipCountries:    "ownership_countries",
	OwnershipDetails:      "ownership_details",
	ReferenceFileStatus:   "reference_file_status",
	RidgeMonitoringStatus: "ridge_monitoring_status",
	Tags:                  "tags",
	UpdateTime:            "update_time",
	WhitelistedFbUsers:    "whitelisted_fb_users",
	WhitelistedIgUsers:    "whitelisted_ig_users",
}
