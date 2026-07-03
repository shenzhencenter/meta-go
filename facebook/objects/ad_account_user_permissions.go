package objects

import (
	"time"
)

type AdAccountUserPermissions struct {
	Business        *Business               `json:"business,omitempty"`
	BusinessPersona *map[string]interface{} `json:"business_persona,omitempty"`
	CreatedBy       *User                   `json:"created_by,omitempty"`
	CreatedTime     *time.Time              `json:"created_time,omitempty"`
	Email           *string                 `json:"email,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Tasks           *[]string               `json:"tasks,omitempty"`
	UpdatedBy       *User                   `json:"updated_by,omitempty"`
	UpdatedTime     *time.Time              `json:"updated_time,omitempty"`
	User            *User                   `json:"user,omitempty"`
}
