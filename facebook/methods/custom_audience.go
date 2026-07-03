package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func DeleteCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceAdaccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
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

func GetCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params GetCustomAudienceAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
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

func CreateCustomAudienceAdaccounts(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceAdaccountsParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
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

func GetCustomAudienceAds(ctx context.Context, client *core.Client, id string, params GetCustomAudienceAdsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), &out)
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

func GetCustomAudienceHealth(ctx context.Context, client *core.Client, id string, params GetCustomAudienceHealthParams) (*core.Cursor[objects.CustomAudienceHealth], error) {
	var out core.Cursor[objects.CustomAudienceHealth]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "health"), params.ToParams(), &out)
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

func GetCustomAudienceSalts(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSaltsParams) (*core.Cursor[objects.CustomAudienceSalts], error) {
	var out core.Cursor[objects.CustomAudienceSalts]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "salts"), params.ToParams(), &out)
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

func CreateCustomAudienceSalts(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceSaltsParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "salts"), params.ToParams(), &out)
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

func GetCustomAudienceSessions(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSessionsParams) (*core.Cursor[objects.CustomAudienceSession], error) {
	var out core.Cursor[objects.CustomAudienceSession]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sessions"), params.ToParams(), &out)
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

func GetCustomAudienceSharedAccountInfo(ctx context.Context, client *core.Client, id string, params GetCustomAudienceSharedAccountInfoParams) (*core.Cursor[objects.CustomAudiencesharedAccountInfo], error) {
	var out core.Cursor[objects.CustomAudiencesharedAccountInfo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_account_info"), params.ToParams(), &out)
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

func DeleteCustomAudienceUsers(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "users"), params.ToParams(), &out)
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

func CreateCustomAudienceUsers(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceUsersParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "users"), params.ToParams(), &out)
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

func CreateCustomAudienceUsersreplace(ctx context.Context, client *core.Client, id string, params CreateCustomAudienceUsersreplaceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "usersreplace"), params.ToParams(), &out)
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

func DeleteCustomAudience(ctx context.Context, client *core.Client, id string, params DeleteCustomAudienceParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetCustomAudience(ctx context.Context, client *core.Client, id string, params GetCustomAudienceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateCustomAudience(ctx context.Context, client *core.Client, id string, params UpdateCustomAudienceParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
