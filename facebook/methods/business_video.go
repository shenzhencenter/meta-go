package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBusinessVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessVideo(ctx context.Context, client *core.Client, id string, params GetBusinessVideoParams) (*objects.BusinessVideo, error) {
	var out objects.BusinessVideo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
