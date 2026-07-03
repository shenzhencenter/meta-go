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

var ExtendedCreditFields = struct {
	AllocatedAmount                 string
	Balance                         string
	CreditAvailable                 string
	CreditType                      string
	ID                              string
	IsAccessRevoked                 string
	IsAutomatedExperience           string
	LegalEntityName                 string
	LiableAddress                   string
	LiableBizName                   string
	MaxBalance                      string
	OnlineMaxBalance                string
	OwnerBusiness                   string
	OwnerBusinessName               string
	PartitionFrom                   string
	ReceivingCreditAllocationConfig string
	SendBillToAddress               string
	SendBillToBizName               string
	SoldToAddress                   string
}{
	AllocatedAmount:                 "allocated_amount",
	Balance:                         "balance",
	CreditAvailable:                 "credit_available",
	CreditType:                      "credit_type",
	ID:                              "id",
	IsAccessRevoked:                 "is_access_revoked",
	IsAutomatedExperience:           "is_automated_experience",
	LegalEntityName:                 "legal_entity_name",
	LiableAddress:                   "liable_address",
	LiableBizName:                   "liable_biz_name",
	MaxBalance:                      "max_balance",
	OnlineMaxBalance:                "online_max_balance",
	OwnerBusiness:                   "owner_business",
	OwnerBusinessName:               "owner_business_name",
	PartitionFrom:                   "partition_from",
	ReceivingCreditAllocationConfig: "receiving_credit_allocation_config",
	SendBillToAddress:               "send_bill_to_address",
	SendBillToBizName:               "send_bill_to_biz_name",
	SoldToAddress:                   "sold_to_address",
}
