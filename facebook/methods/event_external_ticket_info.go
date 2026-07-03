package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventExternalTicketInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventExternalTicketInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventExternalTicketInfoBatchCall(id string, params GetEventExternalTicketInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventExternalTicketInfoBatchRequest(id string, params GetEventExternalTicketInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.EventExternalTicketInfo] {
	return core.NewBatchRequest[objects.EventExternalTicketInfo](GetEventExternalTicketInfoBatchCall(id, params, options...))
}

func DecodeGetEventExternalTicketInfoBatchResponse(response *core.BatchResponse) (*objects.EventExternalTicketInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventExternalTicketInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventExternalTicketInfoWithResponse(ctx context.Context, client *core.Client, id string, params GetEventExternalTicketInfoParams) (*objects.EventExternalTicketInfo, *core.Response, error) {
	var out objects.EventExternalTicketInfo
	call := GetEventExternalTicketInfoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventExternalTicketInfo(ctx context.Context, client *core.Client, id string, params GetEventExternalTicketInfoParams) (*objects.EventExternalTicketInfo, error) {
	out, _, err := GetEventExternalTicketInfoWithResponse(ctx, client, id, params)
	return out, err
}
