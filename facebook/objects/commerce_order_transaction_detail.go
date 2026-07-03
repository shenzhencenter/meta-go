package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CommerceOrderTransactionDetailFields = struct {
	MerchantOrderID   string
	NetPaymentAmount  string
	OrderCreated      string
	OrderDetails      string
	OrderID           string
	PayoutReferenceID string
	PostalCode        string
	ProcessingFee     string
	State             string
	TaxRate           string
	TransactionDate   string
	TransactionType   string
	TransferID        string
}{
	MerchantOrderID:   "merchant_order_id",
	NetPaymentAmount:  "net_payment_amount",
	OrderCreated:      "order_created",
	OrderDetails:      "order_details",
	OrderID:           "order_id",
	PayoutReferenceID: "payout_reference_id",
	PostalCode:        "postal_code",
	ProcessingFee:     "processing_fee",
	State:             "state",
	TaxRate:           "tax_rate",
	TransactionDate:   "transaction_date",
	TransactionType:   "transaction_type",
	TransferID:        "transfer_id",
}
