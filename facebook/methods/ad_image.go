package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdImageBatchCall(id string, params GetAdImageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdImageBatchRequest(id string, params GetAdImageParams, options ...core.BatchOption) *core.BatchRequest[objects.AdImage] {
	return core.NewBatchRequest[objects.AdImage](GetAdImageBatchCall(id, params, options...))
}

func DecodeGetAdImageBatchResponse(response *core.BatchResponse) (*objects.AdImage, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdImage
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdImageWithResponse(ctx context.Context, client *core.Client, id string, params GetAdImageParams) (*objects.AdImage, *core.Response, error) {
	var out objects.AdImage
	call := GetAdImageBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdImage(ctx context.Context, client *core.Client, id string, params GetAdImageParams) (*objects.AdImage, error) {
	out, _, err := GetAdImageWithResponse(ctx, client, id, params)
	return out, err
}
