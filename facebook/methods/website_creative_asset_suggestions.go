package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeAssetSuggestionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeAssetSuggestionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeAssetSuggestions(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSuggestionsParams) (*objects.WebsiteCreativeAssetSuggestions, error) {
	var out objects.WebsiteCreativeAssetSuggestions
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
