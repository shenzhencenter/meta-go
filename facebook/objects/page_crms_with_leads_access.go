package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageCrmsWithLeadsAccess struct {
	CanAccessLeads  *bool    `json:"can_access_leads,omitempty"`
	ID              *core.ID `json:"id,omitempty"`
	IntegrationType *string  `json:"integration_type,omitempty"`
	Name            *string  `json:"name,omitempty"`
}
