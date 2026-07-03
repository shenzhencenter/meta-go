package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWorkExperienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWorkExperienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWorkExperienceBatchCall(id string, params GetWorkExperienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWorkExperienceBatchRequest(id string, params GetWorkExperienceParams, options ...core.BatchOption) *core.BatchRequest[objects.WorkExperience] {
	return core.NewBatchRequest[objects.WorkExperience](GetWorkExperienceBatchCall(id, params, options...))
}

func DecodeGetWorkExperienceBatchResponse(response *core.BatchResponse) (*objects.WorkExperience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WorkExperience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWorkExperience(ctx context.Context, client *core.Client, id string, params GetWorkExperienceParams) (*objects.WorkExperience, error) {
	var out objects.WorkExperience
	call := GetWorkExperienceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
