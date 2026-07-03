package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Organization struct {
	ID              *core.ID  `json:"id,omitempty"`
	LegalEntityName *string   `json:"legal_entity_name,omitempty"`
	OwnerBusiness   *Business `json:"owner_business,omitempty"`
}

var OrganizationFields = struct {
	ID              string
	LegalEntityName string
	OwnerBusiness   string
}{
	ID:              "id",
	LegalEntityName: "legal_entity_name",
	OwnerBusiness:   "owner_business",
}
