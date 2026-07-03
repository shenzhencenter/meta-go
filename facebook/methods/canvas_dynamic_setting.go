package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCanvasDynamicSettingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasDynamicSettingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasDynamicSettingBatchCall(id string, params GetCanvasDynamicSettingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCanvasDynamicSettingBatchRequest(id string, params GetCanvasDynamicSettingParams, options ...core.BatchOption) *core.BatchRequest[objects.CanvasDynamicSetting] {
	return core.NewBatchRequest[objects.CanvasDynamicSetting](GetCanvasDynamicSettingBatchCall(id, params, options...))
}

func DecodeGetCanvasDynamicSettingBatchResponse(response *core.BatchResponse) (*objects.CanvasDynamicSetting, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CanvasDynamicSetting
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCanvasDynamicSetting(ctx context.Context, client *core.Client, id string, params GetCanvasDynamicSettingParams) (*objects.CanvasDynamicSetting, error) {
	var out objects.CanvasDynamicSetting
	call := GetCanvasDynamicSettingBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
