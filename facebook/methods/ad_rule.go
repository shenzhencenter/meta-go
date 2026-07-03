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

func CreateAdRuleExecute(ctx context.Context, client *core.Client, id string, params CreateAdRuleExecuteParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "execute"), params.ToParams(), &out)
	return &out, err
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

func GetAdRuleHistory(ctx context.Context, client *core.Client, id string, params GetAdRuleHistoryParams) (*core.Cursor[objects.AdRuleHistory], error) {
	var out core.Cursor[objects.AdRuleHistory]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "history"), params.ToParams(), &out)
	return &out, err
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

func CreateAdRulePreview(ctx context.Context, client *core.Client, id string, params CreateAdRulePreviewParams) (*objects.AdRule, error) {
	var out objects.AdRule
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "preview"), params.ToParams(), &out)
	return &out, err
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

func DeleteAdRule(ctx context.Context, client *core.Client, id string, params DeleteAdRuleParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetAdRule(ctx context.Context, client *core.Client, id string, params GetAdRuleParams) (*objects.AdRule, error) {
	var out objects.AdRule
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateAdRule(ctx context.Context, client *core.Client, id string, params UpdateAdRuleParams) (*objects.AdRule, error) {
	var out objects.AdRule
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
