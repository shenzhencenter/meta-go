package objects

type PageAppWithLeadsAccess struct {
	CanAccessLeads *bool   `json:"can_access_leads,omitempty"`
	Type           *string `json:"type,omitempty"`
}

var PageAppWithLeadsAccessFields = struct {
	CanAccessLeads string
	Type           string
}{
	CanAccessLeads: "can_access_leads",
	Type:           "type",
}
