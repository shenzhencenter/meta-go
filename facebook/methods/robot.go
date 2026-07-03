package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetRobotParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRobotParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRobotBatchCall(id string, params GetRobotParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetRobotBatchRequest(id string, params GetRobotParams, options ...core.BatchOption) *core.BatchRequest[objects.Robot] {
	return core.NewBatchRequest[objects.Robot](GetRobotBatchCall(id, params, options...))
}

func DecodeGetRobotBatchResponse(response *core.BatchResponse) (*objects.Robot, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Robot
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetRobot(ctx context.Context, client *core.Client, id string, params GetRobotParams) (*objects.Robot, error) {
	var out objects.Robot
	call := GetRobotBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
