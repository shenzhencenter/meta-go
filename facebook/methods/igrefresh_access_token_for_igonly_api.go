package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGRefreshAccessTokenForIGOnlyAPIParams struct {
	AccessToken string      `facebook:"access_token"`
	GrantType   string      `facebook:"grant_type"`
	Extra       core.Params `facebook:"-"`
}

func (params GetIGRefreshAccessTokenForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["access_token"] = params.AccessToken
	out["grant_type"] = params.GrantType
	return out
}

func GetIGRefreshAccessTokenForIGOnlyAPIBatchCall(id string, params GetIGRefreshAccessTokenForIGOnlyAPIParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGRefreshAccessTokenForIGOnlyAPIBatchRequest(id string, params GetIGRefreshAccessTokenForIGOnlyAPIParams, options ...core.BatchOption) *core.BatchRequest[objects.RefreshAccessToken] {
	return core.NewBatchRequest[objects.RefreshAccessToken](GetIGRefreshAccessTokenForIGOnlyAPIBatchCall(id, params, options...))
}

func DecodeGetIGRefreshAccessTokenForIGOnlyAPIBatchResponse(response *core.BatchResponse) (*objects.RefreshAccessToken, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.RefreshAccessToken
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGRefreshAccessTokenForIGOnlyAPIWithResponse(ctx context.Context, client *core.Client, id string, params GetIGRefreshAccessTokenForIGOnlyAPIParams) (*objects.RefreshAccessToken, *core.Response, error) {
	var out objects.RefreshAccessToken
	call := GetIGRefreshAccessTokenForIGOnlyAPIBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGRefreshAccessTokenForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGRefreshAccessTokenForIGOnlyAPIParams) (*objects.RefreshAccessToken, error) {
	out, _, err := GetIGRefreshAccessTokenForIGOnlyAPIWithResponse(ctx, client, id, params)
	return out, err
}
