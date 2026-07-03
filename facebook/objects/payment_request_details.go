package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PaymentRequestDetails struct {
	Amount           *map[string]interface{} `json:"amount,omitempty"`
	CreationTime     *int                    `json:"creation_time,omitempty"`
	Note             *string                 `json:"note,omitempty"`
	PaymentRequestID *core.ID                `json:"payment_request_id,omitempty"`
	ReceiverID       *core.ID                `json:"receiver_id,omitempty"`
	ReferenceNumber  *string                 `json:"reference_number,omitempty"`
	SenderID         *core.ID                `json:"sender_id,omitempty"`
	Status           *string                 `json:"status,omitempty"`
	TransactionTime  *int                    `json:"transaction_time,omitempty"`
}

var PaymentRequestDetailsFields = struct {
	Amount           string
	CreationTime     string
	Note             string
	PaymentRequestID string
	ReceiverID       string
	ReferenceNumber  string
	SenderID         string
	Status           string
	TransactionTime  string
}{
	Amount:           "amount",
	CreationTime:     "creation_time",
	Note:             "note",
	PaymentRequestID: "payment_request_id",
	ReceiverID:       "receiver_id",
	ReferenceNumber:  "reference_number",
	SenderID:         "sender_id",
	Status:           "status",
	TransactionTime:  "transaction_time",
}
