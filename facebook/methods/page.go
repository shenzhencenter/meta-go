package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageAbTestsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageAbTestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageAbTestsBatchCall(id string, params GetPageAbTestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ab_tests"), params.ToParams(), options...)
}

func NewGetPageAbTestsBatchRequest(id string, params GetPageAbTestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePostExperiment]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePostExperiment]](GetPageAbTestsBatchCall(id, params, options...))
}

func DecodeGetPageAbTestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePostExperiment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePostExperiment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageAbTestsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAbTestsParams) (*core.Cursor[objects.PagePostExperiment], *core.Response, error) {
	var out core.Cursor[objects.PagePostExperiment]
	call := GetPageAbTestsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAbTests(ctx context.Context, client *core.Client, id string, params GetPageAbTestsParams) (*core.Cursor[objects.PagePostExperiment], error) {
	out, _, err := GetPageAbTestsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageAbTestsParams struct {
	ControlVideoID               core.ID                                    `facebook:"control_video_id"`
	Description                  string                                     `facebook:"description"`
	Duration                     uint64                                     `facebook:"duration"`
	ExperimentVideoIds           []core.ID                                  `facebook:"experiment_video_ids"`
	Name                         string                                     `facebook:"name"`
	OptimizationGoal             enums.PageabTestsOptimizationGoalEnumParam `facebook:"optimization_goal"`
	ScheduledExperimentTimestamp *uint64                                    `facebook:"scheduled_experiment_timestamp"`
	Extra                        core.Params                                `facebook:"-"`
}

func (params CreatePageAbTestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["control_video_id"] = params.ControlVideoID
	out["description"] = params.Description
	out["duration"] = params.Duration
	out["experiment_video_ids"] = params.ExperimentVideoIds
	out["name"] = params.Name
	out["optimization_goal"] = params.OptimizationGoal
	if params.ScheduledExperimentTimestamp != nil {
		out["scheduled_experiment_timestamp"] = *params.ScheduledExperimentTimestamp
	}
	return out
}

func CreatePageAbTestsBatchCall(id string, params CreatePageAbTestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ab_tests"), params.ToParams(), options...)
}

func NewCreatePageAbTestsBatchRequest(id string, params CreatePageAbTestsParams, options ...core.BatchOption) *core.BatchRequest[objects.PagePostExperiment] {
	return core.NewBatchRequest[objects.PagePostExperiment](CreatePageAbTestsBatchCall(id, params, options...))
}

func DecodeCreatePageAbTestsBatchResponse(response *core.BatchResponse) (*objects.PagePostExperiment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PagePostExperiment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageAbTestsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageAbTestsParams) (*objects.PagePostExperiment, *core.Response, error) {
	var out objects.PagePostExperiment
	call := CreatePageAbTestsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageAbTests(ctx context.Context, client *core.Client, id string, params CreatePageAbTestsParams) (*objects.PagePostExperiment, error) {
	out, _, err := CreatePageAbTestsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageAdsEligibilityParams struct {
	AdsAccountID *core.ID    `facebook:"ads_account_id"`
	Extra        core.Params `facebook:"-"`
}

func (params GetPageAdsEligibilityParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdsAccountID != nil {
		out["ads_account_id"] = *params.AdsAccountID
	}
	return out
}

func GetPageAdsEligibilityBatchCall(id string, params GetPageAdsEligibilityParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_eligibility"), params.ToParams(), options...)
}

func NewGetPageAdsEligibilityBatchRequest(id string, params GetPageAdsEligibilityParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsEligibility]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsEligibility]](GetPageAdsEligibilityBatchCall(id, params, options...))
}

func DecodeGetPageAdsEligibilityBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsEligibility], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsEligibility]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageAdsEligibilityWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAdsEligibilityParams) (*core.Cursor[objects.AdsEligibility], *core.Response, error) {
	var out core.Cursor[objects.AdsEligibility]
	call := GetPageAdsEligibilityBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAdsEligibility(ctx context.Context, client *core.Client, id string, params GetPageAdsEligibilityParams) (*core.Cursor[objects.AdsEligibility], error) {
	out, _, err := GetPageAdsEligibilityWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageAdsPostsParams struct {
	ExcludeDynamicAds   *bool       `facebook:"exclude_dynamic_ads"`
	IncludeInlineCreate *bool       `facebook:"include_inline_create"`
	Since               *core.Time  `facebook:"since"`
	Until               *core.Time  `facebook:"until"`
	Extra               core.Params `facebook:"-"`
}

func (params GetPageAdsPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ExcludeDynamicAds != nil {
		out["exclude_dynamic_ads"] = *params.ExcludeDynamicAds
	}
	if params.IncludeInlineCreate != nil {
		out["include_inline_create"] = *params.IncludeInlineCreate
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPageAdsPostsBatchCall(id string, params GetPageAdsPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_posts"), params.ToParams(), options...)
}

func NewGetPageAdsPostsBatchRequest(id string, params GetPageAdsPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPageAdsPostsBatchCall(id, params, options...))
}

func DecodeGetPageAdsPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageAdsPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAdsPostsParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPageAdsPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAdsPosts(ctx context.Context, client *core.Client, id string, params GetPageAdsPostsParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPageAdsPostsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeletePageAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeletePageAgenciesBatchCall(id string, params DeletePageAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewDeletePageAgenciesBatchRequest(id string, params DeletePageAgenciesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageAgenciesBatchCall(id, params, options...))
}

func DecodeDeletePageAgenciesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageAgenciesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageAgencies(ctx context.Context, client *core.Client, id string, params DeletePageAgenciesParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageAgenciesBatchCall(id string, params GetPageAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetPageAgenciesBatchRequest(id string, params GetPageAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetPageAgenciesBatchCall(id, params, options...))
}

func DecodeGetPageAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetPageAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetPageAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAgencies(ctx context.Context, client *core.Client, id string, params GetPageAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetPageAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageAgenciesParams struct {
	Business       string                                       `facebook:"business"`
	PermittedTasks *[]enums.PageagenciesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                  `facebook:"-"`
}

func (params CreatePageAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	if params.PermittedTasks != nil {
		out["permitted_tasks"] = *params.PermittedTasks
	}
	return out
}

func CreatePageAgenciesBatchCall(id string, params CreatePageAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewCreatePageAgenciesBatchRequest(id string, params CreatePageAgenciesParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageAgenciesBatchCall(id, params, options...))
}

func DecodeCreatePageAgenciesBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageAgenciesParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageAgencies(ctx context.Context, client *core.Client, id string, params CreatePageAgenciesParams) (*objects.Page, error) {
	out, _, err := CreatePageAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageAlbumsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageAlbumsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageAlbumsBatchCall(id string, params GetPageAlbumsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "albums"), params.ToParams(), options...)
}

func NewGetPageAlbumsBatchRequest(id string, params GetPageAlbumsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Album]] {
	return core.NewBatchRequest[core.Cursor[objects.Album]](GetPageAlbumsBatchCall(id, params, options...))
}

func DecodeGetPageAlbumsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Album], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Album]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageAlbumsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAlbumsParams) (*core.Cursor[objects.Album], *core.Response, error) {
	var out core.Cursor[objects.Album]
	call := GetPageAlbumsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAlbums(ctx context.Context, client *core.Client, id string, params GetPageAlbumsParams) (*core.Cursor[objects.Album], error) {
	out, _, err := GetPageAlbumsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageArExperienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageArExperienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageArExperienceBatchCall(id string, params GetPageArExperienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ar_experience"), params.ToParams(), options...)
}

func NewGetPageArExperienceBatchRequest(id string, params GetPageArExperienceParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ArAdsDataContainer]] {
	return core.NewBatchRequest[core.Cursor[objects.ArAdsDataContainer]](GetPageArExperienceBatchCall(id, params, options...))
}

func DecodeGetPageArExperienceBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ArAdsDataContainer], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ArAdsDataContainer]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageArExperienceWithResponse(ctx context.Context, client *core.Client, id string, params GetPageArExperienceParams) (*core.Cursor[objects.ArAdsDataContainer], *core.Response, error) {
	var out core.Cursor[objects.ArAdsDataContainer]
	call := GetPageArExperienceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageArExperience(ctx context.Context, client *core.Client, id string, params GetPageArExperienceParams) (*core.Cursor[objects.ArAdsDataContainer], error) {
	out, _, err := GetPageArExperienceWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeletePageAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeletePageAssignedUsersBatchCall(id string, params DeletePageAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewDeletePageAssignedUsersBatchRequest(id string, params DeletePageAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageAssignedUsersBatchCall(id, params, options...))
}

func DecodeDeletePageAssignedUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageAssignedUsersParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageAssignedUsers(ctx context.Context, client *core.Client, id string, params DeletePageAssignedUsersParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetPageAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetPageAssignedUsersBatchCall(id string, params GetPageAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetPageAssignedUsersBatchRequest(id string, params GetPageAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetPageAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetPageAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
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

func GetPageAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetPageAssignedUsersParams) (*core.Cursor[objects.AssignedUser], *core.Response, error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetPageAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageAssignedUsers(ctx context.Context, client *core.Client, id string, params GetPageAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	out, _, err := GetPageAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageAssignedUsersParams struct {
	Tasks *[]enums.PageassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                      `facebook:"user"`
	Extra core.Params                              `facebook:"-"`
}

func (params CreatePageAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	out["user"] = params.User
	return out
}

func CreatePageAssignedUsersBatchCall(id string, params CreatePageAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreatePageAssignedUsersBatchRequest(id string, params CreatePageAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreatePageAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageAssignedUsersWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageAssignedUsersParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageAssignedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageAssignedUsers(ctx context.Context, client *core.Client, id string, params CreatePageAssignedUsersParams) (*objects.Page, error) {
	out, _, err := CreatePageAssignedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageBlockedParams struct {
	Asid  *string     `facebook:"asid"`
	Psid  *int        `facebook:"psid"`
	UID   *core.ID    `facebook:"uid"`
	User  *int        `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeletePageBlockedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Asid != nil {
		out["asid"] = *params.Asid
	}
	if params.Psid != nil {
		out["psid"] = *params.Psid
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.User != nil {
		out["user"] = *params.User
	}
	return out
}

func DeletePageBlockedBatchCall(id string, params DeletePageBlockedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "blocked"), params.ToParams(), options...)
}

func NewDeletePageBlockedBatchRequest(id string, params DeletePageBlockedParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageBlockedBatchCall(id, params, options...))
}

func DecodeDeletePageBlockedBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageBlockedWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageBlockedParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageBlockedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageBlocked(ctx context.Context, client *core.Client, id string, params DeletePageBlockedParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageBlockedWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageBlockedParams struct {
	UID   *core.ID    `facebook:"uid"`
	User  *int        `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params GetPageBlockedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.User != nil {
		out["user"] = *params.User
	}
	return out
}

func GetPageBlockedBatchCall(id string, params GetPageBlockedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "blocked"), params.ToParams(), options...)
}

func NewGetPageBlockedBatchRequest(id string, params GetPageBlockedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPageBlockedBatchCall(id, params, options...))
}

func DecodeGetPageBlockedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Profile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageBlockedWithResponse(ctx context.Context, client *core.Client, id string, params GetPageBlockedParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetPageBlockedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageBlocked(ctx context.Context, client *core.Client, id string, params GetPageBlockedParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetPageBlockedWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageBlockedParams struct {
	Asid  *[]string   `facebook:"asid"`
	Psid  *[]int      `facebook:"psid"`
	UID   *[]core.ID  `facebook:"uid"`
	User  *[]string   `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params CreatePageBlockedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Asid != nil {
		out["asid"] = *params.Asid
	}
	if params.Psid != nil {
		out["psid"] = *params.Psid
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.User != nil {
		out["user"] = *params.User
	}
	return out
}

func CreatePageBlockedBatchCall(id string, params CreatePageBlockedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "blocked"), params.ToParams(), options...)
}

func NewCreatePageBlockedBatchRequest(id string, params CreatePageBlockedParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageBlockedBatchCall(id, params, options...))
}

func DecodeCreatePageBlockedBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageBlockedWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageBlockedParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageBlockedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageBlocked(ctx context.Context, client *core.Client, id string, params CreatePageBlockedParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageBlockedWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageBusinessMessagingFeatureStatusParams struct {
	Features []map[string]interface{} `facebook:"features"`
	Extra    core.Params              `facebook:"-"`
}

func (params CreatePageBusinessMessagingFeatureStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["features"] = params.Features
	return out
}

func CreatePageBusinessMessagingFeatureStatusBatchCall(id string, params CreatePageBusinessMessagingFeatureStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), options...)
}

func NewCreatePageBusinessMessagingFeatureStatusBatchRequest(id string, params CreatePageBusinessMessagingFeatureStatusParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageBusinessMessagingFeatureStatusBatchCall(id, params, options...))
}

func DecodeCreatePageBusinessMessagingFeatureStatusBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageBusinessMessagingFeatureStatusWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageBusinessMessagingFeatureStatusParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageBusinessMessagingFeatureStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params CreatePageBusinessMessagingFeatureStatusParams) (*objects.Page, error) {
	out, _, err := CreatePageBusinessMessagingFeatureStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageBusinessprojectsParams struct {
	Business *string     `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetPageBusinessprojectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Business != nil {
		out["business"] = *params.Business
	}
	return out
}

func GetPageBusinessprojectsBatchCall(id string, params GetPageBusinessprojectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "businessprojects"), params.ToParams(), options...)
}

func NewGetPageBusinessprojectsBatchRequest(id string, params GetPageBusinessprojectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessProject]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessProject]](GetPageBusinessprojectsBatchCall(id, params, options...))
}

func DecodeGetPageBusinessprojectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessProject], error) {
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

func GetPageBusinessprojectsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageBusinessprojectsParams) (*core.Cursor[objects.BusinessProject], *core.Response, error) {
	var out core.Cursor[objects.BusinessProject]
	call := GetPageBusinessprojectsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageBusinessprojects(ctx context.Context, client *core.Client, id string, params GetPageBusinessprojectsParams) (*core.Cursor[objects.BusinessProject], error) {
	out, _, err := GetPageBusinessprojectsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCallMetricsParams struct {
	AudioStats                   map[string]interface{}                      `facebook:"audio_stats"`
	CallEndedTime                core.Time                                   `facebook:"call_ended_time"`
	CallID                       core.ID                                     `facebook:"call_id"`
	EndCallReason                enums.PagecallMetricsEndCallReasonEnumParam `facebook:"end_call_reason"`
	EndCallSubreason             *string                                     `facebook:"end_call_subreason"`
	FirstAudioPacketReceivedTime core.Time                                   `facebook:"first_audio_packet_received_time"`
	FirstVideoPacketReceivedTime *core.Time                                  `facebook:"first_video_packet_received_time"`
	Platform                     enums.PagecallMetricsPlatformEnumParam      `facebook:"platform"`
	VideoStats                   *map[string]interface{}                     `facebook:"video_stats"`
	Extra                        core.Params                                 `facebook:"-"`
}

func (params CreatePageCallMetricsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["audio_stats"] = params.AudioStats
	out["call_ended_time"] = params.CallEndedTime
	out["call_id"] = params.CallID
	out["end_call_reason"] = params.EndCallReason
	if params.EndCallSubreason != nil {
		out["end_call_subreason"] = *params.EndCallSubreason
	}
	out["first_audio_packet_received_time"] = params.FirstAudioPacketReceivedTime
	if params.FirstVideoPacketReceivedTime != nil {
		out["first_video_packet_received_time"] = *params.FirstVideoPacketReceivedTime
	}
	out["platform"] = params.Platform
	if params.VideoStats != nil {
		out["video_stats"] = *params.VideoStats
	}
	return out
}

func CreatePageCallMetricsBatchCall(id string, params CreatePageCallMetricsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "call_metrics"), params.ToParams(), options...)
}

func NewCreatePageCallMetricsBatchRequest(id string, params CreatePageCallMetricsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageCallMetricsBatchCall(id, params, options...))
}

func DecodeCreatePageCallMetricsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageCallMetricsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCallMetricsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageCallMetricsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCallMetrics(ctx context.Context, client *core.Client, id string, params CreatePageCallMetricsParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageCallMetricsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCallToActionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCallToActionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCallToActionsBatchCall(id string, params GetPageCallToActionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "call_to_actions"), params.ToParams(), options...)
}

func NewGetPageCallToActionsBatchRequest(id string, params GetPageCallToActionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PageCallToAction]] {
	return core.NewBatchRequest[core.Cursor[objects.PageCallToAction]](GetPageCallToActionsBatchCall(id, params, options...))
}

func DecodeGetPageCallToActionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PageCallToAction], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PageCallToAction]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCallToActionsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCallToActionsParams) (*core.Cursor[objects.PageCallToAction], *core.Response, error) {
	var out core.Cursor[objects.PageCallToAction]
	call := GetPageCallToActionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCallToActions(ctx context.Context, client *core.Client, id string, params GetPageCallToActionsParams) (*core.Cursor[objects.PageCallToAction], error) {
	out, _, err := GetPageCallToActionsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCallsParams struct {
	Action      enums.PagecallsActionEnumParam    `facebook:"action"`
	CallID      *core.ID                          `facebook:"call_id"`
	FromVersion *uint64                           `facebook:"from_version"`
	Platform    *enums.PagecallsPlatformEnumParam `facebook:"platform"`
	Session     *map[string]interface{}           `facebook:"session"`
	To          *string                           `facebook:"to"`
	ToVersion   *uint64                           `facebook:"to_version"`
	Tracks      *[]map[string]interface{}         `facebook:"tracks"`
	Extra       core.Params                       `facebook:"-"`
}

func (params CreatePageCallsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["action"] = params.Action
	if params.CallID != nil {
		out["call_id"] = *params.CallID
	}
	if params.FromVersion != nil {
		out["from_version"] = *params.FromVersion
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.Session != nil {
		out["session"] = *params.Session
	}
	if params.To != nil {
		out["to"] = *params.To
	}
	if params.ToVersion != nil {
		out["to_version"] = *params.ToVersion
	}
	if params.Tracks != nil {
		out["tracks"] = *params.Tracks
	}
	return out
}

func CreatePageCallsBatchCall(id string, params CreatePageCallsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "calls"), params.ToParams(), options...)
}

func NewCreatePageCallsBatchRequest(id string, params CreatePageCallsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageCallsBatchCall(id, params, options...))
}

func DecodeCreatePageCallsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageCallsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCallsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageCallsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCalls(ctx context.Context, client *core.Client, id string, params CreatePageCallsParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageCallsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCanvasElementsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCanvasElementsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCanvasElementsBatchCall(id string, params GetPageCanvasElementsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "canvas_elements"), params.ToParams(), options...)
}

func NewGetPageCanvasElementsBatchRequest(id string, params GetPageCanvasElementsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CanvasBodyElement]] {
	return core.NewBatchRequest[core.Cursor[objects.CanvasBodyElement]](GetPageCanvasElementsBatchCall(id, params, options...))
}

func DecodeGetPageCanvasElementsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CanvasBodyElement], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CanvasBodyElement]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCanvasElementsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCanvasElementsParams) (*core.Cursor[objects.CanvasBodyElement], *core.Response, error) {
	var out core.Cursor[objects.CanvasBodyElement]
	call := GetPageCanvasElementsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCanvasElements(ctx context.Context, client *core.Client, id string, params GetPageCanvasElementsParams) (*core.Cursor[objects.CanvasBodyElement], error) {
	out, _, err := GetPageCanvasElementsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCanvasElementsParams struct {
	CanvasButton        *map[string]interface{} `facebook:"canvas_button"`
	CanvasCarousel      *map[string]interface{} `facebook:"canvas_carousel"`
	CanvasExistingPost  *map[string]interface{} `facebook:"canvas_existing_post"`
	CanvasFooter        *map[string]interface{} `facebook:"canvas_footer"`
	CanvasHeader        *map[string]interface{} `facebook:"canvas_header"`
	CanvasLeadForm      *map[string]interface{} `facebook:"canvas_lead_form"`
	CanvasPhoto         *map[string]interface{} `facebook:"canvas_photo"`
	CanvasProductList   *map[string]interface{} `facebook:"canvas_product_list"`
	CanvasProductSet    *map[string]interface{} `facebook:"canvas_product_set"`
	CanvasStoreLocator  *map[string]interface{} `facebook:"canvas_store_locator"`
	CanvasTemplateVideo *map[string]interface{} `facebook:"canvas_template_video"`
	CanvasText          *map[string]interface{} `facebook:"canvas_text"`
	CanvasVideo         *map[string]interface{} `facebook:"canvas_video"`
	Extra               core.Params             `facebook:"-"`
}

func (params CreatePageCanvasElementsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CanvasButton != nil {
		out["canvas_button"] = *params.CanvasButton
	}
	if params.CanvasCarousel != nil {
		out["canvas_carousel"] = *params.CanvasCarousel
	}
	if params.CanvasExistingPost != nil {
		out["canvas_existing_post"] = *params.CanvasExistingPost
	}
	if params.CanvasFooter != nil {
		out["canvas_footer"] = *params.CanvasFooter
	}
	if params.CanvasHeader != nil {
		out["canvas_header"] = *params.CanvasHeader
	}
	if params.CanvasLeadForm != nil {
		out["canvas_lead_form"] = *params.CanvasLeadForm
	}
	if params.CanvasPhoto != nil {
		out["canvas_photo"] = *params.CanvasPhoto
	}
	if params.CanvasProductList != nil {
		out["canvas_product_list"] = *params.CanvasProductList
	}
	if params.CanvasProductSet != nil {
		out["canvas_product_set"] = *params.CanvasProductSet
	}
	if params.CanvasStoreLocator != nil {
		out["canvas_store_locator"] = *params.CanvasStoreLocator
	}
	if params.CanvasTemplateVideo != nil {
		out["canvas_template_video"] = *params.CanvasTemplateVideo
	}
	if params.CanvasText != nil {
		out["canvas_text"] = *params.CanvasText
	}
	if params.CanvasVideo != nil {
		out["canvas_video"] = *params.CanvasVideo
	}
	return out
}

func CreatePageCanvasElementsBatchCall(id string, params CreatePageCanvasElementsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "canvas_elements"), params.ToParams(), options...)
}

func NewCreatePageCanvasElementsBatchRequest(id string, params CreatePageCanvasElementsParams, options ...core.BatchOption) *core.BatchRequest[objects.CanvasBodyElement] {
	return core.NewBatchRequest[objects.CanvasBodyElement](CreatePageCanvasElementsBatchCall(id, params, options...))
}

func DecodeCreatePageCanvasElementsBatchResponse(response *core.BatchResponse) (*objects.CanvasBodyElement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CanvasBodyElement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageCanvasElementsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCanvasElementsParams) (*objects.CanvasBodyElement, *core.Response, error) {
	var out objects.CanvasBodyElement
	call := CreatePageCanvasElementsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCanvasElements(ctx context.Context, client *core.Client, id string, params CreatePageCanvasElementsParams) (*objects.CanvasBodyElement, error) {
	out, _, err := CreatePageCanvasElementsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCanvasesParams struct {
	IsHidden    *bool       `facebook:"is_hidden"`
	IsPublished *bool       `facebook:"is_published"`
	Extra       core.Params `facebook:"-"`
}

func (params GetPageCanvasesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	return out
}

func GetPageCanvasesBatchCall(id string, params GetPageCanvasesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "canvases"), params.ToParams(), options...)
}

func NewGetPageCanvasesBatchRequest(id string, params GetPageCanvasesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Canvas]] {
	return core.NewBatchRequest[core.Cursor[objects.Canvas]](GetPageCanvasesBatchCall(id, params, options...))
}

func DecodeGetPageCanvasesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Canvas], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Canvas]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCanvasesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCanvasesParams) (*core.Cursor[objects.Canvas], *core.Response, error) {
	var out core.Cursor[objects.Canvas]
	call := GetPageCanvasesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCanvases(ctx context.Context, client *core.Client, id string, params GetPageCanvasesParams) (*core.Cursor[objects.Canvas], error) {
	out, _, err := GetPageCanvasesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCanvasesParams struct {
	BackgroundColor           *string     `facebook:"background_color"`
	BodyElementIds            *[]core.ID  `facebook:"body_element_ids"`
	EnableSwipeToOpen         *bool       `facebook:"enable_swipe_to_open"`
	HeroAssetFacebookPostID   *core.ID    `facebook:"hero_asset_facebook_post_id"`
	HeroAssetInstagramMediaID *core.ID    `facebook:"hero_asset_instagram_media_id"`
	IsHidden                  *bool       `facebook:"is_hidden"`
	IsPublished               *bool       `facebook:"is_published"`
	Name                      *string     `facebook:"name"`
	SourceTemplateID          *core.ID    `facebook:"source_template_id"`
	Extra                     core.Params `facebook:"-"`
}

func (params CreatePageCanvasesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BackgroundColor != nil {
		out["background_color"] = *params.BackgroundColor
	}
	if params.BodyElementIds != nil {
		out["body_element_ids"] = *params.BodyElementIds
	}
	if params.EnableSwipeToOpen != nil {
		out["enable_swipe_to_open"] = *params.EnableSwipeToOpen
	}
	if params.HeroAssetFacebookPostID != nil {
		out["hero_asset_facebook_post_id"] = *params.HeroAssetFacebookPostID
	}
	if params.HeroAssetInstagramMediaID != nil {
		out["hero_asset_instagram_media_id"] = *params.HeroAssetInstagramMediaID
	}
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SourceTemplateID != nil {
		out["source_template_id"] = *params.SourceTemplateID
	}
	return out
}

func CreatePageCanvasesBatchCall(id string, params CreatePageCanvasesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "canvases"), params.ToParams(), options...)
}

func NewCreatePageCanvasesBatchRequest(id string, params CreatePageCanvasesParams, options ...core.BatchOption) *core.BatchRequest[objects.Canvas] {
	return core.NewBatchRequest[objects.Canvas](CreatePageCanvasesBatchCall(id, params, options...))
}

func DecodeCreatePageCanvasesBatchResponse(response *core.BatchResponse) (*objects.Canvas, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Canvas
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageCanvasesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCanvasesParams) (*objects.Canvas, *core.Response, error) {
	var out objects.Canvas
	call := CreatePageCanvasesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCanvases(ctx context.Context, client *core.Client, id string, params CreatePageCanvasesParams) (*objects.Canvas, error) {
	out, _, err := CreatePageCanvasesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageChatPluginParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageChatPluginParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageChatPluginBatchCall(id string, params GetPageChatPluginParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "chat_plugin"), params.ToParams(), options...)
}

func NewGetPageChatPluginBatchRequest(id string, params GetPageChatPluginParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ChatPlugin]] {
	return core.NewBatchRequest[core.Cursor[objects.ChatPlugin]](GetPageChatPluginBatchCall(id, params, options...))
}

func DecodeGetPageChatPluginBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ChatPlugin], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ChatPlugin]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageChatPluginWithResponse(ctx context.Context, client *core.Client, id string, params GetPageChatPluginParams) (*core.Cursor[objects.ChatPlugin], *core.Response, error) {
	var out core.Cursor[objects.ChatPlugin]
	call := GetPageChatPluginBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageChatPlugin(ctx context.Context, client *core.Client, id string, params GetPageChatPluginParams) (*core.Cursor[objects.ChatPlugin], error) {
	out, _, err := GetPageChatPluginWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCommerceMerchantSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCommerceMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCommerceMerchantSettingsBatchCall(id string, params GetPageCommerceMerchantSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_merchant_settings"), params.ToParams(), options...)
}

func NewGetPageCommerceMerchantSettingsBatchRequest(id string, params GetPageCommerceMerchantSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceMerchantSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceMerchantSettings]](GetPageCommerceMerchantSettingsBatchCall(id, params, options...))
}

func DecodeGetPageCommerceMerchantSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceMerchantSettings], error) {
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

func GetPageCommerceMerchantSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCommerceMerchantSettingsParams) (*core.Cursor[objects.CommerceMerchantSettings], *core.Response, error) {
	var out core.Cursor[objects.CommerceMerchantSettings]
	call := GetPageCommerceMerchantSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params GetPageCommerceMerchantSettingsParams) (*core.Cursor[objects.CommerceMerchantSettings], error) {
	out, _, err := GetPageCommerceMerchantSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCommerceOrdersParams struct {
	Filters       *[]enums.PagecommerceOrdersFiltersEnumParam `facebook:"filters"`
	State         *[]enums.PagecommerceOrdersStateEnumParam   `facebook:"state"`
	UpdatedAfter  *core.Time                                  `facebook:"updated_after"`
	UpdatedBefore *core.Time                                  `facebook:"updated_before"`
	Extra         core.Params                                 `facebook:"-"`
}

func (params GetPageCommerceOrdersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.State != nil {
		out["state"] = *params.State
	}
	if params.UpdatedAfter != nil {
		out["updated_after"] = *params.UpdatedAfter
	}
	if params.UpdatedBefore != nil {
		out["updated_before"] = *params.UpdatedBefore
	}
	return out
}

func GetPageCommerceOrdersBatchCall(id string, params GetPageCommerceOrdersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_orders"), params.ToParams(), options...)
}

func NewGetPageCommerceOrdersBatchRequest(id string, params GetPageCommerceOrdersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrder]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrder]](GetPageCommerceOrdersBatchCall(id, params, options...))
}

func DecodeGetPageCommerceOrdersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrder], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrder]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCommerceOrdersWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCommerceOrdersParams) (*core.Cursor[objects.CommerceOrder], *core.Response, error) {
	var out core.Cursor[objects.CommerceOrder]
	call := GetPageCommerceOrdersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCommerceOrders(ctx context.Context, client *core.Client, id string, params GetPageCommerceOrdersParams) (*core.Cursor[objects.CommerceOrder], error) {
	out, _, err := GetPageCommerceOrdersWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCommercePayoutsParams struct {
	EndTime   *core.Time  `facebook:"end_time"`
	StartTime *core.Time  `facebook:"start_time"`
	Extra     core.Params `facebook:"-"`
}

func (params GetPageCommercePayoutsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetPageCommercePayoutsBatchCall(id string, params GetPageCommercePayoutsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_payouts"), params.ToParams(), options...)
}

func NewGetPageCommercePayoutsBatchRequest(id string, params GetPageCommercePayoutsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommercePayout]] {
	return core.NewBatchRequest[core.Cursor[objects.CommercePayout]](GetPageCommercePayoutsBatchCall(id, params, options...))
}

func DecodeGetPageCommercePayoutsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommercePayout], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommercePayout]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCommercePayoutsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCommercePayoutsParams) (*core.Cursor[objects.CommercePayout], *core.Response, error) {
	var out core.Cursor[objects.CommercePayout]
	call := GetPageCommercePayoutsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCommercePayouts(ctx context.Context, client *core.Client, id string, params GetPageCommercePayoutsParams) (*core.Cursor[objects.CommercePayout], error) {
	out, _, err := GetPageCommercePayoutsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCommerceTransactionsParams struct {
	EndTime           *core.Time  `facebook:"end_time"`
	PayoutReferenceID *core.ID    `facebook:"payout_reference_id"`
	StartTime         *core.Time  `facebook:"start_time"`
	Extra             core.Params `facebook:"-"`
}

func (params GetPageCommerceTransactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.PayoutReferenceID != nil {
		out["payout_reference_id"] = *params.PayoutReferenceID
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetPageCommerceTransactionsBatchCall(id string, params GetPageCommerceTransactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_transactions"), params.ToParams(), options...)
}

func NewGetPageCommerceTransactionsBatchRequest(id string, params GetPageCommerceTransactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderTransactionDetail]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderTransactionDetail]](GetPageCommerceTransactionsBatchCall(id, params, options...))
}

func DecodeGetPageCommerceTransactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderTransactionDetail], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderTransactionDetail]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCommerceTransactionsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCommerceTransactionsParams) (*core.Cursor[objects.CommerceOrderTransactionDetail], *core.Response, error) {
	var out core.Cursor[objects.CommerceOrderTransactionDetail]
	call := GetPageCommerceTransactionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCommerceTransactions(ctx context.Context, client *core.Client, id string, params GetPageCommerceTransactionsParams) (*core.Cursor[objects.CommerceOrderTransactionDetail], error) {
	out, _, err := GetPageCommerceTransactionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageConversationsParams struct {
	Folder   *string                                   `facebook:"folder"`
	Platform *enums.PageconversationsPlatformEnumParam `facebook:"platform"`
	Tags     *[]string                                 `facebook:"tags"`
	UserID   *core.ID                                  `facebook:"user_id"`
	Extra    core.Params                               `facebook:"-"`
}

func (params GetPageConversationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	return out
}

func GetPageConversationsBatchCall(id string, params GetPageConversationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversations"), params.ToParams(), options...)
}

func NewGetPageConversationsBatchRequest(id string, params GetPageConversationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UnifiedThread]] {
	return core.NewBatchRequest[core.Cursor[objects.UnifiedThread]](GetPageConversationsBatchCall(id, params, options...))
}

func DecodeGetPageConversationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UnifiedThread], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UnifiedThread]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageConversationsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageConversationsParams) (*core.Cursor[objects.UnifiedThread], *core.Response, error) {
	var out core.Cursor[objects.UnifiedThread]
	call := GetPageConversationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageConversations(ctx context.Context, client *core.Client, id string, params GetPageConversationsParams) (*core.Cursor[objects.UnifiedThread], error) {
	out, _, err := GetPageConversationsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCopyrightManualClaimsParams struct {
	Action           *enums.PagecopyrightManualClaimsActionEnumParam          `facebook:"action"`
	ActionReason     *enums.PagecopyrightManualClaimsActionReasonEnumParam    `facebook:"action_reason"`
	Countries        *map[string]interface{}                                  `facebook:"countries"`
	MatchContentType enums.PagecopyrightManualClaimsMatchContentTypeEnumParam `facebook:"match_content_type"`
	MatchedAssetID   core.ID                                                  `facebook:"matched_asset_id"`
	ReferenceAssetID core.ID                                                  `facebook:"reference_asset_id"`
	SelectedSegments *[]map[string]interface{}                                `facebook:"selected_segments"`
	Extra            core.Params                                              `facebook:"-"`
}

func (params CreatePageCopyrightManualClaimsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Action != nil {
		out["action"] = *params.Action
	}
	if params.ActionReason != nil {
		out["action_reason"] = *params.ActionReason
	}
	if params.Countries != nil {
		out["countries"] = *params.Countries
	}
	out["match_content_type"] = params.MatchContentType
	out["matched_asset_id"] = params.MatchedAssetID
	out["reference_asset_id"] = params.ReferenceAssetID
	if params.SelectedSegments != nil {
		out["selected_segments"] = *params.SelectedSegments
	}
	return out
}

func CreatePageCopyrightManualClaimsBatchCall(id string, params CreatePageCopyrightManualClaimsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "copyright_manual_claims"), params.ToParams(), options...)
}

func NewCreatePageCopyrightManualClaimsBatchRequest(id string, params CreatePageCopyrightManualClaimsParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyrightMatch] {
	return core.NewBatchRequest[objects.VideoCopyrightMatch](CreatePageCopyrightManualClaimsBatchCall(id, params, options...))
}

func DecodeCreatePageCopyrightManualClaimsBatchResponse(response *core.BatchResponse) (*objects.VideoCopyrightMatch, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyrightMatch
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageCopyrightManualClaimsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCopyrightManualClaimsParams) (*objects.VideoCopyrightMatch, *core.Response, error) {
	var out objects.VideoCopyrightMatch
	call := CreatePageCopyrightManualClaimsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCopyrightManualClaims(ctx context.Context, client *core.Client, id string, params CreatePageCopyrightManualClaimsParams) (*objects.VideoCopyrightMatch, error) {
	out, _, err := CreatePageCopyrightManualClaimsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCrosspostWhitelistedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCrosspostWhitelistedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCrosspostWhitelistedPagesBatchCall(id string, params GetPageCrosspostWhitelistedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "crosspost_whitelisted_pages"), params.ToParams(), options...)
}

func NewGetPageCrosspostWhitelistedPagesBatchRequest(id string, params GetPageCrosspostWhitelistedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPageCrosspostWhitelistedPagesBatchCall(id, params, options...))
}

func DecodeGetPageCrosspostWhitelistedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPageCrosspostWhitelistedPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCrosspostWhitelistedPagesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPageCrosspostWhitelistedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCrosspostWhitelistedPages(ctx context.Context, client *core.Client, id string, params GetPageCrosspostWhitelistedPagesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPageCrosspostWhitelistedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCtxOptimizationEligibilityParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCtxOptimizationEligibilityParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCtxOptimizationEligibilityBatchCall(id string, params GetPageCtxOptimizationEligibilityParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ctx_optimization_eligibility"), params.ToParams(), options...)
}

func NewGetPageCtxOptimizationEligibilityBatchRequest(id string, params GetPageCtxOptimizationEligibilityParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CTXOptimizationEligibility]] {
	return core.NewBatchRequest[core.Cursor[objects.CTXOptimizationEligibility]](GetPageCtxOptimizationEligibilityBatchCall(id, params, options...))
}

func DecodeGetPageCtxOptimizationEligibilityBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CTXOptimizationEligibility], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CTXOptimizationEligibility]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCtxOptimizationEligibilityWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCtxOptimizationEligibilityParams) (*core.Cursor[objects.CTXOptimizationEligibility], *core.Response, error) {
	var out core.Cursor[objects.CTXOptimizationEligibility]
	call := GetPageCtxOptimizationEligibilityBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCtxOptimizationEligibility(ctx context.Context, client *core.Client, id string, params GetPageCtxOptimizationEligibilityParams) (*core.Cursor[objects.CTXOptimizationEligibility], error) {
	out, _, err := GetPageCtxOptimizationEligibilityWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCustomLabelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageCustomLabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageCustomLabelsBatchCall(id string, params GetPageCustomLabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "custom_labels"), params.ToParams(), options...)
}

func NewGetPageCustomLabelsBatchRequest(id string, params GetPageCustomLabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PageUserMessageThreadLabel]] {
	return core.NewBatchRequest[core.Cursor[objects.PageUserMessageThreadLabel]](GetPageCustomLabelsBatchCall(id, params, options...))
}

func DecodeGetPageCustomLabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PageUserMessageThreadLabel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PageUserMessageThreadLabel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCustomLabelsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCustomLabelsParams) (*core.Cursor[objects.PageUserMessageThreadLabel], *core.Response, error) {
	var out core.Cursor[objects.PageUserMessageThreadLabel]
	call := GetPageCustomLabelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCustomLabels(ctx context.Context, client *core.Client, id string, params GetPageCustomLabelsParams) (*core.Cursor[objects.PageUserMessageThreadLabel], error) {
	out, _, err := GetPageCustomLabelsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCustomLabelsParams struct {
	Name          *string     `facebook:"name"`
	PageLabelName string      `facebook:"page_label_name"`
	Extra         core.Params `facebook:"-"`
}

func (params CreatePageCustomLabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	out["page_label_name"] = params.PageLabelName
	return out
}

func CreatePageCustomLabelsBatchCall(id string, params CreatePageCustomLabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "custom_labels"), params.ToParams(), options...)
}

func NewCreatePageCustomLabelsBatchRequest(id string, params CreatePageCustomLabelsParams, options ...core.BatchOption) *core.BatchRequest[objects.PageUserMessageThreadLabel] {
	return core.NewBatchRequest[objects.PageUserMessageThreadLabel](CreatePageCustomLabelsBatchCall(id, params, options...))
}

func DecodeCreatePageCustomLabelsBatchResponse(response *core.BatchResponse) (*objects.PageUserMessageThreadLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageUserMessageThreadLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageCustomLabelsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCustomLabelsParams) (*objects.PageUserMessageThreadLabel, *core.Response, error) {
	var out objects.PageUserMessageThreadLabel
	call := CreatePageCustomLabelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCustomLabels(ctx context.Context, client *core.Client, id string, params CreatePageCustomLabelsParams) (*objects.PageUserMessageThreadLabel, error) {
	out, _, err := CreatePageCustomLabelsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageCustomUserSettingsParams struct {
	Params []enums.PagecustomUserSettingsParamsEnumParam `facebook:"params"`
	Psid   string                                        `facebook:"psid"`
	Extra  core.Params                                   `facebook:"-"`
}

func (params DeletePageCustomUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["params"] = params.Params
	out["psid"] = params.Psid
	return out
}

func DeletePageCustomUserSettingsBatchCall(id string, params DeletePageCustomUserSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "custom_user_settings"), params.ToParams(), options...)
}

func NewDeletePageCustomUserSettingsBatchRequest(id string, params DeletePageCustomUserSettingsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageCustomUserSettingsBatchCall(id, params, options...))
}

func DecodeDeletePageCustomUserSettingsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageCustomUserSettingsWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageCustomUserSettingsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageCustomUserSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageCustomUserSettings(ctx context.Context, client *core.Client, id string, params DeletePageCustomUserSettingsParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageCustomUserSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageCustomUserSettingsParams struct {
	Psid  string      `facebook:"psid"`
	Extra core.Params `facebook:"-"`
}

func (params GetPageCustomUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["psid"] = params.Psid
	return out
}

func GetPageCustomUserSettingsBatchCall(id string, params GetPageCustomUserSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "custom_user_settings"), params.ToParams(), options...)
}

func NewGetPageCustomUserSettingsBatchRequest(id string, params GetPageCustomUserSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomUserSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomUserSettings]](GetPageCustomUserSettingsBatchCall(id, params, options...))
}

func DecodeGetPageCustomUserSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomUserSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomUserSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageCustomUserSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageCustomUserSettingsParams) (*core.Cursor[objects.CustomUserSettings], *core.Response, error) {
	var out core.Cursor[objects.CustomUserSettings]
	call := GetPageCustomUserSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageCustomUserSettings(ctx context.Context, client *core.Client, id string, params GetPageCustomUserSettingsParams) (*core.Cursor[objects.CustomUserSettings], error) {
	out, _, err := GetPageCustomUserSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageCustomUserSettingsParams struct {
	PersistentMenu *[]map[string]interface{} `facebook:"persistent_menu"`
	Psid           string                    `facebook:"psid"`
	Extra          core.Params               `facebook:"-"`
}

func (params CreatePageCustomUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PersistentMenu != nil {
		out["persistent_menu"] = *params.PersistentMenu
	}
	out["psid"] = params.Psid
	return out
}

func CreatePageCustomUserSettingsBatchCall(id string, params CreatePageCustomUserSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "custom_user_settings"), params.ToParams(), options...)
}

func NewCreatePageCustomUserSettingsBatchRequest(id string, params CreatePageCustomUserSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageCustomUserSettingsBatchCall(id, params, options...))
}

func DecodeCreatePageCustomUserSettingsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageCustomUserSettingsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageCustomUserSettingsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageCustomUserSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageCustomUserSettings(ctx context.Context, client *core.Client, id string, params CreatePageCustomUserSettingsParams) (*objects.Page, error) {
	out, _, err := CreatePageCustomUserSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageDatasetBatchCall(id string, params GetPageDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewGetPageDatasetBatchRequest(id string, params GetPageDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Dataset]] {
	return core.NewBatchRequest[core.Cursor[objects.Dataset]](GetPageDatasetBatchCall(id, params, options...))
}

func DecodeGetPageDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Dataset], error) {
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

func GetPageDatasetWithResponse(ctx context.Context, client *core.Client, id string, params GetPageDatasetParams) (*core.Cursor[objects.Dataset], *core.Response, error) {
	var out core.Cursor[objects.Dataset]
	call := GetPageDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageDataset(ctx context.Context, client *core.Client, id string, params GetPageDatasetParams) (*core.Cursor[objects.Dataset], error) {
	out, _, err := GetPageDatasetWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageDatasetParams struct {
	DatasetName *string     `facebook:"dataset_name"`
	Extra       core.Params `facebook:"-"`
}

func (params CreatePageDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatasetName != nil {
		out["dataset_name"] = *params.DatasetName
	}
	return out
}

func CreatePageDatasetBatchCall(id string, params CreatePageDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewCreatePageDatasetBatchRequest(id string, params CreatePageDatasetParams, options ...core.BatchOption) *core.BatchRequest[objects.Dataset] {
	return core.NewBatchRequest[objects.Dataset](CreatePageDatasetBatchCall(id, params, options...))
}

func DecodeCreatePageDatasetBatchResponse(response *core.BatchResponse) (*objects.Dataset, error) {
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

func CreatePageDatasetWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageDatasetParams) (*objects.Dataset, *core.Response, error) {
	var out objects.Dataset
	call := CreatePageDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageDataset(ctx context.Context, client *core.Client, id string, params CreatePageDatasetParams) (*objects.Dataset, error) {
	out, _, err := CreatePageDatasetWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageEventsParams struct {
	EventStateFilter *[]enums.PageeventsEventStateFilterEnumParam `facebook:"event_state_filter"`
	IncludeCanceled  *bool                                        `facebook:"include_canceled"`
	TimeFilter       *enums.PageeventsTimeFilterEnumParam         `facebook:"time_filter"`
	Type             *enums.PageeventsTypeEnumParam               `facebook:"type"`
	Extra            core.Params                                  `facebook:"-"`
}

func (params GetPageEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EventStateFilter != nil {
		out["event_state_filter"] = *params.EventStateFilter
	}
	if params.IncludeCanceled != nil {
		out["include_canceled"] = *params.IncludeCanceled
	}
	if params.TimeFilter != nil {
		out["time_filter"] = *params.TimeFilter
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetPageEventsBatchCall(id string, params GetPageEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "events"), params.ToParams(), options...)
}

func NewGetPageEventsBatchRequest(id string, params GetPageEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Event]] {
	return core.NewBatchRequest[core.Cursor[objects.Event]](GetPageEventsBatchCall(id, params, options...))
}

func DecodeGetPageEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Event], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Event]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageEventsParams) (*core.Cursor[objects.Event], *core.Response, error) {
	var out core.Cursor[objects.Event]
	call := GetPageEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageEvents(ctx context.Context, client *core.Client, id string, params GetPageEventsParams) (*core.Cursor[objects.Event], error) {
	out, _, err := GetPageEventsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageExtendThreadControlParams struct {
	Duration  *uint64                `facebook:"duration"`
	Recipient map[string]interface{} `facebook:"recipient"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreatePageExtendThreadControlParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Duration != nil {
		out["duration"] = *params.Duration
	}
	out["recipient"] = params.Recipient
	return out
}

func CreatePageExtendThreadControlBatchCall(id string, params CreatePageExtendThreadControlParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "extend_thread_control"), params.ToParams(), options...)
}

func NewCreatePageExtendThreadControlBatchRequest(id string, params CreatePageExtendThreadControlParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageExtendThreadControlBatchCall(id, params, options...))
}

func DecodeCreatePageExtendThreadControlBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageExtendThreadControlWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageExtendThreadControlParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageExtendThreadControlBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageExtendThreadControl(ctx context.Context, client *core.Client, id string, params CreatePageExtendThreadControlParams) (*objects.Page, error) {
	out, _, err := CreatePageExtendThreadControlWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageFantasyGamesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageFantasyGamesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageFantasyGamesBatchCall(id string, params GetPageFantasyGamesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "fantasy_games"), params.ToParams(), options...)
}

func NewGetPageFantasyGamesBatchRequest(id string, params GetPageFantasyGamesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.FantasyGame]] {
	return core.NewBatchRequest[core.Cursor[objects.FantasyGame]](GetPageFantasyGamesBatchCall(id, params, options...))
}

func DecodeGetPageFantasyGamesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.FantasyGame], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.FantasyGame]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageFantasyGamesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageFantasyGamesParams) (*core.Cursor[objects.FantasyGame], *core.Response, error) {
	var out core.Cursor[objects.FantasyGame]
	call := GetPageFantasyGamesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageFantasyGames(ctx context.Context, client *core.Client, id string, params GetPageFantasyGamesParams) (*core.Cursor[objects.FantasyGame], error) {
	out, _, err := GetPageFantasyGamesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageFeedParams struct {
	IncludeHidden *bool                        `facebook:"include_hidden"`
	Limit         *uint64                      `facebook:"limit"`
	ShowExpired   *bool                        `facebook:"show_expired"`
	With          *enums.PagefeedWithEnumParam `facebook:"with"`
	Extra         core.Params                  `facebook:"-"`
}

func (params GetPageFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetPageFeedBatchCall(id string, params GetPageFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewGetPageFeedBatchRequest(id string, params GetPageFeedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPageFeedBatchCall(id, params, options...))
}

func DecodeGetPageFeedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageFeedWithResponse(ctx context.Context, client *core.Client, id string, params GetPageFeedParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPageFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageFeed(ctx context.Context, client *core.Client, id string, params GetPageFeedParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPageFeedWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageFeedParams struct {
	Actions                   *map[string]interface{}                          `facebook:"actions"`
	AlbumID                   *core.ID                                         `facebook:"album_id"`
	AndroidKeyHash            *string                                          `facebook:"android_key_hash"`
	ApplicationID             *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID      *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	Asset3dID                 *core.ID                                         `facebook:"asset3d_id"`
	AssociatedID              *core.ID                                         `facebook:"associated_id"`
	AttachPlaceSuggestion     *bool                                            `facebook:"attach_place_suggestion"`
	AttachedMedia             *[]map[string]interface{}                        `facebook:"attached_media"`
	AudienceExp               *bool                                            `facebook:"audience_exp"`
	BackdatedTime             *core.Time                                       `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.PagefeedBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	BreakingNews              *bool                                            `facebook:"breaking_news"`
	BreakingNewsExpiration    *uint64                                          `facebook:"breaking_news_expiration"`
	CallToAction              *map[string]interface{}                          `facebook:"call_to_action"`
	Caption                   *string                                          `facebook:"caption"`
	ChildAttachments          *[]map[string]interface{}                        `facebook:"child_attachments"`
	ClientMutationID          *core.ID                                         `facebook:"client_mutation_id"`
	ComposerEntryPicker       *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint        *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime         *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog  *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID         *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface     *string                                          `facebook:"composer_source_surface"`
	ComposerType              *string                                          `facebook:"composer_type"`
	ConnectionClass           *string                                          `facebook:"connection_class"`
	ContentAttachment         *string                                          `facebook:"content_attachment"`
	Coordinates               *map[string]interface{}                          `facebook:"coordinates"`
	CtaLink                   *string                                          `facebook:"cta_link"`
	CtaType                   *string                                          `facebook:"cta_type"`
	Description               *string                                          `facebook:"description"`
	DirectShareStatus         *uint64                                          `facebook:"direct_share_status"`
	EnforceLinkOwnership      *bool                                            `facebook:"enforce_link_ownership"`
	ExpandedHeight            *uint64                                          `facebook:"expanded_height"`
	ExpandedWidth             *uint64                                          `facebook:"expanded_width"`
	FeedTargeting             *map[string]interface{}                          `facebook:"feed_targeting"`
	Formatting                *enums.PagefeedFormattingEnumParam               `facebook:"formatting"`
	FunFactPromptID           *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID          *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Height                    *uint64                                          `facebook:"height"`
	HomeCheckinCityID         *map[string]interface{}                          `facebook:"home_checkin_city_id"`
	ImageCrops                *map[string]interface{}                          `facebook:"image_crops"`
	ImplicitWithTags          *[]int                                           `facebook:"implicit_with_tags"`
	InstantGameEntryPointData *string                                          `facebook:"instant_game_entry_point_data"`
	IosBundleID               *core.ID                                         `facebook:"ios_bundle_id"`
	IsBackoutDraft            *bool                                            `facebook:"is_backout_draft"`
	IsBoostIntended           *bool                                            `facebook:"is_boost_intended"`
	IsExplicitLocation        *bool                                            `facebook:"is_explicit_location"`
	IsExplicitShare           *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost        *bool                                            `facebook:"is_group_linking_post"`
	IsPhotoContainer          *bool                                            `facebook:"is_photo_container"`
	Link                      *string                                          `facebook:"link"`
	LocationSourceID          *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy             *bool                                            `facebook:"manual_privacy"`
	Message                   *string                                          `facebook:"message"`
	MultiShareEndCard         *bool                                            `facebook:"multi_share_end_card"`
	MultiShareOptimized       *bool                                            `facebook:"multi_share_optimized"`
	Name                      *string                                          `facebook:"name"`
	NectarModule              *string                                          `facebook:"nectar_module"`
	ObjectAttachment          *string                                          `facebook:"object_attachment"`
	OgActionTypeID            *core.ID                                         `facebook:"og_action_type_id"`
	OgHideObjectAttachment    *bool                                            `facebook:"og_hide_object_attachment"`
	OgIconID                  *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                  *string                                          `facebook:"og_phrase"`
	OgSetProfileBadge         *bool                                            `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism     *string                                          `facebook:"og_suggestion_mechanism"`
	PageRecommendation        *string                                          `facebook:"page_recommendation"`
	Picture                   *string                                          `facebook:"picture"`
	Place                     *map[string]interface{}                          `facebook:"place"`
	PlaceAttachmentSetting    *enums.PagefeedPlaceAttachmentSettingEnumParam   `facebook:"place_attachment_setting"`
	PlaceList                 *string                                          `facebook:"place_list"`
	PlaceListData             *[]interface{}                                   `facebook:"place_list_data"`
	PostSurfacesBlacklist     *[]enums.PagefeedPostSurfacesBlacklistEnumParam  `facebook:"post_surfaces_blacklist"`
	PostingToRedspace         *enums.PagefeedPostingToRedspaceEnumParam        `facebook:"posting_to_redspace"`
	Privacy                   *string                                          `facebook:"privacy"`
	PromptID                  *core.ID                                         `facebook:"prompt_id"`
	PromptTrackingString      *string                                          `facebook:"prompt_tracking_string"`
	Properties                *map[string]interface{}                          `facebook:"properties"`
	ProxiedAppID              *core.ID                                         `facebook:"proxied_app_id"`
	PublishEventID            *core.ID                                         `facebook:"publish_event_id"`
	Published                 *bool                                            `facebook:"published"`
	Quote                     *string                                          `facebook:"quote"`
	Ref                       *[]string                                        `facebook:"ref"`
	ReferenceableImageIds     *[]core.ID                                       `facebook:"referenceable_image_ids"`
	ReferralID                *core.ID                                         `facebook:"referral_id"`
	ScheduledPublishTime      *core.Time                                       `facebook:"scheduled_publish_time"`
	Source                    *string                                          `facebook:"source"`
	SponsorID                 *core.ID                                         `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                                          `facebook:"sponsor_relationship"`
	SuggestedPlaceID          *map[string]interface{}                          `facebook:"suggested_place_id"`
	Tags                      *[]int                                           `facebook:"tags"`
	TargetSurface             *enums.PagefeedTargetSurfaceEnumParam            `facebook:"target_surface"`
	Targeting                 *map[string]interface{}                          `facebook:"targeting"`
	TextFormatMetadata        *string                                          `facebook:"text_format_metadata"`
	TextFormatPresetID        *core.ID                                         `facebook:"text_format_preset_id"`
	TextOnlyPlace             *string                                          `facebook:"text_only_place"`
	Thumbnail                 *core.FileParam                                  `facebook:"thumbnail"`
	TimeSinceOriginalPost     *uint64                                          `facebook:"time_since_original_post"`
	Title                     *string                                          `facebook:"title"`
	TrackingInfo              *string                                          `facebook:"tracking_info"`
	UnpublishedContentType    *enums.PagefeedUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	UserSelectedTags          *bool                                            `facebook:"user_selected_tags"`
	VideoStartTimeMs          *uint64                                          `facebook:"video_start_time_ms"`
	ViewerCoordinates         *map[string]interface{}                          `facebook:"viewer_coordinates"`
	Width                     *uint64                                          `facebook:"width"`
	Extra                     core.Params                                      `facebook:"-"`
}

func (params CreatePageFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Actions != nil {
		out["actions"] = *params.Actions
	}
	if params.AlbumID != nil {
		out["album_id"] = *params.AlbumID
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.Asset3dID != nil {
		out["asset3d_id"] = *params.Asset3dID
	}
	if params.AssociatedID != nil {
		out["associated_id"] = *params.AssociatedID
	}
	if params.AttachPlaceSuggestion != nil {
		out["attach_place_suggestion"] = *params.AttachPlaceSuggestion
	}
	if params.AttachedMedia != nil {
		out["attached_media"] = *params.AttachedMedia
	}
	if params.AudienceExp != nil {
		out["audience_exp"] = *params.AudienceExp
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.BreakingNews != nil {
		out["breaking_news"] = *params.BreakingNews
	}
	if params.BreakingNewsExpiration != nil {
		out["breaking_news_expiration"] = *params.BreakingNewsExpiration
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ChildAttachments != nil {
		out["child_attachments"] = *params.ChildAttachments
	}
	if params.ClientMutationID != nil {
		out["client_mutation_id"] = *params.ClientMutationID
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
	if params.ConnectionClass != nil {
		out["connection_class"] = *params.ConnectionClass
	}
	if params.ContentAttachment != nil {
		out["content_attachment"] = *params.ContentAttachment
	}
	if params.Coordinates != nil {
		out["coordinates"] = *params.Coordinates
	}
	if params.CtaLink != nil {
		out["cta_link"] = *params.CtaLink
	}
	if params.CtaType != nil {
		out["cta_type"] = *params.CtaType
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.EnforceLinkOwnership != nil {
		out["enforce_link_ownership"] = *params.EnforceLinkOwnership
	}
	if params.ExpandedHeight != nil {
		out["expanded_height"] = *params.ExpandedHeight
	}
	if params.ExpandedWidth != nil {
		out["expanded_width"] = *params.ExpandedWidth
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.HomeCheckinCityID != nil {
		out["home_checkin_city_id"] = *params.HomeCheckinCityID
	}
	if params.ImageCrops != nil {
		out["image_crops"] = *params.ImageCrops
	}
	if params.ImplicitWithTags != nil {
		out["implicit_with_tags"] = *params.ImplicitWithTags
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsBackoutDraft != nil {
		out["is_backout_draft"] = *params.IsBackoutDraft
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPhotoContainer != nil {
		out["is_photo_container"] = *params.IsPhotoContainer
	}
	if params.Link != nil {
		out["link"] = *params.Link
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.MultiShareEndCard != nil {
		out["multi_share_end_card"] = *params.MultiShareEndCard
	}
	if params.MultiShareOptimized != nil {
		out["multi_share_optimized"] = *params.MultiShareOptimized
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectAttachment != nil {
		out["object_attachment"] = *params.ObjectAttachment
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgHideObjectAttachment != nil {
		out["og_hide_object_attachment"] = *params.OgHideObjectAttachment
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
	if params.OgSetProfileBadge != nil {
		out["og_set_profile_badge"] = *params.OgSetProfileBadge
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.PageRecommendation != nil {
		out["page_recommendation"] = *params.PageRecommendation
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.PlaceAttachmentSetting != nil {
		out["place_attachment_setting"] = *params.PlaceAttachmentSetting
	}
	if params.PlaceList != nil {
		out["place_list"] = *params.PlaceList
	}
	if params.PlaceListData != nil {
		out["place_list_data"] = *params.PlaceListData
	}
	if params.PostSurfacesBlacklist != nil {
		out["post_surfaces_blacklist"] = *params.PostSurfacesBlacklist
	}
	if params.PostingToRedspace != nil {
		out["posting_to_redspace"] = *params.PostingToRedspace
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PromptID != nil {
		out["prompt_id"] = *params.PromptID
	}
	if params.PromptTrackingString != nil {
		out["prompt_tracking_string"] = *params.PromptTrackingString
	}
	if params.Properties != nil {
		out["properties"] = *params.Properties
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Quote != nil {
		out["quote"] = *params.Quote
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	if params.ReferenceableImageIds != nil {
		out["referenceable_image_ids"] = *params.ReferenceableImageIds
	}
	if params.ReferralID != nil {
		out["referral_id"] = *params.ReferralID
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.SuggestedPlaceID != nil {
		out["suggested_place_id"] = *params.SuggestedPlaceID
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.TargetSurface != nil {
		out["target_surface"] = *params.TargetSurface
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.TextFormatPresetID != nil {
		out["text_format_preset_id"] = *params.TextFormatPresetID
	}
	if params.TextOnlyPlace != nil {
		out["text_only_place"] = *params.TextOnlyPlace
	}
	if params.Thumbnail != nil {
		out["thumbnail"] = *params.Thumbnail
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TrackingInfo != nil {
		out["tracking_info"] = *params.TrackingInfo
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.ViewerCoordinates != nil {
		out["viewer_coordinates"] = *params.ViewerCoordinates
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func CreatePageFeedBatchCall(id string, params CreatePageFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewCreatePageFeedBatchRequest(id string, params CreatePageFeedParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageFeedBatchCall(id, params, options...))
}

func DecodeCreatePageFeedBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageFeedWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageFeedParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageFeed(ctx context.Context, client *core.Client, id string, params CreatePageFeedParams) (*objects.Page, error) {
	out, _, err := CreatePageFeedWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageGlobalBrandChildrenParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageGlobalBrandChildrenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageGlobalBrandChildrenBatchCall(id string, params GetPageGlobalBrandChildrenParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "global_brand_children"), params.ToParams(), options...)
}

func NewGetPageGlobalBrandChildrenBatchRequest(id string, params GetPageGlobalBrandChildrenParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPageGlobalBrandChildrenBatchCall(id, params, options...))
}

func DecodeGetPageGlobalBrandChildrenBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPageGlobalBrandChildrenWithResponse(ctx context.Context, client *core.Client, id string, params GetPageGlobalBrandChildrenParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPageGlobalBrandChildrenBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageGlobalBrandChildren(ctx context.Context, client *core.Client, id string, params GetPageGlobalBrandChildrenParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPageGlobalBrandChildrenWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageImageCopyrightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageImageCopyrightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageImageCopyrightsBatchCall(id string, params GetPageImageCopyrightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "image_copyrights"), params.ToParams(), options...)
}

func NewGetPageImageCopyrightsBatchRequest(id string, params GetPageImageCopyrightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ImageCopyright]] {
	return core.NewBatchRequest[core.Cursor[objects.ImageCopyright]](GetPageImageCopyrightsBatchCall(id, params, options...))
}

func DecodeGetPageImageCopyrightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ImageCopyright], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ImageCopyright]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageImageCopyrightsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageImageCopyrightsParams) (*core.Cursor[objects.ImageCopyright], *core.Response, error) {
	var out core.Cursor[objects.ImageCopyright]
	call := GetPageImageCopyrightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageImageCopyrights(ctx context.Context, client *core.Client, id string, params GetPageImageCopyrightsParams) (*core.Cursor[objects.ImageCopyright], error) {
	out, _, err := GetPageImageCopyrightsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageImageCopyrightsParams struct {
	Artist                      *string                                          `facebook:"artist"`
	AttributionLink             *string                                          `facebook:"attribution_link"`
	Creator                     *string                                          `facebook:"creator"`
	CustomID                    *core.ID                                         `facebook:"custom_id"`
	Description                 *string                                          `facebook:"description"`
	Filename                    string                                           `facebook:"filename"`
	GeoOwnership                []enums.PageimageCopyrightsGeoOwnershipEnumParam `facebook:"geo_ownership"`
	OriginalContentCreationDate *uint64                                          `facebook:"original_content_creation_date"`
	ReferencePhoto              string                                           `facebook:"reference_photo"`
	Title                       *string                                          `facebook:"title"`
	Extra                       core.Params                                      `facebook:"-"`
}

func (params CreatePageImageCopyrightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Artist != nil {
		out["artist"] = *params.Artist
	}
	if params.AttributionLink != nil {
		out["attribution_link"] = *params.AttributionLink
	}
	if params.Creator != nil {
		out["creator"] = *params.Creator
	}
	if params.CustomID != nil {
		out["custom_id"] = *params.CustomID
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	out["filename"] = params.Filename
	out["geo_ownership"] = params.GeoOwnership
	if params.OriginalContentCreationDate != nil {
		out["original_content_creation_date"] = *params.OriginalContentCreationDate
	}
	out["reference_photo"] = params.ReferencePhoto
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func CreatePageImageCopyrightsBatchCall(id string, params CreatePageImageCopyrightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "image_copyrights"), params.ToParams(), options...)
}

func NewCreatePageImageCopyrightsBatchRequest(id string, params CreatePageImageCopyrightsParams, options ...core.BatchOption) *core.BatchRequest[objects.ImageCopyright] {
	return core.NewBatchRequest[objects.ImageCopyright](CreatePageImageCopyrightsBatchCall(id, params, options...))
}

func DecodeCreatePageImageCopyrightsBatchResponse(response *core.BatchResponse) (*objects.ImageCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ImageCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageImageCopyrightsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageImageCopyrightsParams) (*objects.ImageCopyright, *core.Response, error) {
	var out objects.ImageCopyright
	call := CreatePageImageCopyrightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageImageCopyrights(ctx context.Context, client *core.Client, id string, params CreatePageImageCopyrightsParams) (*objects.ImageCopyright, error) {
	out, _, err := CreatePageImageCopyrightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageIndexedVideosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageIndexedVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageIndexedVideosBatchCall(id string, params GetPageIndexedVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "indexed_videos"), params.ToParams(), options...)
}

func NewGetPageIndexedVideosBatchRequest(id string, params GetPageIndexedVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetPageIndexedVideosBatchCall(id, params, options...))
}

func DecodeGetPageIndexedVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageIndexedVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetPageIndexedVideosParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetPageIndexedVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageIndexedVideos(ctx context.Context, client *core.Client, id string, params GetPageIndexedVideosParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetPageIndexedVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageInsightsParams struct {
	Breakdown                 *[]map[string]interface{}              `facebook:"breakdown"`
	DatePreset                *enums.PageinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric                    *[]map[string]interface{}              `facebook:"metric"`
	Period                    *enums.PageinsightsPeriodEnumParam     `facebook:"period"`
	ShowDescriptionFromAPIDoc *bool                                  `facebook:"show_description_from_api_doc"`
	Since                     *core.Time                             `facebook:"since"`
	Until                     *core.Time                             `facebook:"until"`
	Extra                     core.Params                            `facebook:"-"`
}

func (params GetPageInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	if params.Period != nil {
		out["period"] = *params.Period
	}
	if params.ShowDescriptionFromAPIDoc != nil {
		out["show_description_from_api_doc"] = *params.ShowDescriptionFromAPIDoc
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPageInsightsBatchCall(id string, params GetPageInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetPageInsightsBatchRequest(id string, params GetPageInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetPageInsightsBatchCall(id, params, options...))
}

func DecodeGetPageInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetPageInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageInsights(ctx context.Context, client *core.Client, id string, params GetPageInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetPageInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageInstagramAccountsBatchCall(id string, params GetPageInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_accounts"), params.ToParams(), options...)
}

func NewGetPageInstagramAccountsBatchRequest(id string, params GetPageInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetPageInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetPageInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
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

func GetPageInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageInstagramAccountsParams) (*core.Cursor[objects.IGUser], *core.Response, error) {
	var out core.Cursor[objects.IGUser]
	call := GetPageInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetPageInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	out, _, err := GetPageInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageLeadgenFormsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageLeadgenFormsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageLeadgenFormsBatchCall(id string, params GetPageLeadgenFormsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "leadgen_forms"), params.ToParams(), options...)
}

func NewGetPageLeadgenFormsBatchRequest(id string, params GetPageLeadgenFormsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LeadgenForm]] {
	return core.NewBatchRequest[core.Cursor[objects.LeadgenForm]](GetPageLeadgenFormsBatchCall(id, params, options...))
}

func DecodeGetPageLeadgenFormsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LeadgenForm], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.LeadgenForm]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageLeadgenFormsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageLeadgenFormsParams) (*core.Cursor[objects.LeadgenForm], *core.Response, error) {
	var out core.Cursor[objects.LeadgenForm]
	call := GetPageLeadgenFormsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageLeadgenForms(ctx context.Context, client *core.Client, id string, params GetPageLeadgenFormsParams) (*core.Cursor[objects.LeadgenForm], error) {
	out, _, err := GetPageLeadgenFormsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageLeadgenFormsParams struct {
	AllowOrganicLeadRetrieval        *bool                                  `facebook:"allow_organic_lead_retrieval"`
	BlockDisplayForNonTargetedViewer *bool                                  `facebook:"block_display_for_non_targeted_viewer"`
	ContextCard                      *map[string]interface{}                `facebook:"context_card"`
	CoverPhoto                       *core.FileParam                        `facebook:"cover_photo"`
	CustomDisclaimer                 *map[string]interface{}                `facebook:"custom_disclaimer"`
	FollowUpActionURL                *string                                `facebook:"follow_up_action_url"`
	IsForCanvas                      *bool                                  `facebook:"is_for_canvas"`
	IsOptimizedForQuality            *bool                                  `facebook:"is_optimized_for_quality"`
	IsPhoneSmsVerifyEnabled          *bool                                  `facebook:"is_phone_sms_verify_enabled"`
	Locale                           *enums.PageleadgenFormsLocaleEnumParam `facebook:"locale"`
	Name                             string                                 `facebook:"name"`
	PrivacyPolicy                    *map[string]interface{}                `facebook:"privacy_policy"`
	QuestionPageCustomHeadline       *string                                `facebook:"question_page_custom_headline"`
	Questions                        []map[string]interface{}               `facebook:"questions"`
	ShouldEnforceWorkEmail           *bool                                  `facebook:"should_enforce_work_email"`
	ThankYouPage                     *map[string]interface{}                `facebook:"thank_you_page"`
	TrackingParameters               *map[string]interface{}                `facebook:"tracking_parameters"`
	UploadGatedFile                  *core.FileParam                        `facebook:"upload_gated_file"`
	Extra                            core.Params                            `facebook:"-"`
}

func (params CreatePageLeadgenFormsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowOrganicLeadRetrieval != nil {
		out["allow_organic_lead_retrieval"] = *params.AllowOrganicLeadRetrieval
	}
	if params.BlockDisplayForNonTargetedViewer != nil {
		out["block_display_for_non_targeted_viewer"] = *params.BlockDisplayForNonTargetedViewer
	}
	if params.ContextCard != nil {
		out["context_card"] = *params.ContextCard
	}
	if params.CoverPhoto != nil {
		out["cover_photo"] = *params.CoverPhoto
	}
	if params.CustomDisclaimer != nil {
		out["custom_disclaimer"] = *params.CustomDisclaimer
	}
	if params.FollowUpActionURL != nil {
		out["follow_up_action_url"] = *params.FollowUpActionURL
	}
	if params.IsForCanvas != nil {
		out["is_for_canvas"] = *params.IsForCanvas
	}
	if params.IsOptimizedForQuality != nil {
		out["is_optimized_for_quality"] = *params.IsOptimizedForQuality
	}
	if params.IsPhoneSmsVerifyEnabled != nil {
		out["is_phone_sms_verify_enabled"] = *params.IsPhoneSmsVerifyEnabled
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	out["name"] = params.Name
	if params.PrivacyPolicy != nil {
		out["privacy_policy"] = *params.PrivacyPolicy
	}
	if params.QuestionPageCustomHeadline != nil {
		out["question_page_custom_headline"] = *params.QuestionPageCustomHeadline
	}
	out["questions"] = params.Questions
	if params.ShouldEnforceWorkEmail != nil {
		out["should_enforce_work_email"] = *params.ShouldEnforceWorkEmail
	}
	if params.ThankYouPage != nil {
		out["thank_you_page"] = *params.ThankYouPage
	}
	if params.TrackingParameters != nil {
		out["tracking_parameters"] = *params.TrackingParameters
	}
	if params.UploadGatedFile != nil {
		out["upload_gated_file"] = *params.UploadGatedFile
	}
	return out
}

func CreatePageLeadgenFormsBatchCall(id string, params CreatePageLeadgenFormsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "leadgen_forms"), params.ToParams(), options...)
}

func NewCreatePageLeadgenFormsBatchRequest(id string, params CreatePageLeadgenFormsParams, options ...core.BatchOption) *core.BatchRequest[objects.LeadgenForm] {
	return core.NewBatchRequest[objects.LeadgenForm](CreatePageLeadgenFormsBatchCall(id, params, options...))
}

func DecodeCreatePageLeadgenFormsBatchResponse(response *core.BatchResponse) (*objects.LeadgenForm, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LeadgenForm
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageLeadgenFormsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageLeadgenFormsParams) (*objects.LeadgenForm, *core.Response, error) {
	var out objects.LeadgenForm
	call := CreatePageLeadgenFormsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageLeadgenForms(ctx context.Context, client *core.Client, id string, params CreatePageLeadgenFormsParams) (*objects.LeadgenForm, error) {
	out, _, err := CreatePageLeadgenFormsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageLikesParams struct {
	TargetID *core.ID    `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetPageLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	return out
}

func GetPageLikesBatchCall(id string, params GetPageLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetPageLikesBatchRequest(id string, params GetPageLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPageLikesBatchCall(id, params, options...))
}

func DecodeGetPageLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPageLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageLikesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPageLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageLikes(ctx context.Context, client *core.Client, id string, params GetPageLikesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPageLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageLiveVideosParams struct {
	BroadcastStatus *[]enums.PageliveVideosBroadcastStatusEnumParam `facebook:"broadcast_status"`
	Source          *enums.PageliveVideosSourceEnumParam            `facebook:"source"`
	Extra           core.Params                                     `facebook:"-"`
}

func (params GetPageLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BroadcastStatus != nil {
		out["broadcast_status"] = *params.BroadcastStatus
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetPageLiveVideosBatchCall(id string, params GetPageLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewGetPageLiveVideosBatchRequest(id string, params GetPageLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LiveVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.LiveVideo]](GetPageLiveVideosBatchCall(id, params, options...))
}

func DecodeGetPageLiveVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LiveVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.LiveVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetPageLiveVideosParams) (*core.Cursor[objects.LiveVideo], *core.Response, error) {
	var out core.Cursor[objects.LiveVideo]
	call := GetPageLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageLiveVideos(ctx context.Context, client *core.Client, id string, params GetPageLiveVideosParams) (*core.Cursor[objects.LiveVideo], error) {
	out, _, err := GetPageLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageLiveVideosParams struct {
	ContentTags                *[]string                                        `facebook:"content_tags"`
	CrosspostingActions        *[]map[string]interface{}                        `facebook:"crossposting_actions"`
	CustomLabels               *[]string                                        `facebook:"custom_labels"`
	Description                *string                                          `facebook:"description"`
	EnableBackupIngest         *bool                                            `facebook:"enable_backup_ingest"`
	EncodingSettings           *string                                          `facebook:"encoding_settings"`
	EventParams                *map[string]interface{}                          `facebook:"event_params"`
	FisheyeVideoCropped        *bool                                            `facebook:"fisheye_video_cropped"`
	FrontZRotation             *float64                                         `facebook:"front_z_rotation"`
	GameShow                   *map[string]interface{}                          `facebook:"game_show"`
	IsAudioOnly                *bool                                            `facebook:"is_audio_only"`
	IsSpherical                *bool                                            `facebook:"is_spherical"`
	OriginalFov                *uint64                                          `facebook:"original_fov"`
	Privacy                    *string                                          `facebook:"privacy"`
	Projection                 *enums.PageliveVideosProjectionEnumParam         `facebook:"projection"`
	Published                  *bool                                            `facebook:"published"`
	ScheduleCustomProfileImage *core.FileParam                                  `facebook:"schedule_custom_profile_image"`
	SpatialAudioFormat         *enums.PageliveVideosSpatialAudioFormatEnumParam `facebook:"spatial_audio_format"`
	Status                     *enums.PageliveVideosStatusEnumParam             `facebook:"status"`
	StereoscopicMode           *enums.PageliveVideosStereoscopicModeEnumParam   `facebook:"stereoscopic_mode"`
	StopOnDeleteStream         *bool                                            `facebook:"stop_on_delete_stream"`
	StreamType                 *enums.PageliveVideosStreamTypeEnumParam         `facebook:"stream_type"`
	Targeting                  *map[string]interface{}                          `facebook:"targeting"`
	Title                      *string                                          `facebook:"title"`
	Extra                      core.Params                                      `facebook:"-"`
}

func (params CreatePageLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.CrosspostingActions != nil {
		out["crossposting_actions"] = *params.CrosspostingActions
	}
	if params.CustomLabels != nil {
		out["custom_labels"] = *params.CustomLabels
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableBackupIngest != nil {
		out["enable_backup_ingest"] = *params.EnableBackupIngest
	}
	if params.EncodingSettings != nil {
		out["encoding_settings"] = *params.EncodingSettings
	}
	if params.EventParams != nil {
		out["event_params"] = *params.EventParams
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.GameShow != nil {
		out["game_show"] = *params.GameShow
	}
	if params.IsAudioOnly != nil {
		out["is_audio_only"] = *params.IsAudioOnly
	}
	if params.IsSpherical != nil {
		out["is_spherical"] = *params.IsSpherical
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Projection != nil {
		out["projection"] = *params.Projection
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduleCustomProfileImage != nil {
		out["schedule_custom_profile_image"] = *params.ScheduleCustomProfileImage
	}
	if params.SpatialAudioFormat != nil {
		out["spatial_audio_format"] = *params.SpatialAudioFormat
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StereoscopicMode != nil {
		out["stereoscopic_mode"] = *params.StereoscopicMode
	}
	if params.StopOnDeleteStream != nil {
		out["stop_on_delete_stream"] = *params.StopOnDeleteStream
	}
	if params.StreamType != nil {
		out["stream_type"] = *params.StreamType
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func CreatePageLiveVideosBatchCall(id string, params CreatePageLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewCreatePageLiveVideosBatchRequest(id string, params CreatePageLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](CreatePageLiveVideosBatchCall(id, params, options...))
}

func DecodeCreatePageLiveVideosBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageLiveVideosParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := CreatePageLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageLiveVideos(ctx context.Context, client *core.Client, id string, params CreatePageLiveVideosParams) (*objects.LiveVideo, error) {
	out, _, err := CreatePageLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageLocationsParams struct {
	LocationPageIds []core.ID   `facebook:"location_page_ids"`
	StoreNumbers    []uint64    `facebook:"store_numbers"`
	Extra           core.Params `facebook:"-"`
}

func (params DeletePageLocationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["location_page_ids"] = params.LocationPageIds
	out["store_numbers"] = params.StoreNumbers
	return out
}

func DeletePageLocationsBatchCall(id string, params DeletePageLocationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "locations"), params.ToParams(), options...)
}

func NewDeletePageLocationsBatchRequest(id string, params DeletePageLocationsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageLocationsBatchCall(id, params, options...))
}

func DecodeDeletePageLocationsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageLocationsWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageLocationsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageLocationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageLocations(ctx context.Context, client *core.Client, id string, params DeletePageLocationsParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageLocationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageLocationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageLocationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageLocationsBatchCall(id string, params GetPageLocationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "locations"), params.ToParams(), options...)
}

func NewGetPageLocationsBatchRequest(id string, params GetPageLocationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPageLocationsBatchCall(id, params, options...))
}

func DecodeGetPageLocationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPageLocationsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageLocationsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPageLocationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageLocations(ctx context.Context, client *core.Client, id string, params GetPageLocationsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPageLocationsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageLocationsParams struct {
	AlwaysOpen                  *bool                                             `facebook:"always_open"`
	DeliveryAndPickupOptionInfo *[]string                                         `facebook:"delivery_and_pickup_option_info"`
	DifferentlyOpenOfferings    *map[string]interface{}                           `facebook:"differently_open_offerings"`
	Hours                       *map[string]interface{}                           `facebook:"hours"`
	IgnoreWarnings              *bool                                             `facebook:"ignore_warnings"`
	Location                    *map[string]interface{}                           `facebook:"location"`
	LocationPageID              *core.ID                                          `facebook:"location_page_id"`
	OldStoreNumber              *uint64                                           `facebook:"old_store_number"`
	PageUsername                *string                                           `facebook:"page_username"`
	PermanentlyClosed           *bool                                             `facebook:"permanently_closed"`
	Phone                       *string                                           `facebook:"phone"`
	PickupOptions               *[]enums.PagelocationsPickupOptionsEnumParam      `facebook:"pickup_options"`
	PlaceTopics                 *[]string                                         `facebook:"place_topics"`
	PriceRange                  *string                                           `facebook:"price_range"`
	RecommendationAction        *enums.PagelocationsRecommendationActionEnumParam `facebook:"recommendation_action"`
	RecommendationDs            *string                                           `facebook:"recommendation_ds"`
	RecommendationStoreID       *core.ID                                          `facebook:"recommendation_store_id"`
	StoreCode                   *string                                           `facebook:"store_code"`
	StoreLocationDescriptor     *string                                           `facebook:"store_location_descriptor"`
	StoreName                   *string                                           `facebook:"store_name"`
	StoreNumber                 uint64                                            `facebook:"store_number"`
	TemporaryStatus             *enums.PagelocationsTemporaryStatusEnumParam      `facebook:"temporary_status"`
	Type                        *string                                           `facebook:"type"`
	Website                     *string                                           `facebook:"website"`
	Extra                       core.Params                                       `facebook:"-"`
}

func (params CreatePageLocationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AlwaysOpen != nil {
		out["always_open"] = *params.AlwaysOpen
	}
	if params.DeliveryAndPickupOptionInfo != nil {
		out["delivery_and_pickup_option_info"] = *params.DeliveryAndPickupOptionInfo
	}
	if params.DifferentlyOpenOfferings != nil {
		out["differently_open_offerings"] = *params.DifferentlyOpenOfferings
	}
	if params.Hours != nil {
		out["hours"] = *params.Hours
	}
	if params.IgnoreWarnings != nil {
		out["ignore_warnings"] = *params.IgnoreWarnings
	}
	if params.Location != nil {
		out["location"] = *params.Location
	}
	if params.LocationPageID != nil {
		out["location_page_id"] = *params.LocationPageID
	}
	if params.OldStoreNumber != nil {
		out["old_store_number"] = *params.OldStoreNumber
	}
	if params.PageUsername != nil {
		out["page_username"] = *params.PageUsername
	}
	if params.PermanentlyClosed != nil {
		out["permanently_closed"] = *params.PermanentlyClosed
	}
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.PickupOptions != nil {
		out["pickup_options"] = *params.PickupOptions
	}
	if params.PlaceTopics != nil {
		out["place_topics"] = *params.PlaceTopics
	}
	if params.PriceRange != nil {
		out["price_range"] = *params.PriceRange
	}
	if params.RecommendationAction != nil {
		out["recommendation_action"] = *params.RecommendationAction
	}
	if params.RecommendationDs != nil {
		out["recommendation_ds"] = *params.RecommendationDs
	}
	if params.RecommendationStoreID != nil {
		out["recommendation_store_id"] = *params.RecommendationStoreID
	}
	if params.StoreCode != nil {
		out["store_code"] = *params.StoreCode
	}
	if params.StoreLocationDescriptor != nil {
		out["store_location_descriptor"] = *params.StoreLocationDescriptor
	}
	if params.StoreName != nil {
		out["store_name"] = *params.StoreName
	}
	out["store_number"] = params.StoreNumber
	if params.TemporaryStatus != nil {
		out["temporary_status"] = *params.TemporaryStatus
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Website != nil {
		out["website"] = *params.Website
	}
	return out
}

func CreatePageLocationsBatchCall(id string, params CreatePageLocationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "locations"), params.ToParams(), options...)
}

func NewCreatePageLocationsBatchRequest(id string, params CreatePageLocationsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageLocationsBatchCall(id, params, options...))
}

func DecodeCreatePageLocationsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageLocationsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageLocationsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageLocationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageLocations(ctx context.Context, client *core.Client, id string, params CreatePageLocationsParams) (*objects.Page, error) {
	out, _, err := CreatePageLocationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMediaFingerprintsParams struct {
	UniversalContentID *core.ID    `facebook:"universal_content_id"`
	Extra              core.Params `facebook:"-"`
}

func (params GetPageMediaFingerprintsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UniversalContentID != nil {
		out["universal_content_id"] = *params.UniversalContentID
	}
	return out
}

func GetPageMediaFingerprintsBatchCall(id string, params GetPageMediaFingerprintsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "media_fingerprints"), params.ToParams(), options...)
}

func NewGetPageMediaFingerprintsBatchRequest(id string, params GetPageMediaFingerprintsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MediaFingerprint]] {
	return core.NewBatchRequest[core.Cursor[objects.MediaFingerprint]](GetPageMediaFingerprintsBatchCall(id, params, options...))
}

func DecodeGetPageMediaFingerprintsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MediaFingerprint], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MediaFingerprint]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMediaFingerprintsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMediaFingerprintsParams) (*core.Cursor[objects.MediaFingerprint], *core.Response, error) {
	var out core.Cursor[objects.MediaFingerprint]
	call := GetPageMediaFingerprintsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMediaFingerprints(ctx context.Context, client *core.Client, id string, params GetPageMediaFingerprintsParams) (*core.Cursor[objects.MediaFingerprint], error) {
	out, _, err := GetPageMediaFingerprintsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMediaFingerprintsParams struct {
	FingerprintContentType enums.PagemediaFingerprintsFingerprintContentTypeEnumParam `facebook:"fingerprint_content_type"`
	Metadata               []interface{}                                              `facebook:"metadata"`
	Source                 string                                                     `facebook:"source"`
	Title                  string                                                     `facebook:"title"`
	UniversalContentID     *core.ID                                                   `facebook:"universal_content_id"`
	Extra                  core.Params                                                `facebook:"-"`
}

func (params CreatePageMediaFingerprintsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["fingerprint_content_type"] = params.FingerprintContentType
	out["metadata"] = params.Metadata
	out["source"] = params.Source
	out["title"] = params.Title
	if params.UniversalContentID != nil {
		out["universal_content_id"] = *params.UniversalContentID
	}
	return out
}

func CreatePageMediaFingerprintsBatchCall(id string, params CreatePageMediaFingerprintsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "media_fingerprints"), params.ToParams(), options...)
}

func NewCreatePageMediaFingerprintsBatchRequest(id string, params CreatePageMediaFingerprintsParams, options ...core.BatchOption) *core.BatchRequest[objects.MediaFingerprint] {
	return core.NewBatchRequest[objects.MediaFingerprint](CreatePageMediaFingerprintsBatchCall(id, params, options...))
}

func DecodeCreatePageMediaFingerprintsBatchResponse(response *core.BatchResponse) (*objects.MediaFingerprint, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MediaFingerprint
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMediaFingerprintsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMediaFingerprintsParams) (*objects.MediaFingerprint, *core.Response, error) {
	var out objects.MediaFingerprint
	call := CreatePageMediaFingerprintsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMediaFingerprints(ctx context.Context, client *core.Client, id string, params CreatePageMediaFingerprintsParams) (*objects.MediaFingerprint, error) {
	out, _, err := CreatePageMediaFingerprintsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessageAttachmentsParams struct {
	Message  map[string]interface{}                         `facebook:"message"`
	Platform *enums.PagemessageAttachmentsPlatformEnumParam `facebook:"platform"`
	Extra    core.Params                                    `facebook:"-"`
}

func (params CreatePageMessageAttachmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["message"] = params.Message
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	return out
}

func CreatePageMessageAttachmentsBatchCall(id string, params CreatePageMessageAttachmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "message_attachments"), params.ToParams(), options...)
}

func NewCreatePageMessageAttachmentsBatchRequest(id string, params CreatePageMessageAttachmentsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageMessageAttachmentsBatchCall(id, params, options...))
}

func DecodeCreatePageMessageAttachmentsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageMessageAttachmentsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessageAttachmentsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageMessageAttachmentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessageAttachments(ctx context.Context, client *core.Client, id string, params CreatePageMessageAttachmentsParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageMessageAttachmentsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageMessageTemplatesParams struct {
	Name       string      `facebook:"name"`
	TemplateID *core.ID    `facebook:"template_id"`
	Extra      core.Params `facebook:"-"`
}

func (params DeletePageMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	if params.TemplateID != nil {
		out["template_id"] = *params.TemplateID
	}
	return out
}

func DeletePageMessageTemplatesBatchCall(id string, params DeletePageMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewDeletePageMessageTemplatesBatchRequest(id string, params DeletePageMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageMessageTemplatesBatchCall(id, params, options...))
}

func DecodeDeletePageMessageTemplatesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageMessageTemplatesWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageMessageTemplatesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageMessageTemplatesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageMessageTemplates(ctx context.Context, client *core.Client, id string, params DeletePageMessageTemplatesParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageMessageTemplatesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessageTemplatesParams struct {
	Category      *[]enums.PagemessageTemplatesCategoryEnumParam `facebook:"category"`
	Content       *string                                        `facebook:"content"`
	Language      *[]string                                      `facebook:"language"`
	Name          *string                                        `facebook:"name"`
	NameOrContent *string                                        `facebook:"name_or_content"`
	Status        *[]enums.PagemessageTemplatesStatusEnumParam   `facebook:"status"`
	Extra         core.Params                                    `facebook:"-"`
}

func (params GetPageMessageTemplatesParams) ToParams() core.Params {
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
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetPageMessageTemplatesBatchCall(id string, params GetPageMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewGetPageMessageTemplatesBatchRequest(id string, params GetPageMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerBusinessTemplate]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerBusinessTemplate]](GetPageMessageTemplatesBatchCall(id, params, options...))
}

func DecodeGetPageMessageTemplatesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerBusinessTemplate], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerBusinessTemplate]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessageTemplatesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessageTemplatesParams) (*core.Cursor[objects.MessengerBusinessTemplate], *core.Response, error) {
	var out core.Cursor[objects.MessengerBusinessTemplate]
	call := GetPageMessageTemplatesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessageTemplates(ctx context.Context, client *core.Client, id string, params GetPageMessageTemplatesParams) (*core.Cursor[objects.MessengerBusinessTemplate], error) {
	out, _, err := GetPageMessageTemplatesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessageTemplatesParams struct {
	Category                    enums.PagemessageTemplatesCategoryEnumParam         `facebook:"category"`
	Components                  *[]map[string]interface{}                           `facebook:"components"`
	Language                    string                                              `facebook:"language"`
	LibraryTemplateButtonInputs *[]map[string]interface{}                           `facebook:"library_template_button_inputs"`
	LibraryTemplateName         *string                                             `facebook:"library_template_name"`
	Name                        string                                              `facebook:"name"`
	ParameterFormat             *enums.PagemessageTemplatesParameterFormatEnumParam `facebook:"parameter_format"`
	Extra                       core.Params                                         `facebook:"-"`
}

func (params CreatePageMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["category"] = params.Category
	if params.Components != nil {
		out["components"] = *params.Components
	}
	out["language"] = params.Language
	if params.LibraryTemplateButtonInputs != nil {
		out["library_template_button_inputs"] = *params.LibraryTemplateButtonInputs
	}
	if params.LibraryTemplateName != nil {
		out["library_template_name"] = *params.LibraryTemplateName
	}
	out["name"] = params.Name
	if params.ParameterFormat != nil {
		out["parameter_format"] = *params.ParameterFormat
	}
	return out
}

func CreatePageMessageTemplatesBatchCall(id string, params CreatePageMessageTemplatesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "message_templates"), params.ToParams(), options...)
}

func NewCreatePageMessageTemplatesBatchRequest(id string, params CreatePageMessageTemplatesParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageMessageTemplatesBatchCall(id, params, options...))
}

func DecodeCreatePageMessageTemplatesBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMessageTemplatesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessageTemplatesParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageMessageTemplatesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessageTemplates(ctx context.Context, client *core.Client, id string, params CreatePageMessageTemplatesParams) (*objects.Page, error) {
	out, _, err := CreatePageMessageTemplatesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessagesParams struct {
	Folder           *enums.PagemessagesFolderEnumParam           `facebook:"folder"`
	Message          *map[string]interface{}                      `facebook:"message"`
	MessagingType    *enums.PagemessagesMessagingTypeEnumParam    `facebook:"messaging_type"`
	NotificationType *enums.PagemessagesNotificationTypeEnumParam `facebook:"notification_type"`
	Payload          *string                                      `facebook:"payload"`
	PersonaID        *core.ID                                     `facebook:"persona_id"`
	Recipient        map[string]interface{}                       `facebook:"recipient"`
	ReplyTo          *map[string]interface{}                      `facebook:"reply_to"`
	SenderAction     *enums.PagemessagesSenderActionEnumParam     `facebook:"sender_action"`
	SuggestionAction *enums.PagemessagesSuggestionActionEnumParam `facebook:"suggestion_action"`
	Tag              *map[string]interface{}                      `facebook:"tag"`
	ThreadControl    *map[string]interface{}                      `facebook:"thread_control"`
	Extra            core.Params                                  `facebook:"-"`
}

func (params CreatePageMessagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.MessagingType != nil {
		out["messaging_type"] = *params.MessagingType
	}
	if params.NotificationType != nil {
		out["notification_type"] = *params.NotificationType
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.PersonaID != nil {
		out["persona_id"] = *params.PersonaID
	}
	out["recipient"] = params.Recipient
	if params.ReplyTo != nil {
		out["reply_to"] = *params.ReplyTo
	}
	if params.SenderAction != nil {
		out["sender_action"] = *params.SenderAction
	}
	if params.SuggestionAction != nil {
		out["suggestion_action"] = *params.SuggestionAction
	}
	if params.Tag != nil {
		out["tag"] = *params.Tag
	}
	if params.ThreadControl != nil {
		out["thread_control"] = *params.ThreadControl
	}
	return out
}

func CreatePageMessagesBatchCall(id string, params CreatePageMessagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messages"), params.ToParams(), options...)
}

func NewCreatePageMessagesBatchRequest(id string, params CreatePageMessagesParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageMessagesBatchCall(id, params, options...))
}

func DecodeCreatePageMessagesBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMessagesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessagesParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageMessagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessages(ctx context.Context, client *core.Client, id string, params CreatePageMessagesParams) (*objects.Page, error) {
	out, _, err := CreatePageMessagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessagingFeatureReviewParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageMessagingFeatureReviewParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageMessagingFeatureReviewBatchCall(id string, params GetPageMessagingFeatureReviewParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messaging_feature_review"), params.ToParams(), options...)
}

func NewGetPageMessagingFeatureReviewBatchRequest(id string, params GetPageMessagingFeatureReviewParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessagingFeatureReview]] {
	return core.NewBatchRequest[core.Cursor[objects.MessagingFeatureReview]](GetPageMessagingFeatureReviewBatchCall(id, params, options...))
}

func DecodeGetPageMessagingFeatureReviewBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessagingFeatureReview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessagingFeatureReview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessagingFeatureReviewWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessagingFeatureReviewParams) (*core.Cursor[objects.MessagingFeatureReview], *core.Response, error) {
	var out core.Cursor[objects.MessagingFeatureReview]
	call := GetPageMessagingFeatureReviewBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessagingFeatureReview(ctx context.Context, client *core.Client, id string, params GetPageMessagingFeatureReviewParams) (*core.Cursor[objects.MessagingFeatureReview], error) {
	out, _, err := GetPageMessagingFeatureReviewWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessengerCallPermissionsParams struct {
	Psid  string      `facebook:"psid"`
	Extra core.Params `facebook:"-"`
}

func (params GetPageMessengerCallPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["psid"] = params.Psid
	return out
}

func GetPageMessengerCallPermissionsBatchCall(id string, params GetPageMessengerCallPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messenger_call_permissions"), params.ToParams(), options...)
}

func NewGetPageMessengerCallPermissionsBatchRequest(id string, params GetPageMessengerCallPermissionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerCallPermissions]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerCallPermissions]](GetPageMessengerCallPermissionsBatchCall(id, params, options...))
}

func DecodeGetPageMessengerCallPermissionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerCallPermissions], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerCallPermissions]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessengerCallPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessengerCallPermissionsParams) (*core.Cursor[objects.MessengerCallPermissions], *core.Response, error) {
	var out core.Cursor[objects.MessengerCallPermissions]
	call := GetPageMessengerCallPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessengerCallPermissions(ctx context.Context, client *core.Client, id string, params GetPageMessengerCallPermissionsParams) (*core.Cursor[objects.MessengerCallPermissions], error) {
	out, _, err := GetPageMessengerCallPermissionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessengerCallSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageMessengerCallSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageMessengerCallSettingsBatchCall(id string, params GetPageMessengerCallSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messenger_call_settings"), params.ToParams(), options...)
}

func NewGetPageMessengerCallSettingsBatchRequest(id string, params GetPageMessengerCallSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerCallSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerCallSettings]](GetPageMessengerCallSettingsBatchCall(id, params, options...))
}

func DecodeGetPageMessengerCallSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerCallSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerCallSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessengerCallSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessengerCallSettingsParams) (*core.Cursor[objects.MessengerCallSettings], *core.Response, error) {
	var out core.Cursor[objects.MessengerCallSettings]
	call := GetPageMessengerCallSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessengerCallSettings(ctx context.Context, client *core.Client, id string, params GetPageMessengerCallSettingsParams) (*core.Cursor[objects.MessengerCallSettings], error) {
	out, _, err := GetPageMessengerCallSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessengerCallSettingsParams struct {
	AudioEnabled *bool                   `facebook:"audio_enabled"`
	CallHours    *map[string]interface{} `facebook:"call_hours"`
	CallRouting  *map[string]interface{} `facebook:"call_routing"`
	IconEnabled  *bool                   `facebook:"icon_enabled"`
	VideoEnabled *bool                   `facebook:"video_enabled"`
	Extra        core.Params             `facebook:"-"`
}

func (params CreatePageMessengerCallSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AudioEnabled != nil {
		out["audio_enabled"] = *params.AudioEnabled
	}
	if params.CallHours != nil {
		out["call_hours"] = *params.CallHours
	}
	if params.CallRouting != nil {
		out["call_routing"] = *params.CallRouting
	}
	if params.IconEnabled != nil {
		out["icon_enabled"] = *params.IconEnabled
	}
	if params.VideoEnabled != nil {
		out["video_enabled"] = *params.VideoEnabled
	}
	return out
}

func CreatePageMessengerCallSettingsBatchCall(id string, params CreatePageMessengerCallSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_call_settings"), params.ToParams(), options...)
}

func NewCreatePageMessengerCallSettingsBatchRequest(id string, params CreatePageMessengerCallSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageMessengerCallSettingsBatchCall(id, params, options...))
}

func DecodeCreatePageMessengerCallSettingsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMessengerCallSettingsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessengerCallSettingsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageMessengerCallSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessengerCallSettings(ctx context.Context, client *core.Client, id string, params CreatePageMessengerCallSettingsParams) (*objects.Page, error) {
	out, _, err := CreatePageMessengerCallSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessengerLeadFormsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageMessengerLeadFormsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageMessengerLeadFormsBatchCall(id string, params GetPageMessengerLeadFormsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messenger_lead_forms"), params.ToParams(), options...)
}

func NewGetPageMessengerLeadFormsBatchRequest(id string, params GetPageMessengerLeadFormsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerAdsPartialAutomatedStepList]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerAdsPartialAutomatedStepList]](GetPageMessengerLeadFormsBatchCall(id, params, options...))
}

func DecodeGetPageMessengerLeadFormsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerAdsPartialAutomatedStepList], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerAdsPartialAutomatedStepList]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessengerLeadFormsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessengerLeadFormsParams) (*core.Cursor[objects.MessengerAdsPartialAutomatedStepList], *core.Response, error) {
	var out core.Cursor[objects.MessengerAdsPartialAutomatedStepList]
	call := GetPageMessengerLeadFormsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessengerLeadForms(ctx context.Context, client *core.Client, id string, params GetPageMessengerLeadFormsParams) (*core.Cursor[objects.MessengerAdsPartialAutomatedStepList], error) {
	out, _, err := GetPageMessengerLeadFormsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessengerLeadFormsParams struct {
	AccountID           *core.ID                 `facebook:"account_id"`
	BlockSendAPI        *bool                    `facebook:"block_send_api"`
	ExitKeyphrases      *string                  `facebook:"exit_keyphrases"`
	HandoverAppID       *core.ID                 `facebook:"handover_app_id"`
	HandoverSummary     *bool                    `facebook:"handover_summary"`
	PrivacyURL          *string                  `facebook:"privacy_url"`
	ReminderText        *string                  `facebook:"reminder_text"`
	StepList            []map[string]interface{} `facebook:"step_list"`
	StopQuestionMessage *string                  `facebook:"stop_question_message"`
	TemplateName        *string                  `facebook:"template_name"`
	TrackingParameters  *map[string]interface{}  `facebook:"tracking_parameters"`
	Extra               core.Params              `facebook:"-"`
}

func (params CreatePageMessengerLeadFormsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.BlockSendAPI != nil {
		out["block_send_api"] = *params.BlockSendAPI
	}
	if params.ExitKeyphrases != nil {
		out["exit_keyphrases"] = *params.ExitKeyphrases
	}
	if params.HandoverAppID != nil {
		out["handover_app_id"] = *params.HandoverAppID
	}
	if params.HandoverSummary != nil {
		out["handover_summary"] = *params.HandoverSummary
	}
	if params.PrivacyURL != nil {
		out["privacy_url"] = *params.PrivacyURL
	}
	if params.ReminderText != nil {
		out["reminder_text"] = *params.ReminderText
	}
	out["step_list"] = params.StepList
	if params.StopQuestionMessage != nil {
		out["stop_question_message"] = *params.StopQuestionMessage
	}
	if params.TemplateName != nil {
		out["template_name"] = *params.TemplateName
	}
	if params.TrackingParameters != nil {
		out["tracking_parameters"] = *params.TrackingParameters
	}
	return out
}

func CreatePageMessengerLeadFormsBatchCall(id string, params CreatePageMessengerLeadFormsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_lead_forms"), params.ToParams(), options...)
}

func NewCreatePageMessengerLeadFormsBatchRequest(id string, params CreatePageMessengerLeadFormsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageMessengerLeadFormsBatchCall(id, params, options...))
}

func DecodeCreatePageMessengerLeadFormsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMessengerLeadFormsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessengerLeadFormsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageMessengerLeadFormsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessengerLeadForms(ctx context.Context, client *core.Client, id string, params CreatePageMessengerLeadFormsParams) (*objects.Page, error) {
	out, _, err := CreatePageMessengerLeadFormsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageMessengerProfileParams struct {
	Fields   []enums.PagemessengerProfileFieldsEnumParam  `facebook:"fields"`
	Platform *enums.PagemessengerProfilePlatformEnumParam `facebook:"platform"`
	Extra    core.Params                                  `facebook:"-"`
}

func (params DeletePageMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["fields"] = params.Fields
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	return out
}

func DeletePageMessengerProfileBatchCall(id string, params DeletePageMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewDeletePageMessengerProfileBatchRequest(id string, params DeletePageMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageMessengerProfileBatchCall(id, params, options...))
}

func DecodeDeletePageMessengerProfileBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageMessengerProfileParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageMessengerProfile(ctx context.Context, client *core.Client, id string, params DeletePageMessengerProfileParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageMessengerProfileParams struct {
	Platform *enums.PagemessengerProfilePlatformEnumParam `facebook:"platform"`
	Extra    core.Params                                  `facebook:"-"`
}

func (params GetPageMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	return out
}

func GetPageMessengerProfileBatchCall(id string, params GetPageMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewGetPageMessengerProfileBatchRequest(id string, params GetPageMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerProfile]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerProfile]](GetPageMessengerProfileBatchCall(id, params, options...))
}

func DecodeGetPageMessengerProfileBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerProfile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerProfile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params GetPageMessengerProfileParams) (*core.Cursor[objects.MessengerProfile], *core.Response, error) {
	var out core.Cursor[objects.MessengerProfile]
	call := GetPageMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageMessengerProfile(ctx context.Context, client *core.Client, id string, params GetPageMessengerProfileParams) (*core.Cursor[objects.MessengerProfile], error) {
	out, _, err := GetPageMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageMessengerProfileParams struct {
	AccountLinkingURL  *string                                      `facebook:"account_linking_url"`
	Commands           *[]map[string]interface{}                    `facebook:"commands"`
	Description        *[]map[string]interface{}                    `facebook:"description"`
	GetStarted         *map[string]interface{}                      `facebook:"get_started"`
	IceBreakers        *[]map[string]interface{}                    `facebook:"ice_breakers"`
	PersistentMenu     *[]map[string]interface{}                    `facebook:"persistent_menu"`
	Platform           *enums.PagemessengerProfilePlatformEnumParam `facebook:"platform"`
	WhitelistedDomains *[]string                                    `facebook:"whitelisted_domains"`
	Extra              core.Params                                  `facebook:"-"`
}

func (params CreatePageMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountLinkingURL != nil {
		out["account_linking_url"] = *params.AccountLinkingURL
	}
	if params.Commands != nil {
		out["commands"] = *params.Commands
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GetStarted != nil {
		out["get_started"] = *params.GetStarted
	}
	if params.IceBreakers != nil {
		out["ice_breakers"] = *params.IceBreakers
	}
	if params.PersistentMenu != nil {
		out["persistent_menu"] = *params.PersistentMenu
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.WhitelistedDomains != nil {
		out["whitelisted_domains"] = *params.WhitelistedDomains
	}
	return out
}

func CreatePageMessengerProfileBatchCall(id string, params CreatePageMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewCreatePageMessengerProfileBatchRequest(id string, params CreatePageMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageMessengerProfileBatchCall(id, params, options...))
}

func DecodeCreatePageMessengerProfileBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageMessengerProfileParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageMessengerProfile(ctx context.Context, client *core.Client, id string, params CreatePageMessengerProfileParams) (*objects.Page, error) {
	out, _, err := CreatePageMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageModerateConversationsParams struct {
	Actions []enums.PagemoderateConversationsActionsEnumParam `facebook:"actions"`
	UserIds []map[string]interface{}                          `facebook:"user_ids"`
	Extra   core.Params                                       `facebook:"-"`
}

func (params CreatePageModerateConversationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["actions"] = params.Actions
	out["user_ids"] = params.UserIds
	return out
}

func CreatePageModerateConversationsBatchCall(id string, params CreatePageModerateConversationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "moderate_conversations"), params.ToParams(), options...)
}

func NewCreatePageModerateConversationsBatchRequest(id string, params CreatePageModerateConversationsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageModerateConversationsBatchCall(id, params, options...))
}

func DecodeCreatePageModerateConversationsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageModerateConversationsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageModerateConversationsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageModerateConversationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageModerateConversations(ctx context.Context, client *core.Client, id string, params CreatePageModerateConversationsParams) (*objects.Page, error) {
	out, _, err := CreatePageModerateConversationsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageNlpConfigsParams struct {
	APIVersion           *map[string]interface{}             `facebook:"api_version"`
	CustomToken          *string                             `facebook:"custom_token"`
	Model                *enums.PagenlpConfigsModelEnumParam `facebook:"model"`
	NBest                *uint64                             `facebook:"n_best"`
	NlpEnabled           *bool                               `facebook:"nlp_enabled"`
	OtherLanguageSupport *map[string]interface{}             `facebook:"other_language_support"`
	Verbose              *bool                               `facebook:"verbose"`
	Extra                core.Params                         `facebook:"-"`
}

func (params CreatePageNlpConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.APIVersion != nil {
		out["api_version"] = *params.APIVersion
	}
	if params.CustomToken != nil {
		out["custom_token"] = *params.CustomToken
	}
	if params.Model != nil {
		out["model"] = *params.Model
	}
	if params.NBest != nil {
		out["n_best"] = *params.NBest
	}
	if params.NlpEnabled != nil {
		out["nlp_enabled"] = *params.NlpEnabled
	}
	if params.OtherLanguageSupport != nil {
		out["other_language_support"] = *params.OtherLanguageSupport
	}
	if params.Verbose != nil {
		out["verbose"] = *params.Verbose
	}
	return out
}

func CreatePageNlpConfigsBatchCall(id string, params CreatePageNlpConfigsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "nlp_configs"), params.ToParams(), options...)
}

func NewCreatePageNlpConfigsBatchRequest(id string, params CreatePageNlpConfigsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageNlpConfigsBatchCall(id, params, options...))
}

func DecodeCreatePageNlpConfigsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageNlpConfigsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageNlpConfigsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageNlpConfigsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageNlpConfigs(ctx context.Context, client *core.Client, id string, params CreatePageNlpConfigsParams) (*objects.Page, error) {
	out, _, err := CreatePageNlpConfigsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageNotificationMessageTokensParams struct {
	CustomAudienceIds           *[]core.ID                                                 `facebook:"custom_audience_ids"`
	DoNotReturnDuplicates       *bool                                                      `facebook:"do_not_return_duplicates"`
	HasReceivedMarketingMessage *bool                                                      `facebook:"has_received_marketing_message"`
	OptInSource                 *[]enums.PagenotificationMessageTokensOptInSourceEnumParam `facebook:"opt_in_source"`
	Since                       *core.Time                                                 `facebook:"since"`
	SubscriberTagIds            *[]core.ID                                                 `facebook:"subscriber_tag_ids"`
	Until                       *core.Time                                                 `facebook:"until"`
	Extra                       core.Params                                                `facebook:"-"`
}

func (params GetPageNotificationMessageTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CustomAudienceIds != nil {
		out["custom_audience_ids"] = *params.CustomAudienceIds
	}
	if params.DoNotReturnDuplicates != nil {
		out["do_not_return_duplicates"] = *params.DoNotReturnDuplicates
	}
	if params.HasReceivedMarketingMessage != nil {
		out["has_received_marketing_message"] = *params.HasReceivedMarketingMessage
	}
	if params.OptInSource != nil {
		out["opt_in_source"] = *params.OptInSource
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.SubscriberTagIds != nil {
		out["subscriber_tag_ids"] = *params.SubscriberTagIds
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPageNotificationMessageTokensBatchCall(id string, params GetPageNotificationMessageTokensParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "notification_message_tokens"), params.ToParams(), options...)
}

func NewGetPageNotificationMessageTokensBatchRequest(id string, params GetPageNotificationMessageTokensParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserPageOneTimeOptInTokenSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.UserPageOneTimeOptInTokenSettings]](GetPageNotificationMessageTokensBatchCall(id, params, options...))
}

func DecodeGetPageNotificationMessageTokensBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserPageOneTimeOptInTokenSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageNotificationMessageTokensWithResponse(ctx context.Context, client *core.Client, id string, params GetPageNotificationMessageTokensParams) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], *core.Response, error) {
	var out core.Cursor[objects.UserPageOneTimeOptInTokenSettings]
	call := GetPageNotificationMessageTokensBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageNotificationMessageTokens(ctx context.Context, client *core.Client, id string, params GetPageNotificationMessageTokensParams) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], error) {
	out, _, err := GetPageNotificationMessageTokensWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageNotificationMessagesDevSupportParams struct {
	DeveloperAction enums.PagenotificationMessagesDevSupportDeveloperActionEnumParam `facebook:"developer_action"`
	Recipient       map[string]interface{}                                           `facebook:"recipient"`
	Extra           core.Params                                                      `facebook:"-"`
}

func (params CreatePageNotificationMessagesDevSupportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["developer_action"] = params.DeveloperAction
	out["recipient"] = params.Recipient
	return out
}

func CreatePageNotificationMessagesDevSupportBatchCall(id string, params CreatePageNotificationMessagesDevSupportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "notification_messages_dev_support"), params.ToParams(), options...)
}

func NewCreatePageNotificationMessagesDevSupportBatchRequest(id string, params CreatePageNotificationMessagesDevSupportParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageNotificationMessagesDevSupportBatchCall(id, params, options...))
}

func DecodeCreatePageNotificationMessagesDevSupportBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageNotificationMessagesDevSupportWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageNotificationMessagesDevSupportParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageNotificationMessagesDevSupportBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageNotificationMessagesDevSupport(ctx context.Context, client *core.Client, id string, params CreatePageNotificationMessagesDevSupportParams) (*objects.Page, error) {
	out, _, err := CreatePageNotificationMessagesDevSupportWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePageBackedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePageBackedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePageBackedInstagramAccountsBatchCall(id string, params GetPagePageBackedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "page_backed_instagram_accounts"), params.ToParams(), options...)
}

func NewGetPagePageBackedInstagramAccountsBatchRequest(id string, params GetPagePageBackedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetPagePageBackedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetPagePageBackedInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
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

func GetPagePageBackedInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePageBackedInstagramAccountsParams) (*core.Cursor[objects.IGUser], *core.Response, error) {
	var out core.Cursor[objects.IGUser]
	call := GetPagePageBackedInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePageBackedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetPagePageBackedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	out, _, err := GetPagePageBackedInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePageBackedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreatePagePageBackedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreatePagePageBackedInstagramAccountsBatchCall(id string, params CreatePagePageBackedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "page_backed_instagram_accounts"), params.ToParams(), options...)
}

func NewCreatePagePageBackedInstagramAccountsBatchRequest(id string, params CreatePagePageBackedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUser] {
	return core.NewBatchRequest[objects.IGUser](CreatePagePageBackedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeCreatePagePageBackedInstagramAccountsBatchResponse(response *core.BatchResponse) (*objects.IGUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePageBackedInstagramAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePageBackedInstagramAccountsParams) (*objects.IGUser, *core.Response, error) {
	var out objects.IGUser
	call := CreatePagePageBackedInstagramAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePageBackedInstagramAccounts(ctx context.Context, client *core.Client, id string, params CreatePagePageBackedInstagramAccountsParams) (*objects.IGUser, error) {
	out, _, err := CreatePagePageBackedInstagramAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePageBackedThreadsAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreatePagePageBackedThreadsAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreatePagePageBackedThreadsAccountsBatchCall(id string, params CreatePagePageBackedThreadsAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "page_backed_threads_accounts"), params.ToParams(), options...)
}

func NewCreatePagePageBackedThreadsAccountsBatchRequest(id string, params CreatePagePageBackedThreadsAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePagePageBackedThreadsAccountsBatchCall(id, params, options...))
}

func DecodeCreatePagePageBackedThreadsAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePagePageBackedThreadsAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePageBackedThreadsAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePagePageBackedThreadsAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePageBackedThreadsAccounts(ctx context.Context, client *core.Client, id string, params CreatePagePageBackedThreadsAccountsParams) (*map[string]interface{}, error) {
	out, _, err := CreatePagePageBackedThreadsAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePageWhatsappNumberVerificationParams struct {
	VerificationCode *string     `facebook:"verification_code"`
	WhatsappNumber   string      `facebook:"whatsapp_number"`
	Extra            core.Params `facebook:"-"`
}

func (params CreatePagePageWhatsappNumberVerificationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.VerificationCode != nil {
		out["verification_code"] = *params.VerificationCode
	}
	out["whatsapp_number"] = params.WhatsappNumber
	return out
}

func CreatePagePageWhatsappNumberVerificationBatchCall(id string, params CreatePagePageWhatsappNumberVerificationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "page_whatsapp_number_verification"), params.ToParams(), options...)
}

func NewCreatePagePageWhatsappNumberVerificationBatchRequest(id string, params CreatePagePageWhatsappNumberVerificationParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePagePageWhatsappNumberVerificationBatchCall(id, params, options...))
}

func DecodeCreatePagePageWhatsappNumberVerificationBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePageWhatsappNumberVerificationWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePageWhatsappNumberVerificationParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePagePageWhatsappNumberVerificationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePageWhatsappNumberVerification(ctx context.Context, client *core.Client, id string, params CreatePagePageWhatsappNumberVerificationParams) (*objects.Page, error) {
	out, _, err := CreatePagePageWhatsappNumberVerificationWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePassThreadControlParams struct {
	Metadata    *string                `facebook:"metadata"`
	Recipient   map[string]interface{} `facebook:"recipient"`
	TargetAppID *core.ID               `facebook:"target_app_id"`
	Extra       core.Params            `facebook:"-"`
}

func (params CreatePagePassThreadControlParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	out["recipient"] = params.Recipient
	if params.TargetAppID != nil {
		out["target_app_id"] = *params.TargetAppID
	}
	return out
}

func CreatePagePassThreadControlBatchCall(id string, params CreatePagePassThreadControlParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "pass_thread_control"), params.ToParams(), options...)
}

func NewCreatePagePassThreadControlBatchRequest(id string, params CreatePagePassThreadControlParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePagePassThreadControlBatchCall(id, params, options...))
}

func DecodeCreatePagePassThreadControlBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePassThreadControlWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePassThreadControlParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePagePassThreadControlBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePassThreadControl(ctx context.Context, client *core.Client, id string, params CreatePagePassThreadControlParams) (*objects.Page, error) {
	out, _, err := CreatePagePassThreadControlWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePersonasParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePersonasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePersonasBatchCall(id string, params GetPagePersonasParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "personas"), params.ToParams(), options...)
}

func NewGetPagePersonasBatchRequest(id string, params GetPagePersonasParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Persona]] {
	return core.NewBatchRequest[core.Cursor[objects.Persona]](GetPagePersonasBatchCall(id, params, options...))
}

func DecodeGetPagePersonasBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Persona], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Persona]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePersonasWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePersonasParams) (*core.Cursor[objects.Persona], *core.Response, error) {
	var out core.Cursor[objects.Persona]
	call := GetPagePersonasBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePersonas(ctx context.Context, client *core.Client, id string, params GetPagePersonasParams) (*core.Cursor[objects.Persona], error) {
	out, _, err := GetPagePersonasWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePersonasParams struct {
	Name              string      `facebook:"name"`
	ProfilePictureURL string      `facebook:"profile_picture_url"`
	Extra             core.Params `facebook:"-"`
}

func (params CreatePagePersonasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	out["profile_picture_url"] = params.ProfilePictureURL
	return out
}

func CreatePagePersonasBatchCall(id string, params CreatePagePersonasParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "personas"), params.ToParams(), options...)
}

func NewCreatePagePersonasBatchRequest(id string, params CreatePagePersonasParams, options ...core.BatchOption) *core.BatchRequest[objects.Persona] {
	return core.NewBatchRequest[objects.Persona](CreatePagePersonasBatchCall(id, params, options...))
}

func DecodeCreatePagePersonasBatchResponse(response *core.BatchResponse) (*objects.Persona, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Persona
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePersonasWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePersonasParams) (*objects.Persona, *core.Response, error) {
	var out objects.Persona
	call := CreatePagePersonasBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePersonas(ctx context.Context, client *core.Client, id string, params CreatePagePersonasParams) (*objects.Persona, error) {
	out, _, err := CreatePagePersonasWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePhotoStoriesParams struct {
	PhotoID *core.ID    `facebook:"photo_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreatePagePhotoStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PhotoID != nil {
		out["photo_id"] = *params.PhotoID
	}
	return out
}

func CreatePagePhotoStoriesBatchCall(id string, params CreatePagePhotoStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "photo_stories"), params.ToParams(), options...)
}

func NewCreatePagePhotoStoriesBatchRequest(id string, params CreatePagePhotoStoriesParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePagePhotoStoriesBatchCall(id, params, options...))
}

func DecodeCreatePagePhotoStoriesBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePhotoStoriesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePhotoStoriesParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePagePhotoStoriesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePhotoStories(ctx context.Context, client *core.Client, id string, params CreatePagePhotoStoriesParams) (*objects.Page, error) {
	out, _, err := CreatePagePhotoStoriesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePhotosParams struct {
	BizTagID   *core.ID                       `facebook:"biz_tag_id"`
	BusinessID *core.ID                       `facebook:"business_id"`
	Type       *enums.PagephotosTypeEnumParam `facebook:"type"`
	Extra      core.Params                    `facebook:"-"`
}

func (params GetPagePhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BizTagID != nil {
		out["biz_tag_id"] = *params.BizTagID
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetPagePhotosBatchCall(id string, params GetPagePhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewGetPagePhotosBatchRequest(id string, params GetPagePhotosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Photo]] {
	return core.NewBatchRequest[core.Cursor[objects.Photo]](GetPagePhotosBatchCall(id, params, options...))
}

func DecodeGetPagePhotosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Photo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Photo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePhotosWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePhotosParams) (*core.Cursor[objects.Photo], *core.Response, error) {
	var out core.Cursor[objects.Photo]
	call := GetPagePhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePhotos(ctx context.Context, client *core.Client, id string, params GetPagePhotosParams) (*core.Cursor[objects.Photo], error) {
	out, _, err := GetPagePhotosWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePhotosParams struct {
	Aid                                   *string                                            `facebook:"aid"`
	AllowSphericalPhoto                   *bool                                              `facebook:"allow_spherical_photo"`
	AltTextCustom                         *string                                            `facebook:"alt_text_custom"`
	AndroidKeyHash                        *string                                            `facebook:"android_key_hash"`
	ApplicationID                         *core.ID                                           `facebook:"application_id"`
	Attempt                               *uint64                                            `facebook:"attempt"`
	AudienceExp                           *bool                                              `facebook:"audience_exp"`
	BackdatedTime                         *core.Time                                         `facebook:"backdated_time"`
	BackdatedTimeGranularity              *enums.PagephotosBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	Caption                               *string                                            `facebook:"caption"`
	ComposerSessionID                     *core.ID                                           `facebook:"composer_session_id"`
	DirectShareStatus                     *uint64                                            `facebook:"direct_share_status"`
	FeedTargeting                         *map[string]interface{}                            `facebook:"feed_targeting"`
	FilterType                            *uint64                                            `facebook:"filter_type"`
	FullResIsComingLater                  *bool                                              `facebook:"full_res_is_coming_later"`
	InitialViewHeadingOverrideDegrees     *uint64                                            `facebook:"initial_view_heading_override_degrees"`
	InitialViewPitchOverrideDegrees       *uint64                                            `facebook:"initial_view_pitch_override_degrees"`
	InitialViewVerticalFovOverrideDegrees *uint64                                            `facebook:"initial_view_vertical_fov_override_degrees"`
	IosBundleID                           *core.ID                                           `facebook:"ios_bundle_id"`
	IsExplicitLocation                    *bool                                              `facebook:"is_explicit_location"`
	IsExplicitPlace                       *bool                                              `facebook:"is_explicit_place"`
	LocationSourceID                      *core.ID                                           `facebook:"location_source_id"`
	ManualPrivacy                         *bool                                              `facebook:"manual_privacy"`
	Message                               *string                                            `facebook:"message"`
	Name                                  *string                                            `facebook:"name"`
	NectarModule                          *string                                            `facebook:"nectar_module"`
	NoStory                               *bool                                              `facebook:"no_story"`
	OfflineID                             *core.ID                                           `facebook:"offline_id"`
	OgActionTypeID                        *core.ID                                           `facebook:"og_action_type_id"`
	OgIconID                              *core.ID                                           `facebook:"og_icon_id"`
	OgObjectID                            *core.ID                                           `facebook:"og_object_id"`
	OgPhrase                              *string                                            `facebook:"og_phrase"`
	OgSetProfileBadge                     *bool                                              `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism                 *string                                            `facebook:"og_suggestion_mechanism"`
	ParentMediaID                         *core.ID                                           `facebook:"parent_media_id"`
	Place                                 *map[string]interface{}                            `facebook:"place"`
	Privacy                               *string                                            `facebook:"privacy"`
	ProfileID                             *core.ID                                           `facebook:"profile_id"`
	ProvenanceInfo                        *map[string]interface{}                            `facebook:"provenance_info"`
	ProxiedAppID                          *core.ID                                           `facebook:"proxied_app_id"`
	Published                             *bool                                              `facebook:"published"`
	Qn                                    *string                                            `facebook:"qn"`
	ScheduledPublishTime                  *uint64                                            `facebook:"scheduled_publish_time"`
	SphericalMetadata                     *map[string]interface{}                            `facebook:"spherical_metadata"`
	SponsorID                             *core.ID                                           `facebook:"sponsor_id"`
	SponsorRelationship                   *uint64                                            `facebook:"sponsor_relationship"`
	Tags                                  *[]map[string]interface{}                          `facebook:"tags"`
	TargetID                              *core.ID                                           `facebook:"target_id"`
	Targeting                             *map[string]interface{}                            `facebook:"targeting"`
	Temporary                             *bool                                              `facebook:"temporary"`
	TimeSinceOriginalPost                 *uint64                                            `facebook:"time_since_original_post"`
	UID                                   *core.ID                                           `facebook:"uid"`
	UnpublishedContentType                *enums.PagephotosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	URL                                   *string                                            `facebook:"url"`
	UserSelectedTags                      *bool                                              `facebook:"user_selected_tags"`
	VaultImageID                          *core.ID                                           `facebook:"vault_image_id"`
	Extra                                 core.Params                                        `facebook:"-"`
}

func (params CreatePagePhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Aid != nil {
		out["aid"] = *params.Aid
	}
	if params.AllowSphericalPhoto != nil {
		out["allow_spherical_photo"] = *params.AllowSphericalPhoto
	}
	if params.AltTextCustom != nil {
		out["alt_text_custom"] = *params.AltTextCustom
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.Attempt != nil {
		out["attempt"] = *params.Attempt
	}
	if params.AudienceExp != nil {
		out["audience_exp"] = *params.AudienceExp
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.FilterType != nil {
		out["filter_type"] = *params.FilterType
	}
	if params.FullResIsComingLater != nil {
		out["full_res_is_coming_later"] = *params.FullResIsComingLater
	}
	if params.InitialViewHeadingOverrideDegrees != nil {
		out["initial_view_heading_override_degrees"] = *params.InitialViewHeadingOverrideDegrees
	}
	if params.InitialViewPitchOverrideDegrees != nil {
		out["initial_view_pitch_override_degrees"] = *params.InitialViewPitchOverrideDegrees
	}
	if params.InitialViewVerticalFovOverrideDegrees != nil {
		out["initial_view_vertical_fov_override_degrees"] = *params.InitialViewVerticalFovOverrideDegrees
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitPlace != nil {
		out["is_explicit_place"] = *params.IsExplicitPlace
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
	}
	if params.OfflineID != nil {
		out["offline_id"] = *params.OfflineID
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
	if params.OgSetProfileBadge != nil {
		out["og_set_profile_badge"] = *params.OgSetProfileBadge
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.ParentMediaID != nil {
		out["parent_media_id"] = *params.ParentMediaID
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.ProfileID != nil {
		out["profile_id"] = *params.ProfileID
	}
	if params.ProvenanceInfo != nil {
		out["provenance_info"] = *params.ProvenanceInfo
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Qn != nil {
		out["qn"] = *params.Qn
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.SphericalMetadata != nil {
		out["spherical_metadata"] = *params.SphericalMetadata
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.Temporary != nil {
		out["temporary"] = *params.Temporary
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VaultImageID != nil {
		out["vault_image_id"] = *params.VaultImageID
	}
	return out
}

func CreatePagePhotosBatchCall(id string, params CreatePagePhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewCreatePagePhotosBatchRequest(id string, params CreatePagePhotosParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](CreatePagePhotosBatchCall(id, params, options...))
}

func DecodeCreatePagePhotosBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Photo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePhotosWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePhotosParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := CreatePagePhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePhotos(ctx context.Context, client *core.Client, id string, params CreatePagePhotosParams) (*objects.Photo, error) {
	out, _, err := CreatePagePhotosWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePictureParams struct {
	Height   *int                            `facebook:"height"`
	Redirect *bool                           `facebook:"redirect"`
	Type     *enums.PagepictureTypeEnumParam `facebook:"type"`
	Width    *int                            `facebook:"width"`
	Extra    core.Params                     `facebook:"-"`
}

func (params GetPagePictureParams) ToParams() core.Params {
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

func GetPagePictureBatchCall(id string, params GetPagePictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetPagePictureBatchRequest(id string, params GetPagePictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetPagePictureBatchCall(id, params, options...))
}

func DecodeGetPagePictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
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

func GetPagePictureWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePictureParams) (*core.Cursor[objects.ProfilePictureSource], *core.Response, error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetPagePictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePicture(ctx context.Context, client *core.Client, id string, params GetPagePictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	out, _, err := GetPagePictureWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePagePictureParams struct {
	AndroidKeyHash            *string                 `facebook:"android_key_hash"`
	BurnMediaEffect           *bool                   `facebook:"burn_media_effect"`
	Caption                   *string                 `facebook:"caption"`
	ComposerSessionID         *core.ID                `facebook:"composer_session_id"`
	FrameEntrypoint           *string                 `facebook:"frame_entrypoint"`
	HasUmg                    *bool                   `facebook:"has_umg"`
	Height                    *uint64                 `facebook:"height"`
	IosBundleID               *core.ID                `facebook:"ios_bundle_id"`
	MediaEffectIds            *[]core.ID              `facebook:"media_effect_ids"`
	MediaEffectSourceObjectID *core.ID                `facebook:"media_effect_source_object_id"`
	MsqrdMaskID               *core.ID                `facebook:"msqrd_mask_id"`
	Photo                     *string                 `facebook:"photo"`
	Picture                   *string                 `facebook:"picture"`
	ProfilePicMethod          *string                 `facebook:"profile_pic_method"`
	ProfilePicSource          *string                 `facebook:"profile_pic_source"`
	ProxiedAppID              *core.ID                `facebook:"proxied_app_id"`
	Qn                        *string                 `facebook:"qn"`
	Reuse                     *bool                   `facebook:"reuse"`
	ScaledCropRect            *map[string]interface{} `facebook:"scaled_crop_rect"`
	SetProfilePhotoShield     *string                 `facebook:"set_profile_photo_shield"`
	StickerID                 *core.ID                `facebook:"sticker_id"`
	StickerSourceObjectID     *core.ID                `facebook:"sticker_source_object_id"`
	SuppressStories           *bool                   `facebook:"suppress_stories"`
	Width                     *uint64                 `facebook:"width"`
	X                         *uint64                 `facebook:"x"`
	Y                         *uint64                 `facebook:"y"`
	Extra                     core.Params             `facebook:"-"`
}

func (params CreatePagePictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.BurnMediaEffect != nil {
		out["burn_media_effect"] = *params.BurnMediaEffect
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.FrameEntrypoint != nil {
		out["frame_entrypoint"] = *params.FrameEntrypoint
	}
	if params.HasUmg != nil {
		out["has_umg"] = *params.HasUmg
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.MediaEffectIds != nil {
		out["media_effect_ids"] = *params.MediaEffectIds
	}
	if params.MediaEffectSourceObjectID != nil {
		out["media_effect_source_object_id"] = *params.MediaEffectSourceObjectID
	}
	if params.MsqrdMaskID != nil {
		out["msqrd_mask_id"] = *params.MsqrdMaskID
	}
	if params.Photo != nil {
		out["photo"] = *params.Photo
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.ProfilePicMethod != nil {
		out["profile_pic_method"] = *params.ProfilePicMethod
	}
	if params.ProfilePicSource != nil {
		out["profile_pic_source"] = *params.ProfilePicSource
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.Qn != nil {
		out["qn"] = *params.Qn
	}
	if params.Reuse != nil {
		out["reuse"] = *params.Reuse
	}
	if params.ScaledCropRect != nil {
		out["scaled_crop_rect"] = *params.ScaledCropRect
	}
	if params.SetProfilePhotoShield != nil {
		out["set_profile_photo_shield"] = *params.SetProfilePhotoShield
	}
	if params.StickerID != nil {
		out["sticker_id"] = *params.StickerID
	}
	if params.StickerSourceObjectID != nil {
		out["sticker_source_object_id"] = *params.StickerSourceObjectID
	}
	if params.SuppressStories != nil {
		out["suppress_stories"] = *params.SuppressStories
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	if params.X != nil {
		out["x"] = *params.X
	}
	if params.Y != nil {
		out["y"] = *params.Y
	}
	return out
}

func CreatePagePictureBatchCall(id string, params CreatePagePictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewCreatePagePictureBatchRequest(id string, params CreatePagePictureParams, options ...core.BatchOption) *core.BatchRequest[objects.ProfilePictureSource] {
	return core.NewBatchRequest[objects.ProfilePictureSource](CreatePagePictureBatchCall(id, params, options...))
}

func DecodeCreatePagePictureBatchResponse(response *core.BatchResponse) (*objects.ProfilePictureSource, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProfilePictureSource
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePictureWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePictureParams) (*objects.ProfilePictureSource, *core.Response, error) {
	var out objects.ProfilePictureSource
	call := CreatePagePictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePicture(ctx context.Context, client *core.Client, id string, params CreatePagePictureParams) (*objects.ProfilePictureSource, error) {
	out, _, err := CreatePagePictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePostsParams struct {
	IncludeHidden *bool                         `facebook:"include_hidden"`
	Limit         *uint64                       `facebook:"limit"`
	Q             *string                       `facebook:"q"`
	ShowExpired   *bool                         `facebook:"show_expired"`
	With          *enums.PagepostsWithEnumParam `facebook:"with"`
	Extra         core.Params                   `facebook:"-"`
}

func (params GetPagePostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.Q != nil {
		out["q"] = *params.Q
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetPagePostsBatchCall(id string, params GetPagePostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "posts"), params.ToParams(), options...)
}

func NewGetPagePostsBatchRequest(id string, params GetPagePostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPagePostsBatchCall(id, params, options...))
}

func DecodeGetPagePostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostsParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPagePostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePosts(ctx context.Context, client *core.Client, id string, params GetPagePostsParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPagePostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageProductCatalogsBatchCall(id string, params GetPageProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_catalogs"), params.ToParams(), options...)
}

func NewGetPageProductCatalogsBatchRequest(id string, params GetPageProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetPageProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetPageProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
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

func GetPageProductCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], *core.Response, error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetPageProductCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageProductCatalogs(ctx context.Context, client *core.Client, id string, params GetPageProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	out, _, err := GetPageProductCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePublishedPostsParams struct {
	IncludeHidden *bool                                  `facebook:"include_hidden"`
	Limit         *uint64                                `facebook:"limit"`
	ShowExpired   *bool                                  `facebook:"show_expired"`
	With          *enums.PagepublishedPostsWithEnumParam `facebook:"with"`
	Extra         core.Params                            `facebook:"-"`
}

func (params GetPagePublishedPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetPagePublishedPostsBatchCall(id string, params GetPagePublishedPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "published_posts"), params.ToParams(), options...)
}

func NewGetPagePublishedPostsBatchRequest(id string, params GetPagePublishedPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPagePublishedPostsBatchCall(id, params, options...))
}

func DecodeGetPagePublishedPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePublishedPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePublishedPostsParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPagePublishedPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePublishedPosts(ctx context.Context, client *core.Client, id string, params GetPagePublishedPostsParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPagePublishedPostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageRatingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageRatingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageRatingsBatchCall(id string, params GetPageRatingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ratings"), params.ToParams(), options...)
}

func NewGetPageRatingsBatchRequest(id string, params GetPageRatingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Recommendation]] {
	return core.NewBatchRequest[core.Cursor[objects.Recommendation]](GetPageRatingsBatchCall(id, params, options...))
}

func DecodeGetPageRatingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Recommendation], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Recommendation]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageRatingsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageRatingsParams) (*core.Cursor[objects.Recommendation], *core.Response, error) {
	var out core.Cursor[objects.Recommendation]
	call := GetPageRatingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageRatings(ctx context.Context, client *core.Client, id string, params GetPageRatingsParams) (*core.Cursor[objects.Recommendation], error) {
	out, _, err := GetPageRatingsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageReleaseThreadControlParams struct {
	Recipient map[string]interface{} `facebook:"recipient"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreatePageReleaseThreadControlParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["recipient"] = params.Recipient
	return out
}

func CreatePageReleaseThreadControlBatchCall(id string, params CreatePageReleaseThreadControlParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "release_thread_control"), params.ToParams(), options...)
}

func NewCreatePageReleaseThreadControlBatchRequest(id string, params CreatePageReleaseThreadControlParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageReleaseThreadControlBatchCall(id, params, options...))
}

func DecodeCreatePageReleaseThreadControlBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageReleaseThreadControlWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageReleaseThreadControlParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageReleaseThreadControlBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageReleaseThreadControl(ctx context.Context, client *core.Client, id string, params CreatePageReleaseThreadControlParams) (*objects.Page, error) {
	out, _, err := CreatePageReleaseThreadControlWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageRequestThreadControlParams struct {
	Metadata  *string                `facebook:"metadata"`
	Recipient map[string]interface{} `facebook:"recipient"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreatePageRequestThreadControlParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	out["recipient"] = params.Recipient
	return out
}

func CreatePageRequestThreadControlBatchCall(id string, params CreatePageRequestThreadControlParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "request_thread_control"), params.ToParams(), options...)
}

func NewCreatePageRequestThreadControlBatchRequest(id string, params CreatePageRequestThreadControlParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageRequestThreadControlBatchCall(id, params, options...))
}

func DecodeCreatePageRequestThreadControlBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageRequestThreadControlWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageRequestThreadControlParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageRequestThreadControlBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageRequestThreadControl(ctx context.Context, client *core.Client, id string, params CreatePageRequestThreadControlParams) (*objects.Page, error) {
	out, _, err := CreatePageRequestThreadControlWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageRolesParams struct {
	IncludeDeactivated *bool       `facebook:"include_deactivated"`
	UID                *core.ID    `facebook:"uid"`
	Extra              core.Params `facebook:"-"`
}

func (params GetPageRolesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeDeactivated != nil {
		out["include_deactivated"] = *params.IncludeDeactivated
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	return out
}

func GetPageRolesBatchCall(id string, params GetPageRolesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "roles"), params.ToParams(), options...)
}

func NewGetPageRolesBatchRequest(id string, params GetPageRolesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.User]] {
	return core.NewBatchRequest[core.Cursor[objects.User]](GetPageRolesBatchCall(id, params, options...))
}

func DecodeGetPageRolesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.User], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.User]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageRolesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageRolesParams) (*core.Cursor[objects.User], *core.Response, error) {
	var out core.Cursor[objects.User]
	call := GetPageRolesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageRoles(ctx context.Context, client *core.Client, id string, params GetPageRolesParams) (*core.Cursor[objects.User], error) {
	out, _, err := GetPageRolesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageRtbDynamicPostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageRtbDynamicPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageRtbDynamicPostsBatchCall(id string, params GetPageRtbDynamicPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "rtb_dynamic_posts"), params.ToParams(), options...)
}

func NewGetPageRtbDynamicPostsBatchRequest(id string, params GetPageRtbDynamicPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.RTBDynamicPost]] {
	return core.NewBatchRequest[core.Cursor[objects.RTBDynamicPost]](GetPageRtbDynamicPostsBatchCall(id, params, options...))
}

func DecodeGetPageRtbDynamicPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.RTBDynamicPost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.RTBDynamicPost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageRtbDynamicPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageRtbDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], *core.Response, error) {
	var out core.Cursor[objects.RTBDynamicPost]
	call := GetPageRtbDynamicPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageRtbDynamicPosts(ctx context.Context, client *core.Client, id string, params GetPageRtbDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], error) {
	out, _, err := GetPageRtbDynamicPostsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageScheduledLiveVideoParams struct {
	EndTime   *uint64                                    `facebook:"end_time"`
	StartTime uint64                                     `facebook:"start_time"`
	State     enums.PagescheduledLiveVideoStateEnumParam `facebook:"state"`
	Video     string                                     `facebook:"video"`
	Extra     core.Params                                `facebook:"-"`
}

func (params CreatePageScheduledLiveVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	out["start_time"] = params.StartTime
	out["state"] = params.State
	out["video"] = params.Video
	return out
}

func CreatePageScheduledLiveVideoBatchCall(id string, params CreatePageScheduledLiveVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "scheduled_live_video"), params.ToParams(), options...)
}

func NewCreatePageScheduledLiveVideoBatchRequest(id string, params CreatePageScheduledLiveVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.ScheduledLiveVideo] {
	return core.NewBatchRequest[objects.ScheduledLiveVideo](CreatePageScheduledLiveVideoBatchCall(id, params, options...))
}

func DecodeCreatePageScheduledLiveVideoBatchResponse(response *core.BatchResponse) (*objects.ScheduledLiveVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ScheduledLiveVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageScheduledLiveVideoWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageScheduledLiveVideoParams) (*objects.ScheduledLiveVideo, *core.Response, error) {
	var out objects.ScheduledLiveVideo
	call := CreatePageScheduledLiveVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageScheduledLiveVideo(ctx context.Context, client *core.Client, id string, params CreatePageScheduledLiveVideoParams) (*objects.ScheduledLiveVideo, error) {
	out, _, err := CreatePageScheduledLiveVideoWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageScheduledPostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageScheduledPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageScheduledPostsBatchCall(id string, params GetPageScheduledPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "scheduled_posts"), params.ToParams(), options...)
}

func NewGetPageScheduledPostsBatchRequest(id string, params GetPageScheduledPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPageScheduledPostsBatchCall(id, params, options...))
}

func DecodeGetPageScheduledPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageScheduledPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageScheduledPostsParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPageScheduledPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageScheduledPosts(ctx context.Context, client *core.Client, id string, params GetPageScheduledPostsParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPageScheduledPostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageSettingsBatchCall(id string, params GetPageSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "settings"), params.ToParams(), options...)
}

func NewGetPageSettingsBatchRequest(id string, params GetPageSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PageSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.PageSettings]](GetPageSettingsBatchCall(id, params, options...))
}

func DecodeGetPageSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PageSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PageSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageSettingsParams) (*core.Cursor[objects.PageSettings], *core.Response, error) {
	var out core.Cursor[objects.PageSettings]
	call := GetPageSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageSettings(ctx context.Context, client *core.Client, id string, params GetPageSettingsParams) (*core.Cursor[objects.PageSettings], error) {
	out, _, err := GetPageSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageSettingsParams struct {
	Option *map[string]interface{} `facebook:"option"`
	Extra  core.Params             `facebook:"-"`
}

func (params CreatePageSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Option != nil {
		out["option"] = *params.Option
	}
	return out
}

func CreatePageSettingsBatchCall(id string, params CreatePageSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "settings"), params.ToParams(), options...)
}

func NewCreatePageSettingsBatchRequest(id string, params CreatePageSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageSettingsBatchCall(id, params, options...))
}

func DecodeCreatePageSettingsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageSettingsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageSettingsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageSettings(ctx context.Context, client *core.Client, id string, params CreatePageSettingsParams) (*objects.Page, error) {
	out, _, err := CreatePageSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageShopSetupStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageShopSetupStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageShopSetupStatusBatchCall(id string, params GetPageShopSetupStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shop_setup_status"), params.ToParams(), options...)
}

func NewGetPageShopSetupStatusBatchRequest(id string, params GetPageShopSetupStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceMerchantSettingsSetupStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceMerchantSettingsSetupStatus]](GetPageShopSetupStatusBatchCall(id, params, options...))
}

func DecodeGetPageShopSetupStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceMerchantSettingsSetupStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageShopSetupStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetPageShopSetupStatusParams) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], *core.Response, error) {
	var out core.Cursor[objects.CommerceMerchantSettingsSetupStatus]
	call := GetPageShopSetupStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageShopSetupStatus(ctx context.Context, client *core.Client, id string, params GetPageShopSetupStatusParams) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], error) {
	out, _, err := GetPageShopSetupStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageSpaceParticipantsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageSpaceParticipantsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageSpaceParticipantsBatchCall(id string, params GetPageSpaceParticipantsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "space_participants"), params.ToParams(), options...)
}

func NewGetPageSpaceParticipantsBatchRequest(id string, params GetPageSpaceParticipantsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPageSpaceParticipantsBatchCall(id, params, options...))
}

func DecodeGetPageSpaceParticipantsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPageSpaceParticipantsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageSpaceParticipantsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPageSpaceParticipantsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageSpaceParticipants(ctx context.Context, client *core.Client, id string, params GetPageSpaceParticipantsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPageSpaceParticipantsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageSpaceParticipantsParams struct {
	Recipient map[string]interface{} `facebook:"recipient"`
	SpaceName string                 `facebook:"space_name"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreatePageSpaceParticipantsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["recipient"] = params.Recipient
	out["space_name"] = params.SpaceName
	return out
}

func CreatePageSpaceParticipantsBatchCall(id string, params CreatePageSpaceParticipantsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "space_participants"), params.ToParams(), options...)
}

func NewCreatePageSpaceParticipantsBatchRequest(id string, params CreatePageSpaceParticipantsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageSpaceParticipantsBatchCall(id, params, options...))
}

func DecodeCreatePageSpaceParticipantsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageSpaceParticipantsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageSpaceParticipantsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageSpaceParticipantsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageSpaceParticipants(ctx context.Context, client *core.Client, id string, params CreatePageSpaceParticipantsParams) (*objects.Page, error) {
	out, _, err := CreatePageSpaceParticipantsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageStoreLocationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageStoreLocationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageStoreLocationsBatchCall(id string, params GetPageStoreLocationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "store_locations"), params.ToParams(), options...)
}

func NewGetPageStoreLocationsBatchRequest(id string, params GetPageStoreLocationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.StoreLocation]] {
	return core.NewBatchRequest[core.Cursor[objects.StoreLocation]](GetPageStoreLocationsBatchCall(id, params, options...))
}

func DecodeGetPageStoreLocationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.StoreLocation], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.StoreLocation]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageStoreLocationsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageStoreLocationsParams) (*core.Cursor[objects.StoreLocation], *core.Response, error) {
	var out core.Cursor[objects.StoreLocation]
	call := GetPageStoreLocationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageStoreLocations(ctx context.Context, client *core.Client, id string, params GetPageStoreLocationsParams) (*core.Cursor[objects.StoreLocation], error) {
	out, _, err := GetPageStoreLocationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageStoriesParams struct {
	Since  *core.Time                          `facebook:"since"`
	Status *[]enums.PagestoriesStatusEnumParam `facebook:"status"`
	Until  *core.Time                          `facebook:"until"`
	Extra  core.Params                         `facebook:"-"`
}

func (params GetPageStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPageStoriesBatchCall(id string, params GetPageStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stories"), params.ToParams(), options...)
}

func NewGetPageStoriesBatchRequest(id string, params GetPageStoriesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Stories]] {
	return core.NewBatchRequest[core.Cursor[objects.Stories]](GetPageStoriesBatchCall(id, params, options...))
}

func DecodeGetPageStoriesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Stories], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Stories]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageStoriesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageStoriesParams) (*core.Cursor[objects.Stories], *core.Response, error) {
	var out core.Cursor[objects.Stories]
	call := GetPageStoriesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageStories(ctx context.Context, client *core.Client, id string, params GetPageStoriesParams) (*core.Cursor[objects.Stories], error) {
	out, _, err := GetPageStoriesWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePageSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePageSubscribedAppsBatchCall(id string, params DeletePageSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewDeletePageSubscribedAppsBatchRequest(id string, params DeletePageSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageSubscribedAppsBatchCall(id, params, options...))
}

func DecodeDeletePageSubscribedAppsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageSubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageSubscribedAppsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageSubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageSubscribedApps(ctx context.Context, client *core.Client, id string, params DeletePageSubscribedAppsParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageSubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageSubscribedAppsBatchCall(id string, params GetPageSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewGetPageSubscribedAppsBatchRequest(id string, params GetPageSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetPageSubscribedAppsBatchCall(id, params, options...))
}

func DecodeGetPageSubscribedAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
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

func GetPageSubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageSubscribedAppsParams) (*core.Cursor[objects.Application], *core.Response, error) {
	var out core.Cursor[objects.Application]
	call := GetPageSubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageSubscribedApps(ctx context.Context, client *core.Client, id string, params GetPageSubscribedAppsParams) (*core.Cursor[objects.Application], error) {
	out, _, err := GetPageSubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageSubscribedAppsParams struct {
	SubscribedFields []enums.PagesubscribedAppsSubscribedFieldsEnumParam `facebook:"subscribed_fields"`
	Extra            core.Params                                         `facebook:"-"`
}

func (params CreatePageSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["subscribed_fields"] = params.SubscribedFields
	return out
}

func CreatePageSubscribedAppsBatchCall(id string, params CreatePageSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewCreatePageSubscribedAppsBatchRequest(id string, params CreatePageSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageSubscribedAppsBatchCall(id, params, options...))
}

func DecodeCreatePageSubscribedAppsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageSubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageSubscribedAppsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageSubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageSubscribedApps(ctx context.Context, client *core.Client, id string, params CreatePageSubscribedAppsParams) (*objects.Page, error) {
	out, _, err := CreatePageSubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageTabsParams struct {
	Tab   *[]string   `facebook:"tab"`
	Extra core.Params `facebook:"-"`
}

func (params GetPageTabsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Tab != nil {
		out["tab"] = *params.Tab
	}
	return out
}

func GetPageTabsBatchCall(id string, params GetPageTabsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tabs"), params.ToParams(), options...)
}

func NewGetPageTabsBatchRequest(id string, params GetPageTabsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Tab]] {
	return core.NewBatchRequest[core.Cursor[objects.Tab]](GetPageTabsBatchCall(id, params, options...))
}

func DecodeGetPageTabsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Tab], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Tab]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageTabsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageTabsParams) (*core.Cursor[objects.Tab], *core.Response, error) {
	var out core.Cursor[objects.Tab]
	call := GetPageTabsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageTabs(ctx context.Context, client *core.Client, id string, params GetPageTabsParams) (*core.Cursor[objects.Tab], error) {
	out, _, err := GetPageTabsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageTaggedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageTaggedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageTaggedBatchCall(id string, params GetPageTaggedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tagged"), params.ToParams(), options...)
}

func NewGetPageTaggedBatchRequest(id string, params GetPageTaggedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPageTaggedBatchCall(id, params, options...))
}

func DecodeGetPageTaggedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageTaggedWithResponse(ctx context.Context, client *core.Client, id string, params GetPageTaggedParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPageTaggedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageTagged(ctx context.Context, client *core.Client, id string, params GetPageTaggedParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPageTaggedWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageTakeThreadControlParams struct {
	Metadata  *string                `facebook:"metadata"`
	Recipient map[string]interface{} `facebook:"recipient"`
	Extra     core.Params            `facebook:"-"`
}

func (params CreatePageTakeThreadControlParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	out["recipient"] = params.Recipient
	return out
}

func CreatePageTakeThreadControlBatchCall(id string, params CreatePageTakeThreadControlParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "take_thread_control"), params.ToParams(), options...)
}

func NewCreatePageTakeThreadControlBatchRequest(id string, params CreatePageTakeThreadControlParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageTakeThreadControlBatchCall(id, params, options...))
}

func DecodeCreatePageTakeThreadControlBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageTakeThreadControlWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageTakeThreadControlParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageTakeThreadControlBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageTakeThreadControl(ctx context.Context, client *core.Client, id string, params CreatePageTakeThreadControlParams) (*objects.Page, error) {
	out, _, err := CreatePageTakeThreadControlWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageThreadOwnerParams struct {
	Recipient string      `facebook:"recipient"`
	Extra     core.Params `facebook:"-"`
}

func (params GetPageThreadOwnerParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["recipient"] = params.Recipient
	return out
}

func GetPageThreadOwnerBatchCall(id string, params GetPageThreadOwnerParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "thread_owner"), params.ToParams(), options...)
}

func NewGetPageThreadOwnerBatchRequest(id string, params GetPageThreadOwnerParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PageThreadOwner]] {
	return core.NewBatchRequest[core.Cursor[objects.PageThreadOwner]](GetPageThreadOwnerBatchCall(id, params, options...))
}

func DecodeGetPageThreadOwnerBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PageThreadOwner], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PageThreadOwner]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageThreadOwnerWithResponse(ctx context.Context, client *core.Client, id string, params GetPageThreadOwnerParams) (*core.Cursor[objects.PageThreadOwner], *core.Response, error) {
	var out core.Cursor[objects.PageThreadOwner]
	call := GetPageThreadOwnerBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageThreadOwner(ctx context.Context, client *core.Client, id string, params GetPageThreadOwnerParams) (*core.Cursor[objects.PageThreadOwner], error) {
	out, _, err := GetPageThreadOwnerWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageThreadsParams struct {
	Folder   *string                             `facebook:"folder"`
	Platform *enums.PagethreadsPlatformEnumParam `facebook:"platform"`
	Tags     *[]string                           `facebook:"tags"`
	UserID   *core.ID                            `facebook:"user_id"`
	Extra    core.Params                         `facebook:"-"`
}

func (params GetPageThreadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	return out
}

func GetPageThreadsBatchCall(id string, params GetPageThreadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "threads"), params.ToParams(), options...)
}

func NewGetPageThreadsBatchRequest(id string, params GetPageThreadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UnifiedThread]] {
	return core.NewBatchRequest[core.Cursor[objects.UnifiedThread]](GetPageThreadsBatchCall(id, params, options...))
}

func DecodeGetPageThreadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UnifiedThread], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UnifiedThread]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageThreadsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageThreadsParams) (*core.Cursor[objects.UnifiedThread], *core.Response, error) {
	var out core.Cursor[objects.UnifiedThread]
	call := GetPageThreadsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageThreads(ctx context.Context, client *core.Client, id string, params GetPageThreadsParams) (*core.Cursor[objects.UnifiedThread], error) {
	out, _, err := GetPageThreadsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageUnlinkAccountsParams struct {
	Psid  string      `facebook:"psid"`
	Extra core.Params `facebook:"-"`
}

func (params CreatePageUnlinkAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["psid"] = params.Psid
	return out
}

func CreatePageUnlinkAccountsBatchCall(id string, params CreatePageUnlinkAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "unlink_accounts"), params.ToParams(), options...)
}

func NewCreatePageUnlinkAccountsBatchRequest(id string, params CreatePageUnlinkAccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](CreatePageUnlinkAccountsBatchCall(id, params, options...))
}

func DecodeCreatePageUnlinkAccountsBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageUnlinkAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageUnlinkAccountsParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := CreatePageUnlinkAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageUnlinkAccounts(ctx context.Context, client *core.Client, id string, params CreatePageUnlinkAccountsParams) (*objects.Page, error) {
	out, _, err := CreatePageUnlinkAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageVideoCopyrightRulesParams struct {
	SelectedRuleID *core.ID                                      `facebook:"selected_rule_id"`
	Source         *enums.PagevideoCopyrightRulesSourceEnumParam `facebook:"source"`
	Extra          core.Params                                   `facebook:"-"`
}

func (params GetPageVideoCopyrightRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SelectedRuleID != nil {
		out["selected_rule_id"] = *params.SelectedRuleID
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetPageVideoCopyrightRulesBatchCall(id string, params GetPageVideoCopyrightRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_copyright_rules"), params.ToParams(), options...)
}

func NewGetPageVideoCopyrightRulesBatchRequest(id string, params GetPageVideoCopyrightRulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoCopyrightRule]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoCopyrightRule]](GetPageVideoCopyrightRulesBatchCall(id, params, options...))
}

func DecodeGetPageVideoCopyrightRulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoCopyrightRule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoCopyrightRule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageVideoCopyrightRulesWithResponse(ctx context.Context, client *core.Client, id string, params GetPageVideoCopyrightRulesParams) (*core.Cursor[objects.VideoCopyrightRule], *core.Response, error) {
	var out core.Cursor[objects.VideoCopyrightRule]
	call := GetPageVideoCopyrightRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageVideoCopyrightRules(ctx context.Context, client *core.Client, id string, params GetPageVideoCopyrightRulesParams) (*core.Cursor[objects.VideoCopyrightRule], error) {
	out, _, err := GetPageVideoCopyrightRulesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageVideoCopyrightRulesParams struct {
	ConditionGroups []map[string]interface{} `facebook:"condition_groups"`
	Name            string                   `facebook:"name"`
	Extra           core.Params              `facebook:"-"`
}

func (params CreatePageVideoCopyrightRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["condition_groups"] = params.ConditionGroups
	out["name"] = params.Name
	return out
}

func CreatePageVideoCopyrightRulesBatchCall(id string, params CreatePageVideoCopyrightRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "video_copyright_rules"), params.ToParams(), options...)
}

func NewCreatePageVideoCopyrightRulesBatchRequest(id string, params CreatePageVideoCopyrightRulesParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyrightRule] {
	return core.NewBatchRequest[objects.VideoCopyrightRule](CreatePageVideoCopyrightRulesBatchCall(id, params, options...))
}

func DecodeCreatePageVideoCopyrightRulesBatchResponse(response *core.BatchResponse) (*objects.VideoCopyrightRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyrightRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageVideoCopyrightRulesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageVideoCopyrightRulesParams) (*objects.VideoCopyrightRule, *core.Response, error) {
	var out objects.VideoCopyrightRule
	call := CreatePageVideoCopyrightRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageVideoCopyrightRules(ctx context.Context, client *core.Client, id string, params CreatePageVideoCopyrightRulesParams) (*objects.VideoCopyrightRule, error) {
	out, _, err := CreatePageVideoCopyrightRulesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageVideoCopyrightsParams struct {
	AttributionID              *core.ID                                           `facebook:"attribution_id"`
	ContentCategory            *enums.PagevideoCopyrightsContentCategoryEnumParam `facebook:"content_category"`
	CopyrightContentID         core.ID                                            `facebook:"copyright_content_id"`
	ExcludedOwnershipCountries *[]string                                          `facebook:"excluded_ownership_countries"`
	ExcludedOwnershipSegments  *[]map[string]interface{}                          `facebook:"excluded_ownership_segments"`
	IsReferenceDisabled        *bool                                              `facebook:"is_reference_disabled"`
	IsReferenceVideo           *bool                                              `facebook:"is_reference_video"`
	MonitoringType             *enums.PagevideoCopyrightsMonitoringTypeEnumParam  `facebook:"monitoring_type"`
	OwnershipCountries         *[]string                                          `facebook:"ownership_countries"`
	RuleID                     *core.ID                                           `facebook:"rule_id"`
	Tags                       *[]string                                          `facebook:"tags"`
	WhitelistedIds             *[]core.ID                                         `facebook:"whitelisted_ids"`
	WhitelistedIgUserIds       *[]core.ID                                         `facebook:"whitelisted_ig_user_ids"`
	Extra                      core.Params                                        `facebook:"-"`
}

func (params CreatePageVideoCopyrightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttributionID != nil {
		out["attribution_id"] = *params.AttributionID
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	out["copyright_content_id"] = params.CopyrightContentID
	if params.ExcludedOwnershipCountries != nil {
		out["excluded_ownership_countries"] = *params.ExcludedOwnershipCountries
	}
	if params.ExcludedOwnershipSegments != nil {
		out["excluded_ownership_segments"] = *params.ExcludedOwnershipSegments
	}
	if params.IsReferenceDisabled != nil {
		out["is_reference_disabled"] = *params.IsReferenceDisabled
	}
	if params.IsReferenceVideo != nil {
		out["is_reference_video"] = *params.IsReferenceVideo
	}
	if params.MonitoringType != nil {
		out["monitoring_type"] = *params.MonitoringType
	}
	if params.OwnershipCountries != nil {
		out["ownership_countries"] = *params.OwnershipCountries
	}
	if params.RuleID != nil {
		out["rule_id"] = *params.RuleID
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.WhitelistedIds != nil {
		out["whitelisted_ids"] = *params.WhitelistedIds
	}
	if params.WhitelistedIgUserIds != nil {
		out["whitelisted_ig_user_ids"] = *params.WhitelistedIgUserIds
	}
	return out
}

func CreatePageVideoCopyrightsBatchCall(id string, params CreatePageVideoCopyrightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "video_copyrights"), params.ToParams(), options...)
}

func NewCreatePageVideoCopyrightsBatchRequest(id string, params CreatePageVideoCopyrightsParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyright] {
	return core.NewBatchRequest[objects.VideoCopyright](CreatePageVideoCopyrightsBatchCall(id, params, options...))
}

func DecodeCreatePageVideoCopyrightsBatchResponse(response *core.BatchResponse) (*objects.VideoCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePageVideoCopyrightsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageVideoCopyrightsParams) (*objects.VideoCopyright, *core.Response, error) {
	var out objects.VideoCopyright
	call := CreatePageVideoCopyrightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageVideoCopyrights(ctx context.Context, client *core.Client, id string, params CreatePageVideoCopyrightsParams) (*objects.VideoCopyright, error) {
	out, _, err := CreatePageVideoCopyrightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageVideoListsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageVideoListsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageVideoListsBatchCall(id string, params GetPageVideoListsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_lists"), params.ToParams(), options...)
}

func NewGetPageVideoListsBatchRequest(id string, params GetPageVideoListsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoList]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoList]](GetPageVideoListsBatchCall(id, params, options...))
}

func DecodeGetPageVideoListsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoList], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoList]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageVideoListsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageVideoListsParams) (*core.Cursor[objects.VideoList], *core.Response, error) {
	var out core.Cursor[objects.VideoList]
	call := GetPageVideoListsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageVideoLists(ctx context.Context, client *core.Client, id string, params GetPageVideoListsParams) (*core.Cursor[objects.VideoList], error) {
	out, _, err := GetPageVideoListsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageVideoReelsParams struct {
	Since *core.Time  `facebook:"since"`
	Until *core.Time  `facebook:"until"`
	Extra core.Params `facebook:"-"`
}

func (params GetPageVideoReelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPageVideoReelsBatchCall(id string, params GetPageVideoReelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_reels"), params.ToParams(), options...)
}

func NewGetPageVideoReelsBatchRequest(id string, params GetPageVideoReelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetPageVideoReelsBatchCall(id, params, options...))
}

func DecodeGetPageVideoReelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageVideoReelsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageVideoReelsParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetPageVideoReelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageVideoReels(ctx context.Context, client *core.Client, id string, params GetPageVideoReelsParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetPageVideoReelsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageVideoReelsParams struct {
	Description          *string                                  `facebook:"description"`
	FeedTargeting        *map[string]interface{}                  `facebook:"feed_targeting"`
	Place                *string                                  `facebook:"place"`
	ScheduledPublishTime *core.Time                               `facebook:"scheduled_publish_time"`
	Targeting            *map[string]interface{}                  `facebook:"targeting"`
	Title                *string                                  `facebook:"title"`
	UploadPhase          enums.PagevideoReelsUploadPhaseEnumParam `facebook:"upload_phase"`
	VideoID              *core.ID                                 `facebook:"video_id"`
	VideoState           *enums.PagevideoReelsVideoStateEnumParam `facebook:"video_state"`
	Extra                core.Params                              `facebook:"-"`
}

func (params CreatePageVideoReelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	out["upload_phase"] = params.UploadPhase
	if params.VideoID != nil {
		out["video_id"] = *params.VideoID
	}
	if params.VideoState != nil {
		out["video_state"] = *params.VideoState
	}
	return out
}

func CreatePageVideoReelsBatchCall(id string, params CreatePageVideoReelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "video_reels"), params.ToParams(), options...)
}

func NewCreatePageVideoReelsBatchRequest(id string, params CreatePageVideoReelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreatePageVideoReelsBatchCall(id, params, options...))
}

func DecodeCreatePageVideoReelsBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
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

func CreatePageVideoReelsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageVideoReelsParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreatePageVideoReelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageVideoReels(ctx context.Context, client *core.Client, id string, params CreatePageVideoReelsParams) (*objects.AdVideo, error) {
	out, _, err := CreatePageVideoReelsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageVideoStoriesParams struct {
	Description          *string                                    `facebook:"description"`
	FeedTargeting        *map[string]interface{}                    `facebook:"feed_targeting"`
	Place                *string                                    `facebook:"place"`
	ScheduledPublishTime *core.Time                                 `facebook:"scheduled_publish_time"`
	Targeting            *map[string]interface{}                    `facebook:"targeting"`
	Title                *string                                    `facebook:"title"`
	UploadPhase          enums.PagevideoStoriesUploadPhaseEnumParam `facebook:"upload_phase"`
	VideoID              *core.ID                                   `facebook:"video_id"`
	VideoState           *enums.PagevideoStoriesVideoStateEnumParam `facebook:"video_state"`
	Extra                core.Params                                `facebook:"-"`
}

func (params CreatePageVideoStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	out["upload_phase"] = params.UploadPhase
	if params.VideoID != nil {
		out["video_id"] = *params.VideoID
	}
	if params.VideoState != nil {
		out["video_state"] = *params.VideoState
	}
	return out
}

func CreatePageVideoStoriesBatchCall(id string, params CreatePageVideoStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "video_stories"), params.ToParams(), options...)
}

func NewCreatePageVideoStoriesBatchRequest(id string, params CreatePageVideoStoriesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageVideoStoriesBatchCall(id, params, options...))
}

func DecodeCreatePageVideoStoriesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageVideoStoriesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageVideoStoriesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageVideoStoriesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageVideoStories(ctx context.Context, client *core.Client, id string, params CreatePageVideoStoriesParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageVideoStoriesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageVideosParams struct {
	Type  *enums.PagevideosTypeEnumParam `facebook:"type"`
	Extra core.Params                    `facebook:"-"`
}

func (params GetPageVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetPageVideosBatchCall(id string, params GetPageVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewGetPageVideosBatchRequest(id string, params GetPageVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetPageVideosBatchCall(id, params, options...))
}

func DecodeGetPageVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetPageVideosParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetPageVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageVideos(ctx context.Context, client *core.Client, id string, params GetPageVideosParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetPageVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageVideosParams struct {
	AdBreaks                      *[]interface{}                                   `facebook:"ad_breaks"`
	ApplicationID                 *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                          `facebook:"audio_story_wave_animation_handle"`
	BackdatedPost                 *[]interface{}                                   `facebook:"backdated_post"`
	CallToAction                  *map[string]interface{}                          `facebook:"call_to_action"`
	ComposerEntryPicker           *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                          `facebook:"composer_source_surface"`
	ComposerType                  *string                                          `facebook:"composer_type"`
	ContainerType                 *enums.PagevideosContainerTypeEnumParam          `facebook:"container_type"`
	ContentCategory               *enums.PagevideosContentCategoryEnumParam        `facebook:"content_category"`
	ContentTags                   *[]string                                        `facebook:"content_tags"`
	CreativeTools                 *string                                          `facebook:"creative_tools"`
	CrosspostedVideoID            *core.ID                                         `facebook:"crossposted_video_id"`
	CustomLabels                  *[]string                                        `facebook:"custom_labels"`
	Description                   *string                                          `facebook:"description"`
	DirectShareStatus             *uint64                                          `facebook:"direct_share_status"`
	EditDescriptionSpec           *map[string]interface{}                          `facebook:"edit_description_spec"`
	Embeddable                    *bool                                            `facebook:"embeddable"`
	EndOffset                     *uint64                                          `facebook:"end_offset"`
	Expiration                    *map[string]interface{}                          `facebook:"expiration"`
	FbuploaderVideoFileChunk      *string                                          `facebook:"fbuploader_video_file_chunk"`
	FeedTargeting                 *map[string]interface{}                          `facebook:"feed_targeting"`
	FileSize                      *uint64                                          `facebook:"file_size"`
	FileURL                       *string                                          `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                            `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.PagevideosFormattingEnumParam             `facebook:"formatting"`
	Fov                           *uint64                                          `facebook:"fov"`
	FrontZRotation                *float64                                         `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                      `facebook:"guide"`
	GuideEnabled                  *bool                                            `facebook:"guide_enabled"`
	InitialHeading                *uint64                                          `facebook:"initial_heading"`
	InitialPitch                  *uint64                                          `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                          `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                            `facebook:"is_boost_intended"`
	IsExplicitShare               *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost            *bool                                            `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                            `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                            `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy                 *bool                                            `facebook:"manual_privacy"`
	MultilingualData              *[]map[string]interface{}                        `facebook:"multilingual_data"`
	NoStory                       *bool                                            `facebook:"no_story"`
	OgActionTypeID                *core.ID                                         `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                      *string                                          `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                          `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                          `facebook:"original_fov"`
	OriginalProjectionType        *enums.PagevideosOriginalProjectionTypeEnumParam `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                          `facebook:"partnership_ad_ad_code"`
	PublishEventID                *core.ID                                         `facebook:"publish_event_id"`
	Published                     *bool                                            `facebook:"published"`
	ReferenceOnly                 *bool                                            `facebook:"reference_only"`
	ReferencedStickerID           *core.ID                                         `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                         `facebook:"replace_video_id"`
	ScheduledPublishTime          *uint64                                          `facebook:"scheduled_publish_time"`
	Secret                        *bool                                            `facebook:"secret"`
	SelectedAudioSpec             *map[string]interface{}                          `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                          `facebook:"slideshow_spec"`
	SocialActions                 *bool                                            `facebook:"social_actions"`
	Source                        *string                                          `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                         `facebook:"source_instagram_media_id"`
	SpecifiedDialect              *string                                          `facebook:"specified_dialect"`
	Spherical                     *bool                                            `facebook:"spherical"`
	SponsorID                     *core.ID                                         `facebook:"sponsor_id"`
	SponsorRelationship           *uint64                                          `facebook:"sponsor_relationship"`
	StartOffset                   *uint64                                          `facebook:"start_offset"`
	SwapMode                      *enums.PagevideosSwapModeEnumParam               `facebook:"swap_mode"`
	Targeting                     *map[string]interface{}                          `facebook:"targeting"`
	TextFormatMetadata            *string                                          `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                  `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                          `facebook:"time_since_original_post"`
	Title                         *string                                          `facebook:"title"`
	TranscodeSettingProperties    *string                                          `facebook:"transcode_setting_properties"`
	UniversalVideoID              *core.ID                                         `facebook:"universal_video_id"`
	UnpublishedContentType        *enums.PagevideosUnpublishedContentTypeEnumParam `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.PagevideosUploadPhaseEnumParam            `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                         `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                          `facebook:"upload_setting_properties"`
	VideoAssetID                  *core.ID                                         `facebook:"video_asset_id"`
	VideoFileChunk                *string                                          `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                          `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                          `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                         `facebook:"waterfall_id"`
	Extra                         core.Params                                      `facebook:"-"`
}

func (params CreatePageVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdBreaks != nil {
		out["ad_breaks"] = *params.AdBreaks
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
	if params.BackdatedPost != nil {
		out["backdated_post"] = *params.BackdatedPost
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
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
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.CreativeTools != nil {
		out["creative_tools"] = *params.CreativeTools
	}
	if params.CrosspostedVideoID != nil {
		out["crossposted_video_id"] = *params.CrosspostedVideoID
	}
	if params.CustomLabels != nil {
		out["custom_labels"] = *params.CustomLabels
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
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
	if params.Expiration != nil {
		out["expiration"] = *params.Expiration
	}
	if params.FbuploaderVideoFileChunk != nil {
		out["fbuploader_video_file_chunk"] = *params.FbuploaderVideoFileChunk
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
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
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
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
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.MultilingualData != nil {
		out["multilingual_data"] = *params.MultilingualData
	}
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
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
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ReferenceOnly != nil {
		out["reference_only"] = *params.ReferenceOnly
	}
	if params.ReferencedStickerID != nil {
		out["referenced_sticker_id"] = *params.ReferencedStickerID
	}
	if params.ReplaceVideoID != nil {
		out["replace_video_id"] = *params.ReplaceVideoID
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Secret != nil {
		out["secret"] = *params.Secret
	}
	if params.SelectedAudioSpec != nil {
		out["selected_audio_spec"] = *params.SelectedAudioSpec
	}
	if params.SlideshowSpec != nil {
		out["slideshow_spec"] = *params.SlideshowSpec
	}
	if params.SocialActions != nil {
		out["social_actions"] = *params.SocialActions
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SourceInstagramMediaID != nil {
		out["source_instagram_media_id"] = *params.SourceInstagramMediaID
	}
	if params.SpecifiedDialect != nil {
		out["specified_dialect"] = *params.SpecifiedDialect
	}
	if params.Spherical != nil {
		out["spherical"] = *params.Spherical
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.StartOffset != nil {
		out["start_offset"] = *params.StartOffset
	}
	if params.SwapMode != nil {
		out["swap_mode"] = *params.SwapMode
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
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
	if params.UniversalVideoID != nil {
		out["universal_video_id"] = *params.UniversalVideoID
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
	if params.VideoAssetID != nil {
		out["video_asset_id"] = *params.VideoAssetID
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

func CreatePageVideosBatchCall(id string, params CreatePageVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewCreatePageVideosBatchRequest(id string, params CreatePageVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreatePageVideosBatchCall(id, params, options...))
}

func DecodeCreatePageVideosBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
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

func CreatePageVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageVideosParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreatePageVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageVideos(ctx context.Context, client *core.Client, id string, params CreatePageVideosParams) (*objects.AdVideo, error) {
	out, _, err := CreatePageVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageVisitorPostsParams struct {
	IncludeHidden *bool                                `facebook:"include_hidden"`
	Limit         *uint64                              `facebook:"limit"`
	ShowExpired   *bool                                `facebook:"show_expired"`
	With          *enums.PagevisitorPostsWithEnumParam `facebook:"with"`
	Extra         core.Params                          `facebook:"-"`
}

func (params GetPageVisitorPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetPageVisitorPostsBatchCall(id string, params GetPageVisitorPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "visitor_posts"), params.ToParams(), options...)
}

func NewGetPageVisitorPostsBatchRequest(id string, params GetPageVisitorPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePost]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePost]](GetPageVisitorPostsBatchCall(id, params, options...))
}

func DecodeGetPageVisitorPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageVisitorPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageVisitorPostsParams) (*core.Cursor[objects.PagePost], *core.Response, error) {
	var out core.Cursor[objects.PagePost]
	call := GetPageVisitorPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageVisitorPosts(ctx context.Context, client *core.Client, id string, params GetPageVisitorPostsParams) (*core.Cursor[objects.PagePost], error) {
	out, _, err := GetPageVisitorPostsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePageWelcomeMessageFlowsParams struct {
	FlowID core.ID     `facebook:"flow_id"`
	Extra  core.Params `facebook:"-"`
}

func (params DeletePageWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["flow_id"] = params.FlowID
	return out
}

func DeletePageWelcomeMessageFlowsBatchCall(id string, params DeletePageWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewDeletePageWelcomeMessageFlowsBatchRequest(id string, params DeletePageWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePageWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeDeletePageWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePageWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params DeletePageWelcomeMessageFlowsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePageWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePageWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params DeletePageWelcomeMessageFlowsParams) (*map[string]interface{}, error) {
	out, _, err := DeletePageWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageWelcomeMessageFlowsParams struct {
	AppID  *core.ID    `facebook:"app_id"`
	FlowID *core.ID    `facebook:"flow_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetPageWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.FlowID != nil {
		out["flow_id"] = *params.FlowID
	}
	return out
}

func GetPageWelcomeMessageFlowsBatchCall(id string, params GetPageWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewGetPageWelcomeMessageFlowsBatchRequest(id string, params GetPageWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]] {
	return core.NewBatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]](GetPageWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeGetPageWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
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

func GetPageWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params GetPageWelcomeMessageFlowsParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], *core.Response, error) {
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	call := GetPageWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params GetPageWelcomeMessageFlowsParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	out, _, err := GetPageWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePageWelcomeMessageFlowsParams struct {
	EligiblePlatforms  *[]enums.PagewelcomeMessageFlowsEligiblePlatformsEnumParam `facebook:"eligible_platforms"`
	FlowID             *core.ID                                                   `facebook:"flow_id"`
	Name               *string                                                    `facebook:"name"`
	WelcomeMessageFlow *[]map[string]interface{}                                  `facebook:"welcome_message_flow"`
	Extra              core.Params                                                `facebook:"-"`
}

func (params CreatePageWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EligiblePlatforms != nil {
		out["eligible_platforms"] = *params.EligiblePlatforms
	}
	if params.FlowID != nil {
		out["flow_id"] = *params.FlowID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.WelcomeMessageFlow != nil {
		out["welcome_message_flow"] = *params.WelcomeMessageFlow
	}
	return out
}

func CreatePageWelcomeMessageFlowsBatchCall(id string, params CreatePageWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewCreatePageWelcomeMessageFlowsBatchRequest(id string, params CreatePageWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePageWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeCreatePageWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePageWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePageWelcomeMessageFlowsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreatePageWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePageWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params CreatePageWelcomeMessageFlowsParams) (*map[string]interface{}, error) {
	out, _, err := CreatePageWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPageParams struct {
	AccountLinkingToken *string     `facebook:"account_linking_token"`
	Extra               core.Params `facebook:"-"`
}

func (params GetPageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountLinkingToken != nil {
		out["account_linking_token"] = *params.AccountLinkingToken
	}
	return out
}

func GetPageBatchCall(id string, params GetPageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPageBatchRequest(id string, params GetPageParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](GetPageBatchCall(id, params, options...))
}

func DecodeGetPageBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageWithResponse(ctx context.Context, client *core.Client, id string, params GetPageParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := GetPageBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPage(ctx context.Context, client *core.Client, id string, params GetPageParams) (*objects.Page, error) {
	out, _, err := GetPageWithResponse(ctx, client, id, params)
	return out, err
}

type UpdatePageParams struct {
	About                        *string                        `facebook:"about"`
	AcceptCrosspostingHandshake  *[]map[string]interface{}      `facebook:"accept_crossposting_handshake"`
	AllowSphericalPhoto          *bool                          `facebook:"allow_spherical_photo"`
	Attire                       *enums.PageAttire              `facebook:"attire"`
	BeginCrosspostingHandshake   *[]map[string]interface{}      `facebook:"begin_crossposting_handshake"`
	Bio                          *string                        `facebook:"bio"`
	Caption                      *string                        `facebook:"caption"`
	CategoryList                 *[]string                      `facebook:"category_list"`
	CompanyOverview              *string                        `facebook:"company_overview"`
	ContactAddress               *map[string]interface{}        `facebook:"contact_address"`
	Cover                        *string                        `facebook:"cover"`
	CulinaryTeam                 *string                        `facebook:"culinary_team"`
	DeliveryAndPickupOptionInfo  *[]string                      `facebook:"delivery_and_pickup_option_info"`
	Description                  *string                        `facebook:"description"`
	DifferentlyOpenOfferings     *map[string]interface{}        `facebook:"differently_open_offerings"`
	DirectedBy                   *string                        `facebook:"directed_by"`
	DisplayedMessageResponseTime *string                        `facebook:"displayed_message_response_time"`
	Emails                       *[]string                      `facebook:"emails"`
	FocusX                       *float64                       `facebook:"focus_x"`
	FocusY                       *float64                       `facebook:"focus_y"`
	FoodStyles                   *[]enums.PageFoodStyles        `facebook:"food_styles"`
	GenAiProvenanceType          *enums.PageGenAiProvenanceType `facebook:"gen_ai_provenance_type"`
	GeneralInfo                  *string                        `facebook:"general_info"`
	GeneralManager               *string                        `facebook:"general_manager"`
	Genre                        *string                        `facebook:"genre"`
	Hours                        *map[string]interface{}        `facebook:"hours"`
	IgnoreCoordinateWarnings     *bool                          `facebook:"ignore_coordinate_warnings"`
	Impressum                    *string                        `facebook:"impressum"`
	IsAlwaysOpen                 *bool                          `facebook:"is_always_open"`
	IsPermanentlyClosed          *bool                          `facebook:"is_permanently_closed"`
	IsPublished                  *bool                          `facebook:"is_published"`
	IsWebhooksSubscribed         *bool                          `facebook:"is_webhooks_subscribed"`
	Location                     *map[string]interface{}        `facebook:"location"`
	Menu                         *string                        `facebook:"menu"`
	Mission                      *string                        `facebook:"mission"`
	NoFeedStory                  *bool                          `facebook:"no_feed_story"`
	NoNotification               *bool                          `facebook:"no_notification"`
	OffsetX                      *int                           `facebook:"offset_x"`
	OffsetY                      *int                           `facebook:"offset_y"`
	Parking                      *map[string]interface{}        `facebook:"parking"`
	PaymentOptions               *map[string]interface{}        `facebook:"payment_options"`
	Phone                        *string                        `facebook:"phone"`
	PickupOptions                *[]enums.PagePickupOptions     `facebook:"pickup_options"`
	PlotOutline                  *string                        `facebook:"plot_outline"`
	PriceRange                   *string                        `facebook:"price_range"`
	PriorityHours                *map[string]interface{}        `facebook:"priority_hours"`
	PublicTransit                *string                        `facebook:"public_transit"`
	RestaurantServices           *map[string]interface{}        `facebook:"restaurant_services"`
	RestaurantSpecialties        *map[string]interface{}        `facebook:"restaurant_specialties"`
	Scrape                       *bool                          `facebook:"scrape"`
	ServiceDetails               *string                        `facebook:"service_details"`
	SphericalMetadata            *map[string]interface{}        `facebook:"spherical_metadata"`
	StartInfo                    *map[string]interface{}        `facebook:"start_info"`
	StoreLocationDescriptor      *string                        `facebook:"store_location_descriptor"`
	TemporaryStatus              *enums.PageTemporaryStatus     `facebook:"temporary_status"`
	Website                      *string                        `facebook:"website"`
	ZoomScaleX                   *float64                       `facebook:"zoom_scale_x"`
	ZoomScaleY                   *float64                       `facebook:"zoom_scale_y"`
	Extra                        core.Params                    `facebook:"-"`
}

func (params UpdatePageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.About != nil {
		out["about"] = *params.About
	}
	if params.AcceptCrosspostingHandshake != nil {
		out["accept_crossposting_handshake"] = *params.AcceptCrosspostingHandshake
	}
	if params.AllowSphericalPhoto != nil {
		out["allow_spherical_photo"] = *params.AllowSphericalPhoto
	}
	if params.Attire != nil {
		out["attire"] = *params.Attire
	}
	if params.BeginCrosspostingHandshake != nil {
		out["begin_crossposting_handshake"] = *params.BeginCrosspostingHandshake
	}
	if params.Bio != nil {
		out["bio"] = *params.Bio
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.CategoryList != nil {
		out["category_list"] = *params.CategoryList
	}
	if params.CompanyOverview != nil {
		out["company_overview"] = *params.CompanyOverview
	}
	if params.ContactAddress != nil {
		out["contact_address"] = *params.ContactAddress
	}
	if params.Cover != nil {
		out["cover"] = *params.Cover
	}
	if params.CulinaryTeam != nil {
		out["culinary_team"] = *params.CulinaryTeam
	}
	if params.DeliveryAndPickupOptionInfo != nil {
		out["delivery_and_pickup_option_info"] = *params.DeliveryAndPickupOptionInfo
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DifferentlyOpenOfferings != nil {
		out["differently_open_offerings"] = *params.DifferentlyOpenOfferings
	}
	if params.DirectedBy != nil {
		out["directed_by"] = *params.DirectedBy
	}
	if params.DisplayedMessageResponseTime != nil {
		out["displayed_message_response_time"] = *params.DisplayedMessageResponseTime
	}
	if params.Emails != nil {
		out["emails"] = *params.Emails
	}
	if params.FocusX != nil {
		out["focus_x"] = *params.FocusX
	}
	if params.FocusY != nil {
		out["focus_y"] = *params.FocusY
	}
	if params.FoodStyles != nil {
		out["food_styles"] = *params.FoodStyles
	}
	if params.GenAiProvenanceType != nil {
		out["gen_ai_provenance_type"] = *params.GenAiProvenanceType
	}
	if params.GeneralInfo != nil {
		out["general_info"] = *params.GeneralInfo
	}
	if params.GeneralManager != nil {
		out["general_manager"] = *params.GeneralManager
	}
	if params.Genre != nil {
		out["genre"] = *params.Genre
	}
	if params.Hours != nil {
		out["hours"] = *params.Hours
	}
	if params.IgnoreCoordinateWarnings != nil {
		out["ignore_coordinate_warnings"] = *params.IgnoreCoordinateWarnings
	}
	if params.Impressum != nil {
		out["impressum"] = *params.Impressum
	}
	if params.IsAlwaysOpen != nil {
		out["is_always_open"] = *params.IsAlwaysOpen
	}
	if params.IsPermanentlyClosed != nil {
		out["is_permanently_closed"] = *params.IsPermanentlyClosed
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	if params.IsWebhooksSubscribed != nil {
		out["is_webhooks_subscribed"] = *params.IsWebhooksSubscribed
	}
	if params.Location != nil {
		out["location"] = *params.Location
	}
	if params.Menu != nil {
		out["menu"] = *params.Menu
	}
	if params.Mission != nil {
		out["mission"] = *params.Mission
	}
	if params.NoFeedStory != nil {
		out["no_feed_story"] = *params.NoFeedStory
	}
	if params.NoNotification != nil {
		out["no_notification"] = *params.NoNotification
	}
	if params.OffsetX != nil {
		out["offset_x"] = *params.OffsetX
	}
	if params.OffsetY != nil {
		out["offset_y"] = *params.OffsetY
	}
	if params.Parking != nil {
		out["parking"] = *params.Parking
	}
	if params.PaymentOptions != nil {
		out["payment_options"] = *params.PaymentOptions
	}
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.PickupOptions != nil {
		out["pickup_options"] = *params.PickupOptions
	}
	if params.PlotOutline != nil {
		out["plot_outline"] = *params.PlotOutline
	}
	if params.PriceRange != nil {
		out["price_range"] = *params.PriceRange
	}
	if params.PriorityHours != nil {
		out["priority_hours"] = *params.PriorityHours
	}
	if params.PublicTransit != nil {
		out["public_transit"] = *params.PublicTransit
	}
	if params.RestaurantServices != nil {
		out["restaurant_services"] = *params.RestaurantServices
	}
	if params.RestaurantSpecialties != nil {
		out["restaurant_specialties"] = *params.RestaurantSpecialties
	}
	if params.Scrape != nil {
		out["scrape"] = *params.Scrape
	}
	if params.ServiceDetails != nil {
		out["service_details"] = *params.ServiceDetails
	}
	if params.SphericalMetadata != nil {
		out["spherical_metadata"] = *params.SphericalMetadata
	}
	if params.StartInfo != nil {
		out["start_info"] = *params.StartInfo
	}
	if params.StoreLocationDescriptor != nil {
		out["store_location_descriptor"] = *params.StoreLocationDescriptor
	}
	if params.TemporaryStatus != nil {
		out["temporary_status"] = *params.TemporaryStatus
	}
	if params.Website != nil {
		out["website"] = *params.Website
	}
	if params.ZoomScaleX != nil {
		out["zoom_scale_x"] = *params.ZoomScaleX
	}
	if params.ZoomScaleY != nil {
		out["zoom_scale_y"] = *params.ZoomScaleY
	}
	return out
}

func UpdatePageBatchCall(id string, params UpdatePageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdatePageBatchRequest(id string, params UpdatePageParams, options ...core.BatchOption) *core.BatchRequest[objects.Page] {
	return core.NewBatchRequest[objects.Page](UpdatePageBatchCall(id, params, options...))
}

func DecodeUpdatePageBatchResponse(response *core.BatchResponse) (*objects.Page, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Page
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdatePageWithResponse(ctx context.Context, client *core.Client, id string, params UpdatePageParams) (*objects.Page, *core.Response, error) {
	var out objects.Page
	call := UpdatePageBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdatePage(ctx context.Context, client *core.Client, id string, params UpdatePageParams) (*objects.Page, error) {
	out, _, err := UpdatePageWithResponse(ctx, client, id, params)
	return out, err
}
