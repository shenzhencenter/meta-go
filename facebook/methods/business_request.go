package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessRequestBatchCall(id string, params GetBusinessRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessRequestBatchRequest(id string, params GetBusinessRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessRequest] {
	return core.NewBatchRequest[objects.BusinessRequest](GetBusinessRequestBatchCall(id, params, options...))
}

func DecodeGetBusinessRequestBatchResponse(response *core.BatchResponse) (*objects.BusinessRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessRequest(ctx context.Context, client *core.Client, id string, params GetBusinessRequestParams) (*objects.BusinessRequest, error) {
	var out objects.BusinessRequest
	call := GetBusinessRequestBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
