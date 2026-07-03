package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCreatorAssetCreativeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCreatorAssetCreativeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCreatorAssetCreativeBatchCall(id string, params GetCreatorAssetCreativeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCreatorAssetCreativeBatchRequest(id string, params GetCreatorAssetCreativeParams, options ...core.BatchOption) *core.BatchRequest[objects.CreatorAssetCreative] {
	return core.NewBatchRequest[objects.CreatorAssetCreative](GetCreatorAssetCreativeBatchCall(id, params, options...))
}

func DecodeGetCreatorAssetCreativeBatchResponse(response *core.BatchResponse) (*objects.CreatorAssetCreative, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CreatorAssetCreative
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCreatorAssetCreativeWithResponse(ctx context.Context, client *core.Client, id string, params GetCreatorAssetCreativeParams) (*objects.CreatorAssetCreative, *core.Response, error) {
	var out objects.CreatorAssetCreative
	call := GetCreatorAssetCreativeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCreatorAssetCreative(ctx context.Context, client *core.Client, id string, params GetCreatorAssetCreativeParams) (*objects.CreatorAssetCreative, error) {
	out, _, err := GetCreatorAssetCreativeWithResponse(ctx, client, id, params)
	return out, err
}
