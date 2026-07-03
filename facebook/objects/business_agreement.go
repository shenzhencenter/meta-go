package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAgreement struct {
	ID            *core.ID `json:"id,omitempty"`
	RequestStatus *string  `json:"request_status,omitempty"`
}

var BusinessAgreementFields = struct {
	ID            string
	RequestStatus string
}{
	ID:            "id",
	RequestStatus: "request_status",
}
