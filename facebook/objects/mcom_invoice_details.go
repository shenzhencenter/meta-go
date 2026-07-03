package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type McomInvoiceDetails struct {
	AdditionalAmounts           *[]map[string]interface{} `json:"additional_amounts,omitempty"`
	BuyerNotes                  *string                   `json:"buyer_notes,omitempty"`
	CurrencyAmount              *map[string]interface{}   `json:"currency_amount,omitempty"`
	ExternalInvoiceID           *core.ID                  `json:"external_invoice_id,omitempty"`
	Features                    *map[string]interface{}   `json:"features,omitempty"`
	InvoiceCreated              *int                      `json:"invoice_created,omitempty"`
	InvoiceID                   *core.ID                  `json:"invoice_id,omitempty"`
	InvoiceInstructions         *string                   `json:"invoice_instructions,omitempty"`
	InvoiceInstructionsImageURL *string                   `json:"invoice_instructions_image_url,omitempty"`
	InvoiceUpdated              *int                      `json:"invoice_updated,omitempty"`
	OutstandingAmount           *map[string]interface{}   `json:"outstanding_amount,omitempty"`
	PaidAmount                  *map[string]interface{}   `json:"paid_amount,omitempty"`
	Payments                    *[]map[string]interface{} `json:"payments,omitempty"`
	PlatformLogoURL             *string                   `json:"platform_logo_url,omitempty"`
	PlatformName                *string                   `json:"platform_name,omitempty"`
	ProductItems                *[]map[string]interface{} `json:"product_items,omitempty"`
	ShippingAddress             *map[string]interface{}   `json:"shipping_address,omitempty"`
	Status                      *string                   `json:"status,omitempty"`
	TrackingInfo                *map[string]interface{}   `json:"tracking_info,omitempty"`
}
