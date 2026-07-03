package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCanvasTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasTemplate(ctx context.Context, client *core.Client, id string, params GetCanvasTemplateParams) (*objects.CanvasTemplate, error) {
	var out objects.CanvasTemplate
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
