package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventExternalTicketInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventExternalTicketInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventExternalTicketInfo(ctx context.Context, client *core.Client, id string, params GetEventExternalTicketInfoParams) (*objects.EventExternalTicketInfo, error) {
	var out objects.EventExternalTicketInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
