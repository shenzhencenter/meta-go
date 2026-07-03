package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWithAsset3DParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWithAsset3DParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWithAsset3DBatchCall(id string, params GetWithAsset3DParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWithAsset3DBatchRequest(id string, params GetWithAsset3DParams, options ...core.BatchOption) *core.BatchRequest[objects.WithAsset3D] {
	return core.NewBatchRequest[objects.WithAsset3D](GetWithAsset3DBatchCall(id, params, options...))
}

func DecodeGetWithAsset3DBatchResponse(response *core.BatchResponse) (*objects.WithAsset3D, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WithAsset3D
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWithAsset3DWithResponse(ctx context.Context, client *core.Client, id string, params GetWithAsset3DParams) (*objects.WithAsset3D, *core.Response, error) {
	var out objects.WithAsset3D
	call := GetWithAsset3DBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWithAsset3D(ctx context.Context, client *core.Client, id string, params GetWithAsset3DParams) (*objects.WithAsset3D, error) {
	out, _, err := GetWithAsset3DWithResponse(ctx, client, id, params)
	return out, err
}
