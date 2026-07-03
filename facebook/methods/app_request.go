package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAppRequestParams struct {
	Ids   []core.ID   `facebook:"ids"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteAppRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ids"] = params.Ids
	return out
}

func DeleteAppRequest(ctx context.Context, client *core.Client, id string, params DeleteAppRequestParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAppRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAppRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAppRequest(ctx context.Context, client *core.Client, id string, params GetAppRequestParams) (*objects.AppRequest, error) {
	var out objects.AppRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
