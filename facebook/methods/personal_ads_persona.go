package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPersonalAdsPersonaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPersonalAdsPersonaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPersonalAdsPersonaBatchCall(id string, params GetPersonalAdsPersonaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPersonalAdsPersonaBatchRequest(id string, params GetPersonalAdsPersonaParams, options ...core.BatchOption) *core.BatchRequest[objects.PersonalAdsPersona] {
	return core.NewBatchRequest[objects.PersonalAdsPersona](GetPersonalAdsPersonaBatchCall(id, params, options...))
}

func DecodeGetPersonalAdsPersonaBatchResponse(response *core.BatchResponse) (*objects.PersonalAdsPersona, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PersonalAdsPersona
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPersonalAdsPersona(ctx context.Context, client *core.Client, id string, params GetPersonalAdsPersonaParams) (*objects.PersonalAdsPersona, error) {
	var out objects.PersonalAdsPersona
	call := GetPersonalAdsPersonaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
