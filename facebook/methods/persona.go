package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func DeletePersonaBatchCall(id string, params DeletePersonaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePersonaBatchRequest(id string, params DeletePersonaParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePersonaBatchCall(id, params, options...))
}

func DecodeDeletePersonaBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeletePersonaWithResponse(ctx context.Context, client *core.Client, id string, params DeletePersonaParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePersonaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePersona(ctx context.Context, client *core.Client, id string, params DeletePersonaParams) (*map[string]interface{}, error) {
	out, _, err := DeletePersonaWithResponse(ctx, client, id, params)
	return out, err
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

func GetPersonaBatchCall(id string, params GetPersonaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPersonaBatchRequest(id string, params GetPersonaParams, options ...core.BatchOption) *core.BatchRequest[objects.Persona] {
	return core.NewBatchRequest[objects.Persona](GetPersonaBatchCall(id, params, options...))
}

func DecodeGetPersonaBatchResponse(response *core.BatchResponse) (*objects.Persona, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Persona
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPersonaWithResponse(ctx context.Context, client *core.Client, id string, params GetPersonaParams) (*objects.Persona, *core.Response, error) {
	var out objects.Persona
	call := GetPersonaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPersona(ctx context.Context, client *core.Client, id string, params GetPersonaParams) (*objects.Persona, error) {
	out, _, err := GetPersonaWithResponse(ctx, client, id, params)
	return out, err
}
