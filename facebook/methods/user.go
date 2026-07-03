package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type DeleteUserAccessTokensParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteUserAccessTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteUserAccessTokens(ctx context.Context, client *core.Client, id string, params DeleteUserAccessTokensParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "access_tokens"), params.ToParams(), &out)
	return &out, err
}

type CreateUserAccessTokensParams struct {
	BusinessApp              string      `facebook:"business_app"`
	PageID                   *core.ID    `facebook:"page_id"`
	Scope                    *[]string   `facebook:"scope"`
	SetTokenExpiresInX60Days *bool       `facebook:"set_token_expires_in_60_days"`
	Extra                    core.Params `facebook:"-"`
}

func (params CreateUserAccessTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business_app"] = params.BusinessApp
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	if params.Scope != nil {
		out["scope"] = *params.Scope
	}
	if params.SetTokenExpiresInX60Days != nil {
		out["set_token_expires_in_60_days"] = *params.SetTokenExpiresInX60Days
	}
	return out
}

func CreateUserAccessTokens(ctx context.Context, client *core.Client, id string, params CreateUserAccessTokensParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "access_tokens"), params.ToParams(), &out)
	return &out, err
}

type GetUserAccountsParams struct {
	AdID         *core.ID    `facebook:"ad_id"`
	IsPlace      *bool       `facebook:"is_place"`
	IsPromotable *bool       `facebook:"is_promotable"`
	Extra        core.Params `facebook:"-"`
}

func (params GetUserAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	if params.IsPlace != nil {
		out["is_place"] = *params.IsPlace
	}
	if params.IsPromotable != nil {
		out["is_promotable"] = *params.IsPromotable
	}
	return out
}

func GetUserAccounts(ctx context.Context, client *core.Client, id string, params GetUserAccountsParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateUserAccountsParams struct {
	About                    *string                 `facebook:"about"`
	Address                  *string                 `facebook:"address"`
	Category                 *int                    `facebook:"category"`
	CategoryEnum             *string                 `facebook:"category_enum"`
	CategoryList             *[]string               `facebook:"category_list"`
	CityID                   *core.ID                `facebook:"city_id"`
	Coordinates              *map[string]interface{} `facebook:"coordinates"`
	CoverPhoto               *map[string]interface{} `facebook:"cover_photo"`
	Description              *string                 `facebook:"description"`
	IgnoreCoordinateWarnings *bool                   `facebook:"ignore_coordinate_warnings"`
	Location                 *map[string]interface{} `facebook:"location"`
	Name                     string                  `facebook:"name"`
	Phone                    *string                 `facebook:"phone"`
	Picture                  *string                 `facebook:"picture"`
	Website                  *string                 `facebook:"website"`
	Zip                      *string                 `facebook:"zip"`
	Extra                    core.Params             `facebook:"-"`
}

func (params CreateUserAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.About != nil {
		out["about"] = *params.About
	}
	if params.Address != nil {
		out["address"] = *params.Address
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.CategoryEnum != nil {
		out["category_enum"] = *params.CategoryEnum
	}
	if params.CategoryList != nil {
		out["category_list"] = *params.CategoryList
	}
	if params.CityID != nil {
		out["city_id"] = *params.CityID
	}
	if params.Coordinates != nil {
		out["coordinates"] = *params.Coordinates
	}
	if params.CoverPhoto != nil {
		out["cover_photo"] = *params.CoverPhoto
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.IgnoreCoordinateWarnings != nil {
		out["ignore_coordinate_warnings"] = *params.IgnoreCoordinateWarnings
	}
	if params.Location != nil {
		out["location"] = *params.Location
	}
	out["name"] = params.Name
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.Website != nil {
		out["website"] = *params.Website
	}
	if params.Zip != nil {
		out["zip"] = *params.Zip
	}
	return out
}

func CreateUserAccounts(ctx context.Context, client *core.Client, id string, params CreateUserAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "accounts"), params.ToParams(), &out)
	return &out, err
}

type GetUserAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAdStudies(ctx context.Context, client *core.Client, id string, params GetUserAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), &out)
	return &out, err
}

type CreateUserAdStudiesParams struct {
	Cells              *[]map[string]interface{}         `facebook:"cells"`
	ClientBusiness     *string                           `facebook:"client_business"`
	ConfidenceLevel    *float64                          `facebook:"confidence_level"`
	CooldownStartTime  *int                              `facebook:"cooldown_start_time"`
	CreativeTestConfig *map[string]interface{}           `facebook:"creative_test_config"`
	Description        *string                           `facebook:"description"`
	EndTime            *int                              `facebook:"end_time"`
	Name               *string                           `facebook:"name"`
	Objectives         *[]map[string]interface{}         `facebook:"objectives"`
	ObservationEndTime *int                              `facebook:"observation_end_time"`
	StartTime          *int                              `facebook:"start_time"`
	Type               *enums.UseradStudiesTypeEnumParam `facebook:"type"`
	Viewers            *[]int                            `facebook:"viewers"`
	Extra              core.Params                       `facebook:"-"`
}

func (params CreateUserAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cells != nil {
		out["cells"] = *params.Cells
	}
	if params.ClientBusiness != nil {
		out["client_business"] = *params.ClientBusiness
	}
	if params.ConfidenceLevel != nil {
		out["confidence_level"] = *params.ConfidenceLevel
	}
	if params.CooldownStartTime != nil {
		out["cooldown_start_time"] = *params.CooldownStartTime
	}
	if params.CreativeTestConfig != nil {
		out["creative_test_config"] = *params.CreativeTestConfig
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Objectives != nil {
		out["objectives"] = *params.Objectives
	}
	if params.ObservationEndTime != nil {
		out["observation_end_time"] = *params.ObservationEndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Viewers != nil {
		out["viewers"] = *params.Viewers
	}
	return out
}

func CreateUserAdStudies(ctx context.Context, client *core.Client, id string, params CreateUserAdStudiesParams) (*objects.AdStudy, error) {
	var out objects.AdStudy
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ad_studies"), params.ToParams(), &out)
	return &out, err
}

type GetUserAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAdaccounts(ctx context.Context, client *core.Client, id string, params GetUserAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
	return &out, err
}

type GetUserAlbumsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAlbumsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAlbums(ctx context.Context, client *core.Client, id string, params GetUserAlbumsParams) (*core.Cursor[objects.Album], error) {
	var out core.Cursor[objects.Album]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "albums"), params.ToParams(), &out)
	return &out, err
}

