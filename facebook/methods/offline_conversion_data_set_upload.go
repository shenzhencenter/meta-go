package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetOfflineConversionDataSetUploadProgressBatchCall(id string, params GetOfflineConversionDataSetUploadProgressParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "progress"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetUploadProgressBatchRequest(id string, params GetOfflineConversionDataSetUploadProgressParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.SignalsUploadProgress]] {
	return core.NewBatchRequest[core.Cursor[objects.SignalsUploadProgress]](GetOfflineConversionDataSetUploadProgressBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetUploadProgressBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.SignalsUploadProgress], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.SignalsUploadProgress]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetUploadProgressWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadProgressParams) (*core.Cursor[objects.SignalsUploadProgress], *core.Response, error) {
	var out core.Cursor[objects.SignalsUploadProgress]
	call := GetOfflineConversionDataSetUploadProgressBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetUploadProgress(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadProgressParams) (*core.Cursor[objects.SignalsUploadProgress], error) {
	out, _, err := GetOfflineConversionDataSetUploadProgressWithResponse(ctx, client, id, params)
	return out, err
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

func GetOfflineConversionDataSetUploadPullSessionsBatchCall(id string, params GetOfflineConversionDataSetUploadPullSessionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pull_sessions"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetUploadPullSessionsBatchRequest(id string, params GetOfflineConversionDataSetUploadPullSessionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PartnerIntegrationPullSession]] {
	return core.NewBatchRequest[core.Cursor[objects.PartnerIntegrationPullSession]](GetOfflineConversionDataSetUploadPullSessionsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetUploadPullSessionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PartnerIntegrationPullSession], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PartnerIntegrationPullSession]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetUploadPullSessionsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadPullSessionsParams) (*core.Cursor[objects.PartnerIntegrationPullSession], *core.Response, error) {
	var out core.Cursor[objects.PartnerIntegrationPullSession]
	call := GetOfflineConversionDataSetUploadPullSessionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetUploadPullSessions(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadPullSessionsParams) (*core.Cursor[objects.PartnerIntegrationPullSession], error) {
	out, _, err := GetOfflineConversionDataSetUploadPullSessionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetOfflineConversionDataSetUploadBatchCall(id string, params GetOfflineConversionDataSetUploadParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetUploadBatchRequest(id string, params GetOfflineConversionDataSetUploadParams, options ...core.BatchOption) *core.BatchRequest[objects.OfflineConversionDataSetUpload] {
	return core.NewBatchRequest[objects.OfflineConversionDataSetUpload](GetOfflineConversionDataSetUploadBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetUploadBatchResponse(response *core.BatchResponse) (*objects.OfflineConversionDataSetUpload, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OfflineConversionDataSetUpload
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetUploadWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadParams) (*objects.OfflineConversionDataSetUpload, *core.Response, error) {
	var out objects.OfflineConversionDataSetUpload
	call := GetOfflineConversionDataSetUploadBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetUpload(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadParams) (*objects.OfflineConversionDataSetUpload, error) {
	out, _, err := GetOfflineConversionDataSetUploadWithResponse(ctx, client, id, params)
	return out, err
}
