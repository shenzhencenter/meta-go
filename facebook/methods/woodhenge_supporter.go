package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWoodhengeSupporterParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWoodhengeSupporterParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWoodhengeSupporterBatchCall(id string, params GetWoodhengeSupporterParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWoodhengeSupporterBatchRequest(id string, params GetWoodhengeSupporterParams, options ...core.BatchOption) *core.BatchRequest[objects.WoodhengeSupporter] {
	return core.NewBatchRequest[objects.WoodhengeSupporter](GetWoodhengeSupporterBatchCall(id, params, options...))
}

func DecodeGetWoodhengeSupporterBatchResponse(response *core.BatchResponse) (*objects.WoodhengeSupporter, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WoodhengeSupporter
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWoodhengeSupporterWithResponse(ctx context.Context, client *core.Client, id string, params GetWoodhengeSupporterParams) (*objects.WoodhengeSupporter, *core.Response, error) {
	var out objects.WoodhengeSupporter
	call := GetWoodhengeSupporterBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWoodhengeSupporter(ctx context.Context, client *core.Client, id string, params GetWoodhengeSupporterParams) (*objects.WoodhengeSupporter, error) {
	out, _, err := GetWoodhengeSupporterWithResponse(ctx, client, id, params)
	return out, err
}
