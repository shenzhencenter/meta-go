package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var ExtendedCreditAllocationConfigFields = struct {
	CurrencyAmount      string
	ID                  string
	LiabilityType       string
	OwningBusiness      string
	OwningCredential    string
	PartitionType       string
	ReceivingBusiness   string
	ReceivingCredential string
	RequestStatus       string
	SendBillTo          string
}{
	CurrencyAmount:      "currency_amount",
	ID:                  "id",
	LiabilityType:       "liability_type",
	OwningBusiness:      "owning_business",
	OwningCredential:    "owning_credential",
	PartitionType:       "partition_type",
	ReceivingBusiness:   "receiving_business",
	ReceivingCredential: "receiving_credential",
	RequestStatus:       "request_status",
	SendBillTo:          "send_bill_to",
}
