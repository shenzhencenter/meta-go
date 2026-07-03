package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoListVideosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoListVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoListVideos(ctx context.Context, client *core.Client, id string, params GetVideoListVideosParams) (*core.Cursor[objects.AdVideo], error) {
	var out core.Cursor[objects.AdVideo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), &out)
	return &out, err
}

type GetVideoListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoList(ctx context.Context, client *core.Client, id string, params GetVideoListParams) (*objects.VideoList, error) {
	var out objects.VideoList
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
