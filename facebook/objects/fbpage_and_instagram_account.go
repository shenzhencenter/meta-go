package objects

type FBPageAndInstagramAccount struct {
	AdPermissions      *[]string            `json:"ad_permissions,omitempty"`
	BcPermissionStatus *string              `json:"bc_permission_status,omitempty"`
	BcPermissions      *[]map[string]string `json:"bc_permissions,omitempty"`
	IsManaged          *bool                `json:"is_managed,omitempty"`
	MatchedBy          *string              `json:"matched_by,omitempty"`
}
