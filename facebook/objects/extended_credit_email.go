package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExtendedCreditEmail struct {
	Email *string  `json:"email,omitempty"`
	ID    *core.ID `json:"id,omitempty"`
}
