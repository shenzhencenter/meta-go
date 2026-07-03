package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageDirectIntegrationCrmWithLeadsAccess struct {
	CanAccessLeads *bool    `json:"can_access_leads,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
}
