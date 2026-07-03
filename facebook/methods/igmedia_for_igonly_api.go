package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGMediaForIGOnlyAPIChildrenParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIChildrenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPIChildrenBatchCall(id string, params GetIGMediaForIGOnlyAPIChildrenParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "children"), params.ToParams(), options...)
}

func NewGetIGMediaForIGOnlyAPIChildrenBatchRequest(id string, params GetIGMediaForIGOnlyAPIChildrenParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Media]] {
	return core.NewBatchRequest[core.Cursor[objects.Media]](GetIGMediaForIGOnlyAPIChildrenBatchCall(id, params, options...))
}

func DecodeGetIGMediaForIGOnlyAPIChildrenBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Media], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Media]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaForIGOnlyAPIChildrenWithResponse(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIChildrenParams) (*core.Cursor[objects.Media], *core.Response, error) {
	var out core.Cursor[objects.Media]
	call := GetIGMediaForIGOnlyAPIChildrenBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGMediaForIGOnlyAPIChildren(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIChildrenParams) (*core.Cursor[objects.Media], error) {
	out, _, err := GetIGMediaForIGOnlyAPIChildrenWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGMediaForIGOnlyAPICommentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPICommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPICommentsBatchCall(id string, params GetIGMediaForIGOnlyAPICommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetIGMediaForIGOnlyAPICommentsBatchRequest(id string, params GetIGMediaForIGOnlyAPICommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetIGMediaForIGOnlyAPICommentsBatchCall(id, params, options...))
}

func DecodeGetIGMediaForIGOnlyAPICommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Comment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaForIGOnlyAPICommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPICommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetIGMediaForIGOnlyAPICommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGMediaForIGOnlyAPIComments(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPICommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetIGMediaForIGOnlyAPICommentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGMediaForIGOnlyAPICommentsParams struct {
	Message *string     `facebook:"message"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGMediaForIGOnlyAPICommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateIGMediaForIGOnlyAPICommentsBatchCall(id string, params CreateIGMediaForIGOnlyAPICommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateIGMediaForIGOnlyAPICommentsBatchRequest(id string, params CreateIGMediaForIGOnlyAPICommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphComment] {
	return core.NewBatchRequest[objects.IGGraphComment](CreateIGMediaForIGOnlyAPICommentsBatchCall(id, params, options...))
}

func DecodeCreateIGMediaForIGOnlyAPICommentsBatchResponse(response *core.BatchResponse) (*objects.IGGraphComment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphComment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGMediaForIGOnlyAPICommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGMediaForIGOnlyAPICommentsParams) (*objects.IGGraphComment, *core.Response, error) {
	var out objects.IGGraphComment
	call := CreateIGMediaForIGOnlyAPICommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGMediaForIGOnlyAPIComments(ctx context.Context, client *core.Client, id string, params CreateIGMediaForIGOnlyAPICommentsParams) (*objects.IGGraphComment, error) {
	out, _, err := CreateIGMediaForIGOnlyAPICommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGMediaForIGOnlyAPIInsightsParams struct {
	Breakdown *[]enums.MediainsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric    []enums.MediainsightsMetricEnumParam     `facebook:"metric"`
	Period    *[]enums.MediainsightsPeriodEnumParam    `facebook:"period"`
	Extra     core.Params                              `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	out["metric"] = params.Metric
	if params.Period != nil {
		out["period"] = *params.Period
	}
	return out
}

func GetIGMediaForIGOnlyAPIInsightsBatchCall(id string, params GetIGMediaForIGOnlyAPIInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetIGMediaForIGOnlyAPIInsightsBatchRequest(id string, params GetIGMediaForIGOnlyAPIInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetIGMediaForIGOnlyAPIInsightsBatchCall(id, params, options...))
}

func DecodeGetIGMediaForIGOnlyAPIInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaForIGOnlyAPIInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetIGMediaForIGOnlyAPIInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGMediaForIGOnlyAPIInsights(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetIGMediaForIGOnlyAPIInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGMediaForIGOnlyAPIParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPIBatchCall(id string, params GetIGMediaForIGOnlyAPIParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGMediaForIGOnlyAPIBatchRequest(id string, params GetIGMediaForIGOnlyAPIParams, options ...core.BatchOption) *core.BatchRequest[objects.Media] {
	return core.NewBatchRequest[objects.Media](GetIGMediaForIGOnlyAPIBatchCall(id, params, options...))
}

func DecodeGetIGMediaForIGOnlyAPIBatchResponse(response *core.BatchResponse) (*objects.Media, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Media
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaForIGOnlyAPIWithResponse(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIParams) (*objects.Media, *core.Response, error) {
	var out objects.Media
	call := GetIGMediaForIGOnlyAPIBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGMediaForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIParams) (*objects.Media, error) {
	out, _, err := GetIGMediaForIGOnlyAPIWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateIGMediaForIGOnlyAPIParams struct {
	CommentEnabled bool        `facebook:"comment_enabled"`
	Extra          core.Params `facebook:"-"`
}

func (params UpdateIGMediaForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["comment_enabled"] = params.CommentEnabled
	return out
}

func UpdateIGMediaForIGOnlyAPIBatchCall(id string, params UpdateIGMediaForIGOnlyAPIParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateIGMediaForIGOnlyAPIBatchRequest(id string, params UpdateIGMediaForIGOnlyAPIParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphMedia] {
	return core.NewBatchRequest[objects.IGGraphMedia](UpdateIGMediaForIGOnlyAPIBatchCall(id, params, options...))
}

func DecodeUpdateIGMediaForIGOnlyAPIBatchResponse(response *core.BatchResponse) (*objects.IGGraphMedia, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphMedia
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateIGMediaForIGOnlyAPIWithResponse(ctx context.Context, client *core.Client, id string, params UpdateIGMediaForIGOnlyAPIParams) (*objects.IGGraphMedia, *core.Response, error) {
	var out objects.IGGraphMedia
	call := UpdateIGMediaForIGOnlyAPIBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateIGMediaForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params UpdateIGMediaForIGOnlyAPIParams) (*objects.IGGraphMedia, error) {
	out, _, err := UpdateIGMediaForIGOnlyAPIWithResponse(ctx, client, id, params)
	return out, err
}
