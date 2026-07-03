package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateAdCampaignBudgetSchedulesBudgetSchedulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdCampaignBudgetSchedulesBudgetSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdCampaignBudgetSchedulesBudgetSchedulesBatchCall(id string, params CreateAdCampaignBudgetSchedulesBudgetSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), options...)
}

func NewCreateAdCampaignBudgetSchedulesBudgetSchedulesBatchRequest(id string, params CreateAdCampaignBudgetSchedulesBudgetSchedulesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCampaignBudgetSchedulesPost] {
	return core.NewBatchRequest[objects.AdCampaignBudgetSchedulesPost](CreateAdCampaignBudgetSchedulesBudgetSchedulesBatchCall(id, params, options...))
}

func DecodeCreateAdCampaignBudgetSchedulesBudgetSchedulesBatchResponse(response *core.BatchResponse) (*objects.AdCampaignBudgetSchedulesPost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCampaignBudgetSchedulesPost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdCampaignBudgetSchedulesBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateAdCampaignBudgetSchedulesBudgetSchedulesParams) (*objects.AdCampaignBudgetSchedulesPost, error) {
	var out objects.AdCampaignBudgetSchedulesPost
	call := CreateAdCampaignBudgetSchedulesBudgetSchedulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
