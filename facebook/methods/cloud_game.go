package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCloudGameParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCloudGameParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCloudGameBatchCall(id string, params GetCloudGameParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCloudGameBatchRequest(id string, params GetCloudGameParams, options ...core.BatchOption) *core.BatchRequest[objects.CloudGame] {
	return core.NewBatchRequest[objects.CloudGame](GetCloudGameBatchCall(id, params, options...))
}

func DecodeGetCloudGameBatchResponse(response *core.BatchResponse) (*objects.CloudGame, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CloudGame
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCloudGame(ctx context.Context, client *core.Client, id string, params GetCloudGameParams) (*objects.CloudGame, error) {
	var out objects.CloudGame
	call := GetCloudGameBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
