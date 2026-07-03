package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func CreateAdCampaignBudgetSchedulesBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateAdCampaignBudgetSchedulesBudgetSchedulesParams) (*objects.AdCampaignBudgetSchedulesPost, error) {
	var out objects.AdCampaignBudgetSchedulesPost
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), &out)
	return &out, err
}
