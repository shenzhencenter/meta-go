package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BusinessAgreement struct {
	ID            *core.ID `json:"id,omitempty"`
	RequestStatus *string  `json:"request_status,omitempty"`
}
