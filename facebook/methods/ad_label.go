package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLabelAdcreativesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAdcreativesBatchCall(id string, params GetAdLabelAdcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adcreatives"), params.ToParams(), options...)
}

func NewGetAdLabelAdcreativesBatchRequest(id string, params GetAdLabelAdcreativesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCreative]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCreative]](GetAdLabelAdcreativesBatchCall(id, params, options...))
}

func DecodeGetAdLabelAdcreativesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCreative], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCreative]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLabelAdcreativesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLabelAdcreativesParams) (*core.Cursor[objects.AdCreative], *core.Response, error) {
	var out core.Cursor[objects.AdCreative]
	call := GetAdLabelAdcreativesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLabelAdcreatives(ctx context.Context, client *core.Client, id string, params GetAdLabelAdcreativesParams) (*core.Cursor[objects.AdCreative], error) {
	out, _, err := GetAdLabelAdcreativesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdLabelAdsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAdsBatchCall(id string, params GetAdLabelAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), options...)
}

func NewGetAdLabelAdsBatchRequest(id string, params GetAdLabelAdsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetAdLabelAdsBatchCall(id, params, options...))
}

func DecodeGetAdLabelAdsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Ad]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLabelAdsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsParams) (*core.Cursor[objects.Ad], *core.Response, error) {
	var out core.Cursor[objects.Ad]
	call := GetAdLabelAdsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLabelAds(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsParams) (*core.Cursor[objects.Ad], error) {
	out, _, err := GetAdLabelAdsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdLabelAdsetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelAdsetsBatchCall(id string, params GetAdLabelAdsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), options...)
}

func NewGetAdLabelAdsetsBatchRequest(id string, params GetAdLabelAdsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdLabelAdsetsBatchCall(id, params, options...))
}

func DecodeGetAdLabelAdsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
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

func GetAdLabelAdsetsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsetsParams) (*core.Cursor[objects.AdSet], *core.Response, error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdLabelAdsetsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLabelAdsets(ctx context.Context, client *core.Client, id string, params GetAdLabelAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	out, _, err := GetAdLabelAdsetsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdLabelCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelCampaignsBatchCall(id string, params GetAdLabelCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewGetAdLabelCampaignsBatchRequest(id string, params GetAdLabelCampaignsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Campaign]] {
	return core.NewBatchRequest[core.Cursor[objects.Campaign]](GetAdLabelCampaignsBatchCall(id, params, options...))
}

func DecodeGetAdLabelCampaignsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Campaign], error) {
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

func GetAdLabelCampaignsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLabelCampaignsParams) (*core.Cursor[objects.Campaign], *core.Response, error) {
	var out core.Cursor[objects.Campaign]
	call := GetAdLabelCampaignsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLabelCampaigns(ctx context.Context, client *core.Client, id string, params GetAdLabelCampaignsParams) (*core.Cursor[objects.Campaign], error) {
	out, _, err := GetAdLabelCampaignsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdLabelBatchCall(id string, params DeleteAdLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdLabelBatchRequest(id string, params DeleteAdLabelParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdLabelBatchCall(id, params, options...))
}

func DecodeDeleteAdLabelBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdLabelWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdLabelParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdLabel(ctx context.Context, client *core.Client, id string, params DeleteAdLabelParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdLabelWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLabelBatchCall(id string, params GetAdLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdLabelBatchRequest(id string, params GetAdLabelParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLabel] {
	return core.NewBatchRequest[objects.AdLabel](GetAdLabelBatchCall(id, params, options...))
}

func DecodeGetAdLabelBatchResponse(response *core.BatchResponse) (*objects.AdLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLabelWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLabelParams) (*objects.AdLabel, *core.Response, error) {
	var out objects.AdLabel
	call := GetAdLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLabel(ctx context.Context, client *core.Client, id string, params GetAdLabelParams) (*objects.AdLabel, error) {
	out, _, err := GetAdLabelWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdLabelParams struct {
	Name  string      `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateAdLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	return out
}

func UpdateAdLabelBatchCall(id string, params UpdateAdLabelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdLabelBatchRequest(id string, params UpdateAdLabelParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLabel] {
	return core.NewBatchRequest[objects.AdLabel](UpdateAdLabelBatchCall(id, params, options...))
}

func DecodeUpdateAdLabelBatchResponse(response *core.BatchResponse) (*objects.AdLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdLabelWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdLabelParams) (*objects.AdLabel, *core.Response, error) {
	var out objects.AdLabel
	call := UpdateAdLabelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdLabel(ctx context.Context, client *core.Client, id string, params UpdateAdLabelParams) (*objects.AdLabel, error) {
	out, _, err := UpdateAdLabelWithResponse(ctx, client, id, params)
	return out, err
}
