package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWorkSkillUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWorkSkillUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWorkSkillUsersBatchCall(id string, params GetWorkSkillUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "users"), params.ToParams(), options...)
}

func NewGetWorkSkillUsersBatchRequest(id string, params GetWorkSkillUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.User]] {
	return core.NewBatchRequest[core.Cursor[objects.User]](GetWorkSkillUsersBatchCall(id, params, options...))
}

func DecodeGetWorkSkillUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.User], error) {
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

func GetWorkSkillUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetWorkSkillUsersParams) (*core.Cursor[objects.User], *core.Response, error) {
	var out core.Cursor[objects.User]
	call := GetWorkSkillUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWorkSkillUsers(ctx context.Context, client *core.Client, id string, params GetWorkSkillUsersParams) (*core.Cursor[objects.User], error) {
	out, _, err := GetWorkSkillUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetWorkSkillParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWorkSkillParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWorkSkillBatchCall(id string, params GetWorkSkillParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWorkSkillBatchRequest(id string, params GetWorkSkillParams, options ...core.BatchOption) *core.BatchRequest[objects.WorkSkill] {
	return core.NewBatchRequest[objects.WorkSkill](GetWorkSkillBatchCall(id, params, options...))
}

func DecodeGetWorkSkillBatchResponse(response *core.BatchResponse) (*objects.WorkSkill, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WorkSkill
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWorkSkillWithResponse(ctx context.Context, client *core.Client, id string, params GetWorkSkillParams) (*objects.WorkSkill, *core.Response, error) {
	var out objects.WorkSkill
	call := GetWorkSkillBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWorkSkill(ctx context.Context, client *core.Client, id string, params GetWorkSkillParams) (*objects.WorkSkill, error) {
	out, _, err := GetWorkSkillWithResponse(ctx, client, id, params)
	return out, err
}
