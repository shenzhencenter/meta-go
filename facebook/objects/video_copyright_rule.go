package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type VideoCopyrightRule struct {
	ConditionGroups *[]VideoCopyrightConditionGroup `json:"condition_groups,omitempty"`
	Copyrights      *[]string                       `json:"copyrights,omitempty"`
	CreatedDate     *time.Time                      `json:"created_date,omitempty"`
	Creator         *User                           `json:"creator,omitempty"`
	ID              *core.ID                        `json:"id,omitempty"`
	IsInMigration   *bool                           `json:"is_in_migration,omitempty"`
	Name            *string                         `json:"name,omitempty"`
	ValidityStatus  *string                         `json:"validity_status,omitempty"`
}
