package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AnalyticsConfig struct {
	AnalyticsAccessForAuthorizedAdAccount       *bool                     `json:"analytics_access_for_authorized_ad_account,omitempty"`
	BreakdownsConfig                            *[]map[string]interface{} `json:"breakdowns_config,omitempty"`
	BuiltinFieldsConfig                         *[]map[string]interface{} `json:"builtin_fields_config,omitempty"`
	DeprecatedEventsConfig                      *[]map[string]interface{} `json:"deprecated_events_config,omitempty"`
	EventsConfig                                *[]map[string]interface{} `json:"events_config,omitempty"`
	IosPurchaseValidationSecret                 *string                   `json:"ios_purchase_validation_secret,omitempty"`
	IsAnyRoleAbleToSeeRestrictedInsights        *bool                     `json:"is_any_role_able_to_see_restricted_insights,omitempty"`
	IsImplicitPurchaseLoggingOnAndroidSupported *bool                     `json:"is_implicit_purchase_logging_on_android_supported,omitempty"`
	IsImplicitPurchaseLoggingOnIosSupported     *bool                     `json:"is_implicit_purchase_logging_on_ios_supported,omitempty"`
	IsTrackIosAppUninstallSupported             *bool                     `json:"is_track_ios_app_uninstall_supported,omitempty"`
	JourneyBackfillStatus                       *string                   `json:"journey_backfill_status,omitempty"`
	JourneyConversionEvents                     *[]string                 `json:"journey_conversion_events,omitempty"`
	JourneyEnabled                              *bool                     `json:"journey_enabled,omitempty"`
	JourneyImpactingChangeTime                  *core.Time                `json:"journey_impacting_change_time,omitempty"`
	JourneyTimeout                              *string                   `json:"journey_timeout,omitempty"`
	LatestSdkVersions                           *map[string]string        `json:"latest_sdk_versions,omitempty"`
	LogAndroidImplicitPurchaseEvents            *bool                     `json:"log_android_implicit_purchase_events,omitempty"`
	LogAutomaticAnalyticsEvents                 *bool                     `json:"log_automatic_analytics_events,omitempty"`
	LogImplicitPurchaseEvents                   *bool                     `json:"log_implicit_purchase_events,omitempty"`
	PrevJourneyConversionEvents                 *[]string                 `json:"prev_journey_conversion_events,omitempty"`
	QueryApproximationAccuracyLevel             *string                   `json:"query_approximation_accuracy_level,omitempty"`
	QueryCurrency                               *string                   `json:"query_currency,omitempty"`
	QueryTimezone                               *string                   `json:"query_timezone,omitempty"`
	RecentEventsUpdateTime                      *core.Time                `json:"recent_events_update_time,omitempty"`
	SessionTimeoutInterval                      *uint64                   `json:"session_timeout_interval,omitempty"`
	TrackIosAppUninstall                        *bool                     `json:"track_ios_app_uninstall,omitempty"`
}

var AnalyticsConfigFields = struct {
	AnalyticsAccessForAuthorizedAdAccount       string
	BreakdownsConfig                            string
	BuiltinFieldsConfig                         string
	DeprecatedEventsConfig                      string
	EventsConfig                                string
	IosPurchaseValidationSecret                 string
	IsAnyRoleAbleToSeeRestrictedInsights        string
	IsImplicitPurchaseLoggingOnAndroidSupported string
	IsImplicitPurchaseLoggingOnIosSupported     string
	IsTrackIosAppUninstallSupported             string
	JourneyBackfillStatus                       string
	JourneyConversionEvents                     string
	JourneyEnabled                              string
	JourneyImpactingChangeTime                  string
	JourneyTimeout                              string
	LatestSdkVersions                           string
	LogAndroidImplicitPurchaseEvents            string
	LogAutomaticAnalyticsEvents                 string
	LogImplicitPurchaseEvents                   string
	PrevJourneyConversionEvents                 string
	QueryApproximationAccuracyLevel             string
	QueryCurrency                               string
	QueryTimezone                               string
	RecentEventsUpdateTime                      string
	SessionTimeoutInterval                      string
	TrackIosAppUninstall                        string
}{
	AnalyticsAccessForAuthorizedAdAccount:       "analytics_access_for_authorized_ad_account",
	BreakdownsConfig:                            "breakdowns_config",
	BuiltinFieldsConfig:                         "builtin_fields_config",
	DeprecatedEventsConfig:                      "deprecated_events_config",
	EventsConfig:                                "events_config",
	IosPurchaseValidationSecret:                 "ios_purchase_validation_secret",
	IsAnyRoleAbleToSeeRestrictedInsights:        "is_any_role_able_to_see_restricted_insights",
	IsImplicitPurchaseLoggingOnAndroidSupported: "is_implicit_purchase_logging_on_android_supported",
	IsImplicitPurchaseLoggingOnIosSupported:     "is_implicit_purchase_logging_on_ios_supported",
	IsTrackIosAppUninstallSupported:             "is_track_ios_app_uninstall_supported",
	JourneyBackfillStatus:                       "journey_backfill_status",
	JourneyConversionEvents:                     "journey_conversion_events",
	JourneyEnabled:                              "journey_enabled",
	JourneyImpactingChangeTime:                  "journey_impacting_change_time",
	JourneyTimeout:                              "journey_timeout",
	LatestSdkVersions:                           "latest_sdk_versions",
	LogAndroidImplicitPurchaseEvents:            "log_android_implicit_purchase_events",
	LogAutomaticAnalyticsEvents:                 "log_automatic_analytics_events",
	LogImplicitPurchaseEvents:                   "log_implicit_purchase_events",
	PrevJourneyConversionEvents:                 "prev_journey_conversion_events",
	QueryApproximationAccuracyLevel:             "query_approximation_accuracy_level",
	QueryCurrency:                               "query_currency",
	QueryTimezone:                               "query_timezone",
	RecentEventsUpdateTime:                      "recent_events_update_time",
	SessionTimeoutInterval:                      "session_timeout_interval",
	TrackIosAppUninstall:                        "track_ios_app_uninstall",
}
