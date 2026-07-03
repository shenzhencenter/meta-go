package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteGroupAdminsParams struct {
	UID   core.ID     `facebook:"uid"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteGroupAdminsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["uid"] = params.UID
	return out
}

func DeleteGroupAdminsBatchCall(id string, params DeleteGroupAdminsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "admins"), params.ToParams(), options...)
}

func NewDeleteGroupAdminsBatchRequest(id string, params DeleteGroupAdminsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteGroupAdminsBatchCall(id, params, options...))
}

func DecodeDeleteGroupAdminsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteGroupAdminsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteGroupAdminsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteGroupAdminsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteGroupAdmins(ctx context.Context, client *core.Client, id string, params DeleteGroupAdminsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteGroupAdminsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupAdminsParams struct {
	UID   core.ID     `facebook:"uid"`
	Extra core.Params `facebook:"-"`
}

func (params CreateGroupAdminsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["uid"] = params.UID
	return out
}

func CreateGroupAdminsBatchCall(id string, params CreateGroupAdminsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "admins"), params.ToParams(), options...)
}

func NewCreateGroupAdminsBatchRequest(id string, params CreateGroupAdminsParams, options ...core.BatchOption) *core.BatchRequest[objects.Group] {
	return core.NewBatchRequest[objects.Group](CreateGroupAdminsBatchCall(id, params, options...))
}

func DecodeCreateGroupAdminsBatchResponse(response *core.BatchResponse) (*objects.Group, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Group
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateGroupAdminsWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupAdminsParams) (*objects.Group, *core.Response, error) {
	var out objects.Group
	call := CreateGroupAdminsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupAdmins(ctx context.Context, client *core.Client, id string, params CreateGroupAdminsParams) (*objects.Group, error) {
	out, _, err := CreateGroupAdminsWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupAlbumsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupAlbumsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupAlbumsBatchCall(id string, params GetGroupAlbumsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "albums"), params.ToParams(), options...)
}

func NewGetGroupAlbumsBatchRequest(id string, params GetGroupAlbumsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Album]] {
	return core.NewBatchRequest[core.Cursor[objects.Album]](GetGroupAlbumsBatchCall(id, params, options...))
}

func DecodeGetGroupAlbumsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Album], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Album]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupAlbumsWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupAlbumsParams) (*core.Cursor[objects.Album], *core.Response, error) {
	var out core.Cursor[objects.Album]
	call := GetGroupAlbumsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupAlbums(ctx context.Context, client *core.Client, id string, params GetGroupAlbumsParams) (*core.Cursor[objects.Album], error) {
	out, _, err := GetGroupAlbumsWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupDocsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupDocsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupDocsBatchCall(id string, params GetGroupDocsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "docs"), params.ToParams(), options...)
}

func NewGetGroupDocsBatchRequest(id string, params GetGroupDocsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Doc]] {
	return core.NewBatchRequest[core.Cursor[objects.Doc]](GetGroupDocsBatchCall(id, params, options...))
}

func DecodeGetGroupDocsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Doc], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Doc]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupDocsWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupDocsParams) (*core.Cursor[objects.Doc], *core.Response, error) {
	var out core.Cursor[objects.Doc]
	call := GetGroupDocsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupDocs(ctx context.Context, client *core.Client, id string, params GetGroupDocsParams) (*core.Cursor[objects.Doc], error) {
	out, _, err := GetGroupDocsWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupEventsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupEventsBatchCall(id string, params GetGroupEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "events"), params.ToParams(), options...)
}

func NewGetGroupEventsBatchRequest(id string, params GetGroupEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Event]] {
	return core.NewBatchRequest[core.Cursor[objects.Event]](GetGroupEventsBatchCall(id, params, options...))
}

func DecodeGetGroupEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Event], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Event]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupEventsParams) (*core.Cursor[objects.Event], *core.Response, error) {
	var out core.Cursor[objects.Event]
	call := GetGroupEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupEvents(ctx context.Context, client *core.Client, id string, params GetGroupEventsParams) (*core.Cursor[objects.Event], error) {
	out, _, err := GetGroupEventsWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupFeedParams struct {
	IncludeHidden *bool       `facebook:"include_hidden"`
	Q             *string     `facebook:"q"`
	ShowExpired   *bool       `facebook:"show_expired"`
	Since         *core.Time  `facebook:"since"`
	Until         *core.Time  `facebook:"until"`
	With          *string     `facebook:"with"`
	Extra         core.Params `facebook:"-"`
}

func (params GetGroupFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Q != nil {
		out["q"] = *params.Q
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetGroupFeedBatchCall(id string, params GetGroupFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewGetGroupFeedBatchRequest(id string, params GetGroupFeedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Post]] {
	return core.NewBatchRequest[core.Cursor[objects.Post]](GetGroupFeedBatchCall(id, params, options...))
}

func DecodeGetGroupFeedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Post], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Post]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupFeedWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupFeedParams) (*core.Cursor[objects.Post], *core.Response, error) {
	var out core.Cursor[objects.Post]
	call := GetGroupFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupFeed(ctx context.Context, client *core.Client, id string, params GetGroupFeedParams) (*core.Cursor[objects.Post], error) {
	out, _, err := GetGroupFeedWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupFeedParams struct {
	Actions                   *map[string]interface{}                           `facebook:"actions"`
	AlbumID                   *core.ID                                          `facebook:"album_id"`
	AndroidKeyHash            *string                                           `facebook:"android_key_hash"`
	ApplicationID             *core.ID                                          `facebook:"application_id"`
	AskedFunFactPromptID      *core.ID                                          `facebook:"asked_fun_fact_prompt_id"`
	Asset3dID                 *core.ID                                          `facebook:"asset3d_id"`
	AssociatedID              *core.ID                                          `facebook:"associated_id"`
	AttachPlaceSuggestion     *bool                                             `facebook:"attach_place_suggestion"`
	AttachedMedia             *[]map[string]interface{}                         `facebook:"attached_media"`
	AudienceExp               *bool                                             `facebook:"audience_exp"`
	BackdatedTime             *core.Time                                        `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.GroupfeedBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	BreakingNews              *bool                                             `facebook:"breaking_news"`
	BreakingNewsExpiration    *uint64                                           `facebook:"breaking_news_expiration"`
	CallToAction              *map[string]interface{}                           `facebook:"call_to_action"`
	Caption                   *string                                           `facebook:"caption"`
	ChildAttachments          *[]map[string]interface{}                         `facebook:"child_attachments"`
	ClientMutationID          *core.ID                                          `facebook:"client_mutation_id"`
	ComposerEntryPicker       *string                                           `facebook:"composer_entry_picker"`
	ComposerEntryPoint        *string                                           `facebook:"composer_entry_point"`
	ComposerEntryTime         *uint64                                           `facebook:"composer_entry_time"`
	ComposerSessionEventsLog  *string                                           `facebook:"composer_session_events_log"`
	ComposerSessionID         *core.ID                                          `facebook:"composer_session_id"`
	ComposerSourceSurface     *string                                           `facebook:"composer_source_surface"`
	ComposerType              *string                                           `facebook:"composer_type"`
	ConnectionClass           *string                                           `facebook:"connection_class"`
	ContentAttachment         *string                                           `facebook:"content_attachment"`
	Coordinates               *map[string]interface{}                           `facebook:"coordinates"`
	CtaLink                   *string                                           `facebook:"cta_link"`
	CtaType                   *string                                           `facebook:"cta_type"`
	Description               *string                                           `facebook:"description"`
	DirectShareStatus         *uint64                                           `facebook:"direct_share_status"`
	ExpandedHeight            *uint64                                           `facebook:"expanded_height"`
	ExpandedWidth             *uint64                                           `facebook:"expanded_width"`
	FeedTargeting             *map[string]interface{}                           `facebook:"feed_targeting"`
	Formatting                *enums.GroupfeedFormattingEnumParam               `facebook:"formatting"`
	FunFactPromptID           *core.ID                                          `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID          *core.ID                                          `facebook:"fun_fact_toastee_id"`
	Height                    *uint64                                           `facebook:"height"`
	HomeCheckinCityID         *map[string]interface{}                           `facebook:"home_checkin_city_id"`
	ImageCrops                *map[string]interface{}                           `facebook:"image_crops"`
	ImplicitWithTags          *[]int                                            `facebook:"implicit_with_tags"`
	InstantGameEntryPointData *string                                           `facebook:"instant_game_entry_point_data"`
	IosBundleID               *core.ID                                          `facebook:"ios_bundle_id"`
	IsBackoutDraft            *bool                                             `facebook:"is_backout_draft"`
	IsBoostIntended           *bool                                             `facebook:"is_boost_intended"`
	IsExplicitLocation        *bool                                             `facebook:"is_explicit_location"`
	IsExplicitShare           *bool                                             `facebook:"is_explicit_share"`
	IsGroupLinkingPost        *bool                                             `facebook:"is_group_linking_post"`
	IsPhotoContainer          *bool                                             `facebook:"is_photo_container"`
	Link                      *string                                           `facebook:"link"`
	LocationSourceID          *core.ID                                          `facebook:"location_source_id"`
	ManualPrivacy             *bool                                             `facebook:"manual_privacy"`
	Message                   *string                                           `facebook:"message"`
	MultiShareEndCard         *bool                                             `facebook:"multi_share_end_card"`
	MultiShareOptimized       *bool                                             `facebook:"multi_share_optimized"`
	Name                      *string                                           `facebook:"name"`
	NectarModule              *string                                           `facebook:"nectar_module"`
	ObjectAttachment          *string                                           `facebook:"object_attachment"`
	OgActionTypeID            *core.ID                                          `facebook:"og_action_type_id"`
	OgHideObjectAttachment    *bool                                             `facebook:"og_hide_object_attachment"`
	OgIconID                  *core.ID                                          `facebook:"og_icon_id"`
	OgObjectID                *core.ID                                          `facebook:"og_object_id"`
	OgPhrase                  *string                                           `facebook:"og_phrase"`
	OgSetProfileBadge         *bool                                             `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism     *string                                           `facebook:"og_suggestion_mechanism"`
	PageRecommendation        *string                                           `facebook:"page_recommendation"`
	Picture                   *string                                           `facebook:"picture"`
	Place                     *map[string]interface{}                           `facebook:"place"`
	PlaceAttachmentSetting    *enums.GroupfeedPlaceAttachmentSettingEnumParam   `facebook:"place_attachment_setting"`
	PlaceList                 *string                                           `facebook:"place_list"`
	PlaceListData             *[]interface{}                                    `facebook:"place_list_data"`
	PostSurfacesBlacklist     *[]enums.GroupfeedPostSurfacesBlacklistEnumParam  `facebook:"post_surfaces_blacklist"`
	PostingToRedspace         *enums.GroupfeedPostingToRedspaceEnumParam        `facebook:"posting_to_redspace"`
	Privacy                   *string                                           `facebook:"privacy"`
	PromptID                  *core.ID                                          `facebook:"prompt_id"`
	PromptTrackingString      *string                                           `facebook:"prompt_tracking_string"`
	Properties                *map[string]interface{}                           `facebook:"properties"`
	ProxiedAppID              *core.ID                                          `facebook:"proxied_app_id"`
	PublishEventID            *core.ID                                          `facebook:"publish_event_id"`
	Published                 *bool                                             `facebook:"published"`
	Quote                     *string                                           `facebook:"quote"`
	Ref                       *[]string                                         `facebook:"ref"`
	ReferenceableImageIds     *[]core.ID                                        `facebook:"referenceable_image_ids"`
	ReferralID                *core.ID                                          `facebook:"referral_id"`
	ScheduledPublishTime      *core.Time                                        `facebook:"scheduled_publish_time"`
	Source                    *string                                           `facebook:"source"`
	SponsorID                 *core.ID                                          `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                                           `facebook:"sponsor_relationship"`
	SuggestedPlaceID          *map[string]interface{}                           `facebook:"suggested_place_id"`
	Tags                      *[]int                                            `facebook:"tags"`
	TargetSurface             *enums.GroupfeedTargetSurfaceEnumParam            `facebook:"target_surface"`
	Targeting                 *map[string]interface{}                           `facebook:"targeting"`
	TextFormatMetadata        *string                                           `facebook:"text_format_metadata"`
	TextFormatPresetID        *core.ID                                          `facebook:"text_format_preset_id"`
	TextOnlyPlace             *string                                           `facebook:"text_only_place"`
	Thumbnail                 *core.FileParam                                   `facebook:"thumbnail"`
	TimeSinceOriginalPost     *uint64                                           `facebook:"time_since_original_post"`
	Title                     *string                                           `facebook:"title"`
	TrackingInfo              *string                                           `facebook:"tracking_info"`
	UnpublishedContentType    *enums.GroupfeedUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	UserSelectedTags          *bool                                             `facebook:"user_selected_tags"`
	VideoStartTimeMs          *uint64                                           `facebook:"video_start_time_ms"`
	ViewerCoordinates         *map[string]interface{}                           `facebook:"viewer_coordinates"`
	Width                     *uint64                                           `facebook:"width"`
	Extra                     core.Params                                       `facebook:"-"`
}

func (params CreateGroupFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Actions != nil {
		out["actions"] = *params.Actions
	}
	if params.AlbumID != nil {
		out["album_id"] = *params.AlbumID
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.Asset3dID != nil {
		out["asset3d_id"] = *params.Asset3dID
	}
	if params.AssociatedID != nil {
		out["associated_id"] = *params.AssociatedID
	}
	if params.AttachPlaceSuggestion != nil {
		out["attach_place_suggestion"] = *params.AttachPlaceSuggestion
	}
	if params.AttachedMedia != nil {
		out["attached_media"] = *params.AttachedMedia
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
	if params.BreakingNews != nil {
		out["breaking_news"] = *params.BreakingNews
	}
	if params.BreakingNewsExpiration != nil {
		out["breaking_news_expiration"] = *params.BreakingNewsExpiration
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ChildAttachments != nil {
		out["child_attachments"] = *params.ChildAttachments
	}
	if params.ClientMutationID != nil {
		out["client_mutation_id"] = *params.ClientMutationID
	}
	if params.ComposerEntryPicker != nil {
		out["composer_entry_picker"] = *params.ComposerEntryPicker
	}
	if params.ComposerEntryPoint != nil {
		out["composer_entry_point"] = *params.ComposerEntryPoint
	}
	if params.ComposerEntryTime != nil {
		out["composer_entry_time"] = *params.ComposerEntryTime
	}
	if params.ComposerSessionEventsLog != nil {
		out["composer_session_events_log"] = *params.ComposerSessionEventsLog
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.ComposerSourceSurface != nil {
		out["composer_source_surface"] = *params.ComposerSourceSurface
	}
	if params.ComposerType != nil {
		out["composer_type"] = *params.ComposerType
	}
	if params.ConnectionClass != nil {
		out["connection_class"] = *params.ConnectionClass
	}
	if params.ContentAttachment != nil {
		out["content_attachment"] = *params.ContentAttachment
	}
	if params.Coordinates != nil {
		out["coordinates"] = *params.Coordinates
	}
	if params.CtaLink != nil {
		out["cta_link"] = *params.CtaLink
	}
	if params.CtaType != nil {
		out["cta_type"] = *params.CtaType
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.ExpandedHeight != nil {
		out["expanded_height"] = *params.ExpandedHeight
	}
	if params.ExpandedWidth != nil {
		out["expanded_width"] = *params.ExpandedWidth
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.HomeCheckinCityID != nil {
		out["home_checkin_city_id"] = *params.HomeCheckinCityID
	}
	if params.ImageCrops != nil {
		out["image_crops"] = *params.ImageCrops
	}
	if params.ImplicitWithTags != nil {
		out["implicit_with_tags"] = *params.ImplicitWithTags
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsBackoutDraft != nil {
		out["is_backout_draft"] = *params.IsBackoutDraft
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPhotoContainer != nil {
		out["is_photo_container"] = *params.IsPhotoContainer
	}
	if params.Link != nil {
		out["link"] = *params.Link
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.MultiShareEndCard != nil {
		out["multi_share_end_card"] = *params.MultiShareEndCard
	}
	if params.MultiShareOptimized != nil {
		out["multi_share_optimized"] = *params.MultiShareOptimized
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectAttachment != nil {
		out["object_attachment"] = *params.ObjectAttachment
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgHideObjectAttachment != nil {
		out["og_hide_object_attachment"] = *params.OgHideObjectAttachment
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
	if params.PageRecommendation != nil {
		out["page_recommendation"] = *params.PageRecommendation
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.PlaceAttachmentSetting != nil {
		out["place_attachment_setting"] = *params.PlaceAttachmentSetting
	}
	if params.PlaceList != nil {
		out["place_list"] = *params.PlaceList
	}
	if params.PlaceListData != nil {
		out["place_list_data"] = *params.PlaceListData
	}
	if params.PostSurfacesBlacklist != nil {
		out["post_surfaces_blacklist"] = *params.PostSurfacesBlacklist
	}
	if params.PostingToRedspace != nil {
		out["posting_to_redspace"] = *params.PostingToRedspace
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PromptID != nil {
		out["prompt_id"] = *params.PromptID
	}
	if params.PromptTrackingString != nil {
		out["prompt_tracking_string"] = *params.PromptTrackingString
	}
	if params.Properties != nil {
		out["properties"] = *params.Properties
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Quote != nil {
		out["quote"] = *params.Quote
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	if params.ReferenceableImageIds != nil {
		out["referenceable_image_ids"] = *params.ReferenceableImageIds
	}
	if params.ReferralID != nil {
		out["referral_id"] = *params.ReferralID
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.SuggestedPlaceID != nil {
		out["suggested_place_id"] = *params.SuggestedPlaceID
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.TargetSurface != nil {
		out["target_surface"] = *params.TargetSurface
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.TextFormatPresetID != nil {
		out["text_format_preset_id"] = *params.TextFormatPresetID
	}
	if params.TextOnlyPlace != nil {
		out["text_only_place"] = *params.TextOnlyPlace
	}
	if params.Thumbnail != nil {
		out["thumbnail"] = *params.Thumbnail
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TrackingInfo != nil {
		out["tracking_info"] = *params.TrackingInfo
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.ViewerCoordinates != nil {
		out["viewer_coordinates"] = *params.ViewerCoordinates
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func CreateGroupFeedBatchCall(id string, params CreateGroupFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewCreateGroupFeedBatchRequest(id string, params CreateGroupFeedParams, options ...core.BatchOption) *core.BatchRequest[objects.Post] {
	return core.NewBatchRequest[objects.Post](CreateGroupFeedBatchCall(id, params, options...))
}

func DecodeCreateGroupFeedBatchResponse(response *core.BatchResponse) (*objects.Post, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Post
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateGroupFeedWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupFeedParams) (*objects.Post, *core.Response, error) {
	var out objects.Post
	call := CreateGroupFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupFeed(ctx context.Context, client *core.Client, id string, params CreateGroupFeedParams) (*objects.Post, error) {
	out, _, err := CreateGroupFeedWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupFilesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupFilesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupFilesBatchCall(id string, params GetGroupFilesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "files"), params.ToParams(), options...)
}

func NewGetGroupFilesBatchRequest(id string, params GetGroupFilesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.GroupFile]] {
	return core.NewBatchRequest[core.Cursor[objects.GroupFile]](GetGroupFilesBatchCall(id, params, options...))
}

func DecodeGetGroupFilesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.GroupFile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.GroupFile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupFilesWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupFilesParams) (*core.Cursor[objects.GroupFile], *core.Response, error) {
	var out core.Cursor[objects.GroupFile]
	call := GetGroupFilesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupFiles(ctx context.Context, client *core.Client, id string, params GetGroupFilesParams) (*core.Cursor[objects.GroupFile], error) {
	out, _, err := GetGroupFilesWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupGroupsBatchCall(id string, params GetGroupGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "groups"), params.ToParams(), options...)
}

func NewGetGroupGroupsBatchRequest(id string, params GetGroupGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Group]] {
	return core.NewBatchRequest[core.Cursor[objects.Group]](GetGroupGroupsBatchCall(id, params, options...))
}

func DecodeGetGroupGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Group], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Group]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupGroupsParams) (*core.Cursor[objects.Group], *core.Response, error) {
	var out core.Cursor[objects.Group]
	call := GetGroupGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupGroups(ctx context.Context, client *core.Client, id string, params GetGroupGroupsParams) (*core.Cursor[objects.Group], error) {
	out, _, err := GetGroupGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupGroupsParams struct {
	Admin                     *int                                       `facebook:"admin"`
	Description               *string                                    `facebook:"description"`
	GroupIconID               *core.ID                                   `facebook:"group_icon_id"`
	GroupType                 *enums.GroupgroupsGroupTypeEnumParam       `facebook:"group_type"`
	JoinSetting               *enums.GroupgroupsJoinSettingEnumParam     `facebook:"join_setting"`
	Name                      string                                     `facebook:"name"`
	ParentID                  *core.ID                                   `facebook:"parent_id"`
	PostPermissions           *enums.GroupgroupsPostPermissionsEnumParam `facebook:"post_permissions"`
	PostRequiresAdminApproval *bool                                      `facebook:"post_requires_admin_approval"`
	Privacy                   *string                                    `facebook:"privacy"`
	Ref                       *string                                    `facebook:"ref"`
	Extra                     core.Params                                `facebook:"-"`
}

func (params CreateGroupGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Admin != nil {
		out["admin"] = *params.Admin
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GroupIconID != nil {
		out["group_icon_id"] = *params.GroupIconID
	}
	if params.GroupType != nil {
		out["group_type"] = *params.GroupType
	}
	if params.JoinSetting != nil {
		out["join_setting"] = *params.JoinSetting
	}
	out["name"] = params.Name
	if params.ParentID != nil {
		out["parent_id"] = *params.ParentID
	}
	if params.PostPermissions != nil {
		out["post_permissions"] = *params.PostPermissions
	}
	if params.PostRequiresAdminApproval != nil {
		out["post_requires_admin_approval"] = *params.PostRequiresAdminApproval
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	return out
}

func CreateGroupGroupsBatchCall(id string, params CreateGroupGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "groups"), params.ToParams(), options...)
}

func NewCreateGroupGroupsBatchRequest(id string, params CreateGroupGroupsParams, options ...core.BatchOption) *core.BatchRequest[objects.Group] {
	return core.NewBatchRequest[objects.Group](CreateGroupGroupsBatchCall(id, params, options...))
}

func DecodeCreateGroupGroupsBatchResponse(response *core.BatchResponse) (*objects.Group, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Group
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateGroupGroupsWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupGroupsParams) (*objects.Group, *core.Response, error) {
	var out objects.Group
	call := CreateGroupGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupGroups(ctx context.Context, client *core.Client, id string, params CreateGroupGroupsParams) (*objects.Group, error) {
	out, _, err := CreateGroupGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupLiveVideosParams struct {
	BroadcastStatus *[]enums.GroupliveVideosBroadcastStatusEnumParam `facebook:"broadcast_status"`
	Source          *enums.GroupliveVideosSourceEnumParam            `facebook:"source"`
	Extra           core.Params                                      `facebook:"-"`
}

func (params GetGroupLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BroadcastStatus != nil {
		out["broadcast_status"] = *params.BroadcastStatus
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetGroupLiveVideosBatchCall(id string, params GetGroupLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewGetGroupLiveVideosBatchRequest(id string, params GetGroupLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LiveVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.LiveVideo]](GetGroupLiveVideosBatchCall(id, params, options...))
}

func DecodeGetGroupLiveVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LiveVideo], error) {
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

func GetGroupLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupLiveVideosParams) (*core.Cursor[objects.LiveVideo], *core.Response, error) {
	var out core.Cursor[objects.LiveVideo]
	call := GetGroupLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupLiveVideos(ctx context.Context, client *core.Client, id string, params GetGroupLiveVideosParams) (*core.Cursor[objects.LiveVideo], error) {
	out, _, err := GetGroupLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupLiveVideosParams struct {
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
	Projection                 *enums.GroupliveVideosProjectionEnumParam         `facebook:"projection"`
	Published                  *bool                                             `facebook:"published"`
	ScheduleCustomProfileImage *core.FileParam                                   `facebook:"schedule_custom_profile_image"`
	SpatialAudioFormat         *enums.GroupliveVideosSpatialAudioFormatEnumParam `facebook:"spatial_audio_format"`
	Status                     *enums.GroupliveVideosStatusEnumParam             `facebook:"status"`
	StereoscopicMode           *enums.GroupliveVideosStereoscopicModeEnumParam   `facebook:"stereoscopic_mode"`
	StopOnDeleteStream         *bool                                             `facebook:"stop_on_delete_stream"`
	StreamType                 *enums.GroupliveVideosStreamTypeEnumParam         `facebook:"stream_type"`
	Title                      *string                                           `facebook:"title"`
	Extra                      core.Params                                       `facebook:"-"`
}

func (params CreateGroupLiveVideosParams) ToParams() core.Params {
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

func CreateGroupLiveVideosBatchCall(id string, params CreateGroupLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewCreateGroupLiveVideosBatchRequest(id string, params CreateGroupLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](CreateGroupLiveVideosBatchCall(id, params, options...))
}

func DecodeCreateGroupLiveVideosBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
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

func CreateGroupLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupLiveVideosParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := CreateGroupLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupLiveVideos(ctx context.Context, client *core.Client, id string, params CreateGroupLiveVideosParams) (*objects.LiveVideo, error) {
	out, _, err := CreateGroupLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteGroupMembersParams struct {
	Email  *string     `facebook:"email"`
	Member *int        `facebook:"member"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteGroupMembersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	if params.Member != nil {
		out["member"] = *params.Member
	}
	return out
}

