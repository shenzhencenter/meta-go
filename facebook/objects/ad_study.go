package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudy struct {
	Business                  *Business  `json:"business,omitempty"`
	CanceledTime              *core.Time `json:"canceled_time,omitempty"`
	ClientBusiness            *Business  `json:"client_business,omitempty"`
	CooldownStartTime         *core.Time `json:"cooldown_start_time,omitempty"`
	CreatedBy                 *User      `json:"created_by,omitempty"`
	CreatedTime               *core.Time `json:"created_time,omitempty"`
	Description               *string    `json:"description,omitempty"`
	EndTime                   *core.Time `json:"end_time,omitempty"`
	ID                        *core.ID   `json:"id,omitempty"`
	MeasurementContact        *User      `json:"measurement_contact,omitempty"`
	Name                      *string    `json:"name,omitempty"`
	ObservationEndTime        *core.Time `json:"observation_end_time,omitempty"`
	ResultsFirstAvailableDate *string    `json:"results_first_available_date,omitempty"`
	SalesContact              *User      `json:"sales_contact,omitempty"`
	StartTime                 *core.Time `json:"start_time,omitempty"`
	Type                      *string    `json:"type,omitempty"`
	UpdatedBy                 *User      `json:"updated_by,omitempty"`
	UpdatedTime               *core.Time `json:"updated_time,omitempty"`
}

var AdStudyFields = struct {
	Business                  string
	CanceledTime              string
	ClientBusiness            string
	CooldownStartTime         string
	CreatedBy                 string
	CreatedTime               string
	Description               string
	EndTime                   string
	ID                        string
	MeasurementContact        string
	Name                      string
	ObservationEndTime        string
	ResultsFirstAvailableDate string
	SalesContact              string
	StartTime                 string
	Type                      string
	UpdatedBy                 string
	UpdatedTime               string
}{
	Business:                  "business",
	CanceledTime:              "canceled_time",
	ClientBusiness:            "client_business",
	CooldownStartTime:         "cooldown_start_time",
	CreatedBy:                 "created_by",
	CreatedTime:               "created_time",
	Description:               "description",
	EndTime:                   "end_time",
	ID:                        "id",
	MeasurementContact:        "measurement_contact",
	Name:                      "name",
	ObservationEndTime:        "observation_end_time",
	ResultsFirstAvailableDate: "results_first_available_date",
	SalesContact:              "sales_contact",
	StartTime:                 "start_time",
	Type:                      "type",
	UpdatedBy:                 "updated_by",
	UpdatedTime:               "updated_time",
}
