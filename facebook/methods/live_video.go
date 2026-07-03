package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoBlockedUsersParams struct {
	UID   *core.ID    `facebook:"uid"`
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoBlockedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	return out
}

func GetLiveVideoBlockedUsersBatchCall(id string, params GetLiveVideoBlockedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "blocked_users"), params.ToParams(), options...)
}

func NewGetLiveVideoBlockedUsersBatchRequest(id string, params GetLiveVideoBlockedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.User]] {
	return core.NewBatchRequest[core.Cursor[objects.User]](GetLiveVideoBlockedUsersBatchCall(id, params, options...))
}

func DecodeGetLiveVideoBlockedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.User], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.User]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoBlockedUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoBlockedUsersParams) (*core.Cursor[objects.User], *core.Response, error) {
	var out core.Cursor[objects.User]
	call := GetLiveVideoBlockedUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoBlockedUsers(ctx context.Context, client *core.Client, id string, params GetLiveVideoBlockedUsersParams) (*core.Cursor[objects.User], error) {
	out, _, err := GetLiveVideoBlockedUsersWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoCommentsParams struct {
	Filter     *enums.LivevideocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.LivevideocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.LivevideocommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                                  `facebook:"since"`
	Extra      core.Params                                 `facebook:"-"`
}

func (params GetLiveVideoCommentsParams) ToParams() core.Params {
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

func GetLiveVideoCommentsBatchCall(id string, params GetLiveVideoCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetLiveVideoCommentsBatchRequest(id string, params GetLiveVideoCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetLiveVideoCommentsBatchCall(id, params, options...))
}

func DecodeGetLiveVideoCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetLiveVideoCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetLiveVideoCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoComments(ctx context.Context, client *core.Client, id string, params GetLiveVideoCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetLiveVideoCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoCrosspostSharedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoCrosspostSharedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoCrosspostSharedPagesBatchCall(id string, params GetLiveVideoCrosspostSharedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "crosspost_shared_pages"), params.ToParams(), options...)
}

func NewGetLiveVideoCrosspostSharedPagesBatchRequest(id string, params GetLiveVideoCrosspostSharedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetLiveVideoCrosspostSharedPagesBatchCall(id, params, options...))
}

func DecodeGetLiveVideoCrosspostSharedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetLiveVideoCrosspostSharedPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetLiveVideoCrosspostSharedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoCrosspostSharedPages(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetLiveVideoCrosspostSharedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoCrosspostedBroadcastsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoCrosspostedBroadcastsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoCrosspostedBroadcastsBatchCall(id string, params GetLiveVideoCrosspostedBroadcastsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "crossposted_broadcasts"), params.ToParams(), options...)
}

func NewGetLiveVideoCrosspostedBroadcastsBatchRequest(id string, params GetLiveVideoCrosspostedBroadcastsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LiveVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.LiveVideo]](GetLiveVideoCrosspostedBroadcastsBatchCall(id, params, options...))
}

func DecodeGetLiveVideoCrosspostedBroadcastsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LiveVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.LiveVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoCrosspostedBroadcastsWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostedBroadcastsParams) (*core.Cursor[objects.LiveVideo], *core.Response, error) {
	var out core.Cursor[objects.LiveVideo]
	call := GetLiveVideoCrosspostedBroadcastsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoCrosspostedBroadcasts(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostedBroadcastsParams) (*core.Cursor[objects.LiveVideo], error) {
	out, _, err := GetLiveVideoCrosspostedBroadcastsWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoErrorsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoErrorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoErrorsBatchCall(id string, params GetLiveVideoErrorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "errors"), params.ToParams(), options...)
}

func NewGetLiveVideoErrorsBatchRequest(id string, params GetLiveVideoErrorsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LiveVideoError]] {
	return core.NewBatchRequest[core.Cursor[objects.LiveVideoError]](GetLiveVideoErrorsBatchCall(id, params, options...))
}

func DecodeGetLiveVideoErrorsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LiveVideoError], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.LiveVideoError]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoErrorsWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorsParams) (*core.Cursor[objects.LiveVideoError], *core.Response, error) {
	var out core.Cursor[objects.LiveVideoError]
	call := GetLiveVideoErrorsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoErrors(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorsParams) (*core.Cursor[objects.LiveVideoError], error) {
	out, _, err := GetLiveVideoErrorsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateLiveVideoInputStreamsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateLiveVideoInputStreamsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateLiveVideoInputStreamsBatchCall(id string, params CreateLiveVideoInputStreamsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "input_streams"), params.ToParams(), options...)
}

func NewCreateLiveVideoInputStreamsBatchRequest(id string, params CreateLiveVideoInputStreamsParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideoInputStream] {
	return core.NewBatchRequest[objects.LiveVideoInputStream](CreateLiveVideoInputStreamsBatchCall(id, params, options...))
}

