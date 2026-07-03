package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetAdAccountAccountControlsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAccountControlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAccountControlsBatchCall(id string, params GetAdAccountAccountControlsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "account_controls"), params.ToParams(), options...)
}

func NewGetAdAccountAccountControlsBatchRequest(id string, params GetAdAccountAccountControlsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountBusinessConstraints]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountBusinessConstraints]](GetAdAccountAccountControlsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAccountControlsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountBusinessConstraints], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountBusinessConstraints]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAccountControls(ctx context.Context, client *core.Client, id string, params GetAdAccountAccountControlsParams) (*core.Cursor[objects.AdAccountBusinessConstraints], error) {
	var out core.Cursor[objects.AdAccountBusinessConstraints]
	call := GetAdAccountAccountControlsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAccountControlsParams struct {
	AudienceControls  map[string]interface{}  `facebook:"audience_controls"`
	PlacementControls *map[string]interface{} `facebook:"placement_controls"`
	Extra             core.Params             `facebook:"-"`
}

func (params CreateAdAccountAccountControlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["audience_controls"] = params.AudienceControls
	if params.PlacementControls != nil {
		out["placement_controls"] = *params.PlacementControls
	}
	return out
}

func CreateAdAccountAccountControlsBatchCall(id string, params CreateAdAccountAccountControlsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "account_controls"), params.ToParams(), options...)
}

