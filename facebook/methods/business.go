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

func CreateBusinessAccessTokenBatchCall(id string, params CreateBusinessAccessTokenParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "access_token"), params.ToParams(), options...)
}

func NewCreateBusinessAccessTokenBatchRequest(id string, params CreateBusinessAccessTokenParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessAccessTokenBatchCall(id, params, options...))
}

func DecodeCreateBusinessAccessTokenBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAccessToken(ctx context.Context, client *core.Client, id string, params CreateBusinessAccessTokenParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessAccessTokenBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdAccountInfosBatchCall(id string, params GetBusinessAdAccountInfosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_account_infos"), params.ToParams(), options...)
}

func NewGetBusinessAdAccountInfosBatchRequest(id string, params GetBusinessAdAccountInfosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ALMAdAccountInfo]] {
	return core.NewBatchRequest[core.Cursor[objects.ALMAdAccountInfo]](GetBusinessAdAccountInfosBatchCall(id, params, options...))
}

func DecodeGetBusinessAdAccountInfosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ALMAdAccountInfo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ALMAdAccountInfo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdAccountInfos(ctx context.Context, client *core.Client, id string, params GetBusinessAdAccountInfosParams) (*core.Cursor[objects.ALMAdAccountInfo], error) {
	var out core.Cursor[objects.ALMAdAccountInfo]
	call := GetBusinessAdAccountInfosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessAdAccountsBatchCall(id string, params DeleteBusinessAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "ad_accounts"), params.ToParams(), options...)
}

func NewDeleteBusinessAdAccountsBatchRequest(id string, params DeleteBusinessAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAdAccountsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAdAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAdAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAdAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdCustomDerivedMetricsBatchCall(id string, params GetBusinessAdCustomDerivedMetricsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_custom_derived_metrics"), params.ToParams(), options...)
}

func NewGetBusinessAdCustomDerivedMetricsBatchRequest(id string, params GetBusinessAdCustomDerivedMetricsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCustomDerivedMetrics]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCustomDerivedMetrics]](GetBusinessAdCustomDerivedMetricsBatchCall(id, params, options...))
}

func DecodeGetBusinessAdCustomDerivedMetricsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCustomDerivedMetrics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCustomDerivedMetrics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdCustomDerivedMetrics(ctx context.Context, client *core.Client, id string, params GetBusinessAdCustomDerivedMetricsParams) (*core.Cursor[objects.AdCustomDerivedMetrics], error) {
	var out core.Cursor[objects.AdCustomDerivedMetrics]
	call := GetBusinessAdCustomDerivedMetricsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdReviewRequestsBatchCall(id string, params CreateBusinessAdReviewRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_review_requests"), params.ToParams(), options...)
}

func NewCreateBusinessAdReviewRequestsBatchRequest(id string, params CreateBusinessAdReviewRequestsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateBusinessAdReviewRequestsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdReviewRequestsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateBusinessAdReviewRequests(ctx context.Context, client *core.Client, id string, params CreateBusinessAdReviewRequestsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateBusinessAdReviewRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdStudiesBatchCall(id string, params GetBusinessAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewGetBusinessAdStudiesBatchRequest(id string, params GetBusinessAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudy]](GetBusinessAdStudiesBatchCall(id, params, options...))
}

func DecodeGetBusinessAdStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdStudies(ctx context.Context, client *core.Client, id string, params GetBusinessAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	call := GetBusinessAdStudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdStudiesBatchCall(id string, params CreateBusinessAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewCreateBusinessAdStudiesBatchRequest(id string, params CreateBusinessAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudy] {
	return core.NewBatchRequest[objects.AdStudy](CreateBusinessAdStudiesBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdStudiesBatchResponse(response *core.BatchResponse) (*objects.AdStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdStudies(ctx context.Context, client *core.Client, id string, params CreateBusinessAdStudiesParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	call := CreateBusinessAdStudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdaccountBatchCall(id string, params CreateBusinessAdaccountParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adaccount"), params.ToParams(), options...)
}

func NewCreateBusinessAdaccountBatchRequest(id string, params CreateBusinessAdaccountParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateBusinessAdaccountBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdaccountBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdaccount(ctx context.Context, client *core.Client, id string, params CreateBusinessAdaccountParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateBusinessAdaccountBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAddPhoneNumbersBatchCall(id string, params CreateBusinessAddPhoneNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "add_phone_numbers"), params.ToParams(), options...)
}

func NewCreateBusinessAddPhoneNumbersBatchRequest(id string, params CreateBusinessAddPhoneNumbersParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessAddPhoneNumbersBatchCall(id, params, options...))
}

func DecodeCreateBusinessAddPhoneNumbersBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAddPhoneNumbers(ctx context.Context, client *core.Client, id string, params CreateBusinessAddPhoneNumbersParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessAddPhoneNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdnetworkApplicationsBatchCall(id string, params CreateBusinessAdnetworkApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adnetwork_applications"), params.ToParams(), options...)
}

func NewCreateBusinessAdnetworkApplicationsBatchRequest(id string, params CreateBusinessAdnetworkApplicationsParams, options ...core.BatchOption) *core.BatchRequest[objects.Application] {
	return core.NewBatchRequest[objects.Application](CreateBusinessAdnetworkApplicationsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdnetworkApplicationsBatchResponse(response *core.BatchResponse) (*objects.Application, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Application
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdnetworkApplications(ctx context.Context, client *core.Client, id string, params CreateBusinessAdnetworkApplicationsParams) (*objects.Application, error) {
	var out objects.Application
	call := CreateBusinessAdnetworkApplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdnetworkanalyticsBatchCall(id string, params GetBusinessAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewGetBusinessAdnetworkanalyticsBatchRequest(id string, params GetBusinessAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]](GetBusinessAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeGetBusinessAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params GetBusinessAdnetworkanalyticsParams) (*core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsSyncQueryResult]
	call := GetBusinessAdnetworkanalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdnetworkanalyticsBatchCall(id string, params CreateBusinessAdnetworkanalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adnetworkanalytics"), params.ToParams(), options...)
}

func NewCreateBusinessAdnetworkanalyticsBatchRequest(id string, params CreateBusinessAdnetworkanalyticsParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessAdnetworkanalyticsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdnetworkanalyticsBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdnetworkanalytics(ctx context.Context, client *core.Client, id string, params CreateBusinessAdnetworkanalyticsParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessAdnetworkanalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdnetworkanalyticsResultsBatchCall(id string, params GetBusinessAdnetworkanalyticsResultsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adnetworkanalytics_results"), params.ToParams(), options...)
}

func NewGetBusinessAdnetworkanalyticsResultsBatchRequest(id string, params GetBusinessAdnetworkanalyticsResultsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]](GetBusinessAdnetworkanalyticsResultsBatchCall(id, params, options...))
}

func DecodeGetBusinessAdnetworkanalyticsResultsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdnetworkanalyticsResults(ctx context.Context, client *core.Client, id string, params GetBusinessAdnetworkanalyticsResultsParams) (*core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult], error) {
	var out core.Cursor[objects.AdNetworkAnalyticsAsyncQueryResult]
	call := GetBusinessAdnetworkanalyticsResultsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdsDatasetBatchCall(id string, params GetBusinessAdsDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_dataset"), params.ToParams(), options...)
}

func NewGetBusinessAdsDatasetBatchRequest(id string, params GetBusinessAdsDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsDataset]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsDataset]](GetBusinessAdsDatasetBatchCall(id, params, options...))
}

func DecodeGetBusinessAdsDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsDataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsDataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdsDataset(ctx context.Context, client *core.Client, id string, params GetBusinessAdsDatasetParams) (*core.Cursor[objects.AdsDataset], error) {
	var out core.Cursor[objects.AdsDataset]
	call := GetBusinessAdsDatasetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdsDatasetBatchCall(id string, params CreateBusinessAdsDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ads_dataset"), params.ToParams(), options...)
}

func NewCreateBusinessAdsDatasetBatchRequest(id string, params CreateBusinessAdsDatasetParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessAdsDatasetBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdsDatasetBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdsDataset(ctx context.Context, client *core.Client, id string, params CreateBusinessAdsDatasetParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessAdsDatasetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdsReportingMmmReportsBatchCall(id string, params GetBusinessAdsReportingMmmReportsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_reports"), params.ToParams(), options...)
}

func NewGetBusinessAdsReportingMmmReportsBatchRequest(id string, params GetBusinessAdsReportingMmmReportsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsReportBuilderMMMReport]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsReportBuilderMMMReport]](GetBusinessAdsReportingMmmReportsBatchCall(id, params, options...))
}

func DecodeGetBusinessAdsReportingMmmReportsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsReportBuilderMMMReport], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsReportBuilderMMMReport]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdsReportingMmmReports(ctx context.Context, client *core.Client, id string, params GetBusinessAdsReportingMmmReportsParams) (*core.Cursor[objects.AdsReportBuilderMMMReport], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReport]
	call := GetBusinessAdsReportingMmmReportsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdsReportingMmmSchedulersBatchCall(id string, params GetBusinessAdsReportingMmmSchedulersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_schedulers"), params.ToParams(), options...)
}

func NewGetBusinessAdsReportingMmmSchedulersBatchRequest(id string, params GetBusinessAdsReportingMmmSchedulersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsReportBuilderMMMReportScheduler]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsReportBuilderMMMReportScheduler]](GetBusinessAdsReportingMmmSchedulersBatchCall(id, params, options...))
}

func DecodeGetBusinessAdsReportingMmmSchedulersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsReportBuilderMMMReportScheduler], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsReportBuilderMMMReportScheduler]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdsReportingMmmSchedulers(ctx context.Context, client *core.Client, id string, params GetBusinessAdsReportingMmmSchedulersParams) (*core.Cursor[objects.AdsReportBuilderMMMReportScheduler], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReportScheduler]
	call := GetBusinessAdsReportingMmmSchedulersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAdspixelsBatchCall(id string, params GetBusinessAdspixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adspixels"), params.ToParams(), options...)
}

func NewGetBusinessAdspixelsBatchRequest(id string, params GetBusinessAdspixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetBusinessAdspixelsBatchCall(id, params, options...))
}

func DecodeGetBusinessAdspixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAdspixels(ctx context.Context, client *core.Client, id string, params GetBusinessAdspixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetBusinessAdspixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessAdspixelsBatchCall(id string, params CreateBusinessAdspixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adspixels"), params.ToParams(), options...)
}

func NewCreateBusinessAdspixelsBatchRequest(id string, params CreateBusinessAdspixelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](CreateBusinessAdspixelsBatchCall(id, params, options...))
}

func DecodeCreateBusinessAdspixelsBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessAdspixels(ctx context.Context, client *core.Client, id string, params CreateBusinessAdspixelsParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	call := CreateBusinessAdspixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessAgenciesBatchCall(id string, params DeleteBusinessAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewDeleteBusinessAgenciesBatchRequest(id string, params DeleteBusinessAgenciesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessAgenciesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessAgenciesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessAgencies(ctx context.Context, client *core.Client, id string, params DeleteBusinessAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessAgenciesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAgenciesBatchCall(id string, params GetBusinessAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetBusinessAgenciesBatchRequest(id string, params GetBusinessAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetBusinessAgenciesBatchCall(id, params, options...))
}

func DecodeGetBusinessAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetBusinessAgencies(ctx context.Context, client *core.Client, id string, params GetBusinessAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	call := GetBusinessAgenciesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessAnPlacementsBatchCall(id string, params GetBusinessAnPlacementsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "an_placements"), params.ToParams(), options...)
}

func NewGetBusinessAnPlacementsBatchRequest(id string, params GetBusinessAnPlacementsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPlacement]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPlacement]](GetBusinessAnPlacementsBatchCall(id, params, options...))
}

func DecodeGetBusinessAnPlacementsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPlacement], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPlacement]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAnPlacements(ctx context.Context, client *core.Client, id string, params GetBusinessAnPlacementsParams) (*core.Cursor[objects.AdPlacement], error) {
	var out core.Cursor[objects.AdPlacement]
	call := GetBusinessAnPlacementsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessBlockListDraftsBatchCall(id string, params CreateBusinessBlockListDraftsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "block_list_drafts"), params.ToParams(), options...)
}

