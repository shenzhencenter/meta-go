package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCanvasTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasTemplateBatchCall(id string, params GetCanvasTemplateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCanvasTemplateBatchRequest(id string, params GetCanvasTemplateParams, options ...core.BatchOption) *core.BatchRequest[objects.CanvasTemplate] {
	return core.NewBatchRequest[objects.CanvasTemplate](GetCanvasTemplateBatchCall(id, params, options...))
}

func DecodeGetCanvasTemplateBatchResponse(response *core.BatchResponse) (*objects.CanvasTemplate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CanvasTemplate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCanvasTemplate(ctx context.Context, client *core.Client, id string, params GetCanvasTemplateParams) (*objects.CanvasTemplate, error) {
	var out objects.CanvasTemplate
	call := GetCanvasTemplateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
