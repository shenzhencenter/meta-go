package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGAccessTokenForIGOnlyAPIParams struct {
	AccessToken  string      `facebook:"access_token"`
	ClientSecret string      `facebook:"client_secret"`
	GrantType    string      `facebook:"grant_type"`
	Extra        core.Params `facebook:"-"`
}

func (params GetIGAccessTokenForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["access_token"] = params.AccessToken
	out["client_secret"] = params.ClientSecret
	out["grant_type"] = params.GrantType
	return out
}

func GetIGAccessTokenForIGOnlyAPIBatchCall(id string, params GetIGAccessTokenForIGOnlyAPIParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGAccessTokenForIGOnlyAPIBatchRequest(id string, params GetIGAccessTokenForIGOnlyAPIParams, options ...core.BatchOption) *core.BatchRequest[objects.AccessToken] {
	return core.NewBatchRequest[objects.AccessToken](GetIGAccessTokenForIGOnlyAPIBatchCall(id, params, options...))
}

func DecodeGetIGAccessTokenForIGOnlyAPIBatchResponse(response *core.BatchResponse) (*objects.AccessToken, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AccessToken
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGAccessTokenForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGAccessTokenForIGOnlyAPIParams) (*objects.AccessToken, error) {
	var out objects.AccessToken
	call := GetIGAccessTokenForIGOnlyAPIBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
