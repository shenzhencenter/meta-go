package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBidScheduleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBidScheduleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBidSchedule(ctx context.Context, client *core.Client, id string, params GetBidScheduleParams) (*objects.BidSchedule, error) {
	var out objects.BidSchedule
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
