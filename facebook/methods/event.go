package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetEventComments(ctx context.Context, client *core.Client, id string, params GetEventCommentsParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
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

func GetEventFeed(ctx context.Context, client *core.Client, id string, params GetEventFeedParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), &out)
	return &out, err
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

func GetEventLiveVideos(ctx context.Context, client *core.Client, id string, params GetEventLiveVideosParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), &out)
	return &out, err
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

func CreateEventLiveVideos(ctx context.Context, client *core.Client, id string, params CreateEventLiveVideosParams) (*objects.LiveVideo, error) {
	var out objects.LiveVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), &out)
	return &out, err
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

func GetEventPhotos(ctx context.Context, client *core.Client, id string, params GetEventPhotosParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), &out)
	return &out, err
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

func GetEventPicture(ctx context.Context, client *core.Client, id string, params GetEventPictureParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), &out)
	return &out, err
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

func GetEventPosts(ctx context.Context, client *core.Client, id string, params GetEventPostsParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "posts"), params.ToParams(), &out)
	return &out, err
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

func GetEventRoles(ctx context.Context, client *core.Client, id string, params GetEventRolesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "roles"), params.ToParams(), &out)
	return &out, err
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

func GetEventTicketTiers(ctx context.Context, client *core.Client, id string, params GetEventTicketTiersParams) (*core.Cursor[objects.EventTicketTier], error) {
	var out core.Cursor[objects.EventTicketTier]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ticket_tiers"), params.ToParams(), &out)
	return &out, err
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

func GetEventVideos(ctx context.Context, client *core.Client, id string, params GetEventVideosParams) (*core.Cursor[objects.NullNode], error) {
	var out core.Cursor[objects.NullNode]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), &out)
	return &out, err
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

func GetEvent(ctx context.Context, client *core.Client, id string, params GetEventParams) (*objects.Event, error) {
	var out objects.Event
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
