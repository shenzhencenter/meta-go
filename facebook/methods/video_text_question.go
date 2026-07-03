package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoTextQuestionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoTextQuestionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoTextQuestion(ctx context.Context, client *core.Client, id string, params GetVideoTextQuestionParams) (*objects.VideoTextQuestion, error) {
	var out objects.VideoTextQuestion
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
