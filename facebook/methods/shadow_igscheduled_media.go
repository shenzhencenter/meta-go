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

func GetShadowIGScheduledMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGScheduledMediaParams) (*objects.ShadowIGScheduledMedia, error) {
	var out objects.ShadowIGScheduledMedia
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
