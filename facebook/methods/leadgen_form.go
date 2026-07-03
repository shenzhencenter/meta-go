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

func GetLeadgenFormLeads(ctx context.Context, client *core.Client, id string, params GetLeadgenFormLeadsParams) (*core.Cursor[objects.Lead], error) {
	var out core.Cursor[objects.Lead]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "leads"), params.ToParams(), &out)
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

func GetLeadgenFormTestLeads(ctx context.Context, client *core.Client, id string, params GetLeadgenFormTestLeadsParams) (*core.Cursor[objects.Lead], error) {
	var out core.Cursor[objects.Lead]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "test_leads"), params.ToParams(), &out)
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

func CreateLeadgenFormTestLeads(ctx context.Context, client *core.Client, id string, params CreateLeadgenFormTestLeadsParams) (*objects.Lead, error) {
	var out objects.Lead
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "test_leads"), params.ToParams(), &out)
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

func GetLeadgenForm(ctx context.Context, client *core.Client, id string, params GetLeadgenFormParams) (*objects.LeadgenForm, error) {
	var out objects.LeadgenForm
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateLeadgenForm(ctx context.Context, client *core.Client, id string, params UpdateLeadgenFormParams) (*objects.LeadgenForm, error) {
	var out objects.LeadgenForm
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
