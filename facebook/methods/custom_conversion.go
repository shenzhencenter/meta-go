package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
)

type GetCustomConversionStatsParams struct {
	Aggregation *enums.CustomconversionstatsAggregationEnumParam `facebook:"aggregation"`
	EndTime     *time.Time                                       `facebook:"end_time"`
	StartTime   *time.Time                                       `facebook:"start_time"`
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

func GetCustomConversionStats(ctx context.Context, client *core.Client, id string, params GetCustomConversionStatsParams) (*core.Cursor[objects.CustomConversionStatsResult], error) {
	var out core.Cursor[objects.CustomConversionStatsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), &out)
	return &out, err
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

func DeleteCustomConversion(ctx context.Context, client *core.Client, id string, params DeleteCustomConversionParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetCustomConversion(ctx context.Context, client *core.Client, id string, params GetCustomConversionParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateCustomConversion(ctx context.Context, client *core.Client, id string, params UpdateCustomConversionParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
