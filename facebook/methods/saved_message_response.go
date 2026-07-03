package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSavedMessageResponseParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSavedMessageResponseParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSavedMessageResponse(ctx context.Context, client *core.Client, id string, params GetSavedMessageResponseParams) (*objects.SavedMessageResponse, error) {
	var out objects.SavedMessageResponse
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
