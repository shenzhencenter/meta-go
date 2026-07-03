package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type DeleteCustomAudienceAdaccountsParams struct {
	Adaccounts *[]string   `facebook:"adaccounts"`
	Extra      core.Params `facebook:"-"`
}

func (params DeleteCustomAudienceAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adaccounts != nil {
		out["adaccounts"] = *params.Adaccounts
	}
	return out
}

func DeleteCustomAudienceAdaccountsBatchCall(id string, params DeleteCustomAudienceAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewDeleteCustomAudienceAdaccountsBatchRequest(id string, params DeleteCustomAudienceAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCustomAudienceAdaccountsBatchCall(id, params, options...))
}

func DecodeDeleteCustomAudienceAdaccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceAdaccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteCustomAudienceAdaccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceAdaccountsParams struct {
	Permissions *string     `facebook:"permissions"`
	Extra       core.Params `facebook:"-"`
}

func (params GetCustomAudienceAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Permissions != nil {
		out["permissions"] = *params.Permissions
	}
	return out
}

func GetCustomAudienceAdaccountsBatchCall(id string, params GetCustomAudienceAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetCustomAudienceAdaccountsBatchRequest(id string, params GetCustomAudienceAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetCustomAudienceAdaccountsBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params GetCustomAudienceAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetCustomAudienceAdaccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateCustomAudienceAdaccountsParams struct {
	Adaccounts       *[]string   `facebook:"adaccounts"`
	Permissions      *string     `facebook:"permissions"`
	RelationshipType *[]string   `facebook:"relationship_type"`
	Replace          *bool       `facebook:"replace"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateCustomAudienceAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adaccounts != nil {
		out["adaccounts"] = *params.Adaccounts
	}
	if params.Permissions != nil {
		out["permissions"] = *params.Permissions
	}
	if params.RelationshipType != nil {
		out["relationship_type"] = *params.RelationshipType
	}
	if params.Replace != nil {
		out["replace"] = *params.Replace
	}
	return out
}

func CreateCustomAudienceAdaccountsBatchCall(id string, params CreateCustomAudienceAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewCreateCustomAudienceAdaccountsBatchRequest(id string, params CreateCustomAudienceAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateCustomAudienceAdaccountsBatchCall(id, params, options...))
}

func DecodeCreateCustomAudienceAdaccountsBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceAdaccountsParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateCustomAudienceAdaccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceAdsParams struct {
	EffectiveStatus *[]string   `facebook:"effective_status"`
	Status          *[]string   `facebook:"status"`
	Extra           core.Params `facebook:"-"`
}

func (params GetCustomAudienceAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EffectiveStatus != nil {
		out["effective_status"] = *params.EffectiveStatus
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetCustomAudienceAdsBatchCall(id string, params GetCustomAudienceAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), options...)
}

func NewGetCustomAudienceAdsBatchRequest(id string, params GetCustomAudienceAdsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetCustomAudienceAdsBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceAdsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
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

func GetCustomAudienceAds(ctx context.Context, client *core.Client, id string, params GetCustomAudienceAdsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	call := GetCustomAudienceAdsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceHealthParams struct {
	CalculatedDate           *string     `facebook:"calculated_date"`
	ProcessedDate            *string     `facebook:"processed_date"`
	ValueAggregationDuration *uint64     `facebook:"value_aggregation_duration"`
	ValueCountry             *string     `facebook:"value_country"`
	ValueCurrency            *string     `facebook:"value_currency"`
	ValueVersion             *uint64     `facebook:"value_version"`
	Extra                    core.Params `facebook:"-"`
}

func (params GetCustomAudienceHealthParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CalculatedDate != nil {
		out["calculated_date"] = *params.CalculatedDate
	}
	if params.ProcessedDate != nil {
		out["processed_date"] = *params.ProcessedDate
	}
	if params.ValueAggregationDuration != nil {
		out["value_aggregation_duration"] = *params.ValueAggregationDuration
	}
	if params.ValueCountry != nil {
		out["value_country"] = *params.ValueCountry
	}
	if params.ValueCurrency != nil {
		out["value_currency"] = *params.ValueCurrency
	}
	if params.ValueVersion != nil {
		out["value_version"] = *params.ValueVersion
	}
	return out
}

func GetCustomAudienceHealthBatchCall(id string, params GetCustomAudienceHealthParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "health"), params.ToParams(), options...)
}

func NewGetCustomAudienceHealthBatchRequest(id string, params GetCustomAudienceHealthParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudienceHealth]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudienceHealth]](GetCustomAudienceHealthBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceHealthBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudienceHealth], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudienceHealth]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomAudienceHealth(ctx context.Context, client *core.Client, id string, params GetCustomAudienceHealthParams) (*core.Cursor[objects.CustomAudienceHealth], error) {
	var out core.Cursor[objects.CustomAudienceHealth]
	call := GetCustomAudienceHealthBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceSaltsParams struct {
	Params *[]string   `facebook:"params"`
	Extra  core.Params `facebook:"-"`
}

func (params GetCustomAudienceSaltsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Params != nil {
		out["params"] = *params.Params
	}
	return out
}

func GetCustomAudienceSaltsBatchCall(id string, params GetCustomAudienceSaltsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "salts"), params.ToParams(), options...)
}

func NewGetCustomAudienceSaltsBatchRequest(id string, params GetCustomAudienceSaltsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudienceSalts]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudienceSalts]](GetCustomAudienceSaltsBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceSaltsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudienceSalts], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudienceSalts]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomAudienceSalts(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSaltsParams) (*core.Cursor[objects.CustomAudienceSalts], error) {
	var out core.Cursor[objects.CustomAudienceSalts]
	call := GetCustomAudienceSaltsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateCustomAudienceSaltsParams struct {
	Salt      string      `facebook:"salt"`
	ValidFrom time.Time   `facebook:"valid_from"`
	ValidTo   time.Time   `facebook:"valid_to"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateCustomAudienceSaltsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["salt"] = params.Salt
	out["valid_from"] = params.ValidFrom
	out["valid_to"] = params.ValidTo
	return out
}

func CreateCustomAudienceSaltsBatchCall(id string, params CreateCustomAudienceSaltsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "salts"), params.ToParams(), options...)
}

