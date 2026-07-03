package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Group struct {
	Archived           *bool                   `json:"archived,omitempty"`
	Cover              *CoverPhoto             `json:"cover,omitempty"`
	CreatedTime        *core.Time              `json:"created_time,omitempty"`
	Description        *string                 `json:"description,omitempty"`
	Email              *string                 `json:"email,omitempty"`
	Icon               *string                 `json:"icon,omitempty"`
	ID                 *core.ID                `json:"id,omitempty"`
	Install            *map[string]interface{} `json:"install,omitempty"`
	Link               *string                 `json:"link,omitempty"`
	MemberCount        *uint64                 `json:"member_count,omitempty"`
	MemberRequestCount *uint64                 `json:"member_request_count,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Parent             *map[string]interface{} `json:"parent,omitempty"`
	Permissions        *[]string               `json:"permissions,omitempty"`
	Privacy            *string                 `json:"privacy,omitempty"`
	Purpose            *string                 `json:"purpose,omitempty"`
	Subdomain          *string                 `json:"subdomain,omitempty"`
	UpdatedTime        *core.Time              `json:"updated_time,omitempty"`
	Venue              *Location               `json:"venue,omitempty"`
}

var GroupFields = struct {
	Archived           string
	Cover              string
	CreatedTime        string
	Description        string
	Email              string
	Icon               string
	ID                 string
	Install            string
	Link               string
	MemberCount        string
	MemberRequestCount string
	Name               string
	Parent             string
	Permissions        string
	Privacy            string
	Purpose            string
	Subdomain          string
	UpdatedTime        string
	Venue              string
}{
	Archived:           "archived",
	Cover:              "cover",
	CreatedTime:        "created_time",
	Description:        "description",
	Email:              "email",
	Icon:               "icon",
	ID:                 "id",
	Install:            "install",
	Link:               "link",
	MemberCount:        "member_count",
	MemberRequestCount: "member_request_count",
	Name:               "name",
	Parent:             "parent",
	Permissions:        "permissions",
	Privacy:            "privacy",
	Purpose:            "purpose",
	Subdomain:          "subdomain",
	UpdatedTime:        "updated_time",
	Venue:              "venue",
}
