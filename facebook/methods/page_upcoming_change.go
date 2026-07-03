package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetPageUpcomingChangeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageUpcomingChangeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageUpcomingChange(ctx context.Context, client *core.Client, id string, params GetPageUpcomingChangeParams) (*objects.PageUpcomingChange, error) {
	var out objects.PageUpcomingChange
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
