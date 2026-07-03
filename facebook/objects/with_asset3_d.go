package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WithAsset3D struct {
	ID *core.ID `json:"id,omitempty"`
}

var WithAsset3DFields = struct {
	ID string
}{
	ID: "id",
}
