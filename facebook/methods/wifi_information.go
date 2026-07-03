package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWifiInformationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWifiInformationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWifiInformationBatchCall(id string, params GetWifiInformationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWifiInformationBatchRequest(id string, params GetWifiInformationParams, options ...core.BatchOption) *core.BatchRequest[objects.WifiInformation] {
	return core.NewBatchRequest[objects.WifiInformation](GetWifiInformationBatchCall(id, params, options...))
}

func DecodeGetWifiInformationBatchResponse(response *core.BatchResponse) (*objects.WifiInformation, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WifiInformation
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWifiInformationWithResponse(ctx context.Context, client *core.Client, id string, params GetWifiInformationParams) (*objects.WifiInformation, *core.Response, error) {
	var out objects.WifiInformation
	call := GetWifiInformationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWifiInformation(ctx context.Context, client *core.Client, id string, params GetWifiInformationParams) (*objects.WifiInformation, error) {
	out, _, err := GetWifiInformationWithResponse(ctx, client, id, params)
	return out, err
}
