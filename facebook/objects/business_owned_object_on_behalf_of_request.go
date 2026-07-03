package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessOwnedObjectOnBehalfOfRequest struct {
	BusinessOwnedObject *string   `json:"business_owned_object,omitempty"`
	ID                  *core.ID  `json:"id,omitempty"`
	ReceivingBusiness   *Business `json:"receiving_business,omitempty"`
	RequestingBusiness  *Business `json:"requesting_business,omitempty"`
	Status              *string   `json:"status,omitempty"`
}
