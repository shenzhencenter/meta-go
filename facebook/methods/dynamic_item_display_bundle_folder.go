package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicItemDisplayBundleFolderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicItemDisplayBundleFolderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicItemDisplayBundleFolderBatchCall(id string, params GetDynamicItemDisplayBundleFolderParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDynamicItemDisplayBundleFolderBatchRequest(id string, params GetDynamicItemDisplayBundleFolderParams, options ...core.BatchOption) *core.BatchRequest[objects.DynamicItemDisplayBundleFolder] {
	return core.NewBatchRequest[objects.DynamicItemDisplayBundleFolder](GetDynamicItemDisplayBundleFolderBatchCall(id, params, options...))
}

func DecodeGetDynamicItemDisplayBundleFolderBatchResponse(response *core.BatchResponse) (*objects.DynamicItemDisplayBundleFolder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.DynamicItemDisplayBundleFolder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDynamicItemDisplayBundleFolderWithResponse(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleFolderParams) (*objects.DynamicItemDisplayBundleFolder, *core.Response, error) {
	var out objects.DynamicItemDisplayBundleFolder
	call := GetDynamicItemDisplayBundleFolderBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetDynamicItemDisplayBundleFolder(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleFolderParams) (*objects.DynamicItemDisplayBundleFolder, error) {
	out, _, err := GetDynamicItemDisplayBundleFolderWithResponse(ctx, client, id, params)
	return out, err
}
