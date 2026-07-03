package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type BusinessRoleRequest struct {
	CreatedBy       *map[string]interface{} `json:"created_by,omitempty"`
	CreatedTime     *time.Time              `json:"created_time,omitempty"`
	Email           *string                 `json:"email,omitempty"`
	ExpirationTime  *time.Time              `json:"expiration_time,omitempty"`
	ExpiryTime      *time.Time              `json:"expiry_time,omitempty"`
	FinanceRole     *string                 `json:"finance_role,omitempty"`
	ID              *core.ID                `json:"id,omitempty"`
	InviteLink      *string                 `json:"invite_link,omitempty"`
	InvitedUserType *[]string               `json:"invited_user_type,omitempty"`
	IPRole          *string                 `json:"ip_role,omitempty"`
	Owner           *Business               `json:"owner,omitempty"`
	Role            *string                 `json:"role,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Tasks           *[]string               `json:"tasks,omitempty"`
	UpdatedBy       *map[string]interface{} `json:"updated_by,omitempty"`
	UpdatedTime     *time.Time              `json:"updated_time,omitempty"`
}
