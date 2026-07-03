package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LoomConfig struct {
	ID *core.ID `json:"id,omitempty"`
}

var LoomConfigFields = struct {
	ID string
}{
	ID: "id",
}
