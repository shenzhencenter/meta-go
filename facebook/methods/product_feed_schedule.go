package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetProductFeedScheduleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedScheduleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedSchedule(ctx context.Context, client *core.Client, id string, params GetProductFeedScheduleParams) (*objects.ProductFeedSchedule, error) {
	var out objects.ProductFeedSchedule
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
