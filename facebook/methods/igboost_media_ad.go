package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGBoostMediaAdParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGBoostMediaAdParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGBoostMediaAdBatchCall(id string, params GetIGBoostMediaAdParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGBoostMediaAdBatchRequest(id string, params GetIGBoostMediaAdParams, options ...core.BatchOption) *core.BatchRequest[objects.IGBoostMediaAd] {
	return core.NewBatchRequest[objects.IGBoostMediaAd](GetIGBoostMediaAdBatchCall(id, params, options...))
}

func DecodeGetIGBoostMediaAdBatchResponse(response *core.BatchResponse) (*objects.IGBoostMediaAd, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGBoostMediaAd
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGBoostMediaAd(ctx context.Context, client *core.Client, id string, params GetIGBoostMediaAdParams) (*objects.IGBoostMediaAd, error) {
	var out objects.IGBoostMediaAd
	call := GetIGBoostMediaAdBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
