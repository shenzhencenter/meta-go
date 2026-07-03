package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetPlaceTagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlaceTagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlaceTag(ctx context.Context, client *core.Client, id string, params GetPlaceTagParams) (*objects.PlaceTag, error) {
	var out objects.PlaceTag
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
