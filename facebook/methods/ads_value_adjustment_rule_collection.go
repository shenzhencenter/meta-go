package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams struct {
	Status *enums.AdsvalueadjustmentrulecollectiondeleteRuleSetStatusEnumParam `facebook:"status"`
	Extra  core.Params                                                         `facebook:"-"`
}

func (params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetBatchCall(id string, params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "delete_rule_set"), params.ToParams(), options...)
}

func NewCreateAdsValueAdjustmentRuleCollectionDeleteRuleSetBatchRequest(id string, params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsValueAdjustmentRuleCollection] {
	return core.NewBatchRequest[objects.AdsValueAdjustmentRuleCollection](CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetBatchCall(id, params, options...))
}

func DecodeCreateAdsValueAdjustmentRuleCollectionDeleteRuleSetBatchResponse(response *core.BatchResponse) (*objects.AdsValueAdjustmentRuleCollection, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsValueAdjustmentRuleCollection
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams) (*objects.AdsValueAdjustmentRuleCollection, *core.Response, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	call := CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdsValueAdjustmentRuleCollectionDeleteRuleSet(ctx context.Context, client *core.Client, id string, params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	out, _, err := CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsValueAdjustmentRuleCollectionRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsValueAdjustmentRuleCollectionRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsValueAdjustmentRuleCollectionRulesBatchCall(id string, params GetAdsValueAdjustmentRuleCollectionRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), options...)
}

func NewGetAdsValueAdjustmentRuleCollectionRulesBatchRequest(id string, params GetAdsValueAdjustmentRuleCollectionRulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsValueAdjustmentRulePersona]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsValueAdjustmentRulePersona]](GetAdsValueAdjustmentRuleCollectionRulesBatchCall(id, params, options...))
}

func DecodeGetAdsValueAdjustmentRuleCollectionRulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsValueAdjustmentRulePersona], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsValueAdjustmentRulePersona]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsValueAdjustmentRuleCollectionRulesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionRulesParams) (*core.Cursor[objects.AdsValueAdjustmentRulePersona], *core.Response, error) {
	var out core.Cursor[objects.AdsValueAdjustmentRulePersona]
	call := GetAdsValueAdjustmentRuleCollectionRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsValueAdjustmentRuleCollectionRules(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionRulesParams) (*core.Cursor[objects.AdsValueAdjustmentRulePersona], error) {
	out, _, err := GetAdsValueAdjustmentRuleCollectionRulesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsValueAdjustmentRuleCollectionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsValueAdjustmentRuleCollectionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsValueAdjustmentRuleCollectionBatchCall(id string, params GetAdsValueAdjustmentRuleCollectionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsValueAdjustmentRuleCollectionBatchRequest(id string, params GetAdsValueAdjustmentRuleCollectionParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsValueAdjustmentRuleCollection] {
	return core.NewBatchRequest[objects.AdsValueAdjustmentRuleCollection](GetAdsValueAdjustmentRuleCollectionBatchCall(id, params, options...))
}

func DecodeGetAdsValueAdjustmentRuleCollectionBatchResponse(response *core.BatchResponse) (*objects.AdsValueAdjustmentRuleCollection, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsValueAdjustmentRuleCollection
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsValueAdjustmentRuleCollectionWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, *core.Response, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	call := GetAdsValueAdjustmentRuleCollectionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsValueAdjustmentRuleCollection(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	out, _, err := GetAdsValueAdjustmentRuleCollectionWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdsValueAdjustmentRuleCollectionParams struct {
	EntryPoint       *enums.AdsvalueadjustmentrulecollectionEntryPoint `facebook:"entry_point"`
	IsDefaultSetting *bool                                             `facebook:"is_default_setting"`
	Name             string                                            `facebook:"name"`
	Rules            []map[string]interface{}                          `facebook:"rules"`
	Extra            core.Params                                       `facebook:"-"`
}

func (params UpdateAdsValueAdjustmentRuleCollectionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EntryPoint != nil {
		out["entry_point"] = *params.EntryPoint
	}
	if params.IsDefaultSetting != nil {
		out["is_default_setting"] = *params.IsDefaultSetting
	}
	out["name"] = params.Name
	out["rules"] = params.Rules
	return out
}

func UpdateAdsValueAdjustmentRuleCollectionBatchCall(id string, params UpdateAdsValueAdjustmentRuleCollectionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdsValueAdjustmentRuleCollectionBatchRequest(id string, params UpdateAdsValueAdjustmentRuleCollectionParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsValueAdjustmentRuleCollection] {
	return core.NewBatchRequest[objects.AdsValueAdjustmentRuleCollection](UpdateAdsValueAdjustmentRuleCollectionBatchCall(id, params, options...))
}

func DecodeUpdateAdsValueAdjustmentRuleCollectionBatchResponse(response *core.BatchResponse) (*objects.AdsValueAdjustmentRuleCollection, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsValueAdjustmentRuleCollection
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdsValueAdjustmentRuleCollectionWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, *core.Response, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	call := UpdateAdsValueAdjustmentRuleCollectionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdsValueAdjustmentRuleCollection(ctx context.Context, client *core.Client, id string, params UpdateAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	out, _, err := UpdateAdsValueAdjustmentRuleCollectionWithResponse(ctx, client, id, params)
	return out, err
}
