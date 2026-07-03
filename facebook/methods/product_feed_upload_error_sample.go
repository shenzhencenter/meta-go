package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedUploadErrorSampleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorSampleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorSampleBatchCall(id string, params GetProductFeedUploadErrorSampleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedUploadErrorSampleBatchRequest(id string, params GetProductFeedUploadErrorSampleParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedUploadErrorSample] {
	return core.NewBatchRequest[objects.ProductFeedUploadErrorSample](GetProductFeedUploadErrorSampleBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadErrorSampleBatchResponse(response *core.BatchResponse) (*objects.ProductFeedUploadErrorSample, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedUploadErrorSample
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadErrorSampleWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSampleParams) (*objects.ProductFeedUploadErrorSample, *core.Response, error) {
	var out objects.ProductFeedUploadErrorSample
	call := GetProductFeedUploadErrorSampleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedUploadErrorSample(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSampleParams) (*objects.ProductFeedUploadErrorSample, error) {
	out, _, err := GetProductFeedUploadErrorSampleWithResponse(ctx, client, id, params)
	return out, err
}
