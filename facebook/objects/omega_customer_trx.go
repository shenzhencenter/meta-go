package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
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
	DueDate             *time.Time              `json:"due_date,omitempty"`
	Entity              *string                 `json:"entity,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	InvoiceDate         *time.Time              `json:"invoice_date,omitempty"`
	InvoiceID           *core.ID                `json:"invoice_id,omitempty"`
	InvoiceType         *string                 `json:"invoice_type,omitempty"`
	LiabilityType       *string                 `json:"liability_type,omitempty"`
	PaymentStatus       *string                 `json:"payment_status,omitempty"`
	PaymentTerm         *string                 `json:"payment_term,omitempty"`
	Type                *string                 `json:"type,omitempty"`
}
