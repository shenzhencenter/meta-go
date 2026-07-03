package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMediaCopyrightAttributionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMediaCopyrightAttributionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMediaCopyrightAttributionBatchCall(id string, params GetMediaCopyrightAttributionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMediaCopyrightAttributionBatchRequest(id string, params GetMediaCopyrightAttributionParams, options ...core.BatchOption) *core.BatchRequest[objects.MediaCopyrightAttribution] {
	return core.NewBatchRequest[objects.MediaCopyrightAttribution](GetMediaCopyrightAttributionBatchCall(id, params, options...))
}

func DecodeGetMediaCopyrightAttributionBatchResponse(response *core.BatchResponse) (*objects.MediaCopyrightAttribution, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MediaCopyrightAttribution
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMediaCopyrightAttributionWithResponse(ctx context.Context, client *core.Client, id string, params GetMediaCopyrightAttributionParams) (*objects.MediaCopyrightAttribution, *core.Response, error) {
	var out objects.MediaCopyrightAttribution
	call := GetMediaCopyrightAttributionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetMediaCopyrightAttribution(ctx context.Context, client *core.Client, id string, params GetMediaCopyrightAttributionParams) (*objects.MediaCopyrightAttribution, error) {
	out, _, err := GetMediaCopyrightAttributionWithResponse(ctx, client, id, params)
	return out, err
}
