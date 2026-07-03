package objects

type PageUserWithLeadsAccess struct {
	ActiveOnBusiness *bool   `json:"active_on_business,omitempty"`
	BusinessRole     *string `json:"business_role,omitempty"`
	CanAccessLeads   *bool   `json:"can_access_leads,omitempty"`
	PageRole         *string `json:"page_role,omitempty"`
}
