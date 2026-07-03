package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdSavedLocationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedLocationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedLocationBatchCall(id string, params GetAdSavedLocationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdSavedLocationBatchRequest(id string, params GetAdSavedLocationParams, options ...core.BatchOption) *core.BatchRequest[objects.AdSavedLocation] {
	return core.NewBatchRequest[objects.AdSavedLocation](GetAdSavedLocationBatchCall(id, params, options...))
}

func DecodeGetAdSavedLocationBatchResponse(response *core.BatchResponse) (*objects.AdSavedLocation, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdSavedLocation
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdSavedLocation(ctx context.Context, client *core.Client, id string, params GetAdSavedLocationParams) (*objects.AdSavedLocation, error) {
	var out objects.AdSavedLocation
	call := GetAdSavedLocationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
