package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOffsitePixelParams struct {
	Value *uint64     `facebook:"value"`
	Extra core.Params `facebook:"-"`
}

func (params GetOffsitePixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Value != nil {
		out["value"] = *params.Value
	}
	return out
}

func GetOffsitePixel(ctx context.Context, client *core.Client, id string, params GetOffsitePixelParams) (*objects.OffsitePixel, error) {
	var out objects.OffsitePixel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
