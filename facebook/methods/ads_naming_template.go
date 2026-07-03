package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdsNamingTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsNamingTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsNamingTemplate(ctx context.Context, client *core.Client, id string, params GetAdsNamingTemplateParams) (*objects.AdsNamingTemplate, error) {
	var out objects.AdsNamingTemplate
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
