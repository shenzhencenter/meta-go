package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeletePageUserMessageThreadLabelLabelParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeletePageUserMessageThreadLabelLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeletePageUserMessageThreadLabelLabel(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelLabelParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "label"), params.ToParams(), &out)
	return &out, err
}

type CreatePageUserMessageThreadLabelLabelParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params CreatePageUserMessageThreadLabelLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func CreatePageUserMessageThreadLabelLabel(ctx context.Context, client *core.Client, id string, params CreatePageUserMessageThreadLabelLabelParams) (*objects.PageUserMessageThreadLabel, error) {
	var out objects.PageUserMessageThreadLabel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "label"), params.ToParams(), &out)
	return &out, err
}

type DeletePageUserMessageThreadLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePageUserMessageThreadLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePageUserMessageThreadLabel(ctx context.Context, client *core.Client, id string, params DeletePageUserMessageThreadLabelParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPageUserMessageThreadLabelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageUserMessageThreadLabelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageUserMessageThreadLabel(ctx context.Context, client *core.Client, id string, params GetPageUserMessageThreadLabelParams) (*objects.PageUserMessageThreadLabel, error) {
	var out objects.PageUserMessageThreadLabel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
