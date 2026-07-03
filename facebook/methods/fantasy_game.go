package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFantasyGameParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFantasyGameParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFantasyGameBatchCall(id string, params GetFantasyGameParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFantasyGameBatchRequest(id string, params GetFantasyGameParams, options ...core.BatchOption) *core.BatchRequest[objects.FantasyGame] {
	return core.NewBatchRequest[objects.FantasyGame](GetFantasyGameBatchCall(id, params, options...))
}

func DecodeGetFantasyGameBatchResponse(response *core.BatchResponse) (*objects.FantasyGame, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FantasyGame
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFantasyGameWithResponse(ctx context.Context, client *core.Client, id string, params GetFantasyGameParams) (*objects.FantasyGame, *core.Response, error) {
	var out objects.FantasyGame
	call := GetFantasyGameBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFantasyGame(ctx context.Context, client *core.Client, id string, params GetFantasyGameParams) (*objects.FantasyGame, error) {
	out, _, err := GetFantasyGameWithResponse(ctx, client, id, params)
	return out, err
}
