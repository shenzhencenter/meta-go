package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserContext struct {
	ID *core.ID `json:"id,omitempty"`
}

var UserContextFields = struct {
	ID string
}{
	ID: "id",
}
