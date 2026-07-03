package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetAdMonetizationPropertyAdnetworkanalyticsParams struct {
	AggregationPeriod  *enums.AdmonetizationpropertyadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns         *[]enums.AdmonetizationpropertyadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters            *[]map[string]interface{}                                                 `facebook:"filters"`
	Limit              *uint64                                                                   `facebook:"limit"`
	Metrics            []enums.AdmonetizationpropertyadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn     *enums.AdmonetizationpropertyadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType       *enums.AdmonetizationpropertyadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	ShouldIncludeUntil *bool                                                                     `facebook:"should_include_until"`
	Since              *time.Time                                                                `facebook:"since"`
	Until              *time.Time                                                                `facebook:"until"`
	Extra              core.Params                                                               `facebook:"-"`
}

func (params GetAdMonetizationPropertyAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.ShouldIncludeUntil != nil {
		out["should_include_until"] = *params.ShouldIncludeUntil
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetAdMonetizationPropertyAdnetworkanalyticsBatchCall(id string, params GetAdMonetizationPropertyAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewGetAdMonetizationPropertyAdnetworkanalyticsBatchRequest(id string, params GetAdMonetizationPropertyAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]](GetAdMonetizationPropertyAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeGetAdMonetizationPropertyAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdMonetizationPropertyAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	call := GetAdMonetizationPropertyAdnetworkanalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdMonetizationPropertyAdnetworkanalyticsParams struct {
	AggregationPeriod *enums.AdmonetizationpropertyadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns        *[]enums.AdmonetizationpropertyadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters           *[]map[string]interface{}                                                 `facebook:"filters"`
	Limit             *int                                                                      `facebook:"limit"`
	Metrics           []enums.AdmonetizationpropertyadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn    *enums.AdmonetizationpropertyadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType      *enums.AdmonetizationpropertyadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	Since             *time.Time                                                                `facebook:"since"`
	Until             *time.Time                                                                `facebook:"until"`
	Extra             core.Params                                                               `facebook:"-"`
}

func (params CreateAdMonetizationPropertyAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func CreateAdMonetizationPropertyAdnetworkanalyticsBatchCall(id string, params CreateAdMonetizationPropertyAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewCreateAdMonetizationPropertyAdnetworkanalyticsBatchRequest(id string, params CreateAdMonetizationPropertyAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdMonetizationProperty] {
	return core.NewBatchRequest[objects.AdMonetizationProperty](CreateAdMonetizationPropertyAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeCreateAdMonetizationPropertyAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*objects.AdMonetizationProperty, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdMonetizationProperty
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdMonetizationPropertyAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateAdMonetizationPropertyAdnetworkanalyticsParams) (*objects.AdMonetizationProperty, error) {
	var out objects.AdMonetizationProperty
	call := CreateAdMonetizationPropertyAdnetworkanalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdMonetizationPropertyAdnetworkanalyticsResultsParams struct {
	QueryIds *[]core.ID  `facebook:"query_ids"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdMonetizationPropertyAdnetworkanalyticsResultsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.QueryIds != nil {
		out["query_ids"] = *params.QueryIds
	}
	return out
}

func GetAdMonetizationPropertyAdnetworkanalyticsResultsBatchCall(id string, params GetAdMonetizationPropertyAdnetworkanalyticsResultsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), options...)
}

func NewGetAdMonetizationPropertyAdnetworkanalyticsResultsBatchRequest(id string, params GetAdMonetizationPropertyAdnetworkanalyticsResultsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]](GetAdMonetizationPropertyAdnetworkanalyticsResultsBatchCall(id, params, options...))
}

func DecodeGetAdMonetizationPropertyAdnetworkanalyticsResultsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdMonetizationPropertyAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	call := GetAdMonetizationPropertyAdnetworkanalyticsResultsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdMonetizationPropertyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdMonetizationPropertyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdMonetizationPropertyBatchCall(id string, params GetAdMonetizationPropertyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdMonetizationPropertyBatchRequest(id string, params GetAdMonetizationPropertyParams, options ...core.BatchOption) *core.BatchRequest[objects.AdMonetizationProperty] {
	return core.NewBatchRequest[objects.AdMonetizationProperty](GetAdMonetizationPropertyBatchCall(id, params, options...))
}

func DecodeGetAdMonetizationPropertyBatchResponse(response *core.BatchResponse) (*objects.AdMonetizationProperty, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdMonetizationProperty
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdMonetizationProperty(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyParams) (*objects.AdMonetizationProperty, error) {
	var out objects.AdMonetizationProperty
	call := GetAdMonetizationPropertyBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
