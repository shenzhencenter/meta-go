package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Organization struct {
	ID              *core.ID  `json:"id,omitempty"`
	LegalEntityName *string   `json:"legal_entity_name,omitempty"`
	OwnerBusiness   *Business `json:"owner_business,omitempty"`
}
