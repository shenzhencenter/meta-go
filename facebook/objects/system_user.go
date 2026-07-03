package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type SystemUser struct {
	CreatedBy         *User      `json:"created_by,omitempty"`
	CreatedTime       *core.Time `json:"created_time,omitempty"`
	FinancePermission *string    `json:"finance_permission,omitempty"`
	ID                *core.ID   `json:"id,omitempty"`
	IPPermission      *string    `json:"ip_permission,omitempty"`
	Name              *string    `json:"name,omitempty"`
}

var SystemUserFields = struct {
	CreatedBy         string
	CreatedTime       string
	FinancePermission string
	ID                string
	IPPermission      string
	Name              string
}{
	CreatedBy:         "created_by",
	CreatedTime:       "created_time",
	FinancePermission: "finance_permission",
	ID:                "id",
	IPPermission:      "ip_permission",
	Name:              "name",
}
