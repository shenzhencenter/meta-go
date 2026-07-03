package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdAccountBillingDatePreference struct {
	AdAccount     *AdAccount `json:"ad_account,omitempty"`
	DayOfMonth    *int       `json:"day_of_month,omitempty"`
	ID            *core.ID   `json:"id,omitempty"`
	NextBillDate  *time.Time `json:"next_bill_date,omitempty"`
	TimeCreated   *time.Time `json:"time_created,omitempty"`
	TimeEffective *time.Time `json:"time_effective,omitempty"`
}
