package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PlayableContent struct {
	ID    *core.ID `json:"id,omitempty"`
	Name  *string  `json:"name,omitempty"`
	Owner *Profile `json:"owner,omitempty"`
}

var PlayableContentFields = struct {
	ID    string
	Name  string
	Owner string
}{
	ID:    "id",
	Name:  "name",
	Owner: "owner",
}
