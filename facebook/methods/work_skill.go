package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetWorkSkillUsers(ctx context.Context, client *core.Client, id string, params GetWorkSkillUsersParams) (*core.Cursor[objects.User], error) {
	var out core.Cursor[objects.User]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "users"), params.ToParams(), &out)
	return &out, err
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

func GetWorkSkill(ctx context.Context, client *core.Client, id string, params GetWorkSkillParams) (*objects.WorkSkill, error) {
	var out objects.WorkSkill
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
