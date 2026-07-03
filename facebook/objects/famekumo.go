package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FAMEKumo struct {
	ID *core.ID `json:"id,omitempty"`
}

var FAMEKumoFields = struct {
	ID string
}{
	ID: "id",
}
