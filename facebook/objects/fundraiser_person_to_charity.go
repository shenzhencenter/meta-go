package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FundraiserPersonToCharity struct {
	AmountRaised           *int       `json:"amount_raised,omitempty"`
	CharityID              *core.ID   `json:"charity_id,omitempty"`
	Currency               *string    `json:"currency,omitempty"`
	Description            *string    `json:"description,omitempty"`
	DonationsCount         *int       `json:"donations_count,omitempty"`
	DonorsCount            *int       `json:"donors_count,omitempty"`
	EndTime                *core.Time `json:"end_time,omitempty"`
	ExternalAmountRaised   *int       `json:"external_amount_raised,omitempty"`
	ExternalDonationsCount *int       `json:"external_donations_count,omitempty"`
	ExternalDonorsCount    *int       `json:"external_donors_count,omitempty"`
	ExternalEventName      *string    `json:"external_event_name,omitempty"`
	ExternalEventStartTime *core.Time `json:"external_event_start_time,omitempty"`
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

var FundraiserPersonToCharityFields = struct {
	AmountRaised           string
	CharityID              string
	Currency               string
	Description            string
	DonationsCount         string
	DonorsCount            string
	EndTime                string
	ExternalAmountRaised   string
	ExternalDonationsCount string
	ExternalDonorsCount    string
	ExternalEventName      string
	ExternalEventStartTime string
	ExternalEventURI       string
	ExternalFundraiserURI  string
	ExternalID             string
	GoalAmount             string
	ID                     string
	InternalAmountRaised   string
	InternalDonationsCount string
	InternalDonorsCount    string
	Name                   string
	URI                    string
}{
	AmountRaised:           "amount_raised",
	CharityID:              "charity_id",
	Currency:               "currency",
	Description:            "description",
	DonationsCount:         "donations_count",
	DonorsCount:            "donors_count",
	EndTime:                "end_time",
	ExternalAmountRaised:   "external_amount_raised",
	ExternalDonationsCount: "external_donations_count",
	ExternalDonorsCount:    "external_donors_count",
	ExternalEventName:      "external_event_name",
	ExternalEventStartTime: "external_event_start_time",
	ExternalEventURI:       "external_event_uri",
	ExternalFundraiserURI:  "external_fundraiser_uri",
	ExternalID:             "external_id",
	GoalAmount:             "goal_amount",
	ID:                     "id",
	InternalAmountRaised:   "internal_amount_raised",
	InternalDonationsCount: "internal_donations_count",
	InternalDonorsCount:    "internal_donors_count",
	Name:                   "name",
	URI:                    "uri",
}
