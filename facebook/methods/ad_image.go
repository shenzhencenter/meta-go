package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdImage(ctx context.Context, client *core.Client, id string, params GetAdImageParams) (*objects.AdImage, error) {
	var out objects.AdImage
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
