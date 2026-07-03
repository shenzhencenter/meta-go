package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessTagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessTagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessTag(ctx context.Context, client *core.Client, id string, params GetBusinessTagParams) (*objects.BusinessTag, error) {
	var out objects.BusinessTag
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
