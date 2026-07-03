package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetResearchPollStudyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetResearchPollStudyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetResearchPollStudy(ctx context.Context, client *core.Client, id string, params GetResearchPollStudyParams) (*objects.ResearchPollStudy, error) {
	var out objects.ResearchPollStudy
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
