package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
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

func DeleteApplicationAccountsBatchCall(id string, params DeleteApplicationAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "accounts"), params.ToParams(), options...)
}

func NewDeleteApplicationAccountsBatchRequest(id string, params DeleteApplicationAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteApplicationAccountsBatchCall(id, params, options...))
}

func DecodeDeleteApplicationAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteApplicationAccountsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteApplicationAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteApplicationAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteApplicationAccounts(ctx context.Context, client *core.Client, id string, params DeleteApplicationAccountsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteApplicationAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAccountsBatchCall(id string, params GetApplicationAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "accounts"), params.ToParams(), options...)
}

func NewGetApplicationAccountsBatchRequest(id string, params GetApplicationAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.TestAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.TestAccount]](GetApplicationAccountsBatchCall(id, params, options...))
}

func DecodeGetApplicationAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.TestAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.TestAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAccountsParams) (*core.Cursor[objects.TestAccount], *core.Response, error) {
	var out core.Cursor[objects.TestAccount]
	call := GetApplicationAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAccounts(ctx context.Context, client *core.Client, id string, params GetApplicationAccountsParams) (*core.Cursor[objects.TestAccount], error) {
	out, _, err := GetApplicationAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAccountsBatchCall(id string, params CreateApplicationAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "accounts"), params.ToParams(), options...)
}

func NewCreateApplicationAccountsBatchRequest(id string, params CreateApplicationAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.TestAccount] {
	return core.NewBatchRequest[objects.TestAccount](CreateApplicationAccountsBatchCall(id, params, options...))
}

func DecodeCreateApplicationAccountsBatchResponse(response *core.BatchResponse) (*objects.TestAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.TestAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAccountsParams) (*objects.TestAccount, *core.Response, error) {
	var out objects.TestAccount
	call := CreateApplicationAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAccounts(ctx context.Context, client *core.Client, id string, params CreateApplicationAccountsParams) (*objects.TestAccount, error) {
	out, _, err := CreateApplicationAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationActivitiesBatchCall(id string, params CreateApplicationActivitiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "activities"), params.ToParams(), options...)
}

func NewCreateApplicationActivitiesBatchRequest(id string, params CreateApplicationActivitiesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationActivitiesBatchCall(id, params, options...))
}

func DecodeCreateApplicationActivitiesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationActivitiesWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationActivitiesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationActivitiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationActivities(ctx context.Context, client *core.Client, id string, params CreateApplicationActivitiesParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationActivitiesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAdPlacementGroupsBatchCall(id string, params GetApplicationAdPlacementGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_placement_groups"), params.ToParams(), options...)
}

func NewGetApplicationAdPlacementGroupsBatchRequest(id string, params GetApplicationAdPlacementGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPlacementGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPlacementGroup]](GetApplicationAdPlacementGroupsBatchCall(id, params, options...))
}

func DecodeGetApplicationAdPlacementGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPlacementGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPlacementGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAdPlacementGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAdPlacementGroupsParams) (*core.Cursor[objects.AdPlacementGroup], *core.Response, error) {
	var out core.Cursor[objects.AdPlacementGroup]
	call := GetApplicationAdPlacementGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAdPlacementGroups(ctx context.Context, client *core.Client, id string, params GetApplicationAdPlacementGroupsParams) (*core.Cursor[objects.AdPlacementGroup], error) {
	out, _, err := GetApplicationAdPlacementGroupsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAdnetworkPlacementsBatchCall(id string, params GetApplicationAdnetworkPlacementsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetwork_placements"), params.ToParams(), options...)
}

func NewGetApplicationAdnetworkPlacementsBatchRequest(id string, params GetApplicationAdnetworkPlacementsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPlacement]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPlacement]](GetApplicationAdnetworkPlacementsBatchCall(id, params, options...))
}

func DecodeGetApplicationAdnetworkPlacementsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPlacement], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPlacement]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAdnetworkPlacementsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkPlacementsParams) (*core.Cursor[objects.AdPlacement], *core.Response, error) {
	var out core.Cursor[objects.AdPlacement]
	call := GetApplicationAdnetworkPlacementsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAdnetworkPlacements(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkPlacementsParams) (*core.Cursor[objects.AdPlacement], error) {
	out, _, err := GetApplicationAdnetworkPlacementsWithResponse(ctx, client, id, params)
	return out, err
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
	Since              *core.Time                                                     `facebook:"since"`
	Until              *core.Time                                                     `facebook:"until"`
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

func GetApplicationAdnetworkanalyticsBatchCall(id string, params GetApplicationAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewGetApplicationAdnetworkanalyticsBatchRequest(id string, params GetApplicationAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]](GetApplicationAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeGetApplicationAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAdnetworkanalyticsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], *core.Response, error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	call := GetApplicationAdnetworkanalyticsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	out, _, err := GetApplicationAdnetworkanalyticsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateApplicationAdnetworkanalyticsParams struct {
	AggregationPeriod *enums.ApplicationadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns        *[]enums.ApplicationadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters           *[]map[string]interface{}                                      `facebook:"filters"`
	Limit             *int                                                           `facebook:"limit"`
	Metrics           []enums.ApplicationadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn    *enums.ApplicationadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType      *enums.ApplicationadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	Since             *core.Time                                                     `facebook:"since"`
	Until             *core.Time                                                     `facebook:"until"`
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

func CreateApplicationAdnetworkanalyticsBatchCall(id string, params CreateApplicationAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewCreateApplicationAdnetworkanalyticsBatchRequest(id string, params CreateApplicationAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeCreateApplicationAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAdnetworkanalyticsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAdnetworkanalyticsParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationAdnetworkanalyticsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateApplicationAdnetworkanalyticsParams) (*objects.Application, error) {
	out, _, err := CreateApplicationAdnetworkanalyticsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAdnetworkanalyticsResultsBatchCall(id string, params GetApplicationAdnetworkanalyticsResultsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), options...)
}

func NewGetApplicationAdnetworkanalyticsResultsBatchRequest(id string, params GetApplicationAdnetworkanalyticsResultsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]](GetApplicationAdnetworkanalyticsResultsBatchCall(id, params, options...))
}

func DecodeGetApplicationAdnetworkanalyticsResultsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAdnetworkanalyticsResultsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], *core.Response, error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	call := GetApplicationAdnetworkanalyticsResultsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetApplicationAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	out, _, err := GetApplicationAdnetworkanalyticsResultsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAemAttributionBatchCall(id string, params GetApplicationAemAttributionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "aem_attribution"), params.ToParams(), options...)
}