type CreateUserApplicationsParams struct {
	BusinessApp int         `facebook:"business_app"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateUserApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business_app"] = params.BusinessApp
	return out
}

func CreateUserApplications(ctx context.Context, client *core.Client, id string, params CreateUserApplicationsParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "applications"), params.ToParams(), &out)
	return &out, err
}

type GetUserApprequestformerrecipientsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserApprequestformerrecipientsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserApprequestformerrecipients(ctx context.Context, client *core.Client, id string, params GetUserApprequestformerrecipientsParams) (*core.Cursor[objects.AppRequestFormerRecipient], error) {
	var out core.Cursor[objects.AppRequestFormerRecipient]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "apprequestformerrecipients"), params.ToParams(), &out)
	return &out, err
}

type GetUserApprequestsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserApprequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserApprequests(ctx context.Context, client *core.Client, id string, params GetUserApprequestsParams) (*core.Cursor[objects.AppRequest], error) {
	var out core.Cursor[objects.AppRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "apprequests"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedApplications(ctx context.Context, client *core.Client, id string, params GetUserAssignedApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_applications"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedBusinessAssetGroupsParams struct {
	ContainedAssetID *core.ID    `facebook:"contained_asset_id"`
	Extra            core.Params `facebook:"-"`
}

func (params GetUserAssignedBusinessAssetGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContainedAssetID != nil {
		out["contained_asset_id"] = *params.ContainedAssetID
	}
	return out
}

func GetUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedPagesParams struct {
	Pages *[]uint64   `facebook:"pages"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Pages != nil {
		out["pages"] = *params.Pages
	}
	return out
}

func GetUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetUserAssignedWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetUserBusinessUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserBusinessUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserBusinessUsers(ctx context.Context, client *core.Client, id string, params GetUserBusinessUsersParams) (*core.Cursor[objects.BusinessUser], error) {
	var out core.Cursor[objects.BusinessUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "business_users"), params.ToParams(), &out)
	return &out, err
}

