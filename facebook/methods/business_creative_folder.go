package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessCreativeFolderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCreativeFolderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCreativeFolderBatchCall(id string, params GetBusinessCreativeFolderParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessCreativeFolderBatchRequest(id string, params GetBusinessCreativeFolderParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessCreativeFolder] {
	return core.NewBatchRequest[objects.BusinessCreativeFolder](GetBusinessCreativeFolderBatchCall(id, params, options...))
}

func DecodeGetBusinessCreativeFolderBatchResponse(response *core.BatchResponse) (*objects.BusinessCreativeFolder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessCreativeFolder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCreativeFolderWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessCreativeFolderParams) (*objects.BusinessCreativeFolder, *core.Response, error) {
	var out objects.BusinessCreativeFolder
	call := GetBusinessCreativeFolderBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessCreativeFolder(ctx context.Context, client *core.Client, id string, params GetBusinessCreativeFolderParams) (*objects.BusinessCreativeFolder, error) {
	out, _, err := GetBusinessCreativeFolderWithResponse(ctx, client, id, params)
	return out, err
}
