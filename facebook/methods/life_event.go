package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetLifeEventLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLifeEventLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLifeEventLikes(ctx context.Context, client *core.Client, id string, params GetLifeEventLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetLifeEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLifeEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLifeEvent(ctx context.Context, client *core.Client, id string, params GetLifeEventParams) (*objects.LifeEvent, error) {
	var out objects.LifeEvent
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
