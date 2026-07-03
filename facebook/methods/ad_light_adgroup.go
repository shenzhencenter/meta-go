package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLightAdgroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightAdgroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightAdgroupBatchCall(id string, params GetAdLightAdgroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdLightAdgroupBatchRequest(id string, params GetAdLightAdgroupParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLightAdgroup] {
	return core.NewBatchRequest[objects.AdLightAdgroup](GetAdLightAdgroupBatchCall(id, params, options...))
}

func DecodeGetAdLightAdgroupBatchResponse(response *core.BatchResponse) (*objects.AdLightAdgroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLightAdgroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLightAdgroupWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLightAdgroupParams) (*objects.AdLightAdgroup, *core.Response, error) {
	var out objects.AdLightAdgroup
	call := GetAdLightAdgroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLightAdgroup(ctx context.Context, client *core.Client, id string, params GetAdLightAdgroupParams) (*objects.AdLightAdgroup, error) {
	out, _, err := GetAdLightAdgroupWithResponse(ctx, client, id, params)
	return out, err
}
