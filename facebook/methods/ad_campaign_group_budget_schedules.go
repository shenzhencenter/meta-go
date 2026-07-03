package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesBatchCall(id string, params CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), options...)
}

func NewCreateAdCampaignGroupBudgetSchedulesBudgetSchedulesBatchRequest(id string, params CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCampaignGroupBudgetSchedulesPost] {
	return core.NewBatchRequest[objects.AdCampaignGroupBudgetSchedulesPost](CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesBatchCall(id, params, options...))
}

func DecodeCreateAdCampaignGroupBudgetSchedulesBudgetSchedulesBatchResponse(response *core.BatchResponse) (*objects.AdCampaignGroupBudgetSchedulesPost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCampaignGroupBudgetSchedulesPost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdCampaignGroupBudgetSchedulesBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams) (*objects.AdCampaignGroupBudgetSchedulesPost, error) {
	var out objects.AdCampaignGroupBudgetSchedulesPost
	call := CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
