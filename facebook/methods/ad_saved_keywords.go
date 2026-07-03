package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdSavedKeywordsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedKeywordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedKeywords(ctx context.Context, client *core.Client, id string, params GetAdSavedKeywordsParams) (*objects.AdSavedKeywords, error) {
	var out objects.AdSavedKeywords
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
