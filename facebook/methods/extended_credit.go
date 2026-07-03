package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetExtendedCreditExtendedCreditInvoiceGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditExtendedCreditInvoiceGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id string, params GetExtendedCreditExtendedCreditInvoiceGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "extended_credit_invoice_groups"), params.ToParams(), options...)
}

func NewGetExtendedCreditExtendedCreditInvoiceGroupsBatchRequest(id string, params GetExtendedCreditExtendedCreditInvoiceGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ExtendedCreditInvoiceGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.ExtendedCreditInvoiceGroup]](GetExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditExtendedCreditInvoiceGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ExtendedCreditInvoiceGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ExtendedCreditInvoiceGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditExtendedCreditInvoiceGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditExtendedCreditInvoiceGroupsParams) (*core.Cursor[objects.ExtendedCreditInvoiceGroup], *core.Response, error) {
	var out core.Cursor[objects.ExtendedCreditInvoiceGroup]
	call := GetExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCreditExtendedCreditInvoiceGroups(ctx context.Context, client *core.Client, id string, params GetExtendedCreditExtendedCreditInvoiceGroupsParams) (*core.Cursor[objects.ExtendedCreditInvoiceGroup], error) {
	out, _, err := GetExtendedCreditExtendedCreditInvoiceGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateExtendedCreditExtendedCreditInvoiceGroupsParams struct {
	Emails []string    `facebook:"emails"`
	Name   string      `facebook:"name"`
	Extra  core.Params `facebook:"-"`
}

func (params CreateExtendedCreditExtendedCreditInvoiceGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["emails"] = params.Emails
	out["name"] = params.Name
	return out
}

func CreateExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id string, params CreateExtendedCreditExtendedCreditInvoiceGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "extended_credit_invoice_groups"), params.ToParams(), options...)
}

