package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessOwnedObjectOnBehalfOfRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedObjectOnBehalfOfRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedObjectOnBehalfOfRequestBatchCall(id string, params GetBusinessOwnedObjectOnBehalfOfRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessOwnedObjectOnBehalfOfRequestBatchRequest(id string, params GetBusinessOwnedObjectOnBehalfOfRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessOwnedObjectOnBehalfOfRequest] {
	return core.NewBatchRequest[objects.BusinessOwnedObjectOnBehalfOfRequest](GetBusinessOwnedObjectOnBehalfOfRequestBatchCall(id, params, options...))
}

func DecodeGetBusinessOwnedObjectOnBehalfOfRequestBatchResponse(response *core.BatchResponse) (*objects.BusinessOwnedObjectOnBehalfOfRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessOwnedObjectOnBehalfOfRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessOwnedObjectOnBehalfOfRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedObjectOnBehalfOfRequestParams) (*objects.BusinessOwnedObjectOnBehalfOfRequest, *core.Response, error) {
	var out objects.BusinessOwnedObjectOnBehalfOfRequest
	call := GetBusinessOwnedObjectOnBehalfOfRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessOwnedObjectOnBehalfOfRequest(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedObjectOnBehalfOfRequestParams) (*objects.BusinessOwnedObjectOnBehalfOfRequest, error) {
	out, _, err := GetBusinessOwnedObjectOnBehalfOfRequestWithResponse(ctx, client, id, params)
	return out, err
}
