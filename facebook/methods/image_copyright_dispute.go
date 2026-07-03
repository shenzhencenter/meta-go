package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetImageCopyrightDisputeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetImageCopyrightDisputeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetImageCopyrightDispute(ctx context.Context, client *core.Client, id string, params GetImageCopyrightDisputeParams) (*objects.ImageCopyrightDispute, error) {
	var out objects.ImageCopyrightDispute
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
