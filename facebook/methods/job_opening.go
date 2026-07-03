package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetJobOpeningParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetJobOpeningParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetJobOpening(ctx context.Context, client *core.Client, id string, params GetJobOpeningParams) (*objects.JobOpening, error) {
	var out objects.JobOpening
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
