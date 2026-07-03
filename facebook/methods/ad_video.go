package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdVideoBoostAdsListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoBoostAdsListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoBoostAdsListBatchCall(id string, params GetAdVideoBoostAdsListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "boost_ads_list"), params.ToParams(), options...)
}

func NewGetAdVideoBoostAdsListBatchRequest(id string, params GetAdVideoBoostAdsListParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoBoostMediaAd]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoBoostMediaAd]](GetAdVideoBoostAdsListBatchCall(id, params, options...))
}

func DecodeGetAdVideoBoostAdsListBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoBoostMediaAd], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoBoostMediaAd]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoBoostAdsListWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoBoostAdsListParams) (*core.Cursor[objects.VideoBoostMediaAd], *core.Response, error) {
	var out core.Cursor[objects.VideoBoostMediaAd]
	call := GetAdVideoBoostAdsListBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoBoostAdsList(ctx context.Context, client *core.Client, id string, params GetAdVideoBoostAdsListParams) (*core.Cursor[objects.VideoBoostMediaAd], error) {
	out, _, err := GetAdVideoBoostAdsListWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoCaptionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCaptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCaptionsBatchCall(id string, params GetAdVideoCaptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "captions"), params.ToParams(), options...)
}

func NewGetAdVideoCaptionsBatchRequest(id string, params GetAdVideoCaptionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoCaption]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoCaption]](GetAdVideoCaptionsBatchCall(id, params, options...))
}

func DecodeGetAdVideoCaptionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoCaption], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoCaption]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoCaptionsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoCaptionsParams) (*core.Cursor[objects.VideoCaption], *core.Response, error) {
	var out core.Cursor[objects.VideoCaption]
	call := GetAdVideoCaptionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoCaptions(ctx context.Context, client *core.Client, id string, params GetAdVideoCaptionsParams) (*core.Cursor[objects.VideoCaption], error) {
	out, _, err := GetAdVideoCaptionsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoCaptionsParams struct {
	CaptionsFile    *core.FileParam `facebook:"captions_file"`
	DefaultLocale   *string         `facebook:"default_locale"`
	LocalesToDelete *[]string       `facebook:"locales_to_delete"`
	Extra           core.Params     `facebook:"-"`
}

func (params CreateAdVideoCaptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CaptionsFile != nil {
		out["captions_file"] = *params.CaptionsFile
	}
	if params.DefaultLocale != nil {
		out["default_locale"] = *params.DefaultLocale
	}
	if params.LocalesToDelete != nil {
		out["locales_to_delete"] = *params.LocalesToDelete
	}
	return out
}

func CreateAdVideoCaptionsBatchCall(id string, params CreateAdVideoCaptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "captions"), params.ToParams(), options...)
}

func NewCreateAdVideoCaptionsBatchRequest(id string, params CreateAdVideoCaptionsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdVideoCaptionsBatchCall(id, params, options...))
}

func DecodeCreateAdVideoCaptionsBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdVideoCaptionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoCaptionsParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateAdVideoCaptionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoCaptions(ctx context.Context, client *core.Client, id string, params CreateAdVideoCaptionsParams) (*objects.AdVideo, error) {
	out, _, err := CreateAdVideoCaptionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoCollaboratorsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCollaboratorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCollaboratorsBatchCall(id string, params GetAdVideoCollaboratorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaborators"), params.ToParams(), options...)
}

func NewGetAdVideoCollaboratorsBatchRequest(id string, params GetAdVideoCollaboratorsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoCollaborators]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoCollaborators]](GetAdVideoCollaboratorsBatchCall(id, params, options...))
}

func DecodeGetAdVideoCollaboratorsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoCollaborators], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoCollaborators]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoCollaboratorsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoCollaboratorsParams) (*core.Cursor[objects.VideoCollaborators], *core.Response, error) {
	var out core.Cursor[objects.VideoCollaborators]
	call := GetAdVideoCollaboratorsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoCollaborators(ctx context.Context, client *core.Client, id string, params GetAdVideoCollaboratorsParams) (*core.Cursor[objects.VideoCollaborators], error) {
	out, _, err := GetAdVideoCollaboratorsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoCollaboratorsParams struct {
	TargetID core.ID     `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params CreateAdVideoCollaboratorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["target_id"] = params.TargetID
	return out
}

func CreateAdVideoCollaboratorsBatchCall(id string, params CreateAdVideoCollaboratorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "collaborators"), params.ToParams(), options...)
}

func NewCreateAdVideoCollaboratorsBatchRequest(id string, params CreateAdVideoCollaboratorsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdVideoCollaboratorsBatchCall(id, params, options...))
}

func DecodeCreateAdVideoCollaboratorsBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdVideoCollaboratorsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoCollaboratorsParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateAdVideoCollaboratorsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoCollaborators(ctx context.Context, client *core.Client, id string, params CreateAdVideoCollaboratorsParams) (*objects.AdVideo, error) {
	out, _, err := CreateAdVideoCollaboratorsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoCommentsParams struct {
	Filter     *enums.VideocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.VideocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.VideocommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                              `facebook:"since"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetAdVideoCommentsParams) ToParams() core.Params {
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

func GetAdVideoCommentsBatchCall(id string, params GetAdVideoCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetAdVideoCommentsBatchRequest(id string, params GetAdVideoCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetAdVideoCommentsBatchCall(id, params, options...))
}

func DecodeGetAdVideoCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetAdVideoCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetAdVideoCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoComments(ctx context.Context, client *core.Client, id string, params GetAdVideoCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetAdVideoCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoCommentsParams struct {
	AttachmentID         *core.ID                                         `facebook:"attachment_id"`
	AttachmentShareURL   *string                                          `facebook:"attachment_share_url"`
	AttachmentURL        *string                                          `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.VideocommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
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

func (params CreateAdVideoCommentsParams) ToParams() core.Params {
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

func CreateAdVideoCommentsBatchCall(id string, params CreateAdVideoCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateAdVideoCommentsBatchRequest(id string, params CreateAdVideoCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreateAdVideoCommentsBatchCall(id, params, options...))
}

func DecodeCreateAdVideoCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreateAdVideoCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreateAdVideoCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoComments(ctx context.Context, client *core.Client, id string, params CreateAdVideoCommentsParams) (*objects.Comment, error) {
	out, _, err := CreateAdVideoCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoCrosspostSharedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCrosspostSharedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCrosspostSharedPagesBatchCall(id string, params GetAdVideoCrosspostSharedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "crosspost_shared_pages"), params.ToParams(), options...)
}

func NewGetAdVideoCrosspostSharedPagesBatchRequest(id string, params GetAdVideoCrosspostSharedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetAdVideoCrosspostSharedPagesBatchCall(id, params, options...))
}

func DecodeGetAdVideoCrosspostSharedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetAdVideoCrosspostSharedPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetAdVideoCrosspostSharedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoCrosspostSharedPages(ctx context.Context, client *core.Client, id string, params GetAdVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetAdVideoCrosspostSharedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoLikesBatchCall(id string, params GetAdVideoLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetAdVideoLikesBatchRequest(id string, params GetAdVideoLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetAdVideoLikesBatchCall(id, params, options...))
}

func DecodeGetAdVideoLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetAdVideoLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetAdVideoLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoLikes(ctx context.Context, client *core.Client, id string, params GetAdVideoLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetAdVideoLikesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateAdVideoLikesParams) ToParams() core.Params {
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

func CreateAdVideoLikesBatchCall(id string, params CreateAdVideoLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreateAdVideoLikesBatchRequest(id string, params CreateAdVideoLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdVideoLikesBatchCall(id, params, options...))
}

func DecodeCreateAdVideoLikesBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdVideoLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoLikesParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateAdVideoLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoLikes(ctx context.Context, client *core.Client, id string, params CreateAdVideoLikesParams) (*objects.AdVideo, error) {
	out, _, err := CreateAdVideoLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoPollSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoPollSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoPollSettingsBatchCall(id string, params GetAdVideoPollSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "poll_settings"), params.ToParams(), options...)
}

func NewGetAdVideoPollSettingsBatchRequest(id string, params GetAdVideoPollSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PollSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.PollSettings]](GetAdVideoPollSettingsBatchCall(id, params, options...))
}

func DecodeGetAdVideoPollSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PollSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PollSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoPollSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoPollSettingsParams) (*core.Cursor[objects.PollSettings], *core.Response, error) {
	var out core.Cursor[objects.PollSettings]
	call := GetAdVideoPollSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoPollSettings(ctx context.Context, client *core.Client, id string, params GetAdVideoPollSettingsParams) (*core.Cursor[objects.PollSettings], error) {
	out, _, err := GetAdVideoPollSettingsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoPollsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoPollsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoPollsBatchCall(id string, params GetAdVideoPollsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "polls"), params.ToParams(), options...)
}

func NewGetAdVideoPollsBatchRequest(id string, params GetAdVideoPollsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoPoll]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoPoll]](GetAdVideoPollsBatchCall(id, params, options...))
}

func DecodeGetAdVideoPollsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoPoll], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoPoll]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoPollsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoPollsParams) (*core.Cursor[objects.VideoPoll], *core.Response, error) {
	var out core.Cursor[objects.VideoPoll]
	call := GetAdVideoPollsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoPolls(ctx context.Context, client *core.Client, id string, params GetAdVideoPollsParams) (*core.Cursor[objects.VideoPoll], error) {
	out, _, err := GetAdVideoPollsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoPollsParams struct {
	CloseAfterVoting *bool       `facebook:"close_after_voting"`
	CorrectOption    *uint64     `facebook:"correct_option"`
	DefaultOpen      *bool       `facebook:"default_open"`
	Options          []string    `facebook:"options"`
	Question         string      `facebook:"question"`
	ShowGradient     *bool       `facebook:"show_gradient"`
	ShowResults      *bool       `facebook:"show_results"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateAdVideoPollsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CloseAfterVoting != nil {
		out["close_after_voting"] = *params.CloseAfterVoting
	}
	if params.CorrectOption != nil {
		out["correct_option"] = *params.CorrectOption
	}
	if params.DefaultOpen != nil {
		out["default_open"] = *params.DefaultOpen
	}
	out["options"] = params.Options
	out["question"] = params.Question
	if params.ShowGradient != nil {
		out["show_gradient"] = *params.ShowGradient
	}
	if params.ShowResults != nil {
		out["show_results"] = *params.ShowResults
	}
	return out
}

func CreateAdVideoPollsBatchCall(id string, params CreateAdVideoPollsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "polls"), params.ToParams(), options...)
}

func NewCreateAdVideoPollsBatchRequest(id string, params CreateAdVideoPollsParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoPoll] {
	return core.NewBatchRequest[objects.VideoPoll](CreateAdVideoPollsBatchCall(id, params, options...))
}

func DecodeCreateAdVideoPollsBatchResponse(response *core.BatchResponse) (*objects.VideoPoll, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoPoll
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdVideoPollsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoPollsParams) (*objects.VideoPoll, *core.Response, error) {
	var out objects.VideoPoll
	call := CreateAdVideoPollsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoPolls(ctx context.Context, client *core.Client, id string, params CreateAdVideoPollsParams) (*objects.VideoPoll, error) {
	out, _, err := CreateAdVideoPollsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoSponsorTagsBatchCall(id string, params GetAdVideoSponsorTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), options...)
}

func NewGetAdVideoSponsorTagsBatchRequest(id string, params GetAdVideoSponsorTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetAdVideoSponsorTagsBatchCall(id, params, options...))
}

func DecodeGetAdVideoSponsorTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetAdVideoSponsorTagsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoSponsorTagsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetAdVideoSponsorTagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoSponsorTags(ctx context.Context, client *core.Client, id string, params GetAdVideoSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetAdVideoSponsorTagsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoTagsBatchCall(id string, params GetAdVideoTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), options...)
}

