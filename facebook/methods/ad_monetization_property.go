package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdMonetizationPropertyAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
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

func CreateAdMonetizationPropertyAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateAdMonetizationPropertyAdnetworkanalyticsParams) (*objects.AdMonetizationProperty, error) {
	var out objects.AdMonetizationProperty
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
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

func GetAdMonetizationPropertyAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), &out)
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

func GetAdMonetizationProperty(ctx context.Context, client *core.Client, id string, params GetAdMonetizationPropertyParams) (*objects.AdMonetizationProperty, error) {
	var out objects.AdMonetizationProperty
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
