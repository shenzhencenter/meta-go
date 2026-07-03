package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsPixelAdaccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelAdaccountsBatchCall(id string, params GetAdsPixelAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetAdsPixelAdaccountsBatchRequest(id string, params GetAdsPixelAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetAdsPixelAdaccountsBatchCall(id, params, options...))
}

func DecodeGetAdsPixelAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetAdsPixelAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetAdsPixelAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdsPixelAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetAdsPixelAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdsPixelAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteAdsPixelAgenciesBatchCall(id string, params DeleteAdsPixelAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewDeleteAdsPixelAgenciesBatchRequest(id string, params DeleteAdsPixelAgenciesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdsPixelAgenciesBatchCall(id, params, options...))
}

func DecodeDeleteAdsPixelAgenciesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdsPixelAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelAgenciesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdsPixelAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelAgenciesParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdsPixelAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelAgenciesBatchCall(id string, params GetAdsPixelAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetAdsPixelAgenciesBatchRequest(id string, params GetAdsPixelAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetAdsPixelAgenciesBatchCall(id, params, options...))
}

func DecodeGetAdsPixelAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetAdsPixelAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetAdsPixelAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params GetAdsPixelAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetAdsPixelAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelAgenciesParams struct {
	Business       string                                          `facebook:"business"`
	PermittedTasks []enums.AdspixelagenciesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                     `facebook:"-"`
}

func (params CreateAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	out["permitted_tasks"] = params.PermittedTasks
	return out
}

