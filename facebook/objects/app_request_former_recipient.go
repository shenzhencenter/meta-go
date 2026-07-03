package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppRequestFormerRecipient struct {
	ID          *core.ID `json:"id,omitempty"`
	RecipientID *core.ID `json:"recipient_id,omitempty"`
}

var AppRequestFormerRecipientFields = struct {
	ID          string
	RecipientID string
}{
	ID:          "id",
	RecipientID: "recipient_id",
}
