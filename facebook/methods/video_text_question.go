package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoTextQuestionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoTextQuestionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoTextQuestionBatchCall(id string, params GetVideoTextQuestionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoTextQuestionBatchRequest(id string, params GetVideoTextQuestionParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoTextQuestion] {
	return core.NewBatchRequest[objects.VideoTextQuestion](GetVideoTextQuestionBatchCall(id, params, options...))
}

func DecodeGetVideoTextQuestionBatchResponse(response *core.BatchResponse) (*objects.VideoTextQuestion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoTextQuestion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoTextQuestion(ctx context.Context, client *core.Client, id string, params GetVideoTextQuestionParams) (*objects.VideoTextQuestion, error) {
	var out objects.VideoTextQuestion
	call := GetVideoTextQuestionBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
