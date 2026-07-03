package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IDName struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}

var IDNameFields = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}
