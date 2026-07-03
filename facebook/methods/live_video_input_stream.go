package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoInputStreamParams struct {
	TargetToken *string     `facebook:"target_token"`
	Extra       core.Params `facebook:"-"`
}

func (params GetLiveVideoInputStreamParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetToken != nil {
		out["target_token"] = *params.TargetToken
	}
	return out
}

func GetLiveVideoInputStream(ctx context.Context, client *core.Client, id string, params GetLiveVideoInputStreamParams) (*objects.LiveVideoInputStream, error) {
	var out objects.LiveVideoInputStream
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