func NewGetApplicationAemAttributionBatchRequest(id string, params GetApplicationAemAttributionParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AEMAttribution]] {
	return core.NewBatchRequest[core.Cursor[objects.AEMAttribution]](GetApplicationAemAttributionBatchCall(id, params, options...))
}

func DecodeGetApplicationAemAttributionBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AEMAttribution], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AEMAttribution]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAemAttributionWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAemAttributionParams) (*core.Cursor[objects.AEMAttribution], *core.Response, error) {
	var out core.Cursor[objects.AEMAttribution]
	call := GetApplicationAemAttributionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAemAttribution(ctx context.Context, client *core.Client, id string, params GetApplicationAemAttributionParams) (*core.Cursor[objects.AEMAttribution], error) {
	out, _, err := GetApplicationAemAttributionWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAemConversionConfigsBatchCall(id string, params GetApplicationAemConversionConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "aem_conversion_configs"), params.ToParams(), options...)
}

func NewGetApplicationAemConversionConfigsBatchRequest(id string, params GetApplicationAemConversionConfigsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AEMConversionConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.AEMConversionConfig]](GetApplicationAemConversionConfigsBatchCall(id, params, options...))
}

func DecodeGetApplicationAemConversionConfigsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AEMConversionConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AEMConversionConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAemConversionConfigsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionConfigsParams) (*core.Cursor[objects.AEMConversionConfig], *core.Response, error) {
	var out core.Cursor[objects.AEMConversionConfig]
	call := GetApplicationAemConversionConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAemConversionConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionConfigsParams) (*core.Cursor[objects.AEMConversionConfig], error) {
	out, _, err := GetApplicationAemConversionConfigsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAemConversionFilterBatchCall(id string, params GetApplicationAemConversionFilterParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "aem_conversion_filter"), params.ToParams(), options...)
}

func NewGetApplicationAemConversionFilterBatchRequest(id string, params GetApplicationAemConversionFilterParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AEMConversionFilter]] {
	return core.NewBatchRequest[core.Cursor[objects.AEMConversionFilter]](GetApplicationAemConversionFilterBatchCall(id, params, options...))
}

func DecodeGetApplicationAemConversionFilterBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AEMConversionFilter], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AEMConversionFilter]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAemConversionFilterWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionFilterParams) (*core.Cursor[objects.AEMConversionFilter], *core.Response, error) {
	var out core.Cursor[objects.AEMConversionFilter]
	call := GetApplicationAemConversionFilterBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAemConversionFilter(ctx context.Context, client *core.Client, id string, params GetApplicationAemConversionFilterParams) (*core.Cursor[objects.AEMConversionFilter], error) {
	out, _, err := GetApplicationAemConversionFilterWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAemConversionsBatchCall(id string, params CreateApplicationAemConversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "aem_conversions"), params.ToParams(), options...)
}

func NewCreateApplicationAemConversionsBatchRequest(id string, params CreateApplicationAemConversionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationAemConversionsBatchCall(id, params, options...))
}

func DecodeCreateApplicationAemConversionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAemConversionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAemConversionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationAemConversionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAemConversions(ctx context.Context, client *core.Client, id string, params CreateApplicationAemConversionsParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationAemConversionsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAemSkanReadinessBatchCall(id string, params CreateApplicationAemSkanReadinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "aem_skan_readiness"), params.ToParams(), options...)
}

func NewCreateApplicationAemSkanReadinessBatchRequest(id string, params CreateApplicationAemSkanReadinessParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationAemSkanReadinessBatchCall(id, params, options...))
}

func DecodeCreateApplicationAemSkanReadinessBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAemSkanReadinessWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAemSkanReadinessParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationAemSkanReadinessBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAemSkanReadiness(ctx context.Context, client *core.Client, id string, params CreateApplicationAemSkanReadinessParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationAemSkanReadinessWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAgenciesBatchCall(id string, params GetApplicationAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetApplicationAgenciesBatchRequest(id string, params GetApplicationAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetApplicationAgenciesBatchCall(id, params, options...))
}

func DecodeGetApplicationAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetApplicationAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAgencies(ctx context.Context, client *core.Client, id string, params GetApplicationAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetApplicationAgenciesWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAggregateRevenueBatchCall(id string, params CreateApplicationAggregateRevenueParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "aggregate_revenue"), params.ToParams(), options...)
}

func NewCreateApplicationAggregateRevenueBatchRequest(id string, params CreateApplicationAggregateRevenueParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationAggregateRevenueBatchCall(id, params, options...))
}

func DecodeCreateApplicationAggregateRevenueBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAggregateRevenueWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAggregateRevenueParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationAggregateRevenueBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAggregateRevenue(ctx context.Context, client *core.Client, id string, params CreateApplicationAggregateRevenueParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationAggregateRevenueWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAndroidDialogConfigsBatchCall(id string, params GetApplicationAndroidDialogConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "android_dialog_configs"), params.ToParams(), options...)
}

func NewGetApplicationAndroidDialogConfigsBatchRequest(id string, params GetApplicationAndroidDialogConfigsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationDialogConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationDialogConfig]](GetApplicationAndroidDialogConfigsBatchCall(id, params, options...))
}

func DecodeGetApplicationAndroidDialogConfigsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationDialogConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAndroidDialogConfigsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAndroidDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], *core.Response, error) {
	var out core.Cursor[objects.ApplicationDialogConfig]
	call := GetApplicationAndroidDialogConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAndroidDialogConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationAndroidDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	out, _, err := GetApplicationAndroidDialogConfigsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAppCapiSettingsBatchCall(id string, params GetApplicationAppCapiSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "app_capi_settings"), params.ToParams(), options...)
}

