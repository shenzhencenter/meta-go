package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetUserContextParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserContextParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserContextBatchCall(id string, params GetUserContextParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetUserContextBatchRequest(id string, params GetUserContextParams, options ...core.BatchOption) *core.BatchRequest[objects.UserContext] {
	return core.NewBatchRequest[objects.UserContext](GetUserContextBatchCall(id, params, options...))
}

func DecodeGetUserContextBatchResponse(response *core.BatchResponse) (*objects.UserContext, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.UserContext
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserContext(ctx context.Context, client *core.Client, id string, params GetUserContextParams) (*objects.UserContext, error) {
	var out objects.UserContext
	call := GetUserContextBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
