package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventCommentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventCommentsBatchCall(id string, params GetEventCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetEventCommentsBatchRequest(id string, params GetEventCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventCommentsBatchCall(id, params, options...))
}

func DecodeGetEventCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetEventCommentsParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventComments(ctx context.Context, client *core.Client, id string, params GetEventCommentsParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventFeedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventFeedBatchCall(id string, params GetEventFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewGetEventFeedBatchRequest(id string, params GetEventFeedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventFeedBatchCall(id, params, options...))
}

func DecodeGetEventFeedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventFeedWithResponse(ctx context.Context, client *core.Client, id string, params GetEventFeedParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventFeed(ctx context.Context, client *core.Client, id string, params GetEventFeedParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventFeedWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventLiveVideosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventLiveVideosBatchCall(id string, params GetEventLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewGetEventLiveVideosBatchRequest(id string, params GetEventLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventLiveVideosBatchCall(id, params, options...))
}

func DecodeGetEventLiveVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetEventLiveVideosParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventLiveVideos(ctx context.Context, client *core.Client, id string, params GetEventLiveVideosParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateEventLiveVideosParams struct {
	ContentTags                *[]string                                         `facebook:"content_tags"`
	Description                *string                                           `facebook:"description"`
	EnableBackupIngest         *bool                                             `facebook:"enable_backup_ingest"`
	EncodingSettings           *string                                           `facebook:"encoding_settings"`
	EventParams                *map[string]interface{}                           `facebook:"event_params"`
	FisheyeVideoCropped        *bool                                             `facebook:"fisheye_video_cropped"`
	FrontZRotation             *float64                                          `facebook:"front_z_rotation"`
	IsAudioOnly                *bool                                             `facebook:"is_audio_only"`
	IsSpherical                *bool                                             `facebook:"is_spherical"`
	OriginalFov                *uint64                                           `facebook:"original_fov"`
	Privacy                    *string                                           `facebook:"privacy"`
	Projection                 *enums.EventliveVideosProjectionEnumParam         `facebook:"projection"`
	Published                  *bool                                             `facebook:"published"`
	ScheduleCustomProfileImage *core.FileParam                                   `facebook:"schedule_custom_profile_image"`
	SpatialAudioFormat         *enums.EventliveVideosSpatialAudioFormatEnumParam `facebook:"spatial_audio_format"`
	Status                     *enums.EventliveVideosStatusEnumParam             `facebook:"status"`
	StereoscopicMode           *enums.EventliveVideosStereoscopicModeEnumParam   `facebook:"stereoscopic_mode"`
	StopOnDeleteStream         *bool                                             `facebook:"stop_on_delete_stream"`
	StreamType                 *enums.EventliveVideosStreamTypeEnumParam         `facebook:"stream_type"`
	Title                      *string                                           `facebook:"title"`
	Extra                      core.Params                                       `facebook:"-"`
}

func (params CreateEventLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableBackupIngest != nil {
		out["enable_backup_ingest"] = *params.EnableBackupIngest
	}
	if params.EncodingSettings != nil {
		out["encoding_settings"] = *params.EncodingSettings
	}
	if params.EventParams != nil {
		out["event_params"] = *params.EventParams
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.IsAudioOnly != nil {
		out["is_audio_only"] = *params.IsAudioOnly
	}
	if params.IsSpherical != nil {
		out["is_spherical"] = *params.IsSpherical
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Projection != nil {
		out["projection"] = *params.Projection
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduleCustomProfileImage != nil {
		out["schedule_custom_profile_image"] = *params.ScheduleCustomProfileImage
	}
	if params.SpatialAudioFormat != nil {
		out["spatial_audio_format"] = *params.SpatialAudioFormat
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StereoscopicMode != nil {
		out["stereoscopic_mode"] = *params.StereoscopicMode
	}
	if params.StopOnDeleteStream != nil {
		out["stop_on_delete_stream"] = *params.StopOnDeleteStream
	}
	if params.StreamType != nil {
		out["stream_type"] = *params.StreamType
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func CreateEventLiveVideosBatchCall(id string, params CreateEventLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewCreateEventLiveVideosBatchRequest(id string, params CreateEventLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](CreateEventLiveVideosBatchCall(id, params, options...))
}

func DecodeCreateEventLiveVideosBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
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

func CreateEventLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreateEventLiveVideosParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := CreateEventLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateEventLiveVideos(ctx context.Context, client *core.Client, id string, params CreateEventLiveVideosParams) (*objects.LiveVideo, error) {
	out, _, err := CreateEventLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventPhotosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventPhotosBatchCall(id string, params GetEventPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewGetEventPhotosBatchRequest(id string, params GetEventPhotosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventPhotosBatchCall(id, params, options...))
}

func DecodeGetEventPhotosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventPhotosWithResponse(ctx context.Context, client *core.Client, id string, params GetEventPhotosParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventPhotos(ctx context.Context, client *core.Client, id string, params GetEventPhotosParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventPictureParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventPictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventPictureBatchCall(id string, params GetEventPictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetEventPictureBatchRequest(id string, params GetEventPictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventPictureBatchCall(id, params, options...))
}

func DecodeGetEventPictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventPictureWithResponse(ctx context.Context, client *core.Client, id string, params GetEventPictureParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventPictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventPicture(ctx context.Context, client *core.Client, id string, params GetEventPictureParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventPictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventPostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventPostsBatchCall(id string, params GetEventPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "posts"), params.ToParams(), options...)
}

func NewGetEventPostsBatchRequest(id string, params GetEventPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventPostsBatchCall(id, params, options...))
}

func DecodeGetEventPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetEventPostsParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventPosts(ctx context.Context, client *core.Client, id string, params GetEventPostsParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventPostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventRolesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventRolesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventRolesBatchCall(id string, params GetEventRolesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "roles"), params.ToParams(), options...)
}

func NewGetEventRolesBatchRequest(id string, params GetEventRolesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetEventRolesBatchCall(id, params, options...))
}

func DecodeGetEventRolesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetEventRolesWithResponse(ctx context.Context, client *core.Client, id string, params GetEventRolesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetEventRolesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventRoles(ctx context.Context, client *core.Client, id string, params GetEventRolesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetEventRolesWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventTicketTiersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventTicketTiersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventTicketTiersBatchCall(id string, params GetEventTicketTiersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ticket_tiers"), params.ToParams(), options...)
}

func NewGetEventTicketTiersBatchRequest(id string, params GetEventTicketTiersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.EventTicketTier]] {
	return core.NewBatchRequest[core.Cursor[objects.EventTicketTier]](GetEventTicketTiersBatchCall(id, params, options...))
}

func DecodeGetEventTicketTiersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.EventTicketTier], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.EventTicketTier]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventTicketTiersWithResponse(ctx context.Context, client *core.Client, id string, params GetEventTicketTiersParams) (*core.Cursor[objects.EventTicketTier], *core.Response, error) {
	var out core.Cursor[objects.EventTicketTier]
	call := GetEventTicketTiersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventTicketTiers(ctx context.Context, client *core.Client, id string, params GetEventTicketTiersParams) (*core.Cursor[objects.EventTicketTier], error) {
	out, _, err := GetEventTicketTiersWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventVideosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventVideosBatchCall(id string, params GetEventVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewGetEventVideosBatchRequest(id string, params GetEventVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.NullNode]] {
	return core.NewBatchRequest[core.Cursor[objects.NullNode]](GetEventVideosBatchCall(id, params, options...))
}

func DecodeGetEventVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.NullNode], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.NullNode]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetEventVideosParams) (*core.Cursor[objects.NullNode], *core.Response, error) {
	var out core.Cursor[objects.NullNode]
	call := GetEventVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventVideos(ctx context.Context, client *core.Client, id string, params GetEventVideosParams) (*core.Cursor[objects.NullNode], error) {
	out, _, err := GetEventVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventBatchCall(id string, params GetEventParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventBatchRequest(id string, params GetEventParams, options ...core.BatchOption) *core.BatchRequest[objects.Event] {
	return core.NewBatchRequest[objects.Event](GetEventBatchCall(id, params, options...))
}

func DecodeGetEventBatchResponse(response *core.BatchResponse) (*objects.Event, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Event
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventWithResponse(ctx context.Context, client *core.Client, id string, params GetEventParams) (*objects.Event, *core.Response, error) {
	var out objects.Event
	call := GetEventBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEvent(ctx context.Context, client *core.Client, id string, params GetEventParams) (*objects.Event, error) {
	out, _, err := GetEventWithResponse(ctx, client, id, params)
	return out, err
}
