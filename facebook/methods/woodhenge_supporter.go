package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWoodhengeSupporterParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWoodhengeSupporterParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWoodhengeSupporter(ctx context.Context, client *core.Client, id string, params GetWoodhengeSupporterParams) (*objects.WoodhengeSupporter, error) {
	var out objects.WoodhengeSupporter
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
