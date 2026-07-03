package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAccountAgencyFeeConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAgencyFeeConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAgencyFeeConfigBatchCall(id string, params GetAdAccountAgencyFeeConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAccountAgencyFeeConfigBatchRequest(id string, params GetAdAccountAgencyFeeConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountAgencyFeeConfig] {
	return core.NewBatchRequest[objects.AdAccountAgencyFeeConfig](GetAdAccountAgencyFeeConfigBatchCall(id, params, options...))
}

func DecodeGetAdAccountAgencyFeeConfigBatchResponse(response *core.BatchResponse) (*objects.AdAccountAgencyFeeConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountAgencyFeeConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAgencyFeeConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAccountAgencyFeeConfigParams) (*objects.AdAccountAgencyFeeConfig, *core.Response, error) {
	var out objects.AdAccountAgencyFeeConfig
	call := GetAdAccountAgencyFeeConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAccountAgencyFeeConfig(ctx context.Context, client *core.Client, id string, params GetAdAccountAgencyFeeConfigParams) (*objects.AdAccountAgencyFeeConfig, error) {
	out, _, err := GetAdAccountAgencyFeeConfigWithResponse(ctx, client, id, params)
	return out, err
}
