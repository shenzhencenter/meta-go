package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeletePageUserMessageThreadLabelLabelParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeletePageUserMessageThreadLabelLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeletePageUserMessageThreadLabelLabelBatchCall(id string, params DeletePageUserMessageThreadLabelLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "label"), params.ToParams(), options...)
}

func NewDeletePageUserMessageThreadLabelLabelBatchRequest(id string, params DeletePageUserMessageThreadLabelLabelParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageUserMessageThreadLabelLabelBatchCall(id, params, options...))
}

func DecodeDeletePageUserMessageThreadLabelLabelBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageUserMessageThreadLabelLabelWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelLabelParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageUserMessageThreadLabelLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageUserMessageThreadLabelLabel(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelLabelParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageUserMessageThreadLabelLabelWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageUserMessageThreadLabelLabelParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params CreatePageUserMessageThreadLabelLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func CreatePageUserMessageThreadLabelLabelBatchCall(id string, params CreatePageUserMessageThreadLabelLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "label"), params.ToParams(), options...)
}

func NewCreatePageUserMessageThreadLabelLabelBatchRequest(id string, params CreatePageUserMessageThreadLabelLabelParams, options ...core.BatchOption) *core.BatchRequest[objects.PageUserMessageThreadLabel] {
	return core.NewBatchRequest[objects.PageUserMessageThreadLabel](CreatePageUserMessageThreadLabelLabelBatchCall(id, params, options...))
}

func DecodeCreatePageUserMessageThreadLabelLabelBatchResponse(response *core.BatchResponse) (*objects.PageUserMessageThreadLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageUserMessageThreadLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageUserMessageThreadLabelLabelWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageUserMessageThreadLabelLabelParams) (*objects.PageUserMessageThreadLabel, *core.Response, error) {
	var out objects.PageUserMessageThreadLabel
	call := CreatePageUserMessageThreadLabelLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageUserMessageThreadLabelLabel(ctx context.Context, client *core.Client, id string, params CreatePageUserMessageThreadLabelLabelParams) (*objects.PageUserMessageThreadLabel, error) {
	out, _, err := CreatePageUserMessageThreadLabelLabelWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageUserMessageThreadLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePageUserMessageThreadLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePageUserMessageThreadLabelBatchCall(id string, params DeletePageUserMessageThreadLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePageUserMessageThreadLabelBatchRequest(id string, params DeletePageUserMessageThreadLabelParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageUserMessageThreadLabelBatchCall(id, params, options...))
}

func DecodeDeletePageUserMessageThreadLabelBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageUserMessageThreadLabelWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageUserMessageThreadLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageUserMessageThreadLabel(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageUserMessageThreadLabelWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageUserMessageThreadLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageUserMessageThreadLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageUserMessageThreadLabelBatchCall(id string, params GetPageUserMessageThreadLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPageUserMessageThreadLabelBatchRequest(id string, params GetPageUserMessageThreadLabelParams, options ...core.BatchOption) *core.BatchRequest[objects.PageUserMessageThreadLabel] {
	return core.NewBatchRequest[objects.PageUserMessageThreadLabel](GetPageUserMessageThreadLabelBatchCall(id, params, options...))
}

func DecodeGetPageUserMessageThreadLabelBatchResponse(response *core.BatchResponse) (*objects.PageUserMessageThreadLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageUserMessageThreadLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageUserMessageThreadLabelWithResponse(ctx context.Context, client *core.Client, id string, params GetPageUserMessageThreadLabelParams) (*objects.PageUserMessageThreadLabel, *core.Response, error) {
	var out objects.PageUserMessageThreadLabel
	call := GetPageUserMessageThreadLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageUserMessageThreadLabel(ctx context.Context, client *core.Client, id string, params GetPageUserMessageThreadLabelParams) (*objects.PageUserMessageThreadLabel, error) {
	out, _, err := GetPageUserMessageThreadLabelWithResponse(ctx, client, id, params)
	return out, err
}