func NewCreateBusinessBlockListDraftsBatchRequest(id string, params CreateBusinessBlockListDraftsParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessBlockListDraftsBatchCall(id, params, options...))
}

func DecodeCreateBusinessBlockListDraftsBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessBlockListDrafts(ctx context.Context, client *core.Client, id string, params CreateBusinessBlockListDraftsParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessBlockListDraftsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessBmReviewRequestsBatchCall(id string, params CreateBusinessBmReviewRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "bm_review_requests"), params.ToParams(), options...)
}

func NewCreateBusinessBmReviewRequestsBatchRequest(id string, params CreateBusinessBmReviewRequestsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateBusinessBmReviewRequestsBatchCall(id, params, options...))
}

func DecodeCreateBusinessBmReviewRequestsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateBusinessBmReviewRequests(ctx context.Context, client *core.Client, id string, params CreateBusinessBmReviewRequestsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateBusinessBmReviewRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessBusinessAssetGroupsBatchCall(id string, params GetBusinessBusinessAssetGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "business_asset_groups"), params.ToParams(), options...)
}

func NewGetBusinessBusinessAssetGroupsBatchRequest(id string, params GetBusinessBusinessAssetGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetGroup]](GetBusinessBusinessAssetGroupsBatchCall(id, params, options...))
}

func DecodeGetBusinessBusinessAssetGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAssetGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	call := GetBusinessBusinessAssetGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessBusinessInvoicesBatchCall(id string, params GetBusinessBusinessInvoicesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "business_invoices"), params.ToParams(), options...)
}

func NewGetBusinessBusinessInvoicesBatchRequest(id string, params GetBusinessBusinessInvoicesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OmegaCustomerTrx]] {
	return core.NewBatchRequest[core.Cursor[objects.OmegaCustomerTrx]](GetBusinessBusinessInvoicesBatchCall(id, params, options...))
}

func DecodeGetBusinessBusinessInvoicesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OmegaCustomerTrx], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OmegaCustomerTrx]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessBusinessInvoices(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessInvoicesParams) (*core.Cursor[objects.OmegaCustomerTrx], error) {
	var out core.Cursor[objects.OmegaCustomerTrx]
	call := GetBusinessBusinessInvoicesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessBusinessUsersBatchCall(id string, params GetBusinessBusinessUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "business_users"), params.ToParams(), options...)
}

func NewGetBusinessBusinessUsersBatchRequest(id string, params GetBusinessBusinessUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessUser]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessUser]](GetBusinessBusinessUsersBatchCall(id, params, options...))
}

func DecodeGetBusinessBusinessUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessBusinessUsers(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessUsersParams) (*core.Cursor[objects.BusinessUser], error) {
	var out core.Cursor[objects.BusinessUser]
	call := GetBusinessBusinessUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessBusinessUsersBatchCall(id string, params CreateBusinessBusinessUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "business_users"), params.ToParams(), options...)
}

func NewCreateBusinessBusinessUsersBatchRequest(id string, params CreateBusinessBusinessUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessUser] {
	return core.NewBatchRequest[objects.BusinessUser](CreateBusinessBusinessUsersBatchCall(id, params, options...))
}

func DecodeCreateBusinessBusinessUsersBatchResponse(response *core.BatchResponse) (*objects.BusinessUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessBusinessUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessBusinessUsersParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	call := CreateBusinessBusinessUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessBusinessprojectsBatchCall(id string, params GetBusinessBusinessprojectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "businessprojects"), params.ToParams(), options...)
}

func NewGetBusinessBusinessprojectsBatchRequest(id string, params GetBusinessBusinessprojectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessProject]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessProject]](GetBusinessBusinessprojectsBatchCall(id, params, options...))
}

func DecodeGetBusinessBusinessprojectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessProject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessProject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessBusinessprojects(ctx context.Context, client *core.Client, id string, params GetBusinessBusinessprojectsParams) (*core.Cursor[objects.BusinessProject], error) {
	var out core.Cursor[objects.BusinessProject]
	call := GetBusinessBusinessprojectsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessClaimCustomConversionsBatchCall(id string, params CreateBusinessClaimCustomConversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "claim_custom_conversions"), params.ToParams(), options...)
}

