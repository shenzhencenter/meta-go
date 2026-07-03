package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ExtendedCreditEmail struct {
	Email *string  `json:"email,omitempty"`
	ID    *core.ID `json:"id,omitempty"`
}
