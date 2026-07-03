package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdsPaymentCycle struct {
	AccountID                *core.ID   `json:"account_id,omitempty"`
	CreatedTime              *time.Time `json:"created_time,omitempty"`
	Multiplier               *uint64    `json:"multiplier,omitempty"`
	RequestedThresholdAmount *uint64    `json:"requested_threshold_amount,omitempty"`
	ThresholdAmount          *uint64    `json:"threshold_amount,omitempty"`
	UpdatedTime              *time.Time `json:"updated_time,omitempty"`
}
