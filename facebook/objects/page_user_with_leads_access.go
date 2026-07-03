package objects

type PageUserWithLeadsAccess struct {
	ActiveOnBusiness *bool   `json:"active_on_business,omitempty"`
	BusinessRole     *string `json:"business_role,omitempty"`
	CanAccessLeads   *bool   `json:"can_access_leads,omitempty"`
	PageRole         *string `json:"page_role,omitempty"`
}

var PageUserWithLeadsAccessFields = struct {
	ActiveOnBusiness string
	BusinessRole     string
	CanAccessLeads   string
	PageRole         string
}{
	ActiveOnBusiness: "active_on_business",
	BusinessRole:     "business_role",
	CanAccessLeads:   "can_access_leads",
	PageRole:         "page_role",
}
