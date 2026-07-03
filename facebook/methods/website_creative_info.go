package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeInfoBatchCall(id string, params GetWebsiteCreativeInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWebsiteCreativeInfoBatchRequest(id string, params GetWebsiteCreativeInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.WebsiteCreativeInfo] {
	return core.NewBatchRequest[objects.WebsiteCreativeInfo](GetWebsiteCreativeInfoBatchCall(id, params, options...))
}

func DecodeGetWebsiteCreativeInfoBatchResponse(response *core.BatchResponse) (*objects.WebsiteCreativeInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WebsiteCreativeInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWebsiteCreativeInfo(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeInfoParams) (*objects.WebsiteCreativeInfo, error) {
	var out objects.WebsiteCreativeInfo
	call := GetWebsiteCreativeInfoBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
