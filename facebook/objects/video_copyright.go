package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoCopyright struct {
	ContentCategory                         *string                      `json:"content_category,omitempty"`
	ContentProtectProtectionDisabledReason  *string                      `json:"content_protect_protection_disabled_reason,omitempty"`
	CopyrightContentID                      *core.ID                     `json:"copyright_content_id,omitempty"`
	Creator                                 *User                        `json:"creator,omitempty"`
	DisableProtectionByContentProtectStatus *bool                        `json:"disable_protection_by_content_protect_status,omitempty"`
	ExcludedOwnershipSegments               *[]VideoCopyrightSegment     `json:"excluded_ownership_segments,omitempty"`
	ID                                      *core.ID                     `json:"id,omitempty"`
	InConflict                              *bool                        `json:"in_conflict,omitempty"`
	MonitoringStatus                        *string                      `json:"monitoring_status,omitempty"`
	MonitoringType                          *string                      `json:"monitoring_type,omitempty"`
	OwnershipCountries                      *VideoCopyrightGeoGate       `json:"ownership_countries,omitempty"`
	ReferenceFile                           *CopyrightReferenceContainer `json:"reference_file,omitempty"`
	ReferenceFileDisabled                   *bool                        `json:"reference_file_disabled,omitempty"`
	ReferenceFileDisabledByOps              *bool                        `json:"reference_file_disabled_by_ops,omitempty"`
	ReferenceOwnerID                        *core.ID                     `json:"reference_owner_id,omitempty"`
	RuleIds                                 *[]VideoCopyrightRule        `json:"rule_ids,omitempty"`
	Tags                                    *[]string                    `json:"tags,omitempty"`
	WhitelistedIds                          *[]core.ID                   `json:"whitelisted_ids,omitempty"`
}

var VideoCopyrightFields = struct {
	ContentCategory                         string
	ContentProtectProtectionDisabledReason  string
	CopyrightContentID                      string
	Creator                                 string
	DisableProtectionByContentProtectStatus string
	ExcludedOwnershipSegments               string
	ID                                      string
	InConflict                              string
	MonitoringStatus                        string
	MonitoringType                          string
	OwnershipCountries                      string
	ReferenceFile                           string
	ReferenceFileDisabled                   string
	ReferenceFileDisabledByOps              string
	ReferenceOwnerID                        string
	RuleIds                                 string
	Tags                                    string
	WhitelistedIds                          string
}{
	ContentCategory:                         "content_category",
	ContentProtectProtectionDisabledReason:  "content_protect_protection_disabled_reason",
	CopyrightContentID:                      "copyright_content_id",
	Creator:                                 "creator",
	DisableProtectionByContentProtectStatus: "disable_protection_by_content_protect_status",
	ExcludedOwnershipSegments:               "excluded_ownership_segments",
	ID:                                      "id",
	InConflict:                              "in_conflict",
	MonitoringStatus:                        "monitoring_status",
	MonitoringType:                          "monitoring_type",
	OwnershipCountries:                      "ownership_countries",
	ReferenceFile:                           "reference_file",
	ReferenceFileDisabled:                   "reference_file_disabled",
	ReferenceFileDisabledByOps:              "reference_file_disabled_by_ops",
	ReferenceOwnerID:                        "reference_owner_id",
	RuleIds:                                 "rule_ids",
	Tags:                                    "tags",
	WhitelistedIds:                          "whitelisted_ids",
}
