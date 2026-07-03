package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetVideoCopyrightRuleBatchCall(id string, params GetVideoCopyrightRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoCopyrightRuleBatchRequest(id string, params GetVideoCopyrightRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyrightRule] {
	return core.NewBatchRequest[objects.VideoCopyrightRule](GetVideoCopyrightRuleBatchCall(id, params, options...))
}

func DecodeGetVideoCopyrightRuleBatchResponse(response *core.BatchResponse) (*objects.VideoCopyrightRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyrightRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoCopyrightRuleWithResponse(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightRuleParams) (*objects.VideoCopyrightRule, *core.Response, error) {
	var out objects.VideoCopyrightRule
	call := GetVideoCopyrightRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetVideoCopyrightRule(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightRuleParams) (*objects.VideoCopyrightRule, error) {
	out, _, err := GetVideoCopyrightRuleWithResponse(ctx, client, id, params)
	return out, err
}
