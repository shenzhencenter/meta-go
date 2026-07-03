package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
