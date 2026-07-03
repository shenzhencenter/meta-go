package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeAssetSourceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeAssetSourceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeAssetSourceBatchCall(id string, params GetWebsiteCreativeAssetSourceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWebsiteCreativeAssetSourceBatchRequest(id string, params GetWebsiteCreativeAssetSourceParams, options ...core.BatchOption) *core.BatchRequest[objects.WebsiteCreativeAssetSource] {
	return core.NewBatchRequest[objects.WebsiteCreativeAssetSource](GetWebsiteCreativeAssetSourceBatchCall(id, params, options...))
}

func DecodeGetWebsiteCreativeAssetSourceBatchResponse(response *core.BatchResponse) (*objects.WebsiteCreativeAssetSource, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WebsiteCreativeAssetSource
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWebsiteCreativeAssetSourceWithResponse(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSourceParams) (*objects.WebsiteCreativeAssetSource, *core.Response, error) {
	var out objects.WebsiteCreativeAssetSource
	call := GetWebsiteCreativeAssetSourceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWebsiteCreativeAssetSource(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSourceParams) (*objects.WebsiteCreativeAssetSource, error) {
	out, _, err := GetWebsiteCreativeAssetSourceWithResponse(ctx, client, id, params)
	return out, err
}
