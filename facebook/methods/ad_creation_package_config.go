package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdCreationPackageConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdCreationPackageConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdCreationPackageConfig(ctx context.Context, client *core.Client, id string, params GetAdCreationPackageConfigParams) (*objects.AdCreationPackageConfig, error) {
	var out objects.AdCreationPackageConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
