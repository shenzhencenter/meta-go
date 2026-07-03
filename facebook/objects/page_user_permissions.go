package objects

import (
	"time"
)

type PageUserPermissions struct {
	Business        *Business  `json:"business,omitempty"`
	BusinessPersona *string    `json:"business_persona,omitempty"`
	CreatedBy       *Profile   `json:"created_by,omitempty"`
	CreatedTime     *time.Time `json:"created_time,omitempty"`
	Email           *string    `json:"email,omitempty"`
	Status          *string    `json:"status,omitempty"`
	Tasks           *[]string  `json:"tasks,omitempty"`
	UpdatedBy       *Profile   `json:"updated_by,omitempty"`
	UpdatedTime     *time.Time `json:"updated_time,omitempty"`
	User            *Profile   `json:"user,omitempty"`
}
