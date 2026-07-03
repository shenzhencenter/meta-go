package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdPlacePageSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdPlacePageSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdPlacePageSet(ctx context.Context, client *core.Client, id string, params GetAdPlacePageSetParams) (*objects.AdPlacePageSet, error) {
	var out objects.AdPlacePageSet
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
