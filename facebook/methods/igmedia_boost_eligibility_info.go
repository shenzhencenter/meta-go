package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGMediaBoostEligibilityInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaBoostEligibilityInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaBoostEligibilityInfo(ctx context.Context, client *core.Client, id string, params GetIGMediaBoostEligibilityInfoParams) (*objects.IGMediaBoostEligibilityInfo, error) {
	var out objects.IGMediaBoostEligibilityInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
