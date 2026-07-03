package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessPreVerifiedPhoneNumberPartners(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "partners"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams struct {
	CodeMethod enums.WhatsappbusinesspreverifiedphonenumberrequestCodeCodeMethodEnumParam `facebook:"code_method"`
	Language   string                                                                     `facebook:"language"`
	Extra      core.Params                                                                `facebook:"-"`
}

func (params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["code_method"] = params.CodeMethod
	out["language"] = params.Language
	return out
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCode(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "request_code"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams struct {
	Code  string      `facebook:"code"`
	Extra core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["code"] = params.Code
	return out
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCode(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "verify_code"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteWhatsAppBusinessPreVerifiedPhoneNumber(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessPreVerifiedPhoneNumberParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessPreVerifiedPhoneNumberParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessPreVerifiedPhoneNumber(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberParams) (*objects.WhatsAppBusinessPreVerifiedPhoneNumber, error) {
	var out objects.WhatsAppBusinessPreVerifiedPhoneNumber
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