func NewCreateCustomAudienceSaltsBatchRequest(id string, params CreateCustomAudienceSaltsParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateCustomAudienceSaltsBatchCall(id, params, options...))
}

func DecodeCreateCustomAudienceSaltsBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCustomAudienceSalts(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceSaltsParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateCustomAudienceSaltsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceSessionsParams struct {
	SessionID *core.ID    `facebook:"session_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetCustomAudienceSessionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SessionID != nil {
		out["session_id"] = *params.SessionID
	}
	return out
}

func GetCustomAudienceSessionsBatchCall(id string, params GetCustomAudienceSessionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sessions"), params.ToParams(), options...)
}

func NewGetCustomAudienceSessionsBatchRequest(id string, params GetCustomAudienceSessionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudienceSession]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudienceSession]](GetCustomAudienceSessionsBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceSessionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudienceSession], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudienceSession]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomAudienceSessions(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSessionsParams) (*core.Cursor[objects.CustomAudienceSession], error) {
	var out core.Cursor[objects.CustomAudienceSession]
	call := GetCustomAudienceSessionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceSharedAccountInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCustomAudienceSharedAccountInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCustomAudienceSharedAccountInfoBatchCall(id string, params GetCustomAudienceSharedAccountInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shared_account_info"), params.ToParams(), options...)
}

func NewGetCustomAudienceSharedAccountInfoBatchRequest(id string, params GetCustomAudienceSharedAccountInfoParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudiencesharedAccountInfo]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudiencesharedAccountInfo]](GetCustomAudienceSharedAccountInfoBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceSharedAccountInfoBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudiencesharedAccountInfo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudiencesharedAccountInfo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomAudienceSharedAccountInfo(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSharedAccountInfoParams) (*core.Cursor[objects.CustomAudiencesharedAccountInfo], error) {
	var out core.Cursor[objects.CustomAudiencesharedAccountInfo]
	call := GetCustomAudienceSharedAccountInfoBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteCustomAudienceUsersParams struct {
	Namespace *string                 `facebook:"namespace"`
	Payload   *map[string]interface{} `facebook:"payload"`
	Session   *map[string]interface{} `facebook:"session"`
	Extra     core.Params             `facebook:"-"`
}

func (params DeleteCustomAudienceUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Namespace != nil {
		out["namespace"] = *params.Namespace
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Session != nil {
		out["session"] = *params.Session
	}
	return out
}

func DeleteCustomAudienceUsersBatchCall(id string, params DeleteCustomAudienceUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "users"), params.ToParams(), options...)
}

func NewDeleteCustomAudienceUsersBatchRequest(id string, params DeleteCustomAudienceUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCustomAudienceUsersBatchCall(id, params, options...))
}

func DecodeDeleteCustomAudienceUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCustomAudienceUsers(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteCustomAudienceUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateCustomAudienceUsersParams struct {
	Namespace *string                 `facebook:"namespace"`
	Payload   *map[string]interface{} `facebook:"payload"`
	Session   *map[string]interface{} `facebook:"session"`
	Extra     core.Params             `facebook:"-"`
}

func (params CreateCustomAudienceUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Namespace != nil {
		out["namespace"] = *params.Namespace
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Session != nil {
		out["session"] = *params.Session
	}
	return out
}

func CreateCustomAudienceUsersBatchCall(id string, params CreateCustomAudienceUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "users"), params.ToParams(), options...)
}

func NewCreateCustomAudienceUsersBatchRequest(id string, params CreateCustomAudienceUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateCustomAudienceUsersBatchCall(id, params, options...))
}

func DecodeCreateCustomAudienceUsersBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCustomAudienceUsers(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceUsersParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateCustomAudienceUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateCustomAudienceUsersreplaceParams struct {
	Namespace *string                `facebook:"namespace"`
	Payload   map[string]interface{} `facebook:"payload"`
	Session   map[string]interface{} `facebook:"session"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreateCustomAudienceUsersreplaceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Namespace != nil {
		out["namespace"] = *params.Namespace
	}
	out["payload"] = params.Payload
	out["session"] = params.Session
	return out
}

func CreateCustomAudienceUsersreplaceBatchCall(id string, params CreateCustomAudienceUsersreplaceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "usersreplace"), params.ToParams(), options...)
}

