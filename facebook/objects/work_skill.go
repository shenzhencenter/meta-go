package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WorkSkill struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}

var WorkSkillFields = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}
