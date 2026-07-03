package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBlindPigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBlindPigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBlindPigBatchCall(id string, params GetBlindPigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBlindPigBatchRequest(id string, params GetBlindPigParams, options ...core.BatchOption) *core.BatchRequest[objects.BlindPig] {
	return core.NewBatchRequest[objects.BlindPig](GetBlindPigBatchCall(id, params, options...))
}

func DecodeGetBlindPigBatchResponse(response *core.BatchResponse) (*objects.BlindPig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BlindPig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBlindPig(ctx context.Context, client *core.Client, id string, params GetBlindPigParams) (*objects.BlindPig, error) {
	var out objects.BlindPig
	call := GetBlindPigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
