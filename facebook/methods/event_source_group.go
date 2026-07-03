package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventSourceGroupSharedAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventSourceGroupSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventSourceGroupSharedAccountsBatchCall(id string, params GetEventSourceGroupSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewGetEventSourceGroupSharedAccountsBatchRequest(id string, params GetEventSourceGroupSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetEventSourceGroupSharedAccountsBatchCall(id, params, options...))
}

func DecodeGetEventSourceGroupSharedAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventSourceGroupSharedAccounts(ctx context.Context, client *core.Client, id string, params GetEventSourceGroupSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetEventSourceGroupSharedAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateEventSourceGroupSharedAccountsParams struct {
	Accounts []string    `facebook:"accounts"`
	Extra    core.Params `facebook:"-"`
}

func (params CreateEventSourceGroupSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["accounts"] = params.Accounts
	return out
}

func CreateEventSourceGroupSharedAccountsBatchCall(id string, params CreateEventSourceGroupSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewCreateEventSourceGroupSharedAccountsBatchRequest(id string, params CreateEventSourceGroupSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.EventSourceGroup] {
	return core.NewBatchRequest[objects.EventSourceGroup](CreateEventSourceGroupSharedAccountsBatchCall(id, params, options...))
}

func DecodeCreateEventSourceGroupSharedAccountsBatchResponse(response *core.BatchResponse) (*objects.EventSourceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventSourceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateEventSourceGroupSharedAccounts(ctx context.Context, client *core.Client, id string, params CreateEventSourceGroupSharedAccountsParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	call := CreateEventSourceGroupSharedAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetEventSourceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventSourceGroupBatchCall(id string, params GetEventSourceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventSourceGroupBatchRequest(id string, params GetEventSourceGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.EventSourceGroup] {
	return core.NewBatchRequest[objects.EventSourceGroup](GetEventSourceGroupBatchCall(id, params, options...))
}

func DecodeGetEventSourceGroupBatchResponse(response *core.BatchResponse) (*objects.EventSourceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventSourceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventSourceGroup(ctx context.Context, client *core.Client, id string, params GetEventSourceGroupParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	call := GetEventSourceGroupBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateEventSourceGroupParams struct {
	EventSources []string    `facebook:"event_sources"`
	Name         string      `facebook:"name"`
	Extra        core.Params `facebook:"-"`
}

func (params UpdateEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["event_sources"] = params.EventSources
	out["name"] = params.Name
	return out
}

func UpdateEventSourceGroupBatchCall(id string, params UpdateEventSourceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateEventSourceGroupBatchRequest(id string, params UpdateEventSourceGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.EventSourceGroup] {
	return core.NewBatchRequest[objects.EventSourceGroup](UpdateEventSourceGroupBatchCall(id, params, options...))
}

func DecodeUpdateEventSourceGroupBatchResponse(response *core.BatchResponse) (*objects.EventSourceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventSourceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateEventSourceGroup(ctx context.Context, client *core.Client, id string, params UpdateEventSourceGroupParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	call := UpdateEventSourceGroupBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
