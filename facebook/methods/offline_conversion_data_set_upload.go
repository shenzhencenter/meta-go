package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOfflineConversionDataSetUploadProgressParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetUploadProgressParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetUploadProgress(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadProgressParams) (*core.Cursor[objects.SignalsUploadProgress], error) {
	var out core.Cursor[objects.SignalsUploadProgress]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "progress"), params.ToParams(), &out)
	return &out, err
}

type GetOfflineConversionDataSetUploadPullSessionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetUploadPullSessionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetUploadPullSessions(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadPullSessionsParams) (*core.Cursor[objects.PartnerIntegrationPullSession], error) {
	var out core.Cursor[objects.PartnerIntegrationPullSession]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pull_sessions"), params.ToParams(), &out)
	return &out, err
}

type GetOfflineConversionDataSetUploadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetUploadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetUpload(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadParams) (*objects.OfflineConversionDataSetUpload, error) {
	var out objects.OfflineConversionDataSetUpload
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