func DecodeCreateLiveVideoInputStreamsBatchResponse(response *core.BatchResponse) (*objects.LiveVideoInputStream, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideoInputStream
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateLiveVideoInputStreamsWithResponse(ctx context.Context, client *core.Client, id string, params CreateLiveVideoInputStreamsParams) (*objects.LiveVideoInputStream, *core.Response, error) {
	var out objects.LiveVideoInputStream
	call := CreateLiveVideoInputStreamsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateLiveVideoInputStreams(ctx context.Context, client *core.Client, id string, params CreateLiveVideoInputStreamsParams) (*objects.LiveVideoInputStream, error) {
	out, _, err := CreateLiveVideoInputStreamsWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoPollsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoPollsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoPollsBatchCall(id string, params GetLiveVideoPollsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "polls"), params.ToParams(), options...)
}

func NewGetLiveVideoPollsBatchRequest(id string, params GetLiveVideoPollsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoPoll]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoPoll]](GetLiveVideoPollsBatchCall(id, params, options...))
}

func DecodeGetLiveVideoPollsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoPoll], error) {
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

func GetLiveVideoPollsWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoPollsParams) (*core.Cursor[objects.VideoPoll], *core.Response, error) {
	var out core.Cursor[objects.VideoPoll]
	call := GetLiveVideoPollsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoPolls(ctx context.Context, client *core.Client, id string, params GetLiveVideoPollsParams) (*core.Cursor[objects.VideoPoll], error) {
	out, _, err := GetLiveVideoPollsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateLiveVideoPollsParams struct {
	CloseAfterVoting *bool       `facebook:"close_after_voting"`
	CorrectOption    *uint64     `facebook:"correct_option"`
	DefaultOpen      *bool       `facebook:"default_open"`
	Options          []string    `facebook:"options"`
	Question         string      `facebook:"question"`
	ShowGradient     *bool       `facebook:"show_gradient"`
	ShowResults      *bool       `facebook:"show_results"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateLiveVideoPollsParams) ToParams() core.Params {
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

func CreateLiveVideoPollsBatchCall(id string, params CreateLiveVideoPollsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "polls"), params.ToParams(), options...)
}

func NewCreateLiveVideoPollsBatchRequest(id string, params CreateLiveVideoPollsParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoPoll] {
	return core.NewBatchRequest[objects.VideoPoll](CreateLiveVideoPollsBatchCall(id, params, options...))
}

func DecodeCreateLiveVideoPollsBatchResponse(response *core.BatchResponse) (*objects.VideoPoll, error) {
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

func CreateLiveVideoPollsWithResponse(ctx context.Context, client *core.Client, id string, params CreateLiveVideoPollsParams) (*objects.VideoPoll, *core.Response, error) {
	var out objects.VideoPoll
	call := CreateLiveVideoPollsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateLiveVideoPolls(ctx context.Context, client *core.Client, id string, params CreateLiveVideoPollsParams) (*objects.VideoPoll, error) {
	out, _, err := CreateLiveVideoPollsWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoReactionsParams struct {
	Type  *enums.LivevideoreactionsTypeEnumParam `facebook:"type"`
	Extra core.Params                            `facebook:"-"`
}

func (params GetLiveVideoReactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetLiveVideoReactionsBatchCall(id string, params GetLiveVideoReactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), options...)
}

func NewGetLiveVideoReactionsBatchRequest(id string, params GetLiveVideoReactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetLiveVideoReactionsBatchCall(id, params, options...))
}

func DecodeGetLiveVideoReactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetLiveVideoReactionsWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoReactionsParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetLiveVideoReactionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoReactions(ctx context.Context, client *core.Client, id string, params GetLiveVideoReactionsParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetLiveVideoReactionsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteLiveVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteLiveVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteLiveVideoBatchCall(id string, params DeleteLiveVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteLiveVideoBatchRequest(id string, params DeleteLiveVideoParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteLiveVideoBatchCall(id, params, options...))
}

func DecodeDeleteLiveVideoBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteLiveVideoWithResponse(ctx context.Context, client *core.Client, id string, params DeleteLiveVideoParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteLiveVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteLiveVideo(ctx context.Context, client *core.Client, id string, params DeleteLiveVideoParams) (*map[string]interface{}, error) {
	out, _, err := DeleteLiveVideoWithResponse(ctx, client, id, params)
	return out, err
}

type GetLiveVideoParams struct {
	TargetToken *string     `facebook:"target_token"`
	Extra       core.Params `facebook:"-"`
}

func (params GetLiveVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetToken != nil {
		out["target_token"] = *params.TargetToken
	}
	return out
}

func GetLiveVideoBatchCall(id string, params GetLiveVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLiveVideoBatchRequest(id string, params GetLiveVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](GetLiveVideoBatchCall(id, params, options...))
}

func DecodeGetLiveVideoBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := GetLiveVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideo(ctx context.Context, client *core.Client, id string, params GetLiveVideoParams) (*objects.LiveVideo, error) {
	out, _, err := GetLiveVideoWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateLiveVideoParams struct {
	AllowBmCrossposting          *bool                                          `facebook:"allow_bm_crossposting"`
	ContentTags                  *[]string                                      `facebook:"content_tags"`
	CrossShareToGroupIds         *[]core.ID                                     `facebook:"cross_share_to_group_ids"`
	CrosspostingActions          *[]map[string]interface{}                      `facebook:"crossposting_actions"`
	CustomLabels                 *[]string                                      `facebook:"custom_labels"`
	Description                  *string                                        `facebook:"description"`
	DirectShareStatus            *uint64                                        `facebook:"direct_share_status"`
	Embeddable                   *bool                                          `facebook:"embeddable"`
	EndLiveVideo                 *bool                                          `facebook:"end_live_video"`
	EventParams                  *map[string]interface{}                        `facebook:"event_params"`
	IsAudioOnly                  *bool                                          `facebook:"is_audio_only"`
	IsManualMode                 *bool                                          `facebook:"is_manual_mode"`
	LiveCommentModerationSetting *[]enums.LivevideoLiveCommentModerationSetting `facebook:"live_comment_moderation_setting"`
	MasterIngestStreamID         *core.ID                                       `facebook:"master_ingest_stream_id"`
	OgIconID                     *core.ID                                       `facebook:"og_icon_id"`
	OgPhrase                     *string                                        `facebook:"og_phrase"`
	PersistentStreamKeyStatus    *enums.LivevideoPersistentStreamKeyStatus      `facebook:"persistent_stream_key_status"`
	Place                        *map[string]interface{}                        `facebook:"place"`
	PlannedStartTime             *core.Time                                     `facebook:"planned_start_time"`
	Privacy                      *string                                        `facebook:"privacy"`
	Published                    *bool                                          `facebook:"published"`
	ScheduleCustomProfileImage   *core.FileParam                                `facebook:"schedule_custom_profile_image"`
	ScheduleFeedBackgroundImage  *core.FileParam                                `facebook:"schedule_feed_background_image"`
	SponsorID                    *core.ID                                       `facebook:"sponsor_id"`
	SponsorRelationship          *uint64                                        `facebook:"sponsor_relationship"`
	Status                       *enums.LivevideoStatus                         `facebook:"status"`
	StreamType                   *enums.LivevideoStreamType                     `facebook:"stream_type"`
	Tags                         *[]int                                         `facebook:"tags"`
	Targeting                    *map[string]interface{}                        `facebook:"targeting"`
	Title                        *string                                        `facebook:"title"`
	Extra                        core.Params                                    `facebook:"-"`
}

func (params UpdateLiveVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowBmCrossposting != nil {
		out["allow_bm_crossposting"] = *params.AllowBmCrossposting
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.CrossShareToGroupIds != nil {
		out["cross_share_to_group_ids"] = *params.CrossShareToGroupIds
	}
	if params.CrosspostingActions != nil {
		out["crossposting_actions"] = *params.CrosspostingActions
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
	if params.EndLiveVideo != nil {
		out["end_live_video"] = *params.EndLiveVideo
	}
	if params.EventParams != nil {
		out["event_params"] = *params.EventParams
	}
	if params.IsAudioOnly != nil {
		out["is_audio_only"] = *params.IsAudioOnly
	}
	if params.IsManualMode != nil {
		out["is_manual_mode"] = *params.IsManualMode
	}
	if params.LiveCommentModerationSetting != nil {
		out["live_comment_moderation_setting"] = *params.LiveCommentModerationSetting
	}
	if params.MasterIngestStreamID != nil {
		out["master_ingest_stream_id"] = *params.MasterIngestStreamID
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.PersistentStreamKeyStatus != nil {
		out["persistent_stream_key_status"] = *params.PersistentStreamKeyStatus
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.PlannedStartTime != nil {
		out["planned_start_time"] = *params.PlannedStartTime
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduleCustomProfileImage != nil {
		out["schedule_custom_profile_image"] = *params.ScheduleCustomProfileImage
	}
	if params.ScheduleFeedBackgroundImage != nil {
		out["schedule_feed_background_image"] = *params.ScheduleFeedBackgroundImage
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StreamType != nil {
		out["stream_type"] = *params.StreamType
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func UpdateLiveVideoBatchCall(id string, params UpdateLiveVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateLiveVideoBatchRequest(id string, params UpdateLiveVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](UpdateLiveVideoBatchCall(id, params, options...))
}

func DecodeUpdateLiveVideoBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateLiveVideoWithResponse(ctx context.Context, client *core.Client, id string, params UpdateLiveVideoParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := UpdateLiveVideoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateLiveVideo(ctx context.Context, client *core.Client, id string, params UpdateLiveVideoParams) (*objects.LiveVideo, error) {
	out, _, err := UpdateLiveVideoWithResponse(ctx, client, id, params)
	return out, err
}
