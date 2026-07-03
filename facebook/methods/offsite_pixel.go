package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOffsitePixelParams struct {
	Value *uint64     `facebook:"value"`
	Extra core.Params `facebook:"-"`
}

func (params GetOffsitePixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Value != nil {
		out["value"] = *params.Value
	}
	return out
}

func GetOffsitePixelBatchCall(id string, params GetOffsitePixelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOffsitePixelBatchRequest(id string, params GetOffsitePixelParams, options ...core.BatchOption) *core.BatchRequest[objects.OffsitePixel] {
	return core.NewBatchRequest[objects.OffsitePixel](GetOffsitePixelBatchCall(id, params, options...))
}

func DecodeGetOffsitePixelBatchResponse(response *core.BatchResponse) (*objects.OffsitePixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OffsitePixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOffsitePixelWithResponse(ctx context.Context, client *core.Client, id string, params GetOffsitePixelParams) (*objects.OffsitePixel, *core.Response, error) {
	var out objects.OffsitePixel
	call := GetOffsitePixelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOffsitePixel(ctx context.Context, client *core.Client, id string, params GetOffsitePixelParams) (*objects.OffsitePixel, error) {
	out, _, err := GetOffsitePixelWithResponse(ctx, client, id, params)
	return out, err
}
