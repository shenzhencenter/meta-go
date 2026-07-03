package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixel struct {
	AutomaticMatchingFields    *[]string                      `json:"automatic_matching_fields,omitempty"`
	CanProxy                   *bool                          `json:"can_proxy,omitempty"`
	Code                       *string                        `json:"code,omitempty"`
	Config                     *string                        `json:"config,omitempty"`
	CreationTime               *core.Time                     `json:"creation_time,omitempty"`
	Creator                    *User                          `json:"creator,omitempty"`
	DataUseSetting             *string                        `json:"data_use_setting,omitempty"`
	Description                *string                        `json:"description,omitempty"`
	DuplicateEntries           *int                           `json:"duplicate_entries,omitempty"`
	EnableAutoAssignToAccounts *bool                          `json:"enable_auto_assign_to_accounts,omitempty"`
	EnableAutomaticMatching    *bool                          `json:"enable_automatic_matching,omitempty"`
	EventStats                 *string                        `json:"event_stats,omitempty"`
	EventTimeMax               *int                           `json:"event_time_max,omitempty"`
	EventTimeMin               *int                           `json:"event_time_min,omitempty"`
	FirstPartyCookieStatus     *string                        `json:"first_party_cookie_status,omitempty"`
	HasX1pPixelEvent           *bool                          `json:"has_1p_pixel_event,omitempty"`
	ID                         *core.ID                       `json:"id,omitempty"`
	IsConsolidatedContainer    *bool                          `json:"is_consolidated_container,omitempty"`
	IsCreatedByBusiness        *bool                          `json:"is_created_by_business,omitempty"`
	IsCrm                      *bool                          `json:"is_crm,omitempty"`
	IsMtaUse                   *bool                          `json:"is_mta_use,omitempty"`
	IsRestrictedUse            *bool                          `json:"is_restricted_use,omitempty"`
	IsUnavailable              *bool                          `json:"is_unavailable,omitempty"`
	LastFiredTime              *core.Time                     `json:"last_fired_time,omitempty"`
	LastUploadApp              *string                        `json:"last_upload_app,omitempty"`
	LastUploadAppChangedTime   *int                           `json:"last_upload_app_changed_time,omitempty"`
	MatchRateApprox            *int                           `json:"match_rate_approx,omitempty"`
	MatchedEntries             *int                           `json:"matched_entries,omitempty"`
	Name                       *string                        `json:"name,omitempty"`
	OwnerAdAccount             *AdAccount                     `json:"owner_ad_account,omitempty"`
	OwnerBusiness              *Business                      `json:"owner_business,omitempty"`
	Usage                      *OfflineConversionDataSetUsage `json:"usage,omitempty"`
	UserAccessExpireTime       *core.Time                     `json:"user_access_expire_time,omitempty"`
	ValidEntries               *int                           `json:"valid_entries,omitempty"`
}

var AdsPixelFields = struct {
	AutomaticMatchingFields    string
	CanProxy                   string
	Code                       string
	Config                     string
	CreationTime               string
	Creator                    string
	DataUseSetting             string
	Description                string
	DuplicateEntries           string
	EnableAutoAssignToAccounts string
	EnableAutomaticMatching    string
	EventStats                 string
	EventTimeMax               string
	EventTimeMin               string
	FirstPartyCookieStatus     string
	HasX1pPixelEvent           string
	ID                         string
	IsConsolidatedContainer    string
	IsCreatedByBusiness        string
	IsCrm                      string
	IsMtaUse                   string
	IsRestrictedUse            string
	IsUnavailable              string
	LastFiredTime              string
	LastUploadApp              string
	LastUploadAppChangedTime   string
	MatchRateApprox            string
	MatchedEntries             string
	Name                       string
	OwnerAdAccount             string
	OwnerBusiness              string
	Usage                      string
	UserAccessExpireTime       string
	ValidEntries               string
}{
	AutomaticMatchingFields:    "automatic_matching_fields",
	CanProxy:                   "can_proxy",
	Code:                       "code",
	Config:                     "config",
	CreationTime:               "creation_time",
	Creator:                    "creator",
	DataUseSetting:             "data_use_setting",
	Description:                "description",
	DuplicateEntries:           "duplicate_entries",
	EnableAutoAssignToAccounts: "enable_auto_assign_to_accounts",
	EnableAutomaticMatching:    "enable_automatic_matching",
	EventStats:                 "event_stats",
	EventTimeMax:               "event_time_max",
	EventTimeMin:               "event_time_min",
	FirstPartyCookieStatus:     "first_party_cookie_status",
	HasX1pPixelEvent:           "has_1p_pixel_event",
	ID:                         "id",
	IsConsolidatedContainer:    "is_consolidated_container",
	IsCreatedByBusiness:        "is_created_by_business",
	IsCrm:                      "is_crm",
	IsMtaUse:                   "is_mta_use",
	IsRestrictedUse:            "is_restricted_use",
	IsUnavailable:              "is_unavailable",
	LastFiredTime:              "last_fired_time",
	LastUploadApp:              "last_upload_app",
	LastUploadAppChangedTime:   "last_upload_app_changed_time",
	MatchRateApprox:            "match_rate_approx",
	MatchedEntries:             "matched_entries",
	Name:                       "name",
	OwnerAdAccount:             "owner_ad_account",
	OwnerBusiness:              "owner_business",
	Usage:                      "usage",
	UserAccessExpireTime:       "user_access_expire_time",
	ValidEntries:               "valid_entries",
}
