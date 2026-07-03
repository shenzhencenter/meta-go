package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppRequestFormerRecipient struct {
	ID          *core.ID `json:"id,omitempty"`
	RecipientID *core.ID `json:"recipient_id,omitempty"`
}
