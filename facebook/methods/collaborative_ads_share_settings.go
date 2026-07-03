package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCollaborativeAdsShareSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCollaborativeAdsShareSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCollaborativeAdsShareSettingsBatchCall(id string, params GetCollaborativeAdsShareSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCollaborativeAdsShareSettingsBatchRequest(id string, params GetCollaborativeAdsShareSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CollaborativeAdsShareSettings] {
	return core.NewBatchRequest[objects.CollaborativeAdsShareSettings](GetCollaborativeAdsShareSettingsBatchCall(id, params, options...))
}

func DecodeGetCollaborativeAdsShareSettingsBatchResponse(response *core.BatchResponse) (*objects.CollaborativeAdsShareSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CollaborativeAdsShareSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCollaborativeAdsShareSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetCollaborativeAdsShareSettingsParams) (*objects.CollaborativeAdsShareSettings, *core.Response, error) {
	var out objects.CollaborativeAdsShareSettings
	call := GetCollaborativeAdsShareSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCollaborativeAdsShareSettings(ctx context.Context, client *core.Client, id string, params GetCollaborativeAdsShareSettingsParams) (*objects.CollaborativeAdsShareSettings, error) {
	out, _, err := GetCollaborativeAdsShareSettingsWithResponse(ctx, client, id, params)
	return out, err
}
