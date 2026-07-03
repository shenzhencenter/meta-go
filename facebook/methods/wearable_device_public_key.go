package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWearableDevicePublicKeyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWearableDevicePublicKeyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWearableDevicePublicKeyBatchCall(id string, params GetWearableDevicePublicKeyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWearableDevicePublicKeyBatchRequest(id string, params GetWearableDevicePublicKeyParams, options ...core.BatchOption) *core.BatchRequest[objects.WearableDevicePublicKey] {
	return core.NewBatchRequest[objects.WearableDevicePublicKey](GetWearableDevicePublicKeyBatchCall(id, params, options...))
}

func DecodeGetWearableDevicePublicKeyBatchResponse(response *core.BatchResponse) (*objects.WearableDevicePublicKey, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WearableDevicePublicKey
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWearableDevicePublicKeyWithResponse(ctx context.Context, client *core.Client, id string, params GetWearableDevicePublicKeyParams) (*objects.WearableDevicePublicKey, *core.Response, error) {
	var out objects.WearableDevicePublicKey
	call := GetWearableDevicePublicKeyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWearableDevicePublicKey(ctx context.Context, client *core.Client, id string, params GetWearableDevicePublicKeyParams) (*objects.WearableDevicePublicKey, error) {
	out, _, err := GetWearableDevicePublicKeyWithResponse(ctx, client, id, params)
	return out, err
}