type DeleteUserBusinessesParams struct {
	Business *string     `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Business != nil {
		out["business"] = *params.Business
	}
	return out
}

func DeleteUserBusinesses(ctx context.Context, client *core.Client, id string, params DeleteUserBusinessesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "businesses"), params.ToParams(), &out)
	return &out, err
}

type GetUserBusinessesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserBusinesses(ctx context.Context, client *core.Client, id string, params GetUserBusinessesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "businesses"), params.ToParams(), &out)
	return &out, err
}

type CreateUserBusinessesParams struct {
	ChildBusinessExternalID *core.ID                                         `facebook:"child_business_external_id"`
	Email                   *string                                          `facebook:"email"`
	Name                    string                                           `facebook:"name"`
	PrimaryPage             *string                                          `facebook:"primary_page"`
	SalesRepEmail           *string                                          `facebook:"sales_rep_email"`
	SurveyBusinessType      *enums.UserbusinessesSurveyBusinessTypeEnumParam `facebook:"survey_business_type"`
	SurveyNumAssets         *uint64                                          `facebook:"survey_num_assets"`
	SurveyNumPeople         *uint64                                          `facebook:"survey_num_people"`
	TimezoneID              *enums.UserbusinessesTimezoneIDEnumParam         `facebook:"timezone_id"`
	Vertical                enums.UserbusinessesVerticalEnumParam            `facebook:"vertical"`
	Extra                   core.Params                                      `facebook:"-"`
}

func (params CreateUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	out["name"] = params.Name
	if params.PrimaryPage != nil {
		out["primary_page"] = *params.PrimaryPage
	}
	if params.SalesRepEmail != nil {
		out["sales_rep_email"] = *params.SalesRepEmail
	}
	if params.SurveyBusinessType != nil {
		out["survey_business_type"] = *params.SurveyBusinessType
	}
	if params.SurveyNumAssets != nil {
		out["survey_num_assets"] = *params.SurveyNumAssets
	}
	if params.SurveyNumPeople != nil {
		out["survey_num_people"] = *params.SurveyNumPeople
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	out["vertical"] = params.Vertical
	return out
}

func CreateUserBusinesses(ctx context.Context, client *core.Client, id string, params CreateUserBusinessesParams) (*objects.Business, error) {
	var out objects.Business
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "businesses"), params.ToParams(), &out)
	return &out, err
}

type GetUserConversationsParams struct {
	Folder   *string                                   `facebook:"folder"`
	Platform *enums.UserconversationsPlatformEnumParam `facebook:"platform"`
	Tags     *[]string                                 `facebook:"tags"`
	UserID   *core.ID                                  `facebook:"user_id"`
	Extra    core.Params                               `facebook:"-"`
}

func (params GetUserConversationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	return out
}

func GetUserConversations(ctx context.Context, client *core.Client, id string, params GetUserConversationsParams) (*core.Cursor[objects.UnifiedThread], error) {
	var out core.Cursor[objects.UnifiedThread]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "conversations"), params.ToParams(), &out)
	return &out, err
}

type GetUserCustomLabelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserCustomLabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserCustomLabels(ctx context.Context, client *core.Client, id string, params GetUserCustomLabelsParams) (*core.Cursor[objects.PageUserMessageThreadLabel], error) {
	var out core.Cursor[objects.PageUserMessageThreadLabel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "custom_labels"), params.ToParams(), &out)
	return &out, err
}

type GetUserEventsParams struct {
	IncludeCanceled *bool                          `facebook:"include_canceled"`
	Type            *enums.UsereventsTypeEnumParam `facebook:"type"`
	Extra           core.Params                    `facebook:"-"`
}

func (params GetUserEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeCanceled != nil {
		out["include_canceled"] = *params.IncludeCanceled
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserEvents(ctx context.Context, client *core.Client, id string, params GetUserEventsParams) (*core.Cursor[objects.Event], error) {
	var out core.Cursor[objects.Event]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "events"), params.ToParams(), &out)
	return &out, err
}

type GetUserFeedParams struct {
	IncludeHidden *bool       `facebook:"include_hidden"`
	Q             *string     `facebook:"q"`
	ShowExpired   *bool       `facebook:"show_expired"`
	Since         *time.Time  `facebook:"since"`
	Until         *time.Time  `facebook:"until"`
	With          *string     `facebook:"with"`
	Extra         core.Params `facebook:"-"`
}

func (params GetUserFeedParams) ToParams() core.Params {
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

func GetUserFeed(ctx context.Context, client *core.Client, id string, params GetUserFeedParams) (*core.Cursor[objects.Post], error) {
	var out core.Cursor[objects.Post]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), &out)
	return &out, err
}

type CreateUserFeedParams struct {
	Actions                   *map[string]interface{}                          `facebook:"actions"`
	AlbumID                   *core.ID                                         `facebook:"album_id"`
	AndroidKeyHash            *string                                          `facebook:"android_key_hash"`
	ApplicationID             *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID      *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	Asset3dID                 *core.ID                                         `facebook:"asset3d_id"`
	AssociatedID              *core.ID                                         `facebook:"associated_id"`
	AttachPlaceSuggestion     *bool                                            `facebook:"attach_place_suggestion"`
	AttachedMedia             *[]map[string]interface{}                        `facebook:"attached_media"`
	AudienceExp               *bool                                            `facebook:"audience_exp"`
	BackdatedTime             *time.Time                                       `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.UserfeedBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	BreakingNews              *bool                                            `facebook:"breaking_news"`
	BreakingNewsExpiration    *uint64                                          `facebook:"breaking_news_expiration"`
	CallToAction              *map[string]interface{}                          `facebook:"call_to_action"`
	Caption                   *string                                          `facebook:"caption"`
	ChildAttachments          *[]map[string]interface{}                        `facebook:"child_attachments"`
	ClientMutationID          *core.ID                                         `facebook:"client_mutation_id"`
	ComposerEntryPicker       *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint        *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime         *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog  *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID         *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface     *string                                          `facebook:"composer_source_surface"`
	ComposerType              *string                                          `facebook:"composer_type"`
	ConnectionClass           *string                                          `facebook:"connection_class"`
	ContentAttachment         *string                                          `facebook:"content_attachment"`
	Coordinates               *map[string]interface{}                          `facebook:"coordinates"`
	CtaLink                   *string                                          `facebook:"cta_link"`
	CtaType                   *string                                          `facebook:"cta_type"`
	Description               *string                                          `facebook:"description"`
	DirectShareStatus         *uint64                                          `facebook:"direct_share_status"`
	ExpandedHeight            *uint64                                          `facebook:"expanded_height"`
	ExpandedWidth             *uint64                                          `facebook:"expanded_width"`
	FeedTargeting             *map[string]interface{}                          `facebook:"feed_targeting"`
	Formatting                *enums.UserfeedFormattingEnumParam               `facebook:"formatting"`
	FunFactPromptID           *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID          *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Height                    *uint64                                          `facebook:"height"`
	HomeCheckinCityID         *map[string]interface{}                          `facebook:"home_checkin_city_id"`
	ImageCrops                *map[string]interface{}                          `facebook:"image_crops"`
	ImplicitWithTags          *[]int                                           `facebook:"implicit_with_tags"`
	InstantGameEntryPointData *string                                          `facebook:"instant_game_entry_point_data"`
	IosBundleID               *core.ID                                         `facebook:"ios_bundle_id"`
	IsBackoutDraft            *bool                                            `facebook:"is_backout_draft"`
	IsBoostIntended           *bool                                            `facebook:"is_boost_intended"`
	IsExplicitLocation        *bool                                            `facebook:"is_explicit_location"`
	IsExplicitShare           *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost        *bool                                            `facebook:"is_group_linking_post"`
	IsPhotoContainer          *bool                                            `facebook:"is_photo_container"`
	Link                      *string                                          `facebook:"link"`
	LocationSourceID          *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy             *bool                                            `facebook:"manual_privacy"`
	Message                   *string                                          `facebook:"message"`
	MultiShareEndCard         *bool                                            `facebook:"multi_share_end_card"`
	MultiShareOptimized       *bool                                            `facebook:"multi_share_optimized"`
	Name                      *string                                          `facebook:"name"`
	NectarModule              *string                                          `facebook:"nectar_module"`
	ObjectAttachment          *string                                          `facebook:"object_attachment"`
	OgActionTypeID            *core.ID                                         `facebook:"og_action_type_id"`
	OgHideObjectAttachment    *bool                                            `facebook:"og_hide_object_attachment"`
	OgIconID                  *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                  *string                                          `facebook:"og_phrase"`
	OgSetProfileBadge         *bool                                            `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism     *string                                          `facebook:"og_suggestion_mechanism"`
	PageRecommendation        *string                                          `facebook:"page_recommendation"`
	Picture                   *string                                          `facebook:"picture"`
	Place                     *map[string]interface{}                          `facebook:"place"`
	PlaceAttachmentSetting    *enums.UserfeedPlaceAttachmentSettingEnumParam   `facebook:"place_attachment_setting"`
	PlaceList                 *string                                          `facebook:"place_list"`
	PlaceListData             *[]interface{}                                   `facebook:"place_list_data"`
	PostSurfacesBlacklist     *[]enums.UserfeedPostSurfacesBlacklistEnumParam  `facebook:"post_surfaces_blacklist"`
	PostingToRedspace         *enums.UserfeedPostingToRedspaceEnumParam        `facebook:"posting_to_redspace"`
	Privacy                   *string                                          `facebook:"privacy"`
	PromptID                  *core.ID                                         `facebook:"prompt_id"`
	PromptTrackingString      *string                                          `facebook:"prompt_tracking_string"`
	Properties                *map[string]interface{}                          `facebook:"properties"`
	ProxiedAppID              *core.ID                                         `facebook:"proxied_app_id"`
	PublishEventID            *core.ID                                         `facebook:"publish_event_id"`
	Published                 *bool                                            `facebook:"published"`
	Quote                     *string                                          `facebook:"quote"`
	Ref                       *[]string                                        `facebook:"ref"`
	ReferenceableImageIds     *[]core.ID                                       `facebook:"referenceable_image_ids"`
	ReferralID                *core.ID                                         `facebook:"referral_id"`
	ScheduledPublishTime      *time.Time                                       `facebook:"scheduled_publish_time"`
	Source                    *string                                          `facebook:"source"`
	SponsorID                 *core.ID                                         `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                                          `facebook:"sponsor_relationship"`
	SuggestedPlaceID          *map[string]interface{}                          `facebook:"suggested_place_id"`
	Tags                      *[]int                                           `facebook:"tags"`
	TargetSurface             *enums.UserfeedTargetSurfaceEnumParam            `facebook:"target_surface"`
	Targeting                 *map[string]interface{}                          `facebook:"targeting"`
	TextFormatMetadata        *string                                          `facebook:"text_format_metadata"`
	TextFormatPresetID        *core.ID                                         `facebook:"text_format_preset_id"`
	TextOnlyPlace             *string                                          `facebook:"text_only_place"`
	Thumbnail                 *core.FileParam                                  `facebook:"thumbnail"`
	TimeSinceOriginalPost     *uint64                                          `facebook:"time_since_original_post"`
	Title                     *string                                          `facebook:"title"`
	TrackingInfo              *string                                          `facebook:"tracking_info"`
	UnpublishedContentType    *enums.UserfeedUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	UserSelectedTags          *bool                                            `facebook:"user_selected_tags"`
	VideoStartTimeMs          *uint64                                          `facebook:"video_start_time_ms"`
	ViewerCoordinates         *map[string]interface{}                          `facebook:"viewer_coordinates"`
	Width                     *uint64                                          `facebook:"width"`
	Extra                     core.Params                                      `facebook:"-"`
}

