package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetEducationExperienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEducationExperienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEducationExperience(ctx context.Context, client *core.Client, id string, params GetEducationExperienceParams) (*objects.EducationExperience, error) {
	var out objects.EducationExperience
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
