package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdAsyncRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAsyncRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdAsyncRequest(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAdAsyncRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAsyncRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAsyncRequest(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestParams) (*objects.AdAsyncRequest, error) {
	var out objects.AdAsyncRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