func NewCreateCustomAudienceUsersreplaceBatchRequest(id string, params CreateCustomAudienceUsersreplaceParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateCustomAudienceUsersreplaceBatchCall(id, params, options...))
}

func DecodeCreateCustomAudienceUsersreplaceBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCustomAudienceUsersreplace(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceUsersreplaceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateCustomAudienceUsersreplaceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteCustomAudienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteCustomAudienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteCustomAudienceBatchCall(id string, params DeleteCustomAudienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteCustomAudienceBatchRequest(id string, params DeleteCustomAudienceParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCustomAudienceBatchCall(id, params, options...))
}

func DecodeDeleteCustomAudienceBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCustomAudience(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteCustomAudienceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetCustomAudienceParams struct {
	AdAccountID                *core.ID    `facebook:"ad_account_id"`
	SpecialAdCategories        *[]string   `facebook:"special_ad_categories"`
	SpecialAdCategoryCountries *[]string   `facebook:"special_ad_category_countries"`
	TargetCountries            *[]string   `facebook:"target_countries"`
	Extra                      core.Params `facebook:"-"`
}

func (params GetCustomAudienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountID != nil {
		out["ad_account_id"] = *params.AdAccountID
	}
	if params.SpecialAdCategories != nil {
		out["special_ad_categories"] = *params.SpecialAdCategories
	}
	if params.SpecialAdCategoryCountries != nil {
		out["special_ad_category_countries"] = *params.SpecialAdCategoryCountries
	}
	if params.TargetCountries != nil {
		out["target_countries"] = *params.TargetCountries
	}
	return out
}

func GetCustomAudienceBatchCall(id string, params GetCustomAudienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCustomAudienceBatchRequest(id string, params GetCustomAudienceParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](GetCustomAudienceBatchCall(id, params, options...))
}

func DecodeGetCustomAudienceBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCustomAudience(ctx context.Context, client *core.Client, id string, params GetCustomAudienceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := GetCustomAudienceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateCustomAudienceParams struct {
	AllowedDomains      *[]string                               `facebook:"allowed_domains"`
	AudienceLabels      *[]enums.CustomaudienceAudienceLabels   `facebook:"audience_labels"`
	ClaimObjective      *enums.CustomaudienceClaimObjective     `facebook:"claim_objective"`
	ContentType         *enums.CustomaudienceContentType        `facebook:"content_type"`
	Countries           *string                                 `facebook:"countries"`
	CustomerFileSource  *enums.CustomaudienceCustomerFileSource `facebook:"customer_file_source"`
	Description         *string                                 `facebook:"description"`
	EnableFetchOrCreate *bool                                   `facebook:"enable_fetch_or_create"`
	EventSourceGroup    *string                                 `facebook:"event_source_group"`
	EventSources        *[]map[string]interface{}               `facebook:"event_sources"`
	Exclusions          *[]map[string]interface{}               `facebook:"exclusions"`
	Inclusionoperator   *string                                 `facebook:"inclusionOperator"`
	Inclusions          *[]map[string]interface{}               `facebook:"inclusions"`
	LookalikeSpec       *string                                 `facebook:"lookalike_spec"`
	Name                *string                                 `facebook:"name"`
	OptOutLink          *string                                 `facebook:"opt_out_link"`
	ParentAudienceID    *core.ID                                `facebook:"parent_audience_id"`
	ProductSetID        *core.ID                                `facebook:"product_set_id"`
	RetentionDays       *uint64                                 `facebook:"retention_days"`
	RevSharePolicyID    *core.ID                                `facebook:"rev_share_policy_id"`
	Rule                *string                                 `facebook:"rule"`
	RuleAggregation     *string                                 `facebook:"rule_aggregation"`
	Tags                *[]string                               `facebook:"tags"`
	UseForProducts      *[]enums.CustomaudienceUseForProducts   `facebook:"use_for_products"`
	UseInCampaigns      *bool                                   `facebook:"use_in_campaigns"`
	Extra               core.Params                             `facebook:"-"`
}

func (params UpdateCustomAudienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowedDomains != nil {
		out["allowed_domains"] = *params.AllowedDomains
	}
	if params.AudienceLabels != nil {
		out["audience_labels"] = *params.AudienceLabels
	}
	if params.ClaimObjective != nil {
		out["claim_objective"] = *params.ClaimObjective
	}
	if params.ContentType != nil {
		out["content_type"] = *params.ContentType
	}
	if params.Countries != nil {
		out["countries"] = *params.Countries
	}
	if params.CustomerFileSource != nil {
		out["customer_file_source"] = *params.CustomerFileSource
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableFetchOrCreate != nil {
		out["enable_fetch_or_create"] = *params.EnableFetchOrCreate
	}
	if params.EventSourceGroup != nil {
		out["event_source_group"] = *params.EventSourceGroup
	}
	if params.EventSources != nil {
		out["event_sources"] = *params.EventSources
	}
	if params.Exclusions != nil {
		out["exclusions"] = *params.Exclusions
	}
	if params.Inclusionoperator != nil {
		out["inclusionOperator"] = *params.Inclusionoperator
	}
	if params.Inclusions != nil {
		out["inclusions"] = *params.Inclusions
	}
	if params.LookalikeSpec != nil {
		out["lookalike_spec"] = *params.LookalikeSpec
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OptOutLink != nil {
		out["opt_out_link"] = *params.OptOutLink
	}
	if params.ParentAudienceID != nil {
		out["parent_audience_id"] = *params.ParentAudienceID
	}
	if params.ProductSetID != nil {
		out["product_set_id"] = *params.ProductSetID
	}
	if params.RetentionDays != nil {
		out["retention_days"] = *params.RetentionDays
	}
	if params.RevSharePolicyID != nil {
		out["rev_share_policy_id"] = *params.RevSharePolicyID
	}
	if params.Rule != nil {
		out["rule"] = *params.Rule
	}
	if params.RuleAggregation != nil {
		out["rule_aggregation"] = *params.RuleAggregation
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.UseForProducts != nil {
		out["use_for_products"] = *params.UseForProducts
	}
	if params.UseInCampaigns != nil {
		out["use_in_campaigns"] = *params.UseInCampaigns
	}
	return out
}

func UpdateCustomAudienceBatchCall(id string, params UpdateCustomAudienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCustomAudienceBatchRequest(id string, params UpdateCustomAudienceParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](UpdateCustomAudienceBatchCall(id, params, options...))
}

func DecodeUpdateCustomAudienceBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCustomAudience(ctx context.Context, client *core.Client, id string, params UpdateCustomAudienceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := UpdateCustomAudienceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
