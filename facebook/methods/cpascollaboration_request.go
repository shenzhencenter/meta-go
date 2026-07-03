package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCPASCollaborationRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASCollaborationRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASCollaborationRequest(ctx context.Context, client *core.Client, id string, params GetCPASCollaborationRequestParams) (*objects.CPASCollaborationRequest, error) {
	var out objects.CPASCollaborationRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
