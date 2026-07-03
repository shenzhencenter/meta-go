package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBrandRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBrandRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBrandRequestBatchCall(id string, params GetBrandRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBrandRequestBatchRequest(id string, params GetBrandRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.BrandRequest] {
	return core.NewBatchRequest[objects.BrandRequest](GetBrandRequestBatchCall(id, params, options...))
}

func DecodeGetBrandRequestBatchResponse(response *core.BatchResponse) (*objects.BrandRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BrandRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBrandRequest(ctx context.Context, client *core.Client, id string, params GetBrandRequestParams) (*objects.BrandRequest, error) {
	var out objects.BrandRequest
	call := GetBrandRequestBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
