package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PersonalAdsPersona struct {
	Email        *string  `json:"email,omitempty"`
	FirstName    *string  `json:"first_name,omitempty"`
	ID           *core.ID `json:"id,omitempty"`
	LastName     *string  `json:"last_name,omitempty"`
	PendingEmail *string  `json:"pending_email,omitempty"`
}

var PersonalAdsPersonaFields = struct {
	Email        string
	FirstName    string
	ID           string
	LastName     string
	PendingEmail string
}{
	Email:        "email",
	FirstName:    "first_name",
	ID:           "id",
	LastName:     "last_name",
	PendingEmail: "pending_email",
}
