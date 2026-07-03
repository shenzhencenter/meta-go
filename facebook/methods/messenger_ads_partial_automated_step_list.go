package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMessengerAdsPartialAutomatedStepListStepsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMessengerAdsPartialAutomatedStepListStepsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMessengerAdsPartialAutomatedStepListSteps(ctx context.Context, client *core.Client, id string, params GetMessengerAdsPartialAutomatedStepListStepsParams) (*core.Cursor[objects.MessengerAdsPartialAutomatedStep], error) {
	var out core.Cursor[objects.MessengerAdsPartialAutomatedStep]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "steps"), params.ToParams(), &out)
	return &out, err
}

type GetMessengerAdsPartialAutomatedStepListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMessengerAdsPartialAutomatedStepListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMessengerAdsPartialAutomatedStepList(ctx context.Context, client *core.Client, id string, params GetMessengerAdsPartialAutomatedStepListParams) (*objects.MessengerAdsPartialAutomatedStepList, error) {
	var out objects.MessengerAdsPartialAutomatedStepList
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
