package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type CreateBusinessAccessTokenParams struct {
	AppID                 core.ID              `facebook:"app_id"`
	FbeExternalBusinessID *core.ID             `facebook:"fbe_external_business_id"`
	Scope                 []objects.Permission `facebook:"scope"`
	SystemUserName        *string              `facebook:"system_user_name"`
	Extra                 core.Params          `facebook:"-"`
}

func (params CreateBusinessAccessTokenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_id"] = params.AppID
	if params.FbeExternalBusinessID != nil {
		out["fbe_external_business_id"] = *params.FbeExternalBusinessID
	}
	out["scope"] = params.Scope
	if params.SystemUserName != nil {
		out["system_user_name"] = *params.SystemUserName
	}
	return out
}

func CreateBusinessAccessToken(ctx context.Context, client *core.Client, id string, params CreateBusinessAccessTokenParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "access_token"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdAccountInfosParams struct {
	AdAccountID        *core.ID    `facebook:"ad_account_id"`
	ParentAdvertiserID *core.ID    `facebook:"parent_advertiser_id"`
	UserID             *core.ID    `facebook:"user_id"`
	Extra              core.Params `facebook:"-"`
}

func (params GetBusinessAdAccountInfosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountID != nil {
		out["ad_account_id"] = *params.AdAccountID
	}
	if params.ParentAdvertiserID != nil {
		out["parent_advertiser_id"] = *params.ParentAdvertiserID
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	return out
}

func GetBusinessAdAccountInfos(ctx context.Context, client *core.Client, id string, params GetBusinessAdAccountInfosParams) (*core.Cursor[objects.ALMAdAccountInfo], error) {
	var out core.Cursor[objects.ALMAdAccountInfo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_account_infos"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAdAccountsParams struct {
	AdaccountID core.ID     `facebook:"adaccount_id"`
	Extra       core.Params `facebook:"-"`
}

func (params DeleteBusinessAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["adaccount_id"] = params.AdaccountID
	return out
}

func DeleteBusinessAdAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAdAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdCustomDerivedMetricsParams struct {
	AdhocCustomMetrics *[]string                                           `facebook:"adhoc_custom_metrics"`
	Scope              *enums.BusinessadCustomDerivedMetricsScopeEnumParam `facebook:"scope"`
	Extra              core.Params                                         `facebook:"-"`
}

func (params GetBusinessAdCustomDerivedMetricsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdhocCustomMetrics != nil {
		out["adhoc_custom_metrics"] = *params.AdhocCustomMetrics
	}
	if params.Scope != nil {
		out["scope"] = *params.Scope
	}
	return out
}

func GetBusinessAdCustomDerivedMetrics(ctx context.Context, client *core.Client, id string, params GetBusinessAdCustomDerivedMetricsParams) (*core.Cursor[objects.AdCustomDerivedMetrics], error) {
	var out core.Cursor[objects.AdCustomDerivedMetrics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_custom_derived_metrics"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdReviewRequestsParams struct {
	AdAccountIds *[]core.ID  `facebook:"ad_account_ids"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateBusinessAdReviewRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountIds != nil {
		out["ad_account_ids"] = *params.AdAccountIds
	}
	return out
}

func CreateBusinessAdReviewRequests(ctx context.Context, client *core.Client, id string, params CreateBusinessAdReviewRequestsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ad_review_requests"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAdStudies(ctx context.Context, client *core.Client, id string, params GetBusinessAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdStudiesParams struct {
	Cells              []map[string]interface{}              `facebook:"cells"`
	ClientBusiness     *string                               `facebook:"client_business"`
	ConfidenceLevel    *float64                              `facebook:"confidence_level"`
	CooldownStartTime  *int                                  `facebook:"cooldown_start_time"`
	CreativeTestConfig *map[string]interface{}               `facebook:"creative_test_config"`
	Description        *string                               `facebook:"description"`
	EndTime            int                                   `facebook:"end_time"`
	Name               string                                `facebook:"name"`
	Objectives         *[]map[string]interface{}             `facebook:"objectives"`
	ObservationEndTime *int                                  `facebook:"observation_end_time"`
	StartTime          int                                   `facebook:"start_time"`
	Type               *enums.BusinessadStudiesTypeEnumParam `facebook:"type"`
	Viewers            *[]int                                `facebook:"viewers"`
	Extra              core.Params                           `facebook:"-"`
}

func (params CreateBusinessAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["cells"] = params.Cells
	if params.ClientBusiness != nil {
		out["client_business"] = *params.ClientBusiness
	}
	if params.ConfidenceLevel != nil {
		out["confidence_level"] = *params.ConfidenceLevel
	}
	if params.CooldownStartTime != nil {
		out["cooldown_start_time"] = *params.CooldownStartTime
	}
	if params.CreativeTestConfig != nil {
		out["creative_test_config"] = *params.CreativeTestConfig
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	out["end_time"] = params.EndTime
	out["name"] = params.Name
	if params.Objectives != nil {
		out["objectives"] = *params.Objectives
	}
	if params.ObservationEndTime != nil {
		out["observation_end_time"] = *params.ObservationEndTime
	}
	out["start_time"] = params.StartTime
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Viewers != nil {
		out["viewers"] = *params.Viewers
	}
	return out
}

func CreateBusinessAdStudies(ctx context.Context, client *core.Client, id string, params CreateBusinessAdStudiesParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ad_studies"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdaccountParams struct {
	AdAccountCreatedFromBmFlag *bool                  `facebook:"ad_account_created_from_bm_flag"`
	Currency                   string                 `facebook:"currency"`
	EndAdvertiser              map[string]interface{} `facebook:"end_advertiser"`
	FundingID                  *core.ID               `facebook:"funding_id"`
	Invoice                    *bool                  `facebook:"invoice"`
	InvoiceGroupID             *core.ID               `facebook:"invoice_group_id"`
	InvoicingEmails            *[]string              `facebook:"invoicing_emails"`
	Io                         *bool                  `facebook:"io"`
	MediaAgency                string                 `facebook:"media_agency"`
	Name                       string                 `facebook:"name"`
	Partner                    string                 `facebook:"partner"`
	PoNumber                   *string                `facebook:"po_number"`
	TimezoneID                 core.ID                `facebook:"timezone_id"`
	Extra                      core.Params            `facebook:"-"`
}

func (params CreateBusinessAdaccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountCreatedFromBmFlag != nil {
		out["ad_account_created_from_bm_flag"] = *params.AdAccountCreatedFromBmFlag
	}
	out["currency"] = params.Currency
	out["end_advertiser"] = params.EndAdvertiser
	if params.FundingID != nil {
		out["funding_id"] = *params.FundingID
	}
	if params.Invoice != nil {
		out["invoice"] = *params.Invoice
	}
	if params.InvoiceGroupID != nil {
		out["invoice_group_id"] = *params.InvoiceGroupID
	}
	if params.InvoicingEmails != nil {
		out["invoicing_emails"] = *params.InvoicingEmails
	}
	if params.Io != nil {
		out["io"] = *params.Io
	}
	out["media_agency"] = params.MediaAgency
	out["name"] = params.Name
	out["partner"] = params.Partner
	if params.PoNumber != nil {
		out["po_number"] = *params.PoNumber
	}
	out["timezone_id"] = params.TimezoneID
	return out
}

func CreateBusinessAdaccount(ctx context.Context, client *core.Client, id string, params CreateBusinessAdaccountParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adaccount"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAddPhoneNumbersParams struct {
	PhoneNumber string      `facebook:"phone_number"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateBusinessAddPhoneNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["phone_number"] = params.PhoneNumber
	return out
}

func CreateBusinessAddPhoneNumbers(ctx context.Context, client *core.Client, id string, params CreateBusinessAddPhoneNumbersParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "add_phone_numbers"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdnetworkApplicationsParams struct {
	Name  string      `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params CreateBusinessAdnetworkApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	return out
}

func CreateBusinessAdnetworkApplications(ctx context.Context, client *core.Client, id string, params CreateBusinessAdnetworkApplicationsParams) (*objects.Application, error) {
	var out objects.Application
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adnetwork_applications"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdnetworkanalyticsParams struct {
	AggregationPeriod  *enums.BusinessadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns         *[]enums.BusinessadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters            *[]map[string]interface{}                                   `facebook:"filters"`
	Limit              *uint64                                                     `facebook:"limit"`
	Metrics            []enums.BusinessadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn     *enums.BusinessadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType       *enums.BusinessadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	ShouldIncludeUntil *bool                                                       `facebook:"should_include_until"`
	Since              *time.Time                                                  `facebook:"since"`
	Until              *time.Time                                                  `facebook:"until"`
	Extra              core.Params                                                 `facebook:"-"`
}

func (params GetBusinessAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.ShouldIncludeUntil != nil {
		out["should_include_until"] = *params.ShouldIncludeUntil
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetBusinessAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetBusinessAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdnetworkanalyticsParams struct {
	AggregationPeriod *enums.BusinessadnetworkanalyticsAggregationPeriodEnumParam `facebook:"aggregation_period"`
	Breakdowns        *[]enums.BusinessadnetworkanalyticsBreakdownsEnumParam      `facebook:"breakdowns"`
	Filters           *[]map[string]interface{}                                   `facebook:"filters"`
	Limit             *int                                                        `facebook:"limit"`
	Metrics           []enums.BusinessadnetworkanalyticsMetricsEnumParam          `facebook:"metrics"`
	OrderingColumn    *enums.BusinessadnetworkanalyticsOrderingColumnEnumParam    `facebook:"ordering_column"`
	OrderingType      *enums.BusinessadnetworkanalyticsOrderingTypeEnumParam      `facebook:"ordering_type"`
	Since             *time.Time                                                  `facebook:"since"`
	Until             *time.Time                                                  `facebook:"until"`
	Extra             core.Params                                                 `facebook:"-"`
}

func (params CreateBusinessAdnetworkanalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AggregationPeriod != nil {
		out["aggregation_period"] = *params.AggregationPeriod
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	out["metrics"] = params.Metrics
	if params.OrderingColumn != nil {
		out["ordering_column"] = *params.OrderingColumn
	}
	if params.OrderingType != nil {
		out["ordering_type"] = *params.OrderingType
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func CreateBusinessAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateBusinessAdnetworkanalyticsParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdnetworkanalyticsResultsParams struct {
	QueryIds *[]core.ID  `facebook:"query_ids"`
	Extra    core.Params `facebook:"-"`
}

func (params GetBusinessAdnetworkanalyticsResultsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.QueryIds != nil {
		out["query_ids"] = *params.QueryIds
	}
	return out
}

func GetBusinessAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetBusinessAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdsDatasetParams struct {
	IDFilter   *string                                  `facebook:"id_filter"`
	NameFilter *string                                  `facebook:"name_filter"`
	SortBy     *enums.BusinessadsDatasetSortByEnumParam `facebook:"sort_by"`
	Extra      core.Params                              `facebook:"-"`
}

func (params GetBusinessAdsDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IDFilter != nil {
		out["id_filter"] = *params.IDFilter
	}
	if params.NameFilter != nil {
		out["name_filter"] = *params.NameFilter
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	return out
}

func GetBusinessAdsDataset(ctx context.Context, client *core.Client, id string, params GetBusinessAdsDatasetParams) (*core.Cursor[objects.AdsDataset], error) {
	var out core.Cursor[objects.AdsDataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads_dataset"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdsDatasetParams struct {
	AdAccountID *core.ID    `facebook:"ad_account_id"`
	AppID       *core.ID    `facebook:"app_id"`
	IsCrm       *bool       `facebook:"is_crm"`
	Name        string      `facebook:"name"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateBusinessAdsDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountID != nil {
		out["ad_account_id"] = *params.AdAccountID
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.IsCrm != nil {
		out["is_crm"] = *params.IsCrm
	}
	out["name"] = params.Name
	return out
}

func CreateBusinessAdsDataset(ctx context.Context, client *core.Client, id string, params CreateBusinessAdsDatasetParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ads_dataset"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdsReportingMmmReportsParams struct {
	Filtering *[]map[string]interface{} `facebook:"filtering"`
	Extra     core.Params               `facebook:"-"`
}

func (params GetBusinessAdsReportingMmmReportsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	return out
}

func GetBusinessAdsReportingMmmReports(ctx context.Context, client *core.Client, id string, params GetBusinessAdsReportingMmmReportsParams) (*core.Cursor[objects.AdsReportBuilderMMMReport], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReport]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_reports"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdsReportingMmmSchedulersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAdsReportingMmmSchedulersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAdsReportingMmmSchedulers(ctx context.Context, client *core.Client, id string, params GetBusinessAdsReportingMmmSchedulersParams) (*core.Cursor[objects.AdsReportBuilderMMMReportScheduler], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReportScheduler]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_schedulers"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAdspixelsParams struct {
	IDFilter   *string                                 `facebook:"id_filter"`
	NameFilter *string                                 `facebook:"name_filter"`
	SortBy     *enums.BusinessadspixelsSortByEnumParam `facebook:"sort_by"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetBusinessAdspixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IDFilter != nil {
		out["id_filter"] = *params.IDFilter
	}
	if params.NameFilter != nil {
		out["name_filter"] = *params.NameFilter
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	return out
}

func GetBusinessAdspixels(ctx context.Context, client *core.Client, id string, params GetBusinessAdspixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adspixels"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAdspixelsParams struct {
	IsCrm *bool       `facebook:"is_crm"`
	Name  string      `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params CreateBusinessAdspixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsCrm != nil {
		out["is_crm"] = *params.IsCrm
	}
	out["name"] = params.Name
	return out
}

func CreateBusinessAdspixels(ctx context.Context, client *core.Client, id string, params CreateBusinessAdspixelsParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adspixels"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteBusinessAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteBusinessAgencies(ctx context.Context, client *core.Client, id string, params DeleteBusinessAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAgencies(ctx context.Context, client *core.Client, id string, params GetBusinessAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAnPlacementsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAnPlacementsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAnPlacements(ctx context.Context, client *core.Client, id string, params GetBusinessAnPlacementsParams) (*core.Cursor[objects.AdPlacement], error) {
	var out core.Cursor[objects.AdPlacement]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "an_placements"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessBlockListDraftsParams struct {
	PublisherUrlsFile core.FileParam `facebook:"publisher_urls_file"`
	Extra             core.Params    `facebook:"-"`
}

func (params CreateBusinessBlockListDraftsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["publisher_urls_file"] = params.PublisherUrlsFile
	return out
}

func CreateBusinessBlockListDrafts(ctx context.Context, client *core.Client, id string, params CreateBusinessBlockListDraftsParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "block_list_drafts"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessBmReviewRequestsParams struct {
	BusinessManagerIds []core.ID   `facebook:"business_manager_ids"`
	Extra              core.Params `facebook:"-"`
}

func (params CreateBusinessBmReviewRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business_manager_ids"] = params.BusinessManagerIds
	return out
}

func CreateBusinessBmReviewRequests(ctx context.Context, client *core.Client, id string, params CreateBusinessBmReviewRequestsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "bm_review_requests"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessBusinessAssetGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessBusinessAssetGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "business_asset_groups"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessBusinessInvoicesParams struct {
	EndDate        *string                                      `facebook:"end_date"`
	InvoiceID      *core.ID                                     `facebook:"invoice_id"`
	IssueEndDate   *string                                      `facebook:"issue_end_date"`
	IssueStartDate *string                                      `facebook:"issue_start_date"`
	RootID         *core.ID                                     `facebook:"root_id"`
	StartDate      *string                                      `facebook:"start_date"`
	Type           *enums.BusinessbusinessInvoicesTypeEnumParam `facebook:"type"`
	Extra          core.Params                                  `facebook:"-"`
}

func (params GetBusinessBusinessInvoicesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndDate != nil {
		out["end_date"] = *params.EndDate
	}
	if params.InvoiceID != nil {
		out["invoice_id"] = *params.InvoiceID
	}
	if params.IssueEndDate != nil {
		out["issue_end_date"] = *params.IssueEndDate
	}
	if params.IssueStartDate != nil {
		out["issue_start_date"] = *params.IssueStartDate
	}
	if params.RootID != nil {
		out["root_id"] = *params.RootID
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetBusinessBusinessInvoices(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessInvoicesParams) (*core.Cursor[objects.OmegaCustomerTrx], error) {
	var out core.Cursor[objects.OmegaCustomerTrx]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "business_invoices"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessBusinessUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessBusinessUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessBusinessUsers(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessUsersParams) (*core.Cursor[objects.BusinessUser], error) {
	var out core.Cursor[objects.BusinessUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "business_users"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessBusinessUsersParams struct {
	Email           string                                                 `facebook:"email"`
	InvitedUserType *[]enums.BusinessbusinessUsersInvitedUserTypeEnumParam `facebook:"invited_user_type"`
	Role            *enums.BusinessbusinessUsersRoleEnumParam              `facebook:"role"`
	Tasks           *[]enums.BusinessbusinessUsersTasksEnumParam           `facebook:"tasks"`
	Extra           core.Params                                            `facebook:"-"`
}

func (params CreateBusinessBusinessUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["email"] = params.Email
	if params.InvitedUserType != nil {
		out["invited_user_type"] = *params.InvitedUserType
	}
	if params.Role != nil {
		out["role"] = *params.Role
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	return out
}

func CreateBusinessBusinessUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessBusinessUsersParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "business_users"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessBusinessprojectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessBusinessprojectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessBusinessprojects(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessprojectsParams) (*core.Cursor[objects.BusinessProject], error) {
	var out core.Cursor[objects.BusinessProject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "businessprojects"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessClaimCustomConversionsParams struct {
	CustomConversionID core.ID     `facebook:"custom_conversion_id"`
	Extra              core.Params `facebook:"-"`
}

func (params CreateBusinessClaimCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["custom_conversion_id"] = params.CustomConversionID
	return out
}

func CreateBusinessClaimCustomConversions(ctx context.Context, client *core.Client, id string, params CreateBusinessClaimCustomConversionsParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "claim_custom_conversions"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientAdAccountsParams struct {
	SearchQuery *string     `facebook:"search_query"`
	Extra       core.Params `facebook:"-"`
}

func (params GetBusinessClientAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SearchQuery != nil {
		out["search_query"] = *params.SearchQuery
	}
	return out
}

func GetBusinessClientAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessClientAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientApps(ctx context.Context, client *core.Client, id string, params GetBusinessClientAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_apps"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessClientAppsParams struct {
	AppID map[string]interface{} `facebook:"app_id"`
	Extra core.Params            `facebook:"-"`
}

func (params CreateBusinessClientAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_id"] = params.AppID
	return out
}

func CreateBusinessClientApps(ctx context.Context, client *core.Client, id string, params CreateBusinessClientAppsParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "client_apps"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientInstagramAssetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientInstagramAssetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientInstagramAssets(ctx context.Context, client *core.Client, id string, params GetBusinessClientInstagramAssetsParams) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	var out core.Cursor[objects.InstagramBusinessAsset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_instagram_assets"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientOffsiteSignalContainerBusinessObjectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientOffsiteSignalContainerBusinessObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessClientOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_offsite_signal_container_business_objects"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientPages(ctx context.Context, client *core.Client, id string, params GetBusinessClientPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_pages"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessClientPagesParams struct {
	PageID         core.ID                                             `facebook:"page_id"`
	PermittedTasks *[]enums.BusinessclientPagesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                         `facebook:"-"`
}

func (params CreateBusinessClientPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["page_id"] = params.PageID
	if params.PermittedTasks != nil {
		out["permitted_tasks"] = *params.PermittedTasks
	}
	return out
}

func CreateBusinessClientPages(ctx context.Context, client *core.Client, id string, params CreateBusinessClientPagesParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "client_pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientPixelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientPixels(ctx context.Context, client *core.Client, id string, params GetBusinessClientPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_pixels"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessClientProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClientWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessClientWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "client_whatsapp_business_accounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessClientsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteBusinessClientsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteBusinessClients(ctx context.Context, client *core.Client, id string, params DeleteBusinessClientsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "clients"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessClientsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessClientsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessClients(ctx context.Context, client *core.Client, id string, params GetBusinessClientsParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "clients"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCollaborativeAdsCollaborationRequestsParams struct {
	RequestRole *enums.BusinesscollaborativeAdsCollaborationRequestsRequestRoleEnumParam `facebook:"request_role"`
	Since       *time.Time                                                               `facebook:"since"`
	Source      *enums.BusinesscollaborativeAdsCollaborationRequestsSourceEnumParam      `facebook:"source"`
	Status      *string                                                                  `facebook:"status"`
	Until       *time.Time                                                               `facebook:"until"`
	Extra       core.Params                                                              `facebook:"-"`
}

func (params GetBusinessCollaborativeAdsCollaborationRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.RequestRole != nil {
		out["request_role"] = *params.RequestRole
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetBusinessCollaborativeAdsCollaborationRequests(ctx context.Context, client *core.Client, id string, params GetBusinessCollaborativeAdsCollaborationRequestsParams) (*core.Cursor[objects.CPASCollaborationRequest], error) {
	var out core.Cursor[objects.CPASCollaborationRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborative_ads_collaboration_requests"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCollaborativeAdsSuggestedPartnersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCollaborativeAdsSuggestedPartnersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCollaborativeAdsSuggestedPartners(ctx context.Context, client *core.Client, id string, params GetBusinessCollaborativeAdsSuggestedPartnersParams) (*core.Cursor[objects.CPASAdvertiserPartnershipRecommendation], error) {
	var out core.Cursor[objects.CPASAdvertiserPartnershipRecommendation]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborative_ads_suggested_partners"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCommerceMerchantSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCommerceMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params GetBusinessCommerceMerchantSettingsParams) (*core.Cursor[objects.CommerceMerchantSettings], error) {
	var out core.Cursor[objects.CommerceMerchantSettings]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "commerce_merchant_settings"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCpasBusinessSetupConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCpasBusinessSetupConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCpasBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params GetBusinessCpasBusinessSetupConfigParams) (*core.Cursor[objects.CPASBusinessSetupConfig], error) {
	var out core.Cursor[objects.CPASBusinessSetupConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "cpas_business_setup_config"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessCpasBusinessSetupConfigParams struct {
	AcceptedCollabAdsTos         *bool                   `facebook:"accepted_collab_ads_tos"`
	AdAccounts                   *[]string               `facebook:"ad_accounts"`
	BusinessCapabilitiesStatus   *map[string]interface{} `facebook:"business_capabilities_status"`
	CapabilitiesComplianceStatus *map[string]interface{} `facebook:"capabilities_compliance_status"`
	Extra                        core.Params             `facebook:"-"`
}

func (params CreateBusinessCpasBusinessSetupConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AcceptedCollabAdsTos != nil {
		out["accepted_collab_ads_tos"] = *params.AcceptedCollabAdsTos
	}
	if params.AdAccounts != nil {
		out["ad_accounts"] = *params.AdAccounts
	}
	if params.BusinessCapabilitiesStatus != nil {
		out["business_capabilities_status"] = *params.BusinessCapabilitiesStatus
	}
	if params.CapabilitiesComplianceStatus != nil {
		out["capabilities_compliance_status"] = *params.CapabilitiesComplianceStatus
	}
	return out
}

func CreateBusinessCpasBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params CreateBusinessCpasBusinessSetupConfigParams) (*objects.CPASBusinessSetupConfig, error) {
	var out objects.CPASBusinessSetupConfig
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "cpas_business_setup_config"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCpasMerchantConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCpasMerchantConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCpasMerchantConfig(ctx context.Context, client *core.Client, id string, params GetBusinessCpasMerchantConfigParams) (*core.Cursor[objects.CPASMerchantConfig], error) {
	var out core.Cursor[objects.CPASMerchantConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "cpas_merchant_config"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessCreativeFoldersParams struct {
	Description    *string     `facebook:"description"`
	Name           string      `facebook:"name"`
	ParentFolderID *core.ID    `facebook:"parent_folder_id"`
	Extra          core.Params `facebook:"-"`
}

func (params CreateBusinessCreativeFoldersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	out["name"] = params.Name
	if params.ParentFolderID != nil {
		out["parent_folder_id"] = *params.ParentFolderID
	}
	return out
}

func CreateBusinessCreativeFolders(ctx context.Context, client *core.Client, id string, params CreateBusinessCreativeFoldersParams) (*objects.BusinessCreativeFolder, error) {
	var out objects.BusinessCreativeFolder
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "creative_folders"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessCreditcardsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessCreditcardsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessCreditcards(ctx context.Context, client *core.Client, id string, params GetBusinessCreditcardsParams) (*core.Cursor[objects.CreditCard], error) {
	var out core.Cursor[objects.CreditCard]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "creditcards"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessCustomconversionsParams struct {
	ActionSourceType       *enums.BusinesscustomconversionsActionSourceTypeEnumParam `facebook:"action_source_type"`
	AdvancedRule           *string                                                   `facebook:"advanced_rule"`
	CustomEventType        enums.BusinesscustomconversionsCustomEventTypeEnumParam   `facebook:"custom_event_type"`
	DefaultConversionValue *float64                                                  `facebook:"default_conversion_value"`
	Description            *string                                                   `facebook:"description"`
	EventSourceID          *core.ID                                                  `facebook:"event_source_id"`
	Name                   string                                                    `facebook:"name"`
	Rule                   *string                                                   `facebook:"rule"`
	Extra                  core.Params                                               `facebook:"-"`
}

func (params CreateBusinessCustomconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionSourceType != nil {
		out["action_source_type"] = *params.ActionSourceType
	}
	if params.AdvancedRule != nil {
		out["advanced_rule"] = *params.AdvancedRule
	}
	out["custom_event_type"] = params.CustomEventType
	if params.DefaultConversionValue != nil {
		out["default_conversion_value"] = *params.DefaultConversionValue
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EventSourceID != nil {
		out["event_source_id"] = *params.EventSourceID
	}
	out["name"] = params.Name
	if params.Rule != nil {
		out["rule"] = *params.Rule
	}
	return out
}

func CreateBusinessCustomconversions(ctx context.Context, client *core.Client, id string, params CreateBusinessCustomconversionsParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "customconversions"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessEventSourceGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessEventSourceGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessEventSourceGroups(ctx context.Context, client *core.Client, id string, params GetBusinessEventSourceGroupsParams) (*core.Cursor[objects.EventSourceGroup], error) {
	var out core.Cursor[objects.EventSourceGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "event_source_groups"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessEventSourceGroupsParams struct {
	EventSources []string    `facebook:"event_sources"`
	Name         string      `facebook:"name"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateBusinessEventSourceGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["event_sources"] = params.EventSources
	out["name"] = params.Name
	return out
}

func CreateBusinessEventSourceGroups(ctx context.Context, client *core.Client, id string, params CreateBusinessEventSourceGroupsParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "event_source_groups"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessExtendedcreditapplicationsParams struct {
	OnlyShowPending *bool       `facebook:"only_show_pending"`
	Extra           core.Params `facebook:"-"`
}

func (params GetBusinessExtendedcreditapplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.OnlyShowPending != nil {
		out["only_show_pending"] = *params.OnlyShowPending
	}
	return out
}

func GetBusinessExtendedcreditapplications(ctx context.Context, client *core.Client, id string, params GetBusinessExtendedcreditapplicationsParams) (*core.Cursor[objects.ExtendedCreditApplication], error) {
	var out core.Cursor[objects.ExtendedCreditApplication]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "extendedcreditapplications"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessExtendedcreditsParams struct {
	OrderByIsOwnedCredential *bool       `facebook:"order_by_is_owned_credential"`
	Extra                    core.Params `facebook:"-"`
}

func (params GetBusinessExtendedcreditsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.OrderByIsOwnedCredential != nil {
		out["order_by_is_owned_credential"] = *params.OrderByIsOwnedCredential
	}
	return out
}

func GetBusinessExtendedcredits(ctx context.Context, client *core.Client, id string, params GetBusinessExtendedcreditsParams) (*core.Cursor[objects.ExtendedCredit], error) {
	var out core.Cursor[objects.ExtendedCredit]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "extendedcredits"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessImagesParams struct {
	AdPlacementsValidationOnly *bool                                                  `facebook:"ad_placements_validation_only"`
	Bytes                      *string                                                `facebook:"bytes"`
	CreativeFolderID           core.ID                                                `facebook:"creative_folder_id"`
	Name                       *string                                                `facebook:"name"`
	ValidationAdPlacements     *[]enums.BusinessimagesValidationAdPlacementsEnumParam `facebook:"validation_ad_placements"`
	Extra                      core.Params                                            `facebook:"-"`
}

func (params CreateBusinessImagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdPlacementsValidationOnly != nil {
		out["ad_placements_validation_only"] = *params.AdPlacementsValidationOnly
	}
	if params.Bytes != nil {
		out["bytes"] = *params.Bytes
	}
	out["creative_folder_id"] = params.CreativeFolderID
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.ValidationAdPlacements != nil {
		out["validation_ad_placements"] = *params.ValidationAdPlacements
	}
	return out
}

func CreateBusinessImages(ctx context.Context, client *core.Client, id string, params CreateBusinessImagesParams) (*objects.BusinessImage, error) {
	var out objects.BusinessImage
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "images"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessInitiatedAudienceSharingRequestsParams struct {
	RecipientID   *core.ID                                                              `facebook:"recipient_id"`
	RequestStatus *enums.BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam `facebook:"request_status"`
	Extra         core.Params                                                           `facebook:"-"`
}

func (params GetBusinessInitiatedAudienceSharingRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.RecipientID != nil {
		out["recipient_id"] = *params.RecipientID
	}
	if params.RequestStatus != nil {
		out["request_status"] = *params.RequestStatus
	}
	return out
}

func GetBusinessInitiatedAudienceSharingRequests(ctx context.Context, client *core.Client, id string, params GetBusinessInitiatedAudienceSharingRequestsParams) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "initiated_audience_sharing_requests"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessInstagramAccountsParams struct {
	InstagramAccount string      `facebook:"instagram_account"`
	Extra            core.Params `facebook:"-"`
}

func (params DeleteBusinessInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["instagram_account"] = params.InstagramAccount
	return out
}

func DeleteBusinessInstagramAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessInstagramAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessInstagramBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessInstagramBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessInstagramBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessInstagramBusinessAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "instagram_business_accounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessManagedBusinessesParams struct {
	ExistingClientBusinessID core.ID     `facebook:"existing_client_business_id"`
	Extra                    core.Params `facebook:"-"`
}

func (params DeleteBusinessManagedBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["existing_client_business_id"] = params.ExistingClientBusinessID
	return out
}

func DeleteBusinessManagedBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessManagedBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "managed_businesses"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessManagedBusinessesParams struct {
	ChildBusinessExternalID  *core.ID                                                    `facebook:"child_business_external_id"`
	ExistingClientBusinessID *core.ID                                                    `facebook:"existing_client_business_id"`
	Name                     *string                                                     `facebook:"name"`
	SalesRepEmail            *string                                                     `facebook:"sales_rep_email"`
	SurveyBusinessType       *enums.BusinessmanagedBusinessesSurveyBusinessTypeEnumParam `facebook:"survey_business_type"`
	SurveyNumAssets          *uint64                                                     `facebook:"survey_num_assets"`
	SurveyNumPeople          *uint64                                                     `facebook:"survey_num_people"`
	TimezoneID               *enums.BusinessmanagedBusinessesTimezoneIDEnumParam         `facebook:"timezone_id"`
	Vertical                 *enums.BusinessmanagedBusinessesVerticalEnumParam           `facebook:"vertical"`
	Extra                    core.Params                                                 `facebook:"-"`
}

func (params CreateBusinessManagedBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.ExistingClientBusinessID != nil {
		out["existing_client_business_id"] = *params.ExistingClientBusinessID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SalesRepEmail != nil {
		out["sales_rep_email"] = *params.SalesRepEmail
	}
	if params.SurveyBusinessType != nil {
		out["survey_business_type"] = *params.SurveyBusinessType
	}
	if params.SurveyNumAssets != nil {
		out["survey_num_assets"] = *params.SurveyNumAssets
	}
	if params.SurveyNumPeople != nil {
		out["survey_num_people"] = *params.SurveyNumPeople
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	if params.Vertical != nil {
		out["vertical"] = *params.Vertical
	}
	return out
}

func CreateBusinessManagedBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedBusinessesParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "managed_businesses"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessManagedPartnerAdsFundingSourceDetailsParams struct {
	YearQuarter *string     `facebook:"year_quarter"`
	Extra       core.Params `facebook:"-"`
}

func (params GetBusinessManagedPartnerAdsFundingSourceDetailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.YearQuarter != nil {
		out["year_quarter"] = *params.YearQuarter
	}
	return out
}

func GetBusinessManagedPartnerAdsFundingSourceDetails(ctx context.Context, client *core.Client, id string, params GetBusinessManagedPartnerAdsFundingSourceDetailsParams) (*core.Cursor[objects.FundingSourceDetailsCoupon], error) {
	var out core.Cursor[objects.FundingSourceDetailsCoupon]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "managed_partner_ads_funding_source_details"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessManagedPartnerBusinessSetupParams struct {
	ActiveAdAccountID            *core.ID                  `facebook:"active_ad_account_id"`
	ActivePageID                 *core.ID                  `facebook:"active_page_id"`
	PartnerFacebookPageURL       *string                   `facebook:"partner_facebook_page_url"`
	PartnerRegistrationCountries *[]string                 `facebook:"partner_registration_countries"`
	SellerEmailAddress           *string                   `facebook:"seller_email_address"`
	SellerExternalWebsiteURL     *string                   `facebook:"seller_external_website_url"`
	Template                     *[]map[string]interface{} `facebook:"template"`
	Extra                        core.Params               `facebook:"-"`
}

func (params CreateBusinessManagedPartnerBusinessSetupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActiveAdAccountID != nil {
		out["active_ad_account_id"] = *params.ActiveAdAccountID
	}
	if params.ActivePageID != nil {
		out["active_page_id"] = *params.ActivePageID
	}
	if params.PartnerFacebookPageURL != nil {
		out["partner_facebook_page_url"] = *params.PartnerFacebookPageURL
	}
	if params.PartnerRegistrationCountries != nil {
		out["partner_registration_countries"] = *params.PartnerRegistrationCountries
	}
	if params.SellerEmailAddress != nil {
		out["seller_email_address"] = *params.SellerEmailAddress
	}
	if params.SellerExternalWebsiteURL != nil {
		out["seller_external_website_url"] = *params.SellerExternalWebsiteURL
	}
	if params.Template != nil {
		out["template"] = *params.Template
	}
	return out
}

func CreateBusinessManagedPartnerBusinessSetup(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedPartnerBusinessSetupParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "managed_partner_business_setup"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessManagedPartnerBusinessesParams struct {
	ChildBusinessExternalID *core.ID    `facebook:"child_business_external_id"`
	ChildBusinessID         *core.ID    `facebook:"child_business_id"`
	Extra                   core.Params `facebook:"-"`
}

func (params DeleteBusinessManagedPartnerBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.ChildBusinessID != nil {
		out["child_business_id"] = *params.ChildBusinessID
	}
	return out
}

func DeleteBusinessManagedPartnerBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessManagedPartnerBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "managed_partner_businesses"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessManagedPartnerBusinessesParams struct {
	AdAccountCurrency            *string                                                            `facebook:"ad_account_currency"`
	CatalogID                    core.ID                                                            `facebook:"catalog_id"`
	ChildBusinessExternalID      *core.ID                                                           `facebook:"child_business_external_id"`
	CreditLimit                  *uint64                                                            `facebook:"credit_limit"`
	LineOfCreditID               *core.ID                                                           `facebook:"line_of_credit_id"`
	Name                         string                                                             `facebook:"name"`
	NoAdAccount                  *bool                                                              `facebook:"no_ad_account"`
	PageName                     *string                                                            `facebook:"page_name"`
	PageProfileImageURL          *string                                                            `facebook:"page_profile_image_url"`
	PartitionType                *enums.BusinessmanagedPartnerBusinessesPartitionTypeEnumParam      `facebook:"partition_type"`
	PartnerFacebookPageURL       *string                                                            `facebook:"partner_facebook_page_url"`
	PartnerRegistrationCountries *[]string                                                          `facebook:"partner_registration_countries"`
	SalesRepEmail                *string                                                            `facebook:"sales_rep_email"`
	SellerExternalWebsiteURL     string                                                             `facebook:"seller_external_website_url"`
	SellerTargetingCountries     []string                                                           `facebook:"seller_targeting_countries"`
	SkipPartnerPageCreation      *bool                                                              `facebook:"skip_partner_page_creation"`
	SurveyBusinessType           *enums.BusinessmanagedPartnerBusinessesSurveyBusinessTypeEnumParam `facebook:"survey_business_type"`
	SurveyNumAssets              *uint64                                                            `facebook:"survey_num_assets"`
	SurveyNumPeople              *uint64                                                            `facebook:"survey_num_people"`
	TimezoneID                   *enums.BusinessmanagedPartnerBusinessesTimezoneIDEnumParam         `facebook:"timezone_id"`
	Vertical                     enums.BusinessmanagedPartnerBusinessesVerticalEnumParam            `facebook:"vertical"`
	Extra                        core.Params                                                        `facebook:"-"`
}

func (params CreateBusinessManagedPartnerBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountCurrency != nil {
		out["ad_account_currency"] = *params.AdAccountCurrency
	}
	out["catalog_id"] = params.CatalogID
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.CreditLimit != nil {
		out["credit_limit"] = *params.CreditLimit
	}
	if params.LineOfCreditID != nil {
		out["line_of_credit_id"] = *params.LineOfCreditID
	}
	out["name"] = params.Name
	if params.NoAdAccount != nil {
		out["no_ad_account"] = *params.NoAdAccount
	}
	if params.PageName != nil {
		out["page_name"] = *params.PageName
	}
	if params.PageProfileImageURL != nil {
		out["page_profile_image_url"] = *params.PageProfileImageURL
	}
	if params.PartitionType != nil {
		out["partition_type"] = *params.PartitionType
	}
	if params.PartnerFacebookPageURL != nil {
		out["partner_facebook_page_url"] = *params.PartnerFacebookPageURL
	}
	if params.PartnerRegistrationCountries != nil {
		out["partner_registration_countries"] = *params.PartnerRegistrationCountries
	}
	if params.SalesRepEmail != nil {
		out["sales_rep_email"] = *params.SalesRepEmail
	}
	out["seller_external_website_url"] = params.SellerExternalWebsiteURL
	out["seller_targeting_countries"] = params.SellerTargetingCountries
	if params.SkipPartnerPageCreation != nil {
		out["skip_partner_page_creation"] = *params.SkipPartnerPageCreation
	}
	if params.SurveyBusinessType != nil {
		out["survey_business_type"] = *params.SurveyBusinessType
	}
	if params.SurveyNumAssets != nil {
		out["survey_num_assets"] = *params.SurveyNumAssets
	}
	if params.SurveyNumPeople != nil {
		out["survey_num_people"] = *params.SurveyNumPeople
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	out["vertical"] = params.Vertical
	return out
}

func CreateBusinessManagedPartnerBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedPartnerBusinessesParams) (*objects.ManagedPartnerBusiness, error) {
	var out objects.ManagedPartnerBusiness
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "managed_partner_businesses"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOnboardPartnersToMmLiteParams struct {
	SolutionID *core.ID    `facebook:"solution_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateBusinessOnboardPartnersToMmLiteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SolutionID != nil {
		out["solution_id"] = *params.SolutionID
	}
	return out
}

func CreateBusinessOnboardPartnersToMmLite(ctx context.Context, client *core.Client, id string, params CreateBusinessOnboardPartnersToMmLiteParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "onboard_partners_to_mm_lite"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOpenbridgeConfigurationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOpenbridgeConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params GetBusinessOpenbridgeConfigurationsParams) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	var out core.Cursor[objects.OpenBridgeConfiguration]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOpenbridgeConfigurationsParams struct {
	Active                         *bool                                                                          `facebook:"active"`
	BlockedEventTypes              *[]string                                                                      `facebook:"blocked_event_types"`
	BlockedWebsites                *[]string                                                                      `facebook:"blocked_websites"`
	CapiPublishingState            *enums.BusinessopenbridgeConfigurationsCapiPublishingStateEnumParam            `facebook:"capi_publishing_state"`
	CloudProvider                  *string                                                                        `facebook:"cloud_provider"`
	CloudRegion                    *string                                                                        `facebook:"cloud_region"`
	DestinationID                  *core.ID                                                                       `facebook:"destination_id"`
	Endpoint                       *string                                                                        `facebook:"endpoint"`
	EventEnrichmentAdvertiserState *enums.BusinessopenbridgeConfigurationsEventEnrichmentAdvertiserStateEnumParam `facebook:"event_enrichment_advertiser_state"`
	EventEnrichmentMetaState       *enums.BusinessopenbridgeConfigurationsEventEnrichmentMetaStateEnumParam       `facebook:"event_enrichment_meta_state"`
	EventEnrichmentState           *enums.BusinessopenbridgeConfigurationsEventEnrichmentStateEnumParam           `facebook:"event_enrichment_state"`
	FallbackDomain                 *string                                                                        `facebook:"fallback_domain"`
	HostBusinessID                 *core.ID                                                                       `facebook:"host_business_id"`
	InstanceID                     *core.ID                                                                       `facebook:"instance_id"`
	InstanceVersion                *string                                                                        `facebook:"instance_version"`
	IsSgwInstance                  *bool                                                                          `facebook:"is_sgw_instance"`
	IsSgwPixelFromMetaPixel        *bool                                                                          `facebook:"is_sgw_pixel_from_meta_pixel"`
	PartnerName                    *string                                                                        `facebook:"partner_name"`
	PixelID                        core.ID                                                                        `facebook:"pixel_id"`
	SgwAccountID                   *core.ID                                                                       `facebook:"sgw_account_id"`
	SgwInstanceURL                 *string                                                                        `facebook:"sgw_instance_url"`
	SgwPixelID                     *core.ID                                                                       `facebook:"sgw_pixel_id"`
	Extra                          core.Params                                                                    `facebook:"-"`
}

func (params CreateBusinessOpenbridgeConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Active != nil {
		out["active"] = *params.Active
	}
	if params.BlockedEventTypes != nil {
		out["blocked_event_types"] = *params.BlockedEventTypes
	}
	if params.BlockedWebsites != nil {
		out["blocked_websites"] = *params.BlockedWebsites
	}
	if params.CapiPublishingState != nil {
		out["capi_publishing_state"] = *params.CapiPublishingState
	}
	if params.CloudProvider != nil {
		out["cloud_provider"] = *params.CloudProvider
	}
	if params.CloudRegion != nil {
		out["cloud_region"] = *params.CloudRegion
	}
	if params.DestinationID != nil {
		out["destination_id"] = *params.DestinationID
	}
	if params.Endpoint != nil {
		out["endpoint"] = *params.Endpoint
	}
	if params.EventEnrichmentAdvertiserState != nil {
		out["event_enrichment_advertiser_state"] = *params.EventEnrichmentAdvertiserState
	}
	if params.EventEnrichmentMetaState != nil {
		out["event_enrichment_meta_state"] = *params.EventEnrichmentMetaState
	}
	if params.EventEnrichmentState != nil {
		out["event_enrichment_state"] = *params.EventEnrichmentState
	}
	if params.FallbackDomain != nil {
		out["fallback_domain"] = *params.FallbackDomain
	}
	if params.HostBusinessID != nil {
		out["host_business_id"] = *params.HostBusinessID
	}
	if params.InstanceID != nil {
		out["instance_id"] = *params.InstanceID
	}
	if params.InstanceVersion != nil {
		out["instance_version"] = *params.InstanceVersion
	}
	if params.IsSgwInstance != nil {
		out["is_sgw_instance"] = *params.IsSgwInstance
	}
	if params.IsSgwPixelFromMetaPixel != nil {
		out["is_sgw_pixel_from_meta_pixel"] = *params.IsSgwPixelFromMetaPixel
	}
	if params.PartnerName != nil {
		out["partner_name"] = *params.PartnerName
	}
	out["pixel_id"] = params.PixelID
	if params.SgwAccountID != nil {
		out["sgw_account_id"] = *params.SgwAccountID
	}
	if params.SgwInstanceURL != nil {
		out["sgw_instance_url"] = *params.SgwInstanceURL
	}
	if params.SgwPixelID != nil {
		out["sgw_pixel_id"] = *params.SgwPixelID
	}
	return out
}

func CreateBusinessOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params CreateBusinessOpenbridgeConfigurationsParams) (*objects.OpenBridgeConfiguration, error) {
	var out objects.OpenBridgeConfiguration
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedAdAccountsParams struct {
	IncludeSharedAdAccounts *bool       `facebook:"include_shared_ad_accounts"`
	SearchQuery             *string     `facebook:"search_query"`
	Extra                   core.Params `facebook:"-"`
}

func (params GetBusinessOwnedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeSharedAdAccounts != nil {
		out["include_shared_ad_accounts"] = *params.IncludeSharedAdAccounts
	}
	if params.SearchQuery != nil {
		out["search_query"] = *params.SearchQuery
	}
	return out
}

func GetBusinessOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOwnedAdAccountsParams struct {
	AdaccountID core.ID     `facebook:"adaccount_id"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateBusinessOwnedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["adaccount_id"] = params.AdaccountID
	return out
}

func CreateBusinessOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedAdAccountsParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owned_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedApps(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_apps"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOwnedAppsParams struct {
	AppID map[string]interface{} `facebook:"app_id"`
	Extra core.Params            `facebook:"-"`
}

func (params CreateBusinessOwnedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_id"] = params.AppID
	return out
}

func CreateBusinessOwnedApps(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedAppsParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owned_apps"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessOwnedBusinessesParams struct {
	ClientID core.ID     `facebook:"client_id"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteBusinessOwnedBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["client_id"] = params.ClientID
	return out
}

func DeleteBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessOwnedBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "owned_businesses"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedBusinessesParams struct {
	ChildBusinessExternalID *core.ID    `facebook:"child_business_external_id"`
	ClientUserID            *core.ID    `facebook:"client_user_id"`
	Extra                   core.Params `facebook:"-"`
}

func (params GetBusinessOwnedBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.ClientUserID != nil {
		out["client_user_id"] = *params.ClientUserID
	}
	return out
}

func GetBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedBusinessesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_businesses"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOwnedBusinessesParams struct {
	ChildBusinessExternalID *core.ID                                                    `facebook:"child_business_external_id"`
	Name                    string                                                      `facebook:"name"`
	PagePermittedTasks      *[]enums.BusinessownedBusinessesPagePermittedTasksEnumParam `facebook:"page_permitted_tasks"`
	SalesRepEmail           *string                                                     `facebook:"sales_rep_email"`
	SharedPageID            *core.ID                                                    `facebook:"shared_page_id"`
	ShouldGenerateName      *bool                                                       `facebook:"should_generate_name"`
	SurveyBusinessType      *enums.BusinessownedBusinessesSurveyBusinessTypeEnumParam   `facebook:"survey_business_type"`
	SurveyNumAssets         *uint64                                                     `facebook:"survey_num_assets"`
	SurveyNumPeople         *uint64                                                     `facebook:"survey_num_people"`
	TimezoneID              *enums.BusinessownedBusinessesTimezoneIDEnumParam           `facebook:"timezone_id"`
	Vertical                enums.BusinessownedBusinessesVerticalEnumParam              `facebook:"vertical"`
	Extra                   core.Params                                                 `facebook:"-"`
}

func (params CreateBusinessOwnedBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	out["name"] = params.Name
	if params.PagePermittedTasks != nil {
		out["page_permitted_tasks"] = *params.PagePermittedTasks
	}
	if params.SalesRepEmail != nil {
		out["sales_rep_email"] = *params.SalesRepEmail
	}
	if params.SharedPageID != nil {
		out["shared_page_id"] = *params.SharedPageID
	}
	if params.ShouldGenerateName != nil {
		out["should_generate_name"] = *params.ShouldGenerateName
	}
	if params.SurveyBusinessType != nil {
		out["survey_business_type"] = *params.SurveyBusinessType
	}
	if params.SurveyNumAssets != nil {
		out["survey_num_assets"] = *params.SurveyNumAssets
	}
	if params.SurveyNumPeople != nil {
		out["survey_num_people"] = *params.SurveyNumPeople
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	out["vertical"] = params.Vertical
	return out
}

func CreateBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedBusinessesParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owned_businesses"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedInstagramAssetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedInstagramAssetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedInstagramAssets(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedInstagramAssetsParams) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	var out core.Cursor[objects.InstagramBusinessAsset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_instagram_assets"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_offsite_signal_container_business_objects"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedPages(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_pages"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOwnedPagesParams struct {
	Code       *string     `facebook:"code"`
	EntryPoint *string     `facebook:"entry_point"`
	PageID     core.ID     `facebook:"page_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateBusinessOwnedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Code != nil {
		out["code"] = *params.Code
	}
	if params.EntryPoint != nil {
		out["entry_point"] = *params.EntryPoint
	}
	out["page_id"] = params.PageID
	return out
}

func CreateBusinessOwnedPages(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedPagesParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owned_pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedPixelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedPixels(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_pixels"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessOwnedProductCatalogsParams struct {
	AdditionalVerticalOption   *enums.BusinessownedProductCatalogsAdditionalVerticalOptionEnumParam `facebook:"additional_vertical_option"`
	BusinessMetadata           *map[string]interface{}                                              `facebook:"business_metadata"`
	CatalogSegmentFilter       *map[string]interface{}                                              `facebook:"catalog_segment_filter"`
	CatalogSegmentProductSetID *core.ID                                                             `facebook:"catalog_segment_product_set_id"`
	DaDisplaySettings          *map[string]interface{}                                              `facebook:"da_display_settings"`
	DestinationCatalogSettings *map[string]interface{}                                              `facebook:"destination_catalog_settings"`
	FlightCatalogSettings      *map[string]interface{}                                              `facebook:"flight_catalog_settings"`
	Name                       string                                                               `facebook:"name"`
	ParentCatalogID            *core.ID                                                             `facebook:"parent_catalog_id"`
	PartnerIntegration         *map[string]interface{}                                              `facebook:"partner_integration"`
	StoreCatalogSettings       *map[string]interface{}                                              `facebook:"store_catalog_settings"`
	Vertical                   *enums.BusinessownedProductCatalogsVerticalEnumParam                 `facebook:"vertical"`
	Extra                      core.Params                                                          `facebook:"-"`
}

func (params CreateBusinessOwnedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdditionalVerticalOption != nil {
		out["additional_vertical_option"] = *params.AdditionalVerticalOption
	}
	if params.BusinessMetadata != nil {
		out["business_metadata"] = *params.BusinessMetadata
	}
	if params.CatalogSegmentFilter != nil {
		out["catalog_segment_filter"] = *params.CatalogSegmentFilter
	}
	if params.CatalogSegmentProductSetID != nil {
		out["catalog_segment_product_set_id"] = *params.CatalogSegmentProductSetID
	}
	if params.DaDisplaySettings != nil {
		out["da_display_settings"] = *params.DaDisplaySettings
	}
	if params.DestinationCatalogSettings != nil {
		out["destination_catalog_settings"] = *params.DestinationCatalogSettings
	}
	if params.FlightCatalogSettings != nil {
		out["flight_catalog_settings"] = *params.FlightCatalogSettings
	}
	out["name"] = params.Name
	if params.ParentCatalogID != nil {
		out["parent_catalog_id"] = *params.ParentCatalogID
	}
	if params.PartnerIntegration != nil {
		out["partner_integration"] = *params.PartnerIntegration
	}
	if params.StoreCatalogSettings != nil {
		out["store_catalog_settings"] = *params.StoreCatalogSettings
	}
	if params.Vertical != nil {
		out["vertical"] = *params.Vertical
	}
	return out
}

func CreateBusinessOwnedProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedProductCatalogsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "owned_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessOwnedWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "owned_whatsapp_business_accounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessPagesParams struct {
	PageID core.ID     `facebook:"page_id"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteBusinessPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["page_id"] = params.PageID
	return out
}

func DeleteBusinessPages(ctx context.Context, client *core.Client, id string, params DeleteBusinessPagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPartnerAccountLinkingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPartnerAccountLinkingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPartnerAccountLinking(ctx context.Context, client *core.Client, id string, params GetBusinessPartnerAccountLinkingParams) (*core.Cursor[objects.PartnerAccountLinking], error) {
	var out core.Cursor[objects.PartnerAccountLinking]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "partner_account_linking"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessPartnerPremiumOptionsParams struct {
	CatalogSegmentID                  *core.ID               `facebook:"catalog_segment_id"`
	EnableBasketInsight               bool                   `facebook:"enable_basket_insight"`
	EnableExtendedAudienceRetargeting bool                   `facebook:"enable_extended_audience_retargeting"`
	PartnerBusinessID                 core.ID                `facebook:"partner_business_id"`
	RetailerCustomAudienceConfig      map[string]interface{} `facebook:"retailer_custom_audience_config"`
	VendorID                          *core.ID               `facebook:"vendor_id"`
	Extra                             core.Params            `facebook:"-"`
}

func (params CreateBusinessPartnerPremiumOptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CatalogSegmentID != nil {
		out["catalog_segment_id"] = *params.CatalogSegmentID
	}
	out["enable_basket_insight"] = params.EnableBasketInsight
	out["enable_extended_audience_retargeting"] = params.EnableExtendedAudienceRetargeting
	out["partner_business_id"] = params.PartnerBusinessID
	out["retailer_custom_audience_config"] = params.RetailerCustomAudienceConfig
	if params.VendorID != nil {
		out["vendor_id"] = *params.VendorID
	}
	return out
}

func CreateBusinessPartnerPremiumOptions(ctx context.Context, client *core.Client, id string, params CreateBusinessPartnerPremiumOptionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "partner_premium_options"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPassbackAttributionMetadataConfigsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPassbackAttributionMetadataConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPassbackAttributionMetadataConfigs(ctx context.Context, client *core.Client, id string, params GetBusinessPassbackAttributionMetadataConfigsParams) (*core.Cursor[objects.SignalsAttributionMetadataConfig], error) {
	var out core.Cursor[objects.SignalsAttributionMetadataConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "passback_attribution_metadata_configs"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingClientAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingClientAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingClientAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientAdAccountsParams) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	var out core.Cursor[objects.BusinessAdAccountRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_client_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingClientAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingClientAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingClientApps(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientAppsParams) (*core.Cursor[objects.BusinessApplicationRequest], error) {
	var out core.Cursor[objects.BusinessApplicationRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_client_apps"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingClientPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingClientPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingClientPages(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientPagesParams) (*core.Cursor[objects.BusinessPageRequest], error) {
	var out core.Cursor[objects.BusinessPageRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_client_pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingOwnedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingOwnedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessPendingOwnedAdAccountsParams) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	var out core.Cursor[objects.BusinessAdAccountRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_owned_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingOwnedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingOwnedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingOwnedPages(ctx context.Context, client *core.Client, id string, params GetBusinessPendingOwnedPagesParams) (*core.Cursor[objects.BusinessPageRequest], error) {
	var out core.Cursor[objects.BusinessPageRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_owned_pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessPendingSharedOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_shared_offsite_signal_container_business_objects"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPendingUsersParams struct {
	Email *string     `facebook:"email"`
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessPendingUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	return out
}

func GetBusinessPendingUsers(ctx context.Context, client *core.Client, id string, params GetBusinessPendingUsersParams) (*core.Cursor[objects.BusinessRoleRequest], error) {
	var out core.Cursor[objects.BusinessRoleRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pending_users"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPictureParams struct {
	Height   *int                                `facebook:"height"`
	Redirect *bool                               `facebook:"redirect"`
	Type     *enums.BusinesspictureTypeEnumParam `facebook:"type"`
	Width    *int                                `facebook:"width"`
	Extra    core.Params                         `facebook:"-"`
}

func (params GetBusinessPictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Redirect != nil {
		out["redirect"] = *params.Redirect
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetBusinessPicture(ctx context.Context, client *core.Client, id string, params GetBusinessPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	var out core.Cursor[objects.ProfilePictureSource]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessPixelTosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateBusinessPixelTosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateBusinessPixelTos(ctx context.Context, client *core.Client, id string, params CreateBusinessPixelTosParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "pixel_tos"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessPreverifiedNumbersParams struct {
	CodeVerificationStatus *enums.BusinesspreverifiedNumbersCodeVerificationStatusEnumParam `facebook:"code_verification_status"`
	PhoneNumber            *string                                                          `facebook:"phone_number"`
	Extra                  core.Params                                                      `facebook:"-"`
}

func (params GetBusinessPreverifiedNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CodeVerificationStatus != nil {
		out["code_verification_status"] = *params.CodeVerificationStatus
	}
	if params.PhoneNumber != nil {
		out["phone_number"] = *params.PhoneNumber
	}
	return out
}

func GetBusinessPreverifiedNumbers(ctx context.Context, client *core.Client, id string, params GetBusinessPreverifiedNumbersParams) (*core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber], error) {
	var out core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "preverified_numbers"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessReceivedAudienceSharingRequestsParams struct {
	InitiatorID   *core.ID                                                             `facebook:"initiator_id"`
	RequestStatus *enums.BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam `facebook:"request_status"`
	Extra         core.Params                                                          `facebook:"-"`
}

func (params GetBusinessReceivedAudienceSharingRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.InitiatorID != nil {
		out["initiator_id"] = *params.InitiatorID
	}
	if params.RequestStatus != nil {
		out["request_status"] = *params.RequestStatus
	}
	return out
}

func GetBusinessReceivedAudienceSharingRequests(ctx context.Context, client *core.Client, id string, params GetBusinessReceivedAudienceSharingRequestsParams) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "received_audience_sharing_requests"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessResellerGuidancesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessResellerGuidancesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessResellerGuidances(ctx context.Context, client *core.Client, id string, params GetBusinessResellerGuidancesParams) (*core.Cursor[objects.ResellerGuidance], error) {
	var out core.Cursor[objects.ResellerGuidance]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "reseller_guidances"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams struct {
	EndBusinessID *core.ID    `facebook:"end_business_id"`
	Extra         core.Params `facebook:"-"`
}

func (params GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndBusinessID != nil {
		out["end_business_id"] = *params.EndBusinessID
	}
	return out
}

func GetBusinessSelfCertifiedWhatsappBusinessSubmissions(ctx context.Context, client *core.Client, id string, params GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams) (*core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission], error) {
	var out core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "self_certified_whatsapp_business_submissions"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessSelfCertifyWhatsappBusinessParams struct {
	AverageMonthlyRevenueSpendWithPartner *map[string]interface{}                                             `facebook:"average_monthly_revenue_spend_with_partner"`
	BusinessDocuments                     []core.FileParam                                                    `facebook:"business_documents"`
	BusinessVertical                      *enums.BusinessselfCertifyWhatsappBusinessBusinessVerticalEnumParam `facebook:"business_vertical"`
	EndBusinessAddress                    *map[string]interface{}                                             `facebook:"end_business_address"`
	EndBusinessID                         core.ID                                                             `facebook:"end_business_id"`
	EndBusinessLegalName                  *string                                                             `facebook:"end_business_legal_name"`
	EndBusinessTradeNames                 *[]string                                                           `facebook:"end_business_trade_names"`
	EndBusinessWebsite                    *string                                                             `facebook:"end_business_website"`
	NumBillingCyclesWithPartner           *uint64                                                             `facebook:"num_billing_cycles_with_partner"`
	Extra                                 core.Params                                                         `facebook:"-"`
}

func (params CreateBusinessSelfCertifyWhatsappBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AverageMonthlyRevenueSpendWithPartner != nil {
		out["average_monthly_revenue_spend_with_partner"] = *params.AverageMonthlyRevenueSpendWithPartner
	}
	out["business_documents"] = params.BusinessDocuments
	if params.BusinessVertical != nil {
		out["business_vertical"] = *params.BusinessVertical
	}
	if params.EndBusinessAddress != nil {
		out["end_business_address"] = *params.EndBusinessAddress
	}
	out["end_business_id"] = params.EndBusinessID
	if params.EndBusinessLegalName != nil {
		out["end_business_legal_name"] = *params.EndBusinessLegalName
	}
	if params.EndBusinessTradeNames != nil {
		out["end_business_trade_names"] = *params.EndBusinessTradeNames
	}
	if params.EndBusinessWebsite != nil {
		out["end_business_website"] = *params.EndBusinessWebsite
	}
	if params.NumBillingCyclesWithPartner != nil {
		out["num_billing_cycles_with_partner"] = *params.NumBillingCyclesWithPartner
	}
	return out
}

func CreateBusinessSelfCertifyWhatsappBusiness(ctx context.Context, client *core.Client, id string, params CreateBusinessSelfCertifyWhatsappBusinessParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "self_certify_whatsapp_business"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessSharePreverifiedNumbersParams struct {
	PartnerBusinessID core.ID     `facebook:"partner_business_id"`
	PreverifiedID     core.ID     `facebook:"preverified_id"`
	Extra             core.Params `facebook:"-"`
}

func (params DeleteBusinessSharePreverifiedNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["partner_business_id"] = params.PartnerBusinessID
	out["preverified_id"] = params.PreverifiedID
	return out
}

func DeleteBusinessSharePreverifiedNumbers(ctx context.Context, client *core.Client, id string, params DeleteBusinessSharePreverifiedNumbersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "share_preverified_numbers"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessSharePreverifiedNumbersParams struct {
	PartnerBusinessID core.ID     `facebook:"partner_business_id"`
	PreverifiedID     core.ID     `facebook:"preverified_id"`
	Extra             core.Params `facebook:"-"`
}

func (params CreateBusinessSharePreverifiedNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["partner_business_id"] = params.PartnerBusinessID
	out["preverified_id"] = params.PreverifiedID
	return out
}

func CreateBusinessSharePreverifiedNumbers(ctx context.Context, client *core.Client, id string, params CreateBusinessSharePreverifiedNumbersParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "share_preverified_numbers"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessSystemUserAccessTokensParams struct {
	Asset                    *[]uint64             `facebook:"asset"`
	FetchOnly                *bool                 `facebook:"fetch_only"`
	Scope                    *[]objects.Permission `facebook:"scope"`
	SetTokenExpiresInX60Days *bool                 `facebook:"set_token_expires_in_60_days"`
	SystemUserID             *core.ID              `facebook:"system_user_id"`
	Extra                    core.Params           `facebook:"-"`
}

func (params CreateBusinessSystemUserAccessTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Asset != nil {
		out["asset"] = *params.Asset
	}
	if params.FetchOnly != nil {
		out["fetch_only"] = *params.FetchOnly
	}
	if params.Scope != nil {
		out["scope"] = *params.Scope
	}
	if params.SetTokenExpiresInX60Days != nil {
		out["set_token_expires_in_60_days"] = *params.SetTokenExpiresInX60Days
	}
	if params.SystemUserID != nil {
		out["system_user_id"] = *params.SystemUserID
	}
	return out
}

func CreateBusinessSystemUserAccessTokens(ctx context.Context, client *core.Client, id string, params CreateBusinessSystemUserAccessTokensParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "system_user_access_tokens"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessSystemUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessSystemUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessSystemUsers(ctx context.Context, client *core.Client, id string, params GetBusinessSystemUsersParams) (*core.Cursor[objects.SystemUser], error) {
	var out core.Cursor[objects.SystemUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "system_users"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessSystemUsersParams struct {
	Name         string                                  `facebook:"name"`
	Role         *enums.BusinesssystemUsersRoleEnumParam `facebook:"role"`
	SystemUserID *core.ID                                `facebook:"system_user_id"`
	Extra        core.Params                             `facebook:"-"`
}

func (params CreateBusinessSystemUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	if params.Role != nil {
		out["role"] = *params.Role
	}
	if params.SystemUserID != nil {
		out["system_user_id"] = *params.SystemUserID
	}
	return out
}

func CreateBusinessSystemUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessSystemUsersParams) (*objects.SystemUser, error) {
	var out objects.SystemUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "system_users"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessThirdPartyMeasurementReportDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessThirdPartyMeasurementReportDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessThirdPartyMeasurementReportDataset(ctx context.Context, client *core.Client, id string, params GetBusinessThirdPartyMeasurementReportDatasetParams) (*core.Cursor[objects.ThirdPartyMeasurementReportDataset], error) {
	var out core.Cursor[objects.ThirdPartyMeasurementReportDataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "third_party_measurement_report_dataset"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessVideosParams struct {
	AdPlacementsValidationOnly    *bool                                                  `facebook:"ad_placements_validation_only"`
	ApplicationID                 *core.ID                                               `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                               `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                                `facebook:"audio_story_wave_animation_handle"`
	ChunkSessionID                *core.ID                                               `facebook:"chunk_session_id"`
	ComposerEntryPicker           *string                                                `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                                `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                                `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                                `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                               `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                                `facebook:"composer_source_surface"`
	ComposerType                  *string                                                `facebook:"composer_type"`
	ContainerType                 *enums.BusinessvideosContainerTypeEnumParam            `facebook:"container_type"`
	ContentCategory               *enums.BusinessvideosContentCategoryEnumParam          `facebook:"content_category"`
	CreativeFolderID              core.ID                                                `facebook:"creative_folder_id"`
	CreativeTools                 *string                                                `facebook:"creative_tools"`
	Description                   *string                                                `facebook:"description"`
	EditDescriptionSpec           *map[string]interface{}                                `facebook:"edit_description_spec"`
	Embeddable                    *bool                                                  `facebook:"embeddable"`
	EndOffset                     *uint64                                                `facebook:"end_offset"`
	FbuploaderVideoFileChunk      *string                                                `facebook:"fbuploader_video_file_chunk"`
	FileSize                      *uint64                                                `facebook:"file_size"`
	FileURL                       *string                                                `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                                  `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.BusinessvideosFormattingEnumParam               `facebook:"formatting"`
	Fov                           *uint64                                                `facebook:"fov"`
	FrontZRotation                *float64                                               `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                               `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                               `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                            `facebook:"guide"`
	GuideEnabled                  *bool                                                  `facebook:"guide_enabled"`
	InitialHeading                *uint64                                                `facebook:"initial_heading"`
	InitialPitch                  *uint64                                                `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                                `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                                  `facebook:"is_boost_intended"`
	IsGroupLinkingPost            *bool                                                  `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                                  `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                                  `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                               `facebook:"location_source_id"`
	OgActionTypeID                *core.ID                                               `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                               `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                               `facebook:"og_object_id"`
	OgPhrase                      *string                                                `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                                `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                                `facebook:"original_fov"`
	OriginalProjectionType        *enums.BusinessvideosOriginalProjectionTypeEnumParam   `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                                `facebook:"partnership_ad_ad_code"`
	PublishEventID                *core.ID                                               `facebook:"publish_event_id"`
	ReferencedStickerID           *core.ID                                               `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                               `facebook:"replace_video_id"`
	SelectedAudioSpec             *map[string]interface{}                                `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                                `facebook:"slideshow_spec"`
	Source                        *string                                                `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                               `facebook:"source_instagram_media_id"`
	Spherical                     *bool                                                  `facebook:"spherical"`
	StartOffset                   *uint64                                                `facebook:"start_offset"`
	SwapMode                      *enums.BusinessvideosSwapModeEnumParam                 `facebook:"swap_mode"`
	TextFormatMetadata            *string                                                `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                        `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                                `facebook:"time_since_original_post"`
	Title                         *string                                                `facebook:"title"`
	TranscodeSettingProperties    *string                                                `facebook:"transcode_setting_properties"`
	UnpublishedContentType        *enums.BusinessvideosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.BusinessvideosUploadPhaseEnumParam              `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                               `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                                `facebook:"upload_setting_properties"`
	ValidationAdPlacements        *[]enums.BusinessvideosValidationAdPlacementsEnumParam `facebook:"validation_ad_placements"`
	VideoFileChunk                *string                                                `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                                `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                                `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                               `facebook:"waterfall_id"`
	Extra                         core.Params                                            `facebook:"-"`
}

func (params CreateBusinessVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdPlacementsValidationOnly != nil {
		out["ad_placements_validation_only"] = *params.AdPlacementsValidationOnly
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.AudioStoryWaveAnimationHandle != nil {
		out["audio_story_wave_animation_handle"] = *params.AudioStoryWaveAnimationHandle
	}
	if params.ChunkSessionID != nil {
		out["chunk_session_id"] = *params.ChunkSessionID
	}
	if params.ComposerEntryPicker != nil {
		out["composer_entry_picker"] = *params.ComposerEntryPicker
	}
	if params.ComposerEntryPoint != nil {
		out["composer_entry_point"] = *params.ComposerEntryPoint
	}
	if params.ComposerEntryTime != nil {
		out["composer_entry_time"] = *params.ComposerEntryTime
	}
	if params.ComposerSessionEventsLog != nil {
		out["composer_session_events_log"] = *params.ComposerSessionEventsLog
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.ComposerSourceSurface != nil {
		out["composer_source_surface"] = *params.ComposerSourceSurface
	}
	if params.ComposerType != nil {
		out["composer_type"] = *params.ComposerType
	}
	if params.ContainerType != nil {
		out["container_type"] = *params.ContainerType
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	out["creative_folder_id"] = params.CreativeFolderID
	if params.CreativeTools != nil {
		out["creative_tools"] = *params.CreativeTools
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EditDescriptionSpec != nil {
		out["edit_description_spec"] = *params.EditDescriptionSpec
	}
	if params.Embeddable != nil {
		out["embeddable"] = *params.Embeddable
	}
	if params.EndOffset != nil {
		out["end_offset"] = *params.EndOffset
	}
	if params.FbuploaderVideoFileChunk != nil {
		out["fbuploader_video_file_chunk"] = *params.FbuploaderVideoFileChunk
	}
	if params.FileSize != nil {
		out["file_size"] = *params.FileSize
	}
	if params.FileURL != nil {
		out["file_url"] = *params.FileURL
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.Fov != nil {
		out["fov"] = *params.Fov
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Guide != nil {
		out["guide"] = *params.Guide
	}
	if params.GuideEnabled != nil {
		out["guide_enabled"] = *params.GuideEnabled
	}
	if params.InitialHeading != nil {
		out["initial_heading"] = *params.InitialHeading
	}
	if params.InitialPitch != nil {
		out["initial_pitch"] = *params.InitialPitch
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPartnershipAd != nil {
		out["is_partnership_ad"] = *params.IsPartnershipAd
	}
	if params.IsVoiceClip != nil {
		out["is_voice_clip"] = *params.IsVoiceClip
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgObjectID != nil {
		out["og_object_id"] = *params.OgObjectID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.OriginalProjectionType != nil {
		out["original_projection_type"] = *params.OriginalProjectionType
	}
	if params.PartnershipAdAdCode != nil {
		out["partnership_ad_ad_code"] = *params.PartnershipAdAdCode
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.ReferencedStickerID != nil {
		out["referenced_sticker_id"] = *params.ReferencedStickerID
	}
	if params.ReplaceVideoID != nil {
		out["replace_video_id"] = *params.ReplaceVideoID
	}
	if params.SelectedAudioSpec != nil {
		out["selected_audio_spec"] = *params.SelectedAudioSpec
	}
	if params.SlideshowSpec != nil {
		out["slideshow_spec"] = *params.SlideshowSpec
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SourceInstagramMediaID != nil {
		out["source_instagram_media_id"] = *params.SourceInstagramMediaID
	}
	if params.Spherical != nil {
		out["spherical"] = *params.Spherical
	}
	if params.StartOffset != nil {
		out["start_offset"] = *params.StartOffset
	}
	if params.SwapMode != nil {
		out["swap_mode"] = *params.SwapMode
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.Thumb != nil {
		out["thumb"] = *params.Thumb
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TranscodeSettingProperties != nil {
		out["transcode_setting_properties"] = *params.TranscodeSettingProperties
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UploadPhase != nil {
		out["upload_phase"] = *params.UploadPhase
	}
	if params.UploadSessionID != nil {
		out["upload_session_id"] = *params.UploadSessionID
	}
	if params.UploadSettingProperties != nil {
		out["upload_setting_properties"] = *params.UploadSettingProperties
	}
	if params.ValidationAdPlacements != nil {
		out["validation_ad_placements"] = *params.ValidationAdPlacements
	}
	if params.VideoFileChunk != nil {
		out["video_file_chunk"] = *params.VideoFileChunk
	}
	if params.VideoIDOriginal != nil {
		out["video_id_original"] = *params.VideoIDOriginal
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.WaterfallID != nil {
		out["waterfall_id"] = *params.WaterfallID
	}
	return out
}

func CreateBusinessVideos(ctx context.Context, client *core.Client, id string, params CreateBusinessVideosParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusiness(ctx context.Context, client *core.Client, id string, params GetBusinessParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateBusinessParams struct {
	EntryPoint    *string                      `facebook:"entry_point"`
	Name          *string                      `facebook:"name"`
	PrimaryPage   *string                      `facebook:"primary_page"`
	TimezoneID    *core.ID                     `facebook:"timezone_id"`
	TwoFactorType *enums.BusinessTwoFactorType `facebook:"two_factor_type"`
	Vertical      *enums.BusinessVertical      `facebook:"vertical"`
	Extra         core.Params                  `facebook:"-"`
}

func (params UpdateBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EntryPoint != nil {
		out["entry_point"] = *params.EntryPoint
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.PrimaryPage != nil {
		out["primary_page"] = *params.PrimaryPage
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	if params.TwoFactorType != nil {
		out["two_factor_type"] = *params.TwoFactorType
	}
	if params.Vertical != nil {
		out["vertical"] = *params.Vertical
	}
	return out
}

func UpdateBusiness(ctx context.Context, client *core.Client, id string, params UpdateBusinessParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
