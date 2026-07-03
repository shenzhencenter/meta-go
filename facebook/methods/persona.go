package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeletePersonaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePersonaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePersona(ctx context.Context, client *core.Client, id string, params DeletePersonaParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPersonaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPersonaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPersona(ctx context.Context, client *core.Client, id string, params GetPersonaParams) (*objects.Persona, error) {
	var out objects.Persona
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
