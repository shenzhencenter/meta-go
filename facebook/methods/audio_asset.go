package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAudioAssetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioAssetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioAssetBatchCall(id string, params GetAudioAssetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAudioAssetBatchRequest(id string, params GetAudioAssetParams, options ...core.BatchOption) *core.BatchRequest[objects.AudioAsset] {
	return core.NewBatchRequest[objects.AudioAsset](GetAudioAssetBatchCall(id, params, options...))
}

func DecodeGetAudioAssetBatchResponse(response *core.BatchResponse) (*objects.AudioAsset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AudioAsset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAudioAssetWithResponse(ctx context.Context, client *core.Client, id string, params GetAudioAssetParams) (*objects.AudioAsset, *core.Response, error) {
	var out objects.AudioAsset
	call := GetAudioAssetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAudioAsset(ctx context.Context, client *core.Client, id string, params GetAudioAssetParams) (*objects.AudioAsset, error) {
	out, _, err := GetAudioAssetWithResponse(ctx, client, id, params)
	return out, err
}
