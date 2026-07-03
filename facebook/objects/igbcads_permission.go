package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGBCAdsPermission struct {
	ID             *core.ID `json:"id,omitempty"`
	PermissionType *string  `json:"permission_type,omitempty"`
	Status         *string  `json:"status,omitempty"`
}

var IGBCAdsPermissionFields = struct {
	ID             string
	PermissionType string
	Status         string
}{
	ID:             "id",
	PermissionType: "permission_type",
	Status:         "status",
}
