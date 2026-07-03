package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type DeleteApplicationAccountsParams struct {
	Type  *enums.ApplicationaccountsTypeEnumParam `facebook:"type"`
	UID   core.ID                                 `facebook:"uid"`
	Extra core.Params                             `facebook:"-"`
}

func (params DeleteApplicationAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	out["uid"] = params.UID
	return out
}

func DeleteApplicationAccounts(ctx context.Context, client *core.Client, id string, params DeleteApplicationAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "accounts"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAccountsParams struct {
	Type  *enums.ApplicationaccountsTypeEnumParam `facebook:"type"`
	Extra core.Params                             `facebook:"-"`
}

func (params GetApplicationAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetApplicationAccounts(ctx context.Context, client *core.Client, id string, params GetApplicationAccountsParams) (*core.Cursor[objects.TestAccount], error) {
	var out core.Cursor[objects.TestAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAccountsParams struct {
	Installed        *bool                                   `facebook:"installed"`
	Minor            *bool                                   `facebook:"minor"`
	Name             *string                                 `facebook:"name"`
	OwnerAccessToken *string                                 `facebook:"owner_access_token"`
	Permissions      *[]objects.Permission                   `facebook:"permissions"`
	Type             *enums.ApplicationaccountsTypeEnumParam `facebook:"type"`
	UID              *core.ID                                `facebook:"uid"`
	Extra            core.Params                             `facebook:"-"`
}

func (params CreateApplicationAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Installed != nil {
		out["installed"] = *params.Installed
	}
	if params.Minor != nil {
		out["minor"] = *params.Minor
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OwnerAccessToken != nil {
		out["owner_access_token"] = *params.OwnerAccessToken
	}
	if params.Permissions != nil {
		out["permissions"] = *params.Permissions
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	return out
}

func CreateApplicationAccounts(ctx context.Context, client *core.Client, id string, params CreateApplicationAccountsParams) (*objects.TestAccount, error) {
	var out objects.TestAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationActivitiesParams struct {
	AddToMessagingCustomerBaseForWhatsapp *string                                         `facebook:"add_to_messaging_customer_base_for_whatsapp"`
	AdvertiserID                          *core.ID                                        `facebook:"advertiser_id"`
	AdvertiserTrackingEnabled             *bool                                           `facebook:"advertiser_tracking_enabled"`
	AnonID                                *core.ID                                        `facebook:"anon_id"`
	AppUserID                             *core.ID                                        `facebook:"app_user_id"`
	ApplicationTrackingEnabled            *bool                                           `facebook:"application_tracking_enabled"`
	Attribution                           *string                                         `facebook:"attribution"`
	AttributionReferrer                   *string                                         `facebook:"attribution_referrer"`
	AttributionSources                    *[]map[string]interface{}                       `facebook:"attribution_sources"`
	AutoPublish                           *bool                                           `facebook:"auto_publish"`
	BundleID                              *core.ID                                        `facebook:"bundle_id"`
	BundleShortVersion                    *string                                         `facebook:"bundle_short_version"`
	BundleVersion                         *string                                         `facebook:"bundle_version"`
	CampaignIds                           *core.ID                                        `facebook:"campaign_ids"`
	CircuitBreakerTimeoutMs               *uint64                                         `facebook:"circuit_breaker_timeout_ms"`
	ClickID                               *core.ID                                        `facebook:"click_id"`
	ConsiderViews                         *bool                                           `facebook:"consider_views"`
	CustomEvents                          *[]map[string]interface{}                       `facebook:"custom_events"`
	CustomEventsFile                      *core.FileParam                                 `facebook:"custom_events_file"`
	DataProcessingOptions                 *[]string                                       `facebook:"data_processing_options"`
	DataProcessingOptionsCountry          *uint64                                         `facebook:"data_processing_options_country"`
	DataProcessingOptionsState            *uint64                                         `facebook:"data_processing_options_state"`
	DeviceToken                           *string                                         `facebook:"device_token"`
	Event                                 enums.ApplicationactivitiesEventEnumParam       `facebook:"event"`
	EventID                               *core.ID                                        `facebook:"event_id"`
	Extinfo                               *map[string]interface{}                         `facebook:"extinfo"`
	GoogleInstallReferrer                 *string                                         `facebook:"google_install_referrer"`
	IncludeDwellData                      *bool                                           `facebook:"include_dwell_data"`
	IncludeVideoData                      *bool                                           `facebook:"include_video_data"`
	InstallID                             *core.ID                                        `facebook:"install_id"`
	InstallReferrer                       *string                                         `facebook:"install_referrer"`
	InstallTimestamp                      *float64                                        `facebook:"install_timestamp"`
	InstallerPackage                      *string                                         `facebook:"installer_package"`
	IsCircuitBreakerActive                *bool                                           `facebook:"is_circuit_breaker_active"`
	IsFb                                  *bool                                           `facebook:"is_fb"`
	LimitedDataUse                        *bool                                           `facebook:"limited_data_use"`
	MetaInstallReferrer                   *string                                         `facebook:"meta_install_referrer"`
	MigrationBundle                       *string                                         `facebook:"migration_bundle"`
	OperationalParameters                 *[]map[string]interface{}                       `facebook:"operational_parameters"`
	PageID                                *core.ID                                        `facebook:"page_id"`
	PageScopedUserID                      *core.ID                                        `facebook:"page_scoped_user_id"`
	ReceiptData                           *string                                         `facebook:"receipt_data"`
	SdkVersion                            *string                                         `facebook:"sdk_version"`
	Ud                                    *map[string]interface{}                         `facebook:"ud"`
	URLSchemes                            *[]string                                       `facebook:"url_schemes"`
	UserID                                *core.ID                                        `facebook:"user_id"`
	UserIDType                            *enums.ApplicationactivitiesUserIDTypeEnumParam `facebook:"user_id_type"`
	VendorID                              *core.ID                                        `facebook:"vendor_id"`
	WindowsAttributionID                  *core.ID                                        `facebook:"windows_attribution_id"`
	Extra                                 core.Params                                     `facebook:"-"`
}

func (params CreateApplicationActivitiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AddToMessagingCustomerBaseForWhatsapp != nil {
		out["add_to_messaging_customer_base_for_whatsapp"] = *params.AddToMessagingCustomerBaseForWhatsapp
	}
	if params.AdvertiserID != nil {
		out["advertiser_id"] = *params.AdvertiserID
	}
	if params.AdvertiserTrackingEnabled != nil {
		out["advertiser_tracking_enabled"] = *params.AdvertiserTrackingEnabled
	}
	if params.AnonID != nil {
		out["anon_id"] = *params.AnonID
	}
	if params.AppUserID != nil {
		out["app_user_id"] = *params.AppUserID
	}
	if params.ApplicationTrackingEnabled != nil {
		out["application_tracking_enabled"] = *params.ApplicationTrackingEnabled
	}
	if params.Attribution != nil {
		out["attribution"] = *params.Attribution
	}
	if params.AttributionReferrer != nil {
		out["attribution_referrer"] = *params.AttributionReferrer
	}
	if params.AttributionSources != nil {
		out["attribution_sources"] = *params.AttributionSources
	}
	if params.AutoPublish != nil {
		out["auto_publish"] = *params.AutoPublish
	}
	if params.BundleID != nil {
		out["bundle_id"] = *params.BundleID
	}
	if params.BundleShortVersion != nil {
		out["bundle_short_version"] = *params.BundleShortVersion
	}
	if params.BundleVersion != nil {
		out["bundle_version"] = *params.BundleVersion
	}
	if params.CampaignIds != nil {
		out["campaign_ids"] = *params.CampaignIds
	}
	if params.CircuitBreakerTimeoutMs != nil {
		out["circuit_breaker_timeout_ms"] = *params.CircuitBreakerTimeoutMs
	}
	if params.ClickID != nil {
		out["click_id"] = *params.ClickID
	}
	if params.ConsiderViews != nil {
		out["consider_views"] = *params.ConsiderViews
	}
	if params.CustomEvents != nil {
		out["custom_events"] = *params.CustomEvents
	}
	if params.CustomEventsFile != nil {
		out["custom_events_file"] = *params.CustomEventsFile
	}
	if params.DataProcessingOptions != nil {
		out["data_processing_options"] = *params.DataProcessingOptions
	}
	if params.DataProcessingOptionsCountry != nil {
		out["data_processing_options_country"] = *params.DataProcessingOptionsCountry
	}
	if params.DataProcessingOptionsState != nil {
		out["data_processing_options_state"] = *params.DataProcessingOptionsState
	}
	if params.DeviceToken != nil {
		out["device_token"] = *params.DeviceToken
	}
	out["event"] = params.Event
	if params.EventID != nil {
		out["event_id"] = *params.EventID
	}
	if params.Extinfo != nil {
		out["extinfo"] = *params.Extinfo
	}
	if params.GoogleInstallReferrer != nil {
		out["google_install_referrer"] = *params.GoogleInstallReferrer
	}
	if params.IncludeDwellData != nil {
		out["include_dwell_data"] = *params.IncludeDwellData
	}
	if params.IncludeVideoData != nil {
		out["include_video_data"] = *params.IncludeVideoData
	}
	if params.InstallID != nil {
		out["install_id"] = *params.InstallID
	}
	if params.InstallReferrer != nil {
		out["install_referrer"] = *params.InstallReferrer
	}
	if params.InstallTimestamp != nil {
		out["install_timestamp"] = *params.InstallTimestamp
	}
	if params.InstallerPackage != nil {
		out["installer_package"] = *params.InstallerPackage
	}
	if params.IsCircuitBreakerActive != nil {
		out["is_circuit_breaker_active"] = *params.IsCircuitBreakerActive
	}
	if params.IsFb != nil {
		out["is_fb"] = *params.IsFb
	}
	if params.LimitedDataUse != nil {
		out["limited_data_use"] = *params.LimitedDataUse
	}
	if params.MetaInstallReferrer != nil {
		out["meta_install_referrer"] = *params.MetaInstallReferrer
	}
	if params.MigrationBundle != nil {
		out["migration_bundle"] = *params.MigrationBundle
	}
	if params.OperationalParameters != nil {
		out["operational_parameters"] = *params.OperationalParameters
	}
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	if params.PageScopedUserID != nil {
		out["page_scoped_user_id"] = *params.PageScopedUserID
	}
	if params.ReceiptData != nil {
		out["receipt_data"] = *params.ReceiptData
	}
	if params.SdkVersion != nil {
		out["sdk_version"] = *params.SdkVersion
	}
	if params.Ud != nil {
		out["ud"] = *params.Ud
	}
	if params.URLSchemes != nil {
		out["url_schemes"] = *params.URLSchemes
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	if params.UserIDType != nil {
		out["user_id_type"] = *params.UserIDType
	}
	if params.VendorID != nil {
		out["vendor_id"] = *params.VendorID
	}
	if params.WindowsAttributionID != nil {
		out["windows_attribution_id"] = *params.WindowsAttributionID
	}
	return out
}

func CreateApplicationActivities(ctx context.Context, client *core.Client, id string, params CreateApplicationActivitiesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "activities"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAdPlacementGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAdPlacementGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAdPlacementGroups(ctx context.Context, client *core.Client, id string, params GetApplicationAdPlacementGroupsParams) (*core.Cursor[objects.AdPlacementGroup], error) {
	var out core.Cursor[objects.AdPlacementGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_placement_groups"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAdnetworkPlacementsParams struct {
	RequestID *core.ID    `facebook:"request_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetApplicationAdnetworkPlacementsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.RequestID != nil {
		out["request_id"] = *params.RequestID
	}
	return out
}

func GetApplicationAdnetworkPlacements(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkPlacementsParams) (*core.Cursor[objects.AdPlacement], error) {
	var out core.Cursor[objects.AdPlacement]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetwork_placements"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAdnetworkanalyticsParams struct {
	AggregationPeriod  *enums.ApplicationadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns         *[]enums.ApplicationadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters            *[]map[string]interface{}                                      `facebook:"filters"`
	Limit              *uint64                                                        `facebook:"limit"`
	Metrics            []enums.ApplicationadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn     *enums.ApplicationadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType       *enums.ApplicationadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	ShouldIncludeUntil *bool                                                          `facebook:"should_include_until"`
	Since              *time.Time                                                     `facebook:"since"`
	Until              *time.Time                                                     `facebook:"until"`
	Extra              core.Params                                                    `facebook:"-"`
}

func (params GetApplicationAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.ShouldIncludeUntil != nil {
		out["should_include_until"] = *params.ShouldIncludeUntil
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetApplicationAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAdnetworkanalyticsParams struct {
	AggregationPeriod *enums.ApplicationadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns        *[]enums.ApplicationadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters           *[]map[string]interface{}                                      `facebook:"filters"`
	Limit             *int                                                           `facebook:"limit"`
	Metrics           []enums.ApplicationadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn    *enums.ApplicationadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType      *enums.ApplicationadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	Since             *time.Time                                                     `facebook:"since"`
	Until             *time.Time                                                     `facebook:"until"`
	Extra             core.Params                                                    `facebook:"-"`
}

func (params CreateApplicationAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func CreateApplicationAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateApplicationAdnetworkanalyticsParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAdnetworkanalyticsResultsParams struct {
	QueryIds *[]core.ID  `facebook:"query_ids"`
	Extra    core.Params `facebook:"-"`
}

func (params GetApplicationAdnetworkanalyticsResultsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.QueryIds != nil {
		out["query_ids"] = *params.QueryIds
	}
	return out
}

func GetApplicationAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAemAttributionParams struct {
	AdvertiserIds *[]core.ID  `facebook:"advertiser_ids"`
	FbContentData *string     `facebook:"fb_content_data"`
	Extra         core.Params `facebook:"-"`
}

func (params GetApplicationAemAttributionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdvertiserIds != nil {
		out["advertiser_ids"] = *params.AdvertiserIds
	}
	if params.FbContentData != nil {
		out["fb_content_data"] = *params.FbContentData
	}
	return out
}

func GetApplicationAemAttribution(ctx context.Context, client *core.Client, id string, params GetApplicationAemAttributionParams) (*core.Cursor[objects.AEMAttribution], error) {
	var out core.Cursor[objects.AEMAttribution]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "aem_attribution"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAemConversionConfigsParams struct {
	AdvertiserIds *[]core.ID  `facebook:"advertiser_ids"`
	Extra         core.Params `facebook:"-"`
}

func (params GetApplicationAemConversionConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdvertiserIds != nil {
		out["advertiser_ids"] = *params.AdvertiserIds
	}
	return out
}

func GetApplicationAemConversionConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionConfigsParams) (*core.Cursor[objects.AEMConversionConfig], error) {
	var out core.Cursor[objects.AEMConversionConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "aem_conversion_configs"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAemConversionFilterParams struct {
	CatalogID    *core.ID    `facebook:"catalog_id"`
	FbContentIds *core.ID    `facebook:"fb_content_ids"`
	Extra        core.Params `facebook:"-"`
}

func (params GetApplicationAemConversionFilterParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CatalogID != nil {
		out["catalog_id"] = *params.CatalogID
	}
	if params.FbContentIds != nil {
		out["fb_content_ids"] = *params.FbContentIds
	}
	return out
}

func GetApplicationAemConversionFilter(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionFilterParams) (*core.Cursor[objects.AEMConversionFilter], error) {
	var out core.Cursor[objects.AEMConversionFilter]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "aem_conversion_filter"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAemConversionsParams struct {
	AemConversions []map[string]interface{} `facebook:"aem_conversions"`
	Extra          core.Params              `facebook:"-"`
}

func (params CreateApplicationAemConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["aem_conversions"] = params.AemConversions
	return out
}

func CreateApplicationAemConversions(ctx context.Context, client *core.Client, id string, params CreateApplicationAemConversionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "aem_conversions"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAemSkanReadinessParams struct {
	AppID                core.ID     `facebook:"app_id"`
	IsAemReady           *bool       `facebook:"is_aem_ready"`
	IsAppAemInstallReady *bool       `facebook:"is_app_aem_install_ready"`
	IsAppAemReady        *bool       `facebook:"is_app_aem_ready"`
	IsSkanReady          *bool       `facebook:"is_skan_ready"`
	Message              *string     `facebook:"message"`
	Extra                core.Params `facebook:"-"`
}

func (params CreateApplicationAemSkanReadinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_id"] = params.AppID
	if params.IsAemReady != nil {
		out["is_aem_ready"] = *params.IsAemReady
	}
	if params.IsAppAemInstallReady != nil {
		out["is_app_aem_install_ready"] = *params.IsAppAemInstallReady
	}
	if params.IsAppAemReady != nil {
		out["is_app_aem_ready"] = *params.IsAppAemReady
	}
	if params.IsSkanReady != nil {
		out["is_skan_ready"] = *params.IsSkanReady
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateApplicationAemSkanReadiness(ctx context.Context, client *core.Client, id string, params CreateApplicationAemSkanReadinessParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "aem_skan_readiness"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAgencies(ctx context.Context, client *core.Client, id string, params GetApplicationAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAggregateRevenueParams struct {
	Ecpms     *[]string   `facebook:"ecpms"`
	QueryIds  *[]core.ID  `facebook:"query_ids"`
	RequestID *core.ID    `facebook:"request_id"`
	SyncAPI   *bool       `facebook:"sync_api"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateApplicationAggregateRevenueParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Ecpms != nil {
		out["ecpms"] = *params.Ecpms
	}
	if params.QueryIds != nil {
		out["query_ids"] = *params.QueryIds
	}
	if params.RequestID != nil {
		out["request_id"] = *params.RequestID
	}
	if params.SyncAPI != nil {
		out["sync_api"] = *params.SyncAPI
	}
	return out
}

func CreateApplicationAggregateRevenue(ctx context.Context, client *core.Client, id string, params CreateApplicationAggregateRevenueParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "aggregate_revenue"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAndroidDialogConfigsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAndroidDialogConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAndroidDialogConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationAndroidDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	var out core.Cursor[objects.ApplicationDialogConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "android_dialog_configs"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAppCapiSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAppCapiSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAppCapiSettings(ctx context.Context, client *core.Client, id string, params GetApplicationAppCapiSettingsParams) (*core.Cursor[objects.ApplicationAppCapiSetting], error) {
	var out core.Cursor[objects.ApplicationAppCapiSetting]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "app_capi_settings"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAppEventTypesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAppEventTypesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAppEventTypes(ctx context.Context, client *core.Client, id string, params GetApplicationAppEventTypesParams) (*core.Cursor[objects.ApplicationAppEventTypes], error) {
	var out core.Cursor[objects.ApplicationAppEventTypes]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "app_event_types"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAppIndexingParams struct {
	AppVersion      string                                            `facebook:"app_version"`
	DeviceSessionID *core.ID                                          `facebook:"device_session_id"`
	ExtraInfo       *string                                           `facebook:"extra_info"`
	Platform        enums.ApplicationappIndexingPlatformEnumParam     `facebook:"platform"`
	RequestType     *enums.ApplicationappIndexingRequestTypeEnumParam `facebook:"request_type"`
	Tree            map[string]interface{}                            `facebook:"tree"`
	Extra           core.Params                                       `facebook:"-"`
}

func (params CreateApplicationAppIndexingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_version"] = params.AppVersion
	if params.DeviceSessionID != nil {
		out["device_session_id"] = *params.DeviceSessionID
	}
	if params.ExtraInfo != nil {
		out["extra_info"] = *params.ExtraInfo
	}
	out["platform"] = params.Platform
	if params.RequestType != nil {
		out["request_type"] = *params.RequestType
	}
	out["tree"] = params.Tree
	return out
}

func CreateApplicationAppIndexing(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "app_indexing"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAppIndexingSessionParams struct {
	DeviceSessionID core.ID     `facebook:"device_session_id"`
	Extinfo         *string     `facebook:"extinfo"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateApplicationAppIndexingSessionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["device_session_id"] = params.DeviceSessionID
	if params.Extinfo != nil {
		out["extinfo"] = *params.Extinfo
	}
	return out
}

func CreateApplicationAppIndexingSession(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingSessionParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "app_indexing_session"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAppInstalledGroupsParams struct {
	GroupID *core.ID    `facebook:"group_id"`
	Extra   core.Params `facebook:"-"`
}

func (params GetApplicationAppInstalledGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.GroupID != nil {
		out["group_id"] = *params.GroupID
	}
	return out
}

func GetApplicationAppInstalledGroups(ctx context.Context, client *core.Client, id string, params GetApplicationAppInstalledGroupsParams) (*core.Cursor[objects.Group], error) {
	var out core.Cursor[objects.Group]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "app_installed_groups"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAppPushDeviceTokenParams struct {
	DeviceID    core.ID                                               `facebook:"device_id"`
	DeviceToken string                                                `facebook:"device_token"`
	Platform    *enums.ApplicationappPushDeviceTokenPlatformEnumParam `facebook:"platform"`
	Extra       core.Params                                           `facebook:"-"`
}

func (params CreateApplicationAppPushDeviceTokenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["device_id"] = params.DeviceID
	out["device_token"] = params.DeviceToken
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	return out
}

func CreateApplicationAppPushDeviceToken(ctx context.Context, client *core.Client, id string, params CreateApplicationAppPushDeviceTokenParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "app_push_device_token"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAppassetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationAppassetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationAppassets(ctx context.Context, client *core.Client, id string, params GetApplicationAppassetsParams) (*core.Cursor[objects.CanvasAppAsset], error) {
	var out core.Cursor[objects.CanvasAppAsset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "appassets"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationAssetsParams struct {
	Asset   core.FileParam `facebook:"asset"`
	Comment *string        `facebook:"comment"`
	Type    string         `facebook:"type"`
	Extra   core.Params    `facebook:"-"`
}

func (params CreateApplicationAssetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset"] = params.Asset
	if params.Comment != nil {
		out["comment"] = *params.Comment
	}
	out["type"] = params.Type
	return out
}

func CreateApplicationAssets(ctx context.Context, client *core.Client, id string, params CreateApplicationAssetsParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assets"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationAuthorizedAdaccountsParams struct {
	Business *string     `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetApplicationAuthorizedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Business != nil {
		out["business"] = *params.Business
	}
	return out
}

func GetApplicationAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params GetApplicationAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationButtonAutoDetectionDeviceSelectionParams struct {
	DeviceID *core.ID    `facebook:"device_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetApplicationButtonAutoDetectionDeviceSelectionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DeviceID != nil {
		out["device_id"] = *params.DeviceID
	}
	return out
}

func GetApplicationButtonAutoDetectionDeviceSelection(ctx context.Context, client *core.Client, id string, params GetApplicationButtonAutoDetectionDeviceSelectionParams) (*core.Cursor[objects.ButtonAutoDetectionDeviceSelection], error) {
	var out core.Cursor[objects.ButtonAutoDetectionDeviceSelection]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "button_auto_detection_device_selection"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationCloudbridgeSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationCloudbridgeSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationCloudbridgeSettings(ctx context.Context, client *core.Client, id string, params GetApplicationCloudbridgeSettingsParams) (*core.Cursor[objects.CloudbridgeSetting], error) {
	var out core.Cursor[objects.CloudbridgeSetting]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "cloudbridge_settings"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationCodelessEventMappingsParams struct {
	Mappings       []map[string]interface{}                                      `facebook:"mappings"`
	MutationMethod enums.ApplicationcodelessEventMappingsMutationMethodEnumParam `facebook:"mutation_method"`
	Platform       enums.ApplicationcodelessEventMappingsPlatformEnumParam       `facebook:"platform"`
	PostMethod     *enums.ApplicationcodelessEventMappingsPostMethodEnumParam    `facebook:"post_method"`
	Extra          core.Params                                                   `facebook:"-"`
}

func (params CreateApplicationCodelessEventMappingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["mappings"] = params.Mappings
	out["mutation_method"] = params.MutationMethod
	out["platform"] = params.Platform
	if params.PostMethod != nil {
		out["post_method"] = *params.PostMethod
	}
	return out
}

func CreateApplicationCodelessEventMappings(ctx context.Context, client *core.Client, id string, params CreateApplicationCodelessEventMappingsParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "codeless_event_mappings"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationConnectedClientBusinessesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationConnectedClientBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationConnectedClientBusinesses(ctx context.Context, client *core.Client, id string, params GetApplicationConnectedClientBusinessesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "connected_client_businesses"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationDaChecksParams struct {
	Checks           *[]string                                           `facebook:"checks"`
	ConnectionMethod *enums.ApplicationdaChecksConnectionMethodEnumParam `facebook:"connection_method"`
	Extra            core.Params                                         `facebook:"-"`
}

func (params GetApplicationDaChecksParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Checks != nil {
		out["checks"] = *params.Checks
	}
	if params.ConnectionMethod != nil {
		out["connection_method"] = *params.ConnectionMethod
	}
	return out
}

func GetApplicationDaChecks(ctx context.Context, client *core.Client, id string, params GetApplicationDaChecksParams) (*core.Cursor[objects.DACheck], error) {
	var out core.Cursor[objects.DACheck]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "da_checks"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationDomainReportsParams struct {
	TrackingDomains []string    `facebook:"tracking_domains"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateApplicationDomainReportsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tracking_domains"] = params.TrackingDomains
	return out
}

func CreateApplicationDomainReports(ctx context.Context, client *core.Client, id string, params CreateApplicationDomainReportsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "domain_reports"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationIapPurchasesParams struct {
	OrderID core.ID     `facebook:"order_id"`
	Extra   core.Params `facebook:"-"`
}

func (params GetApplicationIapPurchasesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["order_id"] = params.OrderID
	return out
}

func GetApplicationIapPurchases(ctx context.Context, client *core.Client, id string, params GetApplicationIapPurchasesParams) (*core.Cursor[objects.DCPAppStoreOrder], error) {
	var out core.Cursor[objects.DCPAppStoreOrder]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "iap_purchases"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationIosDialogConfigsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationIosDialogConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationIosDialogConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationIosDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	var out core.Cursor[objects.ApplicationDialogConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ios_dialog_configs"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationLinkedDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationLinkedDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationLinkedDataset(ctx context.Context, client *core.Client, id string, params GetApplicationLinkedDatasetParams) (*core.Cursor[objects.AdsDataset], error) {
	var out core.Cursor[objects.AdsDataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "linked_dataset"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationMmpAuditingParams struct {
	AdvertiserID          *core.ID                  `facebook:"advertiser_id"`
	Attribution           *string                   `facebook:"attribution"`
	AttributionMethod     *string                   `facebook:"attribution_method"`
	AttributionModel      *string                   `facebook:"attribution_model"`
	AttributionReferrer   *string                   `facebook:"attribution_referrer"`
	AuditingToken         *string                   `facebook:"auditing_token"`
	ClickAttrWindow       *uint64                   `facebook:"click_attr_window"`
	CustomEvents          *[]map[string]interface{} `facebook:"custom_events"`
	DeclineReason         *string                   `facebook:"decline_reason"`
	DeviceOs              *string                   `facebook:"device_os"`
	EngagementType        *string                   `facebook:"engagement_type"`
	Event                 string                    `facebook:"event"`
	EventID               *core.ID                  `facebook:"event_id"`
	EventReportedTime     *uint64                   `facebook:"event_reported_time"`
	FbAdID                *core.ID                  `facebook:"fb_ad_id"`
	FbAdgroupID           *core.ID                  `facebook:"fb_adgroup_id"`
	FbClickTime           *uint64                   `facebook:"fb_click_time"`
	FbViewTime            *uint64                   `facebook:"fb_view_time"`
	GoogleInstallReferrer *string                   `facebook:"google_install_referrer"`
	InactivityWindowHours *uint64                   `facebook:"inactivity_window_hours"`
	InstallID             *core.ID                  `facebook:"install_id"`
	IsFb                  bool                      `facebook:"is_fb"`
	MetaInstallReferrer   *string                   `facebook:"meta_install_referrer"`
	UsedInstallReferrer   *bool                     `facebook:"used_install_referrer"`
	ViewAttrWindow        *uint64                   `facebook:"view_attr_window"`
	Extra                 core.Params               `facebook:"-"`
}

func (params CreateApplicationMmpAuditingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdvertiserID != nil {
		out["advertiser_id"] = *params.AdvertiserID
	}
	if params.Attribution != nil {
		out["attribution"] = *params.Attribution
	}
	if params.AttributionMethod != nil {
		out["attribution_method"] = *params.AttributionMethod
	}
	if params.AttributionModel != nil {
		out["attribution_model"] = *params.AttributionModel
	}
	if params.AttributionReferrer != nil {
		out["attribution_referrer"] = *params.AttributionReferrer
	}
	if params.AuditingToken != nil {
		out["auditing_token"] = *params.AuditingToken
	}
	if params.ClickAttrWindow != nil {
		out["click_attr_window"] = *params.ClickAttrWindow
	}
	if params.CustomEvents != nil {
		out["custom_events"] = *params.CustomEvents
	}
	if params.DeclineReason != nil {
		out["decline_reason"] = *params.DeclineReason
	}
	if params.DeviceOs != nil {
		out["device_os"] = *params.DeviceOs
	}
	if params.EngagementType != nil {
		out["engagement_type"] = *params.EngagementType
	}
	out["event"] = params.Event
	if params.EventID != nil {
		out["event_id"] = *params.EventID
	}
	if params.EventReportedTime != nil {
		out["event_reported_time"] = *params.EventReportedTime
	}
	if params.FbAdID != nil {
		out["fb_ad_id"] = *params.FbAdID
	}
	if params.FbAdgroupID != nil {
		out["fb_adgroup_id"] = *params.FbAdgroupID
	}
	if params.FbClickTime != nil {
		out["fb_click_time"] = *params.FbClickTime
	}
	if params.FbViewTime != nil {
		out["fb_view_time"] = *params.FbViewTime
	}
	if params.GoogleInstallReferrer != nil {
		out["google_install_referrer"] = *params.GoogleInstallReferrer
	}
	if params.InactivityWindowHours != nil {
		out["inactivity_window_hours"] = *params.InactivityWindowHours
	}
	if params.InstallID != nil {
		out["install_id"] = *params.InstallID
	}
	out["is_fb"] = params.IsFb
	if params.MetaInstallReferrer != nil {
		out["meta_install_referrer"] = *params.MetaInstallReferrer
	}
	if params.UsedInstallReferrer != nil {
		out["used_install_referrer"] = *params.UsedInstallReferrer
	}
	if params.ViewAttrWindow != nil {
		out["view_attr_window"] = *params.ViewAttrWindow
	}
	return out
}

func CreateApplicationMmpAuditing(ctx context.Context, client *core.Client, id string, params CreateApplicationMmpAuditingParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "mmp_auditing"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationMobileSdkGkParams struct {
	DeviceID   *core.ID                                      `facebook:"device_id"`
	Extinfo    *map[string]interface{}                       `facebook:"extinfo"`
	OsVersion  *string                                       `facebook:"os_version"`
	Platform   enums.ApplicationmobileSdkGkPlatformEnumParam `facebook:"platform"`
	SdkVersion string                                        `facebook:"sdk_version"`
	Extra      core.Params                                   `facebook:"-"`
}

func (params GetApplicationMobileSdkGkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DeviceID != nil {
		out["device_id"] = *params.DeviceID
	}
	if params.Extinfo != nil {
		out["extinfo"] = *params.Extinfo
	}
	if params.OsVersion != nil {
		out["os_version"] = *params.OsVersion
	}
	out["platform"] = params.Platform
	out["sdk_version"] = params.SdkVersion
	return out
}

func GetApplicationMobileSdkGk(ctx context.Context, client *core.Client, id string, params GetApplicationMobileSdkGkParams) (*core.Cursor[objects.MobileSdkGk], error) {
	var out core.Cursor[objects.MobileSdkGk]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "mobile_sdk_gk"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationMonetizedDigitalStoreObjectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationMonetizedDigitalStoreObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationMonetizedDigitalStoreObjects(ctx context.Context, client *core.Client, id string, params GetApplicationMonetizedDigitalStoreObjectsParams) (*core.Cursor[objects.MonetizedDigitalStoreObject], error) {
	var out core.Cursor[objects.MonetizedDigitalStoreObject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "monetized_digital_store_objects"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationMonetizedDigitalStoreObjectsParams struct {
	ContentID core.ID     `facebook:"content_id"`
	Store     string      `facebook:"store"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateApplicationMonetizedDigitalStoreObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["content_id"] = params.ContentID
	out["store"] = params.Store
	return out
}

func CreateApplicationMonetizedDigitalStoreObjects(ctx context.Context, client *core.Client, id string, params CreateApplicationMonetizedDigitalStoreObjectsParams) (*objects.MonetizedDigitalStoreObject, error) {
	var out objects.MonetizedDigitalStoreObject
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "monetized_digital_store_objects"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationObjectTypesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationObjectTypesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationObjectTypes(ctx context.Context, client *core.Client, id string, params GetApplicationObjectTypesParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "object_types"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationObjectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationObjects(ctx context.Context, client *core.Client, id string, params GetApplicationObjectsParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "objects"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationOccludespopupsParams struct {
	Flash *bool       `facebook:"flash"`
	Unity *bool       `facebook:"unity"`
	Extra core.Params `facebook:"-"`
}

func (params CreateApplicationOccludespopupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Flash != nil {
		out["flash"] = *params.Flash
	}
	if params.Unity != nil {
		out["unity"] = *params.Unity
	}
	return out
}

func CreateApplicationOccludespopups(ctx context.Context, client *core.Client, id string, params CreateApplicationOccludespopupsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "occludespopups"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationPermissionsParams struct {
	AndroidKeyHash *string                                        `facebook:"android_key_hash"`
	IosBundleID    *core.ID                                       `facebook:"ios_bundle_id"`
	Permission     *[]objects.Permission                          `facebook:"permission"`
	ProxiedAppID   *core.ID                                       `facebook:"proxied_app_id"`
	Status         *[]enums.ApplicationpermissionsStatusEnumParam `facebook:"status"`
	Extra          core.Params                                    `facebook:"-"`
}

func (params GetApplicationPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.Permission != nil {
		out["permission"] = *params.Permission
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetApplicationPermissions(ctx context.Context, client *core.Client, id string, params GetApplicationPermissionsParams) (*core.Cursor[objects.ApplicationPermission], error) {
	var out core.Cursor[objects.ApplicationPermission]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "permissions"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationProductsParams struct {
	ProductIds *[]core.ID  `facebook:"product_ids"`
	Extra      core.Params `facebook:"-"`
}

func (params GetApplicationProductsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ProductIds != nil {
		out["product_ids"] = *params.ProductIds
	}
	return out
}

func GetApplicationProducts(ctx context.Context, client *core.Client, id string, params GetApplicationProductsParams) (*core.Cursor[objects.GamesIAPProduct], error) {
	var out core.Cursor[objects.GamesIAPProduct]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationPurchasesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationPurchasesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationPurchases(ctx context.Context, client *core.Client, id string, params GetApplicationPurchasesParams) (*core.Cursor[objects.GamesIAPOrder], error) {
	var out core.Cursor[objects.GamesIAPOrder]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "purchases"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationRolesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationRolesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationRoles(ctx context.Context, client *core.Client, id string, params GetApplicationRolesParams) (*core.Cursor[objects.ApplicationRoles], error) {
	var out core.Cursor[objects.ApplicationRoles]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "roles"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationServerDomainInfosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationServerDomainInfosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationServerDomainInfos(ctx context.Context, client *core.Client, id string, params GetApplicationServerDomainInfosParams) (*core.Cursor[objects.ApplicationServerDomainInfos], error) {
	var out core.Cursor[objects.ApplicationServerDomainInfos]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "server_domain_infos"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationSgwDatasetStatusParams struct {
	DatasetID core.ID     `facebook:"dataset_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetApplicationSgwDatasetStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["dataset_id"] = params.DatasetID
	return out
}

func GetApplicationSgwDatasetStatus(ctx context.Context, client *core.Client, id string, params GetApplicationSgwDatasetStatusParams) (*core.Cursor[objects.ApplicationSgwDatasetStatus], error) {
	var out core.Cursor[objects.ApplicationSgwDatasetStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sgw_dataset_status"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationSgwInstallDeferralLinkParams struct {
	ClientIP  *string     `facebook:"client_ip"`
	DatasetID core.ID     `facebook:"dataset_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetApplicationSgwInstallDeferralLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ClientIP != nil {
		out["client_ip"] = *params.ClientIP
	}
	out["dataset_id"] = params.DatasetID
	return out
}

func GetApplicationSgwInstallDeferralLink(ctx context.Context, client *core.Client, id string, params GetApplicationSgwInstallDeferralLinkParams) (*core.Cursor[objects.ApplicationSgwInstallDeferralLink], error) {
	var out core.Cursor[objects.ApplicationSgwInstallDeferralLink]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sgw_install_deferral_link"), params.ToParams(), &out)
	return &out, err
}

type DeleteApplicationSubscriptionsParams struct {
	Fields *[]string   `facebook:"fields"`
	Object *string     `facebook:"object"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteApplicationSubscriptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Object != nil {
		out["object"] = *params.Object
	}
	return out
}

func DeleteApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params DeleteApplicationSubscriptionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "subscriptions"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationSubscriptionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetApplicationSubscriptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params GetApplicationSubscriptionsParams) (*core.Cursor[objects.ApplicationSubscriptions], error) {
	var out core.Cursor[objects.ApplicationSubscriptions]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "subscriptions"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationSubscriptionsParams struct {
	CallbackURL              *string     `facebook:"callback_url"`
	Fields                   *[]string   `facebook:"fields"`
	IncludeClientCertificate *bool       `facebook:"include_client_certificate"`
	IncludeValues            *bool       `facebook:"include_values"`
	Object                   string      `facebook:"object"`
	VerifyToken              *string     `facebook:"verify_token"`
	Extra                    core.Params `facebook:"-"`
}

func (params CreateApplicationSubscriptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CallbackURL != nil {
		out["callback_url"] = *params.CallbackURL
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.IncludeClientCertificate != nil {
		out["include_client_certificate"] = *params.IncludeClientCertificate
	}
	if params.IncludeValues != nil {
		out["include_values"] = *params.IncludeValues
	}
	out["object"] = params.Object
	if params.VerifyToken != nil {
		out["verify_token"] = *params.VerifyToken
	}
	return out
}

func CreateApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params CreateApplicationSubscriptionsParams) (*objects.ApplicationSubscriptions, error) {
	var out objects.ApplicationSubscriptions
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "subscriptions"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationThreatPrivacyGroupsMemberParams struct {
	Description *string     `facebook:"description"`
	GroupID     *core.ID    `facebook:"group_id"`
	Name        *string     `facebook:"name"`
	Extra       core.Params `facebook:"-"`
}

func (params GetApplicationThreatPrivacyGroupsMemberParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GroupID != nil {
		out["group_id"] = *params.GroupID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func GetApplicationThreatPrivacyGroupsMember(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsMemberParams) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	var out core.Cursor[objects.ThreatPrivacyGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "threat_privacy_groups_member"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationThreatPrivacyGroupsOwnerParams struct {
	Description *string     `facebook:"description"`
	GroupID     *core.ID    `facebook:"group_id"`
	Name        *string     `facebook:"name"`
	Extra       core.Params `facebook:"-"`
}

func (params GetApplicationThreatPrivacyGroupsOwnerParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GroupID != nil {
		out["group_id"] = *params.GroupID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func GetApplicationThreatPrivacyGroupsOwner(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsOwnerParams) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	var out core.Cursor[objects.ThreatPrivacyGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "threat_privacy_groups_owner"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationUploadsParams struct {
	FileLength  *uint64                                       `facebook:"file_length"`
	FileName    *map[string]interface{}                       `facebook:"file_name"`
	FileType    *map[string]interface{}                       `facebook:"file_type"`
	SessionType *enums.ApplicationuploadsSessionTypeEnumParam `facebook:"session_type"`
	Extra       core.Params                                   `facebook:"-"`
}

func (params CreateApplicationUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.FileLength != nil {
		out["file_length"] = *params.FileLength
	}
	if params.FileName != nil {
		out["file_name"] = *params.FileName
	}
	if params.FileType != nil {
		out["file_type"] = *params.FileType
	}
	if params.SessionType != nil {
		out["session_type"] = *params.SessionType
	}
	return out
}

func CreateApplicationUploads(ctx context.Context, client *core.Client, id string, params CreateApplicationUploadsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "uploads"), params.ToParams(), &out)
	return &out, err
}

type CreateApplicationWhatsappBusinessSolutionParams struct {
	OwnerPermissions   []enums.ApplicationwhatsappBusinessSolutionOwnerPermissionsEnumParam   `facebook:"owner_permissions"`
	PartnerAppID       core.ID                                                                `facebook:"partner_app_id"`
	PartnerPermissions []enums.ApplicationwhatsappBusinessSolutionPartnerPermissionsEnumParam `facebook:"partner_permissions"`
	SolutionName       string                                                                 `facebook:"solution_name"`
	Extra              core.Params                                                            `facebook:"-"`
}

func (params CreateApplicationWhatsappBusinessSolutionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["owner_permissions"] = params.OwnerPermissions
	out["partner_app_id"] = params.PartnerAppID
	out["partner_permissions"] = params.PartnerPermissions
	out["solution_name"] = params.SolutionName
	return out
}

func CreateApplicationWhatsappBusinessSolution(ctx context.Context, client *core.Client, id string, params CreateApplicationWhatsappBusinessSolutionParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "whatsapp_business_solution"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationWhatsappBusinessSolutionsParams struct {
	Role  *enums.ApplicationwhatsappBusinessSolutionsRoleEnumParam `facebook:"role"`
	Extra core.Params                                              `facebook:"-"`
}

func (params GetApplicationWhatsappBusinessSolutionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Role != nil {
		out["role"] = *params.Role
	}
	return out
}

func GetApplicationWhatsappBusinessSolutions(ctx context.Context, client *core.Client, id string, params GetApplicationWhatsappBusinessSolutionsParams) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "whatsapp_business_solutions"), params.ToParams(), &out)
	return &out, err
}

type GetApplicationParams struct {
	AdvertiserID *core.ID    `facebook:"advertiser_id"`
	Extra        core.Params `facebook:"-"`
}

func (params GetApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdvertiserID != nil {
		out["advertiser_id"] = *params.AdvertiserID
	}
	return out
}

func GetApplication(ctx context.Context, client *core.Client, id string, params GetApplicationParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateApplicationParams struct {
	AllowCycleAppSecret        *bool                           `facebook:"allow_cycle_app_secret"`
	AnPlatforms                *[]enums.ApplicationAnPlatforms `facebook:"an_platforms"`
	AppDomains                 *[]string                       `facebook:"app_domains"`
	AppName                    *string                         `facebook:"app_name"`
	AppType                    *bool                           `facebook:"app_type"`
	AuthDialogHeadline         *string                         `facebook:"auth_dialog_headline"`
	AuthDialogPermsExplanation *string                         `facebook:"auth_dialog_perms_explanation"`
	AuthReferralEnabled        *bool                           `facebook:"auth_referral_enabled"`
	AuthReferralExtendedPerms  *[]string                       `facebook:"auth_referral_extended_perms"`
	AuthReferralFriendPerms    *[]string                       `facebook:"auth_referral_friend_perms"`
	AuthReferralResponseType   *string                         `facebook:"auth_referral_response_type"`
	AuthReferralUserPerms      *[]string                       `facebook:"auth_referral_user_perms"`
	CanvasFluidHeight          *bool                           `facebook:"canvas_fluid_height"`
	CanvasFluidWidth           *bool                           `facebook:"canvas_fluid_width"`
	CanvasURL                  *string                         `facebook:"canvas_url"`
	ContactEmail               *string                         `facebook:"contact_email"`
	DeauthCallbackURL          *string                         `facebook:"deauth_callback_url"`
	MobileWebURL               *string                         `facebook:"mobile_web_url"`
	Namespace                  *string                         `facebook:"namespace"`
	PageTabDefaultName         *string                         `facebook:"page_tab_default_name"`
	PrivacyPolicyURL           *string                         `facebook:"privacy_policy_url"`
	Restrictions               *string                         `facebook:"restrictions"`
	SecureCanvasURL            *string                         `facebook:"secure_canvas_url"`
	SecurePageTabURL           *string                         `facebook:"secure_page_tab_url"`
	ServerIPWhitelist          *[]string                       `facebook:"server_ip_whitelist"`
	TermsOfServiceURL          *string                         `facebook:"terms_of_service_url"`
	URLSchemeSuffix            *string                         `facebook:"url_scheme_suffix"`
	UserSupportEmail           *string                         `facebook:"user_support_email"`
	UserSupportURL             *string                         `facebook:"user_support_url"`
	WebsiteURL                 *string                         `facebook:"website_url"`
	Extra                      core.Params                     `facebook:"-"`
}

func (params UpdateApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowCycleAppSecret != nil {
		out["allow_cycle_app_secret"] = *params.AllowCycleAppSecret
	}
	if params.AnPlatforms != nil {
		out["an_platforms"] = *params.AnPlatforms
	}
	if params.AppDomains != nil {
		out["app_domains"] = *params.AppDomains
	}
	if params.AppName != nil {
		out["app_name"] = *params.AppName
	}
	if params.AppType != nil {
		out["app_type"] = *params.AppType
	}
	if params.AuthDialogHeadline != nil {
		out["auth_dialog_headline"] = *params.AuthDialogHeadline
	}
	if params.AuthDialogPermsExplanation != nil {
		out["auth_dialog_perms_explanation"] = *params.AuthDialogPermsExplanation
	}
	if params.AuthReferralEnabled != nil {
		out["auth_referral_enabled"] = *params.AuthReferralEnabled
	}
	if params.AuthReferralExtendedPerms != nil {
		out["auth_referral_extended_perms"] = *params.AuthReferralExtendedPerms
	}
	if params.AuthReferralFriendPerms != nil {
		out["auth_referral_friend_perms"] = *params.AuthReferralFriendPerms
	}
	if params.AuthReferralResponseType != nil {
		out["auth_referral_response_type"] = *params.AuthReferralResponseType
	}
	if params.AuthReferralUserPerms != nil {
		out["auth_referral_user_perms"] = *params.AuthReferralUserPerms
	}
	if params.CanvasFluidHeight != nil {
		out["canvas_fluid_height"] = *params.CanvasFluidHeight
	}
	if params.CanvasFluidWidth != nil {
		out["canvas_fluid_width"] = *params.CanvasFluidWidth
	}
	if params.CanvasURL != nil {
		out["canvas_url"] = *params.CanvasURL
	}
	if params.ContactEmail != nil {
		out["contact_email"] = *params.ContactEmail
	}
	if params.DeauthCallbackURL != nil {
		out["deauth_callback_url"] = *params.DeauthCallbackURL
	}
	if params.MobileWebURL != nil {
		out["mobile_web_url"] = *params.MobileWebURL
	}
	if params.Namespace != nil {
		out["namespace"] = *params.Namespace
	}
	if params.PageTabDefaultName != nil {
		out["page_tab_default_name"] = *params.PageTabDefaultName
	}
	if params.PrivacyPolicyURL != nil {
		out["privacy_policy_url"] = *params.PrivacyPolicyURL
	}
	if params.Restrictions != nil {
		out["restrictions"] = *params.Restrictions
	}
	if params.SecureCanvasURL != nil {
		out["secure_canvas_url"] = *params.SecureCanvasURL
	}
	if params.SecurePageTabURL != nil {
		out["secure_page_tab_url"] = *params.SecurePageTabURL
	}
	if params.ServerIPWhitelist != nil {
		out["server_ip_whitelist"] = *params.ServerIPWhitelist
	}
	if params.TermsOfServiceURL != nil {
		out["terms_of_service_url"] = *params.TermsOfServiceURL
	}
	if params.URLSchemeSuffix != nil {
		out["url_scheme_suffix"] = *params.URLSchemeSuffix
	}
	if params.UserSupportEmail != nil {
		out["user_support_email"] = *params.UserSupportEmail
	}
	if params.UserSupportURL != nil {
		out["user_support_url"] = *params.UserSupportURL
	}
	if params.WebsiteURL != nil {
		out["website_url"] = *params.WebsiteURL
	}
	return out
}

func UpdateApplication(ctx context.Context, client *core.Client, id string, params UpdateApplicationParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
