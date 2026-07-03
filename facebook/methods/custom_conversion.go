package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCustomConversionStatsParams struct {
	Aggregation *enums.CustomconversionstatsAggregationEnumParam `facebook:"aggregation"`
	EndTime     *core.Time                                       `facebook:"end_time"`
	StartTime   *core.Time                                       `facebook:"start_time"`
	Extra       core.Params                                      `facebook:"-"`
}

func (params GetCustomConversionStatsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Aggregation != nil {
		out["aggregation"] = *params.Aggregation
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetCustomConversionStatsBatchCall(id string, params GetCustomConversionStatsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), options...)
}

func NewGetCustomConversionStatsBatchRequest(id string, params GetCustomConversionStatsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomConversionStatsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomConversionStatsResult]](GetCustomConversionStatsBatchCall(id, params, options...))
}

func DecodeGetCustomConversionStatsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomConversionStatsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomConversionStatsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomConversionStatsWithResponse(ctx context.Context, client *core.Client, id string, params GetCustomConversionStatsParams) (*core.Cursor[objects.CustomConversionStatsResult], *core.Response, error) {
	var out core.Cursor[objects.CustomConversionStatsResult]
	call := GetCustomConversionStatsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCustomConversionStats(ctx context.Context, client *core.Client, id string, params GetCustomConversionStatsParams) (*core.Cursor[objects.CustomConversionStatsResult], error) {
	out, _, err := GetCustomConversionStatsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteCustomConversionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteCustomConversionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteCustomConversionBatchCall(id string, params DeleteCustomConversionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteCustomConversionBatchRequest(id string, params DeleteCustomConversionParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCustomConversionBatchCall(id, params, options...))
}

func DecodeDeleteCustomConversionBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCustomConversionWithResponse(ctx context.Context, client *core.Client, id string, params DeleteCustomConversionParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteCustomConversionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteCustomConversion(ctx context.Context, client *core.Client, id string, params DeleteCustomConversionParams) (*map[string]interface{}, error) {
	out, _, err := DeleteCustomConversionWithResponse(ctx, client, id, params)
	return out, err
}

type GetCustomConversionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCustomConversionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCustomConversionBatchCall(id string, params GetCustomConversionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCustomConversionBatchRequest(id string, params GetCustomConversionParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomConversion] {
	return core.NewBatchRequest[objects.CustomConversion](GetCustomConversionBatchCall(id, params, options...))
}

func DecodeGetCustomConversionBatchResponse(response *core.BatchResponse) (*objects.CustomConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomConversionWithResponse(ctx context.Context, client *core.Client, id string, params GetCustomConversionParams) (*objects.CustomConversion, *core.Response, error) {
	var out objects.CustomConversion
	call := GetCustomConversionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCustomConversion(ctx context.Context, client *core.Client, id string, params GetCustomConversionParams) (*objects.CustomConversion, error) {
	out, _, err := GetCustomConversionWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateCustomConversionParams struct {
	DefaultConversionValue *float64    `facebook:"default_conversion_value"`
	Description            *string     `facebook:"description"`
	Name                   *string     `facebook:"name"`
	Extra                  core.Params `facebook:"-"`
}

func (params UpdateCustomConversionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DefaultConversionValue != nil {
		out["default_conversion_value"] = *params.DefaultConversionValue
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateCustomConversionBatchCall(id string, params UpdateCustomConversionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCustomConversionBatchRequest(id string, params UpdateCustomConversionParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomConversion] {
	return core.NewBatchRequest[objects.CustomConversion](UpdateCustomConversionBatchCall(id, params, options...))
}

func DecodeUpdateCustomConversionBatchResponse(response *core.BatchResponse) (*objects.CustomConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCustomConversionWithResponse(ctx context.Context, client *core.Client, id string, params UpdateCustomConversionParams) (*objects.CustomConversion, *core.Response, error) {
	var out objects.CustomConversion
	call := UpdateCustomConversionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateCustomConversion(ctx context.Context, client *core.Client, id string, params UpdateCustomConversionParams) (*objects.CustomConversion, error) {
	out, _, err := UpdateCustomConversionWithResponse(ctx, client, id, params)
	return out, err
}
