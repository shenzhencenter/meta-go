package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessObjectTransferOwnershipAgreement struct {
	ID                 *core.ID  `json:"id,omitempty"`
	ReceivingBusiness  *Business `json:"receiving_business,omitempty"`
	RequestingBusiness *Business `json:"requesting_business,omitempty"`
	Status             *string   `json:"status,omitempty"`
}

var BusinessObjectTransferOwnershipAgreementFields = struct {
	ID                 string
	ReceivingBusiness  string
	RequestingBusiness string
	Status             string
}{
	ID:                 "id",
	ReceivingBusiness:  "receiving_business",
	RequestingBusiness: "requesting_business",
	Status:             "status",
}
