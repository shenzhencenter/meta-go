package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
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

func GetLiveVideoBlockedUsers(ctx context.Context, client *core.Client, id string, params GetLiveVideoBlockedUsersParams) (*core.Cursor[objects.User], error) {
	var out core.Cursor[objects.User]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "blocked_users"), params.ToParams(), &out)
	return &out, err
}

type GetLiveVideoCommentsParams struct {
	Filter     *enums.LivevideocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.LivevideocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.LivevideocommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                                  `facebook:"since"`
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

func GetLiveVideoComments(ctx context.Context, client *core.Client, id string, params GetLiveVideoCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideoCrosspostSharedPages(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "crosspost_shared_pages"), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideoCrosspostedBroadcasts(ctx context.Context, client *core.Client, id string, params GetLiveVideoCrosspostedBroadcastsParams) (*core.Cursor[objects.LiveVideo], error) {
	var out core.Cursor[objects.LiveVideo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "crossposted_broadcasts"), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideoErrors(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorsParams) (*core.Cursor[objects.LiveVideoError], error) {
	var out core.Cursor[objects.LiveVideoError]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "errors"), params.ToParams(), &out)
	return &out, err
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

func CreateLiveVideoInputStreams(ctx context.Context, client *core.Client, id string, params CreateLiveVideoInputStreamsParams) (*objects.LiveVideoInputStream, error) {
	var out objects.LiveVideoInputStream
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "input_streams"), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideoPolls(ctx context.Context, client *core.Client, id string, params GetLiveVideoPollsParams) (*core.Cursor[objects.VideoPoll], error) {
	var out core.Cursor[objects.VideoPoll]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "polls"), params.ToParams(), &out)
	return &out, err
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

func CreateLiveVideoPolls(ctx context.Context, client *core.Client, id string, params CreateLiveVideoPollsParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "polls"), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideoReactions(ctx context.Context, client *core.Client, id string, params GetLiveVideoReactionsParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), &out)
	return &out, err
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

func DeleteLiveVideo(ctx context.Context, client *core.Client, id string, params DeleteLiveVideoParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetLiveVideo(ctx context.Context, client *core.Client, id string, params GetLiveVideoParams) (*objects.LiveVideo, error) {
	var out objects.LiveVideo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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
	PlannedStartTime             *time.Time                                     `facebook:"planned_start_time"`
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

func UpdateLiveVideo(ctx context.Context, client *core.Client, id string, params UpdateLiveVideoParams) (*objects.LiveVideo, error) {
	var out objects.LiveVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
