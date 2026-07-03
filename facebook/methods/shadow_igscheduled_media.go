package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShadowIGScheduledMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShadowIGScheduledMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShadowIGScheduledMediaBatchCall(id string, params GetShadowIGScheduledMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetShadowIGScheduledMediaBatchRequest(id string, params GetShadowIGScheduledMediaParams, options ...core.BatchOption) *core.BatchRequest[objects.ShadowIGScheduledMedia] {
	return core.NewBatchRequest[objects.ShadowIGScheduledMedia](GetShadowIGScheduledMediaBatchCall(id, params, options...))
}

func DecodeGetShadowIGScheduledMediaBatchResponse(response *core.BatchResponse) (*objects.ShadowIGScheduledMedia, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ShadowIGScheduledMedia
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShadowIGScheduledMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGScheduledMediaParams) (*objects.ShadowIGScheduledMedia, error) {
	var out objects.ShadowIGScheduledMedia
	call := GetShadowIGScheduledMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
