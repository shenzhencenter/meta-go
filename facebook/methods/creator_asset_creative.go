package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCreatorAssetCreativeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCreatorAssetCreativeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCreatorAssetCreative(ctx context.Context, client *core.Client, id string, params GetCreatorAssetCreativeParams) (*objects.CreatorAssetCreative, error) {
	var out objects.CreatorAssetCreative
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
