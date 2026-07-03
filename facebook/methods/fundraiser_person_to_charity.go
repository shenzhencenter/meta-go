package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetFundraiserPersonToCharityDonationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFundraiserPersonToCharityDonationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFundraiserPersonToCharityDonations(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityDonationsParams) (*core.Cursor[objects.InternalDonationForApp], error) {
	var out core.Cursor[objects.InternalDonationForApp]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "donations"), params.ToParams(), &out)
	return &out, err
}

type CreateFundraiserPersonToCharityEndFundraiserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateFundraiserPersonToCharityEndFundraiserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateFundraiserPersonToCharityEndFundraiser(ctx context.Context, client *core.Client, id string, params CreateFundraiserPersonToCharityEndFundraiserParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "end_fundraiser"), params.ToParams(), &out)
	return &out, err
}

type GetFundraiserPersonToCharityExternalDonationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFundraiserPersonToCharityExternalDonationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFundraiserPersonToCharityExternalDonations(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityExternalDonationsParams) (*core.Cursor[objects.ExternalAppDonation], error) {
	var out core.Cursor[objects.ExternalAppDonation]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "external_donations"), params.ToParams(), &out)
	return &out, err
}

type CreateFundraiserPersonToCharityExternalDonationsParams struct {
	AmountReceived uint64      `facebook:"amount_received"`
	Currency       string      `facebook:"currency"`
	DonationIDHash string      `facebook:"donation_id_hash"`
	DonationTime   uint64      `facebook:"donation_time"`
	DonorIDHash    string      `facebook:"donor_id_hash"`
	Extra          core.Params `facebook:"-"`
}

func (params CreateFundraiserPersonToCharityExternalDonationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["amount_received"] = params.AmountReceived
	out["currency"] = params.Currency
	out["donation_id_hash"] = params.DonationIDHash
	out["donation_time"] = params.DonationTime
	out["donor_id_hash"] = params.DonorIDHash
	return out
}

func CreateFundraiserPersonToCharityExternalDonations(ctx context.Context, client *core.Client, id string, params CreateFundraiserPersonToCharityExternalDonationsParams) (*objects.ExternalAppDonation, error) {
	var out objects.ExternalAppDonation
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "external_donations"), params.ToParams(), &out)
	return &out, err
}

type GetFundraiserPersonToCharityParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFundraiserPersonToCharityParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFundraiserPersonToCharity(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityParams) (*objects.FundraiserPersonToCharity, error) {
	var out objects.FundraiserPersonToCharity
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateFundraiserPersonToCharityParams struct {
	Description            *string     `facebook:"description"`
	EndTime                *time.Time  `facebook:"end_time"`
	ExternalEventName      *string     `facebook:"external_event_name"`
	ExternalEventStartTime *time.Time  `facebook:"external_event_start_time"`
	ExternalEventURI       *string     `facebook:"external_event_uri"`
	ExternalFundraiserURI  *string     `facebook:"external_fundraiser_uri"`
	ExternalID             *core.ID    `facebook:"external_id"`
	GoalAmount             *uint64     `facebook:"goal_amount"`
	Name                   *string     `facebook:"name"`
	Extra                  core.Params `facebook:"-"`
}

func (params UpdateFundraiserPersonToCharityParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.ExternalEventName != nil {
		out["external_event_name"] = *params.ExternalEventName
	}
	if params.ExternalEventStartTime != nil {
		out["external_event_start_time"] = *params.ExternalEventStartTime
	}
	if params.ExternalEventURI != nil {
		out["external_event_uri"] = *params.ExternalEventURI
	}
	if params.ExternalFundraiserURI != nil {
		out["external_fundraiser_uri"] = *params.ExternalFundraiserURI
	}
	if params.ExternalID != nil {
		out["external_id"] = *params.ExternalID
	}
	if params.GoalAmount != nil {
		out["goal_amount"] = *params.GoalAmount
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateFundraiserPersonToCharity(ctx context.Context, client *core.Client, id string, params UpdateFundraiserPersonToCharityParams) (*objects.FundraiserPersonToCharity, error) {
	var out objects.FundraiserPersonToCharity
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
