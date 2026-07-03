package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func DeleteBusinessRoleRequestBatchCall(id string, params DeleteBusinessRoleRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteBusinessRoleRequestBatchRequest(id string, params DeleteBusinessRoleRequestParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessRoleRequestBatchCall(id, params, options...))
}

func DecodeDeleteBusinessRoleRequestBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessRoleRequestWithResponse(ctx context.Context, client *core.Client, id string, params DeleteBusinessRoleRequestParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteBusinessRoleRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params DeleteBusinessRoleRequestParams) (*map[string]interface{}, error) {
	out, _, err := DeleteBusinessRoleRequestWithResponse(ctx, client, id, params)
	return out, err
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

func GetBusinessRoleRequestBatchCall(id string, params GetBusinessRoleRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessRoleRequestBatchRequest(id string, params GetBusinessRoleRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessRoleRequest] {
	return core.NewBatchRequest[objects.BusinessRoleRequest](GetBusinessRoleRequestBatchCall(id, params, options...))
}

func DecodeGetBusinessRoleRequestBatchResponse(response *core.BatchResponse) (*objects.BusinessRoleRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessRoleRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessRoleRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessRoleRequestParams) (*objects.BusinessRoleRequest, *core.Response, error) {
	var out objects.BusinessRoleRequest
	call := GetBusinessRoleRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params GetBusinessRoleRequestParams) (*objects.BusinessRoleRequest, error) {
	out, _, err := GetBusinessRoleRequestWithResponse(ctx, client, id, params)
	return out, err
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

func UpdateBusinessRoleRequestBatchCall(id string, params UpdateBusinessRoleRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessRoleRequestBatchRequest(id string, params UpdateBusinessRoleRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessRoleRequest] {
	return core.NewBatchRequest[objects.BusinessRoleRequest](UpdateBusinessRoleRequestBatchCall(id, params, options...))
}

func DecodeUpdateBusinessRoleRequestBatchResponse(response *core.BatchResponse) (*objects.BusinessRoleRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessRoleRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusinessRoleRequestWithResponse(ctx context.Context, client *core.Client, id string, params UpdateBusinessRoleRequestParams) (*objects.BusinessRoleRequest, *core.Response, error) {
	var out objects.BusinessRoleRequest
	call := UpdateBusinessRoleRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateBusinessRoleRequest(ctx context.Context, client *core.Client, id string, params UpdateBusinessRoleRequestParams) (*objects.BusinessRoleRequest, error) {
	out, _, err := UpdateBusinessRoleRequestWithResponse(ctx, client, id, params)
	return out, err
}
