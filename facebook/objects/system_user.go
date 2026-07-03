package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type SystemUser struct {
	CreatedBy         *User      `json:"created_by,omitempty"`
	CreatedTime       *time.Time `json:"created_time,omitempty"`
	FinancePermission *string    `json:"finance_permission,omitempty"`
	ID                *core.ID   `json:"id,omitempty"`
	IPPermission      *string    `json:"ip_permission,omitempty"`
	Name              *string    `json:"name,omitempty"`
}
