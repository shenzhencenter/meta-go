package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type McomInvoiceStatus struct {
	BankAccountNumber      *string                 `json:"bank_account_number,omitempty"`
	BankCode               *string                 `json:"bank_code,omitempty"`
	InvoiceID              *core.ID                `json:"invoice_id,omitempty"`
	InvoiceStatus          *string                 `json:"invoice_status,omitempty"`
	PageID                 *core.ID                `json:"page_id,omitempty"`
	PaymentMethod          *string                 `json:"payment_method,omitempty"`
	PaymentType            *string                 `json:"payment_type,omitempty"`
	PayoutAmount           *map[string]interface{} `json:"payout_amount,omitempty"`
	SlipVerificationError  *string                 `json:"slip_verification_error,omitempty"`
	SlipVerificationStatus *string                 `json:"slip_verification_status,omitempty"`
	SofTransferID          *core.ID                `json:"sof_transfer_id,omitempty"`
	SofTransferTimestamp   *int                    `json:"sof_transfer_timestamp,omitempty"`
	TransactionFee         *map[string]interface{} `json:"transaction_fee,omitempty"`
	TransferSlip           *string                 `json:"transfer_slip,omitempty"`
	TransferSlipQrCode     *string                 `json:"transfer_slip_qr_code,omitempty"`
}