func NewGetApplicationAppCapiSettingsBatchRequest(id string, params GetApplicationAppCapiSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationAppCapiSetting]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationAppCapiSetting]](GetApplicationAppCapiSettingsBatchCall(id, params, options...))
}

func DecodeGetApplicationAppCapiSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationAppCapiSetting], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationAppCapiSetting]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAppCapiSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAppCapiSettingsParams) (*core.Cursor[objects.ApplicationAppCapiSetting], *core.Response, error) {
	var out core.Cursor[objects.ApplicationAppCapiSetting]
	call := GetApplicationAppCapiSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAppCapiSettings(ctx context.Context, client *core.Client, id string, params GetApplicationAppCapiSettingsParams) (*core.Cursor[objects.ApplicationAppCapiSetting], error) {
	out, _, err := GetApplicationAppCapiSettingsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAppEventTypesBatchCall(id string, params GetApplicationAppEventTypesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "app_event_types"), params.ToParams(), options...)
}

func NewGetApplicationAppEventTypesBatchRequest(id string, params GetApplicationAppEventTypesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationAppEventTypes]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationAppEventTypes]](GetApplicationAppEventTypesBatchCall(id, params, options...))
}

func DecodeGetApplicationAppEventTypesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationAppEventTypes], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationAppEventTypes]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAppEventTypesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAppEventTypesParams) (*core.Cursor[objects.ApplicationAppEventTypes], *core.Response, error) {
	var out core.Cursor[objects.ApplicationAppEventTypes]
	call := GetApplicationAppEventTypesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAppEventTypes(ctx context.Context, client *core.Client, id string, params GetApplicationAppEventTypesParams) (*core.Cursor[objects.ApplicationAppEventTypes], error) {
	out, _, err := GetApplicationAppEventTypesWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAppIndexingBatchCall(id string, params CreateApplicationAppIndexingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "app_indexing"), params.ToParams(), options...)
}

func NewCreateApplicationAppIndexingBatchRequest(id string, params CreateApplicationAppIndexingParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationAppIndexingBatchCall(id, params, options...))
}

func DecodeCreateApplicationAppIndexingBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAppIndexingWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationAppIndexingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAppIndexing(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingParams) (*objects.Application, error) {
	out, _, err := CreateApplicationAppIndexingWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAppIndexingSessionBatchCall(id string, params CreateApplicationAppIndexingSessionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "app_indexing_session"), params.ToParams(), options...)
}

func NewCreateApplicationAppIndexingSessionBatchRequest(id string, params CreateApplicationAppIndexingSessionParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationAppIndexingSessionBatchCall(id, params, options...))
}

func DecodeCreateApplicationAppIndexingSessionBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAppIndexingSessionWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingSessionParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationAppIndexingSessionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAppIndexingSession(ctx context.Context, client *core.Client, id string, params CreateApplicationAppIndexingSessionParams) (*objects.Application, error) {
	out, _, err := CreateApplicationAppIndexingSessionWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAppInstalledGroupsBatchCall(id string, params GetApplicationAppInstalledGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "app_installed_groups"), params.ToParams(), options...)
}

func NewGetApplicationAppInstalledGroupsBatchRequest(id string, params GetApplicationAppInstalledGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Group]] {
	return core.NewBatchRequest[core.Cursor[objects.Group]](GetApplicationAppInstalledGroupsBatchCall(id, params, options...))
}

func DecodeGetApplicationAppInstalledGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Group], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Group]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAppInstalledGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAppInstalledGroupsParams) (*core.Cursor[objects.Group], *core.Response, error) {
	var out core.Cursor[objects.Group]
	call := GetApplicationAppInstalledGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAppInstalledGroups(ctx context.Context, client *core.Client, id string, params GetApplicationAppInstalledGroupsParams) (*core.Cursor[objects.Group], error) {
	out, _, err := GetApplicationAppInstalledGroupsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAppPushDeviceTokenBatchCall(id string, params CreateApplicationAppPushDeviceTokenParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "app_push_device_token"), params.ToParams(), options...)
}

func NewCreateApplicationAppPushDeviceTokenBatchRequest(id string, params CreateApplicationAppPushDeviceTokenParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationAppPushDeviceTokenBatchCall(id, params, options...))
}

func DecodeCreateApplicationAppPushDeviceTokenBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAppPushDeviceTokenWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAppPushDeviceTokenParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationAppPushDeviceTokenBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAppPushDeviceToken(ctx context.Context, client *core.Client, id string, params CreateApplicationAppPushDeviceTokenParams) (*objects.Application, error) {
	out, _, err := CreateApplicationAppPushDeviceTokenWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAppassetsBatchCall(id string, params GetApplicationAppassetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "appassets"), params.ToParams(), options...)
}

func NewGetApplicationAppassetsBatchRequest(id string, params GetApplicationAppassetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CanvasAppAsset]] {
	return core.NewBatchRequest[core.Cursor[objects.CanvasAppAsset]](GetApplicationAppassetsBatchCall(id, params, options...))
}

func DecodeGetApplicationAppassetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CanvasAppAsset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CanvasAppAsset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAppassetsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAppassetsParams) (*core.Cursor[objects.CanvasAppAsset], *core.Response, error) {
	var out core.Cursor[objects.CanvasAppAsset]
	call := GetApplicationAppassetsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAppassets(ctx context.Context, client *core.Client, id string, params GetApplicationAppassetsParams) (*core.Cursor[objects.CanvasAppAsset], error) {
	out, _, err := GetApplicationAppassetsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationAssetsBatchCall(id string, params CreateApplicationAssetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assets"), params.ToParams(), options...)
}

func NewCreateApplicationAssetsBatchRequest(id string, params CreateApplicationAssetsParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationAssetsBatchCall(id, params, options...))
}

func DecodeCreateApplicationAssetsBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationAssetsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationAssetsParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationAssetsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationAssets(ctx context.Context, client *core.Client, id string, params CreateApplicationAssetsParams) (*objects.Application, error) {
	out, _, err := CreateApplicationAssetsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationAuthorizedAdaccountsBatchCall(id string, params GetApplicationAuthorizedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), options...)
}

