package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountUser struct {
	ID    *core.ID  `json:"id,omitempty"`
	Name  *string   `json:"name,omitempty"`
	Tasks *[]string `json:"tasks,omitempty"`
}

var AdAccountUserFields = struct {
	ID    string
	Name  string
	Tasks string
}{
	ID:    "id",
	Name:  "name",
	Tasks: "tasks",
}
