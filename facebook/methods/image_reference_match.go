package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetImageReferenceMatchParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetImageReferenceMatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetImageReferenceMatch(ctx context.Context, client *core.Client, id string, params GetImageReferenceMatchParams) (*objects.ImageReferenceMatch, error) {
	var out objects.ImageReferenceMatch
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
