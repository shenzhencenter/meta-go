package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CommercePayout struct {
	Amount            *map[string]interface{} `json:"amount,omitempty"`
	PayoutDate        *string                 `json:"payout_date,omitempty"`
	PayoutReferenceID *core.ID                `json:"payout_reference_id,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	TransferID        *core.ID                `json:"transfer_id,omitempty"`
}
