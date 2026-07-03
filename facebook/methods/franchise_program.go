package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFranchiseProgramParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFranchiseProgramParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFranchiseProgram(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramParams) (*objects.FranchiseProgram, error) {
	var out objects.FranchiseProgram
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
