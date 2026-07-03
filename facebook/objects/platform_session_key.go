package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PlatformSessionKey struct {
	ID *core.ID `json:"id,omitempty"`
}

var PlatformSessionKeyFields = struct {
	ID string
}{
	ID: "id",
}
