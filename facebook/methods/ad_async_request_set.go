package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdAsyncRequestSetRequestsParams struct {
	Statuses *[]enums.AdasyncrequestsetrequestsStatusesEnumParam `facebook:"statuses"`
	Extra    core.Params                                         `facebook:"-"`
}

func (params GetAdAsyncRequestSetRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Statuses != nil {
		out["statuses"] = *params.Statuses
	}
	return out
}

func GetAdAsyncRequestSetRequests(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetRequestsParams) (*core.Cursor[objects.AdAsyncRequest], error) {
	var out core.Cursor[objects.AdAsyncRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "requests"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdAsyncRequestSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestSetParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAdAsyncRequestSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, error) {
	var out objects.AdAsyncRequestSet
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdAsyncRequestSetParams struct {
	Name             *string                                  `facebook:"name"`
	NotificationMode *enums.AdasyncrequestsetNotificationMode `facebook:"notification_mode"`
	NotificationURI  *string                                  `facebook:"notification_uri"`
	Extra            core.Params                              `facebook:"-"`
}

func (params UpdateAdAsyncRequestSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NotificationMode != nil {
		out["notification_mode"] = *params.NotificationMode
	}
	if params.NotificationURI != nil {
		out["notification_uri"] = *params.NotificationURI
	}
	return out
}

func UpdateAdAsyncRequestSet(ctx context.Context, client *core.Client, id string, params UpdateAdAsyncRequestSetParams) (*objects.AdAsyncRequestSet, error) {
	var out objects.AdAsyncRequestSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
