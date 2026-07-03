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

func GetPlaceTopicBatchCall(id string, params GetPlaceTopicParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPlaceTopicBatchRequest(id string, params GetPlaceTopicParams, options ...core.BatchOption) *core.BatchRequest[objects.PlaceTopic] {
	return core.NewBatchRequest[objects.PlaceTopic](GetPlaceTopicBatchCall(id, params, options...))
}

func DecodeGetPlaceTopicBatchResponse(response *core.BatchResponse) (*objects.PlaceTopic, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PlaceTopic
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPlaceTopicWithResponse(ctx context.Context, client *core.Client, id string, params GetPlaceTopicParams) (*objects.PlaceTopic, *core.Response, error) {
	var out objects.PlaceTopic
	call := GetPlaceTopicBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPlaceTopic(ctx context.Context, client *core.Client, id string, params GetPlaceTopicParams) (*objects.PlaceTopic, error) {
	out, _, err := GetPlaceTopicWithResponse(ctx, client, id, params)
	return out, err
}
