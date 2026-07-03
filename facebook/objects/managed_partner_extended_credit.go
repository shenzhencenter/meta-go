package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ManagedPartnerExtendedCredit struct {
	ID                              *core.ID                        `json:"id,omitempty"`
	MaxBalance                      *CurrencyAmount                 `json:"max_balance,omitempty"`
	ReceivingCreditAllocationConfig *ExtendedCreditAllocationConfig `json:"receiving_credit_allocation_config,omitempty"`
}
