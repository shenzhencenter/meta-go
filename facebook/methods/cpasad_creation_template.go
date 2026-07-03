package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASAdCreationTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASAdCreationTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASAdCreationTemplateBatchCall(id string, params GetCPASAdCreationTemplateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASAdCreationTemplateBatchRequest(id string, params GetCPASAdCreationTemplateParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASAdCreationTemplate] {
	return core.NewBatchRequest[objects.CPASAdCreationTemplate](GetCPASAdCreationTemplateBatchCall(id, params, options...))
}

func DecodeGetCPASAdCreationTemplateBatchResponse(response *core.BatchResponse) (*objects.CPASAdCreationTemplate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASAdCreationTemplate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASAdCreationTemplateWithResponse(ctx context.Context, client *core.Client, id string, params GetCPASAdCreationTemplateParams) (*objects.CPASAdCreationTemplate, *core.Response, error) {
	var out objects.CPASAdCreationTemplate
	call := GetCPASAdCreationTemplateBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCPASAdCreationTemplate(ctx context.Context, client *core.Client, id string, params GetCPASAdCreationTemplateParams) (*objects.CPASAdCreationTemplate, error) {
	out, _, err := GetCPASAdCreationTemplateWithResponse(ctx, client, id, params)
	return out, err
}
