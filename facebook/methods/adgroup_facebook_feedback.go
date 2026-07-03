package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdgroupFacebookFeedbackCommentsParams struct {
	Order *enums.AdgroupfacebookfeedbackcommentsOrderEnumParam `facebook:"order"`
	Extra core.Params                                          `facebook:"-"`
}

func (params GetAdgroupFacebookFeedbackCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	return out
}

func GetAdgroupFacebookFeedbackCommentsBatchCall(id string, params GetAdgroupFacebookFeedbackCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetAdgroupFacebookFeedbackCommentsBatchRequest(id string, params GetAdgroupFacebookFeedbackCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetAdgroupFacebookFeedbackCommentsBatchCall(id, params, options...))
}

func DecodeGetAdgroupFacebookFeedbackCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetAdgroupFacebookFeedbackCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdgroupFacebookFeedbackCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetAdgroupFacebookFeedbackCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdgroupFacebookFeedbackComments(ctx context.Context, client *core.Client, id string, params GetAdgroupFacebookFeedbackCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetAdgroupFacebookFeedbackCommentsWithResponse(ctx, client, id, params)
	return out, err
}
