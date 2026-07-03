package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func DeleteInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
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

func GetInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
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

func CreateInstagramBusinessAssetAgencies(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAgenciesParams) (*objects.InstagramBusinessAsset, error) {
	var out objects.InstagramBusinessAsset
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
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

func DeleteInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteInstagramBusinessAssetAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
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

func GetInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
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

func CreateInstagramBusinessAssetAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateInstagramBusinessAssetAssignedUsersParams) (*objects.InstagramBusinessAsset, error) {
	var out objects.InstagramBusinessAsset
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
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

func GetInstagramBusinessAsset(ctx context.Context, client *core.Client, id string, params GetInstagramBusinessAssetParams) (*objects.InstagramBusinessAsset, error) {
	var out objects.InstagramBusinessAsset
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
