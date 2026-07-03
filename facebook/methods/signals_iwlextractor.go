package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSignalsIWLExtractorParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSignalsIWLExtractorParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSignalsIWLExtractorBatchCall(id string, params GetSignalsIWLExtractorParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSignalsIWLExtractorBatchRequest(id string, params GetSignalsIWLExtractorParams, options ...core.BatchOption) *core.BatchRequest[objects.SignalsIWLExtractor] {
	return core.NewBatchRequest[objects.SignalsIWLExtractor](GetSignalsIWLExtractorBatchCall(id, params, options...))
}

func DecodeGetSignalsIWLExtractorBatchResponse(response *core.BatchResponse) (*objects.SignalsIWLExtractor, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SignalsIWLExtractor
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSignalsIWLExtractor(ctx context.Context, client *core.Client, id string, params GetSignalsIWLExtractorParams) (*objects.SignalsIWLExtractor, error) {
	var out objects.SignalsIWLExtractor
	call := GetSignalsIWLExtractorBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
