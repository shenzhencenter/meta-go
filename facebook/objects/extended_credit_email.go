package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExtendedCreditEmail struct {
	Email *string  `json:"email,omitempty"`
	ID    *core.ID `json:"id,omitempty"`
}

var ExtendedCreditEmailFields = struct {
	Email string
	ID    string
}{
	Email: "email",
	ID:    "id",
}
