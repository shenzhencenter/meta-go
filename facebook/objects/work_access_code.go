package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WorkAccessCode struct {
	Code           *string    `json:"code,omitempty"`
	ExpirationTime *core.Time `json:"expiration_time,omitempty"`
}

var WorkAccessCodeFields = struct {
	Code           string
	ExpirationTime string
}{
	Code:           "code",
	ExpirationTime: "expiration_time",
}
