package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOfflineConversionDataSetAdaccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetOfflineConversionDataSetAdaccountsBatchCall(id string, params GetOfflineConversionDataSetAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetAdaccountsBatchRequest(id string, params GetOfflineConversionDataSetAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetOfflineConversionDataSetAdaccountsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetOfflineConversionDataSetAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetOfflineConversionDataSetAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetAdaccounts(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetOfflineConversionDataSetAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetAgenciesBatchCall(id string, params GetOfflineConversionDataSetAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetAgenciesBatchRequest(id string, params GetOfflineConversionDataSetAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetOfflineConversionDataSetAgenciesBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetOfflineConversionDataSetAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetAgencies(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetOfflineConversionDataSetAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetAudiencesParams struct {
	ActionSource *enums.OfflineconversiondatasetaudiencesActionSourceEnumParam `facebook:"action_source"`
	AdAccount    *string                                                       `facebook:"ad_account"`
	Extra        core.Params                                                   `facebook:"-"`
}

func (params GetOfflineConversionDataSetAudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionSource != nil {
		out["action_source"] = *params.ActionSource
	}
	if params.AdAccount != nil {
		out["ad_account"] = *params.AdAccount
	}
	return out
}

func GetOfflineConversionDataSetAudiencesBatchCall(id string, params GetOfflineConversionDataSetAudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "audiences"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetAudiencesBatchRequest(id string, params GetOfflineConversionDataSetAudiencesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudience]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudience]](GetOfflineConversionDataSetAudiencesBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetAudiencesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudience], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudience]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetAudiencesWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAudiencesParams) (*core.Cursor[objects.CustomAudience], *core.Response, error) {
	var out core.Cursor[objects.CustomAudience]
	call := GetOfflineConversionDataSetAudiencesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetAudiences(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAudiencesParams) (*core.Cursor[objects.CustomAudience], error) {
	out, _, err := GetOfflineConversionDataSetAudiencesWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetCustomconversionsParams struct {
	AdAccount *string     `facebook:"ad_account"`
	Extra     core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetCustomconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccount != nil {
		out["ad_account"] = *params.AdAccount
	}
	return out
}

func GetOfflineConversionDataSetCustomconversionsBatchCall(id string, params GetOfflineConversionDataSetCustomconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "customconversions"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetCustomconversionsBatchRequest(id string, params GetOfflineConversionDataSetCustomconversionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomConversion]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomConversion]](GetOfflineConversionDataSetCustomconversionsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetCustomconversionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomConversion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomConversion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetCustomconversionsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetCustomconversionsParams) (*core.Cursor[objects.CustomConversion], *core.Response, error) {
	var out core.Cursor[objects.CustomConversion]
	call := GetOfflineConversionDataSetCustomconversionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetCustomconversions(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetCustomconversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	out, _, err := GetOfflineConversionDataSetCustomconversionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetServerEventsPermittedBusinessParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetServerEventsPermittedBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetServerEventsPermittedBusinessBatchCall(id string, params GetOfflineConversionDataSetServerEventsPermittedBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "server_events_permitted_business"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetServerEventsPermittedBusinessBatchRequest(id string, params GetOfflineConversionDataSetServerEventsPermittedBusinessParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetOfflineConversionDataSetServerEventsPermittedBusinessBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetServerEventsPermittedBusinessBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetServerEventsPermittedBusinessWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetServerEventsPermittedBusinessParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetOfflineConversionDataSetServerEventsPermittedBusinessBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetServerEventsPermittedBusiness(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetServerEventsPermittedBusinessParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetOfflineConversionDataSetServerEventsPermittedBusinessWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetSharedAccountsParams struct {
	ActionSource enums.OfflineconversiondatasetsharedAccountsActionSourceEnumParam `facebook:"action_source"`
	Business     string                                                            `facebook:"business"`
	Extra        core.Params                                                       `facebook:"-"`
}

func (params GetOfflineConversionDataSetSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["action_source"] = params.ActionSource
	out["business"] = params.Business
	return out
}

func GetOfflineConversionDataSetSharedAccountsBatchCall(id string, params GetOfflineConversionDataSetSharedAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetSharedAccountsBatchRequest(id string, params GetOfflineConversionDataSetSharedAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetOfflineConversionDataSetSharedAccountsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetSharedAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetOfflineConversionDataSetSharedAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetOfflineConversionDataSetSharedAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetSharedAccounts(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetOfflineConversionDataSetSharedAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetSharedAgenciesParams struct {
	ActionSource enums.OfflineconversiondatasetsharedAgenciesActionSourceEnumParam `facebook:"action_source"`
	Extra        core.Params                                                       `facebook:"-"`
}

func (params GetOfflineConversionDataSetSharedAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["action_source"] = params.ActionSource
	return out
}

func GetOfflineConversionDataSetSharedAgenciesBatchCall(id string, params GetOfflineConversionDataSetSharedAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_agencies"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetSharedAgenciesBatchRequest(id string, params GetOfflineConversionDataSetSharedAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetOfflineConversionDataSetSharedAgenciesBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetSharedAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetSharedAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetOfflineConversionDataSetSharedAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetSharedAgencies(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetOfflineConversionDataSetSharedAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetStatsParams struct {
	AggrTime        *enums.OfflineconversiondatasetstatsAggrTimeEnumParam    `facebook:"aggr_time"`
	End             *int                                                     `facebook:"end"`
	Granularity     *enums.OfflineconversiondatasetstatsGranularityEnumParam `facebook:"granularity"`
	SkipEmptyValues *bool                                                    `facebook:"skip_empty_values"`
	Start           *int                                                     `facebook:"start"`
	UserTimezoneID  *core.ID                                                 `facebook:"user_timezone_id"`
	Extra           core.Params                                              `facebook:"-"`
}

func (params GetOfflineConversionDataSetStatsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggrTime != nil {
		out["aggr_time"] = *params.AggrTime
	}
	if params.End != nil {
		out["end"] = *params.End
	}
	if params.Granularity != nil {
		out["granularity"] = *params.Granularity
	}
	if params.SkipEmptyValues != nil {
		out["skip_empty_values"] = *params.SkipEmptyValues
	}
	if params.Start != nil {
		out["start"] = *params.Start
	}
	if params.UserTimezoneID != nil {
		out["user_timezone_id"] = *params.UserTimezoneID
	}
	return out
}

func GetOfflineConversionDataSetStatsBatchCall(id string, params GetOfflineConversionDataSetStatsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetStatsBatchRequest(id string, params GetOfflineConversionDataSetStatsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OfflineConversionDataSetStats]] {
	return core.NewBatchRequest[core.Cursor[objects.OfflineConversionDataSetStats]](GetOfflineConversionDataSetStatsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetStatsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OfflineConversionDataSetStats], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OfflineConversionDataSetStats]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetStatsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetStatsParams) (*core.Cursor[objects.OfflineConversionDataSetStats], *core.Response, error) {
	var out core.Cursor[objects.OfflineConversionDataSetStats]
	call := GetOfflineConversionDataSetStatsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetStats(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetStatsParams) (*core.Cursor[objects.OfflineConversionDataSetStats], error) {
	out, _, err := GetOfflineConversionDataSetStatsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetUploadsParams struct {
	EndTime   *core.Time                                            `facebook:"end_time"`
	Order     *enums.OfflineconversiondatasetuploadsOrderEnumParam  `facebook:"order"`
	SortBy    *enums.OfflineconversiondatasetuploadsSortByEnumParam `facebook:"sort_by"`
	StartTime *core.Time                                            `facebook:"start_time"`
	UploadTag *string                                               `facebook:"upload_tag"`
	Extra     core.Params                                           `facebook:"-"`
}

func (params GetOfflineConversionDataSetUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.UploadTag != nil {
		out["upload_tag"] = *params.UploadTag
	}
	return out
}

func GetOfflineConversionDataSetUploadsBatchCall(id string, params GetOfflineConversionDataSetUploadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "uploads"), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetUploadsBatchRequest(id string, params GetOfflineConversionDataSetUploadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OfflineConversionDataSetUpload]] {
	return core.NewBatchRequest[core.Cursor[objects.OfflineConversionDataSetUpload]](GetOfflineConversionDataSetUploadsBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetUploadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetUploadsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], *core.Response, error) {
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	call := GetOfflineConversionDataSetUploadsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSetUploads(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	out, _, err := GetOfflineConversionDataSetUploadsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineConversionDataSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineConversionDataSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineConversionDataSetBatchCall(id string, params GetOfflineConversionDataSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOfflineConversionDataSetBatchRequest(id string, params GetOfflineConversionDataSetParams, options ...core.BatchOption) *core.BatchRequest[objects.OfflineConversionDataSet] {
	return core.NewBatchRequest[objects.OfflineConversionDataSet](GetOfflineConversionDataSetBatchCall(id, params, options...))
}

func DecodeGetOfflineConversionDataSetBatchResponse(response *core.BatchResponse) (*objects.OfflineConversionDataSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OfflineConversionDataSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineConversionDataSetWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetParams) (*objects.OfflineConversionDataSet, *core.Response, error) {
	var out objects.OfflineConversionDataSet
	call := GetOfflineConversionDataSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineConversionDataSet(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetParams) (*objects.OfflineConversionDataSet, error) {
	out, _, err := GetOfflineConversionDataSetWithResponse(ctx, client, id, params)
	return out, err
}
