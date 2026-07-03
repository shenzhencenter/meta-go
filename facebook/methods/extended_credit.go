package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetExtendedCreditExtendedCreditInvoiceGroups(ctx context.Context, client *core.Client, id string, params GetExtendedCreditExtendedCreditInvoiceGroupsParams) (*core.Cursor[objects.ExtendedCreditInvoiceGroup], error) {
	var out core.Cursor[objects.ExtendedCreditInvoiceGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "extended_credit_invoice_groups"), params.ToParams(), &out)
	return &out, err
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

func CreateExtendedCreditExtendedCreditInvoiceGroups(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditExtendedCreditInvoiceGroupsParams) (*objects.ExtendedCreditInvoiceGroup, error) {
	var out objects.ExtendedCreditInvoiceGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "extended_credit_invoice_groups"), params.ToParams(), &out)
	return &out, err
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

func GetExtendedCreditOwningCreditAllocationConfigs(ctx context.Context, client *core.Client, id string, params GetExtendedCreditOwningCreditAllocationConfigsParams) (*core.Cursor[objects.ExtendedCreditAllocationConfig], error) {
	var out core.Cursor[objects.ExtendedCreditAllocationConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owning_credit_allocation_configs"), params.ToParams(), &out)
	return &out, err
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

func CreateExtendedCreditOwningCreditAllocationConfigs(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditOwningCreditAllocationConfigsParams) (*objects.ExtendedCreditAllocationConfig, error) {
	var out objects.ExtendedCreditAllocationConfig
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owning_credit_allocation_configs"), params.ToParams(), &out)
	return &out, err
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

func CreateExtendedCreditWhatsappCreditAttach(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditAttachParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "whatsapp_credit_attach"), params.ToParams(), &out)
	return &out, err
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

func CreateExtendedCreditWhatsappCreditSharing(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "whatsapp_credit_sharing"), params.ToParams(), &out)
	return &out, err
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

func CreateExtendedCreditWhatsappCreditSharingAndAttach(ctx context.Context, client *core.Client, id string, params CreateExtendedCreditWhatsappCreditSharingAndAttachParams) (*objects.ExtendedCreditAllocationConfig, error) {
	var out objects.ExtendedCreditAllocationConfig
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "whatsapp_credit_sharing_and_attach"), params.ToParams(), &out)
	return &out, err
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

func GetExtendedCredit(ctx context.Context, client *core.Client, id string, params GetExtendedCreditParams) (*objects.ExtendedCredit, error) {
	var out objects.ExtendedCredit
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
