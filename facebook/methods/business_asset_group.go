package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteBusinessAssetGroupAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteBusinessAssetGroupAssignedUsersBatchCall(id string, params DeleteBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupAssignedUsersBatchRequest(id string, params DeleteBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupAssignedUsersBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupAssignedUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupAssignedUsersParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupAssignedUsersParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetBusinessAssetGroupAssignedUsersBatchCall(id string, params GetBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupAssignedUsersBatchRequest(id string, params GetBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetBusinessAssetGroupAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
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

func GetBusinessAssetGroupAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupAssignedUsersParams) (*core.Cursor[objects.AssignedUser], *core.Response, error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetBusinessAssetGroupAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	out, _, err := GetBusinessAssetGroupAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupAssignedUsersParams struct {
	AdaccountTasks                *[]enums.BusinessassetgroupassignedUsersAdaccountTasksEnumParam                `facebook:"adaccount_tasks"`
	OfflineConversionDataSetTasks *[]enums.BusinessassetgroupassignedUsersOfflineConversionDataSetTasksEnumParam `facebook:"offline_conversion_data_set_tasks"`
	PageTasks                     *[]enums.BusinessassetgroupassignedUsersPageTasksEnumParam                     `facebook:"page_tasks"`
	PixelTasks                    *[]enums.BusinessassetgroupassignedUsersPixelTasksEnumParam                    `facebook:"pixel_tasks"`
	User                          int                                                                            `facebook:"user"`
	Extra                         core.Params                                                                    `facebook:"-"`
}

func (params CreateBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdaccountTasks != nil {
		out["adaccount_tasks"] = *params.AdaccountTasks
	}
	if params.OfflineConversionDataSetTasks != nil {
		out["offline_conversion_data_set_tasks"] = *params.OfflineConversionDataSetTasks
	}
	if params.PageTasks != nil {
		out["page_tasks"] = *params.PageTasks
	}
	if params.PixelTasks != nil {
		out["pixel_tasks"] = *params.PixelTasks
	}
	out["user"] = params.User
	return out
}

func CreateBusinessAssetGroupAssignedUsersBatchCall(id string, params CreateBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupAssignedUsersBatchRequest(id string, params CreateBusinessAssetGroupAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupAssignedUsersParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupAssignedUsersParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedAdaccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedAdaccountsBatchCall(id string, params DeleteBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedAdaccountsBatchRequest(id string, params DeleteBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedAdaccountsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedAdaccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedAdaccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedAdaccountsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedAdaccountsBatchCall(id string, params GetBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedAdaccountsBatchRequest(id string, params GetBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetBusinessAssetGroupContainedAdaccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetBusinessAssetGroupContainedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetBusinessAssetGroupContainedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetBusinessAssetGroupContainedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedAdaccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedAdaccountsBatchCall(id string, params CreateBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedAdaccountsBatchRequest(id string, params CreateBusinessAssetGroupContainedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedAdaccountsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedAdaccountsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedAdaccountsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedAdaccountsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedApplicationsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedApplicationsBatchCall(id string, params DeleteBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_applications"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedApplicationsBatchRequest(id string, params DeleteBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedApplicationsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedApplicationsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedApplicationsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedApplicationsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedApplicationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedApplicationsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedApplicationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedApplicationsBatchCall(id string, params GetBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_applications"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedApplicationsBatchRequest(id string, params GetBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetBusinessAssetGroupContainedApplicationsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedApplicationsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedApplicationsParams) (*core.Cursor[objects.Application], *core.Response, error) {
	var out core.Cursor[objects.Application]
	call := GetBusinessAssetGroupContainedApplicationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedApplicationsParams) (*core.Cursor[objects.Application], error) {
	out, _, err := GetBusinessAssetGroupContainedApplicationsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedApplicationsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedApplicationsBatchCall(id string, params CreateBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_applications"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedApplicationsBatchRequest(id string, params CreateBusinessAssetGroupContainedApplicationsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedApplicationsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedApplicationsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedApplicationsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedApplicationsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedApplicationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedApplicationsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedApplicationsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedCustomConversionsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedCustomConversionsBatchCall(id string, params DeleteBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedCustomConversionsBatchRequest(id string, params DeleteBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedCustomConversionsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedCustomConversionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedCustomConversionsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedCustomConversionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedCustomConversionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedCustomConversionsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedCustomConversionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedCustomConversionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedCustomConversionsBatchCall(id string, params GetBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedCustomConversionsBatchRequest(id string, params GetBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomConversion]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomConversion]](GetBusinessAssetGroupContainedCustomConversionsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedCustomConversionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomConversion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomConversion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedCustomConversionsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedCustomConversionsParams) (*core.Cursor[objects.CustomConversion], *core.Response, error) {
	var out core.Cursor[objects.CustomConversion]
	call := GetBusinessAssetGroupContainedCustomConversionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedCustomConversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	out, _, err := GetBusinessAssetGroupContainedCustomConversionsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedCustomConversionsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedCustomConversionsBatchCall(id string, params CreateBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedCustomConversionsBatchRequest(id string, params CreateBusinessAssetGroupContainedCustomConversionsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedCustomConversionsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedCustomConversionsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedCustomConversionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedCustomConversionsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedCustomConversionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedCustomConversionsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedCustomConversionsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedInstagramAccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedInstagramAccountsBatchCall(id string, params DeleteBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedInstagramAccountsBatchRequest(id string, params DeleteBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedInstagramAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedInstagramAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedInstagramAccountsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedInstagramAccountsBatchCall(id string, params GetBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedInstagramAccountsBatchRequest(id string, params GetBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedInstagramAccountsParams) (*core.Cursor[objects.IGUser], *core.Response, error) {
	var out core.Cursor[objects.IGUser]
	call := GetBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	out, _, err := GetBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedInstagramAccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedInstagramAccountsBatchCall(id string, params CreateBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedInstagramAccountsBatchRequest(id string, params CreateBusinessAssetGroupContainedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedInstagramAccountsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedInstagramAccountsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedInstagramAccountsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedPagesParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedPagesBatchCall(id string, params DeleteBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_pages"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedPagesBatchRequest(id string, params DeleteBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedPagesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedPagesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedPagesWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPagesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPagesParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedPagesBatchCall(id string, params GetBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_pages"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedPagesBatchRequest(id string, params GetBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetBusinessAssetGroupContainedPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPagesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetBusinessAssetGroupContainedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPagesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetBusinessAssetGroupContainedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedPagesParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedPagesBatchCall(id string, params CreateBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_pages"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedPagesBatchRequest(id string, params CreateBusinessAssetGroupContainedPagesParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedPagesBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedPagesBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedPagesWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPagesParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPagesParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedPixelsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedPixelsBatchCall(id string, params DeleteBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_pixels"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedPixelsBatchRequest(id string, params DeleteBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedPixelsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedPixelsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedPixelsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPixelsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedPixelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPixelsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedPixelsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedPixelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedPixelsBatchCall(id string, params GetBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_pixels"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedPixelsBatchRequest(id string, params GetBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetBusinessAssetGroupContainedPixelsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedPixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedPixelsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPixelsParams) (*core.Cursor[objects.AdsPixel], *core.Response, error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetBusinessAssetGroupContainedPixelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	out, _, err := GetBusinessAssetGroupContainedPixelsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedPixelsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedPixelsBatchCall(id string, params CreateBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_pixels"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedPixelsBatchRequest(id string, params CreateBusinessAssetGroupContainedPixelsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedPixelsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedPixelsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedPixelsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPixelsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedPixelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPixelsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedPixelsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteBusinessAssetGroupContainedProductCatalogsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedProductCatalogsBatchCall(id string, params DeleteBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), options...)
}

func NewDeleteBusinessAssetGroupContainedProductCatalogsBatchRequest(id string, params DeleteBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAssetGroupContainedProductCatalogsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAssetGroupContainedProductCatalogsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAssetGroupContainedProductCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedProductCatalogsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessAssetGroupContainedProductCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedProductCatalogsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessAssetGroupContainedProductCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupContainedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedProductCatalogsBatchCall(id string, params GetBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupContainedProductCatalogsBatchRequest(id string, params GetBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetBusinessAssetGroupContainedProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupContainedProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalog]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupContainedProductCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], *core.Response, error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetBusinessAssetGroupContainedProductCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	out, _, err := GetBusinessAssetGroupContainedProductCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateBusinessAssetGroupContainedProductCatalogsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedProductCatalogsBatchCall(id string, params CreateBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), options...)
}

func NewCreateBusinessAssetGroupContainedProductCatalogsBatchRequest(id string, params CreateBusinessAssetGroupContainedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](CreateBusinessAssetGroupContainedProductCatalogsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAssetGroupContainedProductCatalogsBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAssetGroupContainedProductCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedProductCatalogsParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := CreateBusinessAssetGroupContainedProductCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedProductCatalogsParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := CreateBusinessAssetGroupContainedProductCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type GetBusinessAssetGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupBatchCall(id string, params GetBusinessAssetGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessAssetGroupBatchRequest(id string, params GetBusinessAssetGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](GetBusinessAssetGroupBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetGroupBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetGroupWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := GetBusinessAssetGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAssetGroup(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := GetBusinessAssetGroupWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateBusinessAssetGroupParams struct {
	Name  *string     `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateBusinessAssetGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateBusinessAssetGroupBatchCall(id string, params UpdateBusinessAssetGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessAssetGroupBatchRequest(id string, params UpdateBusinessAssetGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetGroup] {
	return core.NewBatchRequest[objects.BusinessAssetGroup](UpdateBusinessAssetGroupBatchCall(id, params, options...))
}

func DecodeUpdateBusinessAssetGroupBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusinessAssetGroupWithResponse(ctx context.Context, client *core.Client, id string, params UpdateBusinessAssetGroupParams) (*objects.BusinessAssetGroup, *core.Response, error) {
	var out objects.BusinessAssetGroup
	call := UpdateBusinessAssetGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateBusinessAssetGroup(ctx context.Context, client *core.Client, id string, params UpdateBusinessAssetGroupParams) (*objects.BusinessAssetGroup, error) {
	out, _, err := UpdateBusinessAssetGroupWithResponse(ctx, client, id, params)
	return out, err
}
