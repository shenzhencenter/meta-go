package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteLeadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteLeadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteLeadBatchCall(id string, params DeleteLeadParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteLeadBatchRequest(id string, params DeleteLeadParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteLeadBatchCall(id, params, options...))
}

func DecodeDeleteLeadBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteLead(ctx context.Context, client *core.Client, id string, params DeleteLeadParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteLeadBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLeadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadBatchCall(id string, params GetLeadParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLeadBatchRequest(id string, params GetLeadParams, options ...core.BatchOption) *core.BatchRequest[objects.Lead] {
	return core.NewBatchRequest[objects.Lead](GetLeadBatchCall(id, params, options...))
}

func DecodeGetLeadBatchResponse(response *core.BatchResponse) (*objects.Lead, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Lead
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLead(ctx context.Context, client *core.Client, id string, params GetLeadParams) (*objects.Lead, error) {
	var out objects.Lead
	call := GetLeadBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
