package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOfflineTermsOfServiceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineTermsOfServiceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineTermsOfServiceBatchCall(id string, params GetOfflineTermsOfServiceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOfflineTermsOfServiceBatchRequest(id string, params GetOfflineTermsOfServiceParams, options ...core.BatchOption) *core.BatchRequest[objects.OfflineTermsOfService] {
	return core.NewBatchRequest[objects.OfflineTermsOfService](GetOfflineTermsOfServiceBatchCall(id, params, options...))
}

func DecodeGetOfflineTermsOfServiceBatchResponse(response *core.BatchResponse) (*objects.OfflineTermsOfService, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OfflineTermsOfService
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineTermsOfService(ctx context.Context, client *core.Client, id string, params GetOfflineTermsOfServiceParams) (*objects.OfflineTermsOfService, error) {
	var out objects.OfflineTermsOfService
	call := GetOfflineTermsOfServiceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