func (params CreateUserFeedParams) ToParams() core.Params {
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

func CreateUserFeed(ctx context.Context, client *core.Client, id string, params CreateUserFeedParams) (*objects.Post, error) {
	var out objects.Post
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "feed"), params.ToParams(), &out)
	return &out, err
}

type GetUserFriendsParams struct {
	UID   *core.ID    `facebook:"uid"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserFriendsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	return out
}

func GetUserFriends(ctx context.Context, client *core.Client, id string, params GetUserFriendsParams) (*core.Cursor[objects.User], error) {
	var out core.Cursor[objects.User]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "friends"), params.ToParams(), &out)
	return &out, err
}

type GetUserFundraisersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserFundraisersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserFundraisers(ctx context.Context, client *core.Client, id string, params GetUserFundraisersParams) (*core.Cursor[objects.FundraiserPersonToCharity], error) {
	var out core.Cursor[objects.FundraiserPersonToCharity]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "fundraisers"), params.ToParams(), &out)
	return &out, err
}

type CreateUserFundraisersParams struct {
	CharityID              *core.ID                                     `facebook:"charity_id"`
	CoverPhoto             *core.FileParam                              `facebook:"cover_photo"`
	Currency               string                                       `facebook:"currency"`
	Description            string                                       `facebook:"description"`
	EndTime                time.Time                                    `facebook:"end_time"`
	ExternalEventName      *string                                      `facebook:"external_event_name"`
	ExternalEventStartTime *time.Time                                   `facebook:"external_event_start_time"`
	ExternalEventURI       *string                                      `facebook:"external_event_uri"`
	ExternalFundraiserURI  *string                                      `facebook:"external_fundraiser_uri"`
	ExternalID             core.ID                                      `facebook:"external_id"`
	FundraiserType         enums.UserfundraisersFundraiserTypeEnumParam `facebook:"fundraiser_type"`
	GoalAmount             uint64                                       `facebook:"goal_amount"`
	Name                   string                                       `facebook:"name"`
	PageID                 *core.ID                                     `facebook:"page_id"`
	Extra                  core.Params                                  `facebook:"-"`
}

func (params CreateUserFundraisersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CharityID != nil {
		out["charity_id"] = *params.CharityID
	}
	if params.CoverPhoto != nil {
		out["cover_photo"] = *params.CoverPhoto
	}
	out["currency"] = params.Currency
	out["description"] = params.Description
	out["end_time"] = params.EndTime
	if params.ExternalEventName != nil {
		out["external_event_name"] = *params.ExternalEventName
	}
	if params.ExternalEventStartTime != nil {
		out["external_event_start_time"] = *params.ExternalEventStartTime
	}
	if params.ExternalEventURI != nil {
		out["external_event_uri"] = *params.ExternalEventURI
	}
	if params.ExternalFundraiserURI != nil {
		out["external_fundraiser_uri"] = *params.ExternalFundraiserURI
	}
	out["external_id"] = params.ExternalID
	out["fundraiser_type"] = params.FundraiserType
	out["goal_amount"] = params.GoalAmount
	out["name"] = params.Name
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	return out
}

func CreateUserFundraisers(ctx context.Context, client *core.Client, id string, params CreateUserFundraisersParams) (*objects.FundraiserPersonToCharity, error) {
	var out objects.FundraiserPersonToCharity
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "fundraisers"), params.ToParams(), &out)
	return &out, err
}

type GetUserGroupsParams struct {
	AdminOnly *bool       `facebook:"admin_only"`
	Parent    *string     `facebook:"parent"`
	Extra     core.Params `facebook:"-"`
}

func (params GetUserGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdminOnly != nil {
		out["admin_only"] = *params.AdminOnly
	}
	if params.Parent != nil {
		out["parent"] = *params.Parent
	}
	return out
}

func GetUserGroups(ctx context.Context, client *core.Client, id string, params GetUserGroupsParams) (*core.Cursor[objects.Group], error) {
	var out core.Cursor[objects.Group]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "groups"), params.ToParams(), &out)
	return &out, err
}

type GetUserIdsForAppsParams struct {
	App   *int        `facebook:"app"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.App != nil {
		out["app"] = *params.App
	}
	return out
}

