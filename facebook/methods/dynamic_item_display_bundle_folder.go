package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetDynamicItemDisplayBundleFolderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicItemDisplayBundleFolderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicItemDisplayBundleFolder(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleFolderParams) (*objects.DynamicItemDisplayBundleFolder, error) {
	var out objects.DynamicItemDisplayBundleFolder
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
