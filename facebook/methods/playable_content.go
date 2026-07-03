package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetPlayableContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlayableContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlayableContent(ctx context.Context, client *core.Client, id string, params GetPlayableContentParams) (*objects.PlayableContent, error) {
	var out objects.PlayableContent
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
