package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetHotelRoomPricingVariablesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelRoomPricingVariablesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelRoomPricingVariablesBatchCall(id string, params GetHotelRoomPricingVariablesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "pricing_variables"), params.ToParams(), options...)
}

func NewGetHotelRoomPricingVariablesBatchRequest(id string, params GetHotelRoomPricingVariablesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicPriceConfigByDate]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicPriceConfigByDate]](GetHotelRoomPricingVariablesBatchCall(id, params, options...))
}

func DecodeGetHotelRoomPricingVariablesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicPriceConfigByDate], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.DynamicPriceConfigByDate]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHotelRoomPricingVariablesWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelRoomPricingVariablesParams) (*core.Cursor[objects.DynamicPriceConfigByDate], *core.Response, error) {
	var out core.Cursor[objects.DynamicPriceConfigByDate]
	call := GetHotelRoomPricingVariablesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelRoomPricingVariables(ctx context.Context, client *core.Client, id string, params GetHotelRoomPricingVariablesParams) (*core.Cursor[objects.DynamicPriceConfigByDate], error) {
	out, _, err := GetHotelRoomPricingVariablesWithResponse(ctx, client, id, params)
	return out, err
}

type GetHotelRoomParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelRoomParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelRoomBatchCall(id string, params GetHotelRoomParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetHotelRoomBatchRequest(id string, params GetHotelRoomParams, options ...core.BatchOption) *core.BatchRequest[objects.HotelRoom] {
	return core.NewBatchRequest[objects.HotelRoom](GetHotelRoomBatchCall(id, params, options...))
}

func DecodeGetHotelRoomBatchResponse(response *core.BatchResponse) (*objects.HotelRoom, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.HotelRoom
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHotelRoomWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelRoomParams) (*objects.HotelRoom, *core.Response, error) {
	var out objects.HotelRoom
	call := GetHotelRoomBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelRoom(ctx context.Context, client *core.Client, id string, params GetHotelRoomParams) (*objects.HotelRoom, error) {
	out, _, err := GetHotelRoomWithResponse(ctx, client, id, params)
	return out, err
}
