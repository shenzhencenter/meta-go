package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeleteLeadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteLeadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteLead(ctx context.Context, client *core.Client, id string, params DeleteLeadParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetLeadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLead(ctx context.Context, client *core.Client, id string, params GetLeadParams) (*objects.Lead, error) {
	var out objects.Lead
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
