package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type Group struct {
	Archived           *bool                   `json:"archived,omitempty"`
	Cover              *CoverPhoto             `json:"cover,omitempty"`
	CreatedTime        *time.Time              `json:"created_time,omitempty"`
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
	UpdatedTime        *time.Time              `json:"updated_time,omitempty"`
	Venue              *Location               `json:"venue,omitempty"`
}
