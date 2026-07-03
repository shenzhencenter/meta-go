package objects

type IPObject struct {
	IPPermission *string                 `json:"ip_permission,omitempty"`
	User         *map[string]interface{} `json:"user,omitempty"`
}