func NewCreateBusinessClaimCustomConversionsBatchRequest(id string, params CreateBusinessClaimCustomConversionsParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomConversion] {
	return core.NewBatchRequest[objects.CustomConversion](CreateBusinessClaimCustomConversionsBatchCall(id, params, options...))
}

func DecodeCreateBusinessClaimCustomConversionsBatchResponse(response *core.BatchResponse) (*objects.CustomConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessClaimCustomConversions(ctx context.Context, client *core.Client, id string, params CreateBusinessClaimCustomConversionsParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	call := CreateBusinessClaimCustomConversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientAdAccountsBatchCall(id string, params GetBusinessClientAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_ad_accounts"), params.ToParams(), options...)
}

func NewGetBusinessClientAdAccountsBatchRequest(id string, params GetBusinessClientAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetBusinessClientAdAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetBusinessClientAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessClientAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetBusinessClientAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientAppsBatchCall(id string, params GetBusinessClientAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_apps"), params.ToParams(), options...)
}

func NewGetBusinessClientAppsBatchRequest(id string, params GetBusinessClientAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetBusinessClientAppsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientApps(ctx context.Context, client *core.Client, id string, params GetBusinessClientAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetBusinessClientAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessClientAppsBatchCall(id string, params CreateBusinessClientAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "client_apps"), params.ToParams(), options...)
}

func NewCreateBusinessClientAppsBatchRequest(id string, params CreateBusinessClientAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessClientAppsBatchCall(id, params, options...))
}

func DecodeCreateBusinessClientAppsBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessClientApps(ctx context.Context, client *core.Client, id string, params CreateBusinessClientAppsParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessClientAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientInstagramAssetsBatchCall(id string, params GetBusinessClientInstagramAssetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_instagram_assets"), params.ToParams(), options...)
}

func NewGetBusinessClientInstagramAssetsBatchRequest(id string, params GetBusinessClientInstagramAssetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InstagramBusinessAsset]] {
	return core.NewBatchRequest[core.Cursor[objects.InstagramBusinessAsset]](GetBusinessClientInstagramAssetsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientInstagramAssetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InstagramBusinessAsset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientInstagramAssets(ctx context.Context, client *core.Client, id string, params GetBusinessClientInstagramAssetsParams) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	var out core.Cursor[objects.InstagramBusinessAsset]
	call := GetBusinessClientInstagramAssetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientOffsiteSignalContainerBusinessObjectsBatchCall(id string, params GetBusinessClientOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_offsite_signal_container_business_objects"), params.ToParams(), options...)
}

func NewGetBusinessClientOffsiteSignalContainerBusinessObjectsBatchRequest(id string, params GetBusinessClientOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]] {
	return core.NewBatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]](GetBusinessClientOffsiteSignalContainerBusinessObjectsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientOffsiteSignalContainerBusinessObjectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessClientOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	call := GetBusinessClientOffsiteSignalContainerBusinessObjectsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientPagesBatchCall(id string, params GetBusinessClientPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_pages"), params.ToParams(), options...)
}

func NewGetBusinessClientPagesBatchRequest(id string, params GetBusinessClientPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetBusinessClientPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessClientPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientPages(ctx context.Context, client *core.Client, id string, params GetBusinessClientPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetBusinessClientPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessClientPagesBatchCall(id string, params CreateBusinessClientPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "client_pages"), params.ToParams(), options...)
}

func NewCreateBusinessClientPagesBatchRequest(id string, params CreateBusinessClientPagesParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessClientPagesBatchCall(id, params, options...))
}

func DecodeCreateBusinessClientPagesBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessClientPages(ctx context.Context, client *core.Client, id string, params CreateBusinessClientPagesParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessClientPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientPixelsBatchCall(id string, params GetBusinessClientPixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_pixels"), params.ToParams(), options...)
}

func NewGetBusinessClientPixelsBatchRequest(id string, params GetBusinessClientPixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetBusinessClientPixelsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientPixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientPixels(ctx context.Context, client *core.Client, id string, params GetBusinessClientPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetBusinessClientPixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientProductCatalogsBatchCall(id string, params GetBusinessClientProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_product_catalogs"), params.ToParams(), options...)
}

func NewGetBusinessClientProductCatalogsBatchRequest(id string, params GetBusinessClientProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetBusinessClientProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalog]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessClientProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetBusinessClientProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientWhatsappBusinessAccountsBatchCall(id string, params GetBusinessClientWhatsappBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "client_whatsapp_business_accounts"), params.ToParams(), options...)
}

func NewGetBusinessClientWhatsappBusinessAccountsBatchRequest(id string, params GetBusinessClientWhatsappBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]](GetBusinessClientWhatsappBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientWhatsappBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessClientWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessClientWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	call := GetBusinessClientWhatsappBusinessAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessClientsBatchCall(id string, params DeleteBusinessClientsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "clients"), params.ToParams(), options...)
}

func NewDeleteBusinessClientsBatchRequest(id string, params DeleteBusinessClientsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessClientsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessClientsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessClients(ctx context.Context, client *core.Client, id string, params DeleteBusinessClientsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessClientsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessClientsBatchCall(id string, params GetBusinessClientsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "clients"), params.ToParams(), options...)
}

func NewGetBusinessClientsBatchRequest(id string, params GetBusinessClientsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetBusinessClientsBatchCall(id, params, options...))
}

func DecodeGetBusinessClientsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetBusinessClients(ctx context.Context, client *core.Client, id string, params GetBusinessClientsParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	call := GetBusinessClientsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCollaborativeAdsCollaborationRequestsBatchCall(id string, params GetBusinessCollaborativeAdsCollaborationRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaborative_ads_collaboration_requests"), params.ToParams(), options...)
}

func NewGetBusinessCollaborativeAdsCollaborationRequestsBatchRequest(id string, params GetBusinessCollaborativeAdsCollaborationRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CPASCollaborationRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.CPASCollaborationRequest]](GetBusinessCollaborativeAdsCollaborationRequestsBatchCall(id, params, options...))
}

func DecodeGetBusinessCollaborativeAdsCollaborationRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CPASCollaborationRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CPASCollaborationRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCollaborativeAdsCollaborationRequests(ctx context.Context, client *core.Client, id string, params GetBusinessCollaborativeAdsCollaborationRequestsParams) (*core.Cursor[objects.CPASCollaborationRequest], error) {
	var out core.Cursor[objects.CPASCollaborationRequest]
	call := GetBusinessCollaborativeAdsCollaborationRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCollaborativeAdsSuggestedPartnersBatchCall(id string, params GetBusinessCollaborativeAdsSuggestedPartnersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaborative_ads_suggested_partners"), params.ToParams(), options...)
}

func NewGetBusinessCollaborativeAdsSuggestedPartnersBatchRequest(id string, params GetBusinessCollaborativeAdsSuggestedPartnersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CPASAdvertiserPartnershipRecommendation]] {
	return core.NewBatchRequest[core.Cursor[objects.CPASAdvertiserPartnershipRecommendation]](GetBusinessCollaborativeAdsSuggestedPartnersBatchCall(id, params, options...))
}

func DecodeGetBusinessCollaborativeAdsSuggestedPartnersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CPASAdvertiserPartnershipRecommendation], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CPASAdvertiserPartnershipRecommendation]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCollaborativeAdsSuggestedPartners(ctx context.Context, client *core.Client, id string, params GetBusinessCollaborativeAdsSuggestedPartnersParams) (*core.Cursor[objects.CPASAdvertiserPartnershipRecommendation], error) {
	var out core.Cursor[objects.CPASAdvertiserPartnershipRecommendation]
	call := GetBusinessCollaborativeAdsSuggestedPartnersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCommerceMerchantSettingsBatchCall(id string, params GetBusinessCommerceMerchantSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_merchant_settings"), params.ToParams(), options...)
}

func NewGetBusinessCommerceMerchantSettingsBatchRequest(id string, params GetBusinessCommerceMerchantSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceMerchantSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceMerchantSettings]](GetBusinessCommerceMerchantSettingsBatchCall(id, params, options...))
}

func DecodeGetBusinessCommerceMerchantSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceMerchantSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceMerchantSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params GetBusinessCommerceMerchantSettingsParams) (*core.Cursor[objects.CommerceMerchantSettings], error) {
	var out core.Cursor[objects.CommerceMerchantSettings]
	call := GetBusinessCommerceMerchantSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCpasBusinessSetupConfigBatchCall(id string, params GetBusinessCpasBusinessSetupConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "cpas_business_setup_config"), params.ToParams(), options...)
}

func NewGetBusinessCpasBusinessSetupConfigBatchRequest(id string, params GetBusinessCpasBusinessSetupConfigParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CPASBusinessSetupConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.CPASBusinessSetupConfig]](GetBusinessCpasBusinessSetupConfigBatchCall(id, params, options...))
}

func DecodeGetBusinessCpasBusinessSetupConfigBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CPASBusinessSetupConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CPASBusinessSetupConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCpasBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params GetBusinessCpasBusinessSetupConfigParams) (*core.Cursor[objects.CPASBusinessSetupConfig], error) {
	var out core.Cursor[objects.CPASBusinessSetupConfig]
	call := GetBusinessCpasBusinessSetupConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessCpasBusinessSetupConfigBatchCall(id string, params CreateBusinessCpasBusinessSetupConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "cpas_business_setup_config"), params.ToParams(), options...)
}

func NewCreateBusinessCpasBusinessSetupConfigBatchRequest(id string, params CreateBusinessCpasBusinessSetupConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASBusinessSetupConfig] {
	return core.NewBatchRequest[objects.CPASBusinessSetupConfig](CreateBusinessCpasBusinessSetupConfigBatchCall(id, params, options...))
}

func DecodeCreateBusinessCpasBusinessSetupConfigBatchResponse(response *core.BatchResponse) (*objects.CPASBusinessSetupConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASBusinessSetupConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessCpasBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params CreateBusinessCpasBusinessSetupConfigParams) (*objects.CPASBusinessSetupConfig, error) {
	var out objects.CPASBusinessSetupConfig
	call := CreateBusinessCpasBusinessSetupConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCpasMerchantConfigBatchCall(id string, params GetBusinessCpasMerchantConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "cpas_merchant_config"), params.ToParams(), options...)
}

func NewGetBusinessCpasMerchantConfigBatchRequest(id string, params GetBusinessCpasMerchantConfigParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CPASMerchantConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.CPASMerchantConfig]](GetBusinessCpasMerchantConfigBatchCall(id, params, options...))
}

func DecodeGetBusinessCpasMerchantConfigBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CPASMerchantConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CPASMerchantConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCpasMerchantConfig(ctx context.Context, client *core.Client, id string, params GetBusinessCpasMerchantConfigParams) (*core.Cursor[objects.CPASMerchantConfig], error) {
	var out core.Cursor[objects.CPASMerchantConfig]
	call := GetBusinessCpasMerchantConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessCreativeFoldersBatchCall(id string, params CreateBusinessCreativeFoldersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "creative_folders"), params.ToParams(), options...)
}

func NewCreateBusinessCreativeFoldersBatchRequest(id string, params CreateBusinessCreativeFoldersParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessCreativeFolder] {
	return core.NewBatchRequest[objects.BusinessCreativeFolder](CreateBusinessCreativeFoldersBatchCall(id, params, options...))
}

func DecodeCreateBusinessCreativeFoldersBatchResponse(response *core.BatchResponse) (*objects.BusinessCreativeFolder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessCreativeFolder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessCreativeFolders(ctx context.Context, client *core.Client, id string, params CreateBusinessCreativeFoldersParams) (*objects.BusinessCreativeFolder, error) {
	var out objects.BusinessCreativeFolder
	call := CreateBusinessCreativeFoldersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessCreditcardsBatchCall(id string, params GetBusinessCreditcardsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "creditcards"), params.ToParams(), options...)
}

func NewGetBusinessCreditcardsBatchRequest(id string, params GetBusinessCreditcardsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CreditCard]] {
	return core.NewBatchRequest[core.Cursor[objects.CreditCard]](GetBusinessCreditcardsBatchCall(id, params, options...))
}

func DecodeGetBusinessCreditcardsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CreditCard], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CreditCard]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessCreditcards(ctx context.Context, client *core.Client, id string, params GetBusinessCreditcardsParams) (*core.Cursor[objects.CreditCard], error) {
	var out core.Cursor[objects.CreditCard]
	call := GetBusinessCreditcardsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessCustomconversionsBatchCall(id string, params CreateBusinessCustomconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "customconversions"), params.ToParams(), options...)
}

func NewCreateBusinessCustomconversionsBatchRequest(id string, params CreateBusinessCustomconversionsParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomConversion] {
	return core.NewBatchRequest[objects.CustomConversion](CreateBusinessCustomconversionsBatchCall(id, params, options...))
}

func DecodeCreateBusinessCustomconversionsBatchResponse(response *core.BatchResponse) (*objects.CustomConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessCustomconversions(ctx context.Context, client *core.Client, id string, params CreateBusinessCustomconversionsParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	call := CreateBusinessCustomconversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessEventSourceGroupsBatchCall(id string, params GetBusinessEventSourceGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "event_source_groups"), params.ToParams(), options...)
}

func NewGetBusinessEventSourceGroupsBatchRequest(id string, params GetBusinessEventSourceGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.EventSourceGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.EventSourceGroup]](GetBusinessEventSourceGroupsBatchCall(id, params, options...))
}

func DecodeGetBusinessEventSourceGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.EventSourceGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.EventSourceGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessEventSourceGroups(ctx context.Context, client *core.Client, id string, params GetBusinessEventSourceGroupsParams) (*core.Cursor[objects.EventSourceGroup], error) {
	var out core.Cursor[objects.EventSourceGroup]
	call := GetBusinessEventSourceGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessEventSourceGroupsBatchCall(id string, params CreateBusinessEventSourceGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "event_source_groups"), params.ToParams(), options...)
}

func NewCreateBusinessEventSourceGroupsBatchRequest(id string, params CreateBusinessEventSourceGroupsParams, options ...core.BatchOption) *core.BatchRequest[objects.EventSourceGroup] {
	return core.NewBatchRequest[objects.EventSourceGroup](CreateBusinessEventSourceGroupsBatchCall(id, params, options...))
}

func DecodeCreateBusinessEventSourceGroupsBatchResponse(response *core.BatchResponse) (*objects.EventSourceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventSourceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessEventSourceGroups(ctx context.Context, client *core.Client, id string, params CreateBusinessEventSourceGroupsParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	call := CreateBusinessEventSourceGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessExtendedcreditapplicationsBatchCall(id string, params GetBusinessExtendedcreditapplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "extendedcreditapplications"), params.ToParams(), options...)
}

func NewGetBusinessExtendedcreditapplicationsBatchRequest(id string, params GetBusinessExtendedcreditapplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ExtendedCreditApplication]] {
	return core.NewBatchRequest[core.Cursor[objects.ExtendedCreditApplication]](GetBusinessExtendedcreditapplicationsBatchCall(id, params, options...))
}

func DecodeGetBusinessExtendedcreditapplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ExtendedCreditApplication], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ExtendedCreditApplication]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessExtendedcreditapplications(ctx context.Context, client *core.Client, id string, params GetBusinessExtendedcreditapplicationsParams) (*core.Cursor[objects.ExtendedCreditApplication], error) {
	var out core.Cursor[objects.ExtendedCreditApplication]
	call := GetBusinessExtendedcreditapplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessExtendedcreditsBatchCall(id string, params GetBusinessExtendedcreditsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "extendedcredits"), params.ToParams(), options...)
}

func NewGetBusinessExtendedcreditsBatchRequest(id string, params GetBusinessExtendedcreditsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ExtendedCredit]] {
	return core.NewBatchRequest[core.Cursor[objects.ExtendedCredit]](GetBusinessExtendedcreditsBatchCall(id, params, options...))
}

func DecodeGetBusinessExtendedcreditsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ExtendedCredit], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ExtendedCredit]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessExtendedcredits(ctx context.Context, client *core.Client, id string, params GetBusinessExtendedcreditsParams) (*core.Cursor[objects.ExtendedCredit], error) {
	var out core.Cursor[objects.ExtendedCredit]
	call := GetBusinessExtendedcreditsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessImagesBatchCall(id string, params CreateBusinessImagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "images"), params.ToParams(), options...)
}

