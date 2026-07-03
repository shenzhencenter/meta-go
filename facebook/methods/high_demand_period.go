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

func DeleteHighDemandPeriod(ctx context.Context, client *core.Client, id string, params DeleteHighDemandPeriodParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetHighDemandPeriod(ctx context.Context, client *core.Client, id string, params GetHighDemandPeriodParams) (*objects.HighDemandPeriod, error) {
	var out objects.HighDemandPeriod
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateHighDemandPeriod(ctx context.Context, client *core.Client, id string, params UpdateHighDemandPeriodParams) (*objects.HighDemandPeriod, error) {
	var out objects.HighDemandPeriod
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
