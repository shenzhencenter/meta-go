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

func GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersBatchCall(id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "partners"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessPreVerifiedPhoneNumberPartnersBatchRequest(id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessPreVerifiedPhoneNumberPartnersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersWithResponse(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWhatsAppBusinessPreVerifiedPhoneNumberPartners(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetWhatsAppBusinessPreVerifiedPhoneNumberPartnersWithResponse(ctx, client, id, params)
	return out, err
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

func CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeBatchCall(id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "request_code"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeBatchRequest(id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeWithResponse(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCode(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeParams) (*map[string]interface{}, error) {
	out, _, err := CreateWhatsAppBusinessPreVerifiedPhoneNumberRequestCodeWithResponse(ctx, client, id, params)
	return out, err
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

func CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeBatchCall(id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "verify_code"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeBatchRequest(id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeWithResponse(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCode(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeParams) (*map[string]interface{}, error) {
	out, _, err := CreateWhatsAppBusinessPreVerifiedPhoneNumberVerifyCodeWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id string, params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessPreVerifiedPhoneNumberBatchRequest(id string, params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessPreVerifiedPhoneNumberBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteWhatsAppBusinessPreVerifiedPhoneNumberWithResponse(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteWhatsAppBusinessPreVerifiedPhoneNumber(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessPreVerifiedPhoneNumberParams) (*map[string]interface{}, error) {
	out, _, err := DeleteWhatsAppBusinessPreVerifiedPhoneNumberWithResponse(ctx, client, id, params)
	return out, err
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

func GetWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessPreVerifiedPhoneNumberBatchRequest(id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessPreVerifiedPhoneNumber] {
	return core.NewBatchRequest[objects.WhatsAppBusinessPreVerifiedPhoneNumber](GetWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessPreVerifiedPhoneNumberBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessPreVerifiedPhoneNumber, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessPreVerifiedPhoneNumber
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessPreVerifiedPhoneNumberWithResponse(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberParams) (*objects.WhatsAppBusinessPreVerifiedPhoneNumber, *core.Response, error) {
	var out objects.WhatsAppBusinessPreVerifiedPhoneNumber
	call := GetWhatsAppBusinessPreVerifiedPhoneNumberBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWhatsAppBusinessPreVerifiedPhoneNumber(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessPreVerifiedPhoneNumberParams) (*objects.WhatsAppBusinessPreVerifiedPhoneNumber, error) {
	out, _, err := GetWhatsAppBusinessPreVerifiedPhoneNumberWithResponse(ctx, client, id, params)
	return out, err
}
