package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAsyncRequestSetRequestsParams struct {
	Statuses *[]enums.AdasyncrequestsetrequestsStatusesEnumParam `facebook:"statuses"`
	Extra    core.Params                                         `facebook:"-"`
}

func (params GetAdAsyncRequestSetRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Statuses != nil {
		out["statuses"] = *params.Statuses
	}
	return out
}

func GetAdAsyncRequestSetRequestsBatchCall(id string, params GetAdAsyncRequestSetRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "requests"), params.ToParams(), options...)
}

func NewGetAdAsyncRequestSetRequestsBatchRequest(id string, params GetAdAsyncRequestSetRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAsyncRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAsyncRequest]](GetAdAsyncRequestSetRequestsBatchCall(id, params, options...))
}

func DecodeGetAdAsyncRequestSetRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAsyncRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAsyncRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAsyncRequestSetRequestsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetRequestsParams) (*core.Cursor[objects.AdAsyncRequest], *core.Response, error) {
	var out core.Cursor[objects.AdAsyncRequest]
	call := GetAdAsyncRequestSetRequestsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAsyncRequestSetRequests(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetRequestsParams) (*core.Cursor[objects.AdAsyncRequest], error) {
	out, _, err := GetAdAsyncRequestSetRequestsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdAsyncRequestSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdAsyncRequestSetBatchCall(id string, params DeleteAdAsyncRequestSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdAsyncRequestSetBatchRequest(id string, params DeleteAdAsyncRequestSetParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAsyncRequestSetBatchCall(id, params, options...))
}

func DecodeDeleteAdAsyncRequestSetBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAsyncRequestSetWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestSetParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdAsyncRequestSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestSetParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdAsyncRequestSetWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdAsyncRequestSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAsyncRequestSetBatchCall(id string, params GetAdAsyncRequestSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAsyncRequestSetBatchRequest(id string, params GetAdAsyncRequestSetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAsyncRequestSet] {
	return core.NewBatchRequest[objects.AdAsyncRequestSet](GetAdAsyncRequestSetBatchCall(id, params, options...))
}

func DecodeGetAdAsyncRequestSetBatchResponse(response *core.BatchResponse) (*objects.AdAsyncRequestSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAsyncRequestSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAsyncRequestSetWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, *core.Response, error) {
	var out objects.AdAsyncRequestSet
	call := GetAdAsyncRequestSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, error) {
	out, _, err := GetAdAsyncRequestSetWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdAsyncRequestSetParams struct {
	Name             *string                                  `facebook:"name"`
	NotificationMode *enums.AdasyncrequestsetNotificationMode `facebook:"notification_mode"`
	NotificationURI  *string                                  `facebook:"notification_uri"`
	Extra            core.Params                              `facebook:"-"`
}

func (params UpdateAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NotificationMode != nil {
		out["notification_mode"] = *params.NotificationMode
	}
	if params.NotificationURI != nil {
		out["notification_uri"] = *params.NotificationURI
	}
	return out
}

func UpdateAdAsyncRequestSetBatchCall(id string, params UpdateAdAsyncRequestSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdAsyncRequestSetBatchRequest(id string, params UpdateAdAsyncRequestSetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAsyncRequestSet] {
	return core.NewBatchRequest[objects.AdAsyncRequestSet](UpdateAdAsyncRequestSetBatchCall(id, params, options...))
}

func DecodeUpdateAdAsyncRequestSetBatchResponse(response *core.BatchResponse) (*objects.AdAsyncRequestSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAsyncRequestSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdAsyncRequestSetWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, *core.Response, error) {
	var out objects.AdAsyncRequestSet
	call := UpdateAdAsyncRequestSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params UpdateAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, error) {
	out, _, err := UpdateAdAsyncRequestSetWithResponse(ctx, client, id, params)
	return out, err
}
