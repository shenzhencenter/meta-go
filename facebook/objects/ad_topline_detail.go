package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdToplineDetail struct {
	ActiveStatus    *int       `json:"active_status,omitempty"`
	AdAccountID     *core.ID   `json:"ad_account_id,omitempty"`
	FlightEndDate   *time.Time `json:"flight_end_date,omitempty"`
	FlightStartDate *time.Time `json:"flight_start_date,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	IoNumber        *int       `json:"io_number,omitempty"`
	LineNumber      *int       `json:"line_number,omitempty"`
	Price           *float64   `json:"price,omitempty"`
	Quantity        *float64   `json:"quantity,omitempty"`
	SfDetailLineID  *core.ID   `json:"sf_detail_line_id,omitempty"`
	SublineID       *core.ID   `json:"subline_id,omitempty"`
	Targets         *string    `json:"targets,omitempty"`
	TimeCreated     *time.Time `json:"time_created,omitempty"`
	TimeUpdated     *time.Time `json:"time_updated,omitempty"`
}
