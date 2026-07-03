package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdSavedLocationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedLocationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedLocation(ctx context.Context, client *core.Client, id string, params GetAdSavedLocationParams) (*objects.AdSavedLocation, error) {
	var out objects.AdSavedLocation
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