func GetUserIdsForApps(ctx context.Context, client *core.Client, id string, params GetUserIdsForAppsParams) (*core.Cursor[objects.UserIDForApp], error) {
	var out core.Cursor[objects.UserIDForApp]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ids_for_apps"), params.ToParams(), &out)
	return &out, err
}

type GetUserIdsForBusinessParams struct {
	App   *int        `facebook:"app"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.App != nil {
		out["app"] = *params.App
	}
	return out
}

func GetUserIdsForBusiness(ctx context.Context, client *core.Client, id string, params GetUserIdsForBusinessParams) (*core.Cursor[objects.UserIDForApp], error) {
	var out core.Cursor[objects.UserIDForApp]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ids_for_business"), params.ToParams(), &out)
	return &out, err
}

type GetUserIdsForPagesParams struct {
	Page  *int        `facebook:"page"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Page != nil {
		out["page"] = *params.Page
	}
	return out
}

func GetUserIdsForPages(ctx context.Context, client *core.Client, id string, params GetUserIdsForPagesParams) (*core.Cursor[objects.UserIDForPage], error) {
	var out core.Cursor[objects.UserIDForPage]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ids_for_pages"), params.ToParams(), &out)
	return &out, err
}

type GetUserLikesParams struct {
	TargetID *core.ID    `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetUserLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	return out
}

func GetUserLikes(ctx context.Context, client *core.Client, id string, params GetUserLikesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetUserLiveVideosParams struct {
	BroadcastStatus *[]enums.UserliveVideosBroadcastStatusEnumParam `facebook:"broadcast_status"`
	Source          *enums.UserliveVideosSourceEnumParam            `facebook:"source"`
	Extra           core.Params                                     `facebook:"-"`
}

func (params GetUserLiveVideosParams) ToParams() core.Params {
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

func GetUserLiveVideos(ctx context.Context, client *core.Client, id string, params GetUserLiveVideosParams) (*core.Cursor[objects.LiveVideo], error) {
	var out core.Cursor[objects.LiveVideo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), &out)
	return &out, err
}

type CreateUserLiveVideosParams struct {
	ContentTags                *[]string                                        `facebook:"content_tags"`
	Description                *string                                          `facebook:"description"`
	EnableBackupIngest         *bool                                            `facebook:"enable_backup_ingest"`
	EncodingSettings           *string                                          `facebook:"encoding_settings"`
	EventParams                *map[string]interface{}                          `facebook:"event_params"`
	FisheyeVideoCropped        *bool                                            `facebook:"fisheye_video_cropped"`
	FrontZRotation             *float64                                         `facebook:"front_z_rotation"`
	IsAudioOnly                *bool                                            `facebook:"is_audio_only"`
	IsSpherical                *bool                                            `facebook:"is_spherical"`
	OriginalFov                *uint64                                          `facebook:"original_fov"`
	Privacy                    *string                                          `facebook:"privacy"`
	Projection                 *enums.UserliveVideosProjectionEnumParam         `facebook:"projection"`
	Published                  *bool                                            `facebook:"published"`
	ScheduleCustomProfileImage *core.FileParam                                  `facebook:"schedule_custom_profile_image"`
	SpatialAudioFormat         *enums.UserliveVideosSpatialAudioFormatEnumParam `facebook:"spatial_audio_format"`
	Status                     *enums.UserliveVideosStatusEnumParam             `facebook:"status"`
	StereoscopicMode           *enums.UserliveVideosStereoscopicModeEnumParam   `facebook:"stereoscopic_mode"`
	StopOnDeleteStream         *bool                                            `facebook:"stop_on_delete_stream"`
	StreamType                 *enums.UserliveVideosStreamTypeEnumParam         `facebook:"stream_type"`
	Title                      *string                                          `facebook:"title"`
	Extra                      core.Params                                      `facebook:"-"`
}

func (params CreateUserLiveVideosParams) ToParams() core.Params {
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

func CreateUserLiveVideos(ctx context.Context, client *core.Client, id string, params CreateUserLiveVideosParams) (*objects.LiveVideo, error) {
	var out objects.LiveVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), &out)
	return &out, err
}

type CreateUserMessengerDesktopPerformanceTracesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateUserMessengerDesktopPerformanceTracesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateUserMessengerDesktopPerformanceTraces(ctx context.Context, client *core.Client, id string, params CreateUserMessengerDesktopPerformanceTracesParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "messenger_desktop_performance_traces"), params.ToParams(), &out)
	return &out, err
}

type CreateUserMessengerKidsAccountsUnreadBadgeParams struct {
	ProxiedAppID core.ID     `facebook:"proxied_app_id"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateUserMessengerKidsAccountsUnreadBadgeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["proxied_app_id"] = params.ProxiedAppID
	return out
}

