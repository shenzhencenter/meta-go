package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdExportPresetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdExportPresetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdExportPresetBatchCall(id string, params GetAdExportPresetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdExportPresetBatchRequest(id string, params GetAdExportPresetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdExportPreset] {
	return core.NewBatchRequest[objects.AdExportPreset](GetAdExportPresetBatchCall(id, params, options...))
}

func DecodeGetAdExportPresetBatchResponse(response *core.BatchResponse) (*objects.AdExportPreset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdExportPreset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdExportPreset(ctx context.Context, client *core.Client, id string, params GetAdExportPresetParams) (*objects.AdExportPreset, error) {
	var out objects.AdExportPreset
	call := GetAdExportPresetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
