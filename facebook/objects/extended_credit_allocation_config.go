package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ExtendedCreditAllocationConfig struct {
	CurrencyAmount      *CurrencyAmount `json:"currency_amount,omitempty"`
	ID                  *core.ID        `json:"id,omitempty"`
	LiabilityType       *string         `json:"liability_type,omitempty"`
	OwningBusiness      *Business       `json:"owning_business,omitempty"`
	OwningCredential    *ExtendedCredit `json:"owning_credential,omitempty"`
	PartitionType       *string         `json:"partition_type,omitempty"`
	ReceivingBusiness   *Business       `json:"receiving_business,omitempty"`
	ReceivingCredential *ExtendedCredit `json:"receiving_credential,omitempty"`
	RequestStatus       *string         `json:"request_status,omitempty"`
	SendBillTo          *string         `json:"send_bill_to,omitempty"`
}
