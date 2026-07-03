package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessRoleRequest struct {
	CreatedBy       *map[string]interface{} `json:"created_by,omitempty"`
	CreatedTime     *core.Time              `json:"created_time,omitempty"`
	Email           *string                 `json:"email,omitempty"`
	ExpirationTime  *core.Time              `json:"expiration_time,omitempty"`
	ExpiryTime      *core.Time              `json:"expiry_time,omitempty"`
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
	UpdatedTime     *core.Time              `json:"updated_time,omitempty"`
}

var BusinessRoleRequestFields = struct {
	CreatedBy       string
	CreatedTime     string
	Email           string
	ExpirationTime  string
	ExpiryTime      string
	FinanceRole     string
	ID              string
	InviteLink      string
	InvitedUserType string
	IPRole          string
	Owner           string
	Role            string
	Status          string
	Tasks           string
	UpdatedBy       string
	UpdatedTime     string
}{
	CreatedBy:       "created_by",
	CreatedTime:     "created_time",
	Email:           "email",
	ExpirationTime:  "expiration_time",
	ExpiryTime:      "expiry_time",
	FinanceRole:     "finance_role",
	ID:              "id",
	InviteLink:      "invite_link",
	InvitedUserType: "invited_user_type",
	IPRole:          "ip_role",
	Owner:           "owner",
	Role:            "role",
	Status:          "status",
	Tasks:           "tasks",
	UpdatedBy:       "updated_by",
	UpdatedTime:     "updated_time",
}
