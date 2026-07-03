package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCopyrightMediaMisuseParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCopyrightMediaMisuseParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCopyrightMediaMisuseBatchCall(id string, params GetCopyrightMediaMisuseParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCopyrightMediaMisuseBatchRequest(id string, params GetCopyrightMediaMisuseParams, options ...core.BatchOption) *core.BatchRequest[objects.CopyrightMediaMisuse] {
	return core.NewBatchRequest[objects.CopyrightMediaMisuse](GetCopyrightMediaMisuseBatchCall(id, params, options...))
}

func DecodeGetCopyrightMediaMisuseBatchResponse(response *core.BatchResponse) (*objects.CopyrightMediaMisuse, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CopyrightMediaMisuse
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCopyrightMediaMisuseWithResponse(ctx context.Context, client *core.Client, id string, params GetCopyrightMediaMisuseParams) (*objects.CopyrightMediaMisuse, *core.Response, error) {
	var out objects.CopyrightMediaMisuse
	call := GetCopyrightMediaMisuseBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCopyrightMediaMisuse(ctx context.Context, client *core.Client, id string, params GetCopyrightMediaMisuseParams) (*objects.CopyrightMediaMisuse, error) {
	out, _, err := GetCopyrightMediaMisuseWithResponse(ctx, client, id, params)
	return out, err
}
