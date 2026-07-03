package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
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

func GetOfflineConversionDataSetAdaccounts(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetAgencies(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetAudiences(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetAudiencesParams) (*core.Cursor[objects.CustomAudience], error) {
	var out core.Cursor[objects.CustomAudience]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "audiences"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetCustomconversions(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetCustomconversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	var out core.Cursor[objects.CustomConversion]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "customconversions"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetServerEventsPermittedBusiness(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetServerEventsPermittedBusinessParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "server_events_permitted_business"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetSharedAccounts(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetSharedAgencies(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetSharedAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_agencies"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSetStats(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetStatsParams) (*core.Cursor[objects.OfflineConversionDataSetStats], error) {
	var out core.Cursor[objects.OfflineConversionDataSetStats]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), &out)
	return &out, err
}

type GetOfflineConversionDataSetUploadsParams struct {
	EndTime   *time.Time                                            `facebook:"end_time"`
	Order     *enums.OfflineconversiondatasetuploadsOrderEnumParam  `facebook:"order"`
	SortBy    *enums.OfflineconversiondatasetuploadsSortByEnumParam `facebook:"sort_by"`
	StartTime *time.Time                                            `facebook:"start_time"`
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

func GetOfflineConversionDataSetUploads(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "uploads"), params.ToParams(), &out)
	return &out, err
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

func GetOfflineConversionDataSet(ctx context.Context, client *core.Client, id string, params GetOfflineConversionDataSetParams) (*objects.OfflineConversionDataSet, error) {
	var out objects.OfflineConversionDataSet
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
