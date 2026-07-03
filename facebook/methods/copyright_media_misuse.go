package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCopyrightMediaMisuseParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCopyrightMediaMisuseParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCopyrightMediaMisuse(ctx context.Context, client *core.Client, id string, params GetCopyrightMediaMisuseParams) (*objects.CopyrightMediaMisuse, error) {
	var out objects.CopyrightMediaMisuse
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
