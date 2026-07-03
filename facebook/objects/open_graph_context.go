package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OpenGraphContext struct {
	ID *core.ID `json:"id,omitempty"`
}

var OpenGraphContextFields = struct {
	ID string
}{
	ID: "id",
}
