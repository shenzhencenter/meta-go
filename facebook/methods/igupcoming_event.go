package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetIGUpcomingEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUpcomingEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUpcomingEventBatchCall(id string, params GetIGUpcomingEventParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGUpcomingEventBatchRequest(id string, params GetIGUpcomingEventParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUpcomingEvent] {
	return core.NewBatchRequest[objects.IGUpcomingEvent](GetIGUpcomingEventBatchCall(id, params, options...))
}

func DecodeGetIGUpcomingEventBatchResponse(response *core.BatchResponse) (*objects.IGUpcomingEvent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUpcomingEvent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUpcomingEvent(ctx context.Context, client *core.Client, id string, params GetIGUpcomingEventParams) (*objects.IGUpcomingEvent, error) {
	var out objects.IGUpcomingEvent
	call := GetIGUpcomingEventBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateIGUpcomingEventParams struct {
	EndTime                *time.Time                                   `facebook:"end_time"`
	NotificationSubtypes   *[]enums.IgupcomingeventNotificationSubtypes `facebook:"notification_subtypes"`
	NotificationTargetTime *enums.IgupcomingeventNotificationTargetTime `facebook:"notification_target_time"`
	StartTime              time.Time                                    `facebook:"start_time"`
	Title                  string                                       `facebook:"title"`
	Extra                  core.Params                                  `facebook:"-"`
}

func (params UpdateIGUpcomingEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.NotificationSubtypes != nil {
		out["notification_subtypes"] = *params.NotificationSubtypes
	}
	if params.NotificationTargetTime != nil {
		out["notification_target_time"] = *params.NotificationTargetTime
	}
	out["start_time"] = params.StartTime
	out["title"] = params.Title
	return out
}

func UpdateIGUpcomingEventBatchCall(id string, params UpdateIGUpcomingEventParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateIGUpcomingEventBatchRequest(id string, params UpdateIGUpcomingEventParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUpcomingEvent] {
	return core.NewBatchRequest[objects.IGUpcomingEvent](UpdateIGUpcomingEventBatchCall(id, params, options...))
}

func DecodeUpdateIGUpcomingEventBatchResponse(response *core.BatchResponse) (*objects.IGUpcomingEvent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUpcomingEvent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateIGUpcomingEvent(ctx context.Context, client *core.Client, id string, params UpdateIGUpcomingEventParams) (*objects.IGUpcomingEvent, error) {
	var out objects.IGUpcomingEvent
	call := UpdateIGUpcomingEventBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
