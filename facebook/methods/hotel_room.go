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

func GetHotelRoomPricingVariables(ctx context.Context, client *core.Client, id string, params GetHotelRoomPricingVariablesParams) (*core.Cursor[objects.DynamicPriceConfigByDate], error) {
	var out core.Cursor[objects.DynamicPriceConfigByDate]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pricing_variables"), params.ToParams(), &out)
	return &out, err
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

func GetHotelRoom(ctx context.Context, client *core.Client, id string, params GetHotelRoomParams) (*objects.HotelRoom, error) {
	var out objects.HotelRoom
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