func NewGetApplicationAuthorizedAdaccountsBatchRequest(id string, params GetApplicationAuthorizedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetApplicationAuthorizedAdaccountsBatchCall(id, params, options...))
}

func DecodeGetApplicationAuthorizedAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationAuthorizedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetApplicationAuthorizedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params GetApplicationAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetApplicationAuthorizedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationButtonAutoDetectionDeviceSelectionBatchCall(id string, params GetApplicationButtonAutoDetectionDeviceSelectionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "button_auto_detection_device_selection"), params.ToParams(), options...)
}

func NewGetApplicationButtonAutoDetectionDeviceSelectionBatchRequest(id string, params GetApplicationButtonAutoDetectionDeviceSelectionParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ButtonAutoDetectionDeviceSelection]] {
	return core.NewBatchRequest[core.Cursor[objects.ButtonAutoDetectionDeviceSelection]](GetApplicationButtonAutoDetectionDeviceSelectionBatchCall(id, params, options...))
}

func DecodeGetApplicationButtonAutoDetectionDeviceSelectionBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ButtonAutoDetectionDeviceSelection], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ButtonAutoDetectionDeviceSelection]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationButtonAutoDetectionDeviceSelectionWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationButtonAutoDetectionDeviceSelectionParams) (*core.Cursor[objects.ButtonAutoDetectionDeviceSelection], *core.Response, error) {
	var out core.Cursor[objects.ButtonAutoDetectionDeviceSelection]
	call := GetApplicationButtonAutoDetectionDeviceSelectionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationButtonAutoDetectionDeviceSelection(ctx context.Context, client *core.Client, id string, params GetApplicationButtonAutoDetectionDeviceSelectionParams) (*core.Cursor[objects.ButtonAutoDetectionDeviceSelection], error) {
	out, _, err := GetApplicationButtonAutoDetectionDeviceSelectionWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationCloudbridgeSettingsBatchCall(id string, params GetApplicationCloudbridgeSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "cloudbridge_settings"), params.ToParams(), options...)
}

func NewGetApplicationCloudbridgeSettingsBatchRequest(id string, params GetApplicationCloudbridgeSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CloudbridgeSetting]] {
	return core.NewBatchRequest[core.Cursor[objects.CloudbridgeSetting]](GetApplicationCloudbridgeSettingsBatchCall(id, params, options...))
}

func DecodeGetApplicationCloudbridgeSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CloudbridgeSetting], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CloudbridgeSetting]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationCloudbridgeSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationCloudbridgeSettingsParams) (*core.Cursor[objects.CloudbridgeSetting], *core.Response, error) {
	var out core.Cursor[objects.CloudbridgeSetting]
	call := GetApplicationCloudbridgeSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationCloudbridgeSettings(ctx context.Context, client *core.Client, id string, params GetApplicationCloudbridgeSettingsParams) (*core.Cursor[objects.CloudbridgeSetting], error) {
	out, _, err := GetApplicationCloudbridgeSettingsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationCodelessEventMappingsBatchCall(id string, params CreateApplicationCodelessEventMappingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "codeless_event_mappings"), params.ToParams(), options...)
}

func NewCreateApplicationCodelessEventMappingsBatchRequest(id string, params CreateApplicationCodelessEventMappingsParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationCodelessEventMappingsBatchCall(id, params, options...))
}

func DecodeCreateApplicationCodelessEventMappingsBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationCodelessEventMappingsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationCodelessEventMappingsParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationCodelessEventMappingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationCodelessEventMappings(ctx context.Context, client *core.Client, id string, params CreateApplicationCodelessEventMappingsParams) (*objects.Application, error) {
	out, _, err := CreateApplicationCodelessEventMappingsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationConnectedClientBusinessesBatchCall(id string, params GetApplicationConnectedClientBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "connected_client_businesses"), params.ToParams(), options...)
}

func NewGetApplicationConnectedClientBusinessesBatchRequest(id string, params GetApplicationConnectedClientBusinessesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetApplicationConnectedClientBusinessesBatchCall(id, params, options...))
}

func DecodeGetApplicationConnectedClientBusinessesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationConnectedClientBusinessesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationConnectedClientBusinessesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetApplicationConnectedClientBusinessesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationConnectedClientBusinesses(ctx context.Context, client *core.Client, id string, params GetApplicationConnectedClientBusinessesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetApplicationConnectedClientBusinessesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationDaChecksBatchCall(id string, params GetApplicationDaChecksParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "da_checks"), params.ToParams(), options...)
}

func NewGetApplicationDaChecksBatchRequest(id string, params GetApplicationDaChecksParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DACheck]] {
	return core.NewBatchRequest[core.Cursor[objects.DACheck]](GetApplicationDaChecksBatchCall(id, params, options...))
}

func DecodeGetApplicationDaChecksBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DACheck], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.DACheck]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationDaChecksWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationDaChecksParams) (*core.Cursor[objects.DACheck], *core.Response, error) {
	var out core.Cursor[objects.DACheck]
	call := GetApplicationDaChecksBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationDaChecks(ctx context.Context, client *core.Client, id string, params GetApplicationDaChecksParams) (*core.Cursor[objects.DACheck], error) {
	out, _, err := GetApplicationDaChecksWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationDomainReportsBatchCall(id string, params CreateApplicationDomainReportsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "domain_reports"), params.ToParams(), options...)
}

func NewCreateApplicationDomainReportsBatchRequest(id string, params CreateApplicationDomainReportsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationDomainReportsBatchCall(id, params, options...))
}

func DecodeCreateApplicationDomainReportsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationDomainReportsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationDomainReportsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationDomainReportsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationDomainReports(ctx context.Context, client *core.Client, id string, params CreateApplicationDomainReportsParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationDomainReportsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationIapPurchasesBatchCall(id string, params GetApplicationIapPurchasesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "iap_purchases"), params.ToParams(), options...)
}

func NewGetApplicationIapPurchasesBatchRequest(id string, params GetApplicationIapPurchasesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DCPAppStoreOrder]] {
	return core.NewBatchRequest[core.Cursor[objects.DCPAppStoreOrder]](GetApplicationIapPurchasesBatchCall(id, params, options...))
}

func DecodeGetApplicationIapPurchasesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DCPAppStoreOrder], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.DCPAppStoreOrder]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationIapPurchasesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationIapPurchasesParams) (*core.Cursor[objects.DCPAppStoreOrder], *core.Response, error) {
	var out core.Cursor[objects.DCPAppStoreOrder]
	call := GetApplicationIapPurchasesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationIapPurchases(ctx context.Context, client *core.Client, id string, params GetApplicationIapPurchasesParams) (*core.Cursor[objects.DCPAppStoreOrder], error) {
	out, _, err := GetApplicationIapPurchasesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationIosDialogConfigsBatchCall(id string, params GetApplicationIosDialogConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ios_dialog_configs"), params.ToParams(), options...)
}

func NewGetApplicationIosDialogConfigsBatchRequest(id string, params GetApplicationIosDialogConfigsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationDialogConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationDialogConfig]](GetApplicationIosDialogConfigsBatchCall(id, params, options...))
}

func DecodeGetApplicationIosDialogConfigsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationDialogConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationIosDialogConfigsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationIosDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], *core.Response, error) {
	var out core.Cursor[objects.ApplicationDialogConfig]
	call := GetApplicationIosDialogConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationIosDialogConfigs(ctx context.Context, client *core.Client, id string, params GetApplicationIosDialogConfigsParams) (*core.Cursor[objects.ApplicationDialogConfig], error) {
	out, _, err := GetApplicationIosDialogConfigsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationLinkedDatasetBatchCall(id string, params GetApplicationLinkedDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "linked_dataset"), params.ToParams(), options...)
}

func NewGetApplicationLinkedDatasetBatchRequest(id string, params GetApplicationLinkedDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsDataset]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsDataset]](GetApplicationLinkedDatasetBatchCall(id, params, options...))
}

func DecodeGetApplicationLinkedDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsDataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsDataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationLinkedDatasetWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationLinkedDatasetParams) (*core.Cursor[objects.AdsDataset], *core.Response, error) {
	var out core.Cursor[objects.AdsDataset]
	call := GetApplicationLinkedDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationLinkedDataset(ctx context.Context, client *core.Client, id string, params GetApplicationLinkedDatasetParams) (*core.Cursor[objects.AdsDataset], error) {
	out, _, err := GetApplicationLinkedDatasetWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationMmpAuditingBatchCall(id string, params CreateApplicationMmpAuditingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "mmp_auditing"), params.ToParams(), options...)
}

func NewCreateApplicationMmpAuditingBatchRequest(id string, params CreateApplicationMmpAuditingParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationMmpAuditingBatchCall(id, params, options...))
}

func DecodeCreateApplicationMmpAuditingBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationMmpAuditingWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationMmpAuditingParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationMmpAuditingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationMmpAuditing(ctx context.Context, client *core.Client, id string, params CreateApplicationMmpAuditingParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationMmpAuditingWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationMobileSdkGkBatchCall(id string, params GetApplicationMobileSdkGkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "mobile_sdk_gk"), params.ToParams(), options...)
}

func NewGetApplicationMobileSdkGkBatchRequest(id string, params GetApplicationMobileSdkGkParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MobileSdkGk]] {
	return core.NewBatchRequest[core.Cursor[objects.MobileSdkGk]](GetApplicationMobileSdkGkBatchCall(id, params, options...))
}

func DecodeGetApplicationMobileSdkGkBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MobileSdkGk], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MobileSdkGk]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationMobileSdkGkWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationMobileSdkGkParams) (*core.Cursor[objects.MobileSdkGk], *core.Response, error) {
	var out core.Cursor[objects.MobileSdkGk]
	call := GetApplicationMobileSdkGkBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationMobileSdkGk(ctx context.Context, client *core.Client, id string, params GetApplicationMobileSdkGkParams) (*core.Cursor[objects.MobileSdkGk], error) {
	out, _, err := GetApplicationMobileSdkGkWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationMonetizedDigitalStoreObjectsBatchCall(id string, params GetApplicationMonetizedDigitalStoreObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "monetized_digital_store_objects"), params.ToParams(), options...)
}

func NewGetApplicationMonetizedDigitalStoreObjectsBatchRequest(id string, params GetApplicationMonetizedDigitalStoreObjectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MonetizedDigitalStoreObject]] {
	return core.NewBatchRequest[core.Cursor[objects.MonetizedDigitalStoreObject]](GetApplicationMonetizedDigitalStoreObjectsBatchCall(id, params, options...))
}

func DecodeGetApplicationMonetizedDigitalStoreObjectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MonetizedDigitalStoreObject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MonetizedDigitalStoreObject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationMonetizedDigitalStoreObjectsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationMonetizedDigitalStoreObjectsParams) (*core.Cursor[objects.MonetizedDigitalStoreObject], *core.Response, error) {
	var out core.Cursor[objects.MonetizedDigitalStoreObject]
	call := GetApplicationMonetizedDigitalStoreObjectsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationMonetizedDigitalStoreObjects(ctx context.Context, client *core.Client, id string, params GetApplicationMonetizedDigitalStoreObjectsParams) (*core.Cursor[objects.MonetizedDigitalStoreObject], error) {
	out, _, err := GetApplicationMonetizedDigitalStoreObjectsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationMonetizedDigitalStoreObjectsBatchCall(id string, params CreateApplicationMonetizedDigitalStoreObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "monetized_digital_store_objects"), params.ToParams(), options...)
}

