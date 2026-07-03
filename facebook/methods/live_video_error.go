package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoErrorParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoErrorParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoErrorBatchCall(id string, params GetLiveVideoErrorParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLiveVideoErrorBatchRequest(id string, params GetLiveVideoErrorParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideoError] {
	return core.NewBatchRequest[objects.LiveVideoError](GetLiveVideoErrorBatchCall(id, params, options...))
}

func DecodeGetLiveVideoErrorBatchResponse(response *core.BatchResponse) (*objects.LiveVideoError, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideoError
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoErrorWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorParams) (*objects.LiveVideoError, *core.Response, error) {
	var out objects.LiveVideoError
	call := GetLiveVideoErrorBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoError(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorParams) (*objects.LiveVideoError, error) {
	out, _, err := GetLiveVideoErrorWithResponse(ctx, client, id, params)
	return out, err
}
