package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPaymentCycle struct {
	AccountID                *core.ID   `json:"account_id,omitempty"`
	CreatedTime              *core.Time `json:"created_time,omitempty"`
	Multiplier               *uint64    `json:"multiplier,omitempty"`
	RequestedThresholdAmount *uint64    `json:"requested_threshold_amount,omitempty"`
	ThresholdAmount          *uint64    `json:"threshold_amount,omitempty"`
	UpdatedTime              *core.Time `json:"updated_time,omitempty"`
}

var AdsPaymentCycleFields = struct {
	AccountID                string
	CreatedTime              string
	Multiplier               string
	RequestedThresholdAmount string
	ThresholdAmount          string
	UpdatedTime              string
}{
	AccountID:                "account_id",
	CreatedTime:              "created_time",
	Multiplier:               "multiplier",
	RequestedThresholdAmount: "requested_threshold_amount",
	ThresholdAmount:          "threshold_amount",
	UpdatedTime:              "updated_time",
}
