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

func GetFundraiserPersonToCharityDonationsBatchCall(id string, params GetFundraiserPersonToCharityDonationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "donations"), params.ToParams(), options...)
}

func NewGetFundraiserPersonToCharityDonationsBatchRequest(id string, params GetFundraiserPersonToCharityDonationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InternalDonationForApp]] {
	return core.NewBatchRequest[core.Cursor[objects.InternalDonationForApp]](GetFundraiserPersonToCharityDonationsBatchCall(id, params, options...))
}

func DecodeGetFundraiserPersonToCharityDonationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InternalDonationForApp], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InternalDonationForApp]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFundraiserPersonToCharityDonations(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityDonationsParams) (*core.Cursor[objects.InternalDonationForApp], error) {
	var out core.Cursor[objects.InternalDonationForApp]
	call := GetFundraiserPersonToCharityDonationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateFundraiserPersonToCharityEndFundraiserBatchCall(id string, params CreateFundraiserPersonToCharityEndFundraiserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "end_fundraiser"), params.ToParams(), options...)
}

func NewCreateFundraiserPersonToCharityEndFundraiserBatchRequest(id string, params CreateFundraiserPersonToCharityEndFundraiserParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateFundraiserPersonToCharityEndFundraiserBatchCall(id, params, options...))
}

func DecodeCreateFundraiserPersonToCharityEndFundraiserBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateFundraiserPersonToCharityEndFundraiser(ctx context.Context, client *core.Client, id string, params CreateFundraiserPersonToCharityEndFundraiserParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateFundraiserPersonToCharityEndFundraiserBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetFundraiserPersonToCharityExternalDonationsBatchCall(id string, params GetFundraiserPersonToCharityExternalDonationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "external_donations"), params.ToParams(), options...)
}

func NewGetFundraiserPersonToCharityExternalDonationsBatchRequest(id string, params GetFundraiserPersonToCharityExternalDonationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ExternalAppDonation]] {
	return core.NewBatchRequest[core.Cursor[objects.ExternalAppDonation]](GetFundraiserPersonToCharityExternalDonationsBatchCall(id, params, options...))
}

func DecodeGetFundraiserPersonToCharityExternalDonationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ExternalAppDonation], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ExternalAppDonation]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFundraiserPersonToCharityExternalDonations(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityExternalDonationsParams) (*core.Cursor[objects.ExternalAppDonation], error) {
	var out core.Cursor[objects.ExternalAppDonation]
	call := GetFundraiserPersonToCharityExternalDonationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateFundraiserPersonToCharityExternalDonationsBatchCall(id string, params CreateFundraiserPersonToCharityExternalDonationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "external_donations"), params.ToParams(), options...)
}

func NewCreateFundraiserPersonToCharityExternalDonationsBatchRequest(id string, params CreateFundraiserPersonToCharityExternalDonationsParams, options ...core.BatchOption) *core.BatchRequest[objects.ExternalAppDonation] {
	return core.NewBatchRequest[objects.ExternalAppDonation](CreateFundraiserPersonToCharityExternalDonationsBatchCall(id, params, options...))
}

func DecodeCreateFundraiserPersonToCharityExternalDonationsBatchResponse(response *core.BatchResponse) (*objects.ExternalAppDonation, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExternalAppDonation
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateFundraiserPersonToCharityExternalDonations(ctx context.Context, client *core.Client, id string, params CreateFundraiserPersonToCharityExternalDonationsParams) (*objects.ExternalAppDonation, error) {
	var out objects.ExternalAppDonation
	call := CreateFundraiserPersonToCharityExternalDonationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetFundraiserPersonToCharityBatchCall(id string, params GetFundraiserPersonToCharityParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFundraiserPersonToCharityBatchRequest(id string, params GetFundraiserPersonToCharityParams, options ...core.BatchOption) *core.BatchRequest[objects.FundraiserPersonToCharity] {
	return core.NewBatchRequest[objects.FundraiserPersonToCharity](GetFundraiserPersonToCharityBatchCall(id, params, options...))
}

func DecodeGetFundraiserPersonToCharityBatchResponse(response *core.BatchResponse) (*objects.FundraiserPersonToCharity, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FundraiserPersonToCharity
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFundraiserPersonToCharity(ctx context.Context, client *core.Client, id string, params GetFundraiserPersonToCharityParams) (*objects.FundraiserPersonToCharity, error) {
	var out objects.FundraiserPersonToCharity
	call := GetFundraiserPersonToCharityBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateFundraiserPersonToCharityBatchCall(id string, params UpdateFundraiserPersonToCharityParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateFundraiserPersonToCharityBatchRequest(id string, params UpdateFundraiserPersonToCharityParams, options ...core.BatchOption) *core.BatchRequest[objects.FundraiserPersonToCharity] {
	return core.NewBatchRequest[objects.FundraiserPersonToCharity](UpdateFundraiserPersonToCharityBatchCall(id, params, options...))
}

func DecodeUpdateFundraiserPersonToCharityBatchResponse(response *core.BatchResponse) (*objects.FundraiserPersonToCharity, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FundraiserPersonToCharity
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateFundraiserPersonToCharity(ctx context.Context, client *core.Client, id string, params UpdateFundraiserPersonToCharityParams) (*objects.FundraiserPersonToCharity, error) {
	var out objects.FundraiserPersonToCharity
	call := UpdateFundraiserPersonToCharityBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
