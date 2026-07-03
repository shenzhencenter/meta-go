package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
)

type GetRTBDynamicPostCommentsParams struct {
	Filter     *enums.RtbdynamicpostcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.RtbdynamicpostcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.RtbdynamicpostcommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                                       `facebook:"since"`
	Extra      core.Params                                      `facebook:"-"`
}

func (params GetRTBDynamicPostCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.LiveFilter != nil {
		out["live_filter"] = *params.LiveFilter
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetRTBDynamicPostComments(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type GetRTBDynamicPostLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRTBDynamicPostLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRTBDynamicPostLikes(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetRTBDynamicPostParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRTBDynamicPostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRTBDynamicPost(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostParams) (*objects.RTBDynamicPost, error) {
	var out objects.RTBDynamicPost
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
