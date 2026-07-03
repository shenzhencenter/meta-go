package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Avatar struct {
	ID *core.ID `json:"id,omitempty"`
}

var AvatarFields = struct {
	ID string
}{
	ID: "id",
}
