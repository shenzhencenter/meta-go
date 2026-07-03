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

func GetInstagramUserUpcomingEventsBatchCall(id string, params GetInstagramUserUpcomingEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "upcoming_events"), params.ToParams(), options...)
}

func NewGetInstagramUserUpcomingEventsBatchRequest(id string, params GetInstagramUserUpcomingEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUpcomingEvent]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUpcomingEvent]](GetInstagramUserUpcomingEventsBatchCall(id, params, options...))
}

func DecodeGetInstagramUserUpcomingEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUpcomingEvent], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUpcomingEvent]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetInstagramUserUpcomingEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetInstagramUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], *core.Response, error) {
	var out core.Cursor[objects.IGUpcomingEvent]
	call := GetInstagramUserUpcomingEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetInstagramUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params GetInstagramUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], error) {
	out, _, err := GetInstagramUserUpcomingEventsWithResponse(ctx, client, id, params)
	return out, err
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

func GetInstagramUserBatchCall(id string, params GetInstagramUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetInstagramUserBatchRequest(id string, params GetInstagramUserParams, options ...core.BatchOption) *core.BatchRequest[objects.InstagramUser] {
	return core.NewBatchRequest[objects.InstagramUser](GetInstagramUserBatchCall(id, params, options...))
}

func DecodeGetInstagramUserBatchResponse(response *core.BatchResponse) (*objects.InstagramUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.InstagramUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetInstagramUserWithResponse(ctx context.Context, client *core.Client, id string, params GetInstagramUserParams) (*objects.InstagramUser, *core.Response, error) {
	var out objects.InstagramUser
	call := GetInstagramUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetInstagramUser(ctx context.Context, client *core.Client, id string, params GetInstagramUserParams) (*objects.InstagramUser, error) {
	out, _, err := GetInstagramUserWithResponse(ctx, client, id, params)
	return out, err
}