func DeleteGroupMembersBatchCall(id string, params DeleteGroupMembersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "members"), params.ToParams(), options...)
}

func NewDeleteGroupMembersBatchRequest(id string, params DeleteGroupMembersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteGroupMembersBatchCall(id, params, options...))
}

func DecodeDeleteGroupMembersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteGroupMembersWithResponse(ctx context.Context, client *core.Client, id string, params DeleteGroupMembersParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteGroupMembersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteGroupMembers(ctx context.Context, client *core.Client, id string, params DeleteGroupMembersParams) (*map[string]interface{}, error) {
	out, _, err := DeleteGroupMembersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupMembersParams struct {
	Email  *string     `facebook:"email"`
	From   *int        `facebook:"from"`
	Member *int        `facebook:"member"`
	Rate   *uint64     `facebook:"rate"`
	Source *string     `facebook:"source"`
	Extra  core.Params `facebook:"-"`
}

func (params CreateGroupMembersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	if params.From != nil {
		out["from"] = *params.From
	}
	if params.Member != nil {
		out["member"] = *params.Member
	}
	if params.Rate != nil {
		out["rate"] = *params.Rate
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func CreateGroupMembersBatchCall(id string, params CreateGroupMembersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "members"), params.ToParams(), options...)
}

func NewCreateGroupMembersBatchRequest(id string, params CreateGroupMembersParams, options ...core.BatchOption) *core.BatchRequest[objects.Group] {
	return core.NewBatchRequest[objects.Group](CreateGroupMembersBatchCall(id, params, options...))
}

func DecodeCreateGroupMembersBatchResponse(response *core.BatchResponse) (*objects.Group, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Group
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateGroupMembersWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupMembersParams) (*objects.Group, *core.Response, error) {
	var out objects.Group
	call := CreateGroupMembersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupMembers(ctx context.Context, client *core.Client, id string, params CreateGroupMembersParams) (*objects.Group, error) {
	out, _, err := CreateGroupMembersWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupOptedInMembersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGroupOptedInMembersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGroupOptedInMembersBatchCall(id string, params GetGroupOptedInMembersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "opted_in_members"), params.ToParams(), options...)
}

func NewGetGroupOptedInMembersBatchRequest(id string, params GetGroupOptedInMembersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.User]] {
	return core.NewBatchRequest[core.Cursor[objects.User]](GetGroupOptedInMembersBatchCall(id, params, options...))
}

func DecodeGetGroupOptedInMembersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.User], error) {
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

func GetGroupOptedInMembersWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupOptedInMembersParams) (*core.Cursor[objects.User], *core.Response, error) {
	var out core.Cursor[objects.User]
	call := GetGroupOptedInMembersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupOptedInMembers(ctx context.Context, client *core.Client, id string, params GetGroupOptedInMembersParams) (*core.Cursor[objects.User], error) {
	out, _, err := GetGroupOptedInMembersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupPhotosParams struct {
	Aid                                   *string                                             `facebook:"aid"`
	AllowSphericalPhoto                   *bool                                               `facebook:"allow_spherical_photo"`
	AltTextCustom                         *string                                             `facebook:"alt_text_custom"`
	AndroidKeyHash                        *string                                             `facebook:"android_key_hash"`
	ApplicationID                         *core.ID                                            `facebook:"application_id"`
	Attempt                               *uint64                                             `facebook:"attempt"`
	AudienceExp                           *bool                                               `facebook:"audience_exp"`
	BackdatedTime                         *core.Time                                          `facebook:"backdated_time"`
	BackdatedTimeGranularity              *enums.GroupphotosBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
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
	UnpublishedContentType                *enums.GroupphotosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	URL                                   *string                                             `facebook:"url"`
	UserSelectedTags                      *bool                                               `facebook:"user_selected_tags"`
	VaultImageID                          *core.ID                                            `facebook:"vault_image_id"`
	Extra                                 core.Params                                         `facebook:"-"`
}

func (params CreateGroupPhotosParams) ToParams() core.Params {
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

func CreateGroupPhotosBatchCall(id string, params CreateGroupPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewCreateGroupPhotosBatchRequest(id string, params CreateGroupPhotosParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](CreateGroupPhotosBatchCall(id, params, options...))
}

func DecodeCreateGroupPhotosBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
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

func CreateGroupPhotosWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupPhotosParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := CreateGroupPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupPhotos(ctx context.Context, client *core.Client, id string, params CreateGroupPhotosParams) (*objects.Photo, error) {
	out, _, err := CreateGroupPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupPictureParams struct {
	Height   *int                             `facebook:"height"`
	Redirect *bool                            `facebook:"redirect"`
	Type     *enums.GrouppictureTypeEnumParam `facebook:"type"`
	Width    *int                             `facebook:"width"`
	Extra    core.Params                      `facebook:"-"`
}

func (params GetGroupPictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Redirect != nil {
		out["redirect"] = *params.Redirect
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetGroupPictureBatchCall(id string, params GetGroupPictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetGroupPictureBatchRequest(id string, params GetGroupPictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetGroupPictureBatchCall(id, params, options...))
}

func DecodeGetGroupPictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
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

func GetGroupPictureWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupPictureParams) (*core.Cursor[objects.ProfilePictureSource], *core.Response, error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetGroupPictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupPicture(ctx context.Context, client *core.Client, id string, params GetGroupPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	out, _, err := GetGroupPictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupVideosParams struct {
	Type  *enums.GroupvideosTypeEnumParam `facebook:"type"`
	Extra core.Params                     `facebook:"-"`
}

func (params GetGroupVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetGroupVideosBatchCall(id string, params GetGroupVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewGetGroupVideosBatchRequest(id string, params GetGroupVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetGroupVideosBatchCall(id, params, options...))
}

func DecodeGetGroupVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupVideosParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetGroupVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroupVideos(ctx context.Context, client *core.Client, id string, params GetGroupVideosParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetGroupVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateGroupVideosParams struct {
	ApplicationID                 *core.ID                                          `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                          `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                           `facebook:"audio_story_wave_animation_handle"`
	ComposerEntryPicker           *string                                           `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                           `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                           `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                           `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                          `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                           `facebook:"composer_source_surface"`
	ComposerType                  *string                                           `facebook:"composer_type"`
	ContainerType                 *enums.GroupvideosContainerTypeEnumParam          `facebook:"container_type"`
	ContentCategory               *enums.GroupvideosContentCategoryEnumParam        `facebook:"content_category"`
	CreativeTools                 *string                                           `facebook:"creative_tools"`
	Description                   *string                                           `facebook:"description"`
	EditDescriptionSpec           *map[string]interface{}                           `facebook:"edit_description_spec"`
	Embeddable                    *bool                                             `facebook:"embeddable"`
	EndOffset                     *uint64                                           `facebook:"end_offset"`
	FbuploaderVideoFileChunk      *string                                           `facebook:"fbuploader_video_file_chunk"`
	FileSize                      *uint64                                           `facebook:"file_size"`
	FileURL                       *string                                           `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                             `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.GroupvideosFormattingEnumParam             `facebook:"formatting"`
	Fov                           *uint64                                           `facebook:"fov"`
	FrontZRotation                *float64                                          `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                          `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                          `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                       `facebook:"guide"`
	GuideEnabled                  *bool                                             `facebook:"guide_enabled"`
	InitialHeading                *uint64                                           `facebook:"initial_heading"`
	InitialPitch                  *uint64                                           `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                           `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                             `facebook:"is_boost_intended"`
	IsExplicitShare               *bool                                             `facebook:"is_explicit_share"`
	IsGroupLinkingPost            *bool                                             `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                             `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                             `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                          `facebook:"location_source_id"`
	ManualPrivacy                 *bool                                             `facebook:"manual_privacy"`
	OgActionTypeID                *core.ID                                          `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                          `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                          `facebook:"og_object_id"`
	OgPhrase                      *string                                           `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                           `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                           `facebook:"original_fov"`
	OriginalProjectionType        *enums.GroupvideosOriginalProjectionTypeEnumParam `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                           `facebook:"partnership_ad_ad_code"`
	PublishEventID                *core.ID                                          `facebook:"publish_event_id"`
	Published                     *bool                                             `facebook:"published"`
	ReferencedStickerID           *core.ID                                          `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                          `facebook:"replace_video_id"`
	ScheduledPublishTime          *uint64                                           `facebook:"scheduled_publish_time"`
	SelectedAudioSpec             *map[string]interface{}                           `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                           `facebook:"slideshow_spec"`
	Source                        *string                                           `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                          `facebook:"source_instagram_media_id"`
	Spherical                     *bool                                             `facebook:"spherical"`
	StartOffset                   *uint64                                           `facebook:"start_offset"`
	SwapMode                      *enums.GroupvideosSwapModeEnumParam               `facebook:"swap_mode"`
	TextFormatMetadata            *string                                           `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                   `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                           `facebook:"time_since_original_post"`
	Title                         *string                                           `facebook:"title"`
	TranscodeSettingProperties    *string                                           `facebook:"transcode_setting_properties"`
	UnpublishedContentType        *enums.GroupvideosUnpublishedContentTypeEnumParam `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.GroupvideosUploadPhaseEnumParam            `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                          `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                           `facebook:"upload_setting_properties"`
	VideoFileChunk                *string                                           `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                           `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                           `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                          `facebook:"waterfall_id"`
	Extra                         core.Params                                       `facebook:"-"`
}

func (params CreateGroupVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.AudioStoryWaveAnimationHandle != nil {
		out["audio_story_wave_animation_handle"] = *params.AudioStoryWaveAnimationHandle
	}
	if params.ComposerEntryPicker != nil {
		out["composer_entry_picker"] = *params.ComposerEntryPicker
	}
	if params.ComposerEntryPoint != nil {
		out["composer_entry_point"] = *params.ComposerEntryPoint
	}
	if params.ComposerEntryTime != nil {
		out["composer_entry_time"] = *params.ComposerEntryTime
	}
	if params.ComposerSessionEventsLog != nil {
		out["composer_session_events_log"] = *params.ComposerSessionEventsLog
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.ComposerSourceSurface != nil {
		out["composer_source_surface"] = *params.ComposerSourceSurface
	}
	if params.ComposerType != nil {
		out["composer_type"] = *params.ComposerType
	}
	if params.ContainerType != nil {
		out["container_type"] = *params.ContainerType
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	if params.CreativeTools != nil {
		out["creative_tools"] = *params.CreativeTools
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EditDescriptionSpec != nil {
		out["edit_description_spec"] = *params.EditDescriptionSpec
	}
	if params.Embeddable != nil {
		out["embeddable"] = *params.Embeddable
	}
	if params.EndOffset != nil {
		out["end_offset"] = *params.EndOffset
	}
	if params.FbuploaderVideoFileChunk != nil {
		out["fbuploader_video_file_chunk"] = *params.FbuploaderVideoFileChunk
	}
	if params.FileSize != nil {
		out["file_size"] = *params.FileSize
	}
	if params.FileURL != nil {
		out["file_url"] = *params.FileURL
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.Fov != nil {
		out["fov"] = *params.Fov
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Guide != nil {
		out["guide"] = *params.Guide
	}
	if params.GuideEnabled != nil {
		out["guide_enabled"] = *params.GuideEnabled
	}
	if params.InitialHeading != nil {
		out["initial_heading"] = *params.InitialHeading
	}
	if params.InitialPitch != nil {
		out["initial_pitch"] = *params.InitialPitch
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPartnershipAd != nil {
		out["is_partnership_ad"] = *params.IsPartnershipAd
	}
	if params.IsVoiceClip != nil {
		out["is_voice_clip"] = *params.IsVoiceClip
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
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
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.OriginalProjectionType != nil {
		out["original_projection_type"] = *params.OriginalProjectionType
	}
	if params.PartnershipAdAdCode != nil {
		out["partnership_ad_ad_code"] = *params.PartnershipAdAdCode
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ReferencedStickerID != nil {
		out["referenced_sticker_id"] = *params.ReferencedStickerID
	}
	if params.ReplaceVideoID != nil {
		out["replace_video_id"] = *params.ReplaceVideoID
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.SelectedAudioSpec != nil {
		out["selected_audio_spec"] = *params.SelectedAudioSpec
	}
	if params.SlideshowSpec != nil {
		out["slideshow_spec"] = *params.SlideshowSpec
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SourceInstagramMediaID != nil {
		out["source_instagram_media_id"] = *params.SourceInstagramMediaID
	}
	if params.Spherical != nil {
		out["spherical"] = *params.Spherical
	}
	if params.StartOffset != nil {
		out["start_offset"] = *params.StartOffset
	}
	if params.SwapMode != nil {
		out["swap_mode"] = *params.SwapMode
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.Thumb != nil {
		out["thumb"] = *params.Thumb
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TranscodeSettingProperties != nil {
		out["transcode_setting_properties"] = *params.TranscodeSettingProperties
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UploadPhase != nil {
		out["upload_phase"] = *params.UploadPhase
	}
	if params.UploadSessionID != nil {
		out["upload_session_id"] = *params.UploadSessionID
	}
	if params.UploadSettingProperties != nil {
		out["upload_setting_properties"] = *params.UploadSettingProperties
	}
	if params.VideoFileChunk != nil {
		out["video_file_chunk"] = *params.VideoFileChunk
	}
	if params.VideoIDOriginal != nil {
		out["video_id_original"] = *params.VideoIDOriginal
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.WaterfallID != nil {
		out["waterfall_id"] = *params.WaterfallID
	}
	return out
}

func CreateGroupVideosBatchCall(id string, params CreateGroupVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewCreateGroupVideosBatchRequest(id string, params CreateGroupVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateGroupVideosBatchCall(id, params, options...))
}

func DecodeCreateGroupVideosBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
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

func CreateGroupVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreateGroupVideosParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateGroupVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateGroupVideos(ctx context.Context, client *core.Client, id string, params CreateGroupVideosParams) (*objects.AdVideo, error) {
	out, _, err := CreateGroupVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetGroupParams struct {
	IconSize *enums.GroupIconSize `facebook:"icon_size"`
	Extra    core.Params          `facebook:"-"`
}

func (params GetGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IconSize != nil {
		out["icon_size"] = *params.IconSize
	}
	return out
}

func GetGroupBatchCall(id string, params GetGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetGroupBatchRequest(id string, params GetGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.Group] {
	return core.NewBatchRequest[objects.Group](GetGroupBatchCall(id, params, options...))
}

func DecodeGetGroupBatchResponse(response *core.BatchResponse) (*objects.Group, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Group
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGroupWithResponse(ctx context.Context, client *core.Client, id string, params GetGroupParams) (*objects.Group, *core.Response, error) {
	var out objects.Group
	call := GetGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGroup(ctx context.Context, client *core.Client, id string, params GetGroupParams) (*objects.Group, error) {
	out, _, err := GetGroupWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateGroupParams struct {
	Cover                     *string                     `facebook:"cover"`
	CoverURL                  *string                     `facebook:"cover_url"`
	Description               *string                     `facebook:"description"`
	FocusX                    *float64                    `facebook:"focus_x"`
	FocusY                    *float64                    `facebook:"focus_y"`
	GroupIcon                 *string                     `facebook:"group_icon"`
	IsOfficialGroup           *bool                       `facebook:"is_official_group"`
	JoinSetting               *enums.GroupJoinSetting     `facebook:"join_setting"`
	Name                      *string                     `facebook:"name"`
	NoFeedStory               *bool                       `facebook:"no_feed_story"`
	OffsetY                   *int                        `facebook:"offset_y"`
	PostPermissions           *enums.GroupPostPermissions `facebook:"post_permissions"`
	PostRequiresAdminApproval *bool                       `facebook:"post_requires_admin_approval"`
	Privacy                   *string                     `facebook:"privacy"`
	Purpose                   *enums.GroupPurpose         `facebook:"purpose"`
	UpdateViewTime            *bool                       `facebook:"update_view_time"`
	Extra                     core.Params                 `facebook:"-"`
}

func (params UpdateGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cover != nil {
		out["cover"] = *params.Cover
	}
	if params.CoverURL != nil {
		out["cover_url"] = *params.CoverURL
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.FocusX != nil {
		out["focus_x"] = *params.FocusX
	}
	if params.FocusY != nil {
		out["focus_y"] = *params.FocusY
	}
	if params.GroupIcon != nil {
		out["group_icon"] = *params.GroupIcon
	}
	if params.IsOfficialGroup != nil {
		out["is_official_group"] = *params.IsOfficialGroup
	}
	if params.JoinSetting != nil {
		out["join_setting"] = *params.JoinSetting
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NoFeedStory != nil {
		out["no_feed_story"] = *params.NoFeedStory
	}
	if params.OffsetY != nil {
		out["offset_y"] = *params.OffsetY
	}
	if params.PostPermissions != nil {
		out["post_permissions"] = *params.PostPermissions
	}
	if params.PostRequiresAdminApproval != nil {
		out["post_requires_admin_approval"] = *params.PostRequiresAdminApproval
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Purpose != nil {
		out["purpose"] = *params.Purpose
	}
	if params.UpdateViewTime != nil {
		out["update_view_time"] = *params.UpdateViewTime
	}
	return out
}

func UpdateGroupBatchCall(id string, params UpdateGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateGroupBatchRequest(id string, params UpdateGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.Group] {
	return core.NewBatchRequest[objects.Group](UpdateGroupBatchCall(id, params, options...))
}

func DecodeUpdateGroupBatchResponse(response *core.BatchResponse) (*objects.Group, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Group
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateGroupWithResponse(ctx context.Context, client *core.Client, id string, params UpdateGroupParams) (*objects.Group, *core.Response, error) {
	var out objects.Group
	call := UpdateGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateGroup(ctx context.Context, client *core.Client, id string, params UpdateGroupParams) (*objects.Group, error) {
	out, _, err := UpdateGroupWithResponse(ctx, client, id, params)
	return out, err
}
