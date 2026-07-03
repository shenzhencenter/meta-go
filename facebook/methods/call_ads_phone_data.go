package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCallAdsPhoneDataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCallAdsPhoneDataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCallAdsPhoneDataBatchCall(id string, params GetCallAdsPhoneDataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCallAdsPhoneDataBatchRequest(id string, params GetCallAdsPhoneDataParams, options ...core.BatchOption) *core.BatchRequest[objects.CallAdsPhoneData] {
	return core.NewBatchRequest[objects.CallAdsPhoneData](GetCallAdsPhoneDataBatchCall(id, params, options...))
}

func DecodeGetCallAdsPhoneDataBatchResponse(response *core.BatchResponse) (*objects.CallAdsPhoneData, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CallAdsPhoneData
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCallAdsPhoneDataWithResponse(ctx context.Context, client *core.Client, id string, params GetCallAdsPhoneDataParams) (*objects.CallAdsPhoneData, *core.Response, error) {
	var out objects.CallAdsPhoneData
	call := GetCallAdsPhoneDataBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCallAdsPhoneData(ctx context.Context, client *core.Client, id string, params GetCallAdsPhoneDataParams) (*objects.CallAdsPhoneData, error) {
	out, _, err := GetCallAdsPhoneDataWithResponse(ctx, client, id, params)
	return out, err
}
