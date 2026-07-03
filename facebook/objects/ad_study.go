package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdStudy struct {
	Business                  *Business  `json:"business,omitempty"`
	CanceledTime              *time.Time `json:"canceled_time,omitempty"`
	ClientBusiness            *Business  `json:"client_business,omitempty"`
	CooldownStartTime         *time.Time `json:"cooldown_start_time,omitempty"`
	CreatedBy                 *User      `json:"created_by,omitempty"`
	CreatedTime               *time.Time `json:"created_time,omitempty"`
	Description               *string    `json:"description,omitempty"`
	EndTime                   *time.Time `json:"end_time,omitempty"`
	ID                        *core.ID   `json:"id,omitempty"`
	MeasurementContact        *User      `json:"measurement_contact,omitempty"`
	Name                      *string    `json:"name,omitempty"`
	ObservationEndTime        *time.Time `json:"observation_end_time,omitempty"`
	ResultsFirstAvailableDate *string    `json:"results_first_available_date,omitempty"`
	SalesContact              *User      `json:"sales_contact,omitempty"`
	StartTime                 *time.Time `json:"start_time,omitempty"`
	Type                      *string    `json:"type,omitempty"`
	UpdatedBy                 *User      `json:"updated_by,omitempty"`
	UpdatedTime               *time.Time `json:"updated_time,omitempty"`
}