func NewGetAdVideoTagsBatchRequest(id string, params GetAdVideoTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.TaggableSubject]] {
	return core.NewBatchRequest[core.Cursor[objects.TaggableSubject]](GetAdVideoTagsBatchCall(id, params, options...))
}

func DecodeGetAdVideoTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.TaggableSubject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.TaggableSubject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoTagsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoTagsParams) (*core.Cursor[objects.TaggableSubject], *core.Response, error) {
	var out core.Cursor[objects.TaggableSubject]
	call := GetAdVideoTagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoTags(ctx context.Context, client *core.Client, id string, params GetAdVideoTagsParams) (*core.Cursor[objects.TaggableSubject], error) {
	out, _, err := GetAdVideoTagsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoThumbnailsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoThumbnailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoThumbnailsBatchCall(id string, params GetAdVideoThumbnailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "thumbnails"), params.ToParams(), options...)
}

func NewGetAdVideoThumbnailsBatchRequest(id string, params GetAdVideoThumbnailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoThumbnail]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoThumbnail]](GetAdVideoThumbnailsBatchCall(id, params, options...))
}

func DecodeGetAdVideoThumbnailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoThumbnail], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoThumbnail]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoThumbnailsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoThumbnailsParams) (*core.Cursor[objects.VideoThumbnail], *core.Response, error) {
	var out core.Cursor[objects.VideoThumbnail]
	call := GetAdVideoThumbnailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoThumbnails(ctx context.Context, client *core.Client, id string, params GetAdVideoThumbnailsParams) (*core.Cursor[objects.VideoThumbnail], error) {
	out, _, err := GetAdVideoThumbnailsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdVideoThumbnailsParams struct {
	IsPreferred *bool          `facebook:"is_preferred"`
	Source      core.FileParam `facebook:"source"`
	Extra       core.Params    `facebook:"-"`
}

func (params CreateAdVideoThumbnailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsPreferred != nil {
		out["is_preferred"] = *params.IsPreferred
	}
	out["source"] = params.Source
	return out
}

func CreateAdVideoThumbnailsBatchCall(id string, params CreateAdVideoThumbnailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "thumbnails"), params.ToParams(), options...)
}

