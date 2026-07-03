package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateStatusLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateStatusLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttributionIDV2 != nil {
		out["attribution_id_v2"] = *params.AttributionIDV2
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.Notify != nil {
		out["notify"] = *params.Notify
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreateStatusLikesBatchCall(id string, params CreateStatusLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreateStatusLikesBatchRequest(id string, params CreateStatusLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.Status] {
	return core.NewBatchRequest[objects.Status](CreateStatusLikesBatchCall(id, params, options...))
}

func DecodeCreateStatusLikesBatchResponse(response *core.BatchResponse) (*objects.Status, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Status
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateStatusLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreateStatusLikesParams) (*objects.Status, *core.Response, error) {
	var out objects.Status
	call := CreateStatusLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateStatusLikes(ctx context.Context, client *core.Client, id string, params CreateStatusLikesParams) (*objects.Status, error) {
	out, _, err := CreateStatusLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStatusBatchCall(id string, params GetStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetStatusBatchRequest(id string, params GetStatusParams, options ...core.BatchOption) *core.BatchRequest[objects.Status] {
	return core.NewBatchRequest[objects.Status](GetStatusBatchCall(id, params, options...))
}

func DecodeGetStatusBatchResponse(response *core.BatchResponse) (*objects.Status, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Status
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetStatusParams) (*objects.Status, *core.Response, error) {
	var out objects.Status
	call := GetStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetStatus(ctx context.Context, client *core.Client, id string, params GetStatusParams) (*objects.Status, error) {
	out, _, err := GetStatusWithResponse(ctx, client, id, params)
	return out, err
}
