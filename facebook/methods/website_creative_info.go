package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeInfo(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeInfoParams) (*objects.WebsiteCreativeInfo, error) {
	var out objects.WebsiteCreativeInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
