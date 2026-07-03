package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetWhatsAppBusinessAccountActivitiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountActivitiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountActivitiesBatchCall(id string, params GetWhatsAppBusinessAccountActivitiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "activities"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountActivitiesBatchRequest(id string, params GetWhatsAppBusinessAccountActivitiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessActivityHistory]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessActivityHistory]](GetWhatsAppBusinessAccountActivitiesBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountActivitiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessActivityHistory], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessActivityHistory]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountActivities(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountActivitiesParams) (*core.Cursor[objects.WhatsAppBusinessActivityHistory], error) {
	var out core.Cursor[objects.WhatsAppBusinessActivityHistory]
	call := GetWhatsAppBusinessAccountActivitiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteWhatsAppBusinessAccountAssignedUsersBatchCall(id string, params DeleteWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountAssignedUsersBatchRequest(id string, params DeleteWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetWhatsAppBusinessAccountAssignedUsersBatchCall(id string, params GetWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountAssignedUsersBatchRequest(id string, params GetWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetWhatsAppBusinessAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AssignedUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetWhatsAppBusinessAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountAssignedUsersParams struct {
	Tasks []enums.WhatsappbusinessaccountassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                                        `facebook:"user"`
	Extra core.Params                                                `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tasks"] = params.Tasks
	out["user"] = params.User
	return out
}

func CreateWhatsAppBusinessAccountAssignedUsersBatchCall(id string, params CreateWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountAssignedUsersBatchRequest(id string, params CreateWhatsAppBusinessAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountAssignedUsersParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountAudiencesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountAudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountAudiencesBatchCall(id string, params GetWhatsAppBusinessAccountAudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "audiences"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountAudiencesBatchRequest(id string, params GetWhatsAppBusinessAccountAudiencesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAudience]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAudience]](GetWhatsAppBusinessAccountAudiencesBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountAudiencesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAudience], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAudience]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountAudiences(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountAudiencesParams) (*core.Cursor[objects.WhatsAppBusinessAudience], error) {
	var out core.Cursor[objects.WhatsAppBusinessAudience]
	call := GetWhatsAppBusinessAccountAudiencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams struct {
	Features []map[string]interface{} `facebook:"features"`
	Extra    core.Params              `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["features"] = params.Features
	return out
}

func CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusBatchCall(id string, params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusBatchRequest(id string, params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountCallAnalyticsParams struct {
	CountryCodes *[]string                                                         `facebook:"country_codes"`
	Dimensions   *[]enums.WhatsappbusinessaccountcallAnalyticsDimensionsEnumParam  `facebook:"dimensions"`
	Directions   *[]enums.WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam  `facebook:"directions"`
	End          uint64                                                            `facebook:"end"`
	Granularity  enums.WhatsappbusinessaccountcallAnalyticsGranularityEnumParam    `facebook:"granularity"`
	MetricTypes  *[]enums.WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam `facebook:"metric_types"`
	PhoneNumbers *[]string                                                         `facebook:"phone_numbers"`
	Start        uint64                                                            `facebook:"start"`
	Tiers        *[]string                                                         `facebook:"tiers"`
	Extra        core.Params                                                       `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountCallAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	if params.Directions != nil {
		out["directions"] = *params.Directions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	out["start"] = params.Start
	if params.Tiers != nil {
		out["tiers"] = *params.Tiers
	}
	return out
}

func GetWhatsAppBusinessAccountCallAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountCallAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "call_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountCallAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountCallAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABACallAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABACallAnalytics]](GetWhatsAppBusinessAccountCallAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountCallAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABACallAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABACallAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountCallAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountCallAnalyticsParams) (*core.Cursor[objects.WABACallAnalytics], error) {
	var out core.Cursor[objects.WABACallAnalytics]
	call := GetWhatsAppBusinessAccountCallAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountConversationAnalyticsParams struct {
	ConversationCategories *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam `facebook:"conversation_categories"`
	ConversationDirections *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationDirectionsEnumParam `facebook:"conversation_directions"`
	ConversationTypes      *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationTypesEnumParam      `facebook:"conversation_types"`
	CountryCodes           *[]string                                                                            `facebook:"country_codes"`
	Dimensions             *[]enums.WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam             `facebook:"dimensions"`
	End                    uint64                                                                               `facebook:"end"`
	Granularity            enums.WhatsappbusinessaccountconversationAnalyticsGranularityEnumParam               `facebook:"granularity"`
	MetricTypes            *[]enums.WhatsappbusinessaccountconversationAnalyticsMetricTypesEnumParam            `facebook:"metric_types"`
	PhoneNumbers           *[]string                                                                            `facebook:"phone_numbers"`
	Start                  uint64                                                                               `facebook:"start"`
	Extra                  core.Params                                                                          `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountConversationAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ConversationCategories != nil {
		out["conversation_categories"] = *params.ConversationCategories
	}
	if params.ConversationDirections != nil {
		out["conversation_directions"] = *params.ConversationDirections
	}
	if params.ConversationTypes != nil {
		out["conversation_types"] = *params.ConversationTypes
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	out["start"] = params.Start
	return out
}

func GetWhatsAppBusinessAccountConversationAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountConversationAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversation_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountConversationAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountConversationAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABAConversationAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABAConversationAnalytics]](GetWhatsAppBusinessAccountConversationAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountConversationAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABAConversationAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABAConversationAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountConversationAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountConversationAnalyticsParams) (*core.Cursor[objects.WABAConversationAnalytics], error) {
	var out core.Cursor[objects.WABAConversationAnalytics]
	call := GetWhatsAppBusinessAccountConversationAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountDatasetBatchCall(id string, params GetWhatsAppBusinessAccountDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountDatasetBatchRequest(id string, params GetWhatsAppBusinessAccountDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Dataset]] {
	return core.NewBatchRequest[core.Cursor[objects.Dataset]](GetWhatsAppBusinessAccountDatasetBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Dataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Dataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountDataset(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountDatasetParams) (*core.Cursor[objects.Dataset], error) {
	var out core.Cursor[objects.Dataset]
	call := GetWhatsAppBusinessAccountDatasetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountDatasetParams struct {
	DatasetName *string     `facebook:"dataset_name"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatasetName != nil {
		out["dataset_name"] = *params.DatasetName
	}
	return out
}

func CreateWhatsAppBusinessAccountDatasetBatchCall(id string, params CreateWhatsAppBusinessAccountDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountDatasetBatchRequest(id string, params CreateWhatsAppBusinessAccountDatasetParams, options ...core.BatchOption) *core.BatchRequest[objects.Dataset] {
	return core.NewBatchRequest[objects.Dataset](CreateWhatsAppBusinessAccountDatasetBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountDatasetBatchResponse(response *core.BatchResponse) (*objects.Dataset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Dataset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountDataset(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountDatasetParams) (*objects.Dataset, error) {
	var out objects.Dataset
	call := CreateWhatsAppBusinessAccountDatasetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountDegreesOfFreedomSpecBatchCall(id string, params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "degrees_of_freedom_spec"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountDegreesOfFreedomSpecBatchRequest(id string, params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn]] {
	return core.NewBatchRequest[core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn]](GetWhatsAppBusinessAccountDegreesOfFreedomSpecBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountDegreesOfFreedomSpecBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountDegreesOfFreedomSpec(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams) (*core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn], error) {
	var out core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn]
	call := GetWhatsAppBusinessAccountDegreesOfFreedomSpecBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountFlowsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountFlowsBatchCall(id string, params GetWhatsAppBusinessAccountFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "flows"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountFlowsBatchRequest(id string, params GetWhatsAppBusinessAccountFlowsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppExtension]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppExtension]](GetWhatsAppBusinessAccountFlowsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountFlowsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppExtension], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppExtension]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountFlows(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountFlowsParams) (*core.Cursor[objects.WhatsAppExtension], error) {
	var out core.Cursor[objects.WhatsAppExtension]
	call := GetWhatsAppBusinessAccountFlowsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountFlowsParams struct {
	Categories  []enums.WhatsappbusinessaccountflowsCategoriesEnumParam `facebook:"categories"`
	CloneFlowID *core.ID                                                `facebook:"clone_flow_id"`
	EndpointURI *string                                                 `facebook:"endpoint_uri"`
	FlowJSON    *string                                                 `facebook:"flow_json"`
	Name        string                                                  `facebook:"name"`
	Publish     *bool                                                   `facebook:"publish"`
	Extra       core.Params                                             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["categories"] = params.Categories
	if params.CloneFlowID != nil {
		out["clone_flow_id"] = *params.CloneFlowID
	}
	if params.EndpointURI != nil {
		out["endpoint_uri"] = *params.EndpointURI
	}
	if params.FlowJSON != nil {
		out["flow_json"] = *params.FlowJSON
	}
	out["name"] = params.Name
	if params.Publish != nil {
		out["publish"] = *params.Publish
	}
	return out
}

func CreateWhatsAppBusinessAccountFlowsBatchCall(id string, params CreateWhatsAppBusinessAccountFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "flows"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountFlowsBatchRequest(id string, params CreateWhatsAppBusinessAccountFlowsParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppExtension] {
	return core.NewBatchRequest[objects.WhatsAppExtension](CreateWhatsAppBusinessAccountFlowsBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountFlowsBatchResponse(response *core.BatchResponse) (*objects.WhatsAppExtension, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppExtension
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountFlows(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountFlowsParams) (*objects.WhatsAppExtension, error) {
	var out objects.WhatsAppExtension
	call := CreateWhatsAppBusinessAccountFlowsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	RedirectURL       *string     `facebook:"redirect_url"`
	Extra             core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	if params.RedirectURL != nil {
		out["redirect_url"] = *params.RedirectURL
	}
	return out
}

func CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkBatchCall(id string, params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "generate_payment_configuration_oauth_link"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkBatchRequest(id string, params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLink(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountGroupAnalyticsParams struct {
	End         time.Time                                                          `facebook:"end"`
	Granularity *enums.WhatsappbusinessaccountgroupAnalyticsGranularityEnumParam   `facebook:"granularity"`
	GroupIds    []core.ID                                                          `facebook:"group_ids"`
	MetricTypes *[]enums.WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam `facebook:"metric_types"`
	Start       time.Time                                                          `facebook:"start"`
	Extra       core.Params                                                        `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountGroupAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	if params.Granularity != nil {
		out["granularity"] = *params.Granularity
	}
	out["group_ids"] = params.GroupIds
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	out["start"] = params.Start
	return out
}

func GetWhatsAppBusinessAccountGroupAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountGroupAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "group_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountGroupAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountGroupAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABAGroupAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABAGroupAnalytics]](GetWhatsAppBusinessAccountGroupAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountGroupAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABAGroupAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABAGroupAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountGroupAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountGroupAnalyticsParams) (*core.Cursor[objects.WABAGroupAnalytics], error) {
	var out core.Cursor[objects.WABAGroupAnalytics]
	call := GetWhatsAppBusinessAccountGroupAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountMessageCampaignsBatchCall(id string, params GetWhatsAppBusinessAccountMessageCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "message_campaigns"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountMessageCampaignsBatchRequest(id string, params GetWhatsAppBusinessAccountMessageCampaignsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessCampaign]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessCampaign]](GetWhatsAppBusinessAccountMessageCampaignsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountMessageCampaignsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessCampaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessCampaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountMessageCampaigns(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageCampaignsParams) (*core.Cursor[objects.WhatsAppBusinessCampaign], error) {
	var out core.Cursor[objects.WhatsAppBusinessCampaign]
	call := GetWhatsAppBusinessAccountMessageCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMessageSamplesParams struct {
	Interactive *map[string]interface{}                                  `facebook:"interactive"`
	Text        *map[string]interface{}                                  `facebook:"text"`
	Type        enums.WhatsappbusinessaccountmessageSamplesTypeEnumParam `facebook:"type"`
	Extra       core.Params                                              `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMessageSamplesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Interactive != nil {
		out["interactive"] = *params.Interactive
	}
	if params.Text != nil {
		out["text"] = *params.Text
	}
	out["type"] = params.Type
	return out
}

func CreateWhatsAppBusinessAccountMessageSamplesBatchCall(id string, params CreateWhatsAppBusinessAccountMessageSamplesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "message_samples"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountMessageSamplesBatchRequest(id string, params CreateWhatsAppBusinessAccountMessageSamplesParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountMessageSamplesBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountMessageSamplesBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountMessageSamples(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMessageSamplesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountMessageSamplesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageTemplatePreviewsParams struct {
	AddSecurityRecommendation *bool                                                                       `facebook:"add_security_recommendation"`
	BusinessName              *string                                                                     `facebook:"business_name"`
	ButtonTypes               *[]enums.WhatsappbusinessaccountmessageTemplatePreviewsButtonTypesEnumParam `facebook:"button_types"`
	Category                  enums.WhatsappbusinessaccountmessageTemplatePreviewsCategoryEnumParam       `facebook:"category"`
	CodeExpirationMinutes     *uint64                                                                     `facebook:"code_expiration_minutes"`
	Languages                 *[]string                                                                   `facebook:"languages"`
	Extra                     core.Params                                                                 `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AddSecurityRecommendation != nil {
		out["add_security_recommendation"] = *params.AddSecurityRecommendation
	}
	if params.BusinessName != nil {
		out["business_name"] = *params.BusinessName
	}
	if params.ButtonTypes != nil {
		out["button_types"] = *params.ButtonTypes
	}
	out["category"] = params.Category
	if params.CodeExpirationMinutes != nil {
		out["code_expiration_minutes"] = *params.CodeExpirationMinutes
	}
	if params.Languages != nil {
		out["languages"] = *params.Languages
	}
	return out
}

func GetWhatsAppBusinessAccountMessageTemplatePreviewsBatchCall(id string, params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "message_template_previews"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountMessageTemplatePreviewsBatchRequest(id string, params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview]](GetWhatsAppBusinessAccountMessageTemplatePreviewsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountMessageTemplatePreviewsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountMessageTemplatePreviews(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams) (*core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview]
	call := GetWhatsAppBusinessAccountMessageTemplatePreviewsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountMessageTemplatesParams struct {
	HsmID  *core.ID    `facebook:"hsm_id"`
	HsmIds *[]core.ID  `facebook:"hsm_ids"`
	Name   *string     `facebook:"name"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.HsmID != nil {
		out["hsm_id"] = *params.HsmID
	}
	if params.HsmIds != nil {
		out["hsm_ids"] = *params.HsmIds
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func DeleteWhatsAppBusinessAccountMessageTemplatesBatchCall(id string, params DeleteWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountMessageTemplatesBatchRequest(id string, params DeleteWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountMessageTemplatesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountMessageTemplatesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageTemplatesParams struct {
	Category      *[]enums.WhatsappbusinessaccountmessageTemplatesCategoryEnumParam     `facebook:"category"`
	Content       *string                                                               `facebook:"content"`
	Language      *[]string                                                             `facebook:"language"`
	Name          *string                                                               `facebook:"name"`
	NameOrContent *string                                                               `facebook:"name_or_content"`
	QualityScore  *[]enums.WhatsappbusinessaccountmessageTemplatesQualityScoreEnumParam `facebook:"quality_score"`
	Since         *time.Time                                                            `facebook:"since"`
	Source        *enums.WhatsappbusinessaccountmessageTemplatesSourceEnumParam         `facebook:"source"`
	Status        *[]enums.WhatsappbusinessaccountmessageTemplatesStatusEnumParam       `facebook:"status"`
	Until         *time.Time                                                            `facebook:"until"`
	Extra         core.Params                                                           `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.Content != nil {
		out["content"] = *params.Content
	}
	if params.Language != nil {
		out["language"] = *params.Language
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NameOrContent != nil {
		out["name_or_content"] = *params.NameOrContent
	}
	if params.QualityScore != nil {
		out["quality_score"] = *params.QualityScore
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

func GetWhatsAppBusinessAccountMessageTemplatesBatchCall(id string, params GetWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountMessageTemplatesBatchRequest(id string, params GetWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessHSM]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessHSM]](GetWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountMessageTemplatesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessHSM], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessHSM]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageTemplatesParams) (*core.Cursor[objects.WhatsAppBusinessHSM], error) {
	var out core.Cursor[objects.WhatsAppBusinessHSM]
	call := GetWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMessageTemplatesParams struct {
	AllowCategoryChange         *bool                                                                  `facebook:"allow_category_change"`
	BidSpec                     *map[string]interface{}                                                `facebook:"bid_spec"`
	Category                    enums.WhatsappbusinessaccountmessageTemplatesCategoryEnumParam         `facebook:"category"`
	Components                  *[]map[string]interface{}                                              `facebook:"components"`
	CreativeSourcingSpec        *map[string]interface{}                                                `facebook:"creative_sourcing_spec"`
	CtaURLLinkTrackingOptedOut  *bool                                                                  `facebook:"cta_url_link_tracking_opted_out"`
	DegreesOfFreedomSpec        *map[string]interface{}                                                `facebook:"degrees_of_freedom_spec"`
	DisplayFormat               *enums.WhatsappbusinessaccountmessageTemplatesDisplayFormatEnumParam   `facebook:"display_format"`
	IsPrimaryDeviceDeliveryOnly *bool                                                                  `facebook:"is_primary_device_delivery_only"`
	Language                    string                                                                 `facebook:"language"`
	LibraryTemplateBodyInputs   *map[string]interface{}                                                `facebook:"library_template_body_inputs"`
	LibraryTemplateButtonInputs *[]map[string]interface{}                                              `facebook:"library_template_button_inputs"`
	LibraryTemplateName         *string                                                                `facebook:"library_template_name"`
	MessageSendTTLSeconds       *uint64                                                                `facebook:"message_send_ttl_seconds"`
	Name                        string                                                                 `facebook:"name"`
	OptimizationSpec            *map[string]interface{}                                                `facebook:"optimization_spec"`
	ParameterFormat             *enums.WhatsappbusinessaccountmessageTemplatesParameterFormatEnumParam `facebook:"parameter_format"`
	ProductSetID                *core.ID                                                               `facebook:"product_set_id"`
	SendType                    *enums.WhatsappbusinessaccountmessageTemplatesSendTypeEnumParam        `facebook:"send_type"`
	SubCategory                 *enums.WhatsappbusinessaccountmessageTemplatesSubCategoryEnumParam     `facebook:"sub_category"`
	Extra                       core.Params                                                            `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowCategoryChange != nil {
		out["allow_category_change"] = *params.AllowCategoryChange
	}
	if params.BidSpec != nil {
		out["bid_spec"] = *params.BidSpec
	}
	out["category"] = params.Category
	if params.Components != nil {
		out["components"] = *params.Components
	}
	if params.CreativeSourcingSpec != nil {
		out["creative_sourcing_spec"] = *params.CreativeSourcingSpec
	}
	if params.CtaURLLinkTrackingOptedOut != nil {
		out["cta_url_link_tracking_opted_out"] = *params.CtaURLLinkTrackingOptedOut
	}
	if params.DegreesOfFreedomSpec != nil {
		out["degrees_of_freedom_spec"] = *params.DegreesOfFreedomSpec
	}
	if params.DisplayFormat != nil {
		out["display_format"] = *params.DisplayFormat
	}
	if params.IsPrimaryDeviceDeliveryOnly != nil {
		out["is_primary_device_delivery_only"] = *params.IsPrimaryDeviceDeliveryOnly
	}
	out["language"] = params.Language
	if params.LibraryTemplateBodyInputs != nil {
		out["library_template_body_inputs"] = *params.LibraryTemplateBodyInputs
	}
	if params.LibraryTemplateButtonInputs != nil {
		out["library_template_button_inputs"] = *params.LibraryTemplateButtonInputs
	}
	if params.LibraryTemplateName != nil {
		out["library_template_name"] = *params.LibraryTemplateName
	}
	if params.MessageSendTTLSeconds != nil {
		out["message_send_ttl_seconds"] = *params.MessageSendTTLSeconds
	}
	out["name"] = params.Name
	if params.OptimizationSpec != nil {
		out["optimization_spec"] = *params.OptimizationSpec
	}
	if params.ParameterFormat != nil {
		out["parameter_format"] = *params.ParameterFormat
	}
	if params.ProductSetID != nil {
		out["product_set_id"] = *params.ProductSetID
	}
	if params.SendType != nil {
		out["send_type"] = *params.SendType
	}
	if params.SubCategory != nil {
		out["sub_category"] = *params.SubCategory
	}
	return out
}

func CreateWhatsAppBusinessAccountMessageTemplatesBatchCall(id string, params CreateWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountMessageTemplatesBatchRequest(id string, params CreateWhatsAppBusinessAccountMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountMessageTemplatesBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountMessageTemplatesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMigrateFlowsParams struct {
	SourceFlowNames *[]string   `facebook:"source_flow_names"`
	SourceWabaID    core.ID     `facebook:"source_waba_id"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMigrateFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SourceFlowNames != nil {
		out["source_flow_names"] = *params.SourceFlowNames
	}
	out["source_waba_id"] = params.SourceWabaID
	return out
}

func CreateWhatsAppBusinessAccountMigrateFlowsBatchCall(id string, params CreateWhatsAppBusinessAccountMigrateFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "migrate_flows"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountMigrateFlowsBatchRequest(id string, params CreateWhatsAppBusinessAccountMigrateFlowsParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountMigrateFlowsBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountMigrateFlowsBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountMigrateFlows(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMigrateFlowsParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountMigrateFlowsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams struct {
	Count        *uint64     `facebook:"count"`
	PageNumber   *uint64     `facebook:"page_number"`
	SourceWabaID core.ID     `facebook:"source_waba_id"`
	TemplateIds  *[]core.ID  `facebook:"template_ids"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Count != nil {
		out["count"] = *params.Count
	}
	if params.PageNumber != nil {
		out["page_number"] = *params.PageNumber
	}
	out["source_waba_id"] = params.SourceWabaID
	if params.TemplateIds != nil {
		out["template_ids"] = *params.TemplateIds
	}
	return out
}

func CreateWhatsAppBusinessAccountMigrateMessageTemplatesBatchCall(id string, params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "migrate_message_templates"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountMigrateMessageTemplatesBatchRequest(id string, params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountMigrateMessageTemplatesBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountMigrateMessageTemplatesBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountMigrateMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountMigrateMessageTemplatesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	Extra             core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	return out
}

func DeleteWhatsAppBusinessAccountPaymentConfigurationBatchCall(id string, params DeleteWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "payment_configuration"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountPaymentConfigurationBatchRequest(id string, params DeleteWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountPaymentConfigurationBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountPaymentConfigurationParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	Extra             core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	return out
}

func GetWhatsAppBusinessAccountPaymentConfigurationBatchCall(id string, params GetWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "payment_configuration"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountPaymentConfigurationBatchRequest(id string, params GetWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration]](GetWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountPaymentConfigurationBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPaymentConfigurationParams) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration]
	call := GetWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName    string                                                                  `facebook:"configuration_name"`
	DataEndpointURL      *string                                                                 `facebook:"data_endpoint_url"`
	MerchantCategoryCode *string                                                                 `facebook:"merchant_category_code"`
	MerchantVpa          *string                                                                 `facebook:"merchant_vpa"`
	ProviderName         *enums.WhatsappbusinessaccountpaymentConfigurationProviderNameEnumParam `facebook:"provider_name"`
	PurposeCode          *string                                                                 `facebook:"purpose_code"`
	RedirectURL          *string                                                                 `facebook:"redirect_url"`
	Extra                core.Params                                                             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	if params.DataEndpointURL != nil {
		out["data_endpoint_url"] = *params.DataEndpointURL
	}
	if params.MerchantCategoryCode != nil {
		out["merchant_category_code"] = *params.MerchantCategoryCode
	}
	if params.MerchantVpa != nil {
		out["merchant_vpa"] = *params.MerchantVpa
	}
	if params.ProviderName != nil {
		out["provider_name"] = *params.ProviderName
	}
	if params.PurposeCode != nil {
		out["purpose_code"] = *params.PurposeCode
	}
	if params.RedirectURL != nil {
		out["redirect_url"] = *params.RedirectURL
	}
	return out
}

func CreateWhatsAppBusinessAccountPaymentConfigurationBatchCall(id string, params CreateWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "payment_configuration"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountPaymentConfigurationBatchRequest(id string, params CreateWhatsAppBusinessAccountPaymentConfigurationParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountPaymentConfigurationBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountPaymentConfigurationParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountPaymentConfigurationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPaymentConfigurationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPaymentConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountPaymentConfigurationsBatchCall(id string, params GetWhatsAppBusinessAccountPaymentConfigurationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "payment_configurations"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountPaymentConfigurationsBatchRequest(id string, params GetWhatsAppBusinessAccountPaymentConfigurationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations]](GetWhatsAppBusinessAccountPaymentConfigurationsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountPaymentConfigurationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountPaymentConfigurations(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPaymentConfigurationsParams) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations]
	call := GetWhatsAppBusinessAccountPaymentConfigurationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPhoneNumbersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPhoneNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountPhoneNumbersBatchCall(id string, params GetWhatsAppBusinessAccountPhoneNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "phone_numbers"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountPhoneNumbersBatchRequest(id string, params GetWhatsAppBusinessAccountPhoneNumbersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus]](GetWhatsAppBusinessAccountPhoneNumbersBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountPhoneNumbersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountPhoneNumbers(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPhoneNumbersParams) (*core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus]
	call := GetWhatsAppBusinessAccountPhoneNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountPhoneNumbersParams struct {
	Cc                 *string     `facebook:"cc"`
	MigratePhoneNumber *bool       `facebook:"migrate_phone_number"`
	PhoneNumber        *string     `facebook:"phone_number"`
	PreverifiedID      *core.ID    `facebook:"preverified_id"`
	VerifiedName       *string     `facebook:"verified_name"`
	Extra              core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountPhoneNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cc != nil {
		out["cc"] = *params.Cc
	}
	if params.MigratePhoneNumber != nil {
		out["migrate_phone_number"] = *params.MigratePhoneNumber
	}
	if params.PhoneNumber != nil {
		out["phone_number"] = *params.PhoneNumber
	}
	if params.PreverifiedID != nil {
		out["preverified_id"] = *params.PreverifiedID
	}
	if params.VerifiedName != nil {
		out["verified_name"] = *params.VerifiedName
	}
	return out
}

func CreateWhatsAppBusinessAccountPhoneNumbersBatchCall(id string, params CreateWhatsAppBusinessAccountPhoneNumbersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "phone_numbers"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountPhoneNumbersBatchRequest(id string, params CreateWhatsAppBusinessAccountPhoneNumbersParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccountToNumberCurrentStatus] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccountToNumberCurrentStatus](CreateWhatsAppBusinessAccountPhoneNumbersBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountPhoneNumbersBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccountToNumberCurrentStatus, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccountToNumberCurrentStatus
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountPhoneNumbers(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountPhoneNumbersParams) (*objects.WhatsAppBusinessAccountToNumberCurrentStatus, error) {
	var out objects.WhatsAppBusinessAccountToNumberCurrentStatus
	call := CreateWhatsAppBusinessAccountPhoneNumbersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPricingAnalyticsParams struct {
	CountryCodes      *[]string                                                                  `facebook:"country_codes"`
	Dimensions        *[]enums.WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam        `facebook:"dimensions"`
	End               uint64                                                                     `facebook:"end"`
	Granularity       enums.WhatsappbusinessaccountpricingAnalyticsGranularityEnumParam          `facebook:"granularity"`
	MetricTypes       *[]enums.WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam       `facebook:"metric_types"`
	PhoneNumbers      *[]string                                                                  `facebook:"phone_numbers"`
	PricingCategories *[]enums.WhatsappbusinessaccountpricingAnalyticsPricingCategoriesEnumParam `facebook:"pricing_categories"`
	PricingTypes      *[]enums.WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam      `facebook:"pricing_types"`
	Start             uint64                                                                     `facebook:"start"`
	Tiers             *[]string                                                                  `facebook:"tiers"`
	Extra             core.Params                                                                `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPricingAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	if params.PricingCategories != nil {
		out["pricing_categories"] = *params.PricingCategories
	}
	if params.PricingTypes != nil {
		out["pricing_types"] = *params.PricingTypes
	}
	out["start"] = params.Start
	if params.Tiers != nil {
		out["tiers"] = *params.Tiers
	}
	return out
}

func GetWhatsAppBusinessAccountPricingAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountPricingAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pricing_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountPricingAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountPricingAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABAPricingAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABAPricingAnalytics]](GetWhatsAppBusinessAccountPricingAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountPricingAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABAPricingAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABAPricingAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountPricingAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPricingAnalyticsParams) (*core.Cursor[objects.WABAPricingAnalytics], error) {
	var out core.Cursor[objects.WABAPricingAnalytics]
	call := GetWhatsAppBusinessAccountPricingAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountProductCatalogsParams struct {
	CatalogID core.ID     `facebook:"catalog_id"`
	Extra     core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["catalog_id"] = params.CatalogID
	return out
}

func DeleteWhatsAppBusinessAccountProductCatalogsBatchCall(id string, params DeleteWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "product_catalogs"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountProductCatalogsBatchRequest(id string, params DeleteWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountProductCatalogsBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountProductCatalogsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountProductCatalogsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountProductCatalogsBatchCall(id string, params GetWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_catalogs"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountProductCatalogsBatchRequest(id string, params GetWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetWhatsAppBusinessAccountProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
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

func GetWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetWhatsAppBusinessAccountProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountProductCatalogsParams struct {
	CatalogID core.ID     `facebook:"catalog_id"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["catalog_id"] = params.CatalogID
	return out
}

func CreateWhatsAppBusinessAccountProductCatalogsBatchCall(id string, params CreateWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "product_catalogs"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountProductCatalogsBatchRequest(id string, params CreateWhatsAppBusinessAccountProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductCatalog] {
	return core.NewBatchRequest[objects.ProductCatalog](CreateWhatsAppBusinessAccountProductCatalogsBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountProductCatalogsBatchResponse(response *core.BatchResponse) (*objects.ProductCatalog, error) {
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

func CreateWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountProductCatalogsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	call := CreateWhatsAppBusinessAccountProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSchedulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSchedulesBatchCall(id string, params GetWhatsAppBusinessAccountSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "schedules"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountSchedulesBatchRequest(id string, params GetWhatsAppBusinessAccountSchedulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessCampaignSchedule]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessCampaignSchedule]](GetWhatsAppBusinessAccountSchedulesBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountSchedulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessCampaignSchedule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessCampaignSchedule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountSchedules(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSchedulesParams) (*core.Cursor[objects.WhatsAppBusinessCampaignSchedule], error) {
	var out core.Cursor[objects.WhatsAppBusinessCampaignSchedule]
	call := GetWhatsAppBusinessAccountSchedulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams struct {
	AppID      *core.ID    `facebook:"app_id"`
	SolutionID *core.ID    `facebook:"solution_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.SolutionID != nil {
		out["solution_id"] = *params.SolutionID
	}
	return out
}

func CreateWhatsAppBusinessAccountSetSolutionMigrationIntentBatchCall(id string, params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "set_solution_migration_intent"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountSetSolutionMigrationIntentBatchRequest(id string, params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccountMigrationIntent] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccountMigrationIntent](CreateWhatsAppBusinessAccountSetSolutionMigrationIntentBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountSetSolutionMigrationIntentBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccountMigrationIntent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccountMigrationIntent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountSetSolutionMigrationIntent(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams) (*objects.WhatsAppBusinessAccountMigrationIntent, error) {
	var out objects.WhatsAppBusinessAccountMigrationIntent
	call := CreateWhatsAppBusinessAccountSetSolutionMigrationIntentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSolutionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSolutionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSolutionsBatchCall(id string, params GetWhatsAppBusinessAccountSolutionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "solutions"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountSolutionsBatchRequest(id string, params GetWhatsAppBusinessAccountSolutionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessSolution]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessSolution]](GetWhatsAppBusinessAccountSolutionsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountSolutionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountSolutions(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSolutionsParams) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	call := GetWhatsAppBusinessAccountSolutionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteWhatsAppBusinessAccountSubscribedAppsBatchCall(id string, params DeleteWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountSubscribedAppsBatchRequest(id string, params DeleteWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountSubscribedAppsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSubscribedAppsBatchCall(id string, params GetWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountSubscribedAppsBatchRequest(id string, params GetWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppApplication]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppApplication]](GetWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppApplication], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppApplication]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSubscribedAppsParams) (*core.Cursor[objects.WhatsAppApplication], error) {
	var out core.Cursor[objects.WhatsAppApplication]
	call := GetWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountSubscribedAppsParams struct {
	OverrideCallbackURI *string     `facebook:"override_callback_uri"`
	VerifyToken         *string     `facebook:"verify_token"`
	Extra               core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.OverrideCallbackURI != nil {
		out["override_callback_uri"] = *params.OverrideCallbackURI
	}
	if params.VerifyToken != nil {
		out["verify_token"] = *params.VerifyToken
	}
	return out
}

func CreateWhatsAppBusinessAccountSubscribedAppsBatchCall(id string, params CreateWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountSubscribedAppsBatchRequest(id string, params CreateWhatsAppBusinessAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountSubscribedAppsParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateAnalyticsParams struct {
	End             time.Time                                                           `facebook:"end"`
	Granularity     enums.WhatsappbusinessaccounttemplateAnalyticsGranularityEnumParam  `facebook:"granularity"`
	MetricTypes     *[]string                                                           `facebook:"metric_types"`
	ProductType     *enums.WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam `facebook:"product_type"`
	Start           time.Time                                                           `facebook:"start"`
	TemplateIds     []core.ID                                                           `facebook:"template_ids"`
	UseWabaTimezone *bool                                                               `facebook:"use_waba_timezone"`
	Extra           core.Params                                                         `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	out["start"] = params.Start
	out["template_ids"] = params.TemplateIds
	if params.UseWabaTimezone != nil {
		out["use_waba_timezone"] = *params.UseWabaTimezone
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountTemplateAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "template_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountTemplateAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountTemplateAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABATemplateAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABATemplateAnalytics]](GetWhatsAppBusinessAccountTemplateAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountTemplateAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABATemplateAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABATemplateAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountTemplateAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateAnalyticsParams) (*core.Cursor[objects.WABATemplateAnalytics], error) {
	var out core.Cursor[objects.WABATemplateAnalytics]
	call := GetWhatsAppBusinessAccountTemplateAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams struct {
	End              time.Time                                                               `facebook:"end"`
	Granularity      enums.WhatsappbusinessaccounttemplateGroupAnalyticsGranularityEnumParam `facebook:"granularity"`
	MetricTypes      *[]string                                                               `facebook:"metric_types"`
	Start            time.Time                                                               `facebook:"start"`
	TemplateGroupIds []core.ID                                                               `facebook:"template_group_ids"`
	UseWabaTimezone  *bool                                                                   `facebook:"use_waba_timezone"`
	Extra            core.Params                                                             `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	out["start"] = params.Start
	out["template_group_ids"] = params.TemplateGroupIds
	if params.UseWabaTimezone != nil {
		out["use_waba_timezone"] = *params.UseWabaTimezone
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateGroupAnalyticsBatchCall(id string, params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "template_group_analytics"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountTemplateGroupAnalyticsBatchRequest(id string, params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WABATemplateGroupAnalytics]] {
	return core.NewBatchRequest[core.Cursor[objects.WABATemplateGroupAnalytics]](GetWhatsAppBusinessAccountTemplateGroupAnalyticsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountTemplateGroupAnalyticsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WABATemplateGroupAnalytics], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WABATemplateGroupAnalytics]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountTemplateGroupAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams) (*core.Cursor[objects.WABATemplateGroupAnalytics], error) {
	var out core.Cursor[objects.WABATemplateGroupAnalytics]
	call := GetWhatsAppBusinessAccountTemplateGroupAnalyticsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateGroupsBatchCall(id string, params GetWhatsAppBusinessAccountTemplateGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "template_groups"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountTemplateGroupsBatchRequest(id string, params GetWhatsAppBusinessAccountTemplateGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessMessagingTemplateGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessMessagingTemplateGroup]](GetWhatsAppBusinessAccountTemplateGroupsBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountTemplateGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessMessagingTemplateGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessMessagingTemplateGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountTemplateGroups(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateGroupsParams) (*core.Cursor[objects.BusinessMessagingTemplateGroup], error) {
	var out core.Cursor[objects.BusinessMessagingTemplateGroup]
	call := GetWhatsAppBusinessAccountTemplateGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountTemplateGroupsParams struct {
	Description               string      `facebook:"description"`
	Name                      string      `facebook:"name"`
	WhatsappBusinessTemplates []string    `facebook:"whatsapp_business_templates"`
	Extra                     core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountTemplateGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["description"] = params.Description
	out["name"] = params.Name
	out["whatsapp_business_templates"] = params.WhatsappBusinessTemplates
	return out
}

func CreateWhatsAppBusinessAccountTemplateGroupsBatchCall(id string, params CreateWhatsAppBusinessAccountTemplateGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "template_groups"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountTemplateGroupsBatchRequest(id string, params CreateWhatsAppBusinessAccountTemplateGroupsParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessMessagingTemplateGroup] {
	return core.NewBatchRequest[objects.BusinessMessagingTemplateGroup](CreateWhatsAppBusinessAccountTemplateGroupsBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountTemplateGroupsBatchResponse(response *core.BatchResponse) (*objects.BusinessMessagingTemplateGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessMessagingTemplateGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountTemplateGroups(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountTemplateGroupsParams) (*objects.BusinessMessagingTemplateGroup, error) {
	var out objects.BusinessMessagingTemplateGroup
	call := CreateWhatsAppBusinessAccountTemplateGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams struct {
	Category              enums.WhatsappbusinessaccountupsertMessageTemplatesCategoryEnumParam `facebook:"category"`
	Components            []map[string]interface{}                                             `facebook:"components"`
	Languages             []string                                                             `facebook:"languages"`
	MessageSendTTLSeconds *uint64                                                              `facebook:"message_send_ttl_seconds"`
	Name                  string                                                               `facebook:"name"`
	Extra                 core.Params                                                          `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["category"] = params.Category
	out["components"] = params.Components
	out["languages"] = params.Languages
	if params.MessageSendTTLSeconds != nil {
		out["message_send_ttl_seconds"] = *params.MessageSendTTLSeconds
	}
	out["name"] = params.Name
	return out
}

func CreateWhatsAppBusinessAccountUpsertMessageTemplatesBatchCall(id string, params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "upsert_message_templates"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountUpsertMessageTemplatesBatchRequest(id string, params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](CreateWhatsAppBusinessAccountUpsertMessageTemplatesBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountUpsertMessageTemplatesBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateWhatsAppBusinessAccountUpsertMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := CreateWhatsAppBusinessAccountUpsertMessageTemplatesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	SequenceID core.ID     `facebook:"sequence_id"`
	Extra      core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["sequence_id"] = params.SequenceID
	return out
}

func DeleteWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id string, params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), options...)
}

func NewDeleteWhatsAppBusinessAccountWelcomeMessageSequencesBatchRequest(id string, params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params, options...))
}

func DecodeDeleteWhatsAppBusinessAccountWelcomeMessageSequencesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	AppID      *core.ID    `facebook:"app_id"`
	SequenceID *core.ID    `facebook:"sequence_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.SequenceID != nil {
		out["sequence_id"] = *params.SequenceID
	}
	return out
}

func GetWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id string, params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountWelcomeMessageSequencesBatchRequest(id string, params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]] {
	return core.NewBatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]](GetWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountWelcomeMessageSequencesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	call := GetWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	Name                   *string                 `facebook:"name"`
	SequenceID             *core.ID                `facebook:"sequence_id"`
	WelcomeMessageSequence *map[string]interface{} `facebook:"welcome_message_sequence"`
	Extra                  core.Params             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SequenceID != nil {
		out["sequence_id"] = *params.SequenceID
	}
	if params.WelcomeMessageSequence != nil {
		out["welcome_message_sequence"] = *params.WelcomeMessageSequence
	}
	return out
}

func CreateWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id string, params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), options...)
}

func NewCreateWhatsAppBusinessAccountWelcomeMessageSequencesBatchRequest(id string, params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params, options...))
}

func DecodeCreateWhatsAppBusinessAccountWelcomeMessageSequencesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateWhatsAppBusinessAccountWelcomeMessageSequencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetWhatsAppBusinessAccountParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountBatchCall(id string, params GetWhatsAppBusinessAccountParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessAccountBatchRequest(id string, params GetWhatsAppBusinessAccountParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](GetWhatsAppBusinessAccountBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessAccountBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessAccount(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := GetWhatsAppBusinessAccountBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateWhatsAppBusinessAccountParams struct {
	CreativeOptimizationsEnrollment    *map[string]interface{} `facebook:"creative_optimizations_enrollment"`
	DegreesOfFreedomSpec               *map[string]interface{} `facebook:"degrees_of_freedom_spec"`
	DisableMarketingMessagesOnCloudAPI *bool                   `facebook:"disable_marketing_messages_on_cloud_api"`
	IsEnabledForInsights               *bool                   `facebook:"is_enabled_for_insights"`
	TemplateAutoArchivalEnabled        *bool                   `facebook:"template_auto_archival_enabled"`
	Extra                              core.Params             `facebook:"-"`
}

func (params UpdateWhatsAppBusinessAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CreativeOptimizationsEnrollment != nil {
		out["creative_optimizations_enrollment"] = *params.CreativeOptimizationsEnrollment
	}
	if params.DegreesOfFreedomSpec != nil {
		out["degrees_of_freedom_spec"] = *params.DegreesOfFreedomSpec
	}
	if params.DisableMarketingMessagesOnCloudAPI != nil {
		out["disable_marketing_messages_on_cloud_api"] = *params.DisableMarketingMessagesOnCloudAPI
	}
	if params.IsEnabledForInsights != nil {
		out["is_enabled_for_insights"] = *params.IsEnabledForInsights
	}
	if params.TemplateAutoArchivalEnabled != nil {
		out["template_auto_archival_enabled"] = *params.TemplateAutoArchivalEnabled
	}
	return out
}

func UpdateWhatsAppBusinessAccountBatchCall(id string, params UpdateWhatsAppBusinessAccountParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateWhatsAppBusinessAccountBatchRequest(id string, params UpdateWhatsAppBusinessAccountParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessAccount] {
	return core.NewBatchRequest[objects.WhatsAppBusinessAccount](UpdateWhatsAppBusinessAccountBatchCall(id, params, options...))
}

func DecodeUpdateWhatsAppBusinessAccountBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateWhatsAppBusinessAccount(ctx context.Context, client *core.Client, id string, params UpdateWhatsAppBusinessAccountParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	call := UpdateWhatsAppBusinessAccountBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
