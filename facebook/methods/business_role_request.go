package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeleteBusinessRoleRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteBusinessRoleRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params DeleteBusinessRoleRequestParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetBusinessRoleRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessRoleRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params GetBusinessRoleRequestParams) (*objects.BusinessRoleRequest, error) {
	var out objects.BusinessRoleRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateBusinessRoleRequestParams struct {
	Role  *enums.BusinessrolerequestRole    `facebook:"role"`
	Tasks *[]enums.BusinessrolerequestTasks `facebook:"tasks"`
	Extra core.Params                       `facebook:"-"`
}

func (params UpdateBusinessRoleRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Role != nil {
		out["role"] = *params.Role
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	return out
}

func UpdateBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params UpdateBusinessRoleRequestParams) (*objects.BusinessRoleRequest, error) {
	var out objects.BusinessRoleRequest
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
