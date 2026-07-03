package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAvatarParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAvatarParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAvatarBatchCall(id string, params GetAvatarParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAvatarBatchRequest(id string, params GetAvatarParams, options ...core.BatchOption) *core.BatchRequest[objects.Avatar] {
	return core.NewBatchRequest[objects.Avatar](GetAvatarBatchCall(id, params, options...))
}

func DecodeGetAvatarBatchResponse(response *core.BatchResponse) (*objects.Avatar, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Avatar
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAvatarWithResponse(ctx context.Context, client *core.Client, id string, params GetAvatarParams) (*objects.Avatar, *core.Response, error) {
	var out objects.Avatar
	call := GetAvatarBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAvatar(ctx context.Context, client *core.Client, id string, params GetAvatarParams) (*objects.Avatar, error) {
	out, _, err := GetAvatarWithResponse(ctx, client, id, params)
	return out, err
}
