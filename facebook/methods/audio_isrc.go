package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAudioIsrc(ctx context.Context, client *core.Client, id string, params GetAudioIsrcParams) (*objects.AudioIsrc, error) {
	var out objects.AudioIsrc
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
