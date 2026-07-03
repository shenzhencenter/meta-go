package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShadowIGHashtagRecentMediaParams struct {
	UserID core.ID     `facebook:"user_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagRecentMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_id"] = params.UserID
	return out
}

func GetShadowIGHashtagRecentMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagRecentMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "recent_media"), params.ToParams(), &out)
	return &out, err
}

type GetShadowIGHashtagTopMediaParams struct {
	UserID core.ID     `facebook:"user_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagTopMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_id"] = params.UserID
	return out
}

func GetShadowIGHashtagTopMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagTopMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "top_media"), params.ToParams(), &out)
	return &out, err
}

type GetShadowIGHashtagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShadowIGHashtag(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagParams) (*objects.ShadowIGHashtag, error) {
	var out objects.ShadowIGHashtag
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
