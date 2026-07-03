package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessImageBatchCall(id string, params GetBusinessImageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessImageBatchRequest(id string, params GetBusinessImageParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessImage] {
	return core.NewBatchRequest[objects.BusinessImage](GetBusinessImageBatchCall(id, params, options...))
}

func DecodeGetBusinessImageBatchResponse(response *core.BatchResponse) (*objects.BusinessImage, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessImage
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessImageWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessImageParams) (*objects.BusinessImage, *core.Response, error) {
	var out objects.BusinessImage
	call := GetBusinessImageBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessImage(ctx context.Context, client *core.Client, id string, params GetBusinessImageParams) (*objects.BusinessImage, error) {
	out, _, err := GetBusinessImageWithResponse(ctx, client, id, params)
	return out, err
}
