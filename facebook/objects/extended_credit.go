package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExtendedCredit struct {
	AllocatedAmount                 *CurrencyAmount                 `json:"allocated_amount,omitempty"`
	Balance                         *CurrencyAmount                 `json:"balance,omitempty"`
	CreditAvailable                 *CurrencyAmount                 `json:"credit_available,omitempty"`
	CreditType                      *string                         `json:"credit_type,omitempty"`
	ID                              *core.ID                        `json:"id,omitempty"`
	IsAccessRevoked                 *bool                           `json:"is_access_revoked,omitempty"`
	IsAutomatedExperience           *bool                           `json:"is_automated_experience,omitempty"`
	LegalEntityName                 *string                         `json:"legal_entity_name,omitempty"`
	LiableAddress                   *CRMAddress                     `json:"liable_address,omitempty"`
	LiableBizName                   *string                         `json:"liable_biz_name,omitempty"`
	MaxBalance                      *CurrencyAmount                 `json:"max_balance,omitempty"`
	OnlineMaxBalance                *CurrencyAmount                 `json:"online_max_balance,omitempty"`
	OwnerBusiness                   *Business                       `json:"owner_business,omitempty"`
	OwnerBusinessName               *string                         `json:"owner_business_name,omitempty"`
	PartitionFrom                   *string                         `json:"partition_from,omitempty"`
	ReceivingCreditAllocationConfig *ExtendedCreditAllocationConfig `json:"receiving_credit_allocation_config,omitempty"`
	SendBillToAddress               *CRMAddress                     `json:"send_bill_to_address,omitempty"`
	SendBillToBizName               *string                         `json:"send_bill_to_biz_name,omitempty"`
	SoldToAddress                   *CRMAddress                     `json:"sold_to_address,omitempty"`
}
