package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAccountPaymentDetails struct {
	Amount           *CurrencyAmount         `json:"amount,omitempty"`
	CreateDate       *int                    `json:"create_date,omitempty"`
	ID               *core.ID                `json:"id,omitempty"`
	LastActionStatus *string                 `json:"last_action_status,omitempty"`
	Metadata         *map[string]interface{} `json:"metadata,omitempty"`
	PaymentDetailsID *core.ID                `json:"payment_details_id,omitempty"`
}
