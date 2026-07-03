package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetEventSourceGroupSharedAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventSourceGroupSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventSourceGroupSharedAccounts(ctx context.Context, client *core.Client, id string, params GetEventSourceGroupSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateEventSourceGroupSharedAccountsParams struct {
	Accounts []string    `facebook:"accounts"`
	Extra    core.Params `facebook:"-"`
}

func (params CreateEventSourceGroupSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["accounts"] = params.Accounts
	return out
}

func CreateEventSourceGroupSharedAccounts(ctx context.Context, client *core.Client, id string, params CreateEventSourceGroupSharedAccountsParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetEventSourceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventSourceGroup(ctx context.Context, client *core.Client, id string, params GetEventSourceGroupParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateEventSourceGroupParams struct {
	EventSources []string    `facebook:"event_sources"`
	Name         string      `facebook:"name"`
	Extra        core.Params `facebook:"-"`
}

func (params UpdateEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["event_sources"] = params.EventSources
	out["name"] = params.Name
	return out
}

func UpdateEventSourceGroup(ctx context.Context, client *core.Client, id string, params UpdateEventSourceGroupParams) (*objects.EventSourceGroup, error) {
	var out objects.EventSourceGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