func NewCreateAdAccountAccountControlsBatchRequest(id string, params CreateAdAccountAccountControlsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountBusinessConstraints] {
	return core.NewBatchRequest[objects.AdAccountBusinessConstraints](CreateAdAccountAccountControlsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAccountControlsBatchResponse(response *core.BatchResponse) (*objects.AdAccountBusinessConstraints, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountBusinessConstraints
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAccountControls(ctx context.Context, client *core.Client, id string, params CreateAdAccountAccountControlsParams) (*objects.AdAccountBusinessConstraints, error) {
	var out objects.AdAccountBusinessConstraints
	call := CreateAdAccountAccountControlsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountActivitiesParams struct {
	AddChildren *bool                                         `facebook:"add_children"`
	After       *string                                       `facebook:"after"`
	BusinessID  *core.ID                                      `facebook:"business_id"`
	Category    *enums.AdaccountactivitiesCategoryEnumParam   `facebook:"category"`
	DataSource  *enums.AdaccountactivitiesDataSourceEnumParam `facebook:"data_source"`
	ExtraOids   *[]string                                     `facebook:"extra_oids"`
	Limit       *int                                          `facebook:"limit"`
	Oid         *string                                       `facebook:"oid"`
	Since       *time.Time                                    `facebook:"since"`
	UID         *core.ID                                      `facebook:"uid"`
	Until       *time.Time                                    `facebook:"until"`
	Extra       core.Params                                   `facebook:"-"`
}

func (params GetAdAccountActivitiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AddChildren != nil {
		out["add_children"] = *params.AddChildren
	}
	if params.After != nil {
		out["after"] = *params.After
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.DataSource != nil {
		out["data_source"] = *params.DataSource
	}
	if params.ExtraOids != nil {
		out["extra_oids"] = *params.ExtraOids
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.Oid != nil {
		out["oid"] = *params.Oid
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetAdAccountActivitiesBatchCall(id string, params GetAdAccountActivitiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "activities"), params.ToParams(), options...)
}

func NewGetAdAccountActivitiesBatchRequest(id string, params GetAdAccountActivitiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdActivity]] {
	return core.NewBatchRequest[core.Cursor[objects.AdActivity]](GetAdAccountActivitiesBatchCall(id, params, options...))
}

func DecodeGetAdAccountActivitiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdActivity], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdActivity]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountActivities(ctx context.Context, client *core.Client, id string, params GetAdAccountActivitiesParams) (*core.Cursor[objects.AdActivity], error) {
	var out core.Cursor[objects.AdActivity]
	call := GetAdAccountActivitiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdPlacePageSetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdPlacePageSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdPlacePageSetsBatchCall(id string, params GetAdAccountAdPlacePageSetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_place_page_sets"), params.ToParams(), options...)
}

func NewGetAdAccountAdPlacePageSetsBatchRequest(id string, params GetAdAccountAdPlacePageSetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPlacePageSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPlacePageSet]](GetAdAccountAdPlacePageSetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdPlacePageSetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPlacePageSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPlacePageSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdPlacePageSets(ctx context.Context, client *core.Client, id string, params GetAdAccountAdPlacePageSetsParams) (*core.Cursor[objects.AdPlacePageSet], error) {
	var out core.Cursor[objects.AdPlacePageSet]
	call := GetAdAccountAdPlacePageSetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdPlacePageSetsParams struct {
	LocationTypes    *[]enums.AdaccountadPlacePageSetsLocationTypesEnumParam  `facebook:"location_types"`
	Name             string                                                   `facebook:"name"`
	ParentPage       string                                                   `facebook:"parent_page"`
	TargetedAreaType *enums.AdaccountadPlacePageSetsTargetedAreaTypeEnumParam `facebook:"targeted_area_type"`
	Extra            core.Params                                              `facebook:"-"`
}

func (params CreateAdAccountAdPlacePageSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.LocationTypes != nil {
		out["location_types"] = *params.LocationTypes
	}
	out["name"] = params.Name
	out["parent_page"] = params.ParentPage
	if params.TargetedAreaType != nil {
		out["targeted_area_type"] = *params.TargetedAreaType
	}
	return out
}

func CreateAdAccountAdPlacePageSetsBatchCall(id string, params CreateAdAccountAdPlacePageSetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_place_page_sets"), params.ToParams(), options...)
}

func NewCreateAdAccountAdPlacePageSetsBatchRequest(id string, params CreateAdAccountAdPlacePageSetsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdPlacePageSet] {
	return core.NewBatchRequest[objects.AdPlacePageSet](CreateAdAccountAdPlacePageSetsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdPlacePageSetsBatchResponse(response *core.BatchResponse) (*objects.AdPlacePageSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdPlacePageSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdPlacePageSets(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdPlacePageSetsParams) (*objects.AdPlacePageSet, error) {
	var out objects.AdPlacePageSet
	call := CreateAdAccountAdPlacePageSetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdPlacePageSetsAsyncParams struct {
	LocationTypes    *[]enums.AdaccountadPlacePageSetsAsyncLocationTypesEnumParam  `facebook:"location_types"`
	Name             string                                                        `facebook:"name"`
	ParentPage       string                                                        `facebook:"parent_page"`
	TargetedAreaType *enums.AdaccountadPlacePageSetsAsyncTargetedAreaTypeEnumParam `facebook:"targeted_area_type"`
	Extra            core.Params                                                   `facebook:"-"`
}

func (params CreateAdAccountAdPlacePageSetsAsyncParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.LocationTypes != nil {
		out["location_types"] = *params.LocationTypes
	}
	out["name"] = params.Name
	out["parent_page"] = params.ParentPage
	if params.TargetedAreaType != nil {
		out["targeted_area_type"] = *params.TargetedAreaType
	}
	return out
}

func CreateAdAccountAdPlacePageSetsAsyncBatchCall(id string, params CreateAdAccountAdPlacePageSetsAsyncParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_place_page_sets_async"), params.ToParams(), options...)
}

func NewCreateAdAccountAdPlacePageSetsAsyncBatchRequest(id string, params CreateAdAccountAdPlacePageSetsAsyncParams, options ...core.BatchOption) *core.BatchRequest[objects.AdPlacePageSet] {
	return core.NewBatchRequest[objects.AdPlacePageSet](CreateAdAccountAdPlacePageSetsAsyncBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdPlacePageSetsAsyncBatchResponse(response *core.BatchResponse) (*objects.AdPlacePageSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdPlacePageSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdPlacePageSetsAsync(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdPlacePageSetsAsyncParams) (*objects.AdPlacePageSet, error) {
	var out objects.AdPlacePageSet
	call := CreateAdAccountAdPlacePageSetsAsyncBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdSavedKeywordsParams struct {
	Fields *[]string   `facebook:"fields"`
	Extra  core.Params `facebook:"-"`
}

func (params GetAdAccountAdSavedKeywordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	return out
}

func GetAdAccountAdSavedKeywordsBatchCall(id string, params GetAdAccountAdSavedKeywordsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_saved_keywords"), params.ToParams(), options...)
}

func NewGetAdAccountAdSavedKeywordsBatchRequest(id string, params GetAdAccountAdSavedKeywordsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSavedKeywords]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSavedKeywords]](GetAdAccountAdSavedKeywordsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdSavedKeywordsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSavedKeywords], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSavedKeywords]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdSavedKeywords(ctx context.Context, client *core.Client, id string, params GetAdAccountAdSavedKeywordsParams) (*core.Cursor[objects.AdSavedKeywords], error) {
	var out core.Cursor[objects.AdSavedKeywords]
	call := GetAdAccountAdSavedKeywordsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdStudiesBatchCall(id string, params GetAdAccountAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewGetAdAccountAdStudiesBatchRequest(id string, params GetAdAccountAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudy]](GetAdAccountAdStudiesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdStudies(ctx context.Context, client *core.Client, id string, params GetAdAccountAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	call := GetAdAccountAdStudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdcloudplayablesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdcloudplayablesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdcloudplayablesBatchCall(id string, params GetAdAccountAdcloudplayablesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adcloudplayables"), params.ToParams(), options...)
}

func NewGetAdAccountAdcloudplayablesBatchRequest(id string, params GetAdAccountAdcloudplayablesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CloudGame]] {
	return core.NewBatchRequest[core.Cursor[objects.CloudGame]](GetAdAccountAdcloudplayablesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdcloudplayablesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CloudGame], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CloudGame]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdcloudplayables(ctx context.Context, client *core.Client, id string, params GetAdAccountAdcloudplayablesParams) (*core.Cursor[objects.CloudGame], error) {
	var out core.Cursor[objects.CloudGame]
	call := GetAdAccountAdcloudplayablesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdcreativesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdcreativesBatchCall(id string, params GetAdAccountAdcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adcreatives"), params.ToParams(), options...)
}

func NewGetAdAccountAdcreativesBatchRequest(id string, params GetAdAccountAdcreativesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCreative]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCreative]](GetAdAccountAdcreativesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdcreativesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCreative], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCreative]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdcreatives(ctx context.Context, client *core.Client, id string, params GetAdAccountAdcreativesParams) (*core.Cursor[objects.AdCreative], error) {
	var out core.Cursor[objects.AdCreative]
	call := GetAdAccountAdcreativesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdcreativesParams struct {
	ActorID                          *core.ID                                                   `facebook:"actor_id"`
	AdDisclaimerSpec                 *map[string]interface{}                                    `facebook:"ad_disclaimer_spec"`
	Adlabels                         *[]map[string]interface{}                                  `facebook:"adlabels"`
	ApplinkTreatment                 *enums.AdaccountadcreativesApplinkTreatmentEnumParam       `facebook:"applink_treatment"`
	AssetFeedSpec                    *map[string]interface{}                                    `facebook:"asset_feed_spec"`
	AuthorizationCategory            *enums.AdaccountadcreativesAuthorizationCategoryEnumParam  `facebook:"authorization_category"`
	Body                             *string                                                    `facebook:"body"`
	BrandedContent                   *map[string]interface{}                                    `facebook:"branded_content"`
	BrandedContentSponsorPageID      *core.ID                                                   `facebook:"branded_content_sponsor_page_id"`
	BundleFolderID                   *core.ID                                                   `facebook:"bundle_folder_id"`
	CallToAction                     *map[string]interface{}                                    `facebook:"call_to_action"`
	CategorizationCriteria           *enums.AdaccountadcreativesCategorizationCriteriaEnumParam `facebook:"categorization_criteria"`
	CategoryMediaSource              *enums.AdaccountadcreativesCategoryMediaSourceEnumParam    `facebook:"category_media_source"`
	ContextualMultiAds               *map[string]interface{}                                    `facebook:"contextual_multi_ads"`
	CreativeSourcingSpec             *map[string]interface{}                                    `facebook:"creative_sourcing_spec"`
	DegreesOfFreedomSpec             *map[string]interface{}                                    `facebook:"degrees_of_freedom_spec"`
	DestinationSetID                 *core.ID                                                   `facebook:"destination_set_id"`
	DestinationSpec                  *map[string]interface{}                                    `facebook:"destination_spec"`
	DynamicAdVoice                   *enums.AdaccountadcreativesDynamicAdVoiceEnumParam         `facebook:"dynamic_ad_voice"`
	EnableLaunchInstantApp           *bool                                                      `facebook:"enable_launch_instant_app"`
	ExecutionOptions                 *[]enums.AdaccountadcreativesExecutionOptionsEnumParam     `facebook:"execution_options"`
	ExistingPostTitle                *string                                                    `facebook:"existing_post_title"`
	FacebookBrandedContent           *map[string]interface{}                                    `facebook:"facebook_branded_content"`
	FormatTransformationSpec         *[]map[string]interface{}                                  `facebook:"format_transformation_spec"`
	GenerativeAssetSpec              *map[string]interface{}                                    `facebook:"generative_asset_spec"`
	ImageCrops                       *map[string]interface{}                                    `facebook:"image_crops"`
	ImageFile                        *string                                                    `facebook:"image_file"`
	ImageHash                        *string                                                    `facebook:"image_hash"`
	ImageURL                         *string                                                    `facebook:"image_url"`
	InstagramBrandedContent          *map[string]interface{}                                    `facebook:"instagram_branded_content"`
	InstagramPermalinkURL            *string                                                    `facebook:"instagram_permalink_url"`
	InstagramUserID                  *core.ID                                                   `facebook:"instagram_user_id"`
	InteractiveComponentsSpec        *map[string]interface{}                                    `facebook:"interactive_components_spec"`
	IsDcoInternal                    *bool                                                      `facebook:"is_dco_internal"`
	LinkOgID                         *core.ID                                                   `facebook:"link_og_id"`
	LinkURL                          *string                                                    `facebook:"link_url"`
	MarketingMessageStructuredSpec   *map[string]interface{}                                    `facebook:"marketing_message_structured_spec"`
	MediaSourcingSpec                *map[string]interface{}                                    `facebook:"media_sourcing_spec"`
	Name                             *string                                                    `facebook:"name"`
	ObjectID                         *core.ID                                                   `facebook:"object_id"`
	ObjectStoryID                    *core.ID                                                   `facebook:"object_story_id"`
	ObjectStorySpec                  *objects.AdCreativeObjectStorySpec                         `facebook:"object_story_spec"`
	ObjectType                       *string                                                    `facebook:"object_type"`
	ObjectURL                        *string                                                    `facebook:"object_url"`
	OmnichannelLinkSpec              *map[string]interface{}                                    `facebook:"omnichannel_link_spec"`
	PageWelcomeMessage               *string                                                    `facebook:"page_welcome_message"`
	PlacePageSetID                   *core.ID                                                   `facebook:"place_page_set_id"`
	PlatformCustomizations           *map[string]interface{}                                    `facebook:"platform_customizations"`
	PlayableAssetID                  *core.ID                                                   `facebook:"playable_asset_id"`
	PortraitCustomizations           *map[string]interface{}                                    `facebook:"portrait_customizations"`
	ProductSetID                     *core.ID                                                   `facebook:"product_set_id"`
	ProductSuggestionSettings        *map[string]interface{}                                    `facebook:"product_suggestion_settings"`
	RecommenderSettings              *map[string]interface{}                                    `facebook:"recommender_settings"`
	RegionalRegulationDisclaimerSpec *map[string]interface{}                                    `facebook:"regional_regulation_disclaimer_spec"`
	SourceFacebookPostID             *core.ID                                                   `facebook:"source_facebook_post_id"`
	SourceInstagramMediaID           *core.ID                                                   `facebook:"source_instagram_media_id"`
	TemplateURL                      *string                                                    `facebook:"template_url"`
	TemplateURLSpec                  *string                                                    `facebook:"template_url_spec"`
	ThumbnailURL                     *string                                                    `facebook:"thumbnail_url"`
	Title                            *string                                                    `facebook:"title"`
	URLTags                          *string                                                    `facebook:"url_tags"`
	UsePageActorOverride             *bool                                                      `facebook:"use_page_actor_override"`
	WamoWhatsappIdentitySpec         *map[string]interface{}                                    `facebook:"wamo_whatsapp_identity_spec"`
	Extra                            core.Params                                                `facebook:"-"`
}

func (params CreateAdAccountAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActorID != nil {
		out["actor_id"] = *params.ActorID
	}
	if params.AdDisclaimerSpec != nil {
		out["ad_disclaimer_spec"] = *params.AdDisclaimerSpec
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.ApplinkTreatment != nil {
		out["applink_treatment"] = *params.ApplinkTreatment
	}
	if params.AssetFeedSpec != nil {
		out["asset_feed_spec"] = *params.AssetFeedSpec
	}
	if params.AuthorizationCategory != nil {
		out["authorization_category"] = *params.AuthorizationCategory
	}
	if params.Body != nil {
		out["body"] = *params.Body
	}
	if params.BrandedContent != nil {
		out["branded_content"] = *params.BrandedContent
	}
	if params.BrandedContentSponsorPageID != nil {
		out["branded_content_sponsor_page_id"] = *params.BrandedContentSponsorPageID
	}
	if params.BundleFolderID != nil {
		out["bundle_folder_id"] = *params.BundleFolderID
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.CategorizationCriteria != nil {
		out["categorization_criteria"] = *params.CategorizationCriteria
	}
	if params.CategoryMediaSource != nil {
		out["category_media_source"] = *params.CategoryMediaSource
	}
	if params.ContextualMultiAds != nil {
		out["contextual_multi_ads"] = *params.ContextualMultiAds
	}
	if params.CreativeSourcingSpec != nil {
		out["creative_sourcing_spec"] = *params.CreativeSourcingSpec
	}
	if params.DegreesOfFreedomSpec != nil {
		out["degrees_of_freedom_spec"] = *params.DegreesOfFreedomSpec
	}
	if params.DestinationSetID != nil {
		out["destination_set_id"] = *params.DestinationSetID
	}
	if params.DestinationSpec != nil {
		out["destination_spec"] = *params.DestinationSpec
	}
	if params.DynamicAdVoice != nil {
		out["dynamic_ad_voice"] = *params.DynamicAdVoice
	}
	if params.EnableLaunchInstantApp != nil {
		out["enable_launch_instant_app"] = *params.EnableLaunchInstantApp
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.ExistingPostTitle != nil {
		out["existing_post_title"] = *params.ExistingPostTitle
	}
	if params.FacebookBrandedContent != nil {
		out["facebook_branded_content"] = *params.FacebookBrandedContent
	}
	if params.FormatTransformationSpec != nil {
		out["format_transformation_spec"] = *params.FormatTransformationSpec
	}
	if params.GenerativeAssetSpec != nil {
		out["generative_asset_spec"] = *params.GenerativeAssetSpec
	}
	if params.ImageCrops != nil {
		out["image_crops"] = *params.ImageCrops
	}
	if params.ImageFile != nil {
		out["image_file"] = *params.ImageFile
	}
	if params.ImageHash != nil {
		out["image_hash"] = *params.ImageHash
	}
	if params.ImageURL != nil {
		out["image_url"] = *params.ImageURL
	}
	if params.InstagramBrandedContent != nil {
		out["instagram_branded_content"] = *params.InstagramBrandedContent
	}
	if params.InstagramPermalinkURL != nil {
		out["instagram_permalink_url"] = *params.InstagramPermalinkURL
	}
	if params.InstagramUserID != nil {
		out["instagram_user_id"] = *params.InstagramUserID
	}
	if params.InteractiveComponentsSpec != nil {
		out["interactive_components_spec"] = *params.InteractiveComponentsSpec
	}
	if params.IsDcoInternal != nil {
		out["is_dco_internal"] = *params.IsDcoInternal
	}
	if params.LinkOgID != nil {
		out["link_og_id"] = *params.LinkOgID
	}
	if params.LinkURL != nil {
		out["link_url"] = *params.LinkURL
	}
	if params.MarketingMessageStructuredSpec != nil {
		out["marketing_message_structured_spec"] = *params.MarketingMessageStructuredSpec
	}
	if params.MediaSourcingSpec != nil {
		out["media_sourcing_spec"] = *params.MediaSourcingSpec
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	if params.ObjectStoryID != nil {
		out["object_story_id"] = *params.ObjectStoryID
	}
	if params.ObjectStorySpec != nil {
		out["object_story_spec"] = *params.ObjectStorySpec
	}
	if params.ObjectType != nil {
		out["object_type"] = *params.ObjectType
	}
	if params.ObjectURL != nil {
		out["object_url"] = *params.ObjectURL
	}
	if params.OmnichannelLinkSpec != nil {
		out["omnichannel_link_spec"] = *params.OmnichannelLinkSpec
	}
	if params.PageWelcomeMessage != nil {
		out["page_welcome_message"] = *params.PageWelcomeMessage
	}
	if params.PlacePageSetID != nil {
		out["place_page_set_id"] = *params.PlacePageSetID
	}
	if params.PlatformCustomizations != nil {
		out["platform_customizations"] = *params.PlatformCustomizations
	}
	if params.PlayableAssetID != nil {
		out["playable_asset_id"] = *params.PlayableAssetID
	}
	if params.PortraitCustomizations != nil {
		out["portrait_customizations"] = *params.PortraitCustomizations
	}
	if params.ProductSetID != nil {
		out["product_set_id"] = *params.ProductSetID
	}
	if params.ProductSuggestionSettings != nil {
		out["product_suggestion_settings"] = *params.ProductSuggestionSettings
	}
	if params.RecommenderSettings != nil {
		out["recommender_settings"] = *params.RecommenderSettings
	}
	if params.RegionalRegulationDisclaimerSpec != nil {
		out["regional_regulation_disclaimer_spec"] = *params.RegionalRegulationDisclaimerSpec
	}
	if params.SourceFacebookPostID != nil {
		out["source_facebook_post_id"] = *params.SourceFacebookPostID
	}
	if params.SourceInstagramMediaID != nil {
		out["source_instagram_media_id"] = *params.SourceInstagramMediaID
	}
	if params.TemplateURL != nil {
		out["template_url"] = *params.TemplateURL
	}
	if params.TemplateURLSpec != nil {
		out["template_url_spec"] = *params.TemplateURLSpec
	}
	if params.ThumbnailURL != nil {
		out["thumbnail_url"] = *params.ThumbnailURL
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.URLTags != nil {
		out["url_tags"] = *params.URLTags
	}
	if params.UsePageActorOverride != nil {
		out["use_page_actor_override"] = *params.UsePageActorOverride
	}
	if params.WamoWhatsappIdentitySpec != nil {
		out["wamo_whatsapp_identity_spec"] = *params.WamoWhatsappIdentitySpec
	}
	return out
}

func CreateAdAccountAdcreativesBatchCall(id string, params CreateAdAccountAdcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adcreatives"), params.ToParams(), options...)
}

func NewCreateAdAccountAdcreativesBatchRequest(id string, params CreateAdAccountAdcreativesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCreative] {
	return core.NewBatchRequest[objects.AdCreative](CreateAdAccountAdcreativesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdcreativesBatchResponse(response *core.BatchResponse) (*objects.AdCreative, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCreative
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdcreatives(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdcreativesParams) (*objects.AdCreative, error) {
	var out objects.AdCreative
	call := CreateAdAccountAdcreativesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdcreativesbylabelsParams struct {
	AdLabelIds []core.ID                                            `facebook:"ad_label_ids"`
	Operator   *enums.AdaccountadcreativesbylabelsOperatorEnumParam `facebook:"operator"`
	Extra      core.Params                                          `facebook:"-"`
}

func (params GetAdAccountAdcreativesbylabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_label_ids"] = params.AdLabelIds
	if params.Operator != nil {
		out["operator"] = *params.Operator
	}
	return out
}

func GetAdAccountAdcreativesbylabelsBatchCall(id string, params GetAdAccountAdcreativesbylabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adcreativesbylabels"), params.ToParams(), options...)
}

func NewGetAdAccountAdcreativesbylabelsBatchRequest(id string, params GetAdAccountAdcreativesbylabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCreative]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCreative]](GetAdAccountAdcreativesbylabelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdcreativesbylabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCreative], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCreative]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdcreativesbylabels(ctx context.Context, client *core.Client, id string, params GetAdAccountAdcreativesbylabelsParams) (*core.Cursor[objects.AdCreative], error) {
	var out core.Cursor[objects.AdCreative]
	call := GetAdAccountAdcreativesbylabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountAdimagesParams struct {
	Hash    string      `facebook:"hash"`
	ImageID *core.ID    `facebook:"image_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteAdAccountAdimagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["hash"] = params.Hash
	if params.ImageID != nil {
		out["image_id"] = *params.ImageID
	}
	return out
}

func DeleteAdAccountAdimagesBatchCall(id string, params DeleteAdAccountAdimagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "adimages"), params.ToParams(), options...)
}

func NewDeleteAdAccountAdimagesBatchRequest(id string, params DeleteAdAccountAdimagesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountAdimagesBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountAdimagesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountAdimages(ctx context.Context, client *core.Client, id string, params DeleteAdAccountAdimagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountAdimagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdimagesParams struct {
	BizTagID       *core.ID    `facebook:"biz_tag_id"`
	BusinessID     *core.ID    `facebook:"business_id"`
	Hashes         *[]string   `facebook:"hashes"`
	Minheight      *uint64     `facebook:"minheight"`
	Minwidth       *uint64     `facebook:"minwidth"`
	Name           *string     `facebook:"name"`
	SelectedHashes *[]string   `facebook:"selected_hashes"`
	Extra          core.Params `facebook:"-"`
}

func (params GetAdAccountAdimagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BizTagID != nil {
		out["biz_tag_id"] = *params.BizTagID
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.Hashes != nil {
		out["hashes"] = *params.Hashes
	}
	if params.Minheight != nil {
		out["minheight"] = *params.Minheight
	}
	if params.Minwidth != nil {
		out["minwidth"] = *params.Minwidth
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SelectedHashes != nil {
		out["selected_hashes"] = *params.SelectedHashes
	}
	return out
}

func GetAdAccountAdimagesBatchCall(id string, params GetAdAccountAdimagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adimages"), params.ToParams(), options...)
}

func NewGetAdAccountAdimagesBatchRequest(id string, params GetAdAccountAdimagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdImage]] {
	return core.NewBatchRequest[core.Cursor[objects.AdImage]](GetAdAccountAdimagesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdimagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdImage], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdImage]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdimages(ctx context.Context, client *core.Client, id string, params GetAdAccountAdimagesParams) (*core.Cursor[objects.AdImage], error) {
	var out core.Cursor[objects.AdImage]
	call := GetAdAccountAdimagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdimagesParams struct {
	Bytes    *string                 `facebook:"bytes"`
	CopyFrom *map[string]interface{} `facebook:"copy_from"`
	Extra    core.Params             `facebook:"-"`
}

func (params CreateAdAccountAdimagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Bytes != nil {
		out["bytes"] = *params.Bytes
	}
	if params.CopyFrom != nil {
		out["copy_from"] = *params.CopyFrom
	}
	return out
}

func CreateAdAccountAdimagesBatchCall(id string, params CreateAdAccountAdimagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adimages"), params.ToParams(), options...)
}

func NewCreateAdAccountAdimagesBatchRequest(id string, params CreateAdAccountAdimagesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdImage] {
	return core.NewBatchRequest[objects.AdImage](CreateAdAccountAdimagesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdimagesBatchResponse(response *core.BatchResponse) (*objects.AdImage, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdImage
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdimages(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdimagesParams) (*objects.AdImage, error) {
	var out objects.AdImage
	call := CreateAdAccountAdimagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdlabelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdlabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdlabelsBatchCall(id string, params GetAdAccountAdlabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adlabels"), params.ToParams(), options...)
}

func NewGetAdAccountAdlabelsBatchRequest(id string, params GetAdAccountAdlabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdLabel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdLabel]](GetAdAccountAdlabelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdlabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdLabel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdLabel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdlabels(ctx context.Context, client *core.Client, id string, params GetAdAccountAdlabelsParams) (*core.Cursor[objects.AdLabel], error) {
	var out core.Cursor[objects.AdLabel]
	call := GetAdAccountAdlabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdlabelsParams struct {
	Name  string      `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountAdlabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["name"] = params.Name
	return out
}

func CreateAdAccountAdlabelsBatchCall(id string, params CreateAdAccountAdlabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), options...)
}

func NewCreateAdAccountAdlabelsBatchRequest(id string, params CreateAdAccountAdlabelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLabel] {
	return core.NewBatchRequest[objects.AdLabel](CreateAdAccountAdlabelsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdlabelsBatchResponse(response *core.BatchResponse) (*objects.AdLabel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLabel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdlabels(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdlabelsParams) (*objects.AdLabel, error) {
	var out objects.AdLabel
	call := CreateAdAccountAdlabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdplayablesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdplayablesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdplayablesBatchCall(id string, params GetAdAccountAdplayablesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adplayables"), params.ToParams(), options...)
}

func NewGetAdAccountAdplayablesBatchRequest(id string, params GetAdAccountAdplayablesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PlayableContent]] {
	return core.NewBatchRequest[core.Cursor[objects.PlayableContent]](GetAdAccountAdplayablesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdplayablesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PlayableContent], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PlayableContent]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdplayables(ctx context.Context, client *core.Client, id string, params GetAdAccountAdplayablesParams) (*core.Cursor[objects.PlayableContent], error) {
	var out core.Cursor[objects.PlayableContent]
	call := GetAdAccountAdplayablesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdplayablesParams struct {
	AppID     *core.ID        `facebook:"app_id"`
	Name      string          `facebook:"name"`
	SessionID *core.ID        `facebook:"session_id"`
	Source    *core.FileParam `facebook:"source"`
	SourceURL *string         `facebook:"source_url"`
	SourceZip *core.FileParam `facebook:"source_zip"`
	Extra     core.Params     `facebook:"-"`
}

func (params CreateAdAccountAdplayablesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	out["name"] = params.Name
	if params.SessionID != nil {
		out["session_id"] = *params.SessionID
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SourceURL != nil {
		out["source_url"] = *params.SourceURL
	}
	if params.SourceZip != nil {
		out["source_zip"] = *params.SourceZip
	}
	return out
}

func CreateAdAccountAdplayablesBatchCall(id string, params CreateAdAccountAdplayablesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adplayables"), params.ToParams(), options...)
}

func NewCreateAdAccountAdplayablesBatchRequest(id string, params CreateAdAccountAdplayablesParams, options ...core.BatchOption) *core.BatchRequest[objects.PlayableContent] {
	return core.NewBatchRequest[objects.PlayableContent](CreateAdAccountAdplayablesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdplayablesBatchResponse(response *core.BatchResponse) (*objects.PlayableContent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PlayableContent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdplayables(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdplayablesParams) (*objects.PlayableContent, error) {
	var out objects.PlayableContent
	call := CreateAdAccountAdplayablesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdrulesHistoryParams struct {
	Action         *enums.AdaccountadrulesHistoryActionEnumParam         `facebook:"action"`
	EvaluationType *enums.AdaccountadrulesHistoryEvaluationTypeEnumParam `facebook:"evaluation_type"`
	HideNoChanges  *bool                                                 `facebook:"hide_no_changes"`
	ObjectID       *core.ID                                              `facebook:"object_id"`
	Extra          core.Params                                           `facebook:"-"`
}

func (params GetAdAccountAdrulesHistoryParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Action != nil {
		out["action"] = *params.Action
	}
	if params.EvaluationType != nil {
		out["evaluation_type"] = *params.EvaluationType
	}
	if params.HideNoChanges != nil {
		out["hide_no_changes"] = *params.HideNoChanges
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	return out
}

func GetAdAccountAdrulesHistoryBatchCall(id string, params GetAdAccountAdrulesHistoryParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adrules_history"), params.ToParams(), options...)
}

func NewGetAdAccountAdrulesHistoryBatchRequest(id string, params GetAdAccountAdrulesHistoryParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountAdRulesHistory]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountAdRulesHistory]](GetAdAccountAdrulesHistoryBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdrulesHistoryBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountAdRulesHistory], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountAdRulesHistory]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdrulesHistory(ctx context.Context, client *core.Client, id string, params GetAdAccountAdrulesHistoryParams) (*core.Cursor[objects.AdAccountAdRulesHistory], error) {
	var out core.Cursor[objects.AdAccountAdRulesHistory]
	call := GetAdAccountAdrulesHistoryBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdrulesLibraryParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdrulesLibraryParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdrulesLibraryBatchCall(id string, params GetAdAccountAdrulesLibraryParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adrules_library"), params.ToParams(), options...)
}

func NewGetAdAccountAdrulesLibraryBatchRequest(id string, params GetAdAccountAdrulesLibraryParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdRule]] {
	return core.NewBatchRequest[core.Cursor[objects.AdRule]](GetAdAccountAdrulesLibraryBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdrulesLibraryBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdRule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdRule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdrulesLibrary(ctx context.Context, client *core.Client, id string, params GetAdAccountAdrulesLibraryParams) (*core.Cursor[objects.AdRule], error) {
	var out core.Cursor[objects.AdRule]
	call := GetAdAccountAdrulesLibraryBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdrulesLibraryParams struct {
	AccountID        *core.ID                                                `facebook:"account_id"`
	EvaluationSpec   map[string]interface{}                                  `facebook:"evaluation_spec"`
	ExecutionSpec    map[string]interface{}                                  `facebook:"execution_spec"`
	Name             string                                                  `facebook:"name"`
	ScheduleSpec     *map[string]interface{}                                 `facebook:"schedule_spec"`
	Status           *enums.AdaccountadrulesLibraryStatusEnumParam           `facebook:"status"`
	UiCreationSource *enums.AdaccountadrulesLibraryUiCreationSourceEnumParam `facebook:"ui_creation_source"`
	Extra            core.Params                                             `facebook:"-"`
}

func (params CreateAdAccountAdrulesLibraryParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	out["evaluation_spec"] = params.EvaluationSpec
	out["execution_spec"] = params.ExecutionSpec
	out["name"] = params.Name
	if params.ScheduleSpec != nil {
		out["schedule_spec"] = *params.ScheduleSpec
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.UiCreationSource != nil {
		out["ui_creation_source"] = *params.UiCreationSource
	}
	return out
}

func CreateAdAccountAdrulesLibraryBatchCall(id string, params CreateAdAccountAdrulesLibraryParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adrules_library"), params.ToParams(), options...)
}

func NewCreateAdAccountAdrulesLibraryBatchRequest(id string, params CreateAdAccountAdrulesLibraryParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRule] {
	return core.NewBatchRequest[objects.AdRule](CreateAdAccountAdrulesLibraryBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdrulesLibraryBatchResponse(response *core.BatchResponse) (*objects.AdRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdrulesLibrary(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdrulesLibraryParams) (*objects.AdRule, error) {
	var out objects.AdRule
	call := CreateAdAccountAdrulesLibraryBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsParams struct {
	DatePreset      *enums.AdaccountadsDatePresetEnumParam `facebook:"date_preset"`
	EffectiveStatus *[]string                              `facebook:"effective_status"`
	TimeRange       *map[string]interface{}                `facebook:"time_range"`
	UpdatedSince    *int                                   `facebook:"updated_since"`
	Extra           core.Params                            `facebook:"-"`
}

func (params GetAdAccountAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.EffectiveStatus != nil {
		out["effective_status"] = *params.EffectiveStatus
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.UpdatedSince != nil {
		out["updated_since"] = *params.UpdatedSince
	}
	return out
}

func GetAdAccountAdsBatchCall(id string, params GetAdAccountAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), options...)
}

func NewGetAdAccountAdsBatchRequest(id string, params GetAdAccountAdsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetAdAccountAdsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Ad]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAds(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	call := GetAdAccountAdsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdsParams struct {
	AdScheduleEndTime       *time.Time                                     `facebook:"ad_schedule_end_time"`
	AdScheduleStartTime     *time.Time                                     `facebook:"ad_schedule_start_time"`
	Adlabels                *[]map[string]interface{}                      `facebook:"adlabels"`
	AdsetID                 *core.ID                                       `facebook:"adset_id"`
	AdsetSpec               *objects.AdSet                                 `facebook:"adset_spec"`
	AudienceID              *core.ID                                       `facebook:"audience_id"`
	BidAmount               *int                                           `facebook:"bid_amount"`
	ConversionDomain        *string                                        `facebook:"conversion_domain"`
	Creative                objects.AdCreative                             `facebook:"creative"`
	CreativeAssetGroupsSpec *map[string]interface{}                        `facebook:"creative_asset_groups_spec"`
	CreativeAutomationSpec  *map[string]interface{}                        `facebook:"creative_automation_spec"`
	DateFormat              *string                                        `facebook:"date_format"`
	DisplaySequence         *uint64                                        `facebook:"display_sequence"`
	DraftAdgroupID          *core.ID                                       `facebook:"draft_adgroup_id"`
	EngagementAudience      *bool                                          `facebook:"engagement_audience"`
	ExecutionOptions        *[]enums.AdaccountadsExecutionOptionsEnumParam `facebook:"execution_options"`
	IncludeDemolinkHashes   *bool                                          `facebook:"include_demolink_hashes"`
	Name                    string                                         `facebook:"name"`
	Priority                *uint64                                        `facebook:"priority"`
	SourceAdID              *core.ID                                       `facebook:"source_ad_id"`
	Status                  *enums.AdaccountadsStatusEnumParam             `facebook:"status"`
	TrackingSpecs           *map[string]interface{}                        `facebook:"tracking_specs"`
	Extra                   core.Params                                    `facebook:"-"`
}

func (params CreateAdAccountAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdScheduleEndTime != nil {
		out["ad_schedule_end_time"] = *params.AdScheduleEndTime
	}
	if params.AdScheduleStartTime != nil {
		out["ad_schedule_start_time"] = *params.AdScheduleStartTime
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.AdsetID != nil {
		out["adset_id"] = *params.AdsetID
	}
	if params.AdsetSpec != nil {
		out["adset_spec"] = *params.AdsetSpec
	}
	if params.AudienceID != nil {
		out["audience_id"] = *params.AudienceID
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.ConversionDomain != nil {
		out["conversion_domain"] = *params.ConversionDomain
	}
	out["creative"] = params.Creative
	if params.CreativeAssetGroupsSpec != nil {
		out["creative_asset_groups_spec"] = *params.CreativeAssetGroupsSpec
	}
	if params.CreativeAutomationSpec != nil {
		out["creative_automation_spec"] = *params.CreativeAutomationSpec
	}
	if params.DateFormat != nil {
		out["date_format"] = *params.DateFormat
	}
	if params.DisplaySequence != nil {
		out["display_sequence"] = *params.DisplaySequence
	}
	if params.DraftAdgroupID != nil {
		out["draft_adgroup_id"] = *params.DraftAdgroupID
	}
	if params.EngagementAudience != nil {
		out["engagement_audience"] = *params.EngagementAudience
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.IncludeDemolinkHashes != nil {
		out["include_demolink_hashes"] = *params.IncludeDemolinkHashes
	}
	out["name"] = params.Name
	if params.Priority != nil {
		out["priority"] = *params.Priority
	}
	if params.SourceAdID != nil {
		out["source_ad_id"] = *params.SourceAdID
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.TrackingSpecs != nil {
		out["tracking_specs"] = *params.TrackingSpecs
	}
	return out
}

func CreateAdAccountAdsBatchCall(id string, params CreateAdAccountAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ads"), params.ToParams(), options...)
}

func NewCreateAdAccountAdsBatchRequest(id string, params CreateAdAccountAdsParams, options ...core.BatchOption) *core.BatchRequest[objects.Ad] {
	return core.NewBatchRequest[objects.Ad](CreateAdAccountAdsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdsBatchResponse(response *core.BatchResponse) (*objects.Ad, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Ad
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAds(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdsParams) (*objects.Ad, error) {
	var out objects.Ad
	call := CreateAdAccountAdsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsReportingMmmReportsParams struct {
	Filtering *[]map[string]interface{} `facebook:"filtering"`
	Extra     core.Params               `facebook:"-"`
}

func (params GetAdAccountAdsReportingMmmReportsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	return out
}

func GetAdAccountAdsReportingMmmReportsBatchCall(id string, params GetAdAccountAdsReportingMmmReportsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_reports"), params.ToParams(), options...)
}

func NewGetAdAccountAdsReportingMmmReportsBatchRequest(id string, params GetAdAccountAdsReportingMmmReportsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsReportBuilderMMMReport]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsReportBuilderMMMReport]](GetAdAccountAdsReportingMmmReportsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsReportingMmmReportsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsReportBuilderMMMReport], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsReportBuilderMMMReport]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsReportingMmmReports(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsReportingMmmReportsParams) (*core.Cursor[objects.AdsReportBuilderMMMReport], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReport]
	call := GetAdAccountAdsReportingMmmReportsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsReportingMmmSchedulersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAdsReportingMmmSchedulersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAdsReportingMmmSchedulersBatchCall(id string, params GetAdAccountAdsReportingMmmSchedulersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_reporting_mmm_schedulers"), params.ToParams(), options...)
}

func NewGetAdAccountAdsReportingMmmSchedulersBatchRequest(id string, params GetAdAccountAdsReportingMmmSchedulersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsReportBuilderMMMReportScheduler]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsReportBuilderMMMReportScheduler]](GetAdAccountAdsReportingMmmSchedulersBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsReportingMmmSchedulersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsReportBuilderMMMReportScheduler], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsReportBuilderMMMReportScheduler]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsReportingMmmSchedulers(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsReportingMmmSchedulersParams) (*core.Cursor[objects.AdsReportBuilderMMMReportScheduler], error) {
	var out core.Cursor[objects.AdsReportBuilderMMMReportScheduler]
	call := GetAdAccountAdsReportingMmmSchedulersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsVolumeParams struct {
	PageID               *core.ID                                             `facebook:"page_id"`
	RecommendationType   *enums.AdaccountadsVolumeRecommendationTypeEnumParam `facebook:"recommendation_type"`
	ShowBreakdownByActor *bool                                                `facebook:"show_breakdown_by_actor"`
	Extra                core.Params                                          `facebook:"-"`
}

func (params GetAdAccountAdsVolumeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	if params.RecommendationType != nil {
		out["recommendation_type"] = *params.RecommendationType
	}
	if params.ShowBreakdownByActor != nil {
		out["show_breakdown_by_actor"] = *params.ShowBreakdownByActor
	}
	return out
}

func GetAdAccountAdsVolumeBatchCall(id string, params GetAdAccountAdsVolumeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads_volume"), params.ToParams(), options...)
}

func NewGetAdAccountAdsVolumeBatchRequest(id string, params GetAdAccountAdsVolumeParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountAdVolume]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountAdVolume]](GetAdAccountAdsVolumeBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsVolumeBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountAdVolume], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountAdVolume]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsVolume(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsVolumeParams) (*core.Cursor[objects.AdAccountAdVolume], error) {
	var out core.Cursor[objects.AdAccountAdVolume]
	call := GetAdAccountAdsVolumeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsbylabelsParams struct {
	AdLabelIds []core.ID                                    `facebook:"ad_label_ids"`
	Operator   *enums.AdaccountadsbylabelsOperatorEnumParam `facebook:"operator"`
	Extra      core.Params                                  `facebook:"-"`
}

func (params GetAdAccountAdsbylabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_label_ids"] = params.AdLabelIds
	if params.Operator != nil {
		out["operator"] = *params.Operator
	}
	return out
}

func GetAdAccountAdsbylabelsBatchCall(id string, params GetAdAccountAdsbylabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsbylabels"), params.ToParams(), options...)
}

func NewGetAdAccountAdsbylabelsBatchRequest(id string, params GetAdAccountAdsbylabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetAdAccountAdsbylabelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsbylabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Ad]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsbylabels(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsbylabelsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	call := GetAdAccountAdsbylabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsetsParams struct {
	DatePreset      *enums.AdaccountadsetsDatePresetEnumParam        `facebook:"date_preset"`
	EffectiveStatus *[]enums.AdaccountadsetsEffectiveStatusEnumParam `facebook:"effective_status"`
	IsCompleted     *bool                                            `facebook:"is_completed"`
	TimeRange       *map[string]interface{}                          `facebook:"time_range"`
	UpdatedSince    *int                                             `facebook:"updated_since"`
	Extra           core.Params                                      `facebook:"-"`
}

func (params GetAdAccountAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.EffectiveStatus != nil {
		out["effective_status"] = *params.EffectiveStatus
	}
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.UpdatedSince != nil {
		out["updated_since"] = *params.UpdatedSince
	}
	return out
}

func GetAdAccountAdsetsBatchCall(id string, params GetAdAccountAdsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), options...)
}

func NewGetAdAccountAdsetsBatchRequest(id string, params GetAdAccountAdsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdAccountAdsetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsets(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdAccountAdsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdsetsParams struct {
	Adlabels                                     *[]map[string]interface{}                                        `facebook:"adlabels"`
	AdsetSchedule                                *[]map[string]interface{}                                        `facebook:"adset_schedule"`
	AttributionCountType                         *enums.AdaccountadsetsAttributionCountTypeEnumParam              `facebook:"attribution_count_type"`
	AttributionSpec                              *[]map[string]interface{}                                        `facebook:"attribution_spec"`
	AutomaticManualState                         *enums.AdaccountadsetsAutomaticManualStateEnumParam              `facebook:"automatic_manual_state"`
	BidAdjustments                               *map[string]interface{}                                          `facebook:"bid_adjustments"`
	BidAmount                                    *int                                                             `facebook:"bid_amount"`
	BidConstraints                               *map[string]map[string]interface{}                               `facebook:"bid_constraints"`
	BidStrategy                                  *enums.AdaccountadsetsBidStrategyEnumParam                       `facebook:"bid_strategy"`
	BillingEvent                                 *enums.AdaccountadsetsBillingEventEnumParam                      `facebook:"billing_event"`
	BudgetScheduleSpecs                          *[]map[string]interface{}                                        `facebook:"budget_schedule_specs"`
	BudgetSource                                 *enums.AdaccountadsetsBudgetSourceEnumParam                      `facebook:"budget_source"`
	BudgetSplitSetID                             *core.ID                                                         `facebook:"budget_split_set_id"`
	CampaignAttribution                          *map[string]interface{}                                          `facebook:"campaign_attribution"`
	CampaignID                                   *core.ID                                                         `facebook:"campaign_id"`
	CampaignSpec                                 *map[string]interface{}                                          `facebook:"campaign_spec"`
	CostBiddingMode                              *enums.AdaccountadsetsCostBiddingModeEnumParam                   `facebook:"cost_bidding_mode"`
	CreativeSequence                             *[]string                                                        `facebook:"creative_sequence"`
	CreativeSequenceRepetitionPattern            *enums.AdaccountadsetsCreativeSequenceRepetitionPatternEnumParam `facebook:"creative_sequence_repetition_pattern"`
	DailyBudget                                  *uint64                                                          `facebook:"daily_budget"`
	DailyImps                                    *uint64                                                          `facebook:"daily_imps"`
	DailyMinSpendTarget                          *uint64                                                          `facebook:"daily_min_spend_target"`
	DailySpendCap                                *uint64                                                          `facebook:"daily_spend_cap"`
	DateFormat                                   *string                                                          `facebook:"date_format"`
	DestinationType                              *enums.AdaccountadsetsDestinationTypeEnumParam                   `facebook:"destination_type"`
	DsaBeneficiary                               *string                                                          `facebook:"dsa_beneficiary"`
	DsaPayor                                     *string                                                          `facebook:"dsa_payor"`
	EndTime                                      *time.Time                                                       `facebook:"end_time"`
	ExecutionOptions                             *[]enums.AdaccountadsetsExecutionOptionsEnumParam                `facebook:"execution_options"`
	ExistingCustomerBudgetPercentage             *uint64                                                          `facebook:"existing_customer_budget_percentage"`
	FrequencyControlSpecs                        *[]map[string]interface{}                                        `facebook:"frequency_control_specs"`
	FullFunnelExplorationMode                    *enums.AdaccountadsetsFullFunnelExplorationModeEnumParam         `facebook:"full_funnel_exploration_mode"`
	IsBaSkipDelayedEligible                      *bool                                                            `facebook:"is_ba_skip_delayed_eligible"`
	IsBudgetScheduleEnabled                      *bool                                                            `facebook:"is_budget_schedule_enabled"`
	IsDcFollowOptimized                          *bool                                                            `facebook:"is_dc_follow_optimized"`
	IsDynamicCreative                            *bool                                                            `facebook:"is_dynamic_creative"`
	IsIncrementalAttributionEnabled              *bool                                                            `facebook:"is_incremental_attribution_enabled"`
	IsSacCfcaTermsCertified                      *bool                                                            `facebook:"is_sac_cfca_terms_certified"`
	LifetimeBudget                               *uint64                                                          `facebook:"lifetime_budget"`
	LifetimeImps                                 *uint64                                                          `facebook:"lifetime_imps"`
	LifetimeMinSpendTarget                       *uint64                                                          `facebook:"lifetime_min_spend_target"`
	LifetimeSpendCap                             *uint64                                                          `facebook:"lifetime_spend_cap"`
	LineNumber                                   *uint64                                                          `facebook:"line_number"`
	LiveVideoAdCampaignConfig                    *map[string]interface{}                                          `facebook:"live_video_ad_campaign_config"`
	MaxBudgetSpendPercentage                     *uint64                                                          `facebook:"max_budget_spend_percentage"`
	MetaMomentMakerSpec                          *map[string]interface{}                                          `facebook:"meta_moment_maker_spec"`
	MinBudgetSpendPercentage                     *uint64                                                          `facebook:"min_budget_spend_percentage"`
	MultiEventConversionAttributionWindowSeconds *uint64                                                          `facebook:"multi_event_conversion_attribution_window_seconds"`
	MultiOptimizationGoalWeight                  *enums.AdaccountadsetsMultiOptimizationGoalWeightEnumParam       `facebook:"multi_optimization_goal_weight"`
	Name                                         string                                                           `facebook:"name"`
	OptimizationGoal                             *enums.AdaccountadsetsOptimizationGoalEnumParam                  `facebook:"optimization_goal"`
	OptimizationSubEvent                         *enums.AdaccountadsetsOptimizationSubEventEnumParam              `facebook:"optimization_sub_event"`
	PacingType                                   *[]string                                                        `facebook:"pacing_type"`
	PlacementSoftOptOut                          *map[string]interface{}                                          `facebook:"placement_soft_opt_out"`
	PromotedObject                               *map[string]interface{}                                          `facebook:"promoted_object"`
	RbPredictionID                               *core.ID                                                         `facebook:"rb_prediction_id"`
	RegionalRegulatedCategories                  *[]enums.AdaccountadsetsRegionalRegulatedCategoriesEnumParam     `facebook:"regional_regulated_categories"`
	RegionalRegulationIdentities                 *map[string]interface{}                                          `facebook:"regional_regulation_identities"`
	RelativeValue                                *float64                                                         `facebook:"relative_value"`
	RfPredictionID                               *core.ID                                                         `facebook:"rf_prediction_id"`
	SourceAdsetID                                *core.ID                                                         `facebook:"source_adset_id"`
	StartTime                                    *time.Time                                                       `facebook:"start_time"`
	Status                                       *enums.AdaccountadsetsStatusEnumParam                            `facebook:"status"`
	Targeting                                    *objects.Targeting                                               `facebook:"targeting"`
	TimeBasedAdRotationIDBlocks                  *[][]uint64                                                      `facebook:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals                 *[]uint64                                                        `facebook:"time_based_ad_rotation_intervals"`
	TimeStart                                    *time.Time                                                       `facebook:"time_start"`
	TimeStop                                     *time.Time                                                       `facebook:"time_stop"`
	ToplineID                                    *core.ID                                                         `facebook:"topline_id"`
	TrendingTopicsSpec                           *map[string]interface{}                                          `facebook:"trending_topics_spec"`
	TuneForCategory                              *enums.AdaccountadsetsTuneForCategoryEnumParam                   `facebook:"tune_for_category"`
	ValueRuleSetID                               *core.ID                                                         `facebook:"value_rule_set_id"`
	ValueRulesApplied                            *bool                                                            `facebook:"value_rules_applied"`
	Extra                                        core.Params                                                      `facebook:"-"`
}

func (params CreateAdAccountAdsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.AdsetSchedule != nil {
		out["adset_schedule"] = *params.AdsetSchedule
	}
	if params.AttributionCountType != nil {
		out["attribution_count_type"] = *params.AttributionCountType
	}
	if params.AttributionSpec != nil {
		out["attribution_spec"] = *params.AttributionSpec
	}
	if params.AutomaticManualState != nil {
		out["automatic_manual_state"] = *params.AutomaticManualState
	}
	if params.BidAdjustments != nil {
		out["bid_adjustments"] = *params.BidAdjustments
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.BidConstraints != nil {
		out["bid_constraints"] = *params.BidConstraints
	}
	if params.BidStrategy != nil {
		out["bid_strategy"] = *params.BidStrategy
	}
	if params.BillingEvent != nil {
		out["billing_event"] = *params.BillingEvent
	}
	if params.BudgetScheduleSpecs != nil {
		out["budget_schedule_specs"] = *params.BudgetScheduleSpecs
	}
	if params.BudgetSource != nil {
		out["budget_source"] = *params.BudgetSource
	}
	if params.BudgetSplitSetID != nil {
		out["budget_split_set_id"] = *params.BudgetSplitSetID
	}
	if params.CampaignAttribution != nil {
		out["campaign_attribution"] = *params.CampaignAttribution
	}
	if params.CampaignID != nil {
		out["campaign_id"] = *params.CampaignID
	}
	if params.CampaignSpec != nil {
		out["campaign_spec"] = *params.CampaignSpec
	}
	if params.CostBiddingMode != nil {
		out["cost_bidding_mode"] = *params.CostBiddingMode
	}
	if params.CreativeSequence != nil {
		out["creative_sequence"] = *params.CreativeSequence
	}
	if params.CreativeSequenceRepetitionPattern != nil {
		out["creative_sequence_repetition_pattern"] = *params.CreativeSequenceRepetitionPattern
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.DailyImps != nil {
		out["daily_imps"] = *params.DailyImps
	}
	if params.DailyMinSpendTarget != nil {
		out["daily_min_spend_target"] = *params.DailyMinSpendTarget
	}
	if params.DailySpendCap != nil {
		out["daily_spend_cap"] = *params.DailySpendCap
	}
	if params.DateFormat != nil {
		out["date_format"] = *params.DateFormat
	}
	if params.DestinationType != nil {
		out["destination_type"] = *params.DestinationType
	}
	if params.DsaBeneficiary != nil {
		out["dsa_beneficiary"] = *params.DsaBeneficiary
	}
	if params.DsaPayor != nil {
		out["dsa_payor"] = *params.DsaPayor
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.ExistingCustomerBudgetPercentage != nil {
		out["existing_customer_budget_percentage"] = *params.ExistingCustomerBudgetPercentage
	}
	if params.FrequencyControlSpecs != nil {
		out["frequency_control_specs"] = *params.FrequencyControlSpecs
	}
	if params.FullFunnelExplorationMode != nil {
		out["full_funnel_exploration_mode"] = *params.FullFunnelExplorationMode
	}
	if params.IsBaSkipDelayedEligible != nil {
		out["is_ba_skip_delayed_eligible"] = *params.IsBaSkipDelayedEligible
	}
	if params.IsBudgetScheduleEnabled != nil {
		out["is_budget_schedule_enabled"] = *params.IsBudgetScheduleEnabled
	}
	if params.IsDcFollowOptimized != nil {
		out["is_dc_follow_optimized"] = *params.IsDcFollowOptimized
	}
	if params.IsDynamicCreative != nil {
		out["is_dynamic_creative"] = *params.IsDynamicCreative
	}
	if params.IsIncrementalAttributionEnabled != nil {
		out["is_incremental_attribution_enabled"] = *params.IsIncrementalAttributionEnabled
	}
	if params.IsSacCfcaTermsCertified != nil {
		out["is_sac_cfca_terms_certified"] = *params.IsSacCfcaTermsCertified
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.LifetimeImps != nil {
		out["lifetime_imps"] = *params.LifetimeImps
	}
	if params.LifetimeMinSpendTarget != nil {
		out["lifetime_min_spend_target"] = *params.LifetimeMinSpendTarget
	}
	if params.LifetimeSpendCap != nil {
		out["lifetime_spend_cap"] = *params.LifetimeSpendCap
	}
	if params.LineNumber != nil {
		out["line_number"] = *params.LineNumber
	}
	if params.LiveVideoAdCampaignConfig != nil {
		out["live_video_ad_campaign_config"] = *params.LiveVideoAdCampaignConfig
	}
	if params.MaxBudgetSpendPercentage != nil {
		out["max_budget_spend_percentage"] = *params.MaxBudgetSpendPercentage
	}
	if params.MetaMomentMakerSpec != nil {
		out["meta_moment_maker_spec"] = *params.MetaMomentMakerSpec
	}
	if params.MinBudgetSpendPercentage != nil {
		out["min_budget_spend_percentage"] = *params.MinBudgetSpendPercentage
	}
	if params.MultiEventConversionAttributionWindowSeconds != nil {
		out["multi_event_conversion_attribution_window_seconds"] = *params.MultiEventConversionAttributionWindowSeconds
	}
	if params.MultiOptimizationGoalWeight != nil {
		out["multi_optimization_goal_weight"] = *params.MultiOptimizationGoalWeight
	}
	out["name"] = params.Name
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.OptimizationSubEvent != nil {
		out["optimization_sub_event"] = *params.OptimizationSubEvent
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	if params.PlacementSoftOptOut != nil {
		out["placement_soft_opt_out"] = *params.PlacementSoftOptOut
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.RbPredictionID != nil {
		out["rb_prediction_id"] = *params.RbPredictionID
	}
	if params.RegionalRegulatedCategories != nil {
		out["regional_regulated_categories"] = *params.RegionalRegulatedCategories
	}
	if params.RegionalRegulationIdentities != nil {
		out["regional_regulation_identities"] = *params.RegionalRegulationIdentities
	}
	if params.RelativeValue != nil {
		out["relative_value"] = *params.RelativeValue
	}
	if params.RfPredictionID != nil {
		out["rf_prediction_id"] = *params.RfPredictionID
	}
	if params.SourceAdsetID != nil {
		out["source_adset_id"] = *params.SourceAdsetID
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TimeBasedAdRotationIDBlocks != nil {
		out["time_based_ad_rotation_id_blocks"] = *params.TimeBasedAdRotationIDBlocks
	}
	if params.TimeBasedAdRotationIntervals != nil {
		out["time_based_ad_rotation_intervals"] = *params.TimeBasedAdRotationIntervals
	}
	if params.TimeStart != nil {
		out["time_start"] = *params.TimeStart
	}
	if params.TimeStop != nil {
		out["time_stop"] = *params.TimeStop
	}
	if params.ToplineID != nil {
		out["topline_id"] = *params.ToplineID
	}
	if params.TrendingTopicsSpec != nil {
		out["trending_topics_spec"] = *params.TrendingTopicsSpec
	}
	if params.TuneForCategory != nil {
		out["tune_for_category"] = *params.TuneForCategory
	}
	if params.ValueRuleSetID != nil {
		out["value_rule_set_id"] = *params.ValueRuleSetID
	}
	if params.ValueRulesApplied != nil {
		out["value_rules_applied"] = *params.ValueRulesApplied
	}
	return out
}

func CreateAdAccountAdsetsBatchCall(id string, params CreateAdAccountAdsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adsets"), params.ToParams(), options...)
}

func NewCreateAdAccountAdsetsBatchRequest(id string, params CreateAdAccountAdsetsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdSet] {
	return core.NewBatchRequest[objects.AdSet](CreateAdAccountAdsetsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdsetsBatchResponse(response *core.BatchResponse) (*objects.AdSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdsets(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdsetsParams) (*objects.AdSet, error) {
	var out objects.AdSet
	call := CreateAdAccountAdsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdsetsbylabelsParams struct {
	AdLabelIds []core.ID                                       `facebook:"ad_label_ids"`
	Operator   *enums.AdaccountadsetsbylabelsOperatorEnumParam `facebook:"operator"`
	Extra      core.Params                                     `facebook:"-"`
}

func (params GetAdAccountAdsetsbylabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_label_ids"] = params.AdLabelIds
	if params.Operator != nil {
		out["operator"] = *params.Operator
	}
	return out
}

func GetAdAccountAdsetsbylabelsBatchCall(id string, params GetAdAccountAdsetsbylabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsetsbylabels"), params.ToParams(), options...)
}

func NewGetAdAccountAdsetsbylabelsBatchRequest(id string, params GetAdAccountAdsetsbylabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdAccountAdsetsbylabelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdsetsbylabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdsetsbylabels(ctx context.Context, client *core.Client, id string, params GetAdAccountAdsetsbylabelsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdAccountAdsetsbylabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdspixelsParams struct {
	SortBy *enums.AdaccountadspixelsSortByEnumParam `facebook:"sort_by"`
	Extra  core.Params                              `facebook:"-"`
}

func (params GetAdAccountAdspixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	return out
}

func GetAdAccountAdspixelsBatchCall(id string, params GetAdAccountAdspixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adspixels"), params.ToParams(), options...)
}

func NewGetAdAccountAdspixelsBatchRequest(id string, params GetAdAccountAdspixelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsPixel]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsPixel]](GetAdAccountAdspixelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdspixelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsPixel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsPixel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdspixels(ctx context.Context, client *core.Client, id string, params GetAdAccountAdspixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	call := GetAdAccountAdspixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdspixelsParams struct {
	Name  *string     `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountAdspixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func CreateAdAccountAdspixelsBatchCall(id string, params CreateAdAccountAdspixelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adspixels"), params.ToParams(), options...)
}

func NewCreateAdAccountAdspixelsBatchRequest(id string, params CreateAdAccountAdspixelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixel] {
	return core.NewBatchRequest[objects.AdsPixel](CreateAdAccountAdspixelsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdspixelsBatchResponse(response *core.BatchResponse) (*objects.AdsPixel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAdspixels(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdspixelsParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	call := CreateAdAccountAdspixelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdvertisableApplicationsParams struct {
	AppID      *core.ID    `facebook:"app_id"`
	BusinessID *core.ID    `facebook:"business_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetAdAccountAdvertisableApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	return out
}

func GetAdAccountAdvertisableApplicationsBatchCall(id string, params GetAdAccountAdvertisableApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "advertisable_applications"), params.ToParams(), options...)
}

func NewGetAdAccountAdvertisableApplicationsBatchRequest(id string, params GetAdAccountAdvertisableApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetAdAccountAdvertisableApplicationsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdvertisableApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAdvertisableApplications(ctx context.Context, client *core.Client, id string, params GetAdAccountAdvertisableApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetAdAccountAdvertisableApplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountAdvideosParams struct {
	VideoID core.ID     `facebook:"video_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteAdAccountAdvideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["video_id"] = params.VideoID
	return out
}

func DeleteAdAccountAdvideosBatchCall(id string, params DeleteAdAccountAdvideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "advideos"), params.ToParams(), options...)
}

func NewDeleteAdAccountAdvideosBatchRequest(id string, params DeleteAdAccountAdvideosParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountAdvideosBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountAdvideosBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountAdvideos(ctx context.Context, client *core.Client, id string, params DeleteAdAccountAdvideosParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountAdvideosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAdvideosParams struct {
	MaxAspectRatio *float64    `facebook:"max_aspect_ratio"`
	Maxheight      *uint64     `facebook:"maxheight"`
	Maxlength      *uint64     `facebook:"maxlength"`
	Maxwidth       *uint64     `facebook:"maxwidth"`
	MinAspectRatio *float64    `facebook:"min_aspect_ratio"`
	Minheight      *uint64     `facebook:"minheight"`
	Minlength      *uint64     `facebook:"minlength"`
	Minwidth       *uint64     `facebook:"minwidth"`
	Title          *string     `facebook:"title"`
	Extra          core.Params `facebook:"-"`
}

func (params GetAdAccountAdvideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.MaxAspectRatio != nil {
		out["max_aspect_ratio"] = *params.MaxAspectRatio
	}
	if params.Maxheight != nil {
		out["maxheight"] = *params.Maxheight
	}
	if params.Maxlength != nil {
		out["maxlength"] = *params.Maxlength
	}
	if params.Maxwidth != nil {
		out["maxwidth"] = *params.Maxwidth
	}
	if params.MinAspectRatio != nil {
		out["min_aspect_ratio"] = *params.MinAspectRatio
	}
	if params.Minheight != nil {
		out["minheight"] = *params.Minheight
	}
	if params.Minlength != nil {
		out["minlength"] = *params.Minlength
	}
	if params.Minwidth != nil {
		out["minwidth"] = *params.Minwidth
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func GetAdAccountAdvideosBatchCall(id string, params GetAdAccountAdvideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "advideos"), params.ToParams(), options...)
}

func NewGetAdAccountAdvideosBatchRequest(id string, params GetAdAccountAdvideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetAdAccountAdvideosBatchCall(id, params, options...))
}

func DecodeGetAdAccountAdvideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
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

func GetAdAccountAdvideos(ctx context.Context, client *core.Client, id string, params GetAdAccountAdvideosParams) (*core.Cursor[objects.AdVideo], error) {
	var out core.Cursor[objects.AdVideo]
	call := GetAdAccountAdvideosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAdvideosParams struct {
	ApplicationID                 *core.ID                                                `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                                `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                                 `facebook:"audio_story_wave_animation_handle"`
	ChunkSessionID                *core.ID                                                `facebook:"chunk_session_id"`
	ComposerEntryPicker           *string                                                 `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                                 `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                                 `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                                 `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                                `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                                 `facebook:"composer_source_surface"`
	ComposerType                  *string                                                 `facebook:"composer_type"`
	ContainerType                 *enums.AdaccountadvideosContainerTypeEnumParam          `facebook:"container_type"`
	ContentCategory               *enums.AdaccountadvideosContentCategoryEnumParam        `facebook:"content_category"`
	CreativeTools                 *string                                                 `facebook:"creative_tools"`
	Description                   *string                                                 `facebook:"description"`
	EditDescriptionSpec           *map[string]interface{}                                 `facebook:"edit_description_spec"`
	Embeddable                    *bool                                                   `facebook:"embeddable"`
	EndOffset                     *uint64                                                 `facebook:"end_offset"`
	FbuploaderVideoFileChunk      *string                                                 `facebook:"fbuploader_video_file_chunk"`
	FileSize                      *uint64                                                 `facebook:"file_size"`
	FileURL                       *string                                                 `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                                   `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.AdaccountadvideosFormattingEnumParam             `facebook:"formatting"`
	Fov                           *uint64                                                 `facebook:"fov"`
	FrontZRotation                *float64                                                `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                                `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                                `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                             `facebook:"guide"`
	GuideEnabled                  *bool                                                   `facebook:"guide_enabled"`
	InitialHeading                *uint64                                                 `facebook:"initial_heading"`
	InitialPitch                  *uint64                                                 `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                                 `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                                   `facebook:"is_boost_intended"`
	IsGroupLinkingPost            *bool                                                   `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                                   `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                                   `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                                `facebook:"location_source_id"`
	Name                          *string                                                 `facebook:"name"`
	OgActionTypeID                *core.ID                                                `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                                `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                                `facebook:"og_object_id"`
	OgPhrase                      *string                                                 `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                                 `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                                 `facebook:"original_fov"`
	OriginalProjectionType        *enums.AdaccountadvideosOriginalProjectionTypeEnumParam `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                                 `facebook:"partnership_ad_ad_code"`
	PublishEventID                *core.ID                                                `facebook:"publish_event_id"`
	ReferencedStickerID           *core.ID                                                `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                                `facebook:"replace_video_id"`
	SelectedAudioSpec             *map[string]interface{}                                 `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                                 `facebook:"slideshow_spec"`
	Source                        *string                                                 `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                                `facebook:"source_instagram_media_id"`
	Spherical                     *bool                                                   `facebook:"spherical"`
	StartOffset                   *uint64                                                 `facebook:"start_offset"`
	SwapMode                      *enums.AdaccountadvideosSwapModeEnumParam               `facebook:"swap_mode"`
	TextFormatMetadata            *string                                                 `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                         `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                                 `facebook:"time_since_original_post"`
	Title                         *string                                                 `facebook:"title"`
	TranscodeSettingProperties    *string                                                 `facebook:"transcode_setting_properties"`
	UnpublishedContentType        *enums.AdaccountadvideosUnpublishedContentTypeEnumParam `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.AdaccountadvideosUploadPhaseEnumParam            `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                                `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                                 `facebook:"upload_setting_properties"`
	VideoFileChunk                *string                                                 `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                                 `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                                 `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                                `facebook:"waterfall_id"`
	Extra                         core.Params                                             `facebook:"-"`
}

func (params CreateAdAccountAdvideosParams) ToParams() core.Params {
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
	if params.ChunkSessionID != nil {
		out["chunk_session_id"] = *params.ChunkSessionID
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
	if params.Name != nil {
		out["name"] = *params.Name
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
	if params.ReferencedStickerID != nil {
		out["referenced_sticker_id"] = *params.ReferencedStickerID
	}
	if params.ReplaceVideoID != nil {
		out["replace_video_id"] = *params.ReplaceVideoID
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

func CreateAdAccountAdvideosBatchCall(id string, params CreateAdAccountAdvideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "advideos"), params.ToParams(), options...)
}

func NewCreateAdAccountAdvideosBatchRequest(id string, params CreateAdAccountAdvideosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdAccountAdvideosBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAdvideosBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
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

func CreateAdAccountAdvideos(ctx context.Context, client *core.Client, id string, params CreateAdAccountAdvideosParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	call := CreateAdAccountAdvideosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAffectedadsetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAffectedadsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAffectedadsetsBatchCall(id string, params GetAdAccountAffectedadsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "affectedadsets"), params.ToParams(), options...)
}

func NewGetAdAccountAffectedadsetsBatchRequest(id string, params GetAdAccountAffectedadsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdAccountAffectedadsetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAffectedadsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAffectedadsets(ctx context.Context, client *core.Client, id string, params GetAdAccountAffectedadsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdAccountAffectedadsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteAdAccountAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteAdAccountAgenciesBatchCall(id string, params DeleteAdAccountAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewDeleteAdAccountAgenciesBatchRequest(id string, params DeleteAdAccountAgenciesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountAgenciesBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountAgenciesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountAgencies(ctx context.Context, client *core.Client, id string, params DeleteAdAccountAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountAgenciesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAgenciesBatchCall(id string, params GetAdAccountAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetAdAccountAgenciesBatchRequest(id string, params GetAdAccountAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetAdAccountAgenciesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAgencies(ctx context.Context, client *core.Client, id string, params GetAdAccountAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	call := GetAdAccountAgenciesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAgenciesParams struct {
	Business       string                                            `facebook:"business"`
	PermittedTasks *[]enums.AdaccountagenciesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                       `facebook:"-"`
}

func (params CreateAdAccountAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	if params.PermittedTasks != nil {
		out["permitted_tasks"] = *params.PermittedTasks
	}
	return out
}

func CreateAdAccountAgenciesBatchCall(id string, params CreateAdAccountAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewCreateAdAccountAgenciesBatchRequest(id string, params CreateAdAccountAgenciesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountAgenciesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAgenciesBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAgencies(ctx context.Context, client *core.Client, id string, params CreateAdAccountAgenciesParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountAgenciesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountApplicationsBatchCall(id string, params GetAdAccountApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "applications"), params.ToParams(), options...)
}

func NewGetAdAccountApplicationsBatchRequest(id string, params GetAdAccountApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetAdAccountApplicationsBatchCall(id, params, options...))
}

func DecodeGetAdAccountApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountApplications(ctx context.Context, client *core.Client, id string, params GetAdAccountApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetAdAccountApplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteAdAccountAssignedUsersBatchCall(id string, params DeleteAdAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewDeleteAdAccountAssignedUsersBatchRequest(id string, params DeleteAdAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteAdAccountAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdAccountAssignedUsersBatchCall(id string, params GetAdAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewGetAdAccountAssignedUsersBatchRequest(id string, params GetAdAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AssignedUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AssignedUser]](GetAdAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeGetAdAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AssignedUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AssignedUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params GetAdAccountAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	call := GetAdAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAssignedUsersParams struct {
	Tasks *[]enums.AdaccountassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                           `facebook:"user"`
	Extra core.Params                                   `facebook:"-"`
}

func (params CreateAdAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	out["user"] = params.User
	return out
}

func CreateAdAccountAssignedUsersBatchCall(id string, params CreateAdAccountAssignedUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), options...)
}

func NewCreateAdAccountAssignedUsersBatchRequest(id string, params CreateAdAccountAssignedUsersParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountAssignedUsersBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAssignedUsersBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateAdAccountAssignedUsersParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountAssignedUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAsyncBatchRequestsParams struct {
	Adbatch []map[string]interface{} `facebook:"adbatch"`
	Name    string                   `facebook:"name"`
	Extra   core.Params              `facebook:"-"`
}

func (params CreateAdAccountAsyncBatchRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["adbatch"] = params.Adbatch
	out["name"] = params.Name
	return out
}

func CreateAdAccountAsyncBatchRequestsBatchCall(id string, params CreateAdAccountAsyncBatchRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "async_batch_requests"), params.ToParams(), options...)
}

func NewCreateAdAccountAsyncBatchRequestsBatchRequest(id string, params CreateAdAccountAsyncBatchRequestsParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](CreateAdAccountAsyncBatchRequestsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAsyncBatchRequestsBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAsyncBatchRequests(ctx context.Context, client *core.Client, id string, params CreateAdAccountAsyncBatchRequestsParams) (*objects.Campaign, error) {
	var out objects.Campaign
	call := CreateAdAccountAsyncBatchRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAsyncRequestsParams struct {
	Status *enums.AdaccountasyncRequestsStatusEnumParam `facebook:"status"`
	Type   *enums.AdaccountasyncRequestsTypeEnumParam   `facebook:"type"`
	Extra  core.Params                                  `facebook:"-"`
}

func (params GetAdAccountAsyncRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetAdAccountAsyncRequestsBatchCall(id string, params GetAdAccountAsyncRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "async_requests"), params.ToParams(), options...)
}

func NewGetAdAccountAsyncRequestsBatchRequest(id string, params GetAdAccountAsyncRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AsyncRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.AsyncRequest]](GetAdAccountAsyncRequestsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAsyncRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AsyncRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AsyncRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAsyncRequests(ctx context.Context, client *core.Client, id string, params GetAdAccountAsyncRequestsParams) (*core.Cursor[objects.AsyncRequest], error) {
	var out core.Cursor[objects.AsyncRequest]
	call := GetAdAccountAsyncRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAsyncadcreativesParams struct {
	IsCompleted *bool       `facebook:"is_completed"`
	Extra       core.Params `facebook:"-"`
}

func (params GetAdAccountAsyncadcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	return out
}

func GetAdAccountAsyncadcreativesBatchCall(id string, params GetAdAccountAsyncadcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "asyncadcreatives"), params.ToParams(), options...)
}

func NewGetAdAccountAsyncadcreativesBatchRequest(id string, params GetAdAccountAsyncadcreativesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAsyncRequestSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAsyncRequestSet]](GetAdAccountAsyncadcreativesBatchCall(id, params, options...))
}

func DecodeGetAdAccountAsyncadcreativesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAsyncRequestSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAsyncRequestSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAsyncadcreatives(ctx context.Context, client *core.Client, id string, params GetAdAccountAsyncadcreativesParams) (*core.Cursor[objects.AdAsyncRequestSet], error) {
	var out core.Cursor[objects.AdAsyncRequestSet]
	call := GetAdAccountAsyncadcreativesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAsyncadcreativesParams struct {
	CreativeSpec     objects.AdCreative                                        `facebook:"creative_spec"`
	Name             string                                                    `facebook:"name"`
	NotificationMode *enums.AdaccountasyncadcreativesNotificationModeEnumParam `facebook:"notification_mode"`
	NotificationURI  *string                                                   `facebook:"notification_uri"`
	Extra            core.Params                                               `facebook:"-"`
}

func (params CreateAdAccountAsyncadcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["creative_spec"] = params.CreativeSpec
	out["name"] = params.Name
	if params.NotificationMode != nil {
		out["notification_mode"] = *params.NotificationMode
	}
	if params.NotificationURI != nil {
		out["notification_uri"] = *params.NotificationURI
	}
	return out
}

func CreateAdAccountAsyncadcreativesBatchCall(id string, params CreateAdAccountAsyncadcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "asyncadcreatives"), params.ToParams(), options...)
}

func NewCreateAdAccountAsyncadcreativesBatchRequest(id string, params CreateAdAccountAsyncadcreativesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAsyncRequestSet] {
	return core.NewBatchRequest[objects.AdAsyncRequestSet](CreateAdAccountAsyncadcreativesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAsyncadcreativesBatchResponse(response *core.BatchResponse) (*objects.AdAsyncRequestSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAsyncRequestSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAsyncadcreatives(ctx context.Context, client *core.Client, id string, params CreateAdAccountAsyncadcreativesParams) (*objects.AdAsyncRequestSet, error) {
	var out objects.AdAsyncRequestSet
	call := CreateAdAccountAsyncadcreativesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAsyncadrequestsetsParams struct {
	IsCompleted *bool       `facebook:"is_completed"`
	Extra       core.Params `facebook:"-"`
}

func (params GetAdAccountAsyncadrequestsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	return out
}

func GetAdAccountAsyncadrequestsetsBatchCall(id string, params GetAdAccountAsyncadrequestsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "asyncadrequestsets"), params.ToParams(), options...)
}

func NewGetAdAccountAsyncadrequestsetsBatchRequest(id string, params GetAdAccountAsyncadrequestsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAsyncRequestSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAsyncRequestSet]](GetAdAccountAsyncadrequestsetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountAsyncadrequestsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAsyncRequestSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAsyncRequestSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAsyncadrequestsets(ctx context.Context, client *core.Client, id string, params GetAdAccountAsyncadrequestsetsParams) (*core.Cursor[objects.AdAsyncRequestSet], error) {
	var out core.Cursor[objects.AdAsyncRequestSet]
	call := GetAdAccountAsyncadrequestsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountAsyncadrequestsetsParams struct {
	AdSpecs          []map[string]interface{}                                    `facebook:"ad_specs"`
	Name             string                                                      `facebook:"name"`
	NotificationMode *enums.AdaccountasyncadrequestsetsNotificationModeEnumParam `facebook:"notification_mode"`
	NotificationURI  *string                                                     `facebook:"notification_uri"`
	Extra            core.Params                                                 `facebook:"-"`
}

func (params CreateAdAccountAsyncadrequestsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_specs"] = params.AdSpecs
	out["name"] = params.Name
	if params.NotificationMode != nil {
		out["notification_mode"] = *params.NotificationMode
	}
	if params.NotificationURI != nil {
		out["notification_uri"] = *params.NotificationURI
	}
	return out
}

func CreateAdAccountAsyncadrequestsetsBatchCall(id string, params CreateAdAccountAsyncadrequestsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "asyncadrequestsets"), params.ToParams(), options...)
}

func NewCreateAdAccountAsyncadrequestsetsBatchRequest(id string, params CreateAdAccountAsyncadrequestsetsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAsyncRequestSet] {
	return core.NewBatchRequest[objects.AdAsyncRequestSet](CreateAdAccountAsyncadrequestsetsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAsyncadrequestsetsBatchResponse(response *core.BatchResponse) (*objects.AdAsyncRequestSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAsyncRequestSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAsyncadrequestsets(ctx context.Context, client *core.Client, id string, params CreateAdAccountAsyncadrequestsetsParams) (*objects.AdAsyncRequestSet, error) {
	var out objects.AdAsyncRequestSet
	call := CreateAdAccountAsyncadrequestsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountAudienceFunnelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAudienceFunnelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAudienceFunnelBatchCall(id string, params GetAdAccountAudienceFunnelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "audience_funnel"), params.ToParams(), options...)
}

func NewGetAdAccountAudienceFunnelBatchRequest(id string, params GetAdAccountAudienceFunnelParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AudienceFunnel]] {
	return core.NewBatchRequest[core.Cursor[objects.AudienceFunnel]](GetAdAccountAudienceFunnelBatchCall(id, params, options...))
}

func DecodeGetAdAccountAudienceFunnelBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AudienceFunnel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AudienceFunnel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountAudienceFunnel(ctx context.Context, client *core.Client, id string, params GetAdAccountAudienceFunnelParams) (*core.Cursor[objects.AudienceFunnel], error) {
	var out core.Cursor[objects.AudienceFunnel]
	call := GetAdAccountAudienceFunnelBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountBlockListDraftsParams struct {
	PublisherUrlsFile core.FileParam `facebook:"publisher_urls_file"`
	Extra             core.Params    `facebook:"-"`
}

func (params CreateAdAccountBlockListDraftsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["publisher_urls_file"] = params.PublisherUrlsFile
	return out
}

func CreateAdAccountBlockListDraftsBatchCall(id string, params CreateAdAccountBlockListDraftsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "block_list_drafts"), params.ToParams(), options...)
}

func NewCreateAdAccountBlockListDraftsBatchRequest(id string, params CreateAdAccountBlockListDraftsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountBlockListDraftsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountBlockListDraftsBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountBlockListDrafts(ctx context.Context, client *core.Client, id string, params CreateAdAccountBlockListDraftsParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountBlockListDraftsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountBrandSafetyContentFilterLevelsParams struct {
	BrandSafetyContentFilterLevels []enums.AdaccountbrandSafetyContentFilterLevelsBrandSafetyContentFilterLevelsEnumParam `facebook:"brand_safety_content_filter_levels"`
	BusinessID                     *core.ID                                                                               `facebook:"business_id"`
	Extra                          core.Params                                                                            `facebook:"-"`
}

func (params CreateAdAccountBrandSafetyContentFilterLevelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["brand_safety_content_filter_levels"] = params.BrandSafetyContentFilterLevels
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	return out
}

func CreateAdAccountBrandSafetyContentFilterLevelsBatchCall(id string, params CreateAdAccountBrandSafetyContentFilterLevelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "brand_safety_content_filter_levels"), params.ToParams(), options...)
}

func NewCreateAdAccountBrandSafetyContentFilterLevelsBatchRequest(id string, params CreateAdAccountBrandSafetyContentFilterLevelsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountBrandSafetyContentFilterLevelsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountBrandSafetyContentFilterLevelsBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountBrandSafetyContentFilterLevels(ctx context.Context, client *core.Client, id string, params CreateAdAccountBrandSafetyContentFilterLevelsParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountBrandSafetyContentFilterLevelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountBroadtargetingcategoriesParams struct {
	CustomCategoriesOnly *bool       `facebook:"custom_categories_only"`
	Extra                core.Params `facebook:"-"`
}

func (params GetAdAccountBroadtargetingcategoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CustomCategoriesOnly != nil {
		out["custom_categories_only"] = *params.CustomCategoriesOnly
	}
	return out
}

func GetAdAccountBroadtargetingcategoriesBatchCall(id string, params GetAdAccountBroadtargetingcategoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "broadtargetingcategories"), params.ToParams(), options...)
}

func NewGetAdAccountBroadtargetingcategoriesBatchRequest(id string, params GetAdAccountBroadtargetingcategoriesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BroadTargetingCategories]] {
	return core.NewBatchRequest[core.Cursor[objects.BroadTargetingCategories]](GetAdAccountBroadtargetingcategoriesBatchCall(id, params, options...))
}

func DecodeGetAdAccountBroadtargetingcategoriesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BroadTargetingCategories], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BroadTargetingCategories]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountBroadtargetingcategories(ctx context.Context, client *core.Client, id string, params GetAdAccountBroadtargetingcategoriesParams) (*core.Cursor[objects.BroadTargetingCategories], error) {
	var out core.Cursor[objects.BroadTargetingCategories]
	call := GetAdAccountBroadtargetingcategoriesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountBusinessprojectsParams struct {
	Business *string     `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdAccountBusinessprojectsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Business != nil {
		out["business"] = *params.Business
	}
	return out
}

func GetAdAccountBusinessprojectsBatchCall(id string, params GetAdAccountBusinessprojectsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "businessprojects"), params.ToParams(), options...)
}

func NewGetAdAccountBusinessprojectsBatchRequest(id string, params GetAdAccountBusinessprojectsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessProject]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessProject]](GetAdAccountBusinessprojectsBatchCall(id, params, options...))
}

func DecodeGetAdAccountBusinessprojectsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessProject], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessProject]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountBusinessprojects(ctx context.Context, client *core.Client, id string, params GetAdAccountBusinessprojectsParams) (*core.Cursor[objects.BusinessProject], error) {
	var out core.Cursor[objects.BusinessProject]
	call := GetAdAccountBusinessprojectsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountCampaignsParams struct {
	BeforeDate     *time.Time                                      `facebook:"before_date"`
	DeleteOffset   *uint64                                         `facebook:"delete_offset"`
	DeleteStrategy enums.AdaccountcampaignsDeleteStrategyEnumParam `facebook:"delete_strategy"`
	ObjectCount    *int                                            `facebook:"object_count"`
	Extra          core.Params                                     `facebook:"-"`
}

func (params DeleteAdAccountCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BeforeDate != nil {
		out["before_date"] = *params.BeforeDate
	}
	if params.DeleteOffset != nil {
		out["delete_offset"] = *params.DeleteOffset
	}
	out["delete_strategy"] = params.DeleteStrategy
	if params.ObjectCount != nil {
		out["object_count"] = *params.ObjectCount
	}
	return out
}

func DeleteAdAccountCampaignsBatchCall(id string, params DeleteAdAccountCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewDeleteAdAccountCampaignsBatchRequest(id string, params DeleteAdAccountCampaignsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountCampaignsBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountCampaignsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountCampaigns(ctx context.Context, client *core.Client, id string, params DeleteAdAccountCampaignsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCampaignsParams struct {
	DatePreset      *enums.AdaccountcampaignsDatePresetEnumParam        `facebook:"date_preset"`
	EffectiveStatus *[]enums.AdaccountcampaignsEffectiveStatusEnumParam `facebook:"effective_status"`
	IsCompleted     *bool                                               `facebook:"is_completed"`
	TimeRange       *map[string]interface{}                             `facebook:"time_range"`
	Extra           core.Params                                         `facebook:"-"`
}

func (params GetAdAccountCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.EffectiveStatus != nil {
		out["effective_status"] = *params.EffectiveStatus
	}
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetAdAccountCampaignsBatchCall(id string, params GetAdAccountCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewGetAdAccountCampaignsBatchRequest(id string, params GetAdAccountCampaignsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Campaign]] {
	return core.NewBatchRequest[core.Cursor[objects.Campaign]](GetAdAccountCampaignsBatchCall(id, params, options...))
}

func DecodeGetAdAccountCampaignsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Campaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Campaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCampaigns(ctx context.Context, client *core.Client, id string, params GetAdAccountCampaignsParams) (*core.Cursor[objects.Campaign], error) {
	var out core.Cursor[objects.Campaign]
	call := GetAdAccountCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountCampaignsParams struct {
	Adlabels                    *[]map[string]interface{}                                    `facebook:"adlabels"`
	BidStrategy                 *enums.AdaccountcampaignsBidStrategyEnumParam                `facebook:"bid_strategy"`
	BudgetScheduleSpecs         *[]map[string]interface{}                                    `facebook:"budget_schedule_specs"`
	BuyingType                  *string                                                      `facebook:"buying_type"`
	DailyBudget                 *uint64                                                      `facebook:"daily_budget"`
	ExecutionOptions            *[]enums.AdaccountcampaignsExecutionOptionsEnumParam         `facebook:"execution_options"`
	FrequencyControlSpecs       *[]map[string]interface{}                                    `facebook:"frequency_control_specs"`
	IsAdsetBudgetSharingEnabled *bool                                                        `facebook:"is_adset_budget_sharing_enabled"`
	IsBudgetScheduleEnabled     *bool                                                        `facebook:"is_budget_schedule_enabled"`
	IsDirectSendCampaign        *bool                                                        `facebook:"is_direct_send_campaign"`
	IsMessageCampaign           *bool                                                        `facebook:"is_message_campaign"`
	IsMetaMomentMakerEnabled    *bool                                                        `facebook:"is_meta_moment_maker_enabled"`
	IsReelsTrendingAdsEnabled   *bool                                                        `facebook:"is_reels_trending_ads_enabled"`
	IsSkadnetworkAttribution    *bool                                                        `facebook:"is_skadnetwork_attribution"`
	IterativeSplitTestConfigs   *[]map[string]interface{}                                    `facebook:"iterative_split_test_configs"`
	LifetimeBudget              *uint64                                                      `facebook:"lifetime_budget"`
	Name                        *string                                                      `facebook:"name"`
	Objective                   *enums.AdaccountcampaignsObjectiveEnumParam                  `facebook:"objective"`
	PacingType                  *[]string                                                    `facebook:"pacing_type"`
	PromotedObject              *map[string]interface{}                                      `facebook:"promoted_object"`
	SmartPromotionType          *enums.AdaccountcampaignsSmartPromotionTypeEnumParam         `facebook:"smart_promotion_type"`
	SourceCampaignID            *core.ID                                                     `facebook:"source_campaign_id"`
	SpecialAdCategories         []enums.AdaccountcampaignsSpecialAdCategoriesEnumParam       `facebook:"special_ad_categories"`
	SpecialAdCategoryCountry    *[]enums.AdaccountcampaignsSpecialAdCategoryCountryEnumParam `facebook:"special_ad_category_country"`
	SpendCap                    *uint64                                                      `facebook:"spend_cap"`
	StartTime                   *time.Time                                                   `facebook:"start_time"`
	Status                      *enums.AdaccountcampaignsStatusEnumParam                     `facebook:"status"`
	StopTime                    *time.Time                                                   `facebook:"stop_time"`
	ToplineID                   *core.ID                                                     `facebook:"topline_id"`
	Extra                       core.Params                                                  `facebook:"-"`
}

func (params CreateAdAccountCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.BidStrategy != nil {
		out["bid_strategy"] = *params.BidStrategy
	}
	if params.BudgetScheduleSpecs != nil {
		out["budget_schedule_specs"] = *params.BudgetScheduleSpecs
	}
	if params.BuyingType != nil {
		out["buying_type"] = *params.BuyingType
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.FrequencyControlSpecs != nil {
		out["frequency_control_specs"] = *params.FrequencyControlSpecs
	}
	if params.IsAdsetBudgetSharingEnabled != nil {
		out["is_adset_budget_sharing_enabled"] = *params.IsAdsetBudgetSharingEnabled
	}
	if params.IsBudgetScheduleEnabled != nil {
		out["is_budget_schedule_enabled"] = *params.IsBudgetScheduleEnabled
	}
	if params.IsDirectSendCampaign != nil {
		out["is_direct_send_campaign"] = *params.IsDirectSendCampaign
	}
	if params.IsMessageCampaign != nil {
		out["is_message_campaign"] = *params.IsMessageCampaign
	}
	if params.IsMetaMomentMakerEnabled != nil {
		out["is_meta_moment_maker_enabled"] = *params.IsMetaMomentMakerEnabled
	}
	if params.IsReelsTrendingAdsEnabled != nil {
		out["is_reels_trending_ads_enabled"] = *params.IsReelsTrendingAdsEnabled
	}
	if params.IsSkadnetworkAttribution != nil {
		out["is_skadnetwork_attribution"] = *params.IsSkadnetworkAttribution
	}
	if params.IterativeSplitTestConfigs != nil {
		out["iterative_split_test_configs"] = *params.IterativeSplitTestConfigs
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Objective != nil {
		out["objective"] = *params.Objective
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.SmartPromotionType != nil {
		out["smart_promotion_type"] = *params.SmartPromotionType
	}
	if params.SourceCampaignID != nil {
		out["source_campaign_id"] = *params.SourceCampaignID
	}
	out["special_ad_categories"] = params.SpecialAdCategories
	if params.SpecialAdCategoryCountry != nil {
		out["special_ad_category_country"] = *params.SpecialAdCategoryCountry
	}
	if params.SpendCap != nil {
		out["spend_cap"] = *params.SpendCap
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StopTime != nil {
		out["stop_time"] = *params.StopTime
	}
	if params.ToplineID != nil {
		out["topline_id"] = *params.ToplineID
	}
	return out
}

func CreateAdAccountCampaignsBatchCall(id string, params CreateAdAccountCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewCreateAdAccountCampaignsBatchRequest(id string, params CreateAdAccountCampaignsParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](CreateAdAccountCampaignsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountCampaignsBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountCampaigns(ctx context.Context, client *core.Client, id string, params CreateAdAccountCampaignsParams) (*objects.Campaign, error) {
	var out objects.Campaign
	call := CreateAdAccountCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCampaignsbylabelsParams struct {
	AdLabelIds []core.ID                                          `facebook:"ad_label_ids"`
	Operator   *enums.AdaccountcampaignsbylabelsOperatorEnumParam `facebook:"operator"`
	Extra      core.Params                                        `facebook:"-"`
}

func (params GetAdAccountCampaignsbylabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_label_ids"] = params.AdLabelIds
	if params.Operator != nil {
		out["operator"] = *params.Operator
	}
	return out
}

func GetAdAccountCampaignsbylabelsBatchCall(id string, params GetAdAccountCampaignsbylabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "campaignsbylabels"), params.ToParams(), options...)
}

func NewGetAdAccountCampaignsbylabelsBatchRequest(id string, params GetAdAccountCampaignsbylabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Campaign]] {
	return core.NewBatchRequest[core.Cursor[objects.Campaign]](GetAdAccountCampaignsbylabelsBatchCall(id, params, options...))
}

func DecodeGetAdAccountCampaignsbylabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Campaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Campaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCampaignsbylabels(ctx context.Context, client *core.Client, id string, params GetAdAccountCampaignsbylabelsParams) (*core.Cursor[objects.Campaign], error) {
	var out core.Cursor[objects.Campaign]
	call := GetAdAccountCampaignsbylabelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountConnectedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountConnectedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountConnectedInstagramAccountsBatchCall(id string, params GetAdAccountConnectedInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "connected_instagram_accounts"), params.ToParams(), options...)
}

func NewGetAdAccountConnectedInstagramAccountsBatchRequest(id string, params GetAdAccountConnectedInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetAdAccountConnectedInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetAdAccountConnectedInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountConnectedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetAdAccountConnectedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetAdAccountConnectedInstagramAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountConnectedInstagramAccountsWithIabpParams struct {
	BusinessID *core.ID    `facebook:"business_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetAdAccountConnectedInstagramAccountsWithIabpParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	return out
}

func GetAdAccountConnectedInstagramAccountsWithIabpBatchCall(id string, params GetAdAccountConnectedInstagramAccountsWithIabpParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "connected_instagram_accounts_with_iabp"), params.ToParams(), options...)
}

func NewGetAdAccountConnectedInstagramAccountsWithIabpBatchRequest(id string, params GetAdAccountConnectedInstagramAccountsWithIabpParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetAdAccountConnectedInstagramAccountsWithIabpBatchCall(id, params, options...))
}

func DecodeGetAdAccountConnectedInstagramAccountsWithIabpBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountConnectedInstagramAccountsWithIabp(ctx context.Context, client *core.Client, id string, params GetAdAccountConnectedInstagramAccountsWithIabpParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetAdAccountConnectedInstagramAccountsWithIabpBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountConversionGoalsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountConversionGoalsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountConversionGoalsBatchCall(id string, params GetAdAccountConversionGoalsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversion_goals"), params.ToParams(), options...)
}

func NewGetAdAccountConversionGoalsBatchRequest(id string, params GetAdAccountConversionGoalsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsConversionGoal]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsConversionGoal]](GetAdAccountConversionGoalsBatchCall(id, params, options...))
}

func DecodeGetAdAccountConversionGoalsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsConversionGoal], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsConversionGoal]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountConversionGoals(ctx context.Context, client *core.Client, id string, params GetAdAccountConversionGoalsParams) (*core.Cursor[objects.AdsConversionGoal], error) {
	var out core.Cursor[objects.AdsConversionGoal]
	call := GetAdAccountConversionGoalsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCustomaudiencesParams struct {
	BusinessID           *core.ID                  `facebook:"business_id"`
	FetchPrimaryAudience *bool                     `facebook:"fetch_primary_audience"`
	Fields               *[]string                 `facebook:"fields"`
	Filtering            *[]map[string]interface{} `facebook:"filtering"`
	PixelID              *core.ID                  `facebook:"pixel_id"`
	Extra                core.Params               `facebook:"-"`
}

func (params GetAdAccountCustomaudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.FetchPrimaryAudience != nil {
		out["fetch_primary_audience"] = *params.FetchPrimaryAudience
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.PixelID != nil {
		out["pixel_id"] = *params.PixelID
	}
	return out
}

func GetAdAccountCustomaudiencesBatchCall(id string, params GetAdAccountCustomaudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "customaudiences"), params.ToParams(), options...)
}

func NewGetAdAccountCustomaudiencesBatchRequest(id string, params GetAdAccountCustomaudiencesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudience]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudience]](GetAdAccountCustomaudiencesBatchCall(id, params, options...))
}

func DecodeGetAdAccountCustomaudiencesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudience], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudience]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCustomaudiences(ctx context.Context, client *core.Client, id string, params GetAdAccountCustomaudiencesParams) (*core.Cursor[objects.CustomAudience], error) {
	var out core.Cursor[objects.CustomAudience]
	call := GetAdAccountCustomaudiencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountCustomaudiencesParams struct {
	AllowedDomains                *[]string                                                  `facebook:"allowed_domains"`
	AssociatedAudienceID          *core.ID                                                   `facebook:"associated_audience_id"`
	AudienceLabels                *[]enums.AdaccountcustomaudiencesAudienceLabelsEnumParam   `facebook:"audience_labels"`
	ClaimObjective                *enums.AdaccountcustomaudiencesClaimObjectiveEnumParam     `facebook:"claim_objective"`
	ContentType                   *enums.AdaccountcustomaudiencesContentTypeEnumParam        `facebook:"content_type"`
	Countries                     *string                                                    `facebook:"countries"`
	CreationParams                *map[string]interface{}                                    `facebook:"creation_params"`
	CustomerFileSource            *enums.AdaccountcustomaudiencesCustomerFileSourceEnumParam `facebook:"customer_file_source"`
	DatasetID                     *core.ID                                                   `facebook:"dataset_id"`
	Description                   *string                                                    `facebook:"description"`
	EnableFetchOrCreate           *bool                                                      `facebook:"enable_fetch_or_create"`
	EventSourceGroup              *string                                                    `facebook:"event_source_group"`
	EventSources                  *[]map[string]interface{}                                  `facebook:"event_sources"`
	Exclusions                    *[]map[string]interface{}                                  `facebook:"exclusions"`
	FacebookPageID                *core.ID                                                   `facebook:"facebook_page_id"`
	Inclusionoperator             *string                                                    `facebook:"inclusionOperator"`
	Inclusions                    *[]map[string]interface{}                                  `facebook:"inclusions"`
	IsSnapshot                    *bool                                                      `facebook:"is_snapshot"`
	IsValueBased                  *bool                                                      `facebook:"is_value_based"`
	ListOfAccounts                *[]uint64                                                  `facebook:"list_of_accounts"`
	LookalikeSpec                 *string                                                    `facebook:"lookalike_spec"`
	MarketingMessageChannels      *map[string]interface{}                                    `facebook:"marketing_message_channels"`
	Name                          *string                                                    `facebook:"name"`
	OptOutLink                    *string                                                    `facebook:"opt_out_link"`
	OriginAudienceID              *core.ID                                                   `facebook:"origin_audience_id"`
	ParentAudienceID              *core.ID                                                   `facebook:"parent_audience_id"`
	PartnerReferenceKey           *string                                                    `facebook:"partner_reference_key"`
	PixelID                       *core.ID                                                   `facebook:"pixel_id"`
	Prefill                       *bool                                                      `facebook:"prefill"`
	ProductSetID                  *core.ID                                                   `facebook:"product_set_id"`
	RegulatedAudienceSpec         *string                                                    `facebook:"regulated_audience_spec"`
	RetentionDays                 *uint64                                                    `facebook:"retention_days"`
	RevSharePolicyID              *core.ID                                                   `facebook:"rev_share_policy_id"`
	Rule                          *string                                                    `facebook:"rule"`
	RuleAggregation               *string                                                    `facebook:"rule_aggregation"`
	SubscriptionInfo              *[]enums.AdaccountcustomaudiencesSubscriptionInfoEnumParam `facebook:"subscription_info"`
	Subtype                       *enums.AdaccountcustomaudiencesSubtypeEnumParam            `facebook:"subtype"`
	UsageRestriction              *enums.AdaccountcustomaudiencesUsageRestrictionEnumParam   `facebook:"usage_restriction"`
	UseForProducts                *[]enums.AdaccountcustomaudiencesUseForProductsEnumParam   `facebook:"use_for_products"`
	UseInCampaigns                *bool                                                      `facebook:"use_in_campaigns"`
	VideoGroupIds                 *[]core.ID                                                 `facebook:"video_group_ids"`
	WhatsAppBusinessPhoneNumberID *core.ID                                                   `facebook:"whats_app_business_phone_number_id"`
	Extra                         core.Params                                                `facebook:"-"`
}

func (params CreateAdAccountCustomaudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowedDomains != nil {
		out["allowed_domains"] = *params.AllowedDomains
	}
	if params.AssociatedAudienceID != nil {
		out["associated_audience_id"] = *params.AssociatedAudienceID
	}
	if params.AudienceLabels != nil {
		out["audience_labels"] = *params.AudienceLabels
	}
	if params.ClaimObjective != nil {
		out["claim_objective"] = *params.ClaimObjective
	}
	if params.ContentType != nil {
		out["content_type"] = *params.ContentType
	}
	if params.Countries != nil {
		out["countries"] = *params.Countries
	}
	if params.CreationParams != nil {
		out["creation_params"] = *params.CreationParams
	}
	if params.CustomerFileSource != nil {
		out["customer_file_source"] = *params.CustomerFileSource
	}
	if params.DatasetID != nil {
		out["dataset_id"] = *params.DatasetID
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableFetchOrCreate != nil {
		out["enable_fetch_or_create"] = *params.EnableFetchOrCreate
	}
	if params.EventSourceGroup != nil {
		out["event_source_group"] = *params.EventSourceGroup
	}
	if params.EventSources != nil {
		out["event_sources"] = *params.EventSources
	}
	if params.Exclusions != nil {
		out["exclusions"] = *params.Exclusions
	}
	if params.FacebookPageID != nil {
		out["facebook_page_id"] = *params.FacebookPageID
	}
	if params.Inclusionoperator != nil {
		out["inclusionOperator"] = *params.Inclusionoperator
	}
	if params.Inclusions != nil {
		out["inclusions"] = *params.Inclusions
	}
	if params.IsSnapshot != nil {
		out["is_snapshot"] = *params.IsSnapshot
	}
	if params.IsValueBased != nil {
		out["is_value_based"] = *params.IsValueBased
	}
	if params.ListOfAccounts != nil {
		out["list_of_accounts"] = *params.ListOfAccounts
	}
	if params.LookalikeSpec != nil {
		out["lookalike_spec"] = *params.LookalikeSpec
	}
	if params.MarketingMessageChannels != nil {
		out["marketing_message_channels"] = *params.MarketingMessageChannels
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OptOutLink != nil {
		out["opt_out_link"] = *params.OptOutLink
	}
	if params.OriginAudienceID != nil {
		out["origin_audience_id"] = *params.OriginAudienceID
	}
	if params.ParentAudienceID != nil {
		out["parent_audience_id"] = *params.ParentAudienceID
	}
	if params.PartnerReferenceKey != nil {
		out["partner_reference_key"] = *params.PartnerReferenceKey
	}
	if params.PixelID != nil {
		out["pixel_id"] = *params.PixelID
	}
	if params.Prefill != nil {
		out["prefill"] = *params.Prefill
	}
	if params.ProductSetID != nil {
		out["product_set_id"] = *params.ProductSetID
	}
	if params.RegulatedAudienceSpec != nil {
		out["regulated_audience_spec"] = *params.RegulatedAudienceSpec
	}
	if params.RetentionDays != nil {
		out["retention_days"] = *params.RetentionDays
	}
	if params.RevSharePolicyID != nil {
		out["rev_share_policy_id"] = *params.RevSharePolicyID
	}
	if params.Rule != nil {
		out["rule"] = *params.Rule
	}
	if params.RuleAggregation != nil {
		out["rule_aggregation"] = *params.RuleAggregation
	}
	if params.SubscriptionInfo != nil {
		out["subscription_info"] = *params.SubscriptionInfo
	}
	if params.Subtype != nil {
		out["subtype"] = *params.Subtype
	}
	if params.UsageRestriction != nil {
		out["usage_restriction"] = *params.UsageRestriction
	}
	if params.UseForProducts != nil {
		out["use_for_products"] = *params.UseForProducts
	}
	if params.UseInCampaigns != nil {
		out["use_in_campaigns"] = *params.UseInCampaigns
	}
	if params.VideoGroupIds != nil {
		out["video_group_ids"] = *params.VideoGroupIds
	}
	if params.WhatsAppBusinessPhoneNumberID != nil {
		out["whats_app_business_phone_number_id"] = *params.WhatsAppBusinessPhoneNumberID
	}
	return out
}

func CreateAdAccountCustomaudiencesBatchCall(id string, params CreateAdAccountCustomaudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "customaudiences"), params.ToParams(), options...)
}

func NewCreateAdAccountCustomaudiencesBatchRequest(id string, params CreateAdAccountCustomaudiencesParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateAdAccountCustomaudiencesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountCustomaudiencesBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountCustomaudiences(ctx context.Context, client *core.Client, id string, params CreateAdAccountCustomaudiencesParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateAdAccountCustomaudiencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCustomaudiencestosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCustomaudiencestosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCustomaudiencestosBatchCall(id string, params GetAdAccountCustomaudiencestosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "customaudiencestos"), params.ToParams(), options...)
}

func NewGetAdAccountCustomaudiencestosBatchRequest(id string, params GetAdAccountCustomaudiencestosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomAudiencesTOS]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomAudiencesTOS]](GetAdAccountCustomaudiencestosBatchCall(id, params, options...))
}

func DecodeGetAdAccountCustomaudiencestosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomAudiencesTOS], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomAudiencesTOS]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCustomaudiencestos(ctx context.Context, client *core.Client, id string, params GetAdAccountCustomaudiencestosParams) (*core.Cursor[objects.CustomAudiencesTOS], error) {
	var out core.Cursor[objects.CustomAudiencesTOS]
	call := GetAdAccountCustomaudiencestosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountCustomaudiencestosParams struct {
	BusinessID *core.ID    `facebook:"business_id"`
	TosID      core.ID     `facebook:"tos_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateAdAccountCustomaudiencestosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	out["tos_id"] = params.TosID
	return out
}

func CreateAdAccountCustomaudiencestosBatchCall(id string, params CreateAdAccountCustomaudiencestosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "customaudiencestos"), params.ToParams(), options...)
}

func NewCreateAdAccountCustomaudiencestosBatchRequest(id string, params CreateAdAccountCustomaudiencestosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountCustomaudiencestosBatchCall(id, params, options...))
}

func DecodeCreateAdAccountCustomaudiencestosBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountCustomaudiencestos(ctx context.Context, client *core.Client, id string, params CreateAdAccountCustomaudiencestosParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountCustomaudiencestosBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCustomconversionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCustomconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCustomconversionsBatchCall(id string, params GetAdAccountCustomconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "customconversions"), params.ToParams(), options...)
}

func NewGetAdAccountCustomconversionsBatchRequest(id string, params GetAdAccountCustomconversionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CustomConversion]] {
	return core.NewBatchRequest[core.Cursor[objects.CustomConversion]](GetAdAccountCustomconversionsBatchCall(id, params, options...))
}

func DecodeGetAdAccountCustomconversionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CustomConversion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CustomConversion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCustomconversions(ctx context.Context, client *core.Client, id string, params GetAdAccountCustomconversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	var out core.Cursor[objects.CustomConversion]
	call := GetAdAccountCustomconversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountCustomconversionsParams struct {
	ActionSourceType       *enums.AdaccountcustomconversionsActionSourceTypeEnumParam `facebook:"action_source_type"`
	AdvancedRule           *string                                                    `facebook:"advanced_rule"`
	CustomEventType        *enums.AdaccountcustomconversionsCustomEventTypeEnumParam  `facebook:"custom_event_type"`
	DefaultConversionValue *float64                                                   `facebook:"default_conversion_value"`
	Description            *string                                                    `facebook:"description"`
	EventSourceID          *core.ID                                                   `facebook:"event_source_id"`
	Name                   string                                                     `facebook:"name"`
	Rule                   *string                                                    `facebook:"rule"`
	Extra                  core.Params                                                `facebook:"-"`
}

func (params CreateAdAccountCustomconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionSourceType != nil {
		out["action_source_type"] = *params.ActionSourceType
	}
	if params.AdvancedRule != nil {
		out["advanced_rule"] = *params.AdvancedRule
	}
	if params.CustomEventType != nil {
		out["custom_event_type"] = *params.CustomEventType
	}
	if params.DefaultConversionValue != nil {
		out["default_conversion_value"] = *params.DefaultConversionValue
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EventSourceID != nil {
		out["event_source_id"] = *params.EventSourceID
	}
	out["name"] = params.Name
	if params.Rule != nil {
		out["rule"] = *params.Rule
	}
	return out
}

func CreateAdAccountCustomconversionsBatchCall(id string, params CreateAdAccountCustomconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "customconversions"), params.ToParams(), options...)
}

func NewCreateAdAccountCustomconversionsBatchRequest(id string, params CreateAdAccountCustomconversionsParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomConversion] {
	return core.NewBatchRequest[objects.CustomConversion](CreateAdAccountCustomconversionsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountCustomconversionsBatchResponse(response *core.BatchResponse) (*objects.CustomConversion, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomConversion
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountCustomconversions(ctx context.Context, client *core.Client, id string, params CreateAdAccountCustomconversionsParams) (*objects.CustomConversion, error) {
	var out objects.CustomConversion
	call := CreateAdAccountCustomconversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountDeliveryEstimateParams struct {
	OptimizationGoal enums.AdaccountdeliveryEstimateOptimizationGoalEnumParam `facebook:"optimization_goal"`
	PromotedObject   *map[string]interface{}                                  `facebook:"promoted_object"`
	TargetingSpec    objects.Targeting                                        `facebook:"targeting_spec"`
	Extra            core.Params                                              `facebook:"-"`
}

func (params GetAdAccountDeliveryEstimateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["optimization_goal"] = params.OptimizationGoal
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	out["targeting_spec"] = params.TargetingSpec
	return out
}

func GetAdAccountDeliveryEstimateBatchCall(id string, params GetAdAccountDeliveryEstimateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "delivery_estimate"), params.ToParams(), options...)
}

func NewGetAdAccountDeliveryEstimateBatchRequest(id string, params GetAdAccountDeliveryEstimateParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountDeliveryEstimate]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountDeliveryEstimate]](GetAdAccountDeliveryEstimateBatchCall(id, params, options...))
}

func DecodeGetAdAccountDeliveryEstimateBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountDeliveryEstimate], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountDeliveryEstimate]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountDeliveryEstimate(ctx context.Context, client *core.Client, id string, params GetAdAccountDeliveryEstimateParams) (*core.Cursor[objects.AdAccountDeliveryEstimate], error) {
	var out core.Cursor[objects.AdAccountDeliveryEstimate]
	call := GetAdAccountDeliveryEstimateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountDeprecatedtargetingadsetsParams struct {
	Type  *string     `facebook:"type"`
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountDeprecatedtargetingadsetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetAdAccountDeprecatedtargetingadsetsBatchCall(id string, params GetAdAccountDeprecatedtargetingadsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "deprecatedtargetingadsets"), params.ToParams(), options...)
}

func NewGetAdAccountDeprecatedtargetingadsetsBatchRequest(id string, params GetAdAccountDeprecatedtargetingadsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetAdAccountDeprecatedtargetingadsetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountDeprecatedtargetingadsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountDeprecatedtargetingadsets(ctx context.Context, client *core.Client, id string, params GetAdAccountDeprecatedtargetingadsetsParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	call := GetAdAccountDeprecatedtargetingadsetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountDsaRecommendationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountDsaRecommendationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountDsaRecommendationsBatchCall(id string, params GetAdAccountDsaRecommendationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dsa_recommendations"), params.ToParams(), options...)
}

func NewGetAdAccountDsaRecommendationsBatchRequest(id string, params GetAdAccountDsaRecommendationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountDsaRecommendations]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountDsaRecommendations]](GetAdAccountDsaRecommendationsBatchCall(id, params, options...))
}

func DecodeGetAdAccountDsaRecommendationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountDsaRecommendations], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountDsaRecommendations]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountDsaRecommendations(ctx context.Context, client *core.Client, id string, params GetAdAccountDsaRecommendationsParams) (*core.Cursor[objects.AdAccountDsaRecommendations], error) {
	var out core.Cursor[objects.AdAccountDsaRecommendations]
	call := GetAdAccountDsaRecommendationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountGeneratepreviewsParams struct {
	AdFormat             enums.AdaccountgeneratepreviewsAdFormatEnumParam         `facebook:"ad_format"`
	Creative             objects.AdCreative                                       `facebook:"creative"`
	CreativeFeature      *enums.AdaccountgeneratepreviewsCreativeFeatureEnumParam `facebook:"creative_feature"`
	DynamicAssetLabel    *string                                                  `facebook:"dynamic_asset_label"`
	DynamicCreativeSpec  *map[string]interface{}                                  `facebook:"dynamic_creative_spec"`
	DynamicCustomization *map[string]interface{}                                  `facebook:"dynamic_customization"`
	EndDate              *time.Time                                               `facebook:"end_date"`
	Height               *uint64                                                  `facebook:"height"`
	Locale               *string                                                  `facebook:"locale"`
	Message              *map[string]interface{}                                  `facebook:"message"`
	PlacePageID          *core.ID                                                 `facebook:"place_page_id"`
	Post                 *map[string]interface{}                                  `facebook:"post"`
	ProductItemIds       *[]core.ID                                               `facebook:"product_item_ids"`
	RenderType           *enums.AdaccountgeneratepreviewsRenderTypeEnumParam      `facebook:"render_type"`
	StartDate            *time.Time                                               `facebook:"start_date"`
	Width                *uint64                                                  `facebook:"width"`
	Extra                core.Params                                              `facebook:"-"`
}

func (params GetAdAccountGeneratepreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_format"] = params.AdFormat
	out["creative"] = params.Creative
	if params.CreativeFeature != nil {
		out["creative_feature"] = *params.CreativeFeature
	}
	if params.DynamicAssetLabel != nil {
		out["dynamic_asset_label"] = *params.DynamicAssetLabel
	}
	if params.DynamicCreativeSpec != nil {
		out["dynamic_creative_spec"] = *params.DynamicCreativeSpec
	}
	if params.DynamicCustomization != nil {
		out["dynamic_customization"] = *params.DynamicCustomization
	}
	if params.EndDate != nil {
		out["end_date"] = *params.EndDate
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.PlacePageID != nil {
		out["place_page_id"] = *params.PlacePageID
	}
	if params.Post != nil {
		out["post"] = *params.Post
	}
	if params.ProductItemIds != nil {
		out["product_item_ids"] = *params.ProductItemIds
	}
	if params.RenderType != nil {
		out["render_type"] = *params.RenderType
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetAdAccountGeneratepreviewsBatchCall(id string, params GetAdAccountGeneratepreviewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "generatepreviews"), params.ToParams(), options...)
}

func NewGetAdAccountGeneratepreviewsBatchRequest(id string, params GetAdAccountGeneratepreviewsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPreview]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPreview]](GetAdAccountGeneratepreviewsBatchCall(id, params, options...))
}

func DecodeGetAdAccountGeneratepreviewsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPreview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPreview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountGeneratepreviews(ctx context.Context, client *core.Client, id string, params GetAdAccountGeneratepreviewsParams) (*core.Cursor[objects.AdPreview], error) {
	var out core.Cursor[objects.AdPreview]
	call := GetAdAccountGeneratepreviewsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountImpactingAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountImpactingAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountImpactingAdStudiesBatchCall(id string, params GetAdAccountImpactingAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "impacting_ad_studies"), params.ToParams(), options...)
}

func NewGetAdAccountImpactingAdStudiesBatchRequest(id string, params GetAdAccountImpactingAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudy]](GetAdAccountImpactingAdStudiesBatchCall(id, params, options...))
}

func DecodeGetAdAccountImpactingAdStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountImpactingAdStudies(ctx context.Context, client *core.Client, id string, params GetAdAccountImpactingAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	call := GetAdAccountImpactingAdStudiesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdaccountinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdaccountinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdaccountinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdaccountinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdaccountinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                       `facebook:"default_summary"`
	ExportColumns                *[]string                                                   `facebook:"export_columns"`
	ExportFormat                 *string                                                     `facebook:"export_format"`
	ExportName                   *string                                                     `facebook:"export_name"`
	Fields                       *[]string                                                   `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                   `facebook:"filtering"`
	GraphCache                   *bool                                                       `facebook:"graph_cache"`
	Level                        *enums.AdaccountinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                        `facebook:"limit"`
	ProductIDLimit               *int                                                        `facebook:"product_id_limit"`
	Sort                         *[]string                                                   `facebook:"sort"`
	Summary                      *[]string                                                   `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdaccountinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                     `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                     `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                   `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                       `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                       `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                 `facebook:"-"`
}

func (params GetAdAccountInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionAttributionWindows != nil {
		out["action_attribution_windows"] = *params.ActionAttributionWindows
	}
	if params.ActionBreakdowns != nil {
		out["action_breakdowns"] = *params.ActionBreakdowns
	}
	if params.ActionReportTime != nil {
		out["action_report_time"] = *params.ActionReportTime
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.DefaultSummary != nil {
		out["default_summary"] = *params.DefaultSummary
	}
	if params.ExportColumns != nil {
		out["export_columns"] = *params.ExportColumns
	}
	if params.ExportFormat != nil {
		out["export_format"] = *params.ExportFormat
	}
	if params.ExportName != nil {
		out["export_name"] = *params.ExportName
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GraphCache != nil {
		out["graph_cache"] = *params.GraphCache
	}
	if params.Level != nil {
		out["level"] = *params.Level
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ProductIDLimit != nil {
		out["product_id_limit"] = *params.ProductIDLimit
	}
	if params.Sort != nil {
		out["sort"] = *params.Sort
	}
	if params.Summary != nil {
		out["summary"] = *params.Summary
	}
	if params.SummaryActionBreakdowns != nil {
		out["summary_action_breakdowns"] = *params.SummaryActionBreakdowns
	}
	if params.TimeIncrement != nil {
		out["time_increment"] = *params.TimeIncrement
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.TimeRanges != nil {
		out["time_ranges"] = *params.TimeRanges
	}
	if params.UseAccountAttributionSetting != nil {
		out["use_account_attribution_setting"] = *params.UseAccountAttributionSetting
	}
	if params.UseUnifiedAttributionSetting != nil {
		out["use_unified_attribution_setting"] = *params.UseUnifiedAttributionSetting
	}
	return out
}

func GetAdAccountInsightsBatchCall(id string, params GetAdAccountInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetAdAccountInsightsBatchRequest(id string, params GetAdAccountInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsInsights]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsInsights]](GetAdAccountInsightsBatchCall(id, params, options...))
}

func DecodeGetAdAccountInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsInsights], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsInsights]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountInsights(ctx context.Context, client *core.Client, id string, params GetAdAccountInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	var out core.Cursor[objects.AdsInsights]
	call := GetAdAccountInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdaccountinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdaccountinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdaccountinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdaccountinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdaccountinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                       `facebook:"default_summary"`
	ExportColumns                *[]string                                                   `facebook:"export_columns"`
	ExportFormat                 *string                                                     `facebook:"export_format"`
	ExportName                   *string                                                     `facebook:"export_name"`
	Fields                       *[]string                                                   `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                   `facebook:"filtering"`
	GraphCache                   *bool                                                       `facebook:"graph_cache"`
	Level                        *enums.AdaccountinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                        `facebook:"limit"`
	ProductIDLimit               *int                                                        `facebook:"product_id_limit"`
	Sort                         *[]string                                                   `facebook:"sort"`
	Summary                      *[]string                                                   `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdaccountinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                     `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                     `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                   `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                       `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                       `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                 `facebook:"-"`
}

func (params CreateAdAccountInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionAttributionWindows != nil {
		out["action_attribution_windows"] = *params.ActionAttributionWindows
	}
	if params.ActionBreakdowns != nil {
		out["action_breakdowns"] = *params.ActionBreakdowns
	}
	if params.ActionReportTime != nil {
		out["action_report_time"] = *params.ActionReportTime
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.DefaultSummary != nil {
		out["default_summary"] = *params.DefaultSummary
	}
	if params.ExportColumns != nil {
		out["export_columns"] = *params.ExportColumns
	}
	if params.ExportFormat != nil {
		out["export_format"] = *params.ExportFormat
	}
	if params.ExportName != nil {
		out["export_name"] = *params.ExportName
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GraphCache != nil {
		out["graph_cache"] = *params.GraphCache
	}
	if params.Level != nil {
		out["level"] = *params.Level
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ProductIDLimit != nil {
		out["product_id_limit"] = *params.ProductIDLimit
	}
	if params.Sort != nil {
		out["sort"] = *params.Sort
	}
	if params.Summary != nil {
		out["summary"] = *params.Summary
	}
	if params.SummaryActionBreakdowns != nil {
		out["summary_action_breakdowns"] = *params.SummaryActionBreakdowns
	}
	if params.TimeIncrement != nil {
		out["time_increment"] = *params.TimeIncrement
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.TimeRanges != nil {
		out["time_ranges"] = *params.TimeRanges
	}
	if params.UseAccountAttributionSetting != nil {
		out["use_account_attribution_setting"] = *params.UseAccountAttributionSetting
	}
	if params.UseUnifiedAttributionSetting != nil {
		out["use_unified_attribution_setting"] = *params.UseUnifiedAttributionSetting
	}
	return out
}

func CreateAdAccountInsightsBatchCall(id string, params CreateAdAccountInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewCreateAdAccountInsightsBatchRequest(id string, params CreateAdAccountInsightsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdReportRun] {
	return core.NewBatchRequest[objects.AdReportRun](CreateAdAccountInsightsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountInsightsBatchResponse(response *core.BatchResponse) (*objects.AdReportRun, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdReportRun
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountInsights(ctx context.Context, client *core.Client, id string, params CreateAdAccountInsightsParams) (*objects.AdReportRun, error) {
	var out objects.AdReportRun
	call := CreateAdAccountInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountInstagramAccountsBatchCall(id string, params GetAdAccountInstagramAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_accounts"), params.ToParams(), options...)
}

func NewGetAdAccountInstagramAccountsBatchRequest(id string, params GetAdAccountInstagramAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUser]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUser]](GetAdAccountInstagramAccountsBatchCall(id, params, options...))
}

func DecodeGetAdAccountInstagramAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetAdAccountInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	call := GetAdAccountInstagramAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountIosFourteenCampaignLimitsParams struct {
	AppID core.ID     `facebook:"app_id"`
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountIosFourteenCampaignLimitsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["app_id"] = params.AppID
	return out
}

func GetAdAccountIosFourteenCampaignLimitsBatchCall(id string, params GetAdAccountIosFourteenCampaignLimitsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ios_fourteen_campaign_limits"), params.ToParams(), options...)
}

func NewGetAdAccountIosFourteenCampaignLimitsBatchRequest(id string, params GetAdAccountIosFourteenCampaignLimitsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountIosFourteenCampaignLimits]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountIosFourteenCampaignLimits]](GetAdAccountIosFourteenCampaignLimitsBatchCall(id, params, options...))
}

func DecodeGetAdAccountIosFourteenCampaignLimitsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountIosFourteenCampaignLimits], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountIosFourteenCampaignLimits]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountIosFourteenCampaignLimits(ctx context.Context, client *core.Client, id string, params GetAdAccountIosFourteenCampaignLimitsParams) (*core.Cursor[objects.AdAccountIosFourteenCampaignLimits], error) {
	var out core.Cursor[objects.AdAccountIosFourteenCampaignLimits]
	call := GetAdAccountIosFourteenCampaignLimitsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountMatchedSearchApplicationsParams struct {
	AllowIncompleteApp     *bool                                                              `facebook:"allow_incomplete_app"`
	AppStore               enums.AdaccountmatchedSearchApplicationsAppStoreEnumParam          `facebook:"app_store"`
	AppStoreCountry        *string                                                            `facebook:"app_store_country"`
	BusinessID             *core.ID                                                           `facebook:"business_id"`
	IsSkadnetworkSearch    *bool                                                              `facebook:"is_skadnetwork_search"`
	OnlyAppsWithPermission *bool                                                              `facebook:"only_apps_with_permission"`
	QueryTerm              string                                                             `facebook:"query_term"`
	StoresToFilter         *[]enums.AdaccountmatchedSearchApplicationsStoresToFilterEnumParam `facebook:"stores_to_filter"`
	Extra                  core.Params                                                        `facebook:"-"`
}

func (params GetAdAccountMatchedSearchApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowIncompleteApp != nil {
		out["allow_incomplete_app"] = *params.AllowIncompleteApp
	}
	out["app_store"] = params.AppStore
	if params.AppStoreCountry != nil {
		out["app_store_country"] = *params.AppStoreCountry
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.IsSkadnetworkSearch != nil {
		out["is_skadnetwork_search"] = *params.IsSkadnetworkSearch
	}
	if params.OnlyAppsWithPermission != nil {
		out["only_apps_with_permission"] = *params.OnlyAppsWithPermission
	}
	out["query_term"] = params.QueryTerm
	if params.StoresToFilter != nil {
		out["stores_to_filter"] = *params.StoresToFilter
	}
	return out
}

func GetAdAccountMatchedSearchApplicationsBatchCall(id string, params GetAdAccountMatchedSearchApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "matched_search_applications"), params.ToParams(), options...)
}

func NewGetAdAccountMatchedSearchApplicationsBatchRequest(id string, params GetAdAccountMatchedSearchApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData]](GetAdAccountMatchedSearchApplicationsBatchCall(id, params, options...))
}

func DecodeGetAdAccountMatchedSearchApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountMatchedSearchApplications(ctx context.Context, client *core.Client, id string, params GetAdAccountMatchedSearchApplicationsParams) (*core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData], error) {
	var out core.Cursor[objects.AdAccountMatchedSearchApplicationsEdgeData]
	call := GetAdAccountMatchedSearchApplicationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountMaxBidParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountMaxBidParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountMaxBidBatchCall(id string, params GetAdAccountMaxBidParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "max_bid"), params.ToParams(), options...)
}

func NewGetAdAccountMaxBidBatchRequest(id string, params GetAdAccountMaxBidParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountMaxBid]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountMaxBid]](GetAdAccountMaxBidBatchCall(id, params, options...))
}

func DecodeGetAdAccountMaxBidBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountMaxBid], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountMaxBid]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountMaxBid(ctx context.Context, client *core.Client, id string, params GetAdAccountMaxBidParams) (*core.Cursor[objects.AdAccountMaxBid], error) {
	var out core.Cursor[objects.AdAccountMaxBid]
	call := GetAdAccountMaxBidBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountMcmeconversionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountMcmeconversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountMcmeconversionsBatchCall(id string, params GetAdAccountMcmeconversionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "mcmeconversions"), params.ToParams(), options...)
}

func NewGetAdAccountMcmeconversionsBatchRequest(id string, params GetAdAccountMcmeconversionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsMcmeConversion]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsMcmeConversion]](GetAdAccountMcmeconversionsBatchCall(id, params, options...))
}

func DecodeGetAdAccountMcmeconversionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsMcmeConversion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsMcmeConversion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountMcmeconversions(ctx context.Context, client *core.Client, id string, params GetAdAccountMcmeconversionsParams) (*core.Cursor[objects.AdsMcmeConversion], error) {
	var out core.Cursor[objects.AdsMcmeConversion]
	call := GetAdAccountMcmeconversionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountMessageCampaignParams struct {
	BidAmount      *uint64     `facebook:"bid_amount"`
	DailyBudget    *uint64     `facebook:"daily_budget"`
	EndTime        *uint64     `facebook:"end_time"`
	LifetimeBudget *uint64     `facebook:"lifetime_budget"`
	Name           string      `facebook:"name"`
	PageID         core.ID     `facebook:"page_id"`
	PixelID        *core.ID    `facebook:"pixel_id"`
	StartTime      *uint64     `facebook:"start_time"`
	Extra          core.Params `facebook:"-"`
}

func (params CreateAdAccountMessageCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	out["name"] = params.Name
	out["page_id"] = params.PageID
	if params.PixelID != nil {
		out["pixel_id"] = *params.PixelID
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func CreateAdAccountMessageCampaignBatchCall(id string, params CreateAdAccountMessageCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "message_campaign"), params.ToParams(), options...)
}

func NewCreateAdAccountMessageCampaignBatchRequest(id string, params CreateAdAccountMessageCampaignParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdAccountMessageCampaignBatchCall(id, params, options...))
}

func DecodeCreateAdAccountMessageCampaignBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdAccountMessageCampaign(ctx context.Context, client *core.Client, id string, params CreateAdAccountMessageCampaignParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateAdAccountMessageCampaignBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountMessageDeliveryEstimateParams struct {
	BidAmount            *uint64                                                          `facebook:"bid_amount"`
	DailyBudget          *uint64                                                          `facebook:"daily_budget"`
	IsDirectSendCampaign *bool                                                            `facebook:"is_direct_send_campaign"`
	LifetimeBudget       *uint64                                                          `facebook:"lifetime_budget"`
	LifetimeInDays       *uint64                                                          `facebook:"lifetime_in_days"`
	OptimizationGoal     *enums.AdaccountmessageDeliveryEstimateOptimizationGoalEnumParam `facebook:"optimization_goal"`
	PacingType           *enums.AdaccountmessageDeliveryEstimatePacingTypeEnumParam       `facebook:"pacing_type"`
	PromotedObject       map[string]interface{}                                           `facebook:"promoted_object"`
	TargetingSpec        objects.Targeting                                                `facebook:"targeting_spec"`
	Extra                core.Params                                                      `facebook:"-"`
}

func (params GetAdAccountMessageDeliveryEstimateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.IsDirectSendCampaign != nil {
		out["is_direct_send_campaign"] = *params.IsDirectSendCampaign
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.LifetimeInDays != nil {
		out["lifetime_in_days"] = *params.LifetimeInDays
	}
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	out["promoted_object"] = params.PromotedObject
	out["targeting_spec"] = params.TargetingSpec
	return out
}

func GetAdAccountMessageDeliveryEstimateBatchCall(id string, params GetAdAccountMessageDeliveryEstimateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "message_delivery_estimate"), params.ToParams(), options...)
}

func NewGetAdAccountMessageDeliveryEstimateBatchRequest(id string, params GetAdAccountMessageDeliveryEstimateParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessageDeliveryEstimate]] {
	return core.NewBatchRequest[core.Cursor[objects.MessageDeliveryEstimate]](GetAdAccountMessageDeliveryEstimateBatchCall(id, params, options...))
}

func DecodeGetAdAccountMessageDeliveryEstimateBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessageDeliveryEstimate], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessageDeliveryEstimate]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountMessageDeliveryEstimate(ctx context.Context, client *core.Client, id string, params GetAdAccountMessageDeliveryEstimateParams) (*core.Cursor[objects.MessageDeliveryEstimate], error) {
	var out core.Cursor[objects.MessageDeliveryEstimate]
	call := GetAdAccountMessageDeliveryEstimateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountMessagesParams struct {
	CustomAudienceID          *core.ID                `facebook:"custom_audience_id"`
	Message                   *map[string]interface{} `facebook:"message"`
	MessageID                 core.ID                 `facebook:"message_id"`
	MessengerDeliveryData     *map[string]interface{} `facebook:"messenger_delivery_data"`
	MinConversationGapSeconds *uint64                 `facebook:"min_conversation_gap_seconds"`
	Extra                     core.Params             `facebook:"-"`
}

func (params CreateAdAccountMessagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CustomAudienceID != nil {
		out["custom_audience_id"] = *params.CustomAudienceID
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	out["message_id"] = params.MessageID
	if params.MessengerDeliveryData != nil {
		out["messenger_delivery_data"] = *params.MessengerDeliveryData
	}
	if params.MinConversationGapSeconds != nil {
		out["min_conversation_gap_seconds"] = *params.MinConversationGapSeconds
	}
	return out
}

func CreateAdAccountMessagesBatchCall(id string, params CreateAdAccountMessagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messages"), params.ToParams(), options...)
}

func NewCreateAdAccountMessagesBatchRequest(id string, params CreateAdAccountMessagesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdAccountMessagesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountMessagesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdAccountMessages(ctx context.Context, client *core.Client, id string, params CreateAdAccountMessagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateAdAccountMessagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountMinimumBudgetsParams struct {
	BidAmount *int        `facebook:"bid_amount"`
	Extra     core.Params `facebook:"-"`
}

func (params GetAdAccountMinimumBudgetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	return out
}

func GetAdAccountMinimumBudgetsBatchCall(id string, params GetAdAccountMinimumBudgetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "minimum_budgets"), params.ToParams(), options...)
}

func NewGetAdAccountMinimumBudgetsBatchRequest(id string, params GetAdAccountMinimumBudgetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MinimumBudget]] {
	return core.NewBatchRequest[core.Cursor[objects.MinimumBudget]](GetAdAccountMinimumBudgetsBatchCall(id, params, options...))
}

func DecodeGetAdAccountMinimumBudgetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MinimumBudget], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MinimumBudget]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountMinimumBudgets(ctx context.Context, client *core.Client, id string, params GetAdAccountMinimumBudgetsParams) (*core.Cursor[objects.MinimumBudget], error) {
	var out core.Cursor[objects.MinimumBudget]
	call := GetAdAccountMinimumBudgetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountOnbehalfRequestsParams struct {
	Status *enums.AdaccountonbehalfRequestsStatusEnumParam `facebook:"status"`
	Extra  core.Params                                     `facebook:"-"`
}

func (params GetAdAccountOnbehalfRequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetAdAccountOnbehalfRequestsBatchCall(id string, params GetAdAccountOnbehalfRequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "onbehalf_requests"), params.ToParams(), options...)
}

func NewGetAdAccountOnbehalfRequestsBatchRequest(id string, params GetAdAccountOnbehalfRequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest]](GetAdAccountOnbehalfRequestsBatchCall(id, params, options...))
}

func DecodeGetAdAccountOnbehalfRequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountOnbehalfRequests(ctx context.Context, client *core.Client, id string, params GetAdAccountOnbehalfRequestsParams) (*core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest], error) {
	var out core.Cursor[objects.BusinessOwnedObjectOnBehalfOfRequest]
	call := GetAdAccountOnbehalfRequestsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountProductAudiencesParams struct {
	AllowedDomains       *[]string                                               `facebook:"allowed_domains"`
	AssociatedAudienceID *core.ID                                                `facebook:"associated_audience_id"`
	ClaimObjective       *enums.AdaccountproductAudiencesClaimObjectiveEnumParam `facebook:"claim_objective"`
	ContentType          *enums.AdaccountproductAudiencesContentTypeEnumParam    `facebook:"content_type"`
	CreationParams       *map[string]interface{}                                 `facebook:"creation_params"`
	Description          *string                                                 `facebook:"description"`
	EnableFetchOrCreate  *bool                                                   `facebook:"enable_fetch_or_create"`
	EventSourceGroup     *string                                                 `facebook:"event_source_group"`
	EventSources         *[]map[string]interface{}                               `facebook:"event_sources"`
	Exclusions           *[]map[string]interface{}                               `facebook:"exclusions"`
	Inclusionoperator    *string                                                 `facebook:"inclusionOperator"`
	Inclusions           *[]map[string]interface{}                               `facebook:"inclusions"`
	IsSnapshot           *bool                                                   `facebook:"is_snapshot"`
	IsValueBased         *bool                                                   `facebook:"is_value_based"`
	Name                 string                                                  `facebook:"name"`
	OptOutLink           *string                                                 `facebook:"opt_out_link"`
	ParentAudienceID     *core.ID                                                `facebook:"parent_audience_id"`
	ProductSetID         core.ID                                                 `facebook:"product_set_id"`
	RevSharePolicyID     *core.ID                                                `facebook:"rev_share_policy_id"`
	Subtype              *enums.AdaccountproductAudiencesSubtypeEnumParam        `facebook:"subtype"`
	Extra                core.Params                                             `facebook:"-"`
}

func (params CreateAdAccountProductAudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowedDomains != nil {
		out["allowed_domains"] = *params.AllowedDomains
	}
	if params.AssociatedAudienceID != nil {
		out["associated_audience_id"] = *params.AssociatedAudienceID
	}
	if params.ClaimObjective != nil {
		out["claim_objective"] = *params.ClaimObjective
	}
	if params.ContentType != nil {
		out["content_type"] = *params.ContentType
	}
	if params.CreationParams != nil {
		out["creation_params"] = *params.CreationParams
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableFetchOrCreate != nil {
		out["enable_fetch_or_create"] = *params.EnableFetchOrCreate
	}
	if params.EventSourceGroup != nil {
		out["event_source_group"] = *params.EventSourceGroup
	}
	if params.EventSources != nil {
		out["event_sources"] = *params.EventSources
	}
	if params.Exclusions != nil {
		out["exclusions"] = *params.Exclusions
	}
	if params.Inclusionoperator != nil {
		out["inclusionOperator"] = *params.Inclusionoperator
	}
	if params.Inclusions != nil {
		out["inclusions"] = *params.Inclusions
	}
	if params.IsSnapshot != nil {
		out["is_snapshot"] = *params.IsSnapshot
	}
	if params.IsValueBased != nil {
		out["is_value_based"] = *params.IsValueBased
	}
	out["name"] = params.Name
	if params.OptOutLink != nil {
		out["opt_out_link"] = *params.OptOutLink
	}
	if params.ParentAudienceID != nil {
		out["parent_audience_id"] = *params.ParentAudienceID
	}
	out["product_set_id"] = params.ProductSetID
	if params.RevSharePolicyID != nil {
		out["rev_share_policy_id"] = *params.RevSharePolicyID
	}
	if params.Subtype != nil {
		out["subtype"] = *params.Subtype
	}
	return out
}

func CreateAdAccountProductAudiencesBatchCall(id string, params CreateAdAccountProductAudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "product_audiences"), params.ToParams(), options...)
}

func NewCreateAdAccountProductAudiencesBatchRequest(id string, params CreateAdAccountProductAudiencesParams, options ...core.BatchOption) *core.BatchRequest[objects.CustomAudience] {
	return core.NewBatchRequest[objects.CustomAudience](CreateAdAccountProductAudiencesBatchCall(id, params, options...))
}

func DecodeCreateAdAccountProductAudiencesBatchResponse(response *core.BatchResponse) (*objects.CustomAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CustomAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountProductAudiences(ctx context.Context, client *core.Client, id string, params CreateAdAccountProductAudiencesParams) (*objects.CustomAudience, error) {
	var out objects.CustomAudience
	call := CreateAdAccountProductAudiencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountPromotePagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountPromotePagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountPromotePagesBatchCall(id string, params GetAdAccountPromotePagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "promote_pages"), params.ToParams(), options...)
}

func NewGetAdAccountPromotePagesBatchRequest(id string, params GetAdAccountPromotePagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetAdAccountPromotePagesBatchCall(id, params, options...))
}

func DecodeGetAdAccountPromotePagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetAdAccountPromotePages(ctx context.Context, client *core.Client, id string, params GetAdAccountPromotePagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetAdAccountPromotePagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountPublisherBlockListsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountPublisherBlockListsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountPublisherBlockListsBatchCall(id string, params GetAdAccountPublisherBlockListsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "publisher_block_lists"), params.ToParams(), options...)
}

func NewGetAdAccountPublisherBlockListsBatchRequest(id string, params GetAdAccountPublisherBlockListsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PublisherBlockList]] {
	return core.NewBatchRequest[core.Cursor[objects.PublisherBlockList]](GetAdAccountPublisherBlockListsBatchCall(id, params, options...))
}

func DecodeGetAdAccountPublisherBlockListsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PublisherBlockList], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PublisherBlockList]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountPublisherBlockLists(ctx context.Context, client *core.Client, id string, params GetAdAccountPublisherBlockListsParams) (*core.Cursor[objects.PublisherBlockList], error) {
	var out core.Cursor[objects.PublisherBlockList]
	call := GetAdAccountPublisherBlockListsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountPublisherBlockListsParams struct {
	Name  *string     `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountPublisherBlockListsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func CreateAdAccountPublisherBlockListsBatchCall(id string, params CreateAdAccountPublisherBlockListsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "publisher_block_lists"), params.ToParams(), options...)
}

func NewCreateAdAccountPublisherBlockListsBatchRequest(id string, params CreateAdAccountPublisherBlockListsParams, options ...core.BatchOption) *core.BatchRequest[objects.PublisherBlockList] {
	return core.NewBatchRequest[objects.PublisherBlockList](CreateAdAccountPublisherBlockListsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountPublisherBlockListsBatchResponse(response *core.BatchResponse) (*objects.PublisherBlockList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PublisherBlockList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountPublisherBlockLists(ctx context.Context, client *core.Client, id string, params CreateAdAccountPublisherBlockListsParams) (*objects.PublisherBlockList, error) {
	var out objects.PublisherBlockList
	call := CreateAdAccountPublisherBlockListsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountReachestimateParams struct {
	AdgroupIds         *[]core.ID        `facebook:"adgroup_ids"`
	CallerID           *core.ID          `facebook:"caller_id"`
	Concepts           *string           `facebook:"concepts"`
	CreativeActionSpec *string           `facebook:"creative_action_spec"`
	IsDebug            *bool             `facebook:"is_debug"`
	ObjectStoreURL     *string           `facebook:"object_store_url"`
	TargetingSpec      objects.Targeting `facebook:"targeting_spec"`
	Extra              core.Params       `facebook:"-"`
}

func (params GetAdAccountReachestimateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdgroupIds != nil {
		out["adgroup_ids"] = *params.AdgroupIds
	}
	if params.CallerID != nil {
		out["caller_id"] = *params.CallerID
	}
	if params.Concepts != nil {
		out["concepts"] = *params.Concepts
	}
	if params.CreativeActionSpec != nil {
		out["creative_action_spec"] = *params.CreativeActionSpec
	}
	if params.IsDebug != nil {
		out["is_debug"] = *params.IsDebug
	}
	if params.ObjectStoreURL != nil {
		out["object_store_url"] = *params.ObjectStoreURL
	}
	out["targeting_spec"] = params.TargetingSpec
	return out
}

func GetAdAccountReachestimateBatchCall(id string, params GetAdAccountReachestimateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reachestimate"), params.ToParams(), options...)
}

func NewGetAdAccountReachestimateBatchRequest(id string, params GetAdAccountReachestimateParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountReachEstimate]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountReachEstimate]](GetAdAccountReachestimateBatchCall(id, params, options...))
}

func DecodeGetAdAccountReachestimateBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountReachEstimate], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountReachEstimate]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountReachestimate(ctx context.Context, client *core.Client, id string, params GetAdAccountReachestimateParams) (*core.Cursor[objects.AdAccountReachEstimate], error) {
	var out core.Cursor[objects.AdAccountReachEstimate]
	call := GetAdAccountReachestimateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountReachfrequencypredictionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountReachfrequencypredictionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountReachfrequencypredictionsBatchCall(id string, params GetAdAccountReachfrequencypredictionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reachfrequencypredictions"), params.ToParams(), options...)
}

func NewGetAdAccountReachfrequencypredictionsBatchRequest(id string, params GetAdAccountReachfrequencypredictionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ReachFrequencyPrediction]] {
	return core.NewBatchRequest[core.Cursor[objects.ReachFrequencyPrediction]](GetAdAccountReachfrequencypredictionsBatchCall(id, params, options...))
}

func DecodeGetAdAccountReachfrequencypredictionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ReachFrequencyPrediction], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ReachFrequencyPrediction]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountReachfrequencypredictions(ctx context.Context, client *core.Client, id string, params GetAdAccountReachfrequencypredictionsParams) (*core.Cursor[objects.ReachFrequencyPrediction], error) {
	var out core.Cursor[objects.ReachFrequencyPrediction]
	call := GetAdAccountReachfrequencypredictionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountReachfrequencypredictionsParams struct {
	Action                          *enums.AdaccountreachfrequencypredictionsActionEnumParam             `facebook:"action"`
	AdFormats                       *[]map[string]interface{}                                            `facebook:"ad_formats"`
	AuctionEntryOptionIndex         *uint64                                                              `facebook:"auction_entry_option_index"`
	Budget                          *uint64                                                              `facebook:"budget"`
	BuyingType                      *enums.AdaccountreachfrequencypredictionsBuyingTypeEnumParam         `facebook:"buying_type"`
	CampaignGroupID                 *core.ID                                                             `facebook:"campaign_group_id"`
	DayPartingSchedule              *[]map[string]interface{}                                            `facebook:"day_parting_schedule"`
	DealID                          *core.ID                                                             `facebook:"deal_id"`
	DestinationID                   *core.ID                                                             `facebook:"destination_id"`
	DestinationIds                  *[]core.ID                                                           `facebook:"destination_ids"`
	EndTime                         *uint64                                                              `facebook:"end_time"`
	Exceptions                      *bool                                                                `facebook:"exceptions"`
	ExistingCampaignID              *core.ID                                                             `facebook:"existing_campaign_id"`
	ExpirationTime                  *uint64                                                              `facebook:"expiration_time"`
	FrequencyCap                    *uint64                                                              `facebook:"frequency_cap"`
	GrpBuying                       *bool                                                                `facebook:"grp_buying"`
	Impression                      *uint64                                                              `facebook:"impression"`
	InstreamPackages                *[]enums.AdaccountreachfrequencypredictionsInstreamPackagesEnumParam `facebook:"instream_packages"`
	IntervalFrequencyCapResetPeriod *uint64                                                              `facebook:"interval_frequency_cap_reset_period"`
	IsBalancedFrequency             *bool                                                                `facebook:"is_balanced_frequency"`
	IsBonusMedia                    *bool                                                                `facebook:"is_bonus_media"`
	IsConversionGoal                *bool                                                                `facebook:"is_conversion_goal"`
	IsFullView                      *bool                                                                `facebook:"is_full_view"`
	IsHigherAverageFrequency        *bool                                                                `facebook:"is_higher_average_frequency"`
	IsReachAndFrequencyIoBuying     *bool                                                                `facebook:"is_reach_and_frequency_io_buying"`
	IsReservedBuying                *bool                                                                `facebook:"is_reserved_buying"`
	MetaMomentMakerSpec             *map[string]interface{}                                              `facebook:"meta_moment_maker_spec"`
	NumCurvePoints                  *uint64                                                              `facebook:"num_curve_points"`
	Objective                       *string                                                              `facebook:"objective"`
	OptimizationGoal                *string                                                              `facebook:"optimization_goal"`
	PredictionMode                  *uint64                                                              `facebook:"prediction_mode"`
	Reach                           *uint64                                                              `facebook:"reach"`
	RfPredictionID                  *core.ID                                                             `facebook:"rf_prediction_id"`
	RfPredictionIDToRelease         *string                                                              `facebook:"rf_prediction_id_to_release"`
	RfPredictionIDToShare           *string                                                              `facebook:"rf_prediction_id_to_share"`
	StartTime                       *uint64                                                              `facebook:"start_time"`
	StopTime                        *uint64                                                              `facebook:"stop_time"`
	StoryEventType                  *uint64                                                              `facebook:"story_event_type"`
	TargetCpm                       *uint64                                                              `facebook:"target_cpm"`
	TargetFrequency                 *uint64                                                              `facebook:"target_frequency"`
	TargetFrequencyResetPeriod      *uint64                                                              `facebook:"target_frequency_reset_period"`
	TargetSpec                      *objects.Targeting                                                   `facebook:"target_spec"`
	TrendingTopicsSpec              *map[string]interface{}                                              `facebook:"trending_topics_spec"`
	VideoViewLengthConstraint       *uint64                                                              `facebook:"video_view_length_constraint"`
	Extra                           core.Params                                                          `facebook:"-"`
}

func (params CreateAdAccountReachfrequencypredictionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Action != nil {
		out["action"] = *params.Action
	}
	if params.AdFormats != nil {
		out["ad_formats"] = *params.AdFormats
	}
	if params.AuctionEntryOptionIndex != nil {
		out["auction_entry_option_index"] = *params.AuctionEntryOptionIndex
	}
	if params.Budget != nil {
		out["budget"] = *params.Budget
	}
	if params.BuyingType != nil {
		out["buying_type"] = *params.BuyingType
	}
	if params.CampaignGroupID != nil {
		out["campaign_group_id"] = *params.CampaignGroupID
	}
	if params.DayPartingSchedule != nil {
		out["day_parting_schedule"] = *params.DayPartingSchedule
	}
	if params.DealID != nil {
		out["deal_id"] = *params.DealID
	}
	if params.DestinationID != nil {
		out["destination_id"] = *params.DestinationID
	}
	if params.DestinationIds != nil {
		out["destination_ids"] = *params.DestinationIds
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Exceptions != nil {
		out["exceptions"] = *params.Exceptions
	}
	if params.ExistingCampaignID != nil {
		out["existing_campaign_id"] = *params.ExistingCampaignID
	}
	if params.ExpirationTime != nil {
		out["expiration_time"] = *params.ExpirationTime
	}
	if params.FrequencyCap != nil {
		out["frequency_cap"] = *params.FrequencyCap
	}
	if params.GrpBuying != nil {
		out["grp_buying"] = *params.GrpBuying
	}
	if params.Impression != nil {
		out["impression"] = *params.Impression
	}
	if params.InstreamPackages != nil {
		out["instream_packages"] = *params.InstreamPackages
	}
	if params.IntervalFrequencyCapResetPeriod != nil {
		out["interval_frequency_cap_reset_period"] = *params.IntervalFrequencyCapResetPeriod
	}
	if params.IsBalancedFrequency != nil {
		out["is_balanced_frequency"] = *params.IsBalancedFrequency
	}
	if params.IsBonusMedia != nil {
		out["is_bonus_media"] = *params.IsBonusMedia
	}
	if params.IsConversionGoal != nil {
		out["is_conversion_goal"] = *params.IsConversionGoal
	}
	if params.IsFullView != nil {
		out["is_full_view"] = *params.IsFullView
	}
	if params.IsHigherAverageFrequency != nil {
		out["is_higher_average_frequency"] = *params.IsHigherAverageFrequency
	}
	if params.IsReachAndFrequencyIoBuying != nil {
		out["is_reach_and_frequency_io_buying"] = *params.IsReachAndFrequencyIoBuying
	}
	if params.IsReservedBuying != nil {
		out["is_reserved_buying"] = *params.IsReservedBuying
	}
	if params.MetaMomentMakerSpec != nil {
		out["meta_moment_maker_spec"] = *params.MetaMomentMakerSpec
	}
	if params.NumCurvePoints != nil {
		out["num_curve_points"] = *params.NumCurvePoints
	}
	if params.Objective != nil {
		out["objective"] = *params.Objective
	}
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.PredictionMode != nil {
		out["prediction_mode"] = *params.PredictionMode
	}
	if params.Reach != nil {
		out["reach"] = *params.Reach
	}
	if params.RfPredictionID != nil {
		out["rf_prediction_id"] = *params.RfPredictionID
	}
	if params.RfPredictionIDToRelease != nil {
		out["rf_prediction_id_to_release"] = *params.RfPredictionIDToRelease
	}
	if params.RfPredictionIDToShare != nil {
		out["rf_prediction_id_to_share"] = *params.RfPredictionIDToShare
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.StopTime != nil {
		out["stop_time"] = *params.StopTime
	}
	if params.StoryEventType != nil {
		out["story_event_type"] = *params.StoryEventType
	}
	if params.TargetCpm != nil {
		out["target_cpm"] = *params.TargetCpm
	}
	if params.TargetFrequency != nil {
		out["target_frequency"] = *params.TargetFrequency
	}
	if params.TargetFrequencyResetPeriod != nil {
		out["target_frequency_reset_period"] = *params.TargetFrequencyResetPeriod
	}
	if params.TargetSpec != nil {
		out["target_spec"] = *params.TargetSpec
	}
	if params.TrendingTopicsSpec != nil {
		out["trending_topics_spec"] = *params.TrendingTopicsSpec
	}
	if params.VideoViewLengthConstraint != nil {
		out["video_view_length_constraint"] = *params.VideoViewLengthConstraint
	}
	return out
}

func CreateAdAccountReachfrequencypredictionsBatchCall(id string, params CreateAdAccountReachfrequencypredictionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "reachfrequencypredictions"), params.ToParams(), options...)
}

func NewCreateAdAccountReachfrequencypredictionsBatchRequest(id string, params CreateAdAccountReachfrequencypredictionsParams, options ...core.BatchOption) *core.BatchRequest[objects.ReachFrequencyPrediction] {
	return core.NewBatchRequest[objects.ReachFrequencyPrediction](CreateAdAccountReachfrequencypredictionsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountReachfrequencypredictionsBatchResponse(response *core.BatchResponse) (*objects.ReachFrequencyPrediction, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ReachFrequencyPrediction
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountReachfrequencypredictions(ctx context.Context, client *core.Client, id string, params CreateAdAccountReachfrequencypredictionsParams) (*objects.ReachFrequencyPrediction, error) {
	var out objects.ReachFrequencyPrediction
	call := CreateAdAccountReachfrequencypredictionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountRecommendationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountRecommendationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountRecommendationsBatchCall(id string, params GetAdAccountRecommendationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "recommendations"), params.ToParams(), options...)
}

func NewGetAdAccountRecommendationsBatchRequest(id string, params GetAdAccountRecommendationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountRecommendations]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountRecommendations]](GetAdAccountRecommendationsBatchCall(id, params, options...))
}

func DecodeGetAdAccountRecommendationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountRecommendations], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountRecommendations]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountRecommendations(ctx context.Context, client *core.Client, id string, params GetAdAccountRecommendationsParams) (*core.Cursor[objects.AdAccountRecommendations], error) {
	var out core.Cursor[objects.AdAccountRecommendations]
	call := GetAdAccountRecommendationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountRecommendationsParams struct {
	AscFragmentationParameters  *map[string]interface{} `facebook:"asc_fragmentation_parameters"`
	AutoflowParameters          *map[string]interface{} `facebook:"autoflow_parameters"`
	ExtraData                   *map[string]interface{} `facebook:"extra_data"`
	FragmentationParameters     *map[string]interface{} `facebook:"fragmentation_parameters"`
	MusicParameters             *map[string]interface{} `facebook:"music_parameters"`
	RecommendationSignature     string                  `facebook:"recommendation_signature"`
	ScaleGoodCampaignParameters *map[string]interface{} `facebook:"scale_good_campaign_parameters"`
	Extra                       core.Params             `facebook:"-"`
}

func (params CreateAdAccountRecommendationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AscFragmentationParameters != nil {
		out["asc_fragmentation_parameters"] = *params.AscFragmentationParameters
	}
	if params.AutoflowParameters != nil {
		out["autoflow_parameters"] = *params.AutoflowParameters
	}
	if params.ExtraData != nil {
		out["extra_data"] = *params.ExtraData
	}
	if params.FragmentationParameters != nil {
		out["fragmentation_parameters"] = *params.FragmentationParameters
	}
	if params.MusicParameters != nil {
		out["music_parameters"] = *params.MusicParameters
	}
	out["recommendation_signature"] = params.RecommendationSignature
	if params.ScaleGoodCampaignParameters != nil {
		out["scale_good_campaign_parameters"] = *params.ScaleGoodCampaignParameters
	}
	return out
}

func CreateAdAccountRecommendationsBatchCall(id string, params CreateAdAccountRecommendationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "recommendations"), params.ToParams(), options...)
}

func NewCreateAdAccountRecommendationsBatchRequest(id string, params CreateAdAccountRecommendationsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountRecommendations] {
	return core.NewBatchRequest[objects.AdAccountRecommendations](CreateAdAccountRecommendationsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountRecommendationsBatchResponse(response *core.BatchResponse) (*objects.AdAccountRecommendations, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountRecommendations
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountRecommendations(ctx context.Context, client *core.Client, id string, params CreateAdAccountRecommendationsParams) (*objects.AdAccountRecommendations, error) {
	var out objects.AdAccountRecommendations
	call := CreateAdAccountRecommendationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountSavedAudiencesParams struct {
	BusinessID *core.ID                  `facebook:"business_id"`
	Fields     *[]string                 `facebook:"fields"`
	Filtering  *[]map[string]interface{} `facebook:"filtering"`
	Extra      core.Params               `facebook:"-"`
}

func (params GetAdAccountSavedAudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	return out
}

func GetAdAccountSavedAudiencesBatchCall(id string, params GetAdAccountSavedAudiencesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "saved_audiences"), params.ToParams(), options...)
}

func NewGetAdAccountSavedAudiencesBatchRequest(id string, params GetAdAccountSavedAudiencesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.SavedAudience]] {
	return core.NewBatchRequest[core.Cursor[objects.SavedAudience]](GetAdAccountSavedAudiencesBatchCall(id, params, options...))
}

func DecodeGetAdAccountSavedAudiencesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.SavedAudience], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.SavedAudience]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountSavedAudiences(ctx context.Context, client *core.Client, id string, params GetAdAccountSavedAudiencesParams) (*core.Cursor[objects.SavedAudience], error) {
	var out core.Cursor[objects.SavedAudience]
	call := GetAdAccountSavedAudiencesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountSubscribedAppsParams struct {
	AppID *core.ID    `facebook:"app_id"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	return out
}

func DeleteAdAccountSubscribedAppsBatchCall(id string, params DeleteAdAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewDeleteAdAccountSubscribedAppsBatchRequest(id string, params DeleteAdAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params DeleteAdAccountSubscribedAppsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountSubscribedAppsBatchCall(id string, params GetAdAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewGetAdAccountSubscribedAppsBatchRequest(id string, params GetAdAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountSubscribedApps]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountSubscribedApps]](GetAdAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeGetAdAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountSubscribedApps], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountSubscribedApps]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params GetAdAccountSubscribedAppsParams) (*core.Cursor[objects.AdAccountSubscribedApps], error) {
	var out core.Cursor[objects.AdAccountSubscribedApps]
	call := GetAdAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountSubscribedAppsParams struct {
	AppID *core.ID    `facebook:"app_id"`
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	return out
}

func CreateAdAccountSubscribedAppsBatchCall(id string, params CreateAdAccountSubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewCreateAdAccountSubscribedAppsBatchRequest(id string, params CreateAdAccountSubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountSubscribedApps] {
	return core.NewBatchRequest[objects.AdAccountSubscribedApps](CreateAdAccountSubscribedAppsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountSubscribedAppsBatchResponse(response *core.BatchResponse) (*objects.AdAccountSubscribedApps, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountSubscribedApps
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params CreateAdAccountSubscribedAppsParams) (*objects.AdAccountSubscribedApps, error) {
	var out objects.AdAccountSubscribedApps
	call := CreateAdAccountSubscribedAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTargetingbrowseParams struct {
	ExcludedCategory    *string                                                       `facebook:"excluded_category"`
	IncludeNodes        *bool                                                         `facebook:"include_nodes"`
	IsExclusion         *bool                                                         `facebook:"is_exclusion"`
	LimitType           *enums.AdaccounttargetingbrowseLimitTypeEnumParam             `facebook:"limit_type"`
	RegulatedCategories *[]enums.AdaccounttargetingbrowseRegulatedCategoriesEnumParam `facebook:"regulated_categories"`
	RegulatedCountries  *[]enums.AdaccounttargetingbrowseRegulatedCountriesEnumParam  `facebook:"regulated_countries"`
	WhitelistedTypes    *[]enums.AdaccounttargetingbrowseWhitelistedTypesEnumParam    `facebook:"whitelisted_types"`
	Extra               core.Params                                                   `facebook:"-"`
}

func (params GetAdAccountTargetingbrowseParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ExcludedCategory != nil {
		out["excluded_category"] = *params.ExcludedCategory
	}
	if params.IncludeNodes != nil {
		out["include_nodes"] = *params.IncludeNodes
	}
	if params.IsExclusion != nil {
		out["is_exclusion"] = *params.IsExclusion
	}
	if params.LimitType != nil {
		out["limit_type"] = *params.LimitType
	}
	if params.RegulatedCategories != nil {
		out["regulated_categories"] = *params.RegulatedCategories
	}
	if params.RegulatedCountries != nil {
		out["regulated_countries"] = *params.RegulatedCountries
	}
	if params.WhitelistedTypes != nil {
		out["whitelisted_types"] = *params.WhitelistedTypes
	}
	return out
}

func GetAdAccountTargetingbrowseBatchCall(id string, params GetAdAccountTargetingbrowseParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingbrowse"), params.ToParams(), options...)
}

func NewGetAdAccountTargetingbrowseBatchRequest(id string, params GetAdAccountTargetingbrowseParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountTargetingUnified]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountTargetingUnified]](GetAdAccountTargetingbrowseBatchCall(id, params, options...))
}

func DecodeGetAdAccountTargetingbrowseBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountTargetingUnified]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTargetingbrowse(ctx context.Context, client *core.Client, id string, params GetAdAccountTargetingbrowseParams) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	var out core.Cursor[objects.AdAccountTargetingUnified]
	call := GetAdAccountTargetingbrowseBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTargetingsearchParams struct {
	AllowOnlyFatHeadInterests          *bool                                                         `facebook:"allow_only_fat_head_interests"`
	AppStore                           *enums.AdaccounttargetingsearchAppStoreEnumParam              `facebook:"app_store"`
	Countries                          *[]string                                                     `facebook:"countries"`
	IsAccountLevelBrandSafetyExclusion *bool                                                         `facebook:"is_account_level_brand_safety_exclusion"`
	IsAccountLevelEmployerExclusion    *bool                                                         `facebook:"is_account_level_employer_exclusion"`
	IsExclusion                        *bool                                                         `facebook:"is_exclusion"`
	LimitType                          *enums.AdaccounttargetingsearchLimitTypeEnumParam             `facebook:"limit_type"`
	Objective                          *enums.AdaccounttargetingsearchObjectiveEnumParam             `facebook:"objective"`
	PromotedObject                     *map[string]interface{}                                       `facebook:"promoted_object"`
	Q                                  string                                                        `facebook:"q"`
	RegulatedCategories                *[]enums.AdaccounttargetingsearchRegulatedCategoriesEnumParam `facebook:"regulated_categories"`
	RegulatedCountries                 *[]enums.AdaccounttargetingsearchRegulatedCountriesEnumParam  `facebook:"regulated_countries"`
	SessionID                          *core.ID                                                      `facebook:"session_id"`
	TargetingList                      *[]map[string]interface{}                                     `facebook:"targeting_list"`
	WhitelistedTypes                   *[]enums.AdaccounttargetingsearchWhitelistedTypesEnumParam    `facebook:"whitelisted_types"`
	Extra                              core.Params                                                   `facebook:"-"`
}

func (params GetAdAccountTargetingsearchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowOnlyFatHeadInterests != nil {
		out["allow_only_fat_head_interests"] = *params.AllowOnlyFatHeadInterests
	}
	if params.AppStore != nil {
		out["app_store"] = *params.AppStore
	}
	if params.Countries != nil {
		out["countries"] = *params.Countries
	}
	if params.IsAccountLevelBrandSafetyExclusion != nil {
		out["is_account_level_brand_safety_exclusion"] = *params.IsAccountLevelBrandSafetyExclusion
	}
	if params.IsAccountLevelEmployerExclusion != nil {
		out["is_account_level_employer_exclusion"] = *params.IsAccountLevelEmployerExclusion
	}
	if params.IsExclusion != nil {
		out["is_exclusion"] = *params.IsExclusion
	}
	if params.LimitType != nil {
		out["limit_type"] = *params.LimitType
	}
	if params.Objective != nil {
		out["objective"] = *params.Objective
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	out["q"] = params.Q
	if params.RegulatedCategories != nil {
		out["regulated_categories"] = *params.RegulatedCategories
	}
	if params.RegulatedCountries != nil {
		out["regulated_countries"] = *params.RegulatedCountries
	}
	if params.SessionID != nil {
		out["session_id"] = *params.SessionID
	}
	if params.TargetingList != nil {
		out["targeting_list"] = *params.TargetingList
	}
	if params.WhitelistedTypes != nil {
		out["whitelisted_types"] = *params.WhitelistedTypes
	}
	return out
}

func GetAdAccountTargetingsearchBatchCall(id string, params GetAdAccountTargetingsearchParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingsearch"), params.ToParams(), options...)
}

func NewGetAdAccountTargetingsearchBatchRequest(id string, params GetAdAccountTargetingsearchParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountTargetingUnified]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountTargetingUnified]](GetAdAccountTargetingsearchBatchCall(id, params, options...))
}

func DecodeGetAdAccountTargetingsearchBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountTargetingUnified]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTargetingsearch(ctx context.Context, client *core.Client, id string, params GetAdAccountTargetingsearchParams) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	var out core.Cursor[objects.AdAccountTargetingUnified]
	call := GetAdAccountTargetingsearchBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTargetingsentencelinesParams struct {
	DiscardAges                 *bool             `facebook:"discard_ages"`
	DiscardPlacements           *bool             `facebook:"discard_placements"`
	HideTargetingSpecFromReturn *bool             `facebook:"hide_targeting_spec_from_return"`
	TargetingSpec               objects.Targeting `facebook:"targeting_spec"`
	Extra                       core.Params       `facebook:"-"`
}

func (params GetAdAccountTargetingsentencelinesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DiscardAges != nil {
		out["discard_ages"] = *params.DiscardAges
	}
	if params.DiscardPlacements != nil {
		out["discard_placements"] = *params.DiscardPlacements
	}
	if params.HideTargetingSpecFromReturn != nil {
		out["hide_targeting_spec_from_return"] = *params.HideTargetingSpecFromReturn
	}
	out["targeting_spec"] = params.TargetingSpec
	return out
}

func GetAdAccountTargetingsentencelinesBatchCall(id string, params GetAdAccountTargetingsentencelinesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingsentencelines"), params.ToParams(), options...)
}

func NewGetAdAccountTargetingsentencelinesBatchRequest(id string, params GetAdAccountTargetingsentencelinesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.TargetingSentenceLine]] {
	return core.NewBatchRequest[core.Cursor[objects.TargetingSentenceLine]](GetAdAccountTargetingsentencelinesBatchCall(id, params, options...))
}

func DecodeGetAdAccountTargetingsentencelinesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.TargetingSentenceLine], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.TargetingSentenceLine]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTargetingsentencelines(ctx context.Context, client *core.Client, id string, params GetAdAccountTargetingsentencelinesParams) (*core.Cursor[objects.TargetingSentenceLine], error) {
	var out core.Cursor[objects.TargetingSentenceLine]
	call := GetAdAccountTargetingsentencelinesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTargetingsuggestionsParams struct {
	AppStore            *enums.AdaccounttargetingsuggestionsAppStoreEnumParam              `facebook:"app_store"`
	Countries           *[]string                                                          `facebook:"countries"`
	LimitType           *enums.AdaccounttargetingsuggestionsLimitTypeEnumParam             `facebook:"limit_type"`
	Mode                *enums.AdaccounttargetingsuggestionsModeEnumParam                  `facebook:"mode"`
	Objective           *enums.AdaccounttargetingsuggestionsObjectiveEnumParam             `facebook:"objective"`
	Objects             *map[string]interface{}                                            `facebook:"objects"`
	RegulatedCategories *[]enums.AdaccounttargetingsuggestionsRegulatedCategoriesEnumParam `facebook:"regulated_categories"`
	RegulatedCountries  *[]enums.AdaccounttargetingsuggestionsRegulatedCountriesEnumParam  `facebook:"regulated_countries"`
	SessionID           *core.ID                                                           `facebook:"session_id"`
	TargetingList       *[]map[string]interface{}                                          `facebook:"targeting_list"`
	WhitelistedTypes    *[]enums.AdaccounttargetingsuggestionsWhitelistedTypesEnumParam    `facebook:"whitelisted_types"`
	Extra               core.Params                                                        `facebook:"-"`
}

func (params GetAdAccountTargetingsuggestionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppStore != nil {
		out["app_store"] = *params.AppStore
	}
	if params.Countries != nil {
		out["countries"] = *params.Countries
	}
	if params.LimitType != nil {
		out["limit_type"] = *params.LimitType
	}
	if params.Mode != nil {
		out["mode"] = *params.Mode
	}
	if params.Objective != nil {
		out["objective"] = *params.Objective
	}
	if params.Objects != nil {
		out["objects"] = *params.Objects
	}
	if params.RegulatedCategories != nil {
		out["regulated_categories"] = *params.RegulatedCategories
	}
	if params.RegulatedCountries != nil {
		out["regulated_countries"] = *params.RegulatedCountries
	}
	if params.SessionID != nil {
		out["session_id"] = *params.SessionID
	}
	if params.TargetingList != nil {
		out["targeting_list"] = *params.TargetingList
	}
	if params.WhitelistedTypes != nil {
		out["whitelisted_types"] = *params.WhitelistedTypes
	}
	return out
}

func GetAdAccountTargetingsuggestionsBatchCall(id string, params GetAdAccountTargetingsuggestionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingsuggestions"), params.ToParams(), options...)
}

func NewGetAdAccountTargetingsuggestionsBatchRequest(id string, params GetAdAccountTargetingsuggestionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountTargetingUnified]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountTargetingUnified]](GetAdAccountTargetingsuggestionsBatchCall(id, params, options...))
}

func DecodeGetAdAccountTargetingsuggestionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountTargetingUnified]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTargetingsuggestions(ctx context.Context, client *core.Client, id string, params GetAdAccountTargetingsuggestionsParams) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	var out core.Cursor[objects.AdAccountTargetingUnified]
	call := GetAdAccountTargetingsuggestionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTargetingvalidationParams struct {
	IDList        *[]uint64                 `facebook:"id_list"`
	IsExclusion   *bool                     `facebook:"is_exclusion"`
	NameList      *[]string                 `facebook:"name_list"`
	TargetingList *[]map[string]interface{} `facebook:"targeting_list"`
	Extra         core.Params               `facebook:"-"`
}

func (params GetAdAccountTargetingvalidationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IDList != nil {
		out["id_list"] = *params.IDList
	}
	if params.IsExclusion != nil {
		out["is_exclusion"] = *params.IsExclusion
	}
	if params.NameList != nil {
		out["name_list"] = *params.NameList
	}
	if params.TargetingList != nil {
		out["targeting_list"] = *params.TargetingList
	}
	return out
}

func GetAdAccountTargetingvalidationBatchCall(id string, params GetAdAccountTargetingvalidationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingvalidation"), params.ToParams(), options...)
}

func NewGetAdAccountTargetingvalidationBatchRequest(id string, params GetAdAccountTargetingvalidationParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountTargetingUnified]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountTargetingUnified]](GetAdAccountTargetingvalidationBatchCall(id, params, options...))
}

func DecodeGetAdAccountTargetingvalidationBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountTargetingUnified]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTargetingvalidation(ctx context.Context, client *core.Client, id string, params GetAdAccountTargetingvalidationParams) (*core.Cursor[objects.AdAccountTargetingUnified], error) {
	var out core.Cursor[objects.AdAccountTargetingUnified]
	call := GetAdAccountTargetingvalidationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountTrackingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountTrackingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountTrackingBatchCall(id string, params GetAdAccountTrackingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tracking"), params.ToParams(), options...)
}

func NewGetAdAccountTrackingBatchRequest(id string, params GetAdAccountTrackingParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountTrackingData]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountTrackingData]](GetAdAccountTrackingBatchCall(id, params, options...))
}

func DecodeGetAdAccountTrackingBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountTrackingData], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountTrackingData]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountTracking(ctx context.Context, client *core.Client, id string, params GetAdAccountTrackingParams) (*core.Cursor[objects.AdAccountTrackingData], error) {
	var out core.Cursor[objects.AdAccountTrackingData]
	call := GetAdAccountTrackingBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountTrackingParams struct {
	TrackingSpecs map[string]interface{} `facebook:"tracking_specs"`
	Extra         core.Params            `facebook:"-"`
}

func (params CreateAdAccountTrackingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tracking_specs"] = params.TrackingSpecs
	return out
}

func CreateAdAccountTrackingBatchCall(id string, params CreateAdAccountTrackingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "tracking"), params.ToParams(), options...)
}