func NewCreateBusinessImagesBatchRequest(id string, params CreateBusinessImagesParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessImage] {
	return core.NewBatchRequest[objects.BusinessImage](CreateBusinessImagesBatchCall(id, params, options...))
}

func DecodeCreateBusinessImagesBatchResponse(response *core.BatchResponse) (*objects.BusinessImage, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessImage
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessImages(ctx context.Context, client *core.Client, id string, params CreateBusinessImagesParams) (*objects.BusinessImage, error) {
	var out objects.BusinessImage
	call := CreateBusinessImagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessInitiatedAudienceSharingRequestsBatchCall(id string, params GetBusinessInitiatedAudienceSharingRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "initiated_audience_sharing_requests"), params.ToParams(), options...)
}

func NewGetBusinessInitiatedAudienceSharingRequestsBatchRequest(id string, params GetBusinessInitiatedAudienceSharingRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetSharingAgreement]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetSharingAgreement]](GetBusinessInitiatedAudienceSharingRequestsBatchCall(id, params, options...))
}

func DecodeGetBusinessInitiatedAudienceSharingRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessInitiatedAudienceSharingRequests(ctx context.Context, client *core.Client, id string, params GetBusinessInitiatedAudienceSharingRequestsParams) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	call := GetBusinessInitiatedAudienceSharingRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessInstagramAccountsBatchCall(id string, params DeleteBusinessInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "instagram_accounts"), params.ToParams(), options...)
}

func NewDeleteBusinessInstagramAccountsBatchRequest(id string, params DeleteBusinessInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessInstagramAccountsBatchCall(id, params, options...))
}

func DecodeDeleteBusinessInstagramAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessInstagramAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessInstagramAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessInstagramAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessInstagramAccountsBatchCall(id string, params GetBusinessInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_accounts"), params.ToParams(), options...)
}

func NewGetBusinessInstagramAccountsBatchRequest(id string, params GetBusinessInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetBusinessInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetBusinessInstagramAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessInstagramBusinessAccountsBatchCall(id string, params GetBusinessInstagramBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_business_accounts"), params.ToParams(), options...)
}

func NewGetBusinessInstagramBusinessAccountsBatchRequest(id string, params GetBusinessInstagramBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetBusinessInstagramBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessInstagramBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessInstagramBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessInstagramBusinessAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetBusinessInstagramBusinessAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessManagedBusinessesBatchCall(id string, params DeleteBusinessManagedBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "managed_businesses"), params.ToParams(), options...)
}

func NewDeleteBusinessManagedBusinessesBatchRequest(id string, params DeleteBusinessManagedBusinessesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessManagedBusinessesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessManagedBusinessesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessManagedBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessManagedBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessManagedBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessManagedBusinessesBatchCall(id string, params CreateBusinessManagedBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "managed_businesses"), params.ToParams(), options...)
}

func NewCreateBusinessManagedBusinessesBatchRequest(id string, params CreateBusinessManagedBusinessesParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessManagedBusinessesBatchCall(id, params, options...))
}

func DecodeCreateBusinessManagedBusinessesBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessManagedBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedBusinessesParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessManagedBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessManagedPartnerAdsFundingSourceDetailsBatchCall(id string, params GetBusinessManagedPartnerAdsFundingSourceDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "managed_partner_ads_funding_source_details"), params.ToParams(), options...)
}

func NewGetBusinessManagedPartnerAdsFundingSourceDetailsBatchRequest(id string, params GetBusinessManagedPartnerAdsFundingSourceDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.FundingSourceDetailsCoupon]] {
	return core.NewBatchRequest[core.Cursor[objects.FundingSourceDetailsCoupon]](GetBusinessManagedPartnerAdsFundingSourceDetailsBatchCall(id, params, options...))
}

func DecodeGetBusinessManagedPartnerAdsFundingSourceDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.FundingSourceDetailsCoupon], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.FundingSourceDetailsCoupon]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessManagedPartnerAdsFundingSourceDetails(ctx context.Context, client *core.Client, id string, params GetBusinessManagedPartnerAdsFundingSourceDetailsParams) (*core.Cursor[objects.FundingSourceDetailsCoupon], error) {
	var out core.Cursor[objects.FundingSourceDetailsCoupon]
	call := GetBusinessManagedPartnerAdsFundingSourceDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessManagedPartnerBusinessSetupBatchCall(id string, params CreateBusinessManagedPartnerBusinessSetupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "managed_partner_business_setup"), params.ToParams(), options...)
}

func NewCreateBusinessManagedPartnerBusinessSetupBatchRequest(id string, params CreateBusinessManagedPartnerBusinessSetupParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessManagedPartnerBusinessSetupBatchCall(id, params, options...))
}

func DecodeCreateBusinessManagedPartnerBusinessSetupBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessManagedPartnerBusinessSetup(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedPartnerBusinessSetupParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessManagedPartnerBusinessSetupBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessManagedPartnerBusinessesBatchCall(id string, params DeleteBusinessManagedPartnerBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "managed_partner_businesses"), params.ToParams(), options...)
}

func NewDeleteBusinessManagedPartnerBusinessesBatchRequest(id string, params DeleteBusinessManagedPartnerBusinessesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessManagedPartnerBusinessesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessManagedPartnerBusinessesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessManagedPartnerBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessManagedPartnerBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessManagedPartnerBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessManagedPartnerBusinessesBatchCall(id string, params CreateBusinessManagedPartnerBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "managed_partner_businesses"), params.ToParams(), options...)
}

func NewCreateBusinessManagedPartnerBusinessesBatchRequest(id string, params CreateBusinessManagedPartnerBusinessesParams, options ...core.BatchOption) *core.BatchRequest[objects.ManagedPartnerBusiness] {
	return core.NewBatchRequest[objects.ManagedPartnerBusiness](CreateBusinessManagedPartnerBusinessesBatchCall(id, params, options...))
}

func DecodeCreateBusinessManagedPartnerBusinessesBatchResponse(response *core.BatchResponse) (*objects.ManagedPartnerBusiness, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ManagedPartnerBusiness
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessManagedPartnerBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessManagedPartnerBusinessesParams) (*objects.ManagedPartnerBusiness, error) {
	var out objects.ManagedPartnerBusiness
	call := CreateBusinessManagedPartnerBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOnboardPartnersToMmLiteBatchCall(id string, params CreateBusinessOnboardPartnersToMmLiteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "onboard_partners_to_mm_lite"), params.ToParams(), options...)
}

func NewCreateBusinessOnboardPartnersToMmLiteBatchRequest(id string, params CreateBusinessOnboardPartnersToMmLiteParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateBusinessOnboardPartnersToMmLiteBatchCall(id, params, options...))
}

func DecodeCreateBusinessOnboardPartnersToMmLiteBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateBusinessOnboardPartnersToMmLite(ctx context.Context, client *core.Client, id string, params CreateBusinessOnboardPartnersToMmLiteParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateBusinessOnboardPartnersToMmLiteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOpenbridgeConfigurationsBatchCall(id string, params GetBusinessOpenbridgeConfigurationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), options...)
}

func NewGetBusinessOpenbridgeConfigurationsBatchRequest(id string, params GetBusinessOpenbridgeConfigurationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OpenBridgeConfiguration]] {
	return core.NewBatchRequest[core.Cursor[objects.OpenBridgeConfiguration]](GetBusinessOpenbridgeConfigurationsBatchCall(id, params, options...))
}

func DecodeGetBusinessOpenbridgeConfigurationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OpenBridgeConfiguration]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params GetBusinessOpenbridgeConfigurationsParams) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	var out core.Cursor[objects.OpenBridgeConfiguration]
	call := GetBusinessOpenbridgeConfigurationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOpenbridgeConfigurationsBatchCall(id string, params CreateBusinessOpenbridgeConfigurationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), options...)
}

