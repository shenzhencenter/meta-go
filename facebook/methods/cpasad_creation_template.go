package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASAdCreationTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASAdCreationTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASAdCreationTemplate(ctx context.Context, client *core.Client, id string, params GetCPASAdCreationTemplateParams) (*objects.CPASAdCreationTemplate, error) {
	var out objects.CPASAdCreationTemplate
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
