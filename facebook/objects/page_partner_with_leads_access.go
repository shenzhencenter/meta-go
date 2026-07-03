package objects

type PagePartnerWithLeadsAccess struct {
	CanAccessLeads  *bool     `json:"can_access_leads,omitempty"`
	PartnerBusiness *Business `json:"partner_business,omitempty"`
	PermittedTasks  *[]string `json:"permitted_tasks,omitempty"`
}
