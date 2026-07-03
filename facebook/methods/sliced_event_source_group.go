package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSlicedEventSourceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSlicedEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSlicedEventSourceGroup(ctx context.Context, client *core.Client, id string, params GetSlicedEventSourceGroupParams) (*objects.SlicedEventSourceGroup, error) {
	var out objects.SlicedEventSourceGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
