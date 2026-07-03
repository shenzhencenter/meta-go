package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWhitehatFBDLRunParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhitehatFBDLRunParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhitehatFBDLRunBatchCall(id string, params GetWhitehatFBDLRunParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWhitehatFBDLRunBatchRequest(id string, params GetWhitehatFBDLRunParams, options ...core.BatchOption) *core.BatchRequest[objects.WhitehatFBDLRun] {
	return core.NewBatchRequest[objects.WhitehatFBDLRun](GetWhitehatFBDLRunBatchCall(id, params, options...))
}

func DecodeGetWhitehatFBDLRunBatchResponse(response *core.BatchResponse) (*objects.WhitehatFBDLRun, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhitehatFBDLRun
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhitehatFBDLRunWithResponse(ctx context.Context, client *core.Client, id string, params GetWhitehatFBDLRunParams) (*objects.WhitehatFBDLRun, *core.Response, error) {
	var out objects.WhitehatFBDLRun
	call := GetWhitehatFBDLRunBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWhitehatFBDLRun(ctx context.Context, client *core.Client, id string, params GetWhitehatFBDLRunParams) (*objects.WhitehatFBDLRun, error) {
	out, _, err := GetWhitehatFBDLRunWithResponse(ctx, client, id, params)
	return out, err
}