func NewCreateAdVideoThumbnailsBatchRequest(id string, params CreateAdVideoThumbnailsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdVideoThumbnailsBatchCall(id, params, options...))
}

func DecodeCreateAdVideoThumbnailsBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdVideoThumbnailsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdVideoThumbnailsParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateAdVideoThumbnailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdVideoThumbnails(ctx context.Context, client *core.Client, id string, params CreateAdVideoThumbnailsParams) (*objects.AdVideo, error) {
	out, _, err := CreateAdVideoThumbnailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoVideoInsightsParams struct {
	Metric *[]map[string]interface{}                `facebook:"metric"`
	Period *enums.VideovideoInsightsPeriodEnumParam `facebook:"period"`
	Since  *core.Time                               `facebook:"since"`
	Until  *core.Time                               `facebook:"until"`
	Extra  core.Params                              `facebook:"-"`
}

func (params GetAdVideoVideoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
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

func GetAdVideoVideoInsightsBatchCall(id string, params GetAdVideoVideoInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_insights"), params.ToParams(), options...)
}

func NewGetAdVideoVideoInsightsBatchRequest(id string, params GetAdVideoVideoInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetAdVideoVideoInsightsBatchCall(id, params, options...))
}

func DecodeGetAdVideoVideoInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
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

func GetAdVideoVideoInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoVideoInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetAdVideoVideoInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideoVideoInsights(ctx context.Context, client *core.Client, id string, params GetAdVideoVideoInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetAdVideoVideoInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdVideoBatchCall(id string, params DeleteAdVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdVideoBatchRequest(id string, params DeleteAdVideoParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdVideoBatchCall(id, params, options...))
}

func DecodeDeleteAdVideoBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdVideoWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdVideoParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdVideo(ctx context.Context, client *core.Client, id string, params DeleteAdVideoParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdVideoWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoBatchCall(id string, params GetAdVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdVideoBatchRequest(id string, params GetAdVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](GetAdVideoBatchCall(id, params, options...))
}

func DecodeGetAdVideoBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdVideoWithResponse(ctx context.Context, client *core.Client, id string, params GetAdVideoParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := GetAdVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdVideo(ctx context.Context, client *core.Client, id string, params GetAdVideoParams) (*objects.AdVideo, error) {
	out, _, err := GetAdVideoWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdVideoParams struct {
	AdBreaks                  *[]interface{}                       `facebook:"ad_breaks"`
	AllowBmCrossposting       *bool                                `facebook:"allow_bm_crossposting"`
	AllowCrosspostingForPages *[]map[string]interface{}            `facebook:"allow_crossposting_for_pages"`
	BackdatedTime             *core.Time                           `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.VideoBackdatedTimeGranularity `facebook:"backdated_time_granularity"`
	CallToAction              *map[string]interface{}              `facebook:"call_to_action"`
	ContentCategory           *enums.VideoContentCategory          `facebook:"content_category"`
	ContentTags               *[]string                            `facebook:"content_tags"`
	CustomLabels              *[]string                            `facebook:"custom_labels"`
	Description               *string                              `facebook:"description"`
	DirectShareStatus         *uint64                              `facebook:"direct_share_status"`
	Embeddable                *bool                                `facebook:"embeddable"`
	Expiration                *map[string]interface{}              `facebook:"expiration"`
	ExpireNow                 *bool                                `facebook:"expire_now"`
	IncrementPlayCount        *bool                                `facebook:"increment_play_count"`
	Name                      *string                              `facebook:"name"`
	PreferredThumbnailID      *core.ID                             `facebook:"preferred_thumbnail_id"`
	Privacy                   *string                              `facebook:"privacy"`
	PublishToNewsFeed         *bool                                `facebook:"publish_to_news_feed"`
	PublishToVideosTab        *bool                                `facebook:"publish_to_videos_tab"`
	Published                 *bool                                `facebook:"published"`
	ScheduledPublishTime      *uint64                              `facebook:"scheduled_publish_time"`
	SocialActions             *bool                                `facebook:"social_actions"`
	SponsorID                 *core.ID                             `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                              `facebook:"sponsor_relationship"`
	Tags                      *[]string                            `facebook:"tags"`
	Target                    *string                              `facebook:"target"`
	UniversalVideoID          *core.ID                             `facebook:"universal_video_id"`
	Extra                     core.Params                          `facebook:"-"`
}

func (params UpdateAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdBreaks != nil {
		out["ad_breaks"] = *params.AdBreaks
	}
	if params.AllowBmCrossposting != nil {
		out["allow_bm_crossposting"] = *params.AllowBmCrossposting
	}
	if params.AllowCrosspostingForPages != nil {
		out["allow_crossposting_for_pages"] = *params.AllowCrosspostingForPages
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.CustomLabels != nil {
		out["custom_labels"] = *params.CustomLabels
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.Embeddable != nil {
		out["embeddable"] = *params.Embeddable
	}
	if params.Expiration != nil {
		out["expiration"] = *params.Expiration
	}
	if params.ExpireNow != nil {
		out["expire_now"] = *params.ExpireNow
	}
	if params.IncrementPlayCount != nil {
		out["increment_play_count"] = *params.IncrementPlayCount
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.PreferredThumbnailID != nil {
		out["preferred_thumbnail_id"] = *params.PreferredThumbnailID
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PublishToNewsFeed != nil {
		out["publish_to_news_feed"] = *params.PublishToNewsFeed
	}
	if params.PublishToVideosTab != nil {
		out["publish_to_videos_tab"] = *params.PublishToVideosTab
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.SocialActions != nil {
		out["social_actions"] = *params.SocialActions
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.Target != nil {
		out["target"] = *params.Target
	}
	if params.UniversalVideoID != nil {
		out["universal_video_id"] = *params.UniversalVideoID
	}
	return out
}

func UpdateAdVideoBatchCall(id string, params UpdateAdVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdVideoBatchRequest(id string, params UpdateAdVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](UpdateAdVideoBatchCall(id, params, options...))
}

func DecodeUpdateAdVideoBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdVideoWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdVideoParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := UpdateAdVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAdVideo(ctx context.Context, client *core.Client, id string, params UpdateAdVideoParams) (*objects.AdVideo, error) {
	out, _, err := UpdateAdVideoWithResponse(ctx, client, id, params)
	return out, err
}