func CreateUserMessengerKidsAccountsUnreadBadge(ctx context.Context, client *core.Client, id string, params CreateUserMessengerKidsAccountsUnreadBadgeParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "messenger_kids_accounts_unread_badge"), params.ToParams(), &out)
	return &out, err
}

type GetUserMusicParams struct {
	TargetID *core.ID    `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetUserMusicParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	return out
}

func GetUserMusic(ctx context.Context, client *core.Client, id string, params GetUserMusicParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "music"), params.ToParams(), &out)
	return &out, err
}

type CreateUserNotificationsParams struct {
	BotMessagePayloadElements *string                                      `facebook:"bot_message_payload_elements"`
	Filtering                 *[]enums.UsernotificationsFilteringEnumParam `facebook:"filtering"`
	GameMetadata              *string                                      `facebook:"game_metadata"`
	Href                      *map[string]interface{}                      `facebook:"href"`
	Label                     *string                                      `facebook:"label"`
	Message                   *map[string]interface{}                      `facebook:"message"`
	NotifIds                  *[]core.ID                                   `facebook:"notif_ids"`
	NotificationContentID     *core.ID                                     `facebook:"notification_content_id"`
	Payload                   *string                                      `facebook:"payload"`
	Read                      *bool                                        `facebook:"read"`
	Ref                       *string                                      `facebook:"ref"`
	ScheduleInterval          *uint64                                      `facebook:"schedule_interval"`
	Seen                      *bool                                        `facebook:"seen"`
	Template                  *map[string]interface{}                      `facebook:"template"`
	Type                      *enums.UsernotificationsTypeEnumParam        `facebook:"type"`
	Extra                     core.Params                                  `facebook:"-"`
}

func (params CreateUserNotificationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BotMessagePayloadElements != nil {
		out["bot_message_payload_elements"] = *params.BotMessagePayloadElements
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GameMetadata != nil {
		out["game_metadata"] = *params.GameMetadata
	}
	if params.Href != nil {
		out["href"] = *params.Href
	}
	if params.Label != nil {
		out["label"] = *params.Label
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NotifIds != nil {
		out["notif_ids"] = *params.NotifIds
	}
	if params.NotificationContentID != nil {
		out["notification_content_id"] = *params.NotificationContentID
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Read != nil {
		out["read"] = *params.Read
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	if params.ScheduleInterval != nil {
		out["schedule_interval"] = *params.ScheduleInterval
	}
	if params.Seen != nil {
		out["seen"] = *params.Seen
	}
	if params.Template != nil {
		out["template"] = *params.Template
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func CreateUserNotifications(ctx context.Context, client *core.Client, id string, params CreateUserNotificationsParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "notifications"), params.ToParams(), &out)
	return &out, err
}

type GetUserPaymentTransactionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPaymentTransactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPaymentTransactions(ctx context.Context, client *core.Client, id string, params GetUserPaymentTransactionsParams) (*core.Cursor[objects.PaymentEnginePayment], error) {
	var out core.Cursor[objects.PaymentEnginePayment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "payment_transactions"), params.ToParams(), &out)
	return &out, err
}

type DeleteUserPermissionsParams struct {
	Permission *string     `facebook:"permission"`
	Extra      core.Params `facebook:"-"`
}

func (params DeleteUserPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Permission != nil {
		out["permission"] = *params.Permission
	}
	return out
}

func DeleteUserPermissions(ctx context.Context, client *core.Client, id string, params DeleteUserPermissionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "permissions"), params.ToParams(), &out)
	return &out, err
}

type GetUserPermissionsParams struct {
	Permission *string                               `facebook:"permission"`
	Status     *enums.UserpermissionsStatusEnumParam `facebook:"status"`
	Extra      core.Params                           `facebook:"-"`
}

func (params GetUserPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Permission != nil {
		out["permission"] = *params.Permission
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetUserPermissions(ctx context.Context, client *core.Client, id string, params GetUserPermissionsParams) (*core.Cursor[objects.Permission], error) {
	var out core.Cursor[objects.Permission]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "permissions"), params.ToParams(), &out)
	return &out, err
}

type GetUserPersonalAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPersonalAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPersonalAdAccounts(ctx context.Context, client *core.Client, id string, params GetUserPersonalAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "personal_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetUserPhotosParams struct {
	Type  *enums.UserphotosTypeEnumParam `facebook:"type"`
	Extra core.Params                    `facebook:"-"`
}

func (params GetUserPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserPhotos(ctx context.Context, client *core.Client, id string, params GetUserPhotosParams) (*core.Cursor[objects.Photo], error) {
	var out core.Cursor[objects.Photo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), &out)
	return &out, err
}

type CreateUserPhotosParams struct {
	Aid                                   *string                                            `facebook:"aid"`
	AllowSphericalPhoto                   *bool                                              `facebook:"allow_spherical_photo"`
	AltTextCustom                         *string                                            `facebook:"alt_text_custom"`
	AndroidKeyHash                        *string                                            `facebook:"android_key_hash"`
	ApplicationID                         *core.ID                                           `facebook:"application_id"`
	Attempt                               *uint64                                            `facebook:"attempt"`
	AudienceExp                           *bool                                              `facebook:"audience_exp"`
	BackdatedTime                         *time.Time                                         `facebook:"backdated_time"`
	BackdatedTimeGranularity              *enums.UserphotosBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	Caption                               *string                                            `facebook:"caption"`
	ComposerSessionID                     *core.ID                                           `facebook:"composer_session_id"`
	DirectShareStatus                     *uint64                                            `facebook:"direct_share_status"`
	FeedTargeting                         *map[string]interface{}                            `facebook:"feed_targeting"`
	FilterType                            *uint64                                            `facebook:"filter_type"`
	FullResIsComingLater                  *bool                                              `facebook:"full_res_is_coming_later"`
	InitialViewHeadingOverrideDegrees     *uint64                                            `facebook:"initial_view_heading_override_degrees"`
	InitialViewPitchOverrideDegrees       *uint64                                            `facebook:"initial_view_pitch_override_degrees"`
	InitialViewVerticalFovOverrideDegrees *uint64                                            `facebook:"initial_view_vertical_fov_override_degrees"`
	IosBundleID                           *core.ID                                           `facebook:"ios_bundle_id"`
	IsExplicitLocation                    *bool                                              `facebook:"is_explicit_location"`
	IsExplicitPlace                       *bool                                              `facebook:"is_explicit_place"`
	ManualPrivacy                         *bool                                              `facebook:"manual_privacy"`
	Message                               *string                                            `facebook:"message"`
	Name                                  *string                                            `facebook:"name"`
	NoStory                               *bool                                              `facebook:"no_story"`
	OfflineID                             *core.ID                                           `facebook:"offline_id"`
	OgActionTypeID                        *core.ID                                           `facebook:"og_action_type_id"`
	OgIconID                              *core.ID                                           `facebook:"og_icon_id"`
	OgObjectID                            *core.ID                                           `facebook:"og_object_id"`
	OgPhrase                              *string                                            `facebook:"og_phrase"`
	OgSetProfileBadge                     *bool                                              `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism                 *string                                            `facebook:"og_suggestion_mechanism"`
	Place                                 *map[string]interface{}                            `facebook:"place"`
	Privacy                               *string                                            `facebook:"privacy"`
	ProfileID                             *core.ID                                           `facebook:"profile_id"`
	ProvenanceInfo                        *map[string]interface{}                            `facebook:"provenance_info"`
	ProxiedAppID                          *core.ID                                           `facebook:"proxied_app_id"`
	Published                             *bool                                              `facebook:"published"`
	Qn                                    *string                                            `facebook:"qn"`
	ScheduledPublishTime                  *uint64                                            `facebook:"scheduled_publish_time"`
	SphericalMetadata                     *map[string]interface{}                            `facebook:"spherical_metadata"`
	SponsorID                             *core.ID                                           `facebook:"sponsor_id"`
	SponsorRelationship                   *uint64                                            `facebook:"sponsor_relationship"`
	Tags                                  *[]map[string]interface{}                          `facebook:"tags"`
	TargetID                              *core.ID                                           `facebook:"target_id"`
	Targeting                             *map[string]interface{}                            `facebook:"targeting"`
	TimeSinceOriginalPost                 *uint64                                            `facebook:"time_since_original_post"`
	UID                                   *core.ID                                           `facebook:"uid"`
	UnpublishedContentType                *enums.UserphotosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	URL                                   *string                                            `facebook:"url"`
	UserSelectedTags                      *bool                                              `facebook:"user_selected_tags"`
	VaultImageID                          *core.ID                                           `facebook:"vault_image_id"`
	Extra                                 core.Params                                        `facebook:"-"`
}

