package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsDataset struct {
	CanProxy                                               *bool                                `json:"can_proxy,omitempty"`
	CollectionRate                                         *float64                             `json:"collection_rate,omitempty"`
	Config                                                 *string                              `json:"config,omitempty"`
	CreationTime                                           *core.Time                           `json:"creation_time,omitempty"`
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
	LastFiredTime                                          *core.Time                           `json:"last_fired_time,omitempty"`
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
	ServerLastFiredTime                                    *core.Time                           `json:"server_last_fired_time,omitempty"`
	ShowAutomaticEvents                                    *bool                                `json:"show_automatic_events,omitempty"`
	UploadRate                                             *float64                             `json:"upload_rate,omitempty"`
	UploadReminderEligibility                              *bool                                `json:"upload_reminder_eligibility,omitempty"`
	Usage                                                  *OfflineConversionDataSetUsage       `json:"usage,omitempty"`
	ValidEntries                                           *int                                 `json:"valid_entries,omitempty"`
}

var AdsDatasetFields = struct {
	CanProxy                                               string
	CollectionRate                                         string
	Config                                                 string
	CreationTime                                           string
	Creator                                                string
	DatasetID                                              string
	Description                                            string
	DuplicateEntries                                       string
	EnableAutoAssignToAccounts                             string
	EnableAutomaticEvents                                  string
	EnableAutomaticMatching                                string
	EnableRealTimeEventLog                                 string
	EventStats                                             string
	EventTimeMax                                           string
	EventTimeMin                                           string
	FirstPartyCookieStatus                                 string
	HasBapiDomains                                         string
	HasCatalogMicrodataActivity                            string
	HasOfaRedactedKeys                                     string
	HasSentPii                                             string
	ID                                                     string
	IsConsolidatedContainer                                string
	IsCreatedByBusiness                                    string
	IsCrm                                                  string
	IsEligibleForSharingToAdAccount                        string
	IsEligibleForSharingToBusiness                         string
	IsEligibleForValueOptimization                         string
	IsMtaUse                                               string
	IsRestrictedUse                                        string
	IsUnavailable                                          string
	LastFiredTime                                          string
	LastUploadApp                                          string
	LastUploadAppChangedTime                               string
	LastUploadTime                                         string
	LateUploadReminderEligibility                          string
	MatchRateApprox                                        string
	MatchedEntries                                         string
	Name                                                   string
	NoAdsTrackedForWeeklyUploadedEventsReminderEligibility string
	NumActiveAdSetTracked                                  string
	NumRecentOfflineConversionsUploaded                    string
	NumUploads                                             string
	OwnerAdAccount                                         string
	OwnerBusiness                                          string
	PercentageOfLateUploadsInExternalSuboptimalWindow      string
	Permissions                                            string
	ServerLastFiredTime                                    string
	ShowAutomaticEvents                                    string
	UploadRate                                             string
	UploadReminderEligibility                              string
	Usage                                                  string
	ValidEntries                                           string
}{
	CanProxy:                        "can_proxy",
	CollectionRate:                  "collection_rate",
	Config:                          "config",
	CreationTime:                    "creation_time",
	Creator:                         "creator",
	DatasetID:                       "dataset_id",
	Description:                     "description",
	DuplicateEntries:                "duplicate_entries",
	EnableAutoAssignToAccounts:      "enable_auto_assign_to_accounts",
	EnableAutomaticEvents:           "enable_automatic_events",
	EnableAutomaticMatching:         "enable_automatic_matching",
	EnableRealTimeEventLog:          "enable_real_time_event_log",
	EventStats:                      "event_stats",
	EventTimeMax:                    "event_time_max",
	EventTimeMin:                    "event_time_min",
	FirstPartyCookieStatus:          "first_party_cookie_status",
	HasBapiDomains:                  "has_bapi_domains",
	HasCatalogMicrodataActivity:     "has_catalog_microdata_activity",
	HasOfaRedactedKeys:              "has_ofa_redacted_keys",
	HasSentPii:                      "has_sent_pii",
	ID:                              "id",
	IsConsolidatedContainer:         "is_consolidated_container",
	IsCreatedByBusiness:             "is_created_by_business",
	IsCrm:                           "is_crm",
	IsEligibleForSharingToAdAccount: "is_eligible_for_sharing_to_ad_account",
	IsEligibleForSharingToBusiness:  "is_eligible_for_sharing_to_business",
	IsEligibleForValueOptimization:  "is_eligible_for_value_optimization",
	IsMtaUse:                        "is_mta_use",
	IsRestrictedUse:                 "is_restricted_use",
	IsUnavailable:                   "is_unavailable",
	LastFiredTime:                   "last_fired_time",
	LastUploadApp:                   "last_upload_app",
	LastUploadAppChangedTime:        "last_upload_app_changed_time",
	LastUploadTime:                  "last_upload_time",
	LateUploadReminderEligibility:   "late_upload_reminder_eligibility",
	MatchRateApprox:                 "match_rate_approx",
	MatchedEntries:                  "matched_entries",
	Name:                            "name",
	NoAdsTrackedForWeeklyUploadedEventsReminderEligibility: "no_ads_tracked_for_weekly_uploaded_events_reminder_eligibility",
	NumActiveAdSetTracked:               "num_active_ad_set_tracked",
	NumRecentOfflineConversionsUploaded: "num_recent_offline_conversions_uploaded",
	NumUploads:                          "num_uploads",
	OwnerAdAccount:                      "owner_ad_account",
	OwnerBusiness:                       "owner_business",
	PercentageOfLateUploadsInExternalSuboptimalWindow: "percentage_of_late_uploads_in_external_suboptimal_window",
	Permissions:               "permissions",
	ServerLastFiredTime:       "server_last_fired_time",
	ShowAutomaticEvents:       "show_automatic_events",
	UploadRate:                "upload_rate",
	UploadReminderEligibility: "upload_reminder_eligibility",
	Usage:                     "usage",
	ValidEntries:              "valid_entries",
}
