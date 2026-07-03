package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OwnedDomain struct {
	DomainName       *string   `json:"domain_name,omitempty"`
	ID               *core.ID  `json:"id,omitempty"`
	OwnerBusiness    *Business `json:"owner_business,omitempty"`
	Status           *string   `json:"status,omitempty"`
	VerificationCode *string   `json:"verification_code,omitempty"`
}

var OwnedDomainFields = struct {
	DomainName       string
	ID               string
	OwnerBusiness    string
	Status           string
	VerificationCode string
}{
	DomainName:       "domain_name",
	ID:               "id",
	OwnerBusiness:    "owner_business",
	Status:           "status",
	VerificationCode: "verification_code",
}
