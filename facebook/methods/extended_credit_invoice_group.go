package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteExtendedCreditInvoiceGroupAdAccountsParams struct {
	AdAccountID core.ID     `facebook:"ad_account_id"`
	Extra       core.Params `facebook:"-"`
}

func (params DeleteExtendedCreditInvoiceGroupAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_account_id"] = params.AdAccountID
	return out
}

func DeleteExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupAdAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetExtendedCreditInvoiceGroupAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditInvoiceGroupAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateExtendedCreditInvoiceGroupAdAccountsParams struct {
	AdAccountID core.ID     `facebook:"ad_account_id"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateExtendedCreditInvoiceGroupAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_account_id"] = params.AdAccountID
	return out
}

func CreateExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditInvoiceGroupAdAccountsParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteExtendedCreditInvoiceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteExtendedCreditInvoiceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetExtendedCreditInvoiceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditInvoiceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	var out objects.ExtendedCreditInvoiceGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateExtendedCreditInvoiceGroupParams struct {
	Emails *[]string   `facebook:"emails"`
	Name   *string     `facebook:"name"`
	Extra  core.Params `facebook:"-"`
}

func (params UpdateExtendedCreditInvoiceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Emails != nil {
		out["emails"] = *params.Emails
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	var out objects.ExtendedCreditInvoiceGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
