package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type FundraiserPersonToCharity struct {
	AmountRaised           *int       `json:"amount_raised,omitempty"`
	CharityID              *core.ID   `json:"charity_id,omitempty"`
	Currency               *string    `json:"currency,omitempty"`
	Description            *string    `json:"description,omitempty"`
	DonationsCount         *int       `json:"donations_count,omitempty"`
	DonorsCount            *int       `json:"donors_count,omitempty"`
	EndTime                *time.Time `json:"end_time,omitempty"`
	ExternalAmountRaised   *int       `json:"external_amount_raised,omitempty"`
	ExternalDonationsCount *int       `json:"external_donations_count,omitempty"`
	ExternalDonorsCount    *int       `json:"external_donors_count,omitempty"`
	ExternalEventName      *string    `json:"external_event_name,omitempty"`
	ExternalEventStartTime *time.Time `json:"external_event_start_time,omitempty"`
	ExternalEventURI       *string    `json:"external_event_uri,omitempty"`
	ExternalFundraiserURI  *string    `json:"external_fundraiser_uri,omitempty"`
	ExternalID             *core.ID   `json:"external_id,omitempty"`
	GoalAmount             *int       `json:"goal_amount,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	InternalAmountRaised   *int       `json:"internal_amount_raised,omitempty"`
	InternalDonationsCount *int       `json:"internal_donations_count,omitempty"`
	InternalDonorsCount    *int       `json:"internal_donors_count,omitempty"`
	Name                   *string    `json:"name,omitempty"`
	URI                    *string    `json:"uri,omitempty"`
}
