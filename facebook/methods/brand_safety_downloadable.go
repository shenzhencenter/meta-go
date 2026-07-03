package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBrandSafetyDownloadableParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBrandSafetyDownloadableParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBrandSafetyDownloadable(ctx context.Context, client *core.Client, id string, params GetBrandSafetyDownloadableParams) (*objects.BrandSafetyDownloadable, error) {
	var out objects.BrandSafetyDownloadable
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
