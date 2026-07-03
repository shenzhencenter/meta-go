package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
)

type GetWhatsAppBusinessAccountActivitiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountActivitiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountActivities(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountActivitiesParams) (*core.Cursor[objects.WhatsAppBusinessActivityHistory], error) {
	var out core.Cursor[objects.WhatsAppBusinessActivityHistory]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "activities"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountAssignedUsersParams struct {
	Tasks []enums.WhatsappbusinessaccountassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                                        `facebook:"user"`
	Extra core.Params                                                `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tasks"] = params.Tasks
	out["user"] = params.User
	return out
}

func CreateWhatsAppBusinessAccountAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountAssignedUsersParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountAudiencesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountAudiencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountAudiences(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountAudiencesParams) (*core.Cursor[objects.WhatsAppBusinessAudience], error) {
	var out core.Cursor[objects.WhatsAppBusinessAudience]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "audiences"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams struct {
	Features []map[string]interface{} `facebook:"features"`
	Extra    core.Params              `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["features"] = params.Features
	return out
}

func CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountBusinessMessagingFeatureStatusParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountCallAnalyticsParams struct {
	CountryCodes *[]string                                                         `facebook:"country_codes"`
	Dimensions   *[]enums.WhatsappbusinessaccountcallAnalyticsDimensionsEnumParam  `facebook:"dimensions"`
	Directions   *[]enums.WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam  `facebook:"directions"`
	End          uint64                                                            `facebook:"end"`
	Granularity  enums.WhatsappbusinessaccountcallAnalyticsGranularityEnumParam    `facebook:"granularity"`
	MetricTypes  *[]enums.WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam `facebook:"metric_types"`
	PhoneNumbers *[]string                                                         `facebook:"phone_numbers"`
	Start        uint64                                                            `facebook:"start"`
	Tiers        *[]string                                                         `facebook:"tiers"`
	Extra        core.Params                                                       `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountCallAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	if params.Directions != nil {
		out["directions"] = *params.Directions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	out["start"] = params.Start
	if params.Tiers != nil {
		out["tiers"] = *params.Tiers
	}
	return out
}

func GetWhatsAppBusinessAccountCallAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountCallAnalyticsParams) (*core.Cursor[objects.WABACallAnalytics], error) {
	var out core.Cursor[objects.WABACallAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "call_analytics"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountConversationAnalyticsParams struct {
	ConversationCategories *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam `facebook:"conversation_categories"`
	ConversationDirections *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationDirectionsEnumParam `facebook:"conversation_directions"`
	ConversationTypes      *[]enums.WhatsappbusinessaccountconversationAnalyticsConversationTypesEnumParam      `facebook:"conversation_types"`
	CountryCodes           *[]string                                                                            `facebook:"country_codes"`
	Dimensions             *[]enums.WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam             `facebook:"dimensions"`
	End                    uint64                                                                               `facebook:"end"`
	Granularity            enums.WhatsappbusinessaccountconversationAnalyticsGranularityEnumParam               `facebook:"granularity"`
	MetricTypes            *[]enums.WhatsappbusinessaccountconversationAnalyticsMetricTypesEnumParam            `facebook:"metric_types"`
	PhoneNumbers           *[]string                                                                            `facebook:"phone_numbers"`
	Start                  uint64                                                                               `facebook:"start"`
	Extra                  core.Params                                                                          `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountConversationAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ConversationCategories != nil {
		out["conversation_categories"] = *params.ConversationCategories
	}
	if params.ConversationDirections != nil {
		out["conversation_directions"] = *params.ConversationDirections
	}
	if params.ConversationTypes != nil {
		out["conversation_types"] = *params.ConversationTypes
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	out["start"] = params.Start
	return out
}

func GetWhatsAppBusinessAccountConversationAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountConversationAnalyticsParams) (*core.Cursor[objects.WABAConversationAnalytics], error) {
	var out core.Cursor[objects.WABAConversationAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "conversation_analytics"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountDataset(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountDatasetParams) (*core.Cursor[objects.Dataset], error) {
	var out core.Cursor[objects.Dataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "dataset"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountDatasetParams struct {
	DatasetName *string     `facebook:"dataset_name"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatasetName != nil {
		out["dataset_name"] = *params.DatasetName
	}
	return out
}

func CreateWhatsAppBusinessAccountDataset(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountDatasetParams) (*objects.Dataset, error) {
	var out objects.Dataset
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "dataset"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountDegreesOfFreedomSpec(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountDegreesOfFreedomSpecParams) (*core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn], error) {
	var out core.Cursor[objects.MarketingMessagesCreativeOptimizationsOptIn]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "degrees_of_freedom_spec"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountFlowsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountFlows(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountFlowsParams) (*core.Cursor[objects.WhatsAppExtension], error) {
	var out core.Cursor[objects.WhatsAppExtension]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "flows"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountFlowsParams struct {
	Categories  []enums.WhatsappbusinessaccountflowsCategoriesEnumParam `facebook:"categories"`
	CloneFlowID *core.ID                                                `facebook:"clone_flow_id"`
	EndpointURI *string                                                 `facebook:"endpoint_uri"`
	FlowJSON    *string                                                 `facebook:"flow_json"`
	Name        string                                                  `facebook:"name"`
	Publish     *bool                                                   `facebook:"publish"`
	Extra       core.Params                                             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["categories"] = params.Categories
	if params.CloneFlowID != nil {
		out["clone_flow_id"] = *params.CloneFlowID
	}
	if params.EndpointURI != nil {
		out["endpoint_uri"] = *params.EndpointURI
	}
	if params.FlowJSON != nil {
		out["flow_json"] = *params.FlowJSON
	}
	out["name"] = params.Name
	if params.Publish != nil {
		out["publish"] = *params.Publish
	}
	return out
}

func CreateWhatsAppBusinessAccountFlows(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountFlowsParams) (*objects.WhatsAppExtension, error) {
	var out objects.WhatsAppExtension
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "flows"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	RedirectURL       *string     `facebook:"redirect_url"`
	Extra             core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	if params.RedirectURL != nil {
		out["redirect_url"] = *params.RedirectURL
	}
	return out
}

func CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLink(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountGeneratePaymentConfigurationOauthLinkParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "generate_payment_configuration_oauth_link"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountGroupAnalyticsParams struct {
	End         time.Time                                                          `facebook:"end"`
	Granularity *enums.WhatsappbusinessaccountgroupAnalyticsGranularityEnumParam   `facebook:"granularity"`
	GroupIds    []core.ID                                                          `facebook:"group_ids"`
	MetricTypes *[]enums.WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam `facebook:"metric_types"`
	Start       time.Time                                                          `facebook:"start"`
	Extra       core.Params                                                        `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountGroupAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	if params.Granularity != nil {
		out["granularity"] = *params.Granularity
	}
	out["group_ids"] = params.GroupIds
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	out["start"] = params.Start
	return out
}

func GetWhatsAppBusinessAccountGroupAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountGroupAnalyticsParams) (*core.Cursor[objects.WABAGroupAnalytics], error) {
	var out core.Cursor[objects.WABAGroupAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "group_analytics"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountMessageCampaigns(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageCampaignsParams) (*core.Cursor[objects.WhatsAppBusinessCampaign], error) {
	var out core.Cursor[objects.WhatsAppBusinessCampaign]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "message_campaigns"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMessageSamplesParams struct {
	Interactive *map[string]interface{}                                  `facebook:"interactive"`
	Text        *map[string]interface{}                                  `facebook:"text"`
	Type        enums.WhatsappbusinessaccountmessageSamplesTypeEnumParam `facebook:"type"`
	Extra       core.Params                                              `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMessageSamplesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Interactive != nil {
		out["interactive"] = *params.Interactive
	}
	if params.Text != nil {
		out["text"] = *params.Text
	}
	out["type"] = params.Type
	return out
}

func CreateWhatsAppBusinessAccountMessageSamples(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMessageSamplesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "message_samples"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageTemplatePreviewsParams struct {
	AddSecurityRecommendation *bool                                                                       `facebook:"add_security_recommendation"`
	BusinessName              *string                                                                     `facebook:"business_name"`
	ButtonTypes               *[]enums.WhatsappbusinessaccountmessageTemplatePreviewsButtonTypesEnumParam `facebook:"button_types"`
	Category                  enums.WhatsappbusinessaccountmessageTemplatePreviewsCategoryEnumParam       `facebook:"category"`
	CodeExpirationMinutes     *uint64                                                                     `facebook:"code_expiration_minutes"`
	Languages                 *[]string                                                                   `facebook:"languages"`
	Extra                     core.Params                                                                 `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AddSecurityRecommendation != nil {
		out["add_security_recommendation"] = *params.AddSecurityRecommendation
	}
	if params.BusinessName != nil {
		out["business_name"] = *params.BusinessName
	}
	if params.ButtonTypes != nil {
		out["button_types"] = *params.ButtonTypes
	}
	out["category"] = params.Category
	if params.CodeExpirationMinutes != nil {
		out["code_expiration_minutes"] = *params.CodeExpirationMinutes
	}
	if params.Languages != nil {
		out["languages"] = *params.Languages
	}
	return out
}

func GetWhatsAppBusinessAccountMessageTemplatePreviews(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageTemplatePreviewsParams) (*core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountMessageTemplatePreview]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "message_template_previews"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountMessageTemplatesParams struct {
	HsmID  *core.ID    `facebook:"hsm_id"`
	HsmIds *[]core.ID  `facebook:"hsm_ids"`
	Name   *string     `facebook:"name"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.HsmID != nil {
		out["hsm_id"] = *params.HsmID
	}
	if params.HsmIds != nil {
		out["hsm_ids"] = *params.HsmIds
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func DeleteWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountMessageTemplatesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "message_templates"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountMessageTemplatesParams struct {
	Category      *[]enums.WhatsappbusinessaccountmessageTemplatesCategoryEnumParam     `facebook:"category"`
	Content       *string                                                               `facebook:"content"`
	Language      *[]string                                                             `facebook:"language"`
	Name          *string                                                               `facebook:"name"`
	NameOrContent *string                                                               `facebook:"name_or_content"`
	QualityScore  *[]enums.WhatsappbusinessaccountmessageTemplatesQualityScoreEnumParam `facebook:"quality_score"`
	Since         *time.Time                                                            `facebook:"since"`
	Source        *enums.WhatsappbusinessaccountmessageTemplatesSourceEnumParam         `facebook:"source"`
	Status        *[]enums.WhatsappbusinessaccountmessageTemplatesStatusEnumParam       `facebook:"status"`
	Until         *time.Time                                                            `facebook:"until"`
	Extra         core.Params                                                           `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.Content != nil {
		out["content"] = *params.Content
	}
	if params.Language != nil {
		out["language"] = *params.Language
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NameOrContent != nil {
		out["name_or_content"] = *params.NameOrContent
	}
	if params.QualityScore != nil {
		out["quality_score"] = *params.QualityScore
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountMessageTemplatesParams) (*core.Cursor[objects.WhatsAppBusinessHSM], error) {
	var out core.Cursor[objects.WhatsAppBusinessHSM]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "message_templates"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMessageTemplatesParams struct {
	AllowCategoryChange         *bool                                                                  `facebook:"allow_category_change"`
	BidSpec                     *map[string]interface{}                                                `facebook:"bid_spec"`
	Category                    enums.WhatsappbusinessaccountmessageTemplatesCategoryEnumParam         `facebook:"category"`
	Components                  *[]map[string]interface{}                                              `facebook:"components"`
	CreativeSourcingSpec        *map[string]interface{}                                                `facebook:"creative_sourcing_spec"`
	CtaURLLinkTrackingOptedOut  *bool                                                                  `facebook:"cta_url_link_tracking_opted_out"`
	DegreesOfFreedomSpec        *map[string]interface{}                                                `facebook:"degrees_of_freedom_spec"`
	DisplayFormat               *enums.WhatsappbusinessaccountmessageTemplatesDisplayFormatEnumParam   `facebook:"display_format"`
	IsPrimaryDeviceDeliveryOnly *bool                                                                  `facebook:"is_primary_device_delivery_only"`
	Language                    string                                                                 `facebook:"language"`
	LibraryTemplateBodyInputs   *map[string]interface{}                                                `facebook:"library_template_body_inputs"`
	LibraryTemplateButtonInputs *[]map[string]interface{}                                              `facebook:"library_template_button_inputs"`
	LibraryTemplateName         *string                                                                `facebook:"library_template_name"`
	MessageSendTTLSeconds       *uint64                                                                `facebook:"message_send_ttl_seconds"`
	Name                        string                                                                 `facebook:"name"`
	OptimizationSpec            *map[string]interface{}                                                `facebook:"optimization_spec"`
	ParameterFormat             *enums.WhatsappbusinessaccountmessageTemplatesParameterFormatEnumParam `facebook:"parameter_format"`
	ProductSetID                *core.ID                                                               `facebook:"product_set_id"`
	SendType                    *enums.WhatsappbusinessaccountmessageTemplatesSendTypeEnumParam        `facebook:"send_type"`
	SubCategory                 *enums.WhatsappbusinessaccountmessageTemplatesSubCategoryEnumParam     `facebook:"sub_category"`
	Extra                       core.Params                                                            `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowCategoryChange != nil {
		out["allow_category_change"] = *params.AllowCategoryChange
	}
	if params.BidSpec != nil {
		out["bid_spec"] = *params.BidSpec
	}
	out["category"] = params.Category
	if params.Components != nil {
		out["components"] = *params.Components
	}
	if params.CreativeSourcingSpec != nil {
		out["creative_sourcing_spec"] = *params.CreativeSourcingSpec
	}
	if params.CtaURLLinkTrackingOptedOut != nil {
		out["cta_url_link_tracking_opted_out"] = *params.CtaURLLinkTrackingOptedOut
	}
	if params.DegreesOfFreedomSpec != nil {
		out["degrees_of_freedom_spec"] = *params.DegreesOfFreedomSpec
	}
	if params.DisplayFormat != nil {
		out["display_format"] = *params.DisplayFormat
	}
	if params.IsPrimaryDeviceDeliveryOnly != nil {
		out["is_primary_device_delivery_only"] = *params.IsPrimaryDeviceDeliveryOnly
	}
	out["language"] = params.Language
	if params.LibraryTemplateBodyInputs != nil {
		out["library_template_body_inputs"] = *params.LibraryTemplateBodyInputs
	}
	if params.LibraryTemplateButtonInputs != nil {
		out["library_template_button_inputs"] = *params.LibraryTemplateButtonInputs
	}
	if params.LibraryTemplateName != nil {
		out["library_template_name"] = *params.LibraryTemplateName
	}
	if params.MessageSendTTLSeconds != nil {
		out["message_send_ttl_seconds"] = *params.MessageSendTTLSeconds
	}
	out["name"] = params.Name
	if params.OptimizationSpec != nil {
		out["optimization_spec"] = *params.OptimizationSpec
	}
	if params.ParameterFormat != nil {
		out["parameter_format"] = *params.ParameterFormat
	}
	if params.ProductSetID != nil {
		out["product_set_id"] = *params.ProductSetID
	}
	if params.SendType != nil {
		out["send_type"] = *params.SendType
	}
	if params.SubCategory != nil {
		out["sub_category"] = *params.SubCategory
	}
	return out
}

func CreateWhatsAppBusinessAccountMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "message_templates"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMigrateFlowsParams struct {
	SourceFlowNames *[]string   `facebook:"source_flow_names"`
	SourceWabaID    core.ID     `facebook:"source_waba_id"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMigrateFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SourceFlowNames != nil {
		out["source_flow_names"] = *params.SourceFlowNames
	}
	out["source_waba_id"] = params.SourceWabaID
	return out
}

func CreateWhatsAppBusinessAccountMigrateFlows(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMigrateFlowsParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "migrate_flows"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams struct {
	Count        *uint64     `facebook:"count"`
	PageNumber   *uint64     `facebook:"page_number"`
	SourceWabaID core.ID     `facebook:"source_waba_id"`
	TemplateIds  *[]core.ID  `facebook:"template_ids"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Count != nil {
		out["count"] = *params.Count
	}
	if params.PageNumber != nil {
		out["page_number"] = *params.PageNumber
	}
	out["source_waba_id"] = params.SourceWabaID
	if params.TemplateIds != nil {
		out["template_ids"] = *params.TemplateIds
	}
	return out
}

func CreateWhatsAppBusinessAccountMigrateMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountMigrateMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "migrate_message_templates"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	Extra             core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	return out
}

func DeleteWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountPaymentConfigurationParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "payment_configuration"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName string      `facebook:"configuration_name"`
	Extra             core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	return out
}

func GetWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPaymentConfigurationParams) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfiguration]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "payment_configuration"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountPaymentConfigurationParams struct {
	ConfigurationName    string                                                                  `facebook:"configuration_name"`
	DataEndpointURL      *string                                                                 `facebook:"data_endpoint_url"`
	MerchantCategoryCode *string                                                                 `facebook:"merchant_category_code"`
	MerchantVpa          *string                                                                 `facebook:"merchant_vpa"`
	ProviderName         *enums.WhatsappbusinessaccountpaymentConfigurationProviderNameEnumParam `facebook:"provider_name"`
	PurposeCode          *string                                                                 `facebook:"purpose_code"`
	RedirectURL          *string                                                                 `facebook:"redirect_url"`
	Extra                core.Params                                                             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountPaymentConfigurationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["configuration_name"] = params.ConfigurationName
	if params.DataEndpointURL != nil {
		out["data_endpoint_url"] = *params.DataEndpointURL
	}
	if params.MerchantCategoryCode != nil {
		out["merchant_category_code"] = *params.MerchantCategoryCode
	}
	if params.MerchantVpa != nil {
		out["merchant_vpa"] = *params.MerchantVpa
	}
	if params.ProviderName != nil {
		out["provider_name"] = *params.ProviderName
	}
	if params.PurposeCode != nil {
		out["purpose_code"] = *params.PurposeCode
	}
	if params.RedirectURL != nil {
		out["redirect_url"] = *params.RedirectURL
	}
	return out
}

func CreateWhatsAppBusinessAccountPaymentConfiguration(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountPaymentConfigurationParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "payment_configuration"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPaymentConfigurationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPaymentConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountPaymentConfigurations(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPaymentConfigurationsParams) (*core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountPaymentConfigurations]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "payment_configurations"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPhoneNumbersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPhoneNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountPhoneNumbers(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPhoneNumbersParams) (*core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccountToNumberCurrentStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "phone_numbers"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountPhoneNumbersParams struct {
	Cc                 *string     `facebook:"cc"`
	MigratePhoneNumber *bool       `facebook:"migrate_phone_number"`
	PhoneNumber        *string     `facebook:"phone_number"`
	PreverifiedID      *core.ID    `facebook:"preverified_id"`
	VerifiedName       *string     `facebook:"verified_name"`
	Extra              core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountPhoneNumbersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cc != nil {
		out["cc"] = *params.Cc
	}
	if params.MigratePhoneNumber != nil {
		out["migrate_phone_number"] = *params.MigratePhoneNumber
	}
	if params.PhoneNumber != nil {
		out["phone_number"] = *params.PhoneNumber
	}
	if params.PreverifiedID != nil {
		out["preverified_id"] = *params.PreverifiedID
	}
	if params.VerifiedName != nil {
		out["verified_name"] = *params.VerifiedName
	}
	return out
}

func CreateWhatsAppBusinessAccountPhoneNumbers(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountPhoneNumbersParams) (*objects.WhatsAppBusinessAccountToNumberCurrentStatus, error) {
	var out objects.WhatsAppBusinessAccountToNumberCurrentStatus
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "phone_numbers"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountPricingAnalyticsParams struct {
	CountryCodes      *[]string                                                                  `facebook:"country_codes"`
	Dimensions        *[]enums.WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam        `facebook:"dimensions"`
	End               uint64                                                                     `facebook:"end"`
	Granularity       enums.WhatsappbusinessaccountpricingAnalyticsGranularityEnumParam          `facebook:"granularity"`
	MetricTypes       *[]enums.WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam       `facebook:"metric_types"`
	PhoneNumbers      *[]string                                                                  `facebook:"phone_numbers"`
	PricingCategories *[]enums.WhatsappbusinessaccountpricingAnalyticsPricingCategoriesEnumParam `facebook:"pricing_categories"`
	PricingTypes      *[]enums.WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam      `facebook:"pricing_types"`
	Start             uint64                                                                     `facebook:"start"`
	Tiers             *[]string                                                                  `facebook:"tiers"`
	Extra             core.Params                                                                `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountPricingAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CountryCodes != nil {
		out["country_codes"] = *params.CountryCodes
	}
	if params.Dimensions != nil {
		out["dimensions"] = *params.Dimensions
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.PhoneNumbers != nil {
		out["phone_numbers"] = *params.PhoneNumbers
	}
	if params.PricingCategories != nil {
		out["pricing_categories"] = *params.PricingCategories
	}
	if params.PricingTypes != nil {
		out["pricing_types"] = *params.PricingTypes
	}
	out["start"] = params.Start
	if params.Tiers != nil {
		out["tiers"] = *params.Tiers
	}
	return out
}

func GetWhatsAppBusinessAccountPricingAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountPricingAnalyticsParams) (*core.Cursor[objects.WABAPricingAnalytics], error) {
	var out core.Cursor[objects.WABAPricingAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pricing_analytics"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountProductCatalogsParams struct {
	CatalogID core.ID     `facebook:"catalog_id"`
	Extra     core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["catalog_id"] = params.CatalogID
	return out
}

func DeleteWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountProductCatalogsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountProductCatalogsParams struct {
	CatalogID core.ID     `facebook:"catalog_id"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["catalog_id"] = params.CatalogID
	return out
}

func CreateWhatsAppBusinessAccountProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountProductCatalogsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSchedulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSchedules(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSchedulesParams) (*core.Cursor[objects.WhatsAppBusinessCampaignSchedule], error) {
	var out core.Cursor[objects.WhatsAppBusinessCampaignSchedule]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "schedules"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams struct {
	AppID      *core.ID    `facebook:"app_id"`
	SolutionID *core.ID    `facebook:"solution_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.SolutionID != nil {
		out["solution_id"] = *params.SolutionID
	}
	return out
}

func CreateWhatsAppBusinessAccountSetSolutionMigrationIntent(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountSetSolutionMigrationIntentParams) (*objects.WhatsAppBusinessAccountMigrationIntent, error) {
	var out objects.WhatsAppBusinessAccountMigrationIntent
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "set_solution_migration_intent"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSolutionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSolutionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSolutions(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSolutionsParams) (*core.Cursor[objects.WhatsAppBusinessSolution], error) {
	var out core.Cursor[objects.WhatsAppBusinessSolution]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "solutions"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountSubscribedAppsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountSubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountSubscribedAppsParams) (*core.Cursor[objects.WhatsAppApplication], error) {
	var out core.Cursor[objects.WhatsAppApplication]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountSubscribedAppsParams struct {
	OverrideCallbackURI *string     `facebook:"override_callback_uri"`
	VerifyToken         *string     `facebook:"verify_token"`
	Extra               core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountSubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.OverrideCallbackURI != nil {
		out["override_callback_uri"] = *params.OverrideCallbackURI
	}
	if params.VerifyToken != nil {
		out["verify_token"] = *params.VerifyToken
	}
	return out
}

func CreateWhatsAppBusinessAccountSubscribedApps(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountSubscribedAppsParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateAnalyticsParams struct {
	End             time.Time                                                           `facebook:"end"`
	Granularity     enums.WhatsappbusinessaccounttemplateAnalyticsGranularityEnumParam  `facebook:"granularity"`
	MetricTypes     *[]string                                                           `facebook:"metric_types"`
	ProductType     *enums.WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam `facebook:"product_type"`
	Start           time.Time                                                           `facebook:"start"`
	TemplateIds     []core.ID                                                           `facebook:"template_ids"`
	UseWabaTimezone *bool                                                               `facebook:"use_waba_timezone"`
	Extra           core.Params                                                         `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	out["start"] = params.Start
	out["template_ids"] = params.TemplateIds
	if params.UseWabaTimezone != nil {
		out["use_waba_timezone"] = *params.UseWabaTimezone
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateAnalyticsParams) (*core.Cursor[objects.WABATemplateAnalytics], error) {
	var out core.Cursor[objects.WABATemplateAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "template_analytics"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams struct {
	End              time.Time                                                               `facebook:"end"`
	Granularity      enums.WhatsappbusinessaccounttemplateGroupAnalyticsGranularityEnumParam `facebook:"granularity"`
	MetricTypes      *[]string                                                               `facebook:"metric_types"`
	Start            time.Time                                                               `facebook:"start"`
	TemplateGroupIds []core.ID                                                               `facebook:"template_group_ids"`
	UseWabaTimezone  *bool                                                                   `facebook:"use_waba_timezone"`
	Extra            core.Params                                                             `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["end"] = params.End
	out["granularity"] = params.Granularity
	if params.MetricTypes != nil {
		out["metric_types"] = *params.MetricTypes
	}
	out["start"] = params.Start
	out["template_group_ids"] = params.TemplateGroupIds
	if params.UseWabaTimezone != nil {
		out["use_waba_timezone"] = *params.UseWabaTimezone
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateGroupAnalytics(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateGroupAnalyticsParams) (*core.Cursor[objects.WABATemplateGroupAnalytics], error) {
	var out core.Cursor[objects.WABATemplateGroupAnalytics]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "template_group_analytics"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountTemplateGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountTemplateGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccountTemplateGroups(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountTemplateGroupsParams) (*core.Cursor[objects.BusinessMessagingTemplateGroup], error) {
	var out core.Cursor[objects.BusinessMessagingTemplateGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "template_groups"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountTemplateGroupsParams struct {
	Description               string      `facebook:"description"`
	Name                      string      `facebook:"name"`
	WhatsappBusinessTemplates []string    `facebook:"whatsapp_business_templates"`
	Extra                     core.Params `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountTemplateGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["description"] = params.Description
	out["name"] = params.Name
	out["whatsapp_business_templates"] = params.WhatsappBusinessTemplates
	return out
}

func CreateWhatsAppBusinessAccountTemplateGroups(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountTemplateGroupsParams) (*objects.BusinessMessagingTemplateGroup, error) {
	var out objects.BusinessMessagingTemplateGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "template_groups"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams struct {
	Category              enums.WhatsappbusinessaccountupsertMessageTemplatesCategoryEnumParam `facebook:"category"`
	Components            []map[string]interface{}                                             `facebook:"components"`
	Languages             []string                                                             `facebook:"languages"`
	MessageSendTTLSeconds *uint64                                                              `facebook:"message_send_ttl_seconds"`
	Name                  string                                                               `facebook:"name"`
	Extra                 core.Params                                                          `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["category"] = params.Category
	out["components"] = params.Components
	out["languages"] = params.Languages
	if params.MessageSendTTLSeconds != nil {
		out["message_send_ttl_seconds"] = *params.MessageSendTTLSeconds
	}
	out["name"] = params.Name
	return out
}

func CreateWhatsAppBusinessAccountUpsertMessageTemplates(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountUpsertMessageTemplatesParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "upsert_message_templates"), params.ToParams(), &out)
	return &out, err
}

type DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	SequenceID core.ID     `facebook:"sequence_id"`
	Extra      core.Params `facebook:"-"`
}

func (params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["sequence_id"] = params.SequenceID
	return out
}

func DeleteWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params DeleteWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	AppID      *core.ID    `facebook:"app_id"`
	SequenceID *core.ID    `facebook:"sequence_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.SequenceID != nil {
		out["sequence_id"] = *params.SequenceID
	}
	return out
}

func GetWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), &out)
	return &out, err
}

type CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams struct {
	Name                   *string                 `facebook:"name"`
	SequenceID             *core.ID                `facebook:"sequence_id"`
	WelcomeMessageSequence *map[string]interface{} `facebook:"welcome_message_sequence"`
	Extra                  core.Params             `facebook:"-"`
}

func (params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SequenceID != nil {
		out["sequence_id"] = *params.SequenceID
	}
	if params.WelcomeMessageSequence != nil {
		out["welcome_message_sequence"] = *params.WelcomeMessageSequence
	}
	return out
}

func CreateWhatsAppBusinessAccountWelcomeMessageSequences(ctx context.Context, client *core.Client, id string, params CreateWhatsAppBusinessAccountWelcomeMessageSequencesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "welcome_message_sequences"), params.ToParams(), &out)
	return &out, err
}

type GetWhatsAppBusinessAccountParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessAccount(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessAccountParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateWhatsAppBusinessAccountParams struct {
	CreativeOptimizationsEnrollment    *map[string]interface{} `facebook:"creative_optimizations_enrollment"`
	DegreesOfFreedomSpec               *map[string]interface{} `facebook:"degrees_of_freedom_spec"`
	DisableMarketingMessagesOnCloudAPI *bool                   `facebook:"disable_marketing_messages_on_cloud_api"`
	IsEnabledForInsights               *bool                   `facebook:"is_enabled_for_insights"`
	TemplateAutoArchivalEnabled        *bool                   `facebook:"template_auto_archival_enabled"`
	Extra                              core.Params             `facebook:"-"`
}

func (params UpdateWhatsAppBusinessAccountParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CreativeOptimizationsEnrollment != nil {
		out["creative_optimizations_enrollment"] = *params.CreativeOptimizationsEnrollment
	}
	if params.DegreesOfFreedomSpec != nil {
		out["degrees_of_freedom_spec"] = *params.DegreesOfFreedomSpec
	}
	if params.DisableMarketingMessagesOnCloudAPI != nil {
		out["disable_marketing_messages_on_cloud_api"] = *params.DisableMarketingMessagesOnCloudAPI
	}
	if params.IsEnabledForInsights != nil {
		out["is_enabled_for_insights"] = *params.IsEnabledForInsights
	}
	if params.TemplateAutoArchivalEnabled != nil {
		out["template_auto_archival_enabled"] = *params.TemplateAutoArchivalEnabled
	}
	return out
}

func UpdateWhatsAppBusinessAccount(ctx context.Context, client *core.Client, id string, params UpdateWhatsAppBusinessAccountParams) (*objects.WhatsAppBusinessAccount, error) {
	var out objects.WhatsAppBusinessAccount
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
