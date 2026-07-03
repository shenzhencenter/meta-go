package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMessengerBusinessTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMessengerBusinessTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMessengerBusinessTemplateBatchCall(id string, params GetMessengerBusinessTemplateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMessengerBusinessTemplateBatchRequest(id string, params GetMessengerBusinessTemplateParams, options ...core.BatchOption) *core.BatchRequest[objects.MessengerBusinessTemplate] {
	return core.NewBatchRequest[objects.MessengerBusinessTemplate](GetMessengerBusinessTemplateBatchCall(id, params, options...))
}

func DecodeGetMessengerBusinessTemplateBatchResponse(response *core.BatchResponse) (*objects.MessengerBusinessTemplate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MessengerBusinessTemplate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMessengerBusinessTemplateWithResponse(ctx context.Context, client *core.Client, id string, params GetMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, *core.Response, error) {
	var out objects.MessengerBusinessTemplate
	call := GetMessengerBusinessTemplateBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetMessengerBusinessTemplate(ctx context.Context, client *core.Client, id string, params GetMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, error) {
	out, _, err := GetMessengerBusinessTemplateWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateMessengerBusinessTemplateParams struct {
	Components      *[]map[string]interface{}                       `facebook:"components"`
	ParameterFormat *enums.MessengerbusinesstemplateParameterFormat `facebook:"parameter_format"`
	Extra           core.Params                                     `facebook:"-"`
}

func (params UpdateMessengerBusinessTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Components != nil {
		out["components"] = *params.Components
	}
	if params.ParameterFormat != nil {
		out["parameter_format"] = *params.ParameterFormat
	}
	return out
}

func UpdateMessengerBusinessTemplateBatchCall(id string, params UpdateMessengerBusinessTemplateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateMessengerBusinessTemplateBatchRequest(id string, params UpdateMessengerBusinessTemplateParams, options ...core.BatchOption) *core.BatchRequest[objects.MessengerBusinessTemplate] {
	return core.NewBatchRequest[objects.MessengerBusinessTemplate](UpdateMessengerBusinessTemplateBatchCall(id, params, options...))
}

func DecodeUpdateMessengerBusinessTemplateBatchResponse(response *core.BatchResponse) (*objects.MessengerBusinessTemplate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MessengerBusinessTemplate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateMessengerBusinessTemplateWithResponse(ctx context.Context, client *core.Client, id string, params UpdateMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, *core.Response, error) {
	var out objects.MessengerBusinessTemplate
	call := UpdateMessengerBusinessTemplateBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateMessengerBusinessTemplate(ctx context.Context, client *core.Client, id string, params UpdateMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, error) {
	out, _, err := UpdateMessengerBusinessTemplateWithResponse(ctx, client, id, params)
	return out, err
}
