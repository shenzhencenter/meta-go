package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Application struct {
	AamRules                           *string                                `json:"aam_rules,omitempty"`
	AnAdSpaceLimit                     *uint64                                `json:"an_ad_space_limit,omitempty"`
	AnPlatforms                        *[]string                              `json:"an_platforms,omitempty"`
	AndroidKeyHash                     *[]string                              `json:"android_key_hash,omitempty"`
	AndroidSdkErrorCategories          *[]map[string]interface{}              `json:"android_sdk_error_categories,omitempty"`
	AppDomains                         *[]string                              `json:"app_domains,omitempty"`
	AppEventsConfig                    *map[string]interface{}                `json:"app_events_config,omitempty"`
	AppEventsFeatureBitmask            *uint64                                `json:"app_events_feature_bitmask,omitempty"`
	AppEventsSessionTimeout            *uint64                                `json:"app_events_session_timeout,omitempty"`
	AppInstallTracked                  *bool                                  `json:"app_install_tracked,omitempty"`
	AppName                            *string                                `json:"app_name,omitempty"`
	AppSignalsBindingIos               *[]map[string]interface{}              `json:"app_signals_binding_ios,omitempty"`
	AppType                            *uint64                                `json:"app_type,omitempty"`
	AuthDialogDataHelpURL              *string                                `json:"auth_dialog_data_help_url,omitempty"`
	AuthDialogHeadline                 *string                                `json:"auth_dialog_headline,omitempty"`
	AuthDialogPermsExplanation         *string                                `json:"auth_dialog_perms_explanation,omitempty"`
	AuthReferralDefaultActivityPrivacy *string                                `json:"auth_referral_default_activity_privacy,omitempty"`
	AuthReferralEnabled                *uint64                                `json:"auth_referral_enabled,omitempty"`
	AuthReferralExtendedPerms          *[]string                              `json:"auth_referral_extended_perms,omitempty"`
	AuthReferralFriendPerms            *[]string                              `json:"auth_referral_friend_perms,omitempty"`
	AuthReferralResponseType           *string                                `json:"auth_referral_response_type,omitempty"`
	AuthReferralUserPerms              *[]string                              `json:"auth_referral_user_perms,omitempty"`
	AutoEventMappingAndroid            *[]map[string]interface{}              `json:"auto_event_mapping_android,omitempty"`
	AutoEventMappingIos                *[]map[string]interface{}              `json:"auto_event_mapping_ios,omitempty"`
	AutoEventSetupEnabled              *bool                                  `json:"auto_event_setup_enabled,omitempty"`
	AutoLogAppEventsDefault            *bool                                  `json:"auto_log_app_events_default,omitempty"`
	AutoLogAppEventsEnabled            *bool                                  `json:"auto_log_app_events_enabled,omitempty"`
	Business                           *Business                              `json:"business,omitempty"`
	CanvasFluidHeight                  *bool                                  `json:"canvas_fluid_height,omitempty"`
	CanvasFluidWidth                   *uint64                                `json:"canvas_fluid_width,omitempty"`
	CanvasURL                          *string                                `json:"canvas_url,omitempty"`
	Category                           *string                                `json:"category,omitempty"`
	ClientConfig                       *map[string]interface{}                `json:"client_config,omitempty"`
	Company                            *string                                `json:"company,omitempty"`
	ConfigIds                          *[]map[string]interface{}              `json:"config_ids,omitempty"`
	ConfiguredIosSso                   *bool                                  `json:"configured_ios_sso,omitempty"`
	ContactEmail                       *string                                `json:"contact_email,omitempty"`
	CreatedTime                        *core.Time                             `json:"created_time,omitempty"`
	CreatorUID                         *core.ID                               `json:"creator_uid,omitempty"`
	DailyActiveUsers                   *string                                `json:"daily_active_users,omitempty"`
	DailyActiveUsersRank               *uint64                                `json:"daily_active_users_rank,omitempty"`
	DeauthCallbackURL                  *string                                `json:"deauth_callback_url,omitempty"`
	DefaultShareMode                   *string                                `json:"default_share_mode,omitempty"`
	Description                        *string                                `json:"description,omitempty"`
	EnigmaConfig                       *map[string]interface{}                `json:"enigma_config,omitempty"`
	FinancialID                        *core.ID                               `json:"financial_id,omitempty"`
	Gdpv4ChromeCustomTabsEnabled       *bool                                  `json:"gdpv4_chrome_custom_tabs_enabled,omitempty"`
	Gdpv4Enabled                       *bool                                  `json:"gdpv4_enabled,omitempty"`
	Gdpv4NuxContent                    *string                                `json:"gdpv4_nux_content,omitempty"`
	Gdpv4NuxEnabled                    *bool                                  `json:"gdpv4_nux_enabled,omitempty"`
	HasMessengerProduct                *bool                                  `json:"has_messenger_product,omitempty"`
	HostingURL                         *string                                `json:"hosting_url,omitempty"`
	IconURL                            *string                                `json:"icon_url,omitempty"`
	ID                                 *core.ID                               `json:"id,omitempty"`
	IosBundleID                        *[]core.ID                             `json:"ios_bundle_id,omitempty"`
	IosSdkDialogFlows                  *map[string]interface{}                `json:"ios_sdk_dialog_flows,omitempty"`
	IosSdkErrorCategories              *[]map[string]interface{}              `json:"ios_sdk_error_categories,omitempty"`
	IosSfvcAttr                        *bool                                  `json:"ios_sfvc_attr,omitempty"`
	IosSupportsNativeProxyAuthFlow     *bool                                  `json:"ios_supports_native_proxy_auth_flow,omitempty"`
	IosSupportsSystemAuth              *bool                                  `json:"ios_supports_system_auth,omitempty"`
	IpadAppStoreID                     *core.ID                               `json:"ipad_app_store_id,omitempty"`
	IphoneAppStoreID                   *core.ID                               `json:"iphone_app_store_id,omitempty"`
	LatestSdkVersion                   *map[string]interface{}                `json:"latest_sdk_version,omitempty"`
	Link                               *string                                `json:"link,omitempty"`
	LoggingToken                       *string                                `json:"logging_token,omitempty"`
	LogoURL                            *string                                `json:"logo_url,omitempty"`
	Migrations                         *map[string]bool                       `json:"migrations,omitempty"`
	MobileProfileSectionURL            *string                                `json:"mobile_profile_section_url,omitempty"`
	MobileWebURL                       *string                                `json:"mobile_web_url,omitempty"`
	MonthlyActiveUsers                 *string                                `json:"monthly_active_users,omitempty"`
	MonthlyActiveUsersRank             *uint64                                `json:"monthly_active_users_rank,omitempty"`
	Name                               *string                                `json:"name,omitempty"`
	Namespace                          *string                                `json:"namespace,omitempty"`
	ObjectStoreUrls                    *map[string]interface{}                `json:"object_store_urls,omitempty"`
	OwnerBusiness                      *Business                              `json:"owner_business,omitempty"`
	PageTabDefaultName                 *string                                `json:"page_tab_default_name,omitempty"`
	PageTabURL                         *string                                `json:"page_tab_url,omitempty"`
	PhotoURL                           *string                                `json:"photo_url,omitempty"`
	PrivacyPolicyURL                   *string                                `json:"privacy_policy_url,omitempty"`
	ProfileSectionURL                  *string                                `json:"profile_section_url,omitempty"`
	PropertyID                         *core.ID                               `json:"property_id,omitempty"`
	ProtectedModeRules                 *map[string]interface{}                `json:"protected_mode_rules,omitempty"`
	RealTimeModeDevices                *[]string                              `json:"real_time_mode_devices,omitempty"`
	Restrictions                       *map[string]interface{}                `json:"restrictions,omitempty"`
	RestrictiveDataFilterParams        *string                                `json:"restrictive_data_filter_params,omitempty"`
	RestrictiveDataFilterRules         *string                                `json:"restrictive_data_filter_rules,omitempty"`
	SdkUpdateMessage                   *string                                `json:"sdk_update_message,omitempty"`
	SeamlessLogin                      *int                                   `json:"seamless_login,omitempty"`
	SecureCanvasURL                    *string                                `json:"secure_canvas_url,omitempty"`
	SecurePageTabURL                   *string                                `json:"secure_page_tab_url,omitempty"`
	ServerIPWhitelist                  *string                                `json:"server_ip_whitelist,omitempty"`
	SmartLoginBookmarkIconURL          *string                                `json:"smart_login_bookmark_icon_url,omitempty"`
	SmartLoginMenuIconURL              *string                                `json:"smart_login_menu_icon_url,omitempty"`
	SocialDiscovery                    *uint64                                `json:"social_discovery,omitempty"`
	Subcategory                        *string                                `json:"subcategory,omitempty"`
	SuggestedEventsSetting             *string                                `json:"suggested_events_setting,omitempty"`
	SupportedPlatforms                 *[]enums.ApplicationSupportedPlatforms `json:"supported_platforms,omitempty"`
	SupportsApprequestsFastAppSwitch   *map[string]interface{}                `json:"supports_apprequests_fast_app_switch,omitempty"`
	SupportsAttribution                *bool                                  `json:"supports_attribution,omitempty"`
	SupportsImplicitSdkLogging         *bool                                  `json:"supports_implicit_sdk_logging,omitempty"`
	SuppressNativeIosGdp               *bool                                  `json:"suppress_native_ios_gdp,omitempty"`
	TermsOfServiceURL                  *string                                `json:"terms_of_service_url,omitempty"`
	URLSchemeSuffix                    *string                                `json:"url_scheme_suffix,omitempty"`
	UserSupportEmail                   *string                                `json:"user_support_email,omitempty"`
	UserSupportURL                     *string                                `json:"user_support_url,omitempty"`
	WebsiteURL                         *string                                `json:"website_url,omitempty"`
	WeeklyActiveUsers                  *string                                `json:"weekly_active_users,omitempty"`
}

