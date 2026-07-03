package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageDirectIntegrationCrmWithLeadsAccess struct {
	CanAccessLeads *bool    `json:"can_access_leads,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
}

var PageDirectIntegrationCrmWithLeadsAccessFields = struct {
	CanAccessLeads string
	ID             string
	Name           string
}{
	CanAccessLeads: "can_access_leads",
	ID:             "id",
	Name:           "name",
}
