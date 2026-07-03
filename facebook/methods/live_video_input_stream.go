package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoInputStreamParams struct {
	TargetToken *string     `facebook:"target_token"`
	Extra       core.Params `facebook:"-"`
}

func (params GetLiveVideoInputStreamParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetToken != nil {
		out["target_token"] = *params.TargetToken
	}
	return out
}

func GetLiveVideoInputStreamBatchCall(id string, params GetLiveVideoInputStreamParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLiveVideoInputStreamBatchRequest(id string, params GetLiveVideoInputStreamParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideoInputStream] {
	return core.NewBatchRequest[objects.LiveVideoInputStream](GetLiveVideoInputStreamBatchCall(id, params, options...))
}

func DecodeGetLiveVideoInputStreamBatchResponse(response *core.BatchResponse) (*objects.LiveVideoInputStream, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideoInputStream
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoInputStream(ctx context.Context, client *core.Client, id string, params GetLiveVideoInputStreamParams) (*objects.LiveVideoInputStream, error) {
	var out objects.LiveVideoInputStream
	call := GetLiveVideoInputStreamBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
