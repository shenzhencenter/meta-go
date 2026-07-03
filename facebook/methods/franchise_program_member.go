package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFranchiseProgramMemberParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFranchiseProgramMemberParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFranchiseProgramMember(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramMemberParams) (*objects.FranchiseProgramMember, error) {
	var out objects.FranchiseProgramMember
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
