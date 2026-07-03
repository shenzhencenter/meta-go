package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAppLinksParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAppLinksParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAppLinksBatchCall(id string, params GetAppLinksParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAppLinksBatchRequest(id string, params GetAppLinksParams, options ...core.BatchOption) *core.BatchRequest[objects.AppLinks] {
	return core.NewBatchRequest[objects.AppLinks](GetAppLinksBatchCall(id, params, options...))
}

func DecodeGetAppLinksBatchResponse(response *core.BatchResponse) (*objects.AppLinks, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AppLinks
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAppLinksWithResponse(ctx context.Context, client *core.Client, id string, params GetAppLinksParams) (*objects.AppLinks, *core.Response, error) {
	var out objects.AppLinks
	call := GetAppLinksBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAppLinks(ctx context.Context, client *core.Client, id string, params GetAppLinksParams) (*objects.AppLinks, error) {
	out, _, err := GetAppLinksWithResponse(ctx, client, id, params)
	return out, err
}