func NewCreateBusinessOpenbridgeConfigurationsBatchRequest(id string, params CreateBusinessOpenbridgeConfigurationsParams, options ...core.BatchOption) *core.BatchRequest[objects.OpenBridgeConfiguration] {
	return core.NewBatchRequest[objects.OpenBridgeConfiguration](CreateBusinessOpenbridgeConfigurationsBatchCall(id, params, options...))
}

func DecodeCreateBusinessOpenbridgeConfigurationsBatchResponse(response *core.BatchResponse) (*objects.OpenBridgeConfiguration, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OpenBridgeConfiguration
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params CreateBusinessOpenbridgeConfigurationsParams) (*objects.OpenBridgeConfiguration, error) {
	var out objects.OpenBridgeConfiguration
	call := CreateBusinessOpenbridgeConfigurationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedAdAccountsBatchCall(id string, params GetBusinessOwnedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_ad_accounts"), params.ToParams(), options...)
}

func NewGetBusinessOwnedAdAccountsBatchRequest(id string, params GetBusinessOwnedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetBusinessOwnedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetBusinessOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetBusinessOwnedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOwnedAdAccountsBatchCall(id string, params CreateBusinessOwnedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owned_ad_accounts"), params.ToParams(), options...)
}

func NewCreateBusinessOwnedAdAccountsBatchRequest(id string, params CreateBusinessOwnedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessOwnedAdAccountsBatchCall(id, params, options...))
}

func DecodeCreateBusinessOwnedAdAccountsBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedAdAccountsParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessOwnedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedAppsBatchCall(id string, params GetBusinessOwnedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_apps"), params.ToParams(), options...)
}

func NewGetBusinessOwnedAppsBatchRequest(id string, params GetBusinessOwnedAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetBusinessOwnedAppsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedApps(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetBusinessOwnedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOwnedAppsBatchCall(id string, params CreateBusinessOwnedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owned_apps"), params.ToParams(), options...)
}

func NewCreateBusinessOwnedAppsBatchRequest(id string, params CreateBusinessOwnedAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessOwnedAppsBatchCall(id, params, options...))
}

func DecodeCreateBusinessOwnedAppsBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOwnedApps(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedAppsParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessOwnedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessOwnedBusinessesBatchCall(id string, params DeleteBusinessOwnedBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "owned_businesses"), params.ToParams(), options...)
}

func NewDeleteBusinessOwnedBusinessesBatchRequest(id string, params DeleteBusinessOwnedBusinessesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessOwnedBusinessesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessOwnedBusinessesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params DeleteBusinessOwnedBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessOwnedBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedBusinessesBatchCall(id string, params GetBusinessOwnedBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_businesses"), params.ToParams(), options...)
}

func NewGetBusinessOwnedBusinessesBatchRequest(id string, params GetBusinessOwnedBusinessesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetBusinessOwnedBusinessesBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedBusinessesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedBusinessesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	call := GetBusinessOwnedBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOwnedBusinessesBatchCall(id string, params CreateBusinessOwnedBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owned_businesses"), params.ToParams(), options...)
}

func NewCreateBusinessOwnedBusinessesBatchRequest(id string, params CreateBusinessOwnedBusinessesParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessOwnedBusinessesBatchCall(id, params, options...))
}

func DecodeCreateBusinessOwnedBusinessesBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOwnedBusinesses(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedBusinessesParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessOwnedBusinessesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedInstagramAccountsBatchCall(id string, params GetBusinessOwnedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_instagram_accounts"), params.ToParams(), options...)
}

func NewGetBusinessOwnedInstagramAccountsBatchRequest(id string, params GetBusinessOwnedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetBusinessOwnedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetBusinessOwnedInstagramAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedInstagramAssetsBatchCall(id string, params GetBusinessOwnedInstagramAssetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_instagram_assets"), params.ToParams(), options...)
}

func NewGetBusinessOwnedInstagramAssetsBatchRequest(id string, params GetBusinessOwnedInstagramAssetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InstagramBusinessAsset]] {
	return core.NewBatchRequest[core.Cursor[objects.InstagramBusinessAsset]](GetBusinessOwnedInstagramAssetsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedInstagramAssetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InstagramBusinessAsset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedInstagramAssets(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedInstagramAssetsParams) (*core.Cursor[objects.InstagramBusinessAsset], error) {
	var out core.Cursor[objects.InstagramBusinessAsset]
	call := GetBusinessOwnedInstagramAssetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedOffsiteSignalContainerBusinessObjectsBatchCall(id string, params GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_offsite_signal_container_business_objects"), params.ToParams(), options...)
}

func NewGetBusinessOwnedOffsiteSignalContainerBusinessObjectsBatchRequest(id string, params GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]] {
	return core.NewBatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]](GetBusinessOwnedOffsiteSignalContainerBusinessObjectsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedOffsiteSignalContainerBusinessObjectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	call := GetBusinessOwnedOffsiteSignalContainerBusinessObjectsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedPagesBatchCall(id string, params GetBusinessOwnedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_pages"), params.ToParams(), options...)
}

func NewGetBusinessOwnedPagesBatchRequest(id string, params GetBusinessOwnedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetBusinessOwnedPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedPages(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetBusinessOwnedPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOwnedPagesBatchCall(id string, params CreateBusinessOwnedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owned_pages"), params.ToParams(), options...)
}

func NewCreateBusinessOwnedPagesBatchRequest(id string, params CreateBusinessOwnedPagesParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessOwnedPagesBatchCall(id, params, options...))
}

func DecodeCreateBusinessOwnedPagesBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOwnedPages(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedPagesParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessOwnedPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedPixelsBatchCall(id string, params GetBusinessOwnedPixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_pixels"), params.ToParams(), options...)
}

func NewGetBusinessOwnedPixelsBatchRequest(id string, params GetBusinessOwnedPixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetBusinessOwnedPixelsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedPixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedPixels(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetBusinessOwnedPixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedProductCatalogsBatchCall(id string, params GetBusinessOwnedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_product_catalogs"), params.ToParams(), options...)
}

func NewGetBusinessOwnedProductCatalogsBatchRequest(id string, params GetBusinessOwnedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetBusinessOwnedProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalog]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetBusinessOwnedProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessOwnedProductCatalogsBatchCall(id string, params CreateBusinessOwnedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "owned_product_catalogs"), params.ToParams(), options...)
}

func NewCreateBusinessOwnedProductCatalogsBatchRequest(id string, params CreateBusinessOwnedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductCatalog] {
	return core.NewBatchRequest[objects.ProductCatalog](CreateBusinessOwnedProductCatalogsBatchCall(id, params, options...))
}

func DecodeCreateBusinessOwnedProductCatalogsBatchResponse(response *core.BatchResponse) (*objects.ProductCatalog, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductCatalog
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessOwnedProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateBusinessOwnedProductCatalogsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	call := CreateBusinessOwnedProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessOwnedWhatsappBusinessAccountsBatchCall(id string, params GetBusinessOwnedWhatsappBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "owned_whatsapp_business_accounts"), params.ToParams(), options...)
}

func NewGetBusinessOwnedWhatsappBusinessAccountsBatchRequest(id string, params GetBusinessOwnedWhatsappBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]](GetBusinessOwnedWhatsappBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedWhatsappBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	call := GetBusinessOwnedWhatsappBusinessAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessPagesBatchCall(id string, params DeleteBusinessPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "pages"), params.ToParams(), options...)
}

func NewDeleteBusinessPagesBatchRequest(id string, params DeleteBusinessPagesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessPagesBatchCall(id, params, options...))
}

func DecodeDeleteBusinessPagesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessPages(ctx context.Context, client *core.Client, id string, params DeleteBusinessPagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPartnerAccountLinkingBatchCall(id string, params GetBusinessPartnerAccountLinkingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "partner_account_linking"), params.ToParams(), options...)
}

func NewGetBusinessPartnerAccountLinkingBatchRequest(id string, params GetBusinessPartnerAccountLinkingParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PartnerAccountLinking]] {
	return core.NewBatchRequest[core.Cursor[objects.PartnerAccountLinking]](GetBusinessPartnerAccountLinkingBatchCall(id, params, options...))
}

func DecodeGetBusinessPartnerAccountLinkingBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PartnerAccountLinking], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PartnerAccountLinking]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPartnerAccountLinking(ctx context.Context, client *core.Client, id string, params GetBusinessPartnerAccountLinkingParams) (*core.Cursor[objects.PartnerAccountLinking], error) {
	var out core.Cursor[objects.PartnerAccountLinking]
	call := GetBusinessPartnerAccountLinkingBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessPartnerPremiumOptionsBatchCall(id string, params CreateBusinessPartnerPremiumOptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "partner_premium_options"), params.ToParams(), options...)
}

