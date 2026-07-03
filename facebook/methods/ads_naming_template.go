package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsNamingTemplateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsNamingTemplateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsNamingTemplateBatchCall(id string, params GetAdsNamingTemplateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsNamingTemplateBatchRequest(id string, params GetAdsNamingTemplateParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsNamingTemplate] {
	return core.NewBatchRequest[objects.AdsNamingTemplate](GetAdsNamingTemplateBatchCall(id, params, options...))
}

func DecodeGetAdsNamingTemplateBatchResponse(response *core.BatchResponse) (*objects.AdsNamingTemplate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsNamingTemplate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsNamingTemplate(ctx context.Context, client *core.Client, id string, params GetAdsNamingTemplateParams) (*objects.AdsNamingTemplate, error) {
	var out objects.AdsNamingTemplate
	call := GetAdsNamingTemplateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