func NewCreateAdAccountTrackingBatchRequest(id string, params CreateAdAccountTrackingParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](CreateAdAccountTrackingBatchCall(id, params, options...))
}

func DecodeCreateAdAccountTrackingBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountTracking(ctx context.Context, client *core.Client, id string, params CreateAdAccountTrackingParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := CreateAdAccountTrackingBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountUsersBatchCall(id string, params GetAdAccountUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "users"), params.ToParams(), options...)
}

func NewGetAdAccountUsersBatchRequest(id string, params GetAdAccountUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccountUser]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccountUser]](GetAdAccountUsersBatchCall(id, params, options...))
}

func DecodeGetAdAccountUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccountUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccountUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountUsers(ctx context.Context, client *core.Client, id string, params GetAdAccountUsersParams) (*core.Cursor[objects.AdAccountUser], error) {
	var out core.Cursor[objects.AdAccountUser]
	call := GetAdAccountUsersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteAdAccountUsersofanyaudienceParams struct {
	Namespace *string                 `facebook:"namespace"`
	Payload   *map[string]interface{} `facebook:"payload"`
	Session   *map[string]interface{} `facebook:"session"`
	Extra     core.Params             `facebook:"-"`
}

func (params DeleteAdAccountUsersofanyaudienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Namespace != nil {
		out["namespace"] = *params.Namespace
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Session != nil {
		out["session"] = *params.Session
	}
	return out
}

func DeleteAdAccountUsersofanyaudienceBatchCall(id string, params DeleteAdAccountUsersofanyaudienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "usersofanyaudience"), params.ToParams(), options...)
}

