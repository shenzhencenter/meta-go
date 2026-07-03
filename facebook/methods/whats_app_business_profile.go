package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWhatsAppBusinessProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessProfile(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessProfileParams) (*objects.WhatsAppBusinessProfile, error) {
	var out objects.WhatsAppBusinessProfile
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateWhatsAppBusinessProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params UpdateWhatsAppBusinessProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func UpdateWhatsAppBusinessProfile(ctx context.Context, client *core.Client, id string, params UpdateWhatsAppBusinessProfileParams) (*objects.WhatsAppBusinessProfile, error) {
	var out objects.WhatsAppBusinessProfile
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
