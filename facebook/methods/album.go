package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAlbumCommentsParams struct {
	Filter     *enums.AlbumcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.AlbumcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.AlbumcommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                              `facebook:"since"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetAlbumCommentsParams) ToParams() core.Params {
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

func GetAlbumCommentsBatchCall(id string, params GetAlbumCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetAlbumCommentsBatchRequest(id string, params GetAlbumCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetAlbumCommentsBatchCall(id, params, options...))
}

func DecodeGetAlbumCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetAlbumCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetAlbumCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetAlbumCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAlbumComments(ctx context.Context, client *core.Client, id string, params GetAlbumCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetAlbumCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAlbumCommentsParams struct {
	AttachmentID         *core.ID                                         `facebook:"attachment_id"`
	AttachmentShareURL   *string                                          `facebook:"attachment_share_url"`
	AttachmentURL        *string                                          `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.AlbumcommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
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

func (params CreateAlbumCommentsParams) ToParams() core.Params {
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

func CreateAlbumCommentsBatchCall(id string, params CreateAlbumCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateAlbumCommentsBatchRequest(id string, params CreateAlbumCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreateAlbumCommentsBatchCall(id, params, options...))
}

func DecodeCreateAlbumCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreateAlbumCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAlbumCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreateAlbumCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAlbumComments(ctx context.Context, client *core.Client, id string, params CreateAlbumCommentsParams) (*objects.Comment, error) {
	out, _, err := CreateAlbumCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAlbumLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAlbumLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAlbumLikesBatchCall(id string, params GetAlbumLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetAlbumLikesBatchRequest(id string, params GetAlbumLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetAlbumLikesBatchCall(id, params, options...))
}

func DecodeGetAlbumLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetAlbumLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetAlbumLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetAlbumLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAlbumLikes(ctx context.Context, client *core.Client, id string, params GetAlbumLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetAlbumLikesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAlbumLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateAlbumLikesParams) ToParams() core.Params {
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

func CreateAlbumLikesBatchCall(id string, params CreateAlbumLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreateAlbumLikesBatchRequest(id string, params CreateAlbumLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.Album] {
	return core.NewBatchRequest[objects.Album](CreateAlbumLikesBatchCall(id, params, options...))
}

func DecodeCreateAlbumLikesBatchResponse(response *core.BatchResponse) (*objects.Album, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Album
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAlbumLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreateAlbumLikesParams) (*objects.Album, *core.Response, error) {
	var out objects.Album
	call := CreateAlbumLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAlbumLikes(ctx context.Context, client *core.Client, id string, params CreateAlbumLikesParams) (*objects.Album, error) {
	out, _, err := CreateAlbumLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAlbumPhotosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAlbumPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAlbumPhotosBatchCall(id string, params GetAlbumPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewGetAlbumPhotosBatchRequest(id string, params GetAlbumPhotosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Photo]] {
	return core.NewBatchRequest[core.Cursor[objects.Photo]](GetAlbumPhotosBatchCall(id, params, options...))
}

func DecodeGetAlbumPhotosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Photo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Photo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAlbumPhotosWithResponse(ctx context.Context, client *core.Client, id string, params GetAlbumPhotosParams) (*core.Cursor[objects.Photo], *core.Response, error) {
	var out core.Cursor[objects.Photo]
	call := GetAlbumPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAlbumPhotos(ctx context.Context, client *core.Client, id string, params GetAlbumPhotosParams) (*core.Cursor[objects.Photo], error) {
	out, _, err := GetAlbumPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAlbumPhotosParams struct {
	Aid                                   *string                                             `facebook:"aid"`
	AllowSphericalPhoto                   *bool                                               `facebook:"allow_spherical_photo"`
	AltTextCustom                         *string                                             `facebook:"alt_text_custom"`
	AndroidKeyHash                        *string                                             `facebook:"android_key_hash"`
	ApplicationID                         *core.ID                                            `facebook:"application_id"`
	Attempt                               *uint64                                             `facebook:"attempt"`
	AudienceExp                           *bool                                               `facebook:"audience_exp"`
	BackdatedTime                         *core.Time                                          `facebook:"backdated_time"`
	BackdatedTimeGranularity              *enums.AlbumphotosBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	Caption                               *string                                             `facebook:"caption"`
	ComposerSessionID                     *core.ID                                            `facebook:"composer_session_id"`
	DirectShareStatus                     *uint64                                             `facebook:"direct_share_status"`
	FeedTargeting                         *map[string]interface{}                             `facebook:"feed_targeting"`
	FilterType                            *uint64                                             `facebook:"filter_type"`
	FullResIsComingLater                  *bool                                               `facebook:"full_res_is_coming_later"`
	InitialViewHeadingOverrideDegrees     *uint64                                             `facebook:"initial_view_heading_override_degrees"`
	InitialViewPitchOverrideDegrees       *uint64                                             `facebook:"initial_view_pitch_override_degrees"`
	InitialViewVerticalFovOverrideDegrees *uint64                                             `facebook:"initial_view_vertical_fov_override_degrees"`
	IosBundleID                           *core.ID                                            `facebook:"ios_bundle_id"`
	IsExplicitLocation                    *bool                                               `facebook:"is_explicit_location"`
	IsExplicitPlace                       *bool                                               `facebook:"is_explicit_place"`
	ManualPrivacy                         *bool                                               `facebook:"manual_privacy"`
	Message                               *string                                             `facebook:"message"`
	Name                                  *string                                             `facebook:"name"`
	NoStory                               *bool                                               `facebook:"no_story"`
	OfflineID                             *core.ID                                            `facebook:"offline_id"`
	OgActionTypeID                        *core.ID                                            `facebook:"og_action_type_id"`
	OgIconID                              *core.ID                                            `facebook:"og_icon_id"`
	OgObjectID                            *core.ID                                            `facebook:"og_object_id"`
	OgPhrase                              *string                                             `facebook:"og_phrase"`
	OgSetProfileBadge                     *bool                                               `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism                 *string                                             `facebook:"og_suggestion_mechanism"`
	Place                                 *map[string]interface{}                             `facebook:"place"`
	Privacy                               *string                                             `facebook:"privacy"`
	ProfileID                             *core.ID                                            `facebook:"profile_id"`
	ProvenanceInfo                        *map[string]interface{}                             `facebook:"provenance_info"`
	ProxiedAppID                          *core.ID                                            `facebook:"proxied_app_id"`
	Published                             *bool                                               `facebook:"published"`
	Qn                                    *string                                             `facebook:"qn"`
	SphericalMetadata                     *map[string]interface{}                             `facebook:"spherical_metadata"`
	SponsorID                             *core.ID                                            `facebook:"sponsor_id"`
	SponsorRelationship                   *uint64                                             `facebook:"sponsor_relationship"`
	Tags                                  *[]map[string]interface{}                           `facebook:"tags"`
	TargetID                              *core.ID                                            `facebook:"target_id"`
	Targeting                             *map[string]interface{}                             `facebook:"targeting"`
	TimeSinceOriginalPost                 *uint64                                             `facebook:"time_since_original_post"`
	UID                                   *core.ID                                            `facebook:"uid"`
	UnpublishedContentType                *enums.AlbumphotosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	URL                                   *string                                             `facebook:"url"`
	UserSelectedTags                      *bool                                               `facebook:"user_selected_tags"`
	VaultImageID                          *core.ID                                            `facebook:"vault_image_id"`
	Extra                                 core.Params                                         `facebook:"-"`
}

func (params CreateAlbumPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Aid != nil {
		out["aid"] = *params.Aid
	}
	if params.AllowSphericalPhoto != nil {
		out["allow_spherical_photo"] = *params.AllowSphericalPhoto
	}
	if params.AltTextCustom != nil {
		out["alt_text_custom"] = *params.AltTextCustom
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.Attempt != nil {
		out["attempt"] = *params.Attempt
	}
	if params.AudienceExp != nil {
		out["audience_exp"] = *params.AudienceExp
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.FilterType != nil {
		out["filter_type"] = *params.FilterType
	}
	if params.FullResIsComingLater != nil {
		out["full_res_is_coming_later"] = *params.FullResIsComingLater
	}
	if params.InitialViewHeadingOverrideDegrees != nil {
		out["initial_view_heading_override_degrees"] = *params.InitialViewHeadingOverrideDegrees
	}
	if params.InitialViewPitchOverrideDegrees != nil {
		out["initial_view_pitch_override_degrees"] = *params.InitialViewPitchOverrideDegrees
	}
	if params.InitialViewVerticalFovOverrideDegrees != nil {
		out["initial_view_vertical_fov_override_degrees"] = *params.InitialViewVerticalFovOverrideDegrees
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitPlace != nil {
		out["is_explicit_place"] = *params.IsExplicitPlace
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
	}
	if params.OfflineID != nil {
		out["offline_id"] = *params.OfflineID
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgObjectID != nil {
		out["og_object_id"] = *params.OgObjectID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.OgSetProfileBadge != nil {
		out["og_set_profile_badge"] = *params.OgSetProfileBadge
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.ProfileID != nil {
		out["profile_id"] = *params.ProfileID
	}
	if params.ProvenanceInfo != nil {
		out["provenance_info"] = *params.ProvenanceInfo
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Qn != nil {
		out["qn"] = *params.Qn
	}
	if params.SphericalMetadata != nil {
		out["spherical_metadata"] = *params.SphericalMetadata
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
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VaultImageID != nil {
		out["vault_image_id"] = *params.VaultImageID
	}
	return out
}

func CreateAlbumPhotosBatchCall(id string, params CreateAlbumPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewCreateAlbumPhotosBatchRequest(id string, params CreateAlbumPhotosParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](CreateAlbumPhotosBatchCall(id, params, options...))
}

func DecodeCreateAlbumPhotosBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
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

func CreateAlbumPhotosWithResponse(ctx context.Context, client *core.Client, id string, params CreateAlbumPhotosParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := CreateAlbumPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAlbumPhotos(ctx context.Context, client *core.Client, id string, params CreateAlbumPhotosParams) (*objects.Photo, error) {
	out, _, err := CreateAlbumPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type GetAlbumPictureParams struct {
	Redirect *bool                            `facebook:"redirect"`
	Type     *enums.AlbumpictureTypeEnumParam `facebook:"type"`
	Extra    core.Params                      `facebook:"-"`
}

func (params GetAlbumPictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Redirect != nil {
		out["redirect"] = *params.Redirect
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetAlbumPictureBatchCall(id string, params GetAlbumPictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetAlbumPictureBatchRequest(id string, params GetAlbumPictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetAlbumPictureBatchCall(id, params, options...))
}

func DecodeGetAlbumPictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProfilePictureSource]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAlbumPictureWithResponse(ctx context.Context, client *core.Client, id string, params GetAlbumPictureParams) (*core.Cursor[objects.ProfilePictureSource], *core.Response, error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetAlbumPictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAlbumPicture(ctx context.Context, client *core.Client, id string, params GetAlbumPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	out, _, err := GetAlbumPictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetAlbumParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAlbumParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAlbumBatchCall(id string, params GetAlbumParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAlbumBatchRequest(id string, params GetAlbumParams, options ...core.BatchOption) *core.BatchRequest[objects.Album] {
	return core.NewBatchRequest[objects.Album](GetAlbumBatchCall(id, params, options...))
}

func DecodeGetAlbumBatchResponse(response *core.BatchResponse) (*objects.Album, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Album
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAlbumWithResponse(ctx context.Context, client *core.Client, id string, params GetAlbumParams) (*objects.Album, *core.Response, error) {
	var out objects.Album
	call := GetAlbumBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAlbum(ctx context.Context, client *core.Client, id string, params GetAlbumParams) (*objects.Album, error) {
	out, _, err := GetAlbumWithResponse(ctx, client, id, params)
	return out, err
}
