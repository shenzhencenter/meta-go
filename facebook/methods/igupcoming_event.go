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

func GetIGUpcomingEvent(ctx context.Context, client *core.Client, id string, params GetIGUpcomingEventParams) (*objects.IGUpcomingEvent, error) {
	var out objects.IGUpcomingEvent
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateIGUpcomingEvent(ctx context.Context, client *core.Client, id string, params UpdateIGUpcomingEventParams) (*objects.IGUpcomingEvent, error) {
	var out objects.IGUpcomingEvent
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
