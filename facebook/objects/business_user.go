package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BusinessUser struct {
	Business            *Business            `json:"business,omitempty"`
	BusinessRoleRequest *BusinessRoleRequest `json:"business_role_request,omitempty"`
	Email               *string              `json:"email,omitempty"`
	FinancePermission   *string              `json:"finance_permission,omitempty"`
	FirstName           *string              `json:"first_name,omitempty"`
	ID                  *core.ID             `json:"id,omitempty"`
	IPPermission        *string              `json:"ip_permission,omitempty"`
	LastName            *string              `json:"last_name,omitempty"`
	MarkedForRemoval    *bool                `json:"marked_for_removal,omitempty"`
	Name                *string              `json:"name,omitempty"`
	PendingEmail        *string              `json:"pending_email,omitempty"`
	Role                *string              `json:"role,omitempty"`
	Tasks               *[]string            `json:"tasks,omitempty"`
	Title               *string              `json:"title,omitempty"`
	TwoFacStatus        *string              `json:"two_fac_status,omitempty"`
}
