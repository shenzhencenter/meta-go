package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CommerceOrderTransactionDetail struct {
	MerchantOrderID   *core.ID                `json:"merchant_order_id,omitempty"`
	NetPaymentAmount  *map[string]interface{} `json:"net_payment_amount,omitempty"`
	OrderCreated      *string                 `json:"order_created,omitempty"`
	OrderDetails      *CommerceOrder          `json:"order_details,omitempty"`
	OrderID           *core.ID                `json:"order_id,omitempty"`
	PayoutReferenceID *core.ID                `json:"payout_reference_id,omitempty"`
	PostalCode        *string                 `json:"postal_code,omitempty"`
	ProcessingFee     *map[string]interface{} `json:"processing_fee,omitempty"`
	State             *string                 `json:"state,omitempty"`
	TaxRate           *string                 `json:"tax_rate,omitempty"`
	TransactionDate   *string                 `json:"transaction_date,omitempty"`
	TransactionType   *string                 `json:"transaction_type,omitempty"`
	TransferID        *core.ID                `json:"transfer_id,omitempty"`
}
