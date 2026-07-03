package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGBCAdsPermissionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGBCAdsPermissionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGBCAdsPermissionBatchCall(id string, params GetIGBCAdsPermissionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGBCAdsPermissionBatchRequest(id string, params GetIGBCAdsPermissionParams, options ...core.BatchOption) *core.BatchRequest[objects.IGBCAdsPermission] {
	return core.NewBatchRequest[objects.IGBCAdsPermission](GetIGBCAdsPermissionBatchCall(id, params, options...))
}

func DecodeGetIGBCAdsPermissionBatchResponse(response *core.BatchResponse) (*objects.IGBCAdsPermission, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGBCAdsPermission
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGBCAdsPermissionWithResponse(ctx context.Context, client *core.Client, id string, params GetIGBCAdsPermissionParams) (*objects.IGBCAdsPermission, *core.Response, error) {
	var out objects.IGBCAdsPermission
	call := GetIGBCAdsPermissionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGBCAdsPermission(ctx context.Context, client *core.Client, id string, params GetIGBCAdsPermissionParams) (*objects.IGBCAdsPermission, error) {
	out, _, err := GetIGBCAdsPermissionWithResponse(ctx, client, id, params)
	return out, err
}
