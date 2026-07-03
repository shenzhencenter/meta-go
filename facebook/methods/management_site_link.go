package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetManagementSiteLinkParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetManagementSiteLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetManagementSiteLinkBatchCall(id string, params GetManagementSiteLinkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetManagementSiteLinkBatchRequest(id string, params GetManagementSiteLinkParams, options ...core.BatchOption) *core.BatchRequest[objects.ManagementSiteLink] {
	return core.NewBatchRequest[objects.ManagementSiteLink](GetManagementSiteLinkBatchCall(id, params, options...))
}

func DecodeGetManagementSiteLinkBatchResponse(response *core.BatchResponse) (*objects.ManagementSiteLink, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ManagementSiteLink
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetManagementSiteLink(ctx context.Context, client *core.Client, id string, params GetManagementSiteLinkParams) (*objects.ManagementSiteLink, error) {
	var out objects.ManagementSiteLink
	call := GetManagementSiteLinkBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
