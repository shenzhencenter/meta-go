package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicARMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicARMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicARMetadata(ctx context.Context, client *core.Client, id string, params GetDynamicARMetadataParams) (*objects.DynamicARMetadata, error) {
	var out objects.DynamicARMetadata
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
