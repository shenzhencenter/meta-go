package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetMusicVideoCopyright(ctx context.Context, client *core.Client, id string, params GetMusicVideoCopyrightParams) (*objects.MusicVideoCopyright, error) {
	var out objects.MusicVideoCopyright
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