func (params CreateUserPhotosParams) ToParams() core.Params {
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
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
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

func CreateUserPhotos(ctx context.Context, client *core.Client, id string, params CreateUserPhotosParams) (*objects.Photo, error) {
	var out objects.Photo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "photos"), params.ToParams(), &out)
	return &out, err
}

type GetUserPictureParams struct {
	Height   *int                            `facebook:"height"`
	Redirect *bool                           `facebook:"redirect"`
	Type     *enums.UserpictureTypeEnumParam `facebook:"type"`
	Width    *int                            `facebook:"width"`
	Extra    core.Params                     `facebook:"-"`
}

func (params GetUserPictureParams) ToParams() core.Params {
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

func GetUserPicture(ctx context.Context, client *core.Client, id string, params GetUserPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	var out core.Cursor[objects.ProfilePictureSource]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), &out)
	return &out, err
}

type GetUserPostsParams struct {
	IncludeHidden *bool       `facebook:"include_hidden"`
	Q             *string     `facebook:"q"`
	ShowExpired   *bool       `facebook:"show_expired"`
	Since         *time.Time  `facebook:"since"`
	Until         *time.Time  `facebook:"until"`
	With          *string     `facebook:"with"`
	Extra         core.Params `facebook:"-"`
}

func (params GetUserPostsParams) ToParams() core.Params {
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

func GetUserPosts(ctx context.Context, client *core.Client, id string, params GetUserPostsParams) (*core.Cursor[objects.Post], error) {
	var out core.Cursor[objects.Post]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "posts"), params.ToParams(), &out)
	return &out, err
}

