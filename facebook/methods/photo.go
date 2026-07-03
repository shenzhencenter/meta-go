package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPhotoCommentsParams struct {
	Filter     *enums.PhotocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.PhotocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.PhotocommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                              `facebook:"since"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetPhotoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.LiveFilter != nil {
		out["live_filter"] = *params.LiveFilter
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetPhotoCommentsBatchCall(id string, params GetPhotoCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetPhotoCommentsBatchRequest(id string, params GetPhotoCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetPhotoCommentsBatchCall(id, params, options...))
}

func DecodeGetPhotoCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetPhotoCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetPhotoCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetPhotoCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPhotoComments(ctx context.Context, client *core.Client, id string, params GetPhotoCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetPhotoCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePhotoCommentsParams struct {
	AttachmentID         *core.ID                                         `facebook:"attachment_id"`
	AttachmentShareURL   *string                                          `facebook:"attachment_share_url"`
	AttachmentURL        *string                                          `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.PhotocommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FacepileMentionedIds *[]core.ID                                       `facebook:"facepile_mentioned_ids"`
	FeedbackSource       *string                                          `facebook:"feedback_source"`
	IsOffline            *bool                                            `facebook:"is_offline"`
	Message              *string                                          `facebook:"message"`
	NectarModule         *string                                          `facebook:"nectar_module"`
	ObjectID             *core.ID                                         `facebook:"object_id"`
	ParentCommentID      *map[string]interface{}                          `facebook:"parent_comment_id"`
	Text                 *string                                          `facebook:"text"`
	Tracking             *string                                          `facebook:"tracking"`
	Extra                core.Params                                      `facebook:"-"`
}

func (params CreatePhotoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttachmentID != nil {
		out["attachment_id"] = *params.AttachmentID
	}
	if params.AttachmentShareURL != nil {
		out["attachment_share_url"] = *params.AttachmentShareURL
	}
	if params.AttachmentURL != nil {
		out["attachment_url"] = *params.AttachmentURL
	}
	if params.CommentPrivacyValue != nil {
		out["comment_privacy_value"] = *params.CommentPrivacyValue
	}
	if params.FacepileMentionedIds != nil {
		out["facepile_mentioned_ids"] = *params.FacepileMentionedIds
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.IsOffline != nil {
		out["is_offline"] = *params.IsOffline
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	if params.ParentCommentID != nil {
		out["parent_comment_id"] = *params.ParentCommentID
	}
	if params.Text != nil {
		out["text"] = *params.Text
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreatePhotoCommentsBatchCall(id string, params CreatePhotoCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreatePhotoCommentsBatchRequest(id string, params CreatePhotoCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreatePhotoCommentsBatchCall(id, params, options...))
}

func DecodeCreatePhotoCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Comment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePhotoCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePhotoCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreatePhotoCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePhotoComments(ctx context.Context, client *core.Client, id string, params CreatePhotoCommentsParams) (*objects.Comment, error) {
	out, _, err := CreatePhotoCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPhotoInsightsParams struct {
	DatePreset *enums.PhotoinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric     *[]map[string]interface{}               `facebook:"metric"`
	Period     *enums.PhotoinsightsPeriodEnumParam     `facebook:"period"`
	Since      *core.Time                              `facebook:"since"`
	Until      *core.Time                              `facebook:"until"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetPhotoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	if params.Period != nil {
		out["period"] = *params.Period
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetPhotoInsightsBatchCall(id string, params GetPhotoInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetPhotoInsightsBatchRequest(id string, params GetPhotoInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetPhotoInsightsBatchCall(id, params, options...))
}

func DecodeGetPhotoInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
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

func GetPhotoInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetPhotoInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetPhotoInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPhotoInsights(ctx context.Context, client *core.Client, id string, params GetPhotoInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetPhotoInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPhotoLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhotoLikesBatchCall(id string, params GetPhotoLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetPhotoLikesBatchRequest(id string, params GetPhotoLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPhotoLikesBatchCall(id, params, options...))
}

func DecodeGetPhotoLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Profile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPhotoLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetPhotoLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetPhotoLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPhotoLikes(ctx context.Context, client *core.Client, id string, params GetPhotoLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetPhotoLikesWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePhotoLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreatePhotoLikesParams) ToParams() core.Params {
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

func CreatePhotoLikesBatchCall(id string, params CreatePhotoLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreatePhotoLikesBatchRequest(id string, params CreatePhotoLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](CreatePhotoLikesBatchCall(id, params, options...))
}

func DecodeCreatePhotoLikesBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Photo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePhotoLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePhotoLikesParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := CreatePhotoLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePhotoLikes(ctx context.Context, client *core.Client, id string, params CreatePhotoLikesParams) (*objects.Photo, error) {
	out, _, err := CreatePhotoLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetPhotoSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhotoSponsorTagsBatchCall(id string, params GetPhotoSponsorTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), options...)
}

func NewGetPhotoSponsorTagsBatchRequest(id string, params GetPhotoSponsorTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPhotoSponsorTagsBatchCall(id, params, options...))
}

func DecodeGetPhotoSponsorTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPhotoSponsorTagsWithResponse(ctx context.Context, client *core.Client, id string, params GetPhotoSponsorTagsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPhotoSponsorTagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPhotoSponsorTags(ctx context.Context, client *core.Client, id string, params GetPhotoSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPhotoSponsorTagsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePhotoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePhotoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePhotoBatchCall(id string, params DeletePhotoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePhotoBatchRequest(id string, params DeletePhotoParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePhotoBatchCall(id, params, options...))
}

func DecodeDeletePhotoBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeletePhotoWithResponse(ctx context.Context, client *core.Client, id string, params DeletePhotoParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePhotoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePhoto(ctx context.Context, client *core.Client, id string, params DeletePhotoParams) (*map[string]interface{}, error) {
	out, _, err := DeletePhotoWithResponse(ctx, client, id, params)
	return out, err
}

type GetPhotoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPhotoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPhotoBatchCall(id string, params GetPhotoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPhotoBatchRequest(id string, params GetPhotoParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](GetPhotoBatchCall(id, params, options...))
}

func DecodeGetPhotoBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Photo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPhotoWithResponse(ctx context.Context, client *core.Client, id string, params GetPhotoParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := GetPhotoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPhoto(ctx context.Context, client *core.Client, id string, params GetPhotoParams) (*objects.Photo, error) {
	out, _, err := GetPhotoWithResponse(ctx, client, id, params)
	return out, err
}
