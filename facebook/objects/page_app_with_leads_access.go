package objects

type PageAppWithLeadsAccess struct {
	CanAccessLeads *bool   `json:"can_access_leads,omitempty"`
	Type           *string `json:"type,omitempty"`
}
