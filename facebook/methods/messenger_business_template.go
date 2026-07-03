package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetMessengerBusinessTemplate(ctx context.Context, client *core.Client, id string, params GetMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, error) {
	var out objects.MessengerBusinessTemplate
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateMessengerBusinessTemplate(ctx context.Context, client *core.Client, id string, params UpdateMessengerBusinessTemplateParams) (*objects.MessengerBusinessTemplate, error) {
	var out objects.MessengerBusinessTemplate
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
