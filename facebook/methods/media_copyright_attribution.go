package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMediaCopyrightAttributionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMediaCopyrightAttributionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMediaCopyrightAttribution(ctx context.Context, client *core.Client, id string, params GetMediaCopyrightAttributionParams) (*objects.MediaCopyrightAttribution, error) {
	var out objects.MediaCopyrightAttribution
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