func NewCreateBusinessPartnerPremiumOptionsBatchRequest(id string, params CreateBusinessPartnerPremiumOptionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateBusinessPartnerPremiumOptionsBatchCall(id, params, options...))
}

func DecodeCreateBusinessPartnerPremiumOptionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateBusinessPartnerPremiumOptions(ctx context.Context, client *core.Client, id string, params CreateBusinessPartnerPremiumOptionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateBusinessPartnerPremiumOptionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPassbackAttributionMetadataConfigsBatchCall(id string, params GetBusinessPassbackAttributionMetadataConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "passback_attribution_metadata_configs"), params.ToParams(), options...)
}

func NewGetBusinessPassbackAttributionMetadataConfigsBatchRequest(id string, params GetBusinessPassbackAttributionMetadataConfigsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.SignalsAttributionMetadataConfig]] {
	return core.NewBatchRequest[core.Cursor[objects.SignalsAttributionMetadataConfig]](GetBusinessPassbackAttributionMetadataConfigsBatchCall(id, params, options...))
}

func DecodeGetBusinessPassbackAttributionMetadataConfigsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.SignalsAttributionMetadataConfig], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.SignalsAttributionMetadataConfig]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPassbackAttributionMetadataConfigs(ctx context.Context, client *core.Client, id string, params GetBusinessPassbackAttributionMetadataConfigsParams) (*core.Cursor[objects.SignalsAttributionMetadataConfig], error) {
	var out core.Cursor[objects.SignalsAttributionMetadataConfig]
	call := GetBusinessPassbackAttributionMetadataConfigsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingClientAdAccountsBatchCall(id string, params GetBusinessPendingClientAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_client_ad_accounts"), params.ToParams(), options...)
}

func NewGetBusinessPendingClientAdAccountsBatchRequest(id string, params GetBusinessPendingClientAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAdAccountRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAdAccountRequest]](GetBusinessPendingClientAdAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingClientAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAdAccountRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingClientAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientAdAccountsParams) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	var out core.Cursor[objects.BusinessAdAccountRequest]
	call := GetBusinessPendingClientAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingClientAppsBatchCall(id string, params GetBusinessPendingClientAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_client_apps"), params.ToParams(), options...)
}

func NewGetBusinessPendingClientAppsBatchRequest(id string, params GetBusinessPendingClientAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessApplicationRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessApplicationRequest]](GetBusinessPendingClientAppsBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingClientAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessApplicationRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessApplicationRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingClientApps(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientAppsParams) (*core.Cursor[objects.BusinessApplicationRequest], error) {
	var out core.Cursor[objects.BusinessApplicationRequest]
	call := GetBusinessPendingClientAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingClientPagesBatchCall(id string, params GetBusinessPendingClientPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_client_pages"), params.ToParams(), options...)
}

func NewGetBusinessPendingClientPagesBatchRequest(id string, params GetBusinessPendingClientPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessPageRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessPageRequest]](GetBusinessPendingClientPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingClientPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessPageRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessPageRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingClientPages(ctx context.Context, client *core.Client, id string, params GetBusinessPendingClientPagesParams) (*core.Cursor[objects.BusinessPageRequest], error) {
	var out core.Cursor[objects.BusinessPageRequest]
	call := GetBusinessPendingClientPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingOwnedAdAccountsBatchCall(id string, params GetBusinessPendingOwnedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_owned_ad_accounts"), params.ToParams(), options...)
}

func NewGetBusinessPendingOwnedAdAccountsBatchRequest(id string, params GetBusinessPendingOwnedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAdAccountRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAdAccountRequest]](GetBusinessPendingOwnedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingOwnedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAdAccountRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingOwnedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessPendingOwnedAdAccountsParams) (*core.Cursor[objects.BusinessAdAccountRequest], error) {
	var out core.Cursor[objects.BusinessAdAccountRequest]
	call := GetBusinessPendingOwnedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingOwnedPagesBatchCall(id string, params GetBusinessPendingOwnedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_owned_pages"), params.ToParams(), options...)
}

func NewGetBusinessPendingOwnedPagesBatchRequest(id string, params GetBusinessPendingOwnedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessPageRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessPageRequest]](GetBusinessPendingOwnedPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingOwnedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessPageRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessPageRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingOwnedPages(ctx context.Context, client *core.Client, id string, params GetBusinessPendingOwnedPagesParams) (*core.Cursor[objects.BusinessPageRequest], error) {
	var out core.Cursor[objects.BusinessPageRequest]
	call := GetBusinessPendingOwnedPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsBatchCall(id string, params GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_shared_offsite_signal_container_business_objects"), params.ToParams(), options...)
}

func NewGetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsBatchRequest(id string, params GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]] {
	return core.NewBatchRequest[core.Cursor[objects.OffsiteSignalContainerBusinessObject]](GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingSharedOffsiteSignalContainerBusinessObjects(ctx context.Context, client *core.Client, id string, params GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsParams) (*core.Cursor[objects.OffsiteSignalContainerBusinessObject], error) {
	var out core.Cursor[objects.OffsiteSignalContainerBusinessObject]
	call := GetBusinessPendingSharedOffsiteSignalContainerBusinessObjectsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPendingUsersBatchCall(id string, params GetBusinessPendingUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pending_users"), params.ToParams(), options...)
}

func NewGetBusinessPendingUsersBatchRequest(id string, params GetBusinessPendingUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessRoleRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessRoleRequest]](GetBusinessPendingUsersBatchCall(id, params, options...))
}

func DecodeGetBusinessPendingUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessRoleRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessRoleRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPendingUsers(ctx context.Context, client *core.Client, id string, params GetBusinessPendingUsersParams) (*core.Cursor[objects.BusinessRoleRequest], error) {
	var out core.Cursor[objects.BusinessRoleRequest]
	call := GetBusinessPendingUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPictureBatchCall(id string, params GetBusinessPictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetBusinessPictureBatchRequest(id string, params GetBusinessPictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetBusinessPictureBatchCall(id, params, options...))
}

func DecodeGetBusinessPictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProfilePictureSource]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPicture(ctx context.Context, client *core.Client, id string, params GetBusinessPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetBusinessPictureBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessPixelTosBatchCall(id string, params CreateBusinessPixelTosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "pixel_tos"), params.ToParams(), options...)
}

func NewCreateBusinessPixelTosBatchRequest(id string, params CreateBusinessPixelTosParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateBusinessPixelTosBatchCall(id, params, options...))
}

func DecodeCreateBusinessPixelTosBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateBusinessPixelTos(ctx context.Context, client *core.Client, id string, params CreateBusinessPixelTosParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateBusinessPixelTosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessPreverifiedNumbersBatchCall(id string, params GetBusinessPreverifiedNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "preverified_numbers"), params.ToParams(), options...)
}

func NewGetBusinessPreverifiedNumbersBatchRequest(id string, params GetBusinessPreverifiedNumbersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber]](GetBusinessPreverifiedNumbersBatchCall(id, params, options...))
}

func DecodeGetBusinessPreverifiedNumbersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessPreverifiedNumbers(ctx context.Context, client *core.Client, id string, params GetBusinessPreverifiedNumbersParams) (*core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber], error) {
	var out core.Cursor[objects.WhatsAppBusinessPreVerifiedPhoneNumber]
	call := GetBusinessPreverifiedNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessReceivedAudienceSharingRequestsBatchCall(id string, params GetBusinessReceivedAudienceSharingRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "received_audience_sharing_requests"), params.ToParams(), options...)
}

func NewGetBusinessReceivedAudienceSharingRequestsBatchRequest(id string, params GetBusinessReceivedAudienceSharingRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetSharingAgreement]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetSharingAgreement]](GetBusinessReceivedAudienceSharingRequestsBatchCall(id, params, options...))
}

