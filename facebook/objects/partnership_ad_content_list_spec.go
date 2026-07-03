package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PartnershipAdContentListSpec struct {
	ListID *core.ID `json:"list_id,omitempty"`
}

var PartnershipAdContentListSpecFields = struct {
	ListID string
}{
	ListID: "list_id",
}
