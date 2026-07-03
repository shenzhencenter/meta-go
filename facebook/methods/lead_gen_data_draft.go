package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetLeadGenDataDraftParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadGenDataDraftParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadGenDataDraft(ctx context.Context, client *core.Client, id string, params GetLeadGenDataDraftParams) (*objects.LeadGenDataDraft, error) {
	var out objects.LeadGenDataDraft
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
