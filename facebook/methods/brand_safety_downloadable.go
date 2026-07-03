package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBrandSafetyDownloadableParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBrandSafetyDownloadableParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBrandSafetyDownloadableBatchCall(id string, params GetBrandSafetyDownloadableParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBrandSafetyDownloadableBatchRequest(id string, params GetBrandSafetyDownloadableParams, options ...core.BatchOption) *core.BatchRequest[objects.BrandSafetyDownloadable] {
	return core.NewBatchRequest[objects.BrandSafetyDownloadable](GetBrandSafetyDownloadableBatchCall(id, params, options...))
}

func DecodeGetBrandSafetyDownloadableBatchResponse(response *core.BatchResponse) (*objects.BrandSafetyDownloadable, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BrandSafetyDownloadable
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBrandSafetyDownloadableWithResponse(ctx context.Context, client *core.Client, id string, params GetBrandSafetyDownloadableParams) (*objects.BrandSafetyDownloadable, *core.Response, error) {
	var out objects.BrandSafetyDownloadable
	call := GetBrandSafetyDownloadableBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBrandSafetyDownloadable(ctx context.Context, client *core.Client, id string, params GetBrandSafetyDownloadableParams) (*objects.BrandSafetyDownloadable, error) {
	out, _, err := GetBrandSafetyDownloadableWithResponse(ctx, client, id, params)
	return out, err
}