func DecodeGetBusinessReceivedAudienceSharingRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessReceivedAudienceSharingRequests(ctx context.Context, client *core.Client, id string, params GetBusinessReceivedAudienceSharingRequestsParams) (*core.Cursor[objects.BusinessAssetSharingAgreement], error) {
	var out core.Cursor[objects.BusinessAssetSharingAgreement]
	call := GetBusinessReceivedAudienceSharingRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessResellerGuidancesBatchCall(id string, params GetBusinessResellerGuidancesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reseller_guidances"), params.ToParams(), options...)
}

func NewGetBusinessResellerGuidancesBatchRequest(id string, params GetBusinessResellerGuidancesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ResellerGuidance]] {
	return core.NewBatchRequest[core.Cursor[objects.ResellerGuidance]](GetBusinessResellerGuidancesBatchCall(id, params, options...))
}

func DecodeGetBusinessResellerGuidancesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ResellerGuidance], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ResellerGuidance]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessResellerGuidances(ctx context.Context, client *core.Client, id string, params GetBusinessResellerGuidancesParams) (*core.Cursor[objects.ResellerGuidance], error) {
	var out core.Cursor[objects.ResellerGuidance]
	call := GetBusinessResellerGuidancesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessSelfCertifiedWhatsappBusinessSubmissionsBatchCall(id string, params GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "self_certified_whatsapp_business_submissions"), params.ToParams(), options...)
}

func NewGetBusinessSelfCertifiedWhatsappBusinessSubmissionsBatchRequest(id string, params GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission]](GetBusinessSelfCertifiedWhatsappBusinessSubmissionsBatchCall(id, params, options...))
}

func DecodeGetBusinessSelfCertifiedWhatsappBusinessSubmissionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessSelfCertifiedWhatsappBusinessSubmissions(ctx context.Context, client *core.Client, id string, params GetBusinessSelfCertifiedWhatsappBusinessSubmissionsParams) (*core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission], error) {
	var out core.Cursor[objects.WhatsAppBusinessPartnerClientVerificationSubmission]
	call := GetBusinessSelfCertifiedWhatsappBusinessSubmissionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessSelfCertifyWhatsappBusinessBatchCall(id string, params CreateBusinessSelfCertifyWhatsappBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "self_certify_whatsapp_business"), params.ToParams(), options...)
}

func NewCreateBusinessSelfCertifyWhatsappBusinessBatchRequest(id string, params CreateBusinessSelfCertifyWhatsappBusinessParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessSelfCertifyWhatsappBusinessBatchCall(id, params, options...))
}

func DecodeCreateBusinessSelfCertifyWhatsappBusinessBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessSelfCertifyWhatsappBusiness(ctx context.Context, client *core.Client, id string, params CreateBusinessSelfCertifyWhatsappBusinessParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessSelfCertifyWhatsappBusinessBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteBusinessSharePreverifiedNumbersBatchCall(id string, params DeleteBusinessSharePreverifiedNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "share_preverified_numbers"), params.ToParams(), options...)
}

func NewDeleteBusinessSharePreverifiedNumbersBatchRequest(id string, params DeleteBusinessSharePreverifiedNumbersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessSharePreverifiedNumbersBatchCall(id, params, options...))
}

func DecodeDeleteBusinessSharePreverifiedNumbersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteBusinessSharePreverifiedNumbers(ctx context.Context, client *core.Client, id string, params DeleteBusinessSharePreverifiedNumbersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessSharePreverifiedNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessSharePreverifiedNumbersBatchCall(id string, params CreateBusinessSharePreverifiedNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "share_preverified_numbers"), params.ToParams(), options...)
}

func NewCreateBusinessSharePreverifiedNumbersBatchRequest(id string, params CreateBusinessSharePreverifiedNumbersParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessSharePreverifiedNumbersBatchCall(id, params, options...))
}

func DecodeCreateBusinessSharePreverifiedNumbersBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessSharePreverifiedNumbers(ctx context.Context, client *core.Client, id string, params CreateBusinessSharePreverifiedNumbersParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessSharePreverifiedNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessSystemUserAccessTokensBatchCall(id string, params CreateBusinessSystemUserAccessTokensParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "system_user_access_tokens"), params.ToParams(), options...)
}

func NewCreateBusinessSystemUserAccessTokensBatchRequest(id string, params CreateBusinessSystemUserAccessTokensParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateBusinessSystemUserAccessTokensBatchCall(id, params, options...))
}

func DecodeCreateBusinessSystemUserAccessTokensBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessSystemUserAccessTokens(ctx context.Context, client *core.Client, id string, params CreateBusinessSystemUserAccessTokensParams) (*objects.Business, error) {
	var out objects.Business
	call := CreateBusinessSystemUserAccessTokensBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessSystemUsersBatchCall(id string, params GetBusinessSystemUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "system_users"), params.ToParams(), options...)
}

func NewGetBusinessSystemUsersBatchRequest(id string, params GetBusinessSystemUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.SystemUser]] {
	return core.NewBatchRequest[core.Cursor[objects.SystemUser]](GetBusinessSystemUsersBatchCall(id, params, options...))
}

func DecodeGetBusinessSystemUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.SystemUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.SystemUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessSystemUsers(ctx context.Context, client *core.Client, id string, params GetBusinessSystemUsersParams) (*core.Cursor[objects.SystemUser], error) {
	var out core.Cursor[objects.SystemUser]
	call := GetBusinessSystemUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessSystemUsersBatchCall(id string, params CreateBusinessSystemUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "system_users"), params.ToParams(), options...)
}

func NewCreateBusinessSystemUsersBatchRequest(id string, params CreateBusinessSystemUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.SystemUser] {
	return core.NewBatchRequest[objects.SystemUser](CreateBusinessSystemUsersBatchCall(id, params, options...))
}

func DecodeCreateBusinessSystemUsersBatchResponse(response *core.BatchResponse) (*objects.SystemUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SystemUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessSystemUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessSystemUsersParams) (*objects.SystemUser, error) {
	var out objects.SystemUser
	call := CreateBusinessSystemUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessThirdPartyMeasurementReportDatasetBatchCall(id string, params GetBusinessThirdPartyMeasurementReportDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "third_party_measurement_report_dataset"), params.ToParams(), options...)
}

func NewGetBusinessThirdPartyMeasurementReportDatasetBatchRequest(id string, params GetBusinessThirdPartyMeasurementReportDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ThirdPartyMeasurementReportDataset]] {
	return core.NewBatchRequest[core.Cursor[objects.ThirdPartyMeasurementReportDataset]](GetBusinessThirdPartyMeasurementReportDatasetBatchCall(id, params, options...))
}

func DecodeGetBusinessThirdPartyMeasurementReportDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ThirdPartyMeasurementReportDataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ThirdPartyMeasurementReportDataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessThirdPartyMeasurementReportDataset(ctx context.Context, client *core.Client, id string, params GetBusinessThirdPartyMeasurementReportDatasetParams) (*core.Cursor[objects.ThirdPartyMeasurementReportDataset], error) {
	var out core.Cursor[objects.ThirdPartyMeasurementReportDataset]
	call := GetBusinessThirdPartyMeasurementReportDatasetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateBusinessVideosBatchCall(id string, params CreateBusinessVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewCreateBusinessVideosBatchRequest(id string, params CreateBusinessVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateBusinessVideosBatchCall(id, params, options...))
}

func DecodeCreateBusinessVideosBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateBusinessVideos(ctx context.Context, client *core.Client, id string, params CreateBusinessVideosParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	call := CreateBusinessVideosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetBusinessBatchCall(id string, params GetBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessBatchRequest(id string, params GetBusinessParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](GetBusinessBatchCall(id, params, options...))
}

func DecodeGetBusinessBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusiness(ctx context.Context, client *core.Client, id string, params GetBusinessParams) (*objects.Business, error) {
	var out objects.Business
	call := GetBusinessBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateBusinessBatchCall(id string, params UpdateBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessBatchRequest(id string, params UpdateBusinessParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](UpdateBusinessBatchCall(id, params, options...))
}

func DecodeUpdateBusinessBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusiness(ctx context.Context, client *core.Client, id string, params UpdateBusinessParams) (*objects.Business, error) {
	var out objects.Business
	call := UpdateBusinessBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
