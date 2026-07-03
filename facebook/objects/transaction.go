package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Transaction struct {
	AccountID               *core.ID                      `json:"account_id,omitempty"`
	AppAmount               *map[string]interface{}       `json:"app_amount,omitempty"`
	BillingEndTime          *uint64                       `json:"billing_end_time,omitempty"`
	BillingReason           *string                       `json:"billing_reason,omitempty"`
	BillingStartTime        *uint64                       `json:"billing_start_time,omitempty"`
	CardChargeMode          *uint64                       `json:"card_charge_mode,omitempty"`
	ChargeType              *string                       `json:"charge_type,omitempty"`
	CheckoutCampaignGroupID *core.ID                      `json:"checkout_campaign_group_id,omitempty"`
	CredentialID            *core.ID                      `json:"credential_id,omitempty"`
	FaturaID                *core.ID                      `json:"fatura_id,omitempty"`
	ID                      *core.ID                      `json:"id,omitempty"`
	IsBusinessEcCharge      *bool                         `json:"is_business_ec_charge,omitempty"`
	IsFundingEvent          *bool                         `json:"is_funding_event,omitempty"`
	PaymentOption           *string                       `json:"payment_option,omitempty"`
	ProductType             *enums.TransactionProductType `json:"product_type,omitempty"`
	ProviderAmount          *map[string]interface{}       `json:"provider_amount,omitempty"`
	Status                  *string                       `json:"status,omitempty"`
	Time                    *uint64                       `json:"time,omitempty"`
	TrackingID              *core.ID                      `json:"tracking_id,omitempty"`
	TransactionType         *string                       `json:"transaction_type,omitempty"`
	TxType                  *uint64                       `json:"tx_type,omitempty"`
	VatInvoiceID            *core.ID                      `json:"vat_invoice_id,omitempty"`
}

var TransactionFields = struct {
	AccountID               string
	AppAmount               string
	BillingEndTime          string
	BillingReason           string
	BillingStartTime        string
	CardChargeMode          string
	ChargeType              string
	CheckoutCampaignGroupID string
	CredentialID            string
	FaturaID                string
	ID                      string
	IsBusinessEcCharge      string
	IsFundingEvent          string
	PaymentOption           string
	ProductType             string
	ProviderAmount          string
	Status                  string
	Time                    string
	TrackingID              string
	TransactionType         string
	TxType                  string
	VatInvoiceID            string
}{
	AccountID:               "account_id",
	AppAmount:               "app_amount",
	BillingEndTime:          "billing_end_time",
	BillingReason:           "billing_reason",
	BillingStartTime:        "billing_start_time",
	CardChargeMode:          "card_charge_mode",
	ChargeType:              "charge_type",
	CheckoutCampaignGroupID: "checkout_campaign_group_id",
	CredentialID:            "credential_id",
	FaturaID:                "fatura_id",
	ID:                      "id",
	IsBusinessEcCharge:      "is_business_ec_charge",
	IsFundingEvent:          "is_funding_event",
	PaymentOption:           "payment_option",
	ProductType:             "product_type",
	ProviderAmount:          "provider_amount",
	Status:                  "status",
	Time:                    "time",
	TrackingID:              "tracking_id",
	TransactionType:         "transaction_type",
	TxType:                  "tx_type",
	VatInvoiceID:            "vat_invoice_id",
}
