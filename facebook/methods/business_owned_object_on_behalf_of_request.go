package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessOwnedObjectOnBehalfOfRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessOwnedObjectOnBehalfOfRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessOwnedObjectOnBehalfOfRequest(ctx context.Context, client *core.Client, id string, params GetBusinessOwnedObjectOnBehalfOfRequestParams) (*objects.BusinessOwnedObjectOnBehalfOfRequest, error) {
	var out objects.BusinessOwnedObjectOnBehalfOfRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
