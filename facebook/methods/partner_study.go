package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetPartnerStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerStudy(ctx context.Context, client *core.Client, id string, params GetPartnerStudyParams) (*objects.PartnerStudy, error) {
	var out objects.PartnerStudy
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
