package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPlaceTopicParams struct {
	IconSize *enums.PlacetopicIconSize `facebook:"icon_size"`
	Extra    core.Params               `facebook:"-"`
}

func (params GetPlaceTopicParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IconSize != nil {
		out["icon_size"] = *params.IconSize
	}
	return out
}

func GetPlaceTopic(ctx context.Context, client *core.Client, id string, params GetPlaceTopicParams) (*objects.PlaceTopic, error) {
	var out objects.PlaceTopic
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
