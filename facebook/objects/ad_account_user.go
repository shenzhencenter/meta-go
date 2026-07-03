package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountUser struct {
	ID    *core.ID  `json:"id,omitempty"`
	Name  *string   `json:"name,omitempty"`
	Tasks *[]string `json:"tasks,omitempty"`
}
