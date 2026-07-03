package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type CreateAdAccountAccountControlsAccountControlsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountAccountControlsAccountControlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdAccountAccountControlsAccountControls(ctx context.Context, client *core.Client, id string, params CreateAdAccountAccountControlsAccountControlsParams) (*objects.AdAccountAccountControlsPost, error) {
	var out objects.AdAccountAccountControlsPost
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "account_controls"), params.ToParams(), &out)
	return &out, err
}
