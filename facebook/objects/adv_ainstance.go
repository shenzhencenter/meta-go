package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdvAInstance struct {
	ID            *core.ID  `json:"id,omitempty"`
	InstanceType  *string   `json:"instance_type,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OwnerBusiness *Business `json:"owner_business,omitempty"`
}
