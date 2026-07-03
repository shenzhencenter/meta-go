package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountBillingDatePreference struct {
	AdAccount     *AdAccount `json:"ad_account,omitempty"`
	DayOfMonth    *int       `json:"day_of_month,omitempty"`
	ID            *core.ID   `json:"id,omitempty"`
	NextBillDate  *core.Time `json:"next_bill_date,omitempty"`
	TimeCreated   *core.Time `json:"time_created,omitempty"`
	TimeEffective *core.Time `json:"time_effective,omitempty"`
}

var AdAccountBillingDatePreferenceFields = struct {
	AdAccount     string
	DayOfMonth    string
	ID            string
	NextBillDate  string
	TimeCreated   string
	TimeEffective string
}{
	AdAccount:     "ad_account",
	DayOfMonth:    "day_of_month",
	ID:            "id",
	NextBillDate:  "next_bill_date",
	TimeCreated:   "time_created",
	TimeEffective: "time_effective",
}
