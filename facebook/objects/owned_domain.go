package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type OwnedDomain struct {
	DomainName       *string   `json:"domain_name,omitempty"`
	ID               *core.ID  `json:"id,omitempty"`
	OwnerBusiness    *Business `json:"owner_business,omitempty"`
	Status           *string   `json:"status,omitempty"`
	VerificationCode *string   `json:"verification_code,omitempty"`
}
