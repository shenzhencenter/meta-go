package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func CreateAdCampaignGroupBudgetSchedulesBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateAdCampaignGroupBudgetSchedulesBudgetSchedulesParams) (*objects.AdCampaignGroupBudgetSchedulesPost, error) {
	var out objects.AdCampaignGroupBudgetSchedulesPost
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), &out)
	return &out, err
}
