package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGMediaBoostEligibilityInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaBoostEligibilityInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaBoostEligibilityInfoBatchCall(id string, params GetIGMediaBoostEligibilityInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGMediaBoostEligibilityInfoBatchRequest(id string, params GetIGMediaBoostEligibilityInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.IGMediaBoostEligibilityInfo] {
	return core.NewBatchRequest[objects.IGMediaBoostEligibilityInfo](GetIGMediaBoostEligibilityInfoBatchCall(id, params, options...))
}

func DecodeGetIGMediaBoostEligibilityInfoBatchResponse(response *core.BatchResponse) (*objects.IGMediaBoostEligibilityInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGMediaBoostEligibilityInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaBoostEligibilityInfoWithResponse(ctx context.Context, client *core.Client, id string, params GetIGMediaBoostEligibilityInfoParams) (*objects.IGMediaBoostEligibilityInfo, *core.Response, error) {
	var out objects.IGMediaBoostEligibilityInfo
	call := GetIGMediaBoostEligibilityInfoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGMediaBoostEligibilityInfo(ctx context.Context, client *core.Client, id string, params GetIGMediaBoostEligibilityInfoParams) (*objects.IGMediaBoostEligibilityInfo, error) {
	out, _, err := GetIGMediaBoostEligibilityInfoWithResponse(ctx, client, id, params)
	return out, err
}
