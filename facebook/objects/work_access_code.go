package objects

import (
	"time"
)

type WorkAccessCode struct {
	Code           *string    `json:"code,omitempty"`
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
}