func CreateAdsPixelAgenciesBatchCall(id string, params CreateAdsPixelAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewCreateAdsPixelAgenciesBatchRequest(id string, params CreateAdsPixelAgenciesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](CreateAdsPixelAgenciesBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelAgenciesBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdsPixelAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAgenciesParams) (*objects.AdsPixel, *core.Response, error) {
	var out objects.AdsPixel
	call := CreateAdsPixelAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAgenciesParams) (*objects.AdsPixel, error) {
	out, _, err := CreateAdsPixelAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelAhpConfigsParams struct {
	ApplinkAutosetup bool        `facebook:"applink_autosetup"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateAdsPixelAhpConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["applink_autosetup"] = params.ApplinkAutosetup
	return out
}

func CreateAdsPixelAhpConfigsBatchCall(id string, params CreateAdsPixelAhpConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ahp_configs"), params.ToParams(), options...)
}

func NewCreateAdsPixelAhpConfigsBatchRequest(id string, params CreateAdsPixelAhpConfigsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdsPixelAhpConfigsBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelAhpConfigsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdsPixelAhpConfigsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAhpConfigsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateAdsPixelAhpConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelAhpConfigs(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAhpConfigsParams) (*map[string]interface{}, error) {
	out, _, err := CreateAdsPixelAhpConfigsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelAssignedUsersBatchCall(id string, params GetAdsPixelAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetAdsPixelAssignedUsersBatchRequest(id string, params GetAdsPixelAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetAdsPixelAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetAdsPixelAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AssignedUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelAssignedUsersParams) (*core.Cursor[objects.AssignedUser], *core.Response, error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetAdsPixelAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelAssignedUsers(ctx context.Context, client *core.Client, id string, params GetAdsPixelAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	out, _, err := GetAdsPixelAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelAssignedUsersParams struct {
	Tasks []enums.AdspixelassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                         `facebook:"user"`
	Extra core.Params                                 `facebook:"-"`
}

func (params CreateAdsPixelAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tasks"] = params.Tasks
	out["user"] = params.User
	return out
}

func CreateAdsPixelAssignedUsersBatchCall(id string, params CreateAdsPixelAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreateAdsPixelAssignedUsersBatchRequest(id string, params CreateAdsPixelAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](CreateAdsPixelAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdsPixelAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAssignedUsersParams) (*objects.AdsPixel, *core.Response, error) {
	var out objects.AdsPixel
	call := CreateAdsPixelAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAssignedUsersParams) (*objects.AdsPixel, error) {
	out, _, err := CreateAdsPixelAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelDaChecksParams struct {
	Checks           *[]string                                        `facebook:"checks"`
	ConnectionMethod *enums.AdspixeldaChecksConnectionMethodEnumParam `facebook:"connection_method"`
	Extra            core.Params                                      `facebook:"-"`
}

func (params GetAdsPixelDaChecksParams) ToParams() core.Params {
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

func GetAdsPixelDaChecksBatchCall(id string, params GetAdsPixelDaChecksParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "da_checks"), params.ToParams(), options...)
}

func NewGetAdsPixelDaChecksBatchRequest(id string, params GetAdsPixelDaChecksParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DACheck]] {
	return core.NewBatchRequest[core.Cursor[objects.DACheck]](GetAdsPixelDaChecksBatchCall(id, params, options...))
}

func DecodeGetAdsPixelDaChecksBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DACheck], error) {
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

func GetAdsPixelDaChecksWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelDaChecksParams) (*core.Cursor[objects.DACheck], *core.Response, error) {
	var out core.Cursor[objects.DACheck]
	call := GetAdsPixelDaChecksBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelDaChecks(ctx context.Context, client *core.Client, id string, params GetAdsPixelDaChecksParams) (*core.Cursor[objects.DACheck], error) {
	out, _, err := GetAdsPixelDaChecksWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelEventsParams struct {
	Data          []string                  `facebook:"data"`
	NamespaceID   *core.ID                  `facebook:"namespace_id"`
	PartnerAgent  *string                   `facebook:"partner_agent"`
	Platforms     *[]map[string]interface{} `facebook:"platforms"`
	Progress      *map[string]interface{}   `facebook:"progress"`
	TestEventCode *string                   `facebook:"test_event_code"`
	Trace         *uint64                   `facebook:"trace"`
	UploadID      *core.ID                  `facebook:"upload_id"`
	UploadSource  *string                   `facebook:"upload_source"`
	UploadTag     *string                   `facebook:"upload_tag"`
	Extra         core.Params               `facebook:"-"`
}

func (params CreateAdsPixelEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["data"] = params.Data
	if params.NamespaceID != nil {
		out["namespace_id"] = *params.NamespaceID
	}
	if params.PartnerAgent != nil {
		out["partner_agent"] = *params.PartnerAgent
	}
	if params.Platforms != nil {
		out["platforms"] = *params.Platforms
	}
	if params.Progress != nil {
		out["progress"] = *params.Progress
	}
	if params.TestEventCode != nil {
		out["test_event_code"] = *params.TestEventCode
	}
	if params.Trace != nil {
		out["trace"] = *params.Trace
	}
	if params.UploadID != nil {
		out["upload_id"] = *params.UploadID
	}
	if params.UploadSource != nil {
		out["upload_source"] = *params.UploadSource
	}
	if params.UploadTag != nil {
		out["upload_tag"] = *params.UploadTag
	}
	return out
}

func CreateAdsPixelEventsBatchCall(id string, params CreateAdsPixelEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "events"), params.ToParams(), options...)
}

func NewCreateAdsPixelEventsBatchRequest(id string, params CreateAdsPixelEventsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdsPixelEventsBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelEventsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdsPixelEventsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelEventsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateAdsPixelEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelEvents(ctx context.Context, client *core.Client, id string, params CreateAdsPixelEventsParams) (*map[string]interface{}, error) {
	out, _, err := CreateAdsPixelEventsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelOfflineEventUploadsParams struct {
	EndTime   *core.Time                                        `facebook:"end_time"`
	Order     *enums.AdspixelofflineEventUploadsOrderEnumParam  `facebook:"order"`
	SortBy    *enums.AdspixelofflineEventUploadsSortByEnumParam `facebook:"sort_by"`
	StartTime *core.Time                                        `facebook:"start_time"`
	UploadTag *string                                           `facebook:"upload_tag"`
	Extra     core.Params                                       `facebook:"-"`
}

func (params GetAdsPixelOfflineEventUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.UploadTag != nil {
		out["upload_tag"] = *params.UploadTag
	}
	return out
}

func GetAdsPixelOfflineEventUploadsBatchCall(id string, params GetAdsPixelOfflineEventUploadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "offline_event_uploads"), params.ToParams(), options...)
}

func NewGetAdsPixelOfflineEventUploadsBatchRequest(id string, params GetAdsPixelOfflineEventUploadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OfflineConversionDataSetUpload]] {
	return core.NewBatchRequest[core.Cursor[objects.OfflineConversionDataSetUpload]](GetAdsPixelOfflineEventUploadsBatchCall(id, params, options...))
}

func DecodeGetAdsPixelOfflineEventUploadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelOfflineEventUploadsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelOfflineEventUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], *core.Response, error) {
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	call := GetAdsPixelOfflineEventUploadsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelOfflineEventUploads(ctx context.Context, client *core.Client, id string, params GetAdsPixelOfflineEventUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	out, _, err := GetAdsPixelOfflineEventUploadsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelOpenbridgeConfigurationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelOpenbridgeConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelOpenbridgeConfigurationsBatchCall(id string, params GetAdsPixelOpenbridgeConfigurationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), options...)
}

func NewGetAdsPixelOpenbridgeConfigurationsBatchRequest(id string, params GetAdsPixelOpenbridgeConfigurationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OpenBridgeConfiguration]] {
	return core.NewBatchRequest[core.Cursor[objects.OpenBridgeConfiguration]](GetAdsPixelOpenbridgeConfigurationsBatchCall(id, params, options...))
}

func DecodeGetAdsPixelOpenbridgeConfigurationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OpenBridgeConfiguration]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelOpenbridgeConfigurationsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelOpenbridgeConfigurationsParams) (*core.Cursor[objects.OpenBridgeConfiguration], *core.Response, error) {
	var out core.Cursor[objects.OpenBridgeConfiguration]
	call := GetAdsPixelOpenbridgeConfigurationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params GetAdsPixelOpenbridgeConfigurationsParams) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	out, _, err := GetAdsPixelOpenbridgeConfigurationsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelShadowtraffichelperParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdsPixelShadowtraffichelperParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdsPixelShadowtraffichelperBatchCall(id string, params CreateAdsPixelShadowtraffichelperParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "shadowtraffichelper"), params.ToParams(), options...)
}

func NewCreateAdsPixelShadowtraffichelperBatchRequest(id string, params CreateAdsPixelShadowtraffichelperParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdsPixelShadowtraffichelperBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelShadowtraffichelperBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdsPixelShadowtraffichelperWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelShadowtraffichelperParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateAdsPixelShadowtraffichelperBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelShadowtraffichelper(ctx context.Context, client *core.Client, id string, params CreateAdsPixelShadowtraffichelperParams) (*map[string]interface{}, error) {
	out, _, err := CreateAdsPixelShadowtraffichelperWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdsPixelSharedAccountsParams struct {
	AccountID core.ID     `facebook:"account_id"`
	Business  string      `facebook:"business"`
	Extra     core.Params `facebook:"-"`
}

func (params DeleteAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["business"] = params.Business
	return out
}

func DeleteAdsPixelSharedAccountsBatchCall(id string, params DeleteAdsPixelSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewDeleteAdsPixelSharedAccountsBatchRequest(id string, params DeleteAdsPixelSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdsPixelSharedAccountsBatchCall(id, params, options...))
}

func DecodeDeleteAdsPixelSharedAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdsPixelSharedAccountsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelSharedAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdsPixelSharedAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelSharedAccountsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdsPixelSharedAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelSharedAccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelSharedAccountsBatchCall(id string, params GetAdsPixelSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewGetAdsPixelSharedAccountsBatchRequest(id string, params GetAdsPixelSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetAdsPixelSharedAccountsBatchCall(id, params, options...))
}

func DecodeGetAdsPixelSharedAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetAdsPixelSharedAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetAdsPixelSharedAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetAdsPixelSharedAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdsPixelSharedAccountsParams struct {
	AccountID core.ID     `facebook:"account_id"`
	Business  string      `facebook:"business"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["business"] = params.Business
	return out
}

func CreateAdsPixelSharedAccountsBatchCall(id string, params CreateAdsPixelSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewCreateAdsPixelSharedAccountsBatchRequest(id string, params CreateAdsPixelSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](CreateAdsPixelSharedAccountsBatchCall(id, params, options...))
}

func DecodeCreateAdsPixelSharedAccountsBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdsPixelSharedAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsPixelSharedAccountsParams) (*objects.AdsPixel, *core.Response, error) {
	var out objects.AdsPixel
	call := CreateAdsPixelSharedAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params CreateAdsPixelSharedAccountsParams) (*objects.AdsPixel, error) {
	out, _, err := CreateAdsPixelSharedAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelSharedAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelSharedAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelSharedAgenciesBatchCall(id string, params GetAdsPixelSharedAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_agencies"), params.ToParams(), options...)
}

func NewGetAdsPixelSharedAgenciesBatchRequest(id string, params GetAdsPixelSharedAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetAdsPixelSharedAgenciesBatchCall(id, params, options...))
}

func DecodeGetAdsPixelSharedAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetAdsPixelSharedAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetAdsPixelSharedAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelSharedAgencies(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetAdsPixelSharedAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelStatsParams struct {
	Agent       *string                                  `facebook:"agent"`
	Aggregation *enums.AdspixelstatsAggregationEnumParam `facebook:"aggregation"`
	EndTime     *core.Time                               `facebook:"end_time"`
	Event       *string                                  `facebook:"event"`
	EventSource *string                                  `facebook:"event_source"`
	StartTime   *core.Time                               `facebook:"start_time"`
	Extra       core.Params                              `facebook:"-"`
}

func (params GetAdsPixelStatsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Agent != nil {
		out["agent"] = *params.Agent
	}
	if params.Aggregation != nil {
		out["aggregation"] = *params.Aggregation
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Event != nil {
		out["event"] = *params.Event
	}
	if params.EventSource != nil {
		out["event_source"] = *params.EventSource
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetAdsPixelStatsBatchCall(id string, params GetAdsPixelStatsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), options...)
}

func NewGetAdsPixelStatsBatchRequest(id string, params GetAdsPixelStatsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixelStatsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixelStatsResult]](GetAdsPixelStatsBatchCall(id, params, options...))
}

func DecodeGetAdsPixelStatsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixelStatsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixelStatsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelStatsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelStatsParams) (*core.Cursor[objects.AdsPixelStatsResult], *core.Response, error) {
	var out core.Cursor[objects.AdsPixelStatsResult]
	call := GetAdsPixelStatsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixelStats(ctx context.Context, client *core.Client, id string, params GetAdsPixelStatsParams) (*core.Cursor[objects.AdsPixelStatsResult], error) {
	out, _, err := GetAdsPixelStatsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsPixelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelBatchCall(id string, params GetAdsPixelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsPixelBatchRequest(id string, params GetAdsPixelParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](GetAdsPixelBatchCall(id, params, options...))
}

func DecodeGetAdsPixelBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPixelParams) (*objects.AdsPixel, *core.Response, error) {
	var out objects.AdsPixel
	call := GetAdsPixelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPixel(ctx context.Context, client *core.Client, id string, params GetAdsPixelParams) (*objects.AdsPixel, error) {
	out, _, err := GetAdsPixelWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdsPixelParams struct {
	AutomaticMatchingFields       *[]enums.AdspixelAutomaticMatchingFields `facebook:"automatic_matching_fields"`
	DataUseSetting                *enums.AdspixelDataUseSetting            `facebook:"data_use_setting"`
	EnableAutomaticMatching       *bool                                    `facebook:"enable_automatic_matching"`
	FirstPartyCookieStatus        *enums.AdspixelFirstPartyCookieStatus    `facebook:"first_party_cookie_status"`
	Name                          *string                                  `facebook:"name"`
	ServerEventsBusinessIds       *[]core.ID                               `facebook:"server_events_business_ids"`
	ServerEventsBusinessIdsAdd    *[]string                                `facebook:"server_events_business_ids_add"`
	ServerEventsBusinessIdsRemove *[]string                                `facebook:"server_events_business_ids_remove"`
	Extra                         core.Params                              `facebook:"-"`
}

func (params UpdateAdsPixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AutomaticMatchingFields != nil {
		out["automatic_matching_fields"] = *params.AutomaticMatchingFields
	}
	if params.DataUseSetting != nil {
		out["data_use_setting"] = *params.DataUseSetting
	}
	if params.EnableAutomaticMatching != nil {
		out["enable_automatic_matching"] = *params.EnableAutomaticMatching
	}
	if params.FirstPartyCookieStatus != nil {
		out["first_party_cookie_status"] = *params.FirstPartyCookieStatus
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.ServerEventsBusinessIds != nil {
		out["server_events_business_ids"] = *params.ServerEventsBusinessIds
	}
	if params.ServerEventsBusinessIdsAdd != nil {
		out["server_events_business_ids_add"] = *params.ServerEventsBusinessIdsAdd
	}
	if params.ServerEventsBusinessIdsRemove != nil {
		out["server_events_business_ids_remove"] = *params.ServerEventsBusinessIdsRemove
	}
	return out
}

func UpdateAdsPixelBatchCall(id string, params UpdateAdsPixelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdsPixelBatchRequest(id string, params UpdateAdsPixelParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](UpdateAdsPixelBatchCall(id, params, options...))
}

func DecodeUpdateAdsPixelBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdsPixelWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdsPixelParams) (*objects.AdsPixel, *core.Response, error) {
	var out objects.AdsPixel
	call := UpdateAdsPixelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdsPixel(ctx context.Context, client *core.Client, id string, params UpdateAdsPixelParams) (*objects.AdsPixel, error) {
	out, _, err := UpdateAdsPixelWithResponse(ctx, client, id, params)
	return out, err
}
