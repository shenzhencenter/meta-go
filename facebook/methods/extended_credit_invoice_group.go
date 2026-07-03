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

func DeleteExtendedCreditInvoiceGroupAdAccountsBatchCall(id string, params DeleteExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "ad_accounts"), params.ToParams(), options...)
}

func NewDeleteExtendedCreditInvoiceGroupAdAccountsBatchRequest(id string, params DeleteExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params, options...))
}

func DecodeDeleteExtendedCreditInvoiceGroupAdAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupAdAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupAdAccountsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func GetExtendedCreditInvoiceGroupAdAccountsBatchCall(id string, params GetExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_accounts"), params.ToParams(), options...)
}

func NewGetExtendedCreditInvoiceGroupAdAccountsBatchRequest(id string, params GetExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditInvoiceGroupAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupAdAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateExtendedCreditInvoiceGroupAdAccountsBatchCall(id string, params CreateExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_accounts"), params.ToParams(), options...)
}

func NewCreateExtendedCreditInvoiceGroupAdAccountsBatchRequest(id string, params CreateExtendedCreditInvoiceGroupAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditInvoiceGroupAdAccountsBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditInvoiceGroupAdAccountsParams) (*objects.AdAccount, *core.Response, error) {
	var out objects.AdAccount
	call := CreateExtendedCreditInvoiceGroupAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditInvoiceGroupAdAccounts(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditInvoiceGroupAdAccountsParams) (*objects.AdAccount, error) {
	out, _, err := CreateExtendedCreditInvoiceGroupAdAccountsWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteExtendedCreditInvoiceGroupBatchCall(id string, params DeleteExtendedCreditInvoiceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteExtendedCreditInvoiceGroupBatchRequest(id string, params DeleteExtendedCreditInvoiceGroupParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteExtendedCreditInvoiceGroupBatchCall(id, params, options...))
}

func DecodeDeleteExtendedCreditInvoiceGroupBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteExtendedCreditInvoiceGroupWithResponse(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteExtendedCreditInvoiceGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params DeleteExtendedCreditInvoiceGroupParams) (*map[string]interface{}, error) {
	out, _, err := DeleteExtendedCreditInvoiceGroupWithResponse(ctx, client, id, params)
	return out, err
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

func GetExtendedCreditInvoiceGroupBatchCall(id string, params GetExtendedCreditInvoiceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetExtendedCreditInvoiceGroupBatchRequest(id string, params GetExtendedCreditInvoiceGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditInvoiceGroup] {
	return core.NewBatchRequest[objects.ExtendedCreditInvoiceGroup](GetExtendedCreditInvoiceGroupBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditInvoiceGroupBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditInvoiceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditInvoiceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditInvoiceGroupWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, *core.Response, error) {
	var out objects.ExtendedCreditInvoiceGroup
	call := GetExtendedCreditInvoiceGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params GetExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	out, _, err := GetExtendedCreditInvoiceGroupWithResponse(ctx, client, id, params)
	return out, err
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

func UpdateExtendedCreditInvoiceGroupBatchCall(id string, params UpdateExtendedCreditInvoiceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateExtendedCreditInvoiceGroupBatchRequest(id string, params UpdateExtendedCreditInvoiceGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditInvoiceGroup] {
	return core.NewBatchRequest[objects.ExtendedCreditInvoiceGroup](UpdateExtendedCreditInvoiceGroupBatchCall(id, params, options...))
}

func DecodeUpdateExtendedCreditInvoiceGroupBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditInvoiceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditInvoiceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateExtendedCreditInvoiceGroupWithResponse(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, *core.Response, error) {
	var out objects.ExtendedCreditInvoiceGroup
	call := UpdateExtendedCreditInvoiceGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateExtendedCreditInvoiceGroup(ctx context.Context, client *core.Client, id string, params UpdateExtendedCreditInvoiceGroupParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	out, _, err := UpdateExtendedCreditInvoiceGroupWithResponse(ctx, client, id, params)
	return out, err
}
