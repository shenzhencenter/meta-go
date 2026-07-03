package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMusicWorkCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMusicWorkCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMusicWorkCopyrightBatchCall(id string, params GetMusicWorkCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMusicWorkCopyrightBatchRequest(id string, params GetMusicWorkCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.MusicWorkCopyright] {
	return core.NewBatchRequest[objects.MusicWorkCopyright](GetMusicWorkCopyrightBatchCall(id, params, options...))
}

func DecodeGetMusicWorkCopyrightBatchResponse(response *core.BatchResponse) (*objects.MusicWorkCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MusicWorkCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMusicWorkCopyright(ctx context.Context, client *core.Client, id string, params GetMusicWorkCopyrightParams) (*objects.MusicWorkCopyright, error) {
	var out objects.MusicWorkCopyright
	call := GetMusicWorkCopyrightBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