func NewCreateExtendedCreditExtendedCreditInvoiceGroupsBatchRequest(id string, params CreateExtendedCreditExtendedCreditInvoiceGroupsParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditInvoiceGroup] {
	return core.NewBatchRequest[objects.ExtendedCreditInvoiceGroup](CreateExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditExtendedCreditInvoiceGroupsBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditInvoiceGroup, error) {
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

func CreateExtendedCreditExtendedCreditInvoiceGroupsWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditExtendedCreditInvoiceGroupsParams) (*objects.ExtendedCreditInvoiceGroup, *core.Response, error) {
	var out objects.ExtendedCreditInvoiceGroup
	call := CreateExtendedCreditExtendedCreditInvoiceGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditExtendedCreditInvoiceGroups(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditExtendedCreditInvoiceGroupsParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	out, _, err := CreateExtendedCreditExtendedCreditInvoiceGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type GetExtendedCreditOwningCreditAllocationConfigsParams struct {
	ReceivingBusinessID *core.ID    `facebook:"receiving_business_id"`
	Extra               core.Params `facebook:"-"`
}

func (params GetExtendedCreditOwningCreditAllocationConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ReceivingBusinessID != nil {
		out["receiving_business_id"] = *params.ReceivingBusinessID
	}
	return out
}

func GetExtendedCreditOwningCreditAllocationConfigsBatchCall(id string, params GetExtendedCreditOwningCreditAllocationConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owning_credit_allocation_configs"), params.ToParams(), options...)
}

func NewGetExtendedCreditOwningCreditAllocationConfigsBatchRequest(id string, params GetExtendedCreditOwningCreditAllocationConfigsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ExtendedCreditAllocationConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.ExtendedCreditAllocationConfig]](GetExtendedCreditOwningCreditAllocationConfigsBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditOwningCreditAllocationConfigsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ExtendedCreditAllocationConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ExtendedCreditAllocationConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditOwningCreditAllocationConfigsWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditOwningCreditAllocationConfigsParams) (*core.Cursor[objects.ExtendedCreditAllocationConfig], *core.Response, error) {
	var out core.Cursor[objects.ExtendedCreditAllocationConfig]
	call := GetExtendedCreditOwningCreditAllocationConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCreditOwningCreditAllocationConfigs(ctx context.Context, client *core.Client, id string, params GetExtendedCreditOwningCreditAllocationConfigsParams) (*core.Cursor[objects.ExtendedCreditAllocationConfig], error) {
	out, _, err := GetExtendedCreditOwningCreditAllocationConfigsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateExtendedCreditOwningCreditAllocationConfigsParams struct {
	Amount              *map[string]interface{}                                                  `facebook:"amount"`
	LiabilityType       *enums.ExtendedcreditowningCreditAllocationConfigsLiabilityTypeEnumParam `facebook:"liability_type"`
	PartitionType       *enums.ExtendedcreditowningCreditAllocationConfigsPartitionTypeEnumParam `facebook:"partition_type"`
	ReceivingBusinessID core.ID                                                                  `facebook:"receiving_business_id"`
	SendBillTo          *enums.ExtendedcreditowningCreditAllocationConfigsSendBillToEnumParam    `facebook:"send_bill_to"`
	Extra               core.Params                                                              `facebook:"-"`
}

func (params CreateExtendedCreditOwningCreditAllocationConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Amount != nil {
		out["amount"] = *params.Amount
	}
	if params.LiabilityType != nil {
		out["liability_type"] = *params.LiabilityType
	}
	if params.PartitionType != nil {
		out["partition_type"] = *params.PartitionType
	}
	out["receiving_business_id"] = params.ReceivingBusinessID
	if params.SendBillTo != nil {
		out["send_bill_to"] = *params.SendBillTo
	}
	return out
}

func CreateExtendedCreditOwningCreditAllocationConfigsBatchCall(id string, params CreateExtendedCreditOwningCreditAllocationConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owning_credit_allocation_configs"), params.ToParams(), options...)
}

func NewCreateExtendedCreditOwningCreditAllocationConfigsBatchRequest(id string, params CreateExtendedCreditOwningCreditAllocationConfigsParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditAllocationConfig] {
	return core.NewBatchRequest[objects.ExtendedCreditAllocationConfig](CreateExtendedCreditOwningCreditAllocationConfigsBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditOwningCreditAllocationConfigsBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditAllocationConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditAllocationConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateExtendedCreditOwningCreditAllocationConfigsWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditOwningCreditAllocationConfigsParams) (*objects.ExtendedCreditAllocationConfig, *core.Response, error) {
	var out objects.ExtendedCreditAllocationConfig
	call := CreateExtendedCreditOwningCreditAllocationConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditOwningCreditAllocationConfigs(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditOwningCreditAllocationConfigsParams) (*objects.ExtendedCreditAllocationConfig, error) {
	out, _, err := CreateExtendedCreditOwningCreditAllocationConfigsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateExtendedCreditWhatsappCreditAttachParams struct {
	WabaCurrency string      `facebook:"waba_currency"`
	WabaID       core.ID     `facebook:"waba_id"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateExtendedCreditWhatsappCreditAttachParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["waba_currency"] = params.WabaCurrency
	out["waba_id"] = params.WabaID
	return out
}

func CreateExtendedCreditWhatsappCreditAttachBatchCall(id string, params CreateExtendedCreditWhatsappCreditAttachParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "whatsapp_credit_attach"), params.ToParams(), options...)
}

func NewCreateExtendedCreditWhatsappCreditAttachBatchRequest(id string, params CreateExtendedCreditWhatsappCreditAttachParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateExtendedCreditWhatsappCreditAttachBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditWhatsappCreditAttachBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateExtendedCreditWhatsappCreditAttachWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditAttachParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateExtendedCreditWhatsappCreditAttachBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditWhatsappCreditAttach(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditAttachParams) (*map[string]interface{}, error) {
	out, _, err := CreateExtendedCreditWhatsappCreditAttachWithResponse(ctx, client, id, params)
	return out, err
}

type CreateExtendedCreditWhatsappCreditSharingParams struct {
	ReceivingBusinessID core.ID     `facebook:"receiving_business_id"`
	Extra               core.Params `facebook:"-"`
}

func (params CreateExtendedCreditWhatsappCreditSharingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["receiving_business_id"] = params.ReceivingBusinessID
	return out
}

func CreateExtendedCreditWhatsappCreditSharingBatchCall(id string, params CreateExtendedCreditWhatsappCreditSharingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "whatsapp_credit_sharing"), params.ToParams(), options...)
}

func NewCreateExtendedCreditWhatsappCreditSharingBatchRequest(id string, params CreateExtendedCreditWhatsappCreditSharingParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateExtendedCreditWhatsappCreditSharingBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditWhatsappCreditSharingBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateExtendedCreditWhatsappCreditSharingWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateExtendedCreditWhatsappCreditSharingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditWhatsappCreditSharing(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingParams) (*map[string]interface{}, error) {
	out, _, err := CreateExtendedCreditWhatsappCreditSharingWithResponse(ctx, client, id, params)
	return out, err
}

type CreateExtendedCreditWhatsappCreditSharingAndAttachParams struct {
	WabaCurrency string      `facebook:"waba_currency"`
	WabaID       core.ID     `facebook:"waba_id"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateExtendedCreditWhatsappCreditSharingAndAttachParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["waba_currency"] = params.WabaCurrency
	out["waba_id"] = params.WabaID
	return out
}

func CreateExtendedCreditWhatsappCreditSharingAndAttachBatchCall(id string, params CreateExtendedCreditWhatsappCreditSharingAndAttachParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "whatsapp_credit_sharing_and_attach"), params.ToParams(), options...)
}

func NewCreateExtendedCreditWhatsappCreditSharingAndAttachBatchRequest(id string, params CreateExtendedCreditWhatsappCreditSharingAndAttachParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditAllocationConfig] {
	return core.NewBatchRequest[objects.ExtendedCreditAllocationConfig](CreateExtendedCreditWhatsappCreditSharingAndAttachBatchCall(id, params, options...))
}

func DecodeCreateExtendedCreditWhatsappCreditSharingAndAttachBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditAllocationConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditAllocationConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateExtendedCreditWhatsappCreditSharingAndAttachWithResponse(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingAndAttachParams) (*objects.ExtendedCreditAllocationConfig, *core.Response, error) {
	var out objects.ExtendedCreditAllocationConfig
	call := CreateExtendedCreditWhatsappCreditSharingAndAttachBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateExtendedCreditWhatsappCreditSharingAndAttach(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingAndAttachParams) (*objects.ExtendedCreditAllocationConfig, error) {
	out, _, err := CreateExtendedCreditWhatsappCreditSharingAndAttachWithResponse(ctx, client, id, params)
	return out, err
}

type GetExtendedCreditParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditBatchCall(id string, params GetExtendedCreditParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetExtendedCreditBatchRequest(id string, params GetExtendedCreditParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCredit] {
	return core.NewBatchRequest[objects.ExtendedCredit](GetExtendedCreditBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditBatchResponse(response *core.BatchResponse) (*objects.ExtendedCredit, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCredit
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditWithResponse(ctx context.Context, client *core.Client, id string, params GetExtendedCreditParams) (*objects.ExtendedCredit, *core.Response, error) {
	var out objects.ExtendedCredit
	call := GetExtendedCreditBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetExtendedCredit(ctx context.Context, client *core.Client, id string, params GetExtendedCreditParams) (*objects.ExtendedCredit, error) {
	out, _, err := GetExtendedCreditWithResponse(ctx, client, id, params)
	return out, err
}
