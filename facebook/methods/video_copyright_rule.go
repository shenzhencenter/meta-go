package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetVideoCopyrightRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoCopyrightRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoCopyrightRule(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightRuleParams) (*objects.VideoCopyrightRule, error) {
	var out objects.VideoCopyrightRule
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
