package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdvAInstance struct {
	ID            *core.ID  `json:"id,omitempty"`
	InstanceType  *string   `json:"instance_type,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OwnerBusiness *Business `json:"owner_business,omitempty"`
}

var AdvAInstanceFields = struct {
	ID            string
	InstanceType  string
	Name          string
	OwnerBusiness string
}{
	ID:            "id",
	InstanceType:  "instance_type",
	Name:          "name",
	OwnerBusiness: "owner_business",
}
