package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAudioIsrcParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioIsrcParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioIsrcBatchCall(id string, params GetAudioIsrcParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAudioIsrcBatchRequest(id string, params GetAudioIsrcParams, options ...core.BatchOption) *core.BatchRequest[objects.AudioIsrc] {
	return core.NewBatchRequest[objects.AudioIsrc](GetAudioIsrcBatchCall(id, params, options...))
}

func DecodeGetAudioIsrcBatchResponse(response *core.BatchResponse) (*objects.AudioIsrc, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AudioIsrc
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAudioIsrc(ctx context.Context, client *core.Client, id string, params GetAudioIsrcParams) (*objects.AudioIsrc, error) {
	var out objects.AudioIsrc
	call := GetAudioIsrcBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