func NewCreateApplicationMonetizedDigitalStoreObjectsBatchRequest(id string, params CreateApplicationMonetizedDigitalStoreObjectsParams, options ...core.BatchOption) *core.BatchRequest[objects.MonetizedDigitalStoreObject] {
	return core.NewBatchRequest[objects.MonetizedDigitalStoreObject](CreateApplicationMonetizedDigitalStoreObjectsBatchCall(id, params, options...))
}

func DecodeCreateApplicationMonetizedDigitalStoreObjectsBatchResponse(response *core.BatchResponse) (*objects.MonetizedDigitalStoreObject, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MonetizedDigitalStoreObject
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationMonetizedDigitalStoreObjectsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationMonetizedDigitalStoreObjectsParams) (*objects.MonetizedDigitalStoreObject, *core.Response, error) {
	var out objects.MonetizedDigitalStoreObject
	call := CreateApplicationMonetizedDigitalStoreObjectsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationMonetizedDigitalStoreObjects(ctx context.Context, client *core.Client, id string, params CreateApplicationMonetizedDigitalStoreObjectsParams) (*objects.MonetizedDigitalStoreObject, error) {
	out, _, err := CreateApplicationMonetizedDigitalStoreObjectsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationObjectTypesBatchCall(id string, params GetApplicationObjectTypesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "object_types"), params.ToParams(), options...)
}

func NewGetApplicationObjectTypesBatchRequest(id string, params GetApplicationObjectTypesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetApplicationObjectTypesBatchCall(id, params, options...))
}

func DecodeGetApplicationObjectTypesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationObjectTypesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationObjectTypesParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetApplicationObjectTypesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationObjectTypes(ctx context.Context, client *core.Client, id string, params GetApplicationObjectTypesParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetApplicationObjectTypesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationObjectsBatchCall(id string, params GetApplicationObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "objects"), params.ToParams(), options...)
}

func NewGetApplicationObjectsBatchRequest(id string, params GetApplicationObjectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetApplicationObjectsBatchCall(id, params, options...))
}

func DecodeGetApplicationObjectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationObjectsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationObjectsParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetApplicationObjectsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationObjects(ctx context.Context, client *core.Client, id string, params GetApplicationObjectsParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetApplicationObjectsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationOccludespopupsBatchCall(id string, params CreateApplicationOccludespopupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "occludespopups"), params.ToParams(), options...)
}

func NewCreateApplicationOccludespopupsBatchRequest(id string, params CreateApplicationOccludespopupsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationOccludespopupsBatchCall(id, params, options...))
}

func DecodeCreateApplicationOccludespopupsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationOccludespopupsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationOccludespopupsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationOccludespopupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationOccludespopups(ctx context.Context, client *core.Client, id string, params CreateApplicationOccludespopupsParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationOccludespopupsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationPermissionsBatchCall(id string, params GetApplicationPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "permissions"), params.ToParams(), options...)
}

func NewGetApplicationPermissionsBatchRequest(id string, params GetApplicationPermissionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationPermission]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationPermission]](GetApplicationPermissionsBatchCall(id, params, options...))
}

func DecodeGetApplicationPermissionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationPermission], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationPermission]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationPermissionsParams) (*core.Cursor[objects.ApplicationPermission], *core.Response, error) {
	var out core.Cursor[objects.ApplicationPermission]
	call := GetApplicationPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationPermissions(ctx context.Context, client *core.Client, id string, params GetApplicationPermissionsParams) (*core.Cursor[objects.ApplicationPermission], error) {
	out, _, err := GetApplicationPermissionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationProductsBatchCall(id string, params GetApplicationProductsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), options...)
}

func NewGetApplicationProductsBatchRequest(id string, params GetApplicationProductsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.GamesIAPProduct]] {
	return core.NewBatchRequest[core.Cursor[objects.GamesIAPProduct]](GetApplicationProductsBatchCall(id, params, options...))
}

func DecodeGetApplicationProductsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.GamesIAPProduct], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.GamesIAPProduct]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationProductsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationProductsParams) (*core.Cursor[objects.GamesIAPProduct], *core.Response, error) {
	var out core.Cursor[objects.GamesIAPProduct]
	call := GetApplicationProductsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationProducts(ctx context.Context, client *core.Client, id string, params GetApplicationProductsParams) (*core.Cursor[objects.GamesIAPProduct], error) {
	out, _, err := GetApplicationProductsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationPurchasesBatchCall(id string, params GetApplicationPurchasesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "purchases"), params.ToParams(), options...)
}

func NewGetApplicationPurchasesBatchRequest(id string, params GetApplicationPurchasesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.GamesIAPOrder]] {
	return core.NewBatchRequest[core.Cursor[objects.GamesIAPOrder]](GetApplicationPurchasesBatchCall(id, params, options...))
}

func DecodeGetApplicationPurchasesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.GamesIAPOrder], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.GamesIAPOrder]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationPurchasesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationPurchasesParams) (*core.Cursor[objects.GamesIAPOrder], *core.Response, error) {
	var out core.Cursor[objects.GamesIAPOrder]
	call := GetApplicationPurchasesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationPurchases(ctx context.Context, client *core.Client, id string, params GetApplicationPurchasesParams) (*core.Cursor[objects.GamesIAPOrder], error) {
	out, _, err := GetApplicationPurchasesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationRolesBatchCall(id string, params GetApplicationRolesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "roles"), params.ToParams(), options...)
}

func NewGetApplicationRolesBatchRequest(id string, params GetApplicationRolesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationRoles]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationRoles]](GetApplicationRolesBatchCall(id, params, options...))
}

func DecodeGetApplicationRolesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationRoles], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationRoles]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationRolesWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationRolesParams) (*core.Cursor[objects.ApplicationRoles], *core.Response, error) {
	var out core.Cursor[objects.ApplicationRoles]
	call := GetApplicationRolesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationRoles(ctx context.Context, client *core.Client, id string, params GetApplicationRolesParams) (*core.Cursor[objects.ApplicationRoles], error) {
	out, _, err := GetApplicationRolesWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationServerDomainInfosBatchCall(id string, params GetApplicationServerDomainInfosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "server_domain_infos"), params.ToParams(), options...)
}

func NewGetApplicationServerDomainInfosBatchRequest(id string, params GetApplicationServerDomainInfosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationServerDomainInfos]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationServerDomainInfos]](GetApplicationServerDomainInfosBatchCall(id, params, options...))
}

func DecodeGetApplicationServerDomainInfosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationServerDomainInfos], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationServerDomainInfos]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationServerDomainInfosWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationServerDomainInfosParams) (*core.Cursor[objects.ApplicationServerDomainInfos], *core.Response, error) {
	var out core.Cursor[objects.ApplicationServerDomainInfos]
	call := GetApplicationServerDomainInfosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationServerDomainInfos(ctx context.Context, client *core.Client, id string, params GetApplicationServerDomainInfosParams) (*core.Cursor[objects.ApplicationServerDomainInfos], error) {
	out, _, err := GetApplicationServerDomainInfosWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationSgwDatasetStatusBatchCall(id string, params GetApplicationSgwDatasetStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sgw_dataset_status"), params.ToParams(), options...)
}

func NewGetApplicationSgwDatasetStatusBatchRequest(id string, params GetApplicationSgwDatasetStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationSgwDatasetStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationSgwDatasetStatus]](GetApplicationSgwDatasetStatusBatchCall(id, params, options...))
}

func DecodeGetApplicationSgwDatasetStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationSgwDatasetStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationSgwDatasetStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationSgwDatasetStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationSgwDatasetStatusParams) (*core.Cursor[objects.ApplicationSgwDatasetStatus], *core.Response, error) {
	var out core.Cursor[objects.ApplicationSgwDatasetStatus]
	call := GetApplicationSgwDatasetStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationSgwDatasetStatus(ctx context.Context, client *core.Client, id string, params GetApplicationSgwDatasetStatusParams) (*core.Cursor[objects.ApplicationSgwDatasetStatus], error) {
	out, _, err := GetApplicationSgwDatasetStatusWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationSgwInstallDeferralLinkBatchCall(id string, params GetApplicationSgwInstallDeferralLinkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sgw_install_deferral_link"), params.ToParams(), options...)
}

func NewGetApplicationSgwInstallDeferralLinkBatchRequest(id string, params GetApplicationSgwInstallDeferralLinkParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationSgwInstallDeferralLink]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationSgwInstallDeferralLink]](GetApplicationSgwInstallDeferralLinkBatchCall(id, params, options...))
}

func DecodeGetApplicationSgwInstallDeferralLinkBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationSgwInstallDeferralLink], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationSgwInstallDeferralLink]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationSgwInstallDeferralLinkWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationSgwInstallDeferralLinkParams) (*core.Cursor[objects.ApplicationSgwInstallDeferralLink], *core.Response, error) {
	var out core.Cursor[objects.ApplicationSgwInstallDeferralLink]
	call := GetApplicationSgwInstallDeferralLinkBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationSgwInstallDeferralLink(ctx context.Context, client *core.Client, id string, params GetApplicationSgwInstallDeferralLinkParams) (*core.Cursor[objects.ApplicationSgwInstallDeferralLink], error) {
	out, _, err := GetApplicationSgwInstallDeferralLinkWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteApplicationSubscriptionsBatchCall(id string, params DeleteApplicationSubscriptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "subscriptions"), params.ToParams(), options...)
}

func NewDeleteApplicationSubscriptionsBatchRequest(id string, params DeleteApplicationSubscriptionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteApplicationSubscriptionsBatchCall(id, params, options...))
}

func DecodeDeleteApplicationSubscriptionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteApplicationSubscriptionsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteApplicationSubscriptionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteApplicationSubscriptionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params DeleteApplicationSubscriptionsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteApplicationSubscriptionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationSubscriptionsBatchCall(id string, params GetApplicationSubscriptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "subscriptions"), params.ToParams(), options...)
}

func NewGetApplicationSubscriptionsBatchRequest(id string, params GetApplicationSubscriptionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ApplicationSubscriptions]] {
	return core.NewBatchRequest[core.Cursor[objects.ApplicationSubscriptions]](GetApplicationSubscriptionsBatchCall(id, params, options...))
}

func DecodeGetApplicationSubscriptionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ApplicationSubscriptions], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ApplicationSubscriptions]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationSubscriptionsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationSubscriptionsParams) (*core.Cursor[objects.ApplicationSubscriptions], *core.Response, error) {
	var out core.Cursor[objects.ApplicationSubscriptions]
	call := GetApplicationSubscriptionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params GetApplicationSubscriptionsParams) (*core.Cursor[objects.ApplicationSubscriptions], error) {
	out, _, err := GetApplicationSubscriptionsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationSubscriptionsBatchCall(id string, params CreateApplicationSubscriptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "subscriptions"), params.ToParams(), options...)
}

func NewCreateApplicationSubscriptionsBatchRequest(id string, params CreateApplicationSubscriptionsParams, options ...core.BatchOption) *core.BatchRequest[objects.ApplicationSubscriptions] {
	return core.NewBatchRequest[objects.ApplicationSubscriptions](CreateApplicationSubscriptionsBatchCall(id, params, options...))
}

func DecodeCreateApplicationSubscriptionsBatchResponse(response *core.BatchResponse) (*objects.ApplicationSubscriptions, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ApplicationSubscriptions
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationSubscriptionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationSubscriptionsParams) (*objects.ApplicationSubscriptions, *core.Response, error) {
	var out objects.ApplicationSubscriptions
	call := CreateApplicationSubscriptionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationSubscriptions(ctx context.Context, client *core.Client, id string, params CreateApplicationSubscriptionsParams) (*objects.ApplicationSubscriptions, error) {
	out, _, err := CreateApplicationSubscriptionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationThreatPrivacyGroupsMemberBatchCall(id string, params GetApplicationThreatPrivacyGroupsMemberParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "threat_privacy_groups_member"), params.ToParams(), options...)
}

