package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountPaymentDetails struct {
	Amount           *CurrencyAmount         `json:"amount,omitempty"`
	CreateDate       *int                    `json:"create_date,omitempty"`
	ID               *core.ID                `json:"id,omitempty"`
	LastActionStatus *string                 `json:"last_action_status,omitempty"`
	Metadata         *map[string]interface{} `json:"metadata,omitempty"`
	PaymentDetailsID *core.ID                `json:"payment_details_id,omitempty"`
}

var AdAccountPaymentDetailsFields = struct {
	Amount           string
	CreateDate       string
	ID               string
	LastActionStatus string
	Metadata         string
	PaymentDetailsID string
}{
	Amount:           "amount",
	CreateDate:       "create_date",
	ID:               "id",
	LastActionStatus: "last_action_status",
	Metadata:         "metadata",
	PaymentDetailsID: "payment_details_id",
}
