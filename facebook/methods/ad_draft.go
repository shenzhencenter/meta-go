package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdDraftParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdDraftParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdDraft(ctx context.Context, client *core.Client, id string, params GetAdDraftParams) (*objects.AdDraft, error) {
	var out objects.AdDraft
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
