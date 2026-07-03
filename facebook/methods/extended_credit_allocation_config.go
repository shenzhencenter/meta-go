package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteExtendedCreditAllocationConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteExtendedCreditAllocationConfigBatchCall(id string, params DeleteExtendedCreditAllocationConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteExtendedCreditAllocationConfigBatchRequest(id string, params DeleteExtendedCreditAllocationConfigParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteExtendedCreditAllocationConfigBatchCall(id, params, options...))
}

func DecodeDeleteExtendedCreditAllocationConfigBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteExtendedCreditAllocationConfigWithResponse(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditAllocationConfigParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteExtendedCreditAllocationConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditAllocationConfigParams) (*map[string]interface{}, error) {
	out, _, err := DeleteExtendedCreditAllocationConfigWithResponse(ctx, client, id, params)
	return out, err
}

type GetExtendedCreditAllocationConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditAllocationConfigBatchCall(id string, params GetExtendedCreditAllocationConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetExtendedCreditAllocationConfigBatchRequest(id string, params GetExtendedCreditAllocationConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditAllocationConfig] {
	return core.NewBatchRequest[objects.ExtendedCreditAllocationConfig](GetExtendedCreditAllocationConfigBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditAllocationConfigBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditAllocationConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditAllocationConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditAllocationConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, *core.Response, error) {
	var out objects.ExtendedCreditAllocationConfig
	call := GetExtendedCreditAllocationConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params GetExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, error) {
	out, _, err := GetExtendedCreditAllocationConfigWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateExtendedCreditAllocationConfigParams struct {
	Amount *map[string]interface{} `facebook:"amount"`
	Extra  core.Params             `facebook:"-"`
}

func (params UpdateExtendedCreditAllocationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Amount != nil {
		out["amount"] = *params.Amount
	}
	return out
}

func UpdateExtendedCreditAllocationConfigBatchCall(id string, params UpdateExtendedCreditAllocationConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateExtendedCreditAllocationConfigBatchRequest(id string, params UpdateExtendedCreditAllocationConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditAllocationConfig] {
	return core.NewBatchRequest[objects.ExtendedCreditAllocationConfig](UpdateExtendedCreditAllocationConfigBatchCall(id, params, options...))
}

func DecodeUpdateExtendedCreditAllocationConfigBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditAllocationConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditAllocationConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateExtendedCreditAllocationConfigWithResponse(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, *core.Response, error) {
	var out objects.ExtendedCreditAllocationConfig
	call := UpdateExtendedCreditAllocationConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateExtendedCreditAllocationConfig(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditAllocationConfigParams) (*objects.ExtendedCreditAllocationConfig, error) {
	out, _, err := UpdateExtendedCreditAllocationConfigWithResponse(ctx, client, id, params)
	return out, err
}
