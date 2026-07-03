package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAssetGroup struct {
	ID            *core.ID  `json:"id,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OwnerBusiness *Business `json:"owner_business,omitempty"`
}
