package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMusicVideoCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMusicVideoCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMusicVideoCopyrightBatchCall(id string, params GetMusicVideoCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMusicVideoCopyrightBatchRequest(id string, params GetMusicVideoCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.MusicVideoCopyright] {
	return core.NewBatchRequest[objects.MusicVideoCopyright](GetMusicVideoCopyrightBatchCall(id, params, options...))
}

func DecodeGetMusicVideoCopyrightBatchResponse(response *core.BatchResponse) (*objects.MusicVideoCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MusicVideoCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMusicVideoCopyrightWithResponse(ctx context.Context, client *core.Client, id string, params GetMusicVideoCopyrightParams) (*objects.MusicVideoCopyright, *core.Response, error) {
	var out objects.MusicVideoCopyright
	call := GetMusicVideoCopyrightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetMusicVideoCopyright(ctx context.Context, client *core.Client, id string, params GetMusicVideoCopyrightParams) (*objects.MusicVideoCopyright, error) {
	out, _, err := GetMusicVideoCopyrightWithResponse(ctx, client, id, params)
	return out, err
}
