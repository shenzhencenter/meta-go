package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPartnerStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerStudyBatchCall(id string, params GetPartnerStudyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPartnerStudyBatchRequest(id string, params GetPartnerStudyParams, options ...core.BatchOption) *core.BatchRequest[objects.PartnerStudy] {
	return core.NewBatchRequest[objects.PartnerStudy](GetPartnerStudyBatchCall(id, params, options...))
}

func DecodeGetPartnerStudyBatchResponse(response *core.BatchResponse) (*objects.PartnerStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PartnerStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPartnerStudyWithResponse(ctx context.Context, client *core.Client, id string, params GetPartnerStudyParams) (*objects.PartnerStudy, *core.Response, error) {
	var out objects.PartnerStudy
	call := GetPartnerStudyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPartnerStudy(ctx context.Context, client *core.Client, id string, params GetPartnerStudyParams) (*objects.PartnerStudy, error) {
	out, _, err := GetPartnerStudyWithResponse(ctx, client, id, params)
	return out, err
}
