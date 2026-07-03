package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAccountCreationRequestAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCreationRequestAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCreationRequestAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdAccountCreationRequestAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
	return &out, err
}

type GetAdAccountCreationRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCreationRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCreationRequest(ctx context.Context, client *core.Client, id string, params GetAdAccountCreationRequestParams) (*objects.AdAccountCreationRequest, error) {
	var out objects.AdAccountCreationRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