type GetUserRichMediaDocumentsParams struct {
	Query *string     `facebook:"query"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserRichMediaDocumentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Query != nil {
		out["query"] = *params.Query
	}
	return out
}

func GetUserRichMediaDocuments(ctx context.Context, client *core.Client, id string, params GetUserRichMediaDocumentsParams) (*core.Cursor[objects.Canvas], error) {
	var out core.Cursor[objects.Canvas]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "rich_media_documents"), params.ToParams(), &out)
	return &out, err
}

type CreateUserStagingResourcesParams struct {
	File  *core.FileParam `facebook:"file"`
	Extra core.Params     `facebook:"-"`
}

func (params CreateUserStagingResourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.File != nil {
		out["file"] = *params.File
	}
	return out
}

func CreateUserStagingResources(ctx context.Context, client *core.Client, id string, params CreateUserStagingResourcesParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "staging_resources"), params.ToParams(), &out)
	return &out, err
}

type GetUserVideosParams struct {
	Type  *enums.UservideosTypeEnumParam `facebook:"type"`
	Extra core.Params                    `facebook:"-"`
}

func (params GetUserVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserVideos(ctx context.Context, client *core.Client, id string, params GetUserVideosParams) (*core.Cursor[objects.AdVideo], error) {
	var out core.Cursor[objects.AdVideo]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), &out)
	return &out, err
}

type CreateUserVideosParams struct {
	ApplicationID                 *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                          `facebook:"audio_story_wave_animation_handle"`
	ComposerEntryPicker           *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                          `facebook:"composer_source_surface"`
	ComposerType                  *string                                          `facebook:"composer_type"`
	ContainerType                 *enums.UservideosContainerTypeEnumParam          `facebook:"container_type"`
	ContentCategory               *enums.UservideosContentCategoryEnumParam        `facebook:"content_category"`
	CreativeTools                 *string                                          `facebook:"creative_tools"`
	Description                   *string                                          `facebook:"description"`
	DirectShareStatus             *uint64                                          `facebook:"direct_share_status"`
	EditDescriptionSpec           *map[string]interface{}                          `facebook:"edit_description_spec"`
	Embeddable                    *bool                                            `facebook:"embeddable"`
	EndOffset                     *uint64                                          `facebook:"end_offset"`
	FbuploaderVideoFileChunk      *string                                          `facebook:"fbuploader_video_file_chunk"`
	FileSize                      *uint64                                          `facebook:"file_size"`
	FileURL                       *string                                          `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                            `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.UservideosFormattingEnumParam             `facebook:"formatting"`
	Fov                           *uint64                                          `facebook:"fov"`
	FrontZRotation                *float64                                         `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                      `facebook:"guide"`
	GuideEnabled                  *bool                                            `facebook:"guide_enabled"`
	InitialHeading                *uint64                                          `facebook:"initial_heading"`
	InitialPitch                  *uint64                                          `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                          `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                            `facebook:"is_boost_intended"`
	IsExplicitShare               *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost            *bool                                            `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                            `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                            `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy                 *bool                                            `facebook:"manual_privacy"`
	NoStory                       *bool                                            `facebook:"no_story"`
	OgActionTypeID                *core.ID                                         `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                      *string                                          `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                          `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                          `facebook:"original_fov"`
	OriginalProjectionType        *enums.UservideosOriginalProjectionTypeEnumParam `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                          `facebook:"partnership_ad_ad_code"`
	Privacy                       *string                                          `facebook:"privacy"`
	PublishEventID                *core.ID                                         `facebook:"publish_event_id"`
	ReferencedStickerID           *core.ID                                         `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                         `facebook:"replace_video_id"`
	SelectedAudioSpec             *map[string]interface{}                          `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                          `facebook:"slideshow_spec"`
	Source                        *string                                          `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                         `facebook:"source_instagram_media_id"`
	Spherical                     *bool                                            `facebook:"spherical"`
	SponsorID                     *core.ID                                         `facebook:"sponsor_id"`
	StartOffset                   *uint64                                          `facebook:"start_offset"`
	SwapMode                      *enums.UservideosSwapModeEnumParam               `facebook:"swap_mode"`
	TextFormatMetadata            *string                                          `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                  `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                          `facebook:"time_since_original_post"`
	Title                         *string                                          `facebook:"title"`
	TranscodeSettingProperties    *string                                          `facebook:"transcode_setting_properties"`
	UnpublishedContentType        *enums.UservideosUnpublishedContentTypeEnumParam `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.UservideosUploadPhaseEnumParam            `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                         `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                          `facebook:"upload_setting_properties"`
	VideoFileChunk                *string                                          `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                          `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                          `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                         `facebook:"waterfall_id"`
	Extra                         core.Params                                      `facebook:"-"`
}

func (params CreateUserVideosParams) ToParams() core.Params {
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
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
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
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
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
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
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
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
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

func CreateUserVideos(ctx context.Context, client *core.Client, id string, params CreateUserVideosParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), &out)
	return &out, err
}

type DeleteUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteUser(ctx context.Context, client *core.Client, id string, params DeleteUserParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUser(ctx context.Context, client *core.Client, id string, params GetUserParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateUserParams struct {
	EmojiColorPref                  *uint64                                    `facebook:"emoji_color_pref"`
	Firstname                       *string                                    `facebook:"firstname"`
	Lastname                        *string                                    `facebook:"lastname"`
	LocalNewsMegaphoneDismissStatus *enums.UserLocalNewsMegaphoneDismissStatus `facebook:"local_news_megaphone_dismiss_status"`
	LocalNewsSubscriptionStatus     *enums.UserLocalNewsSubscriptionStatus     `facebook:"local_news_subscription_status"`
	Name                            *string                                    `facebook:"name"`
	Password                        *string                                    `facebook:"password"`
	Extra                           core.Params                                `facebook:"-"`
}

func (params UpdateUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EmojiColorPref != nil {
		out["emoji_color_pref"] = *params.EmojiColorPref
	}
	if params.Firstname != nil {
		out["firstname"] = *params.Firstname
	}
	if params.Lastname != nil {
		out["lastname"] = *params.Lastname
	}
	if params.LocalNewsMegaphoneDismissStatus != nil {
		out["local_news_megaphone_dismiss_status"] = *params.LocalNewsMegaphoneDismissStatus
	}
	if params.LocalNewsSubscriptionStatus != nil {
		out["local_news_subscription_status"] = *params.LocalNewsSubscriptionStatus
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Password != nil {
		out["password"] = *params.Password
	}
	return out
}

func UpdateUser(ctx context.Context, client *core.Client, id string, params UpdateUserParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
