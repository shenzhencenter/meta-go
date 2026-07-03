package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateAdRuleExecuteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdRuleExecuteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdRuleExecuteBatchCall(id string, params CreateAdRuleExecuteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "execute"), params.ToParams(), options...)
}

func NewCreateAdRuleExecuteBatchRequest(id string, params CreateAdRuleExecuteParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdRuleExecuteBatchCall(id, params, options...))
}

func DecodeCreateAdRuleExecuteBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdRuleExecuteWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdRuleExecuteParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateAdRuleExecuteBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdRuleExecute(ctx context.Context, client *core.Client, id string, params CreateAdRuleExecuteParams) (*map[string]interface{}, error) {
	out, _, err := CreateAdRuleExecuteWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdRuleHistoryParams struct {
	Action        *enums.AdrulehistoryActionEnumParam `facebook:"action"`
	HideNoChanges *bool                               `facebook:"hide_no_changes"`
	ObjectID      *core.ID                            `facebook:"object_id"`
	Extra         core.Params                         `facebook:"-"`
}

func (params GetAdRuleHistoryParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Action != nil {
		out["action"] = *params.Action
	}
	if params.HideNoChanges != nil {
		out["hide_no_changes"] = *params.HideNoChanges
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	return out
}

func GetAdRuleHistoryBatchCall(id string, params GetAdRuleHistoryParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "history"), params.ToParams(), options...)
}

func NewGetAdRuleHistoryBatchRequest(id string, params GetAdRuleHistoryParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdRuleHistory]] {
	return core.NewBatchRequest[core.Cursor[objects.AdRuleHistory]](GetAdRuleHistoryBatchCall(id, params, options...))
}

func DecodeGetAdRuleHistoryBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdRuleHistory], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdRuleHistory]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdRuleHistoryWithResponse(ctx context.Context, client *core.Client, id string, params GetAdRuleHistoryParams) (*core.Cursor[objects.AdRuleHistory], *core.Response, error) {
	var out core.Cursor[objects.AdRuleHistory]
	call := GetAdRuleHistoryBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdRuleHistory(ctx context.Context, client *core.Client, id string, params GetAdRuleHistoryParams) (*core.Cursor[objects.AdRuleHistory], error) {
	out, _, err := GetAdRuleHistoryWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdRulePreviewParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdRulePreviewParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdRulePreviewBatchCall(id string, params CreateAdRulePreviewParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "preview"), params.ToParams(), options...)
}

func NewCreateAdRulePreviewBatchRequest(id string, params CreateAdRulePreviewParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRule] {
	return core.NewBatchRequest[objects.AdRule](CreateAdRulePreviewBatchCall(id, params, options...))
}

func DecodeCreateAdRulePreviewBatchResponse(response *core.BatchResponse) (*objects.AdRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdRulePreviewWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdRulePreviewParams) (*objects.AdRule, *core.Response, error) {
	var out objects.AdRule
	call := CreateAdRulePreviewBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdRulePreview(ctx context.Context, client *core.Client, id string, params CreateAdRulePreviewParams) (*objects.AdRule, error) {
	out, _, err := CreateAdRulePreviewWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdRuleBatchCall(id string, params DeleteAdRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdRuleBatchRequest(id string, params DeleteAdRuleParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdRuleBatchCall(id, params, options...))
}

func DecodeDeleteAdRuleBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdRuleWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdRuleParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdRule(ctx context.Context, client *core.Client, id string, params DeleteAdRuleParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdRuleWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdRuleBatchCall(id string, params GetAdRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdRuleBatchRequest(id string, params GetAdRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRule] {
	return core.NewBatchRequest[objects.AdRule](GetAdRuleBatchCall(id, params, options...))
}

func DecodeGetAdRuleBatchResponse(response *core.BatchResponse) (*objects.AdRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdRuleWithResponse(ctx context.Context, client *core.Client, id string, params GetAdRuleParams) (*objects.AdRule, *core.Response, error) {
	var out objects.AdRule
	call := GetAdRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdRule(ctx context.Context, client *core.Client, id string, params GetAdRuleParams) (*objects.AdRule, error) {
	out, _, err := GetAdRuleWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdRuleParams struct {
	EvaluationSpec *map[string]interface{} `facebook:"evaluation_spec"`
	ExecutionSpec  *map[string]interface{} `facebook:"execution_spec"`
	Name           *string                 `facebook:"name"`
	ScheduleSpec   *map[string]interface{} `facebook:"schedule_spec"`
	Status         *enums.AdruleStatus     `facebook:"status"`
	Extra          core.Params             `facebook:"-"`
}

func (params UpdateAdRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EvaluationSpec != nil {
		out["evaluation_spec"] = *params.EvaluationSpec
	}
	if params.ExecutionSpec != nil {
		out["execution_spec"] = *params.ExecutionSpec
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.ScheduleSpec != nil {
		out["schedule_spec"] = *params.ScheduleSpec
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func UpdateAdRuleBatchCall(id string, params UpdateAdRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdRuleBatchRequest(id string, params UpdateAdRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRule] {
	return core.NewBatchRequest[objects.AdRule](UpdateAdRuleBatchCall(id, params, options...))
}

func DecodeUpdateAdRuleBatchResponse(response *core.BatchResponse) (*objects.AdRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdRuleWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdRuleParams) (*objects.AdRule, *core.Response, error) {
	var out objects.AdRule
	call := UpdateAdRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdRule(ctx context.Context, client *core.Client, id string, params UpdateAdRuleParams) (*objects.AdRule, error) {
	out, _, err := UpdateAdRuleWithResponse(ctx, client, id, params)
	return out, err
}
