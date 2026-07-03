package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPlaceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlaceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlace(ctx context.Context, client *core.Client, id string, params GetPlaceParams) (*objects.Place, error) {
	var out objects.Place
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
