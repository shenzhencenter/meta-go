package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type OfflineConversionDataSet struct {
	AutomaticMatchingFields    *[]string                      `json:"automatic_matching_fields,omitempty"`
	Business                   *Business                      `json:"business,omitempty"`
	CanProxy                   *bool                          `json:"can_proxy,omitempty"`
	Config                     *string                        `json:"config,omitempty"`
	CreationTime               *time.Time                     `json:"creation_time,omitempty"`
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
	ID                         *core.ID                       `json:"id,omitempty"`
	IsConsolidatedContainer    *bool                          `json:"is_consolidated_container,omitempty"`
	IsCreatedByBusiness        *bool                          `json:"is_created_by_business,omitempty"`
	IsCrm                      *bool                          `json:"is_crm,omitempty"`
	IsMtaUse                   *bool                          `json:"is_mta_use,omitempty"`
	IsRestrictedUse            *bool                          `json:"is_restricted_use,omitempty"`
	IsUnavailable              *bool                          `json:"is_unavailable,omitempty"`
	LastFiredTime              *time.Time                     `json:"last_fired_time,omitempty"`
	LastUploadApp              *string                        `json:"last_upload_app,omitempty"`
	LastUploadAppChangedTime   *int                           `json:"last_upload_app_changed_time,omitempty"`
	MatchRateApprox            *int                           `json:"match_rate_approx,omitempty"`
	MatchedEntries             *int                           `json:"matched_entries,omitempty"`
	Name                       *string                        `json:"name,omitempty"`
	OwnerAdAccount             *AdAccount                     `json:"owner_ad_account,omitempty"`
	OwnerBusiness              *Business                      `json:"owner_business,omitempty"`
	Usage                      *OfflineConversionDataSetUsage `json:"usage,omitempty"`
	ValidEntries               *int                           `json:"valid_entries,omitempty"`
}
