package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdAccountASLSchedule struct {
	AdAccount   *AdAccount `json:"ad_account,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	TimeCreated *time.Time `json:"time_created,omitempty"`
	TimeUpdated *time.Time `json:"time_updated,omitempty"`
}