func NewDeleteAdAccountUsersofanyaudienceBatchRequest(id string, params DeleteAdAccountUsersofanyaudienceParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAccountUsersofanyaudienceBatchCall(id, params, options...))
}

func DecodeDeleteAdAccountUsersofanyaudienceBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAdAccountUsersofanyaudience(ctx context.Context, client *core.Client, id string, params DeleteAdAccountUsersofanyaudienceParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteAdAccountUsersofanyaudienceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountValueRuleSetParams struct {
	ProductType *enums.AdaccountvalueRuleSetProductTypeEnumParam `facebook:"product_type"`
	Status      *enums.AdaccountvalueRuleSetStatusEnumParam      `facebook:"status"`
	Extra       core.Params                                      `facebook:"-"`
}

func (params GetAdAccountValueRuleSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetAdAccountValueRuleSetBatchCall(id string, params GetAdAccountValueRuleSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "value_rule_set"), params.ToParams(), options...)
}

func NewGetAdAccountValueRuleSetBatchRequest(id string, params GetAdAccountValueRuleSetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsValueAdjustmentRuleCollection]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsValueAdjustmentRuleCollection]](GetAdAccountValueRuleSetBatchCall(id, params, options...))
}

func DecodeGetAdAccountValueRuleSetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsValueAdjustmentRuleCollection], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsValueAdjustmentRuleCollection]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountValueRuleSet(ctx context.Context, client *core.Client, id string, params GetAdAccountValueRuleSetParams) (*core.Cursor[objects.AdsValueAdjustmentRuleCollection], error) {
	var out core.Cursor[objects.AdsValueAdjustmentRuleCollection]
	call := GetAdAccountValueRuleSetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountValueRuleSetParams struct {
	EntryPoint  *enums.AdaccountvalueRuleSetEntryPointEnumParam  `facebook:"entry_point"`
	Name        string                                           `facebook:"name"`
	ProductType *enums.AdaccountvalueRuleSetProductTypeEnumParam `facebook:"product_type"`
	Rules       []map[string]interface{}                         `facebook:"rules"`
	Extra       core.Params                                      `facebook:"-"`
}

func (params CreateAdAccountValueRuleSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EntryPoint != nil {
		out["entry_point"] = *params.EntryPoint
	}
	out["name"] = params.Name
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	out["rules"] = params.Rules
	return out
}

func CreateAdAccountValueRuleSetBatchCall(id string, params CreateAdAccountValueRuleSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "value_rule_set"), params.ToParams(), options...)
}

