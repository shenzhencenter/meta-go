package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetImageCopyrightDisputeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetImageCopyrightDisputeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetImageCopyrightDisputeBatchCall(id string, params GetImageCopyrightDisputeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetImageCopyrightDisputeBatchRequest(id string, params GetImageCopyrightDisputeParams, options ...core.BatchOption) *core.BatchRequest[objects.ImageCopyrightDispute] {
	return core.NewBatchRequest[objects.ImageCopyrightDispute](GetImageCopyrightDisputeBatchCall(id, params, options...))
}

func DecodeGetImageCopyrightDisputeBatchResponse(response *core.BatchResponse) (*objects.ImageCopyrightDispute, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ImageCopyrightDispute
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetImageCopyrightDispute(ctx context.Context, client *core.Client, id string, params GetImageCopyrightDisputeParams) (*objects.ImageCopyrightDispute, error) {
	var out objects.ImageCopyrightDispute
	call := GetImageCopyrightDisputeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
