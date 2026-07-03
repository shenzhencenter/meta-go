package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdStudyCellAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellAdaccountsBatchCall(id string, params GetAdStudyCellAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetAdStudyCellAdaccountsBatchRequest(id string, params GetAdStudyCellAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetAdStudyCellAdaccountsBatchCall(id, params, options...))
}

func DecodeGetAdStudyCellAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyCellAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdStudyCellAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetAdStudyCellAdaccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyCellAdsetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellAdsetsBatchCall(id string, params GetAdStudyCellAdsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), options...)
}

func NewGetAdStudyCellAdsetsBatchRequest(id string, params GetAdStudyCellAdsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdStudyCellAdsetsBatchCall(id, params, options...))
}

func DecodeGetAdStudyCellAdsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyCellAdsets(ctx context.Context, client *core.Client, id string, params GetAdStudyCellAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdStudyCellAdsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyCellCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellCampaignsBatchCall(id string, params GetAdStudyCellCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewGetAdStudyCellCampaignsBatchRequest(id string, params GetAdStudyCellCampaignsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Campaign]] {
	return core.NewBatchRequest[core.Cursor[objects.Campaign]](GetAdStudyCellCampaignsBatchCall(id, params, options...))
}

func DecodeGetAdStudyCellCampaignsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Campaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Campaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyCellCampaigns(ctx context.Context, client *core.Client, id string, params GetAdStudyCellCampaignsParams) (*core.Cursor[objects.Campaign], error) {
	var out core.Cursor[objects.Campaign]
	call := GetAdStudyCellCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdStudyCellParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdStudyCellParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdStudyCellBatchCall(id string, params GetAdStudyCellParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdStudyCellBatchRequest(id string, params GetAdStudyCellParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudyCell] {
	return core.NewBatchRequest[objects.AdStudyCell](GetAdStudyCellBatchCall(id, params, options...))
}

func DecodeGetAdStudyCellBatchResponse(response *core.BatchResponse) (*objects.AdStudyCell, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudyCell
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdStudyCell(ctx context.Context, client *core.Client, id string, params GetAdStudyCellParams) (*objects.AdStudyCell, error) {
	var out objects.AdStudyCell
	call := GetAdStudyCellBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateAdStudyCellParams struct {
	Adaccounts       *[]uint64                          `facebook:"adaccounts"`
	Ads              *[]string                          `facebook:"ads"`
	Adsets           *[]string                          `facebook:"adsets"`
	Campaigns        *[]string                          `facebook:"campaigns"`
	CreationTemplate *enums.AdstudycellCreationTemplate `facebook:"creation_template"`
	Description      *string                            `facebook:"description"`
	Name             *string                            `facebook:"name"`
	Extra            core.Params                        `facebook:"-"`
}

func (params UpdateAdStudyCellParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adaccounts != nil {
		out["adaccounts"] = *params.Adaccounts
	}
	if params.Ads != nil {
		out["ads"] = *params.Ads
	}
	if params.Adsets != nil {
		out["adsets"] = *params.Adsets
	}
	if params.Campaigns != nil {
		out["campaigns"] = *params.Campaigns
	}
	if params.CreationTemplate != nil {
		out["creation_template"] = *params.CreationTemplate
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateAdStudyCellBatchCall(id string, params UpdateAdStudyCellParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdStudyCellBatchRequest(id string, params UpdateAdStudyCellParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudyCell] {
	return core.NewBatchRequest[objects.AdStudyCell](UpdateAdStudyCellBatchCall(id, params, options...))
}

func DecodeUpdateAdStudyCellBatchResponse(response *core.BatchResponse) (*objects.AdStudyCell, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudyCell
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdStudyCell(ctx context.Context, client *core.Client, id string, params UpdateAdStudyCellParams) (*objects.AdStudyCell, error) {
	var out objects.AdStudyCell
	call := UpdateAdStudyCellBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
