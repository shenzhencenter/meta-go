package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoCopyrightRule struct {
	ConditionGroups *[]VideoCopyrightConditionGroup `json:"condition_groups,omitempty"`
	Copyrights      *[]string                       `json:"copyrights,omitempty"`
	CreatedDate     *core.Time                      `json:"created_date,omitempty"`
	Creator         *User                           `json:"creator,omitempty"`
	ID              *core.ID                        `json:"id,omitempty"`
	IsInMigration   *bool                           `json:"is_in_migration,omitempty"`
	Name            *string                         `json:"name,omitempty"`
	ValidityStatus  *string                         `json:"validity_status,omitempty"`
}

var VideoCopyrightRuleFields = struct {
	ConditionGroups string
	Copyrights      string
	CreatedDate     string
	Creator         string
	ID              string
	IsInMigration   string
	Name            string
	ValidityStatus  string
}{
	ConditionGroups: "condition_groups",
	Copyrights:      "copyrights",
	CreatedDate:     "created_date",
	Creator:         "creator",
	ID:              "id",
	IsInMigration:   "is_in_migration",
	Name:            "name",
	ValidityStatus:  "validity_status",
}
