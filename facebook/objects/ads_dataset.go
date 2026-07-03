package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdsDataset struct {
	CanProxy                                               *bool                                `json:"can_proxy,omitempty"`
	CollectionRate                                         *float64                             `json:"collection_rate,omitempty"`
	Config                                                 *string                              `json:"config,omitempty"`
	CreationTime                                           *time.Time                           `json:"creation_time,omitempty"`
	Creator                                                *User                                `json:"creator,omitempty"`
	DatasetID                                              *core.ID                             `json:"dataset_id,omitempty"`
	Description                                            *string                              `json:"description,omitempty"`
	DuplicateEntries                                       *int                                 `json:"duplicate_entries,omitempty"`
	EnableAutoAssignToAccounts                             *bool                                `json:"enable_auto_assign_to_accounts,omitempty"`
	EnableAutomaticEvents                                  *bool                                `json:"enable_automatic_events,omitempty"`
	EnableAutomaticMatching                                *bool                                `json:"enable_automatic_matching,omitempty"`
	EnableRealTimeEventLog                                 *bool                                `json:"enable_real_time_event_log,omitempty"`
	EventStats                                             *string                              `json:"event_stats,omitempty"`
	EventTimeMax                                           *int                                 `json:"event_time_max,omitempty"`
	EventTimeMin                                           *int                                 `json:"event_time_min,omitempty"`
	FirstPartyCookieStatus                                 *string                              `json:"first_party_cookie_status,omitempty"`
	HasBapiDomains                                         *bool                                `json:"has_bapi_domains,omitempty"`
	HasCatalogMicrodataActivity                            *bool                                `json:"has_catalog_microdata_activity,omitempty"`
	HasOfaRedactedKeys                                     *bool                                `json:"has_ofa_redacted_keys,omitempty"`
	HasSentPii                                             *bool                                `json:"has_sent_pii,omitempty"`
	ID                                                     *core.ID                             `json:"id,omitempty"`
	IsConsolidatedContainer                                *bool                                `json:"is_consolidated_container,omitempty"`
	IsCreatedByBusiness                                    *bool                                `json:"is_created_by_business,omitempty"`
	IsCrm                                                  *bool                                `json:"is_crm,omitempty"`
	IsEligibleForSharingToAdAccount                        *bool                                `json:"is_eligible_for_sharing_to_ad_account,omitempty"`
	IsEligibleForSharingToBusiness                         *bool                                `json:"is_eligible_for_sharing_to_business,omitempty"`
	IsEligibleForValueOptimization                         *bool                                `json:"is_eligible_for_value_optimization,omitempty"`
	IsMtaUse                                               *bool                                `json:"is_mta_use,omitempty"`
	IsRestrictedUse                                        *bool                                `json:"is_restricted_use,omitempty"`
	IsUnavailable                                          *bool                                `json:"is_unavailable,omitempty"`
	LastFiredTime                                          *time.Time                           `json:"last_fired_time,omitempty"`
	LastUploadApp                                          *string                              `json:"last_upload_app,omitempty"`
	LastUploadAppChangedTime                               *int                                 `json:"last_upload_app_changed_time,omitempty"`
	LastUploadTime                                         *int                                 `json:"last_upload_time,omitempty"`
	LateUploadReminderEligibility                          *bool                                `json:"late_upload_reminder_eligibility,omitempty"`
	MatchRateApprox                                        *int                                 `json:"match_rate_approx,omitempty"`
	MatchedEntries                                         *int                                 `json:"matched_entries,omitempty"`
	Name                                                   *string                              `json:"name,omitempty"`
	NoAdsTrackedForWeeklyUploadedEventsReminderEligibility *bool                                `json:"no_ads_tracked_for_weekly_uploaded_events_reminder_eligibility,omitempty"`
	NumActiveAdSetTracked                                  *int                                 `json:"num_active_ad_set_tracked,omitempty"`
	NumRecentOfflineConversionsUploaded                    *int                                 `json:"num_recent_offline_conversions_uploaded,omitempty"`
	NumUploads                                             *int                                 `json:"num_uploads,omitempty"`
	OwnerAdAccount                                         *AdAccount                           `json:"owner_ad_account,omitempty"`
	OwnerBusiness                                          *Business                            `json:"owner_business,omitempty"`
	PercentageOfLateUploadsInExternalSuboptimalWindow      *int                                 `json:"percentage_of_late_uploads_in_external_suboptimal_window,omitempty"`
	Permissions                                            *OfflineConversionDataSetPermissions `json:"permissions,omitempty"`
	ServerLastFiredTime                                    *time.Time                           `json:"server_last_fired_time,omitempty"`
	ShowAutomaticEvents                                    *bool                                `json:"show_automatic_events,omitempty"`
	UploadRate                                             *float64                             `json:"upload_rate,omitempty"`
	UploadReminderEligibility                              *bool                                `json:"upload_reminder_eligibility,omitempty"`
	Usage                                                  *OfflineConversionDataSetUsage       `json:"usage,omitempty"`
	ValidEntries                                           *int                                 `json:"valid_entries,omitempty"`
}
