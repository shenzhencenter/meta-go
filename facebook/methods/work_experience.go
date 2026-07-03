package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWorkExperienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWorkExperienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWorkExperience(ctx context.Context, client *core.Client, id string, params GetWorkExperienceParams) (*objects.WorkExperience, error) {
	var out objects.WorkExperience
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
