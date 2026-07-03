package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetStoreLocationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStoreLocationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStoreLocationBatchCall(id string, params GetStoreLocationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetStoreLocationBatchRequest(id string, params GetStoreLocationParams, options ...core.BatchOption) *core.BatchRequest[objects.StoreLocation] {
	return core.NewBatchRequest[objects.StoreLocation](GetStoreLocationBatchCall(id, params, options...))
}

func DecodeGetStoreLocationBatchResponse(response *core.BatchResponse) (*objects.StoreLocation, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.StoreLocation
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetStoreLocationWithResponse(ctx context.Context, client *core.Client, id string, params GetStoreLocationParams) (*objects.StoreLocation, *core.Response, error) {
	var out objects.StoreLocation
	call := GetStoreLocationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetStoreLocation(ctx context.Context, client *core.Client, id string, params GetStoreLocationParams) (*objects.StoreLocation, error) {
	out, _, err := GetStoreLocationWithResponse(ctx, client, id, params)
	return out, err
}