var ApplicationFields = struct {
	AamRules                           string
	AnAdSpaceLimit                     string
	AnPlatforms                        string
	AndroidKeyHash                     string
	AndroidSdkErrorCategories          string
	AppDomains                         string
	AppEventsConfig                    string
	AppEventsFeatureBitmask            string
	AppEventsSessionTimeout            string
	AppInstallTracked                  string
	AppName                            string
	AppSignalsBindingIos               string
	AppType                            string
	AuthDialogDataHelpURL              string
	AuthDialogHeadline                 string
	AuthDialogPermsExplanation         string
	AuthReferralDefaultActivityPrivacy string
	AuthReferralEnabled                string
	AuthReferralExtendedPerms          string
	AuthReferralFriendPerms            string
	AuthReferralResponseType           string
	AuthReferralUserPerms              string
	AutoEventMappingAndroid            string
	AutoEventMappingIos                string
	AutoEventSetupEnabled              string
	AutoLogAppEventsDefault            string
	AutoLogAppEventsEnabled            string
	Business                           string
	CanvasFluidHeight                  string
	CanvasFluidWidth                   string
	CanvasURL                          string
	Category                           string
	ClientConfig                       string
	Company                            string
	ConfigIds                          string
	ConfiguredIosSso                   string
	ContactEmail                       string
	CreatedTime                        string
	CreatorUID                         string
	DailyActiveUsers                   string
	DailyActiveUsersRank               string
	DeauthCallbackURL                  string
	DefaultShareMode                   string
	Description                        string
	EnigmaConfig                       string
	FinancialID                        string
	Gdpv4ChromeCustomTabsEnabled       string
	Gdpv4Enabled                       string
	Gdpv4NuxContent                    string
	Gdpv4NuxEnabled                    string
	HasMessengerProduct                string
	HostingURL                         string
	IconURL                            string
	ID                                 string
	IosBundleID                        string
	IosSdkDialogFlows                  string
	IosSdkErrorCategories              string
	IosSfvcAttr                        string
	IosSupportsNativeProxyAuthFlow     string
	IosSupportsSystemAuth              string
	IpadAppStoreID                     string
	IphoneAppStoreID                   string
	LatestSdkVersion                   string
	Link                               string
	LoggingToken                       string
	LogoURL                            string
	Migrations                         string
	MobileProfileSectionURL            string
	MobileWebURL                       string
	MonthlyActiveUsers                 string
	MonthlyActiveUsersRank             string
	Name                               string
	Namespace                          string
	ObjectStoreUrls                    string
	OwnerBusiness                      string
	PageTabDefaultName                 string
	PageTabURL                         string
	PhotoURL                           string
	PrivacyPolicyURL                   string
	ProfileSectionURL                  string
	PropertyID                         string
	ProtectedModeRules                 string
	RealTimeModeDevices                string
	Restrictions                       string
	RestrictiveDataFilterParams        string
	RestrictiveDataFilterRules         string
	SdkUpdateMessage                   string
	SeamlessLogin                      string
	SecureCanvasURL                    string
	SecurePageTabURL                   string
	ServerIPWhitelist                  string
	SmartLoginBookmarkIconURL          string
	SmartLoginMenuIconURL              string
	SocialDiscovery                    string
	Subcategory                        string
	SuggestedEventsSetting             string
	SupportedPlatforms                 string
	SupportsApprequestsFastAppSwitch   string
	SupportsAttribution                string
	SupportsImplicitSdkLogging         string
	SuppressNativeIosGdp               string
	TermsOfServiceURL                  string
	URLSchemeSuffix                    string
	UserSupportEmail                   string
	UserSupportURL                     string
	WebsiteURL                         string
	WeeklyActiveUsers                  string
}{
	AamRules:                           "aam_rules",
	AnAdSpaceLimit:                     "an_ad_space_limit",
	AnPlatforms:                        "an_platforms",
	AndroidKeyHash:                     "android_key_hash",
	AndroidSdkErrorCategories:          "android_sdk_error_categories",
	AppDomains:                         "app_domains",
	AppEventsConfig:                    "app_events_config",
	AppEventsFeatureBitmask:            "app_events_feature_bitmask",
	AppEventsSessionTimeout:            "app_events_session_timeout",
	AppInstallTracked:                  "app_install_tracked",
	AppName:                            "app_name",
	AppSignalsBindingIos:               "app_signals_binding_ios",
	AppType:                            "app_type",
	AuthDialogDataHelpURL:              "auth_dialog_data_help_url",
	AuthDialogHeadline:                 "auth_dialog_headline",
	AuthDialogPermsExplanation:         "auth_dialog_perms_explanation",
	AuthReferralDefaultActivityPrivacy: "auth_referral_default_activity_privacy",
	AuthReferralEnabled:                "auth_referral_enabled",
	AuthReferralExtendedPerms:          "auth_referral_extended_perms",
	AuthReferralFriendPerms:            "auth_referral_friend_perms",
	AuthReferralResponseType:           "auth_referral_response_type",
	AuthReferralUserPerms:              "auth_referral_user_perms",
	AutoEventMappingAndroid:            "auto_event_mapping_android",
	AutoEventMappingIos:                "auto_event_mapping_ios",
	AutoEventSetupEnabled:              "auto_event_setup_enabled",
	AutoLogAppEventsDefault:            "auto_log_app_events_default",
	AutoLogAppEventsEnabled:            "auto_log_app_events_enabled",
	Business:                           "business",
	CanvasFluidHeight:                  "canvas_fluid_height",
	CanvasFluidWidth:                   "canvas_fluid_width",
	CanvasURL:                          "canvas_url",
	Category:                           "category",
	ClientConfig:                       "client_config",
	Company:                            "company",
	ConfigIds:                          "config_ids",
	ConfiguredIosSso:                   "configured_ios_sso",
	ContactEmail:                       "contact_email",
	CreatedTime:                        "created_time",
	CreatorUID:                         "creator_uid",
	DailyActiveUsers:                   "daily_active_users",
	DailyActiveUsersRank:               "daily_active_users_rank",
	DeauthCallbackURL:                  "deauth_callback_url",
	DefaultShareMode:                   "default_share_mode",
	Description:                        "description",
	EnigmaConfig:                       "enigma_config",
	FinancialID:                        "financial_id",
	Gdpv4ChromeCustomTabsEnabled:       "gdpv4_chrome_custom_tabs_enabled",
	Gdpv4Enabled:                       "gdpv4_enabled",
	Gdpv4NuxContent:                    "gdpv4_nux_content",
	Gdpv4NuxEnabled:                    "gdpv4_nux_enabled",
	HasMessengerProduct:                "has_messenger_product",
	HostingURL:                         "hosting_url",
	IconURL:                            "icon_url",
	ID:                                 "id",
	IosBundleID:                        "ios_bundle_id",
	IosSdkDialogFlows:                  "ios_sdk_dialog_flows",
	IosSdkErrorCategories:              "ios_sdk_error_categories",
	IosSfvcAttr:                        "ios_sfvc_attr",
	IosSupportsNativeProxyAuthFlow:     "ios_supports_native_proxy_auth_flow",
	IosSupportsSystemAuth:              "ios_supports_system_auth",
	IpadAppStoreID:                     "ipad_app_store_id",
	IphoneAppStoreID:                   "iphone_app_store_id",
	LatestSdkVersion:                   "latest_sdk_version",
	Link:                               "link",
	LoggingToken:                       "logging_token",
	LogoURL:                            "logo_url",
	Migrations:                         "migrations",
	MobileProfileSectionURL:            "mobile_profile_section_url",
	MobileWebURL:                       "mobile_web_url",
	MonthlyActiveUsers:                 "monthly_active_users",
	MonthlyActiveUsersRank:             "monthly_active_users_rank",
	Name:                               "name",
	Namespace:                          "namespace",
	ObjectStoreUrls:                    "object_store_urls",
	OwnerBusiness:                      "owner_business",
	PageTabDefaultName:                 "page_tab_default_name",
	PageTabURL:                         "page_tab_url",
	PhotoURL:                           "photo_url",
	PrivacyPolicyURL:                   "privacy_policy_url",
	ProfileSectionURL:                  "profile_section_url",
	PropertyID:                         "property_id",
	ProtectedModeRules:                 "protected_mode_rules",
	RealTimeModeDevices:                "real_time_mode_devices",
	Restrictions:                       "restrictions",
	RestrictiveDataFilterParams:        "restrictive_data_filter_params",
	RestrictiveDataFilterRules:         "restrictive_data_filter_rules",
	SdkUpdateMessage:                   "sdk_update_message",
	SeamlessLogin:                      "seamless_login",
	SecureCanvasURL:                    "secure_canvas_url",
	SecurePageTabURL:                   "secure_page_tab_url",
	ServerIPWhitelist:                  "server_ip_whitelist",
	SmartLoginBookmarkIconURL:          "smart_login_bookmark_icon_url",
	SmartLoginMenuIconURL:              "smart_login_menu_icon_url",
	SocialDiscovery:                    "social_discovery",
	Subcategory:                        "subcategory",
	SuggestedEventsSetting:             "suggested_events_setting",
	SupportedPlatforms:                 "supported_platforms",
	SupportsApprequestsFastAppSwitch:   "supports_apprequests_fast_app_switch",
	SupportsAttribution:                "supports_attribution",
	SupportsImplicitSdkLogging:         "supports_implicit_sdk_logging",
	SuppressNativeIosGdp:               "suppress_native_ios_gdp",
	TermsOfServiceURL:                  "terms_of_service_url",
	URLSchemeSuffix:                    "url_scheme_suffix",
	UserSupportEmail:                   "user_support_email",
	UserSupportURL:                     "user_support_url",
	WebsiteURL:                         "website_url",
	WeeklyActiveUsers:                  "weekly_active_users",
}