func NewGetApplicationThreatPrivacyGroupsMemberBatchRequest(id string, params GetApplicationThreatPrivacyGroupsMemberParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ThreatPrivacyGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.ThreatPrivacyGroup]](GetApplicationThreatPrivacyGroupsMemberBatchCall(id, params, options...))
}

func DecodeGetApplicationThreatPrivacyGroupsMemberBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ThreatPrivacyGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationThreatPrivacyGroupsMemberWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsMemberParams) (*core.Cursor[objects.ThreatPrivacyGroup], *core.Response, error) {
	var out core.Cursor[objects.ThreatPrivacyGroup]
	call := GetApplicationThreatPrivacyGroupsMemberBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationThreatPrivacyGroupsMember(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsMemberParams) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	out, _, err := GetApplicationThreatPrivacyGroupsMemberWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationThreatPrivacyGroupsOwnerBatchCall(id string, params GetApplicationThreatPrivacyGroupsOwnerParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "threat_privacy_groups_owner"), params.ToParams(), options...)
}

func NewGetApplicationThreatPrivacyGroupsOwnerBatchRequest(id string, params GetApplicationThreatPrivacyGroupsOwnerParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ThreatPrivacyGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.ThreatPrivacyGroup]](GetApplicationThreatPrivacyGroupsOwnerBatchCall(id, params, options...))
}

func DecodeGetApplicationThreatPrivacyGroupsOwnerBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ThreatPrivacyGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationThreatPrivacyGroupsOwnerWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsOwnerParams) (*core.Cursor[objects.ThreatPrivacyGroup], *core.Response, error) {
	var out core.Cursor[objects.ThreatPrivacyGroup]
	call := GetApplicationThreatPrivacyGroupsOwnerBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationThreatPrivacyGroupsOwner(ctx context.Context, client *core.Client, id string, params GetApplicationThreatPrivacyGroupsOwnerParams) (*core.Cursor[objects.ThreatPrivacyGroup], error) {
	out, _, err := GetApplicationThreatPrivacyGroupsOwnerWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationUploadsBatchCall(id string, params CreateApplicationUploadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "uploads"), params.ToParams(), options...)
}

func NewCreateApplicationUploadsBatchRequest(id string, params CreateApplicationUploadsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateApplicationUploadsBatchCall(id, params, options...))
}

func DecodeCreateApplicationUploadsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationUploadsWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationUploadsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateApplicationUploadsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationUploads(ctx context.Context, client *core.Client, id string, params CreateApplicationUploadsParams) (*map[string]interface{}, error) {
	out, _, err := CreateApplicationUploadsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateApplicationWhatsappBusinessSolutionBatchCall(id string, params CreateApplicationWhatsappBusinessSolutionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "whatsapp_business_solution"), params.ToParams(), options...)
}

func NewCreateApplicationWhatsappBusinessSolutionBatchRequest(id string, params CreateApplicationWhatsappBusinessSolutionParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateApplicationWhatsappBusinessSolutionBatchCall(id, params, options...))
}

func DecodeCreateApplicationWhatsappBusinessSolutionBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateApplicationWhatsappBusinessSolutionWithResponse(ctx context.Context, client *core.Client, id string, params CreateApplicationWhatsappBusinessSolutionParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := CreateApplicationWhatsappBusinessSolutionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateApplicationWhatsappBusinessSolution(ctx context.Context, client *core.Client, id string, params CreateApplicationWhatsappBusinessSolutionParams) (*objects.Application, error) {
	out, _, err := CreateApplicationWhatsappBusinessSolutionWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationWhatsappBusinessSolutionsBatchCall(id string, params GetApplicationWhatsappBusinessSolutionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "whatsapp_business_solutions"), params.ToParams(), options...)
}

func NewGetApplicationWhatsappBusinessSolutionsBatchRequest(id string, params GetApplicationWhatsappBusinessSolutionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessSolution]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessSolution]](GetApplicationWhatsappBusinessSolutionsBatchCall(id, params, options...))
}

func DecodeGetApplicationWhatsappBusinessSolutionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationWhatsappBusinessSolutionsWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationWhatsappBusinessSolutionsParams) (*core.Cursor[objects.WhatsAppBusinessSolution], *core.Response, error) {
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	call := GetApplicationWhatsappBusinessSolutionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplicationWhatsappBusinessSolutions(ctx context.Context, client *core.Client, id string, params GetApplicationWhatsappBusinessSolutionsParams) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	out, _, err := GetApplicationWhatsappBusinessSolutionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetApplicationBatchCall(id string, params GetApplicationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetApplicationBatchRequest(id string, params GetApplicationParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](GetApplicationBatchCall(id, params, options...))
}

func DecodeGetApplicationBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetApplicationWithResponse(ctx context.Context, client *core.Client, id string, params GetApplicationParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := GetApplicationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetApplication(ctx context.Context, client *core.Client, id string, params GetApplicationParams) (*objects.Application, error) {
	out, _, err := GetApplicationWithResponse(ctx, client, id, params)
	return out, err
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

func UpdateApplicationBatchCall(id string, params UpdateApplicationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateApplicationBatchRequest(id string, params UpdateApplicationParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](UpdateApplicationBatchCall(id, params, options...))
}

func DecodeUpdateApplicationBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateApplicationWithResponse(ctx context.Context, client *core.Client, id string, params UpdateApplicationParams) (*objects.Application, *core.Response, error) {
	var out objects.Application
	call := UpdateApplicationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateApplication(ctx context.Context, client *core.Client, id string, params UpdateApplicationParams) (*objects.Application, error) {
	out, _, err := UpdateApplicationWithResponse(ctx, client, id, params)
	return out, err
}
