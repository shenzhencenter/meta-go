package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventTicketSettingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventTicketSettingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventTicketSetting(ctx context.Context, client *core.Client, id string, params GetEventTicketSettingParams) (*objects.EventTicketSetting, error) {
	var out objects.EventTicketSetting
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
