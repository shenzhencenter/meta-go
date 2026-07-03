package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountASLSchedule struct {
	AdAccount   *AdAccount `json:"ad_account,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	TimeCreated *core.Time `json:"time_created,omitempty"`
	TimeUpdated *core.Time `json:"time_updated,omitempty"`
}

var AdAccountASLScheduleFields = struct {
	AdAccount   string
	ID          string
	TimeCreated string
	TimeUpdated string
}{
	AdAccount:   "ad_account",
	ID:          "id",
	TimeCreated: "time_created",
	TimeUpdated: "time_updated",
}
