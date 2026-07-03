package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCopyrightOwnershipTransferParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCopyrightOwnershipTransferParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCopyrightOwnershipTransfer(ctx context.Context, client *core.Client, id string, params GetCopyrightOwnershipTransferParams) (*objects.CopyrightOwnershipTransfer, error) {
	var out objects.CopyrightOwnershipTransfer
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
