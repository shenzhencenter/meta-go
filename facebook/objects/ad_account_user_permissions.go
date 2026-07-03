package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountUserPermissions struct {
	Business        *Business               `json:"business,omitempty"`
	BusinessPersona *map[string]interface{} `json:"business_persona,omitempty"`
	CreatedBy       *User                   `json:"created_by,omitempty"`
	CreatedTime     *core.Time              `json:"created_time,omitempty"`
	Email           *string                 `json:"email,omitempty"`
	Status          *string                 `json:"status,omitempty"`
	Tasks           *[]string               `json:"tasks,omitempty"`
	UpdatedBy       *User                   `json:"updated_by,omitempty"`
	UpdatedTime     *core.Time              `json:"updated_time,omitempty"`
	User            *User                   `json:"user,omitempty"`
}

var AdAccountUserPermissionsFields = struct {
	Business        string
	BusinessPersona string
	CreatedBy       string
	CreatedTime     string
	Email           string
	Status          string
	Tasks           string
	UpdatedBy       string
	UpdatedTime     string
	User            string
}{
	Business:        "business",
	BusinessPersona: "business_persona",
	CreatedBy:       "created_by",
	CreatedTime:     "created_time",
	Email:           "email",
	Status:          "status",
	Tasks:           "tasks",
	UpdatedBy:       "updated_by",
	UpdatedTime:     "updated_time",
	User:            "user",
}
