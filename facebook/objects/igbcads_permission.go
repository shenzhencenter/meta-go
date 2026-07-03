package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type IGBCAdsPermission struct {
	ID             *core.ID `json:"id,omitempty"`
	PermissionType *string  `json:"permission_type,omitempty"`
	Status         *string  `json:"status,omitempty"`
}
