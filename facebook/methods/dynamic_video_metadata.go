package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicVideoMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicVideoMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicVideoMetadata(ctx context.Context, client *core.Client, id string, params GetDynamicVideoMetadataParams) (*objects.DynamicVideoMetadata, error) {
	var out objects.DynamicVideoMetadata
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
