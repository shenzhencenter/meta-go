package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessProjectParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessProjectParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessProjectBatchCall(id string, params GetBusinessProjectParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessProjectBatchRequest(id string, params GetBusinessProjectParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessProject] {
	return core.NewBatchRequest[objects.BusinessProject](GetBusinessProjectBatchCall(id, params, options...))
}

func DecodeGetBusinessProjectBatchResponse(response *core.BatchResponse) (*objects.BusinessProject, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessProject
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessProject(ctx context.Context, client *core.Client, id string, params GetBusinessProjectParams) (*objects.BusinessProject, error) {
	var out objects.BusinessProject
	call := GetBusinessProjectBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
