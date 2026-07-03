package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteInstagramBusinessAssetAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteInstagramBusinessAssetAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteInstagramBusinessAssetAgenciesBatchCall(id string, params DeleteInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewDeleteInstagramBusinessAssetAgenciesBatchRequest(id string, params DeleteInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteInstagramBusinessAssetAgenciesBatchCall(id, params, options...))
}

func DecodeDeleteInstagramBusinessAssetAgenciesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteInstagramBusinessAssetAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAgenciesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteInstagramBusinessAssetAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAgenciesParams) (*map[string]interface{}, error) {
	out, _, err := DeleteInstagramBusinessAssetAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetInstagramBusinessAssetAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetInstagramBusinessAssetAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetInstagramBusinessAssetAgenciesBatchCall(id string, params GetInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetInstagramBusinessAssetAgenciesBatchRequest(id string, params GetInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetInstagramBusinessAssetAgenciesBatchCall(id, params, options...))
}

func DecodeGetInstagramBusinessAssetAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetInstagramBusinessAssetAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetInstagramBusinessAssetAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetInstagramBusinessAssetAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateInstagramBusinessAssetAgenciesParams struct {
	Business       string                                                        `facebook:"business"`
	PermittedTasks []enums.InstagrambusinessassetagenciesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                                   `facebook:"-"`
}

func (params CreateInstagramBusinessAssetAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	out["permitted_tasks"] = params.PermittedTasks
	return out
}

func CreateInstagramBusinessAssetAgenciesBatchCall(id string, params CreateInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewCreateInstagramBusinessAssetAgenciesBatchRequest(id string, params CreateInstagramBusinessAssetAgenciesParams, options ...core.BatchOption) *core.BatchRequest[objects.InstagramBusinessAsset] {
	return core.NewBatchRequest[objects.InstagramBusinessAsset](CreateInstagramBusinessAssetAgenciesBatchCall(id, params, options...))
}

func DecodeCreateInstagramBusinessAssetAgenciesBatchResponse(response *core.BatchResponse) (*objects.InstagramBusinessAsset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.InstagramBusinessAsset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateInstagramBusinessAssetAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAgenciesParams) (*objects.InstagramBusinessAsset, *core.Response, error) {
	var out objects.InstagramBusinessAsset
	call := CreateInstagramBusinessAssetAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAgenciesParams) (*objects.InstagramBusinessAsset, error) {
	out, _, err := CreateInstagramBusinessAssetAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteInstagramBusinessAssetAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteInstagramBusinessAssetAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteInstagramBusinessAssetAssignedUsersBatchCall(id string, params DeleteInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewDeleteInstagramBusinessAssetAssignedUsersBatchRequest(id string, params DeleteInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteInstagramBusinessAssetAssignedUsersBatchCall(id, params, options...))
}

func DecodeDeleteInstagramBusinessAssetAssignedUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteInstagramBusinessAssetAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAssignedUsersParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteInstagramBusinessAssetAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAssignedUsersParams) (*map[string]interface{}, error) {
	out, _, err := DeleteInstagramBusinessAssetAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetInstagramBusinessAssetAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetInstagramBusinessAssetAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetInstagramBusinessAssetAssignedUsersBatchCall(id string, params GetInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetInstagramBusinessAssetAssignedUsersBatchRequest(id string, params GetInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetInstagramBusinessAssetAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetInstagramBusinessAssetAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
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

func GetInstagramBusinessAssetAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAssignedUsersParams) (*core.Cursor[objects.AssignedUser], *core.Response, error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetInstagramBusinessAssetAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	out, _, err := GetInstagramBusinessAssetAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateInstagramBusinessAssetAssignedUsersParams struct {
	Tasks *[]enums.InstagrambusinessassetassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                                        `facebook:"user"`
	Extra core.Params                                                `facebook:"-"`
}

func (params CreateInstagramBusinessAssetAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	out["user"] = params.User
	return out
}

func CreateInstagramBusinessAssetAssignedUsersBatchCall(id string, params CreateInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreateInstagramBusinessAssetAssignedUsersBatchRequest(id string, params CreateInstagramBusinessAssetAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.InstagramBusinessAsset] {
	return core.NewBatchRequest[objects.InstagramBusinessAsset](CreateInstagramBusinessAssetAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreateInstagramBusinessAssetAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.InstagramBusinessAsset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.InstagramBusinessAsset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateInstagramBusinessAssetAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAssignedUsersParams) (*objects.InstagramBusinessAsset, *core.Response, error) {
	var out objects.InstagramBusinessAsset
	call := CreateInstagramBusinessAssetAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAssignedUsersParams) (*objects.InstagramBusinessAsset, error) {
	out, _, err := CreateInstagramBusinessAssetAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetInstagramBusinessAssetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetInstagramBusinessAssetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetInstagramBusinessAssetBatchCall(id string, params GetInstagramBusinessAssetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetInstagramBusinessAssetBatchRequest(id string, params GetInstagramBusinessAssetParams, options ...core.BatchOption) *core.BatchRequest[objects.InstagramBusinessAsset] {
	return core.NewBatchRequest[objects.InstagramBusinessAsset](GetInstagramBusinessAssetBatchCall(id, params, options...))
}

func DecodeGetInstagramBusinessAssetBatchResponse(response *core.BatchResponse) (*objects.InstagramBusinessAsset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.InstagramBusinessAsset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetInstagramBusinessAssetWithResponse(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetParams) (*objects.InstagramBusinessAsset, *core.Response, error) {
	var out objects.InstagramBusinessAsset
	call := GetInstagramBusinessAssetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetInstagramBusinessAsset(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetParams) (*objects.InstagramBusinessAsset, error) {
	out, _, err := GetInstagramBusinessAssetWithResponse(ctx, client, id, params)
	return out, err
}
