package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFBImageCopyrightMatchParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFBImageCopyrightMatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFBImageCopyrightMatch(ctx context.Context, client *core.Client, id string, params GetFBImageCopyrightMatchParams) (*objects.FBImageCopyrightMatch, error) {
	var out objects.FBImageCopyrightMatch
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
