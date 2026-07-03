package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsMcmeConversionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsMcmeConversionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsMcmeConversionBatchCall(id string, params GetAdsMcmeConversionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsMcmeConversionBatchRequest(id string, params GetAdsMcmeConversionParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsMcmeConversion] {
	return core.NewBatchRequest[objects.AdsMcmeConversion](GetAdsMcmeConversionBatchCall(id, params, options...))
}

func DecodeGetAdsMcmeConversionBatchResponse(response *core.BatchResponse) (*objects.AdsMcmeConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsMcmeConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsMcmeConversion(ctx context.Context, client *core.Client, id string, params GetAdsMcmeConversionParams) (*objects.AdsMcmeConversion, error) {
	var out objects.AdsMcmeConversion
	call := GetAdsMcmeConversionBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
