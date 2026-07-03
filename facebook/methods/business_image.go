package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessImage(ctx context.Context, client *core.Client, id string, params GetBusinessImageParams) (*objects.BusinessImage, error) {
	var out objects.BusinessImage
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
