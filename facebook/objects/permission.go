package objects

type Permission struct {
	Permission *string `json:"permission,omitempty"`
	Status     *string `json:"status,omitempty"`
}

var PermissionFields = struct {
	Permission string
	Status     string
}{
	Permission: "permission",
	Status:     "status",
}
