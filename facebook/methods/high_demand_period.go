package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteHighDemandPeriodParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteHighDemandPeriodParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteHighDemandPeriodBatchCall(id string, params DeleteHighDemandPeriodParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteHighDemandPeriodBatchRequest(id string, params DeleteHighDemandPeriodParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteHighDemandPeriodBatchCall(id, params, options...))
}

func DecodeDeleteHighDemandPeriodBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteHighDemandPeriod(ctx context.Context, client *core.Client, id string, params DeleteHighDemandPeriodParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteHighDemandPeriodBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetHighDemandPeriodParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHighDemandPeriodParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHighDemandPeriodBatchCall(id string, params GetHighDemandPeriodParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetHighDemandPeriodBatchRequest(id string, params GetHighDemandPeriodParams, options ...core.BatchOption) *core.BatchRequest[objects.HighDemandPeriod] {
	return core.NewBatchRequest[objects.HighDemandPeriod](GetHighDemandPeriodBatchCall(id, params, options...))
}

func DecodeGetHighDemandPeriodBatchResponse(response *core.BatchResponse) (*objects.HighDemandPeriod, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.HighDemandPeriod
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHighDemandPeriod(ctx context.Context, client *core.Client, id string, params GetHighDemandPeriodParams) (*objects.HighDemandPeriod, error) {
	var out objects.HighDemandPeriod
	call := GetHighDemandPeriodBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateHighDemandPeriodParams struct {
	BudgetValue     *uint64                                `facebook:"budget_value"`
	BudgetValueType *enums.HighdemandperiodBudgetValueType `facebook:"budget_value_type"`
	TimeEnd         *uint64                                `facebook:"time_end"`
	TimeStart       *uint64                                `facebook:"time_start"`
	Extra           core.Params                            `facebook:"-"`
}

func (params UpdateHighDemandPeriodParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BudgetValue != nil {
		out["budget_value"] = *params.BudgetValue
	}
	if params.BudgetValueType != nil {
		out["budget_value_type"] = *params.BudgetValueType
	}
	if params.TimeEnd != nil {
		out["time_end"] = *params.TimeEnd
	}
	if params.TimeStart != nil {
		out["time_start"] = *params.TimeStart
	}
	return out
}

func UpdateHighDemandPeriodBatchCall(id string, params UpdateHighDemandPeriodParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateHighDemandPeriodBatchRequest(id string, params UpdateHighDemandPeriodParams, options ...core.BatchOption) *core.BatchRequest[objects.HighDemandPeriod] {
	return core.NewBatchRequest[objects.HighDemandPeriod](UpdateHighDemandPeriodBatchCall(id, params, options...))
}

func DecodeUpdateHighDemandPeriodBatchResponse(response *core.BatchResponse) (*objects.HighDemandPeriod, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.HighDemandPeriod
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateHighDemandPeriod(ctx context.Context, client *core.Client, id string, params UpdateHighDemandPeriodParams) (*objects.HighDemandPeriod, error) {
	var out objects.HighDemandPeriod
	call := UpdateHighDemandPeriodBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
