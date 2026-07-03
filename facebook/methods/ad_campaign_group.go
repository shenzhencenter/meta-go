package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdCampaignGroupGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdCampaignGroupGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdCampaignGroupGendeleteBatchCall(id string, params DeleteAdCampaignGroupGendeleteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), options...)
}

func NewDeleteAdCampaignGroupGendeleteBatchRequest(id string, params DeleteAdCampaignGroupGendeleteParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCampaignGroupDelete] {
	return core.NewBatchRequest[objects.AdCampaignGroupDelete](DeleteAdCampaignGroupGendeleteBatchCall(id, params, options...))
}

func DecodeDeleteAdCampaignGroupGendeleteBatchResponse(response *core.BatchResponse) (*objects.AdCampaignGroupDelete, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCampaignGroupDelete
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdCampaignGroupGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdCampaignGroupGendeleteParams) (*objects.AdCampaignGroupDelete, error) {
	var out objects.AdCampaignGroupDelete
	call := DeleteAdCampaignGroupGendeleteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
