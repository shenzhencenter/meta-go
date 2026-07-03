package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEducationExperienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEducationExperienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEducationExperienceBatchCall(id string, params GetEducationExperienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEducationExperienceBatchRequest(id string, params GetEducationExperienceParams, options ...core.BatchOption) *core.BatchRequest[objects.EducationExperience] {
	return core.NewBatchRequest[objects.EducationExperience](GetEducationExperienceBatchCall(id, params, options...))
}

func DecodeGetEducationExperienceBatchResponse(response *core.BatchResponse) (*objects.EducationExperience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EducationExperience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEducationExperienceWithResponse(ctx context.Context, client *core.Client, id string, params GetEducationExperienceParams) (*objects.EducationExperience, *core.Response, error) {
	var out objects.EducationExperience
	call := GetEducationExperienceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEducationExperience(ctx context.Context, client *core.Client, id string, params GetEducationExperienceParams) (*objects.EducationExperience, error) {
	out, _, err := GetEducationExperienceWithResponse(ctx, client, id, params)
	return out, err
}
