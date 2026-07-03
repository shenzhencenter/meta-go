package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetResearchPollStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetResearchPollStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetResearchPollStudyBatchCall(id string, params GetResearchPollStudyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetResearchPollStudyBatchRequest(id string, params GetResearchPollStudyParams, options ...core.BatchOption) *core.BatchRequest[objects.ResearchPollStudy] {
	return core.NewBatchRequest[objects.ResearchPollStudy](GetResearchPollStudyBatchCall(id, params, options...))
}

func DecodeGetResearchPollStudyBatchResponse(response *core.BatchResponse) (*objects.ResearchPollStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ResearchPollStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetResearchPollStudy(ctx context.Context, client *core.Client, id string, params GetResearchPollStudyParams) (*objects.ResearchPollStudy, error) {
	var out objects.ResearchPollStudy
	call := GetResearchPollStudyBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
