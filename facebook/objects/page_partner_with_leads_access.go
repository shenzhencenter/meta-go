package objects

type PagePartnerWithLeadsAccess struct {
	CanAccessLeads  *bool     `json:"can_access_leads,omitempty"`
	PartnerBusiness *Business `json:"partner_business,omitempty"`
	PermittedTasks  *[]string `json:"permitted_tasks,omitempty"`
}

var PagePartnerWithLeadsAccessFields = struct {
	CanAccessLeads  string
	PartnerBusiness string
	PermittedTasks  string
}{
	CanAccessLeads:  "can_access_leads",
	PartnerBusiness: "partner_business",
	PermittedTasks:  "permitted_tasks",
}
