package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdToplineDetail struct {
	ActiveStatus    *int       `json:"active_status,omitempty"`
	AdAccountID     *core.ID   `json:"ad_account_id,omitempty"`
	FlightEndDate   *core.Time `json:"flight_end_date,omitempty"`
	FlightStartDate *core.Time `json:"flight_start_date,omitempty"`
	ID              *core.ID   `json:"id,omitempty"`
	IoNumber        *int       `json:"io_number,omitempty"`
	LineNumber      *int       `json:"line_number,omitempty"`
	Price           *float64   `json:"price,omitempty"`
	Quantity        *float64   `json:"quantity,omitempty"`
	SfDetailLineID  *core.ID   `json:"sf_detail_line_id,omitempty"`
	SublineID       *core.ID   `json:"subline_id,omitempty"`
	Targets         *string    `json:"targets,omitempty"`
	TimeCreated     *core.Time `json:"time_created,omitempty"`
	TimeUpdated     *core.Time `json:"time_updated,omitempty"`
}

var AdToplineDetailFields = struct {
	ActiveStatus    string
	AdAccountID     string
	FlightEndDate   string
	FlightStartDate string
	ID              string
	IoNumber        string
	LineNumber      string
	Price           string
	Quantity        string
	SfDetailLineID  string
	SublineID       string
	Targets         string
	TimeCreated     string
	TimeUpdated     string
}{
	ActiveStatus:    "active_status",
	AdAccountID:     "ad_account_id",
	FlightEndDate:   "flight_end_date",
	FlightStartDate: "flight_start_date",
	ID:              "id",
	IoNumber:        "io_number",
	LineNumber:      "line_number",
	Price:           "price",
	Quantity:        "quantity",
	SfDetailLineID:  "sf_detail_line_id",
	SublineID:       "subline_id",
	Targets:         "targets",
	TimeCreated:     "time_created",
	TimeUpdated:     "time_updated",
}
