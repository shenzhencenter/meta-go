package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSiteLinkParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSiteLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSiteLinkBatchCall(id string, params GetSiteLinkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSiteLinkBatchRequest(id string, params GetSiteLinkParams, options ...core.BatchOption) *core.BatchRequest[objects.SiteLink] {
	return core.NewBatchRequest[objects.SiteLink](GetSiteLinkBatchCall(id, params, options...))
}

func DecodeGetSiteLinkBatchResponse(response *core.BatchResponse) (*objects.SiteLink, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SiteLink
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSiteLink(ctx context.Context, client *core.Client, id string, params GetSiteLinkParams) (*objects.SiteLink, error) {
	var out objects.SiteLink
	call := GetSiteLinkBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
