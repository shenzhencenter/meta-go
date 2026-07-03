package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLeadgenFormLeadsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadgenFormLeadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadgenFormLeadsBatchCall(id string, params GetLeadgenFormLeadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "leads"), params.ToParams(), options...)
}

func NewGetLeadgenFormLeadsBatchRequest(id string, params GetLeadgenFormLeadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Lead]] {
	return core.NewBatchRequest[core.Cursor[objects.Lead]](GetLeadgenFormLeadsBatchCall(id, params, options...))
}

func DecodeGetLeadgenFormLeadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Lead], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Lead]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLeadgenFormLeads(ctx context.Context, client *core.Client, id string, params GetLeadgenFormLeadsParams) (*core.Cursor[objects.Lead], error) {
	var out core.Cursor[objects.Lead]
	call := GetLeadgenFormLeadsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLeadgenFormTestLeadsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadgenFormTestLeadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadgenFormTestLeadsBatchCall(id string, params GetLeadgenFormTestLeadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "test_leads"), params.ToParams(), options...)
}

func NewGetLeadgenFormTestLeadsBatchRequest(id string, params GetLeadgenFormTestLeadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Lead]] {
	return core.NewBatchRequest[core.Cursor[objects.Lead]](GetLeadgenFormTestLeadsBatchCall(id, params, options...))
}

func DecodeGetLeadgenFormTestLeadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Lead], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Lead]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLeadgenFormTestLeads(ctx context.Context, client *core.Client, id string, params GetLeadgenFormTestLeadsParams) (*core.Cursor[objects.Lead], error) {
	var out core.Cursor[objects.Lead]
	call := GetLeadgenFormTestLeadsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateLeadgenFormTestLeadsParams struct {
	CustomDisclaimerResponses *[]map[string]interface{} `facebook:"custom_disclaimer_responses"`
	FieldData                 *[]map[string]interface{} `facebook:"field_data"`
	Extra                     core.Params               `facebook:"-"`
}

func (params CreateLeadgenFormTestLeadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CustomDisclaimerResponses != nil {
		out["custom_disclaimer_responses"] = *params.CustomDisclaimerResponses
	}
	if params.FieldData != nil {
		out["field_data"] = *params.FieldData
	}
	return out
}

func CreateLeadgenFormTestLeadsBatchCall(id string, params CreateLeadgenFormTestLeadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "test_leads"), params.ToParams(), options...)
}

func NewCreateLeadgenFormTestLeadsBatchRequest(id string, params CreateLeadgenFormTestLeadsParams, options ...core.BatchOption) *core.BatchRequest[objects.Lead] {
	return core.NewBatchRequest[objects.Lead](CreateLeadgenFormTestLeadsBatchCall(id, params, options...))
}

func DecodeCreateLeadgenFormTestLeadsBatchResponse(response *core.BatchResponse) (*objects.Lead, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Lead
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateLeadgenFormTestLeads(ctx context.Context, client *core.Client, id string, params CreateLeadgenFormTestLeadsParams) (*objects.Lead, error) {
	var out objects.Lead
	call := CreateLeadgenFormTestLeadsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLeadgenFormParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadgenFormParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadgenFormBatchCall(id string, params GetLeadgenFormParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLeadgenFormBatchRequest(id string, params GetLeadgenFormParams, options ...core.BatchOption) *core.BatchRequest[objects.LeadgenForm] {
	return core.NewBatchRequest[objects.LeadgenForm](GetLeadgenFormBatchCall(id, params, options...))
}

func DecodeGetLeadgenFormBatchResponse(response *core.BatchResponse) (*objects.LeadgenForm, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LeadgenForm
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLeadgenForm(ctx context.Context, client *core.Client, id string, params GetLeadgenFormParams) (*objects.LeadgenForm, error) {
	var out objects.LeadgenForm
	call := GetLeadgenFormBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateLeadgenFormParams struct {
	Status *enums.LeadgendataStatus `facebook:"status"`
	Extra  core.Params              `facebook:"-"`
}

func (params UpdateLeadgenFormParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func UpdateLeadgenFormBatchCall(id string, params UpdateLeadgenFormParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateLeadgenFormBatchRequest(id string, params UpdateLeadgenFormParams, options ...core.BatchOption) *core.BatchRequest[objects.LeadgenForm] {
	return core.NewBatchRequest[objects.LeadgenForm](UpdateLeadgenFormBatchCall(id, params, options...))
}

func DecodeUpdateLeadgenFormBatchResponse(response *core.BatchResponse) (*objects.LeadgenForm, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LeadgenForm
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateLeadgenForm(ctx context.Context, client *core.Client, id string, params UpdateLeadgenFormParams) (*objects.LeadgenForm, error) {
	var out objects.LeadgenForm
	call := UpdateLeadgenFormBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
