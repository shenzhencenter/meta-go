package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetInstagramUserUpcomingEventsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetInstagramUserUpcomingEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetInstagramUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params GetInstagramUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], error) {
	var out core.Cursor[objects.IGUpcomingEvent]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "upcoming_events"), params.ToParams(), &out)
	return &out, err
}

type GetInstagramUserParams struct {
	AdgroupID *core.ID    `facebook:"adgroup_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetInstagramUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdgroupID != nil {
		out["adgroup_id"] = *params.AdgroupID
	}
	return out
}

func GetInstagramUser(ctx context.Context, client *core.Client, id string, params GetInstagramUserParams) (*objects.InstagramUser, error) {
	var out objects.InstagramUser
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
