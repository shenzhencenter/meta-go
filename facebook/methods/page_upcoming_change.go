package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageUpcomingChangeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageUpcomingChangeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageUpcomingChangeBatchCall(id string, params GetPageUpcomingChangeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPageUpcomingChangeBatchRequest(id string, params GetPageUpcomingChangeParams, options ...core.BatchOption) *core.BatchRequest[objects.PageUpcomingChange] {
	return core.NewBatchRequest[objects.PageUpcomingChange](GetPageUpcomingChangeBatchCall(id, params, options...))
}

func DecodeGetPageUpcomingChangeBatchResponse(response *core.BatchResponse) (*objects.PageUpcomingChange, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageUpcomingChange
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageUpcomingChangeWithResponse(ctx context.Context, client *core.Client, id string, params GetPageUpcomingChangeParams) (*objects.PageUpcomingChange, *core.Response, error) {
	var out objects.PageUpcomingChange
	call := GetPageUpcomingChangeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageUpcomingChange(ctx context.Context, client *core.Client, id string, params GetPageUpcomingChangeParams) (*objects.PageUpcomingChange, error) {
	out, _, err := GetPageUpcomingChangeWithResponse(ctx, client, id, params)
	return out, err
}