func NewCreateAdAccountValueRuleSetBatchRequest(id string, params CreateAdAccountValueRuleSetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsValueAdjustmentRuleCollection] {
	return core.NewBatchRequest[objects.AdsValueAdjustmentRuleCollection](CreateAdAccountValueRuleSetBatchCall(id, params, options...))
}

func DecodeCreateAdAccountValueRuleSetBatchResponse(response *core.BatchResponse) (*objects.AdsValueAdjustmentRuleCollection, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsValueAdjustmentRuleCollection
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountValueRuleSet(ctx context.Context, client *core.Client, id string, params CreateAdAccountValueRuleSetParams) (*objects.AdsValueAdjustmentRuleCollection, error) {
	var out objects.AdsValueAdjustmentRuleCollection
	call := CreateAdAccountValueRuleSetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountValueRuleSetTranslationParams struct {
	Source *map[string]interface{} `facebook:"source"`
	Extra  core.Params             `facebook:"-"`
}

func (params CreateAdAccountValueRuleSetTranslationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func CreateAdAccountValueRuleSetTranslationBatchCall(id string, params CreateAdAccountValueRuleSetTranslationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "value_rule_set_translation"), params.ToParams(), options...)
}

func NewCreateAdAccountValueRuleSetTranslationBatchRequest(id string, params CreateAdAccountValueRuleSetTranslationParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateAdAccountValueRuleSetTranslationBatchCall(id, params, options...))
}

func DecodeCreateAdAccountValueRuleSetTranslationBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateAdAccountValueRuleSetTranslation(ctx context.Context, client *core.Client, id string, params CreateAdAccountValueRuleSetTranslationParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateAdAccountValueRuleSetTranslationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountVideoAdsParams struct {
	Since *time.Time  `facebook:"since"`
	Until *time.Time  `facebook:"until"`
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountVideoAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetAdAccountVideoAdsBatchCall(id string, params GetAdAccountVideoAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_ads"), params.ToParams(), options...)
}

func NewGetAdAccountVideoAdsBatchRequest(id string, params GetAdAccountVideoAdsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetAdAccountVideoAdsBatchCall(id, params, options...))
}

func DecodeGetAdAccountVideoAdsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
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

func GetAdAccountVideoAds(ctx context.Context, client *core.Client, id string, params GetAdAccountVideoAdsParams) (*core.Cursor[objects.AdVideo], error) {
	var out core.Cursor[objects.AdVideo]
	call := GetAdAccountVideoAdsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateAdAccountVideoAdsParams struct {
	Description *string                                     `facebook:"description"`
	Privacy     *string                                     `facebook:"privacy"`
	Title       *string                                     `facebook:"title"`
	UploadPhase enums.AdaccountvideoAdsUploadPhaseEnumParam `facebook:"upload_phase"`
	VideoID     *core.ID                                    `facebook:"video_id"`
	VideoState  *enums.AdaccountvideoAdsVideoStateEnumParam `facebook:"video_state"`
	Extra       core.Params                                 `facebook:"-"`
}

func (params CreateAdAccountVideoAdsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	out["upload_phase"] = params.UploadPhase
	if params.VideoID != nil {
		out["video_id"] = *params.VideoID
	}
	if params.VideoState != nil {
		out["video_state"] = *params.VideoState
	}
	return out
}

func CreateAdAccountVideoAdsBatchCall(id string, params CreateAdAccountVideoAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "video_ads"), params.ToParams(), options...)
}

func NewCreateAdAccountVideoAdsBatchRequest(id string, params CreateAdAccountVideoAdsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateAdAccountVideoAdsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountVideoAdsBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
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

func CreateAdAccountVideoAds(ctx context.Context, client *core.Client, id string, params CreateAdAccountVideoAdsParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	call := CreateAdAccountVideoAdsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountBatchCall(id string, params GetAdAccountParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAccountBatchRequest(id string, params GetAdAccountParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](GetAdAccountBatchCall(id, params, options...))
}

func DecodeGetAdAccountBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccount(ctx context.Context, client *core.Client, id string, params GetAdAccountParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := GetAdAccountBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateAdAccountParams struct {
	AgencyClientDeclaration *map[string]interface{}   `facebook:"agency_client_declaration"`
	AttributionSpec         *[]map[string]interface{} `facebook:"attribution_spec"`
	BusinessInfo            *map[string]interface{}   `facebook:"business_info"`
	Currency                *enums.AdaccountCurrency  `facebook:"currency"`
	CustomAudienceInfo      *map[string]interface{}   `facebook:"custom_audience_info"`
	DefaultDsaBeneficiary   *string                   `facebook:"default_dsa_beneficiary"`
	DefaultDsaPayor         *string                   `facebook:"default_dsa_payor"`
	EndAdvertiser           *string                   `facebook:"end_advertiser"`
	ExistingCustomers       *[]string                 `facebook:"existing_customers"`
	IsBaSkipDelayedEligible *bool                     `facebook:"is_ba_skip_delayed_eligible"`
	IsNotificationsEnabled  *bool                     `facebook:"is_notifications_enabled"`
	MediaAgency             *string                   `facebook:"media_agency"`
	Name                    *string                   `facebook:"name"`
	Partner                 *string                   `facebook:"partner"`
	SpendCap                *float64                  `facebook:"spend_cap"`
	SpendCapAction          *string                   `facebook:"spend_cap_action"`
	TimezoneID              *core.ID                  `facebook:"timezone_id"`
	TosAccepted             *map[string]interface{}   `facebook:"tos_accepted"`
	Extra                   core.Params               `facebook:"-"`
}

func (params UpdateAdAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AgencyClientDeclaration != nil {
		out["agency_client_declaration"] = *params.AgencyClientDeclaration
	}
	if params.AttributionSpec != nil {
		out["attribution_spec"] = *params.AttributionSpec
	}
	if params.BusinessInfo != nil {
		out["business_info"] = *params.BusinessInfo
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.CustomAudienceInfo != nil {
		out["custom_audience_info"] = *params.CustomAudienceInfo
	}
	if params.DefaultDsaBeneficiary != nil {
		out["default_dsa_beneficiary"] = *params.DefaultDsaBeneficiary
	}
	if params.DefaultDsaPayor != nil {
		out["default_dsa_payor"] = *params.DefaultDsaPayor
	}
	if params.EndAdvertiser != nil {
		out["end_advertiser"] = *params.EndAdvertiser
	}
	if params.ExistingCustomers != nil {
		out["existing_customers"] = *params.ExistingCustomers
	}
	if params.IsBaSkipDelayedEligible != nil {
		out["is_ba_skip_delayed_eligible"] = *params.IsBaSkipDelayedEligible
	}
	if params.IsNotificationsEnabled != nil {
		out["is_notifications_enabled"] = *params.IsNotificationsEnabled
	}
	if params.MediaAgency != nil {
		out["media_agency"] = *params.MediaAgency
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Partner != nil {
		out["partner"] = *params.Partner
	}
	if params.SpendCap != nil {
		out["spend_cap"] = *params.SpendCap
	}
	if params.SpendCapAction != nil {
		out["spend_cap_action"] = *params.SpendCapAction
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	if params.TosAccepted != nil {
		out["tos_accepted"] = *params.TosAccepted
	}
	return out
}

func UpdateAdAccountBatchCall(id string, params UpdateAdAccountParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdAccountBatchRequest(id string, params UpdateAdAccountParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccount] {
	return core.NewBatchRequest[objects.AdAccount](UpdateAdAccountBatchCall(id, params, options...))
}

func DecodeUpdateAdAccountBatchResponse(response *core.BatchResponse) (*objects.AdAccount, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccount
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdAccount(ctx context.Context, client *core.Client, id string, params UpdateAdAccountParams) (*objects.AdAccount, error) {
	var out objects.AdAccount
	call := UpdateAdAccountBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
