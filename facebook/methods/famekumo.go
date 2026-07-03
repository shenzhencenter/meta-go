package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFAMEKumoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFAMEKumoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFAMEKumoBatchCall(id string, params GetFAMEKumoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFAMEKumoBatchRequest(id string, params GetFAMEKumoParams, options ...core.BatchOption) *core.BatchRequest[objects.FAMEKumo] {
	return core.NewBatchRequest[objects.FAMEKumo](GetFAMEKumoBatchCall(id, params, options...))
}

func DecodeGetFAMEKumoBatchResponse(response *core.BatchResponse) (*objects.FAMEKumo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FAMEKumo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFAMEKumo(ctx context.Context, client *core.Client, id string, params GetFAMEKumoParams) (*objects.FAMEKumo, error) {
	var out objects.FAMEKumo
	call := GetFAMEKumoBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
