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

func GetRobot(ctx context.Context, client *core.Client, id string, params GetRobotParams) (*objects.Robot, error) {
	var out objects.Robot
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
