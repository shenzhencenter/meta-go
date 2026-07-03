package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeAssetSourceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeAssetSourceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeAssetSource(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSourceParams) (*objects.WebsiteCreativeAssetSource, error) {
	var out objects.WebsiteCreativeAssetSource
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
