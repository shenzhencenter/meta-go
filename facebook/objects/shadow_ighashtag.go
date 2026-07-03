package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ShadowIGHashtag struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}

var ShadowIGHashtagFields = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}
