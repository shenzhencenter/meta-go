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

func CreateAdsValueAdjustmentRuleCollectionDeleteRuleSet(ctx context.Context, client *core.Client, id string, params CreateAdsValueAdjustmentRuleCollectionDeleteRuleSetParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "delete_rule_set"), params.ToParams(), &out)
	return &out, err
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

func GetAdsValueAdjustmentRuleCollectionRules(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionRulesParams) (*core.Cursor[objects.AdsValueAdjustmentRulePersona], error) {
	var out core.Cursor[objects.AdsValueAdjustmentRulePersona]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), &out)
	return &out, err
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

func GetAdsValueAdjustmentRuleCollection(ctx context.Context, client *core.Client, id string, params GetAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateAdsValueAdjustmentRuleCollection(ctx context.Context, client *core.Client, id string, params UpdateAdsValueAdjustmentRuleCollectionParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
