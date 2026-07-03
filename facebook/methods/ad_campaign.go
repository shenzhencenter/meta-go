package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdCampaignGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdCampaignGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdCampaignGendeleteBatchCall(id string, params DeleteAdCampaignGendeleteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), options...)
}

func NewDeleteAdCampaignGendeleteBatchRequest(id string, params DeleteAdCampaignGendeleteParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCampaignDelete] {
	return core.NewBatchRequest[objects.AdCampaignDelete](DeleteAdCampaignGendeleteBatchCall(id, params, options...))
}

func DecodeDeleteAdCampaignGendeleteBatchResponse(response *core.BatchResponse) (*objects.AdCampaignDelete, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCampaignDelete
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdCampaignGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdCampaignGendeleteParams) (*objects.AdCampaignDelete, error) {
	var out objects.AdCampaignDelete
	call := DeleteAdCampaignGendeleteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
