package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AppRequestFormerRecipient struct {
	ID          *core.ID `json:"id,omitempty"`
	RecipientID *core.ID `json:"recipient_id,omitempty"`
}
