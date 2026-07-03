package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountSpendLimit struct {
	AmountSpent *string  `json:"amount_spent,omitempty"`
	GroupID     *core.ID `json:"group_id,omitempty"`
	LimitID     *core.ID `json:"limit_id,omitempty"`
	LimitValue  *string  `json:"limit_value,omitempty"`
	TimeCreated *uint64  `json:"time_created,omitempty"`
	TimeStart   *uint64  `json:"time_start,omitempty"`
	TimeStop    *uint64  `json:"time_stop,omitempty"`
}
