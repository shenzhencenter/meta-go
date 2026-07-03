package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type OfflineTermsOfService struct {
	AcceptTime   *int     `json:"accept_time,omitempty"`
	ID           *core.ID `json:"id,omitempty"`
	SignedByUser *User    `json:"signed_by_user,omitempty"`
}

var OfflineTermsOfServiceFields = struct {
	AcceptTime   string
	ID           string
	SignedByUser string
}{
	AcceptTime:   "accept_time",
	ID:           "id",
	SignedByUser: "signed_by_user",
}
