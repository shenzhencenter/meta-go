package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OmegaCustomerTrx struct {
	AdAccountIds        *[]core.ID              `json:"ad_account_ids,omitempty"`
	AdvertiserName      *string                 `json:"advertiser_name,omitempty"`
	Amount              *string                 `json:"amount,omitempty"`
	AmountDue           *CurrencyAmount         `json:"amount_due,omitempty"`
	BilledAmountDetails *map[string]interface{} `json:"billed_amount_details,omitempty"`
	BillingPeriod       *string                 `json:"billing_period,omitempty"`
	CdnDownloadURI      *string                 `json:"cdn_download_uri,omitempty"`
	Currency            *string                 `json:"currency,omitempty"`
	DownloadURI         *string                 `json:"download_uri,omitempty"`
	DueDate             *core.Time              `json:"due_date,omitempty"`
	Entity              *string                 `json:"entity,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	InvoiceDate         *core.Time              `json:"invoice_date,omitempty"`
	InvoiceID           *core.ID                `json:"invoice_id,omitempty"`
	InvoiceType         *string                 `json:"invoice_type,omitempty"`
	LiabilityType       *string                 `json:"liability_type,omitempty"`
	PaymentStatus       *string                 `json:"payment_status,omitempty"`
	PaymentTerm         *string                 `json:"payment_term,omitempty"`
	Type                *string                 `json:"type,omitempty"`
}

var OmegaCustomerTrxFields = struct {
	AdAccountIds        string
	AdvertiserName      string
	Amount              string
	AmountDue           string
	BilledAmountDetails string
	BillingPeriod       string
	CdnDownloadURI      string
	Currency            string
	DownloadURI         string
	DueDate             string
	Entity              string
	ID                  string
	InvoiceDate         string
	InvoiceID           string
	InvoiceType         string
	LiabilityType       string
	PaymentStatus       string
	PaymentTerm         string
	Type                string
}{
	AdAccountIds:        "ad_account_ids",
	AdvertiserName:      "advertiser_name",
	Amount:              "amount",
	AmountDue:           "amount_due",
	BilledAmountDetails: "billed_amount_details",
	BillingPeriod:       "billing_period",
	CdnDownloadURI:      "cdn_download_uri",
	Currency:            "currency",
	DownloadURI:         "download_uri",
	DueDate:             "due_date",
	Entity:              "entity",
	ID:                  "id",
	InvoiceDate:         "invoice_date",
	InvoiceID:           "invoice_id",
	InvoiceType:         "invoice_type",
	LiabilityType:       "liability_type",
	PaymentStatus:       "payment_status",
	PaymentTerm:         "payment_term",
	Type:                "type",
}
