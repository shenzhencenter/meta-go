package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type DeleteProductCatalogAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteProductCatalogAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteProductCatalogAgencies(ctx context.Context, client *core.Client, id string, params DeleteProductCatalogAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogAgencies(ctx context.Context, client *core.Client, id string, params GetProductCatalogAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogAgenciesParams struct {
	Business           string                                                     `facebook:"business"`
	EnabledCollabTerms *[]enums.ProductcatalogagenciesEnabledCollabTermsEnumParam `facebook:"enabled_collab_terms"`
	PermittedRoles     *[]enums.ProductcatalogagenciesPermittedRolesEnumParam     `facebook:"permitted_roles"`
	PermittedTasks     *[]enums.ProductcatalogagenciesPermittedTasksEnumParam     `facebook:"permitted_tasks"`
	SkipDefaults       *bool                                                      `facebook:"skip_defaults"`
	UtmSettings        *map[string]interface{}                                    `facebook:"utm_settings"`
	Extra              core.Params                                                `facebook:"-"`
}

func (params CreateProductCatalogAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	if params.EnabledCollabTerms != nil {
		out["enabled_collab_terms"] = *params.EnabledCollabTerms
	}
	if params.PermittedRoles != nil {
		out["permitted_roles"] = *params.PermittedRoles
	}
	if params.PermittedTasks != nil {
		out["permitted_tasks"] = *params.PermittedTasks
	}
	if params.SkipDefaults != nil {
		out["skip_defaults"] = *params.SkipDefaults
	}
	if params.UtmSettings != nil {
		out["utm_settings"] = *params.UtmSettings
	}
	return out
}

func CreateProductCatalogAgencies(ctx context.Context, client *core.Client, id string, params CreateProductCatalogAgenciesParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type DeleteProductCatalogAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteProductCatalogAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteProductCatalogAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteProductCatalogAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetProductCatalogAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetProductCatalogAssignedUsers(ctx context.Context, client *core.Client, id string, params GetProductCatalogAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogAssignedUsersParams struct {
	Tasks []enums.ProductcatalogassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                               `facebook:"user"`
	Extra core.Params                                       `facebook:"-"`
}

func (params CreateProductCatalogAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tasks"] = params.Tasks
	out["user"] = params.User
	return out
}

func CreateProductCatalogAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateProductCatalogAssignedUsersParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogAutomotiveModelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogAutomotiveModelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogAutomotiveModels(ctx context.Context, client *core.Client, id string, params GetProductCatalogAutomotiveModelsParams) (*core.Cursor[objects.AutomotiveModel], error) {
	var out core.Cursor[objects.AutomotiveModel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "automotive_models"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogBatchParams struct {
	AllowUpsert           *bool                    `facebook:"allow_upsert"`
	FbeExternalBusinessID *core.ID                 `facebook:"fbe_external_business_id"`
	Requests              []map[string]interface{} `facebook:"requests"`
	Version               *uint64                  `facebook:"version"`
	Extra                 core.Params              `facebook:"-"`
}

func (params CreateProductCatalogBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	if params.FbeExternalBusinessID != nil {
		out["fbe_external_business_id"] = *params.FbeExternalBusinessID
	}
	out["requests"] = params.Requests
	if params.Version != nil {
		out["version"] = *params.Version
	}
	return out
}

func CreateProductCatalogBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "batch"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogCatalogStoreParams struct {
	Page  string      `facebook:"page"`
	Extra core.Params `facebook:"-"`
}

func (params CreateProductCatalogCatalogStoreParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["page"] = params.Page
	return out
}

func CreateProductCatalogCatalogStore(ctx context.Context, client *core.Client, id string, params CreateProductCatalogCatalogStoreParams) (*objects.StoreCatalogSettings, error) {
	var out objects.StoreCatalogSettings
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "catalog_store"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCategoriesParams struct {
	CategorizationCriteria enums.ProductcatalogcategoriesCategorizationCriteriaEnumParam `facebook:"categorization_criteria"`
	Filter                 *map[string]interface{}                                       `facebook:"filter"`
	Extra                  core.Params                                                   `facebook:"-"`
}

func (params GetProductCatalogCategoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["categorization_criteria"] = params.CategorizationCriteria
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogCategories(ctx context.Context, client *core.Client, id string, params GetProductCatalogCategoriesParams) (*core.Cursor[objects.ProductCatalogCategory], error) {
	var out core.Cursor[objects.ProductCatalogCategory]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "categories"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogCategoriesParams struct {
	Data  []map[string]interface{} `facebook:"data"`
	Extra core.Params              `facebook:"-"`
}

func (params CreateProductCatalogCategoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["data"] = params.Data
	return out
}

func CreateProductCatalogCategories(ctx context.Context, client *core.Client, id string, params CreateProductCatalogCategoriesParams) (*objects.ProductCatalogCategory, error) {
	var out objects.ProductCatalogCategory
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "categories"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCheckBatchRequestStatusParams struct {
	ErrorPriority            *enums.ProductcatalogcheckBatchRequestStatusErrorPriorityEnumParam `facebook:"error_priority"`
	Handle                   string                                                             `facebook:"handle"`
	LoadIdsOfInvalidRequests *bool                                                              `facebook:"load_ids_of_invalid_requests"`
	Extra                    core.Params                                                        `facebook:"-"`
}

func (params GetProductCatalogCheckBatchRequestStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ErrorPriority != nil {
		out["error_priority"] = *params.ErrorPriority
	}
	out["handle"] = params.Handle
	if params.LoadIdsOfInvalidRequests != nil {
		out["load_ids_of_invalid_requests"] = *params.LoadIdsOfInvalidRequests
	}
	return out
}

func GetProductCatalogCheckBatchRequestStatus(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckBatchRequestStatusParams) (*core.Cursor[objects.CheckBatchRequestStatus], error) {
	var out core.Cursor[objects.CheckBatchRequestStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "check_batch_request_status"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCheckMarketplacePartnerDealsStatusParams struct {
	SessionID core.ID     `facebook:"session_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetProductCatalogCheckMarketplacePartnerDealsStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["session_id"] = params.SessionID
	return out
}

func GetProductCatalogCheckMarketplacePartnerDealsStatus(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckMarketplacePartnerDealsStatusParams) (*core.Cursor[objects.ProductCatalogCheckMarketplacePartnerDealsStatus], error) {
	var out core.Cursor[objects.ProductCatalogCheckMarketplacePartnerDealsStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "check_marketplace_partner_deals_status"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCheckMarketplacePartnerSellersStatusParams struct {
	SessionID core.ID     `facebook:"session_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetProductCatalogCheckMarketplacePartnerSellersStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["session_id"] = params.SessionID
	return out
}

func GetProductCatalogCheckMarketplacePartnerSellersStatus(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckMarketplacePartnerSellersStatusParams) (*core.Cursor[objects.ProductCatalogCheckMarketplacePartnerSellersStatus], error) {
	var out core.Cursor[objects.ProductCatalogCheckMarketplacePartnerSellersStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "check_marketplace_partner_sellers_status"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCollaborativeAdsLsbImageBankParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogCollaborativeAdsLsbImageBankParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogCollaborativeAdsLsbImageBank(ctx context.Context, client *core.Client, id string, params GetProductCatalogCollaborativeAdsLsbImageBankParams) (*core.Cursor[objects.CPASLsbImageBank], error) {
	var out core.Cursor[objects.CPASLsbImageBank]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborative_ads_lsb_image_bank"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCollaborativeAdsShareSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogCollaborativeAdsShareSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogCollaborativeAdsShareSettings(ctx context.Context, client *core.Client, id string, params GetProductCatalogCollaborativeAdsShareSettingsParams) (*core.Cursor[objects.CollaborativeAdsShareSettings], error) {
	var out core.Cursor[objects.CollaborativeAdsShareSettings]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborative_ads_share_settings"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogCpasLsbImageBankParams struct {
	AdGroupID        *core.ID    `facebook:"ad_group_id"`
	AgencyBusinessID *core.ID    `facebook:"agency_business_id"`
	BackupImageUrls  []string    `facebook:"backup_image_urls"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateProductCatalogCpasLsbImageBankParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdGroupID != nil {
		out["ad_group_id"] = *params.AdGroupID
	}
	if params.AgencyBusinessID != nil {
		out["agency_business_id"] = *params.AgencyBusinessID
	}
	out["backup_image_urls"] = params.BackupImageUrls
	return out
}

func CreateProductCatalogCpasLsbImageBank(ctx context.Context, client *core.Client, id string, params CreateProductCatalogCpasLsbImageBankParams) (*objects.CPASLsbImageBank, error) {
	var out objects.CPASLsbImageBank
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "cpas_lsb_image_bank"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogCreatorAssetCreativesParams struct {
	ModerationStatus *enums.ProductcatalogcreatorAssetCreativesModerationStatusEnumParam `facebook:"moderation_status"`
	Extra            core.Params                                                         `facebook:"-"`
}

func (params GetProductCatalogCreatorAssetCreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ModerationStatus != nil {
		out["moderation_status"] = *params.ModerationStatus
	}
	return out
}

func GetProductCatalogCreatorAssetCreatives(ctx context.Context, client *core.Client, id string, params GetProductCatalogCreatorAssetCreativesParams) (*core.Cursor[objects.CreatorAssetCreative], error) {
	var out core.Cursor[objects.CreatorAssetCreative]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "creator_asset_creatives"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogDataSourcesParams struct {
	IngestionSourceType *enums.ProductcatalogdataSourcesIngestionSourceTypeEnumParam `facebook:"ingestion_source_type"`
	Extra               core.Params                                                  `facebook:"-"`
}

func (params GetProductCatalogDataSourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IngestionSourceType != nil {
		out["ingestion_source_type"] = *params.IngestionSourceType
	}
	return out
}

func GetProductCatalogDataSources(ctx context.Context, client *core.Client, id string, params GetProductCatalogDataSourcesParams) (*core.Cursor[objects.ProductCatalogDataSource], error) {
	var out core.Cursor[objects.ProductCatalogDataSource]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "data_sources"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogDestinationsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogDestinationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogDestinations(ctx context.Context, client *core.Client, id string, params GetProductCatalogDestinationsParams) (*core.Cursor[objects.Destination], error) {
	var out core.Cursor[objects.Destination]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "destinations"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogDiagnosticsParams struct {
	AffectedChannels *[]enums.ProductcatalogdiagnosticsAffectedChannelsEnumParam `facebook:"affected_channels"`
	AffectedEntities *[]enums.ProductcatalogdiagnosticsAffectedEntitiesEnumParam `facebook:"affected_entities"`
	AffectedFeatures *[]enums.ProductcatalogdiagnosticsAffectedFeaturesEnumParam `facebook:"affected_features"`
	Severities       *[]enums.ProductcatalogdiagnosticsSeveritiesEnumParam       `facebook:"severities"`
	Types            *[]enums.ProductcatalogdiagnosticsTypesEnumParam            `facebook:"types"`
	Extra            core.Params                                                 `facebook:"-"`
}

func (params GetProductCatalogDiagnosticsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AffectedChannels != nil {
		out["affected_channels"] = *params.AffectedChannels
	}
	if params.AffectedEntities != nil {
		out["affected_entities"] = *params.AffectedEntities
	}
	if params.AffectedFeatures != nil {
		out["affected_features"] = *params.AffectedFeatures
	}
	if params.Severities != nil {
		out["severities"] = *params.Severities
	}
	if params.Types != nil {
		out["types"] = *params.Types
	}
	return out
}

func GetProductCatalogDiagnostics(ctx context.Context, client *core.Client, id string, params GetProductCatalogDiagnosticsParams) (*core.Cursor[objects.ProductCatalogDiagnosticGroup], error) {
	var out core.Cursor[objects.ProductCatalogDiagnosticGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "diagnostics"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogEventStatsParams struct {
	Breakdowns *[]enums.ProductcatalogeventStatsBreakdownsEnumParam `facebook:"breakdowns"`
	Extra      core.Params                                          `facebook:"-"`
}

func (params GetProductCatalogEventStatsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	return out
}

func GetProductCatalogEventStats(ctx context.Context, client *core.Client, id string, params GetProductCatalogEventStatsParams) (*core.Cursor[objects.ProductEventStat], error) {
	var out core.Cursor[objects.ProductEventStat]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "event_stats"), params.ToParams(), &out)
	return &out, err
}

type DeleteProductCatalogExternalEventSourcesParams struct {
	ExternalEventSources *map[string]interface{} `facebook:"external_event_sources"`
	Extra                core.Params             `facebook:"-"`
}

func (params DeleteProductCatalogExternalEventSourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ExternalEventSources != nil {
		out["external_event_sources"] = *params.ExternalEventSources
	}
	return out
}

func DeleteProductCatalogExternalEventSources(ctx context.Context, client *core.Client, id string, params DeleteProductCatalogExternalEventSourcesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "external_event_sources"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogExternalEventSourcesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogExternalEventSourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogExternalEventSources(ctx context.Context, client *core.Client, id string, params GetProductCatalogExternalEventSourcesParams) (*core.Cursor[objects.ExternalEventSource], error) {
	var out core.Cursor[objects.ExternalEventSource]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "external_event_sources"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogExternalEventSourcesParams struct {
	ExternalEventSources *map[string]interface{} `facebook:"external_event_sources"`
	Extra                core.Params             `facebook:"-"`
}

func (params CreateProductCatalogExternalEventSourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ExternalEventSources != nil {
		out["external_event_sources"] = *params.ExternalEventSources
	}
	return out
}

func CreateProductCatalogExternalEventSources(ctx context.Context, client *core.Client, id string, params CreateProductCatalogExternalEventSourcesParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "external_event_sources"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogFlightsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogFlightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogFlights(ctx context.Context, client *core.Client, id string, params GetProductCatalogFlightsParams) (*core.Cursor[objects.Flight], error) {
	var out core.Cursor[objects.Flight]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "flights"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogGeolocatedItemsBatchParams struct {
	AllowUpsert *bool                  `facebook:"allow_upsert"`
	ItemType    string                 `facebook:"item_type"`
	Requests    map[string]interface{} `facebook:"requests"`
	Extra       core.Params            `facebook:"-"`
}

func (params CreateProductCatalogGeolocatedItemsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	out["item_type"] = params.ItemType
	out["requests"] = params.Requests
	return out
}

func CreateProductCatalogGeolocatedItemsBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogGeolocatedItemsBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "geolocated_items_batch"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogHomeListingsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogHomeListingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogHomeListings(ctx context.Context, client *core.Client, id string, params GetProductCatalogHomeListingsParams) (*core.Cursor[objects.HomeListing], error) {
	var out core.Cursor[objects.HomeListing]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "home_listings"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogHomeListingsParams struct {
	Address       map[string]interface{}   `facebook:"address"`
	Availability  string                   `facebook:"availability"`
	Currency      string                   `facebook:"currency"`
	Description   *string                  `facebook:"description"`
	HomeListingID core.ID                  `facebook:"home_listing_id"`
	Images        []map[string]interface{} `facebook:"images"`
	ListingType   *string                  `facebook:"listing_type"`
	Name          string                   `facebook:"name"`
	NumBaths      *float64                 `facebook:"num_baths"`
	NumBeds       *float64                 `facebook:"num_beds"`
	NumUnits      *float64                 `facebook:"num_units"`
	Price         float64                  `facebook:"price"`
	PropertyType  *string                  `facebook:"property_type"`
	URL           string                   `facebook:"url"`
	YearBuilt     uint64                   `facebook:"year_built"`
	Extra         core.Params              `facebook:"-"`
}

func (params CreateProductCatalogHomeListingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["address"] = params.Address
	out["availability"] = params.Availability
	out["currency"] = params.Currency
	if params.Description != nil {
		out["description"] = *params.Description
	}
	out["home_listing_id"] = params.HomeListingID
	out["images"] = params.Images
	if params.ListingType != nil {
		out["listing_type"] = *params.ListingType
	}
	out["name"] = params.Name
	if params.NumBaths != nil {
		out["num_baths"] = *params.NumBaths
	}
	if params.NumBeds != nil {
		out["num_beds"] = *params.NumBeds
	}
	if params.NumUnits != nil {
		out["num_units"] = *params.NumUnits
	}
	out["price"] = params.Price
	if params.PropertyType != nil {
		out["property_type"] = *params.PropertyType
	}
	out["url"] = params.URL
	out["year_built"] = params.YearBuilt
	return out
}

func CreateProductCatalogHomeListings(ctx context.Context, client *core.Client, id string, params CreateProductCatalogHomeListingsParams) (*objects.HomeListing, error) {
	var out objects.HomeListing
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "home_listings"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogHotelRoomsBatchParams struct {
	Handle string      `facebook:"handle"`
	Extra  core.Params `facebook:"-"`
}

func (params GetProductCatalogHotelRoomsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["handle"] = params.Handle
	return out
}

func GetProductCatalogHotelRoomsBatch(ctx context.Context, client *core.Client, id string, params GetProductCatalogHotelRoomsBatchParams) (*core.Cursor[objects.ProductCatalogHotelRoomsBatch], error) {
	var out core.Cursor[objects.ProductCatalogHotelRoomsBatch]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "hotel_rooms_batch"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogHotelRoomsBatchParams struct {
	File       *core.FileParam                                      `facebook:"file"`
	Password   *string                                              `facebook:"password"`
	Standard   enums.ProductcataloghotelRoomsBatchStandardEnumParam `facebook:"standard"`
	UpdateOnly *bool                                                `facebook:"update_only"`
	URL        *string                                              `facebook:"url"`
	Username   *string                                              `facebook:"username"`
	Extra      core.Params                                          `facebook:"-"`
}

func (params CreateProductCatalogHotelRoomsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.File != nil {
		out["file"] = *params.File
	}
	if params.Password != nil {
		out["password"] = *params.Password
	}
	out["standard"] = params.Standard
	if params.UpdateOnly != nil {
		out["update_only"] = *params.UpdateOnly
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.Username != nil {
		out["username"] = *params.Username
	}
	return out
}

func CreateProductCatalogHotelRoomsBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogHotelRoomsBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "hotel_rooms_batch"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogHotelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogHotelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogHotels(ctx context.Context, client *core.Client, id string, params GetProductCatalogHotelsParams) (*core.Cursor[objects.Hotel], error) {
	var out core.Cursor[objects.Hotel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "hotels"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogHotelsParams struct {
	Address      map[string]interface{}    `facebook:"address"`
	Applinks     *map[string]interface{}   `facebook:"applinks"`
	BasePrice    *uint64                   `facebook:"base_price"`
	Brand        *string                   `facebook:"brand"`
	Currency     *string                   `facebook:"currency"`
	Description  string                    `facebook:"description"`
	GuestRatings *[]map[string]interface{} `facebook:"guest_ratings"`
	HotelID      *core.ID                  `facebook:"hotel_id"`
	Images       []map[string]interface{}  `facebook:"images"`
	Name         string                    `facebook:"name"`
	Phone        *string                   `facebook:"phone"`
	StarRating   *float64                  `facebook:"star_rating"`
	URL          string                    `facebook:"url"`
	Extra        core.Params               `facebook:"-"`
}

func (params CreateProductCatalogHotelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["address"] = params.Address
	if params.Applinks != nil {
		out["applinks"] = *params.Applinks
	}
	if params.BasePrice != nil {
		out["base_price"] = *params.BasePrice
	}
	if params.Brand != nil {
		out["brand"] = *params.Brand
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	out["description"] = params.Description
	if params.GuestRatings != nil {
		out["guest_ratings"] = *params.GuestRatings
	}
	if params.HotelID != nil {
		out["hotel_id"] = *params.HotelID
	}
	out["images"] = params.Images
	out["name"] = params.Name
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.StarRating != nil {
		out["star_rating"] = *params.StarRating
	}
	out["url"] = params.URL
	return out
}

func CreateProductCatalogHotels(ctx context.Context, client *core.Client, id string, params CreateProductCatalogHotelsParams) (*objects.Hotel, error) {
	var out objects.Hotel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "hotels"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogItemsBatchParams struct {
	AllowUpsert *bool                                               `facebook:"allow_upsert"`
	ItemSubType *enums.ProductcatalogitemsBatchItemSubTypeEnumParam `facebook:"item_sub_type"`
	ItemType    string                                              `facebook:"item_type"`
	Requests    map[string]interface{}                              `facebook:"requests"`
	Version     *uint64                                             `facebook:"version"`
	Extra       core.Params                                         `facebook:"-"`
}

func (params CreateProductCatalogItemsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	if params.ItemSubType != nil {
		out["item_sub_type"] = *params.ItemSubType
	}
	out["item_type"] = params.ItemType
	out["requests"] = params.Requests
	if params.Version != nil {
		out["version"] = *params.Version
	}
	return out
}

func CreateProductCatalogItemsBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogItemsBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "items_batch"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogLocalizedItemsBatchParams struct {
	AllowUpsert *bool                  `facebook:"allow_upsert"`
	ItemType    string                 `facebook:"item_type"`
	Requests    map[string]interface{} `facebook:"requests"`
	Version     *uint64                `facebook:"version"`
	Extra       core.Params            `facebook:"-"`
}

func (params CreateProductCatalogLocalizedItemsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	out["item_type"] = params.ItemType
	out["requests"] = params.Requests
	if params.Version != nil {
		out["version"] = *params.Version
	}
	return out
}

func CreateProductCatalogLocalizedItemsBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogLocalizedItemsBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "localized_items_batch"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogMarketplacePartnerDealsDetailsParams struct {
	Requests map[string]interface{} `facebook:"requests"`
	Extra    core.Params            `facebook:"-"`
}

func (params CreateProductCatalogMarketplacePartnerDealsDetailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["requests"] = params.Requests
	return out
}

func CreateProductCatalogMarketplacePartnerDealsDetails(ctx context.Context, client *core.Client, id string, params CreateProductCatalogMarketplacePartnerDealsDetailsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "marketplace_partner_deals_details"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogMarketplacePartnerSellersDetailsParams struct {
	Requests map[string]interface{} `facebook:"requests"`
	Extra    core.Params            `facebook:"-"`
}

func (params CreateProductCatalogMarketplacePartnerSellersDetailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["requests"] = params.Requests
	return out
}

func CreateProductCatalogMarketplacePartnerSellersDetails(ctx context.Context, client *core.Client, id string, params CreateProductCatalogMarketplacePartnerSellersDetailsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "marketplace_partner_sellers_details"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogMarketplacePartnerSignalsParams struct {
	ConversionType *enums.ProductcatalogmarketplacePartnerSignalsConversionTypeEnumParam `facebook:"conversion_type"`
	EventID        *core.ID                                                              `facebook:"event_id"`
	EventName      enums.ProductcatalogmarketplacePartnerSignalsEventNameEnumParam       `facebook:"event_name"`
	EventSourceURL *string                                                               `facebook:"event_source_url"`
	EventTime      time.Time                                                             `facebook:"event_time"`
	OfferData      *map[string]interface{}                                               `facebook:"offer_data"`
	OrderData      *map[string]interface{}                                               `facebook:"order_data"`
	UserData       map[string]interface{}                                                `facebook:"user_data"`
	Extra          core.Params                                                           `facebook:"-"`
}

func (params CreateProductCatalogMarketplacePartnerSignalsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ConversionType != nil {
		out["conversion_type"] = *params.ConversionType
	}
	if params.EventID != nil {
		out["event_id"] = *params.EventID
	}
	out["event_name"] = params.EventName
	if params.EventSourceURL != nil {
		out["event_source_url"] = *params.EventSourceURL
	}
	out["event_time"] = params.EventTime
	if params.OfferData != nil {
		out["offer_data"] = *params.OfferData
	}
	if params.OrderData != nil {
		out["order_data"] = *params.OrderData
	}
	out["user_data"] = params.UserData
	return out
}

func CreateProductCatalogMarketplacePartnerSignals(ctx context.Context, client *core.Client, id string, params CreateProductCatalogMarketplacePartnerSignalsParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "marketplace_partner_signals"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogMediaTitlesParams struct {
	AdditionalImageUrls *[]string   `facebook:"additional_image_urls"`
	AndroidAppName      *string     `facebook:"android_app_name"`
	AndroidClass        *string     `facebook:"android_class"`
	AndroidPackage      *string     `facebook:"android_package"`
	AndroidURL          *string     `facebook:"android_url"`
	Awards              *[]string   `facebook:"awards"`
	Cast                *[]string   `facebook:"cast"`
	Category            *string     `facebook:"category"`
	Currency            *string     `facebook:"currency"`
	Description         *string     `facebook:"description"`
	Director            *[]string   `facebook:"director"`
	FbProductCategory   *string     `facebook:"fb_product_category"`
	Genre               *[]string   `facebook:"genre"`
	ImageURL            string      `facebook:"image_url"`
	IosAppName          *string     `facebook:"ios_app_name"`
	IosAppStoreID       *core.ID    `facebook:"ios_app_store_id"`
	IosURL              *string     `facebook:"ios_url"`
	IpadAppName         *string     `facebook:"ipad_app_name"`
	IpadAppStoreID      *core.ID    `facebook:"ipad_app_store_id"`
	IpadURL             *string     `facebook:"ipad_url"`
	IphoneAppName       *string     `facebook:"iphone_app_name"`
	IphoneAppStoreID    *core.ID    `facebook:"iphone_app_store_id"`
	IphoneURL           *string     `facebook:"iphone_url"`
	MediaCategory       *string     `facebook:"media_category"`
	Name                string      `facebook:"name"`
	Price               *uint64     `facebook:"price"`
	Rating              *string     `facebook:"rating"`
	ReleaseDate         *string     `facebook:"release_date"`
	RetailerID          core.ID     `facebook:"retailer_id"`
	URL                 *string     `facebook:"url"`
	WindowsPhoneAppID   *core.ID    `facebook:"windows_phone_app_id"`
	WindowsPhoneAppName *string     `facebook:"windows_phone_app_name"`
	WindowsPhoneURL     *string     `facebook:"windows_phone_url"`
	Extra               core.Params `facebook:"-"`
}

func (params CreateProductCatalogMediaTitlesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdditionalImageUrls != nil {
		out["additional_image_urls"] = *params.AdditionalImageUrls
	}
	if params.AndroidAppName != nil {
		out["android_app_name"] = *params.AndroidAppName
	}
	if params.AndroidClass != nil {
		out["android_class"] = *params.AndroidClass
	}
	if params.AndroidPackage != nil {
		out["android_package"] = *params.AndroidPackage
	}
	if params.AndroidURL != nil {
		out["android_url"] = *params.AndroidURL
	}
	if params.Awards != nil {
		out["awards"] = *params.Awards
	}
	if params.Cast != nil {
		out["cast"] = *params.Cast
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Director != nil {
		out["director"] = *params.Director
	}
	if params.FbProductCategory != nil {
		out["fb_product_category"] = *params.FbProductCategory
	}
	if params.Genre != nil {
		out["genre"] = *params.Genre
	}
	out["image_url"] = params.ImageURL
	if params.IosAppName != nil {
		out["ios_app_name"] = *params.IosAppName
	}
	if params.IosAppStoreID != nil {
		out["ios_app_store_id"] = *params.IosAppStoreID
	}
	if params.IosURL != nil {
		out["ios_url"] = *params.IosURL
	}
	if params.IpadAppName != nil {
		out["ipad_app_name"] = *params.IpadAppName
	}
	if params.IpadAppStoreID != nil {
		out["ipad_app_store_id"] = *params.IpadAppStoreID
	}
	if params.IpadURL != nil {
		out["ipad_url"] = *params.IpadURL
	}
	if params.IphoneAppName != nil {
		out["iphone_app_name"] = *params.IphoneAppName
	}
	if params.IphoneAppStoreID != nil {
		out["iphone_app_store_id"] = *params.IphoneAppStoreID
	}
	if params.IphoneURL != nil {
		out["iphone_url"] = *params.IphoneURL
	}
	if params.MediaCategory != nil {
		out["media_category"] = *params.MediaCategory
	}
	out["name"] = params.Name
	if params.Price != nil {
		out["price"] = *params.Price
	}
	if params.Rating != nil {
		out["rating"] = *params.Rating
	}
	if params.ReleaseDate != nil {
		out["release_date"] = *params.ReleaseDate
	}
	out["retailer_id"] = params.RetailerID
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.WindowsPhoneAppID != nil {
		out["windows_phone_app_id"] = *params.WindowsPhoneAppID
	}
	if params.WindowsPhoneAppName != nil {
		out["windows_phone_app_name"] = *params.WindowsPhoneAppName
	}
	if params.WindowsPhoneURL != nil {
		out["windows_phone_url"] = *params.WindowsPhoneURL
	}
	return out
}

func CreateProductCatalogMediaTitles(ctx context.Context, client *core.Client, id string, params CreateProductCatalogMediaTitlesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "media_titles"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogPricingVariablesBatchParams struct {
	Handle string      `facebook:"handle"`
	Extra  core.Params `facebook:"-"`
}

func (params GetProductCatalogPricingVariablesBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["handle"] = params.Handle
	return out
}

func GetProductCatalogPricingVariablesBatch(ctx context.Context, client *core.Client, id string, params GetProductCatalogPricingVariablesBatchParams) (*core.Cursor[objects.ProductCatalogPricingVariablesBatch], error) {
	var out core.Cursor[objects.ProductCatalogPricingVariablesBatch]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "pricing_variables_batch"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogPricingVariablesBatchParams struct {
	File       *core.FileParam                                            `facebook:"file"`
	Password   *string                                                    `facebook:"password"`
	Standard   enums.ProductcatalogpricingVariablesBatchStandardEnumParam `facebook:"standard"`
	UpdateOnly *bool                                                      `facebook:"update_only"`
	URL        *string                                                    `facebook:"url"`
	Username   *string                                                    `facebook:"username"`
	Extra      core.Params                                                `facebook:"-"`
}

func (params CreateProductCatalogPricingVariablesBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.File != nil {
		out["file"] = *params.File
	}
	if params.Password != nil {
		out["password"] = *params.Password
	}
	out["standard"] = params.Standard
	if params.UpdateOnly != nil {
		out["update_only"] = *params.UpdateOnly
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.Username != nil {
		out["username"] = *params.Username
	}
	return out
}

func CreateProductCatalogPricingVariablesBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogPricingVariablesBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "pricing_variables_batch"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogProductFeedsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogProductFeedsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogProductFeeds(ctx context.Context, client *core.Client, id string, params GetProductCatalogProductFeedsParams) (*core.Cursor[objects.ProductFeed], error) {
	var out core.Cursor[objects.ProductFeed]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_feeds"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogProductFeedsParams struct {
	Country                *string                                                       `facebook:"country"`
	DefaultCurrency        *string                                                       `facebook:"default_currency"`
	DeletionEnabled        *bool                                                         `facebook:"deletion_enabled"`
	Delimiter              *enums.ProductcatalogproductFeedsDelimiterEnumParam           `facebook:"delimiter"`
	Encoding               *enums.ProductcatalogproductFeedsEncodingEnumParam            `facebook:"encoding"`
	FeedType               *enums.ProductcatalogproductFeedsFeedTypeEnumParam            `facebook:"feed_type"`
	FileName               *string                                                       `facebook:"file_name"`
	IngestionSourceType    *enums.ProductcatalogproductFeedsIngestionSourceTypeEnumParam `facebook:"ingestion_source_type"`
	ItemSubType            *enums.ProductcatalogproductFeedsItemSubTypeEnumParam         `facebook:"item_sub_type"`
	MigratedFromFeedID     *core.ID                                                      `facebook:"migrated_from_feed_id"`
	Name                   *string                                                       `facebook:"name"`
	OverrideType           *enums.ProductcatalogproductFeedsOverrideTypeEnumParam        `facebook:"override_type"`
	OverrideValue          *string                                                       `facebook:"override_value"`
	PrimaryFeedIds         *[]core.ID                                                    `facebook:"primary_feed_ids"`
	QuotedFieldsMode       *enums.ProductcatalogproductFeedsQuotedFieldsModeEnumParam    `facebook:"quoted_fields_mode"`
	Rules                  *[]string                                                     `facebook:"rules"`
	Schedule               *string                                                       `facebook:"schedule"`
	SelectedOverrideFields *[]string                                                     `facebook:"selected_override_fields"`
	UpdateSchedule         *string                                                       `facebook:"update_schedule"`
	UseCase                *enums.ProductcatalogproductFeedsUseCaseEnumParam             `facebook:"use_case"`
	Extra                  core.Params                                                   `facebook:"-"`
}

func (params CreateProductCatalogProductFeedsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Country != nil {
		out["country"] = *params.Country
	}
	if params.DefaultCurrency != nil {
		out["default_currency"] = *params.DefaultCurrency
	}
	if params.DeletionEnabled != nil {
		out["deletion_enabled"] = *params.DeletionEnabled
	}
	if params.Delimiter != nil {
		out["delimiter"] = *params.Delimiter
	}
	if params.Encoding != nil {
		out["encoding"] = *params.Encoding
	}
	if params.FeedType != nil {
		out["feed_type"] = *params.FeedType
	}
	if params.FileName != nil {
		out["file_name"] = *params.FileName
	}
	if params.IngestionSourceType != nil {
		out["ingestion_source_type"] = *params.IngestionSourceType
	}
	if params.ItemSubType != nil {
		out["item_sub_type"] = *params.ItemSubType
	}
	if params.MigratedFromFeedID != nil {
		out["migrated_from_feed_id"] = *params.MigratedFromFeedID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OverrideType != nil {
		out["override_type"] = *params.OverrideType
	}
	if params.OverrideValue != nil {
		out["override_value"] = *params.OverrideValue
	}
	if params.PrimaryFeedIds != nil {
		out["primary_feed_ids"] = *params.PrimaryFeedIds
	}
	if params.QuotedFieldsMode != nil {
		out["quoted_fields_mode"] = *params.QuotedFieldsMode
	}
	if params.Rules != nil {
		out["rules"] = *params.Rules
	}
	if params.Schedule != nil {
		out["schedule"] = *params.Schedule
	}
	if params.SelectedOverrideFields != nil {
		out["selected_override_fields"] = *params.SelectedOverrideFields
	}
	if params.UpdateSchedule != nil {
		out["update_schedule"] = *params.UpdateSchedule
	}
	if params.UseCase != nil {
		out["use_case"] = *params.UseCase
	}
	return out
}

func CreateProductCatalogProductFeeds(ctx context.Context, client *core.Client, id string, params CreateProductCatalogProductFeedsParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_feeds"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogProductGroupsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogProductGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogProductGroups(ctx context.Context, client *core.Client, id string, params GetProductCatalogProductGroupsParams) (*core.Cursor[objects.ProductGroup], error) {
	var out core.Cursor[objects.ProductGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_groups"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogProductGroupsParams struct {
	RetailerID *core.ID                  `facebook:"retailer_id"`
	Variants   *[]map[string]interface{} `facebook:"variants"`
	Extra      core.Params               `facebook:"-"`
}

func (params CreateProductCatalogProductGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.RetailerID != nil {
		out["retailer_id"] = *params.RetailerID
	}
	if params.Variants != nil {
		out["variants"] = *params.Variants
	}
	return out
}

func CreateProductCatalogProductGroups(ctx context.Context, client *core.Client, id string, params CreateProductCatalogProductGroupsParams) (*objects.ProductGroup, error) {
	var out objects.ProductGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_groups"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogProductSetsParams struct {
	AncestorID  *core.ID    `facebook:"ancestor_id"`
	HasChildren *bool       `facebook:"has_children"`
	ParentID    *core.ID    `facebook:"parent_id"`
	RetailerID  *core.ID    `facebook:"retailer_id"`
	Extra       core.Params `facebook:"-"`
}

func (params GetProductCatalogProductSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AncestorID != nil {
		out["ancestor_id"] = *params.AncestorID
	}
	if params.HasChildren != nil {
		out["has_children"] = *params.HasChildren
	}
	if params.ParentID != nil {
		out["parent_id"] = *params.ParentID
	}
	if params.RetailerID != nil {
		out["retailer_id"] = *params.RetailerID
	}
	return out
}

func GetProductCatalogProductSets(ctx context.Context, client *core.Client, id string, params GetProductCatalogProductSetsParams) (*core.Cursor[objects.ProductSet], error) {
	var out core.Cursor[objects.ProductSet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_sets"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogProductSetsParams struct {
	Filter         *map[string]interface{}   `facebook:"filter"`
	Metadata       *map[string]interface{}   `facebook:"metadata"`
	Name           string                    `facebook:"name"`
	OrderingInfo   *[]uint64                 `facebook:"ordering_info"`
	PublishToShops *[]map[string]interface{} `facebook:"publish_to_shops"`
	RetailerID     *core.ID                  `facebook:"retailer_id"`
	Extra          core.Params               `facebook:"-"`
}

func (params CreateProductCatalogProductSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	out["name"] = params.Name
	if params.OrderingInfo != nil {
		out["ordering_info"] = *params.OrderingInfo
	}
	if params.PublishToShops != nil {
		out["publish_to_shops"] = *params.PublishToShops
	}
	if params.RetailerID != nil {
		out["retailer_id"] = *params.RetailerID
	}
	return out
}

func CreateProductCatalogProductSets(ctx context.Context, client *core.Client, id string, params CreateProductCatalogProductSetsParams) (*objects.ProductSet, error) {
	var out objects.ProductSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_sets"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogProductSetsBatchParams struct {
	Handle string      `facebook:"handle"`
	Extra  core.Params `facebook:"-"`
}

func (params GetProductCatalogProductSetsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["handle"] = params.Handle
	return out
}

func GetProductCatalogProductSetsBatch(ctx context.Context, client *core.Client, id string, params GetProductCatalogProductSetsBatchParams) (*core.Cursor[objects.ProductCatalogProductSetsBatch], error) {
	var out core.Cursor[objects.ProductCatalogProductSetsBatch]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_sets_batch"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogProductsParams struct {
	BulkPagination             *bool                                               `facebook:"bulk_pagination"`
	ErrorPriority              *enums.ProductcatalogproductsErrorPriorityEnumParam `facebook:"error_priority"`
	ErrorType                  *enums.ProductcatalogproductsErrorTypeEnumParam     `facebook:"error_type"`
	Filter                     *map[string]interface{}                             `facebook:"filter"`
	ReturnOnlyApprovedProducts *bool                                               `facebook:"return_only_approved_products"`
	Extra                      core.Params                                         `facebook:"-"`
}

func (params GetProductCatalogProductsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.ErrorPriority != nil {
		out["error_priority"] = *params.ErrorPriority
	}
	if params.ErrorType != nil {
		out["error_type"] = *params.ErrorType
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.ReturnOnlyApprovedProducts != nil {
		out["return_only_approved_products"] = *params.ReturnOnlyApprovedProducts
	}
	return out
}

func GetProductCatalogProducts(ctx context.Context, client *core.Client, id string, params GetProductCatalogProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogProductsParams struct {
	AdditionalImageUrls         *[]string                                                    `facebook:"additional_image_urls"`
	AdditionalVariantAttributes *map[string]interface{}                                      `facebook:"additional_variant_attributes"`
	AgeGroup                    *enums.ProductcatalogproductsAgeGroupEnumParam               `facebook:"age_group"`
	AllowUpsert                 *bool                                                        `facebook:"allow_upsert"`
	AndroidAppName              *string                                                      `facebook:"android_app_name"`
	AndroidClass                *string                                                      `facebook:"android_class"`
	AndroidPackage              *string                                                      `facebook:"android_package"`
	AndroidURL                  *string                                                      `facebook:"android_url"`
	Availability                *enums.ProductcatalogproductsAvailabilityEnumParam           `facebook:"availability"`
	Brand                       *string                                                      `facebook:"brand"`
	Category                    *string                                                      `facebook:"category"`
	CategorySpecificFields      *map[string]interface{}                                      `facebook:"category_specific_fields"`
	CheckoutURL                 *string                                                      `facebook:"checkout_url"`
	Color                       *string                                                      `facebook:"color"`
	CommerceTaxCategory         *enums.ProductcatalogproductsCommerceTaxCategoryEnumParam    `facebook:"commerce_tax_category"`
	Condition                   *enums.ProductcatalogproductsConditionEnumParam              `facebook:"condition"`
	Currency                    string                                                       `facebook:"currency"`
	CustomData                  *map[string]interface{}                                      `facebook:"custom_data"`
	CustomLabelX0               *string                                                      `facebook:"custom_label_0"`
	CustomLabelX1               *string                                                      `facebook:"custom_label_1"`
	CustomLabelX2               *string                                                      `facebook:"custom_label_2"`
	CustomLabelX3               *string                                                      `facebook:"custom_label_3"`
	CustomLabelX4               *string                                                      `facebook:"custom_label_4"`
	CustomNumberX0              *uint64                                                      `facebook:"custom_number_0"`
	CustomNumberX1              *uint64                                                      `facebook:"custom_number_1"`
	CustomNumberX2              *uint64                                                      `facebook:"custom_number_2"`
	CustomNumberX3              *uint64                                                      `facebook:"custom_number_3"`
	CustomNumberX4              *uint64                                                      `facebook:"custom_number_4"`
	Description                 *string                                                      `facebook:"description"`
	ExpirationDate              *string                                                      `facebook:"expiration_date"`
	FbProductCategory           *string                                                      `facebook:"fb_product_category"`
	Gender                      *enums.ProductcatalogproductsGenderEnumParam                 `facebook:"gender"`
	Gtin                        *string                                                      `facebook:"gtin"`
	ImageURL                    *string                                                      `facebook:"image_url"`
	ImporterAddress             *map[string]interface{}                                      `facebook:"importer_address"`
	ImporterName                *string                                                      `facebook:"importer_name"`
	Inventory                   *uint64                                                      `facebook:"inventory"`
	IosAppName                  *string                                                      `facebook:"ios_app_name"`
	IosAppStoreID               *core.ID                                                     `facebook:"ios_app_store_id"`
	IosURL                      *string                                                      `facebook:"ios_url"`
	IpadAppName                 *string                                                      `facebook:"ipad_app_name"`
	IpadAppStoreID              *core.ID                                                     `facebook:"ipad_app_store_id"`
	IpadURL                     *string                                                      `facebook:"ipad_url"`
	IphoneAppName               *string                                                      `facebook:"iphone_app_name"`
	IphoneAppStoreID            *core.ID                                                     `facebook:"iphone_app_store_id"`
	IphoneURL                   *string                                                      `facebook:"iphone_url"`
	LaunchDate                  *string                                                      `facebook:"launch_date"`
	LiveSpecialPrice            *string                                                      `facebook:"live_special_price"`
	ManufacturerInfo            *string                                                      `facebook:"manufacturer_info"`
	ManufacturerPartNumber      *string                                                      `facebook:"manufacturer_part_number"`
	MarkedForProductLaunch      *enums.ProductcatalogproductsMarkedForProductLaunchEnumParam `facebook:"marked_for_product_launch"`
	Material                    *string                                                      `facebook:"material"`
	MobileLink                  *string                                                      `facebook:"mobile_link"`
	Name                        string                                                       `facebook:"name"`
	OrderingIndex               *uint64                                                      `facebook:"ordering_index"`
	OriginCountry               *enums.ProductcatalogproductsOriginCountryEnumParam          `facebook:"origin_country"`
	Pattern                     *string                                                      `facebook:"pattern"`
	Price                       uint64                                                       `facebook:"price"`
	ProductPriorityX0           *float64                                                     `facebook:"product_priority_0"`
	ProductPriorityX1           *float64                                                     `facebook:"product_priority_1"`
	ProductPriorityX2           *float64                                                     `facebook:"product_priority_2"`
	ProductPriorityX3           *float64                                                     `facebook:"product_priority_3"`
	ProductPriorityX4           *float64                                                     `facebook:"product_priority_4"`
	ProductType                 *string                                                      `facebook:"product_type"`
	QuantityToSellOnFacebook    *uint64                                                      `facebook:"quantity_to_sell_on_facebook"`
	RetailerID                  *core.ID                                                     `facebook:"retailer_id"`
	RetailerProductGroupID      *core.ID                                                     `facebook:"retailer_product_group_id"`
	ReturnPolicyDays            *uint64                                                      `facebook:"return_policy_days"`
	RichTextDescription         *string                                                      `facebook:"rich_text_description"`
	SalePrice                   *uint64                                                      `facebook:"sale_price"`
	SalePriceEndDate            *time.Time                                                   `facebook:"sale_price_end_date"`
	SalePriceStartDate          *time.Time                                                   `facebook:"sale_price_start_date"`
	ShortDescription            *string                                                      `facebook:"short_description"`
	Size                        *string                                                      `facebook:"size"`
	StartDate                   *string                                                      `facebook:"start_date"`
	URL                         *string                                                      `facebook:"url"`
	Visibility                  *enums.ProductcatalogproductsVisibilityEnumParam             `facebook:"visibility"`
	WaComplianceCategory        *enums.ProductcatalogproductsWaComplianceCategoryEnumParam   `facebook:"wa_compliance_category"`
	WindowsPhoneAppID           *core.ID                                                     `facebook:"windows_phone_app_id"`
	WindowsPhoneAppName         *string                                                      `facebook:"windows_phone_app_name"`
	WindowsPhoneURL             *string                                                      `facebook:"windows_phone_url"`
	Extra                       core.Params                                                  `facebook:"-"`
}

func (params CreateProductCatalogProductsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdditionalImageUrls != nil {
		out["additional_image_urls"] = *params.AdditionalImageUrls
	}
	if params.AdditionalVariantAttributes != nil {
		out["additional_variant_attributes"] = *params.AdditionalVariantAttributes
	}
	if params.AgeGroup != nil {
		out["age_group"] = *params.AgeGroup
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	if params.AndroidAppName != nil {
		out["android_app_name"] = *params.AndroidAppName
	}
	if params.AndroidClass != nil {
		out["android_class"] = *params.AndroidClass
	}
	if params.AndroidPackage != nil {
		out["android_package"] = *params.AndroidPackage
	}
	if params.AndroidURL != nil {
		out["android_url"] = *params.AndroidURL
	}
	if params.Availability != nil {
		out["availability"] = *params.Availability
	}
	if params.Brand != nil {
		out["brand"] = *params.Brand
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.CategorySpecificFields != nil {
		out["category_specific_fields"] = *params.CategorySpecificFields
	}
	if params.CheckoutURL != nil {
		out["checkout_url"] = *params.CheckoutURL
	}
	if params.Color != nil {
		out["color"] = *params.Color
	}
	if params.CommerceTaxCategory != nil {
		out["commerce_tax_category"] = *params.CommerceTaxCategory
	}
	if params.Condition != nil {
		out["condition"] = *params.Condition
	}
	out["currency"] = params.Currency
	if params.CustomData != nil {
		out["custom_data"] = *params.CustomData
	}
	if params.CustomLabelX0 != nil {
		out["custom_label_0"] = *params.CustomLabelX0
	}
	if params.CustomLabelX1 != nil {
		out["custom_label_1"] = *params.CustomLabelX1
	}
	if params.CustomLabelX2 != nil {
		out["custom_label_2"] = *params.CustomLabelX2
	}
	if params.CustomLabelX3 != nil {
		out["custom_label_3"] = *params.CustomLabelX3
	}
	if params.CustomLabelX4 != nil {
		out["custom_label_4"] = *params.CustomLabelX4
	}
	if params.CustomNumberX0 != nil {
		out["custom_number_0"] = *params.CustomNumberX0
	}
	if params.CustomNumberX1 != nil {
		out["custom_number_1"] = *params.CustomNumberX1
	}
	if params.CustomNumberX2 != nil {
		out["custom_number_2"] = *params.CustomNumberX2
	}
	if params.CustomNumberX3 != nil {
		out["custom_number_3"] = *params.CustomNumberX3
	}
	if params.CustomNumberX4 != nil {
		out["custom_number_4"] = *params.CustomNumberX4
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.ExpirationDate != nil {
		out["expiration_date"] = *params.ExpirationDate
	}
	if params.FbProductCategory != nil {
		out["fb_product_category"] = *params.FbProductCategory
	}
	if params.Gender != nil {
		out["gender"] = *params.Gender
	}
	if params.Gtin != nil {
		out["gtin"] = *params.Gtin
	}
	if params.ImageURL != nil {
		out["image_url"] = *params.ImageURL
	}
	if params.ImporterAddress != nil {
		out["importer_address"] = *params.ImporterAddress
	}
	if params.ImporterName != nil {
		out["importer_name"] = *params.ImporterName
	}
	if params.Inventory != nil {
		out["inventory"] = *params.Inventory
	}
	if params.IosAppName != nil {
		out["ios_app_name"] = *params.IosAppName
	}
	if params.IosAppStoreID != nil {
		out["ios_app_store_id"] = *params.IosAppStoreID
	}
	if params.IosURL != nil {
		out["ios_url"] = *params.IosURL
	}
	if params.IpadAppName != nil {
		out["ipad_app_name"] = *params.IpadAppName
	}
	if params.IpadAppStoreID != nil {
		out["ipad_app_store_id"] = *params.IpadAppStoreID
	}
	if params.IpadURL != nil {
		out["ipad_url"] = *params.IpadURL
	}
	if params.IphoneAppName != nil {
		out["iphone_app_name"] = *params.IphoneAppName
	}
	if params.IphoneAppStoreID != nil {
		out["iphone_app_store_id"] = *params.IphoneAppStoreID
	}
	if params.IphoneURL != nil {
		out["iphone_url"] = *params.IphoneURL
	}
	if params.LaunchDate != nil {
		out["launch_date"] = *params.LaunchDate
	}
	if params.LiveSpecialPrice != nil {
		out["live_special_price"] = *params.LiveSpecialPrice
	}
	if params.ManufacturerInfo != nil {
		out["manufacturer_info"] = *params.ManufacturerInfo
	}
	if params.ManufacturerPartNumber != nil {
		out["manufacturer_part_number"] = *params.ManufacturerPartNumber
	}
	if params.MarkedForProductLaunch != nil {
		out["marked_for_product_launch"] = *params.MarkedForProductLaunch
	}
	if params.Material != nil {
		out["material"] = *params.Material
	}
	if params.MobileLink != nil {
		out["mobile_link"] = *params.MobileLink
	}
	out["name"] = params.Name
	if params.OrderingIndex != nil {
		out["ordering_index"] = *params.OrderingIndex
	}
	if params.OriginCountry != nil {
		out["origin_country"] = *params.OriginCountry
	}
	if params.Pattern != nil {
		out["pattern"] = *params.Pattern
	}
	out["price"] = params.Price
	if params.ProductPriorityX0 != nil {
		out["product_priority_0"] = *params.ProductPriorityX0
	}
	if params.ProductPriorityX1 != nil {
		out["product_priority_1"] = *params.ProductPriorityX1
	}
	if params.ProductPriorityX2 != nil {
		out["product_priority_2"] = *params.ProductPriorityX2
	}
	if params.ProductPriorityX3 != nil {
		out["product_priority_3"] = *params.ProductPriorityX3
	}
	if params.ProductPriorityX4 != nil {
		out["product_priority_4"] = *params.ProductPriorityX4
	}
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	if params.QuantityToSellOnFacebook != nil {
		out["quantity_to_sell_on_facebook"] = *params.QuantityToSellOnFacebook
	}
	if params.RetailerID != nil {
		out["retailer_id"] = *params.RetailerID
	}
	if params.RetailerProductGroupID != nil {
		out["retailer_product_group_id"] = *params.RetailerProductGroupID
	}
	if params.ReturnPolicyDays != nil {
		out["return_policy_days"] = *params.ReturnPolicyDays
	}
	if params.RichTextDescription != nil {
		out["rich_text_description"] = *params.RichTextDescription
	}
	if params.SalePrice != nil {
		out["sale_price"] = *params.SalePrice
	}
	if params.SalePriceEndDate != nil {
		out["sale_price_end_date"] = *params.SalePriceEndDate
	}
	if params.SalePriceStartDate != nil {
		out["sale_price_start_date"] = *params.SalePriceStartDate
	}
	if params.ShortDescription != nil {
		out["short_description"] = *params.ShortDescription
	}
	if params.Size != nil {
		out["size"] = *params.Size
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.Visibility != nil {
		out["visibility"] = *params.Visibility
	}
	if params.WaComplianceCategory != nil {
		out["wa_compliance_category"] = *params.WaComplianceCategory
	}
	if params.WindowsPhoneAppID != nil {
		out["windows_phone_app_id"] = *params.WindowsPhoneAppID
	}
	if params.WindowsPhoneAppName != nil {
		out["windows_phone_app_name"] = *params.WindowsPhoneAppName
	}
	if params.WindowsPhoneURL != nil {
		out["windows_phone_url"] = *params.WindowsPhoneURL
	}
	return out
}

func CreateProductCatalogProducts(ctx context.Context, client *core.Client, id string, params CreateProductCatalogProductsParams) (*objects.ProductItem, error) {
	var out objects.ProductItem
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogUpdateGeneratedImageConfigParams struct {
	Data  []map[string]interface{} `facebook:"data"`
	Extra core.Params              `facebook:"-"`
}

func (params CreateProductCatalogUpdateGeneratedImageConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["data"] = params.Data
	return out
}

func CreateProductCatalogUpdateGeneratedImageConfig(ctx context.Context, client *core.Client, id string, params CreateProductCatalogUpdateGeneratedImageConfigParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "update_generated_image_config"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogVehicleOffersParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogVehicleOffersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogVehicleOffers(ctx context.Context, client *core.Client, id string, params GetProductCatalogVehicleOffersParams) (*core.Cursor[objects.VehicleOffer], error) {
	var out core.Cursor[objects.VehicleOffer]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicle_offers"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogVehiclesParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductCatalogVehiclesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BulkPagination != nil {
		out["bulk_pagination"] = *params.BulkPagination
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	return out
}

func GetProductCatalogVehicles(ctx context.Context, client *core.Client, id string, params GetProductCatalogVehiclesParams) (*core.Cursor[objects.Vehicle], error) {
	var out core.Cursor[objects.Vehicle]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicles"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogVehiclesParams struct {
	Address        map[string]interface{}                              `facebook:"address"`
	Applinks       *map[string]interface{}                             `facebook:"applinks"`
	Availability   *enums.ProductcatalogvehiclesAvailabilityEnumParam  `facebook:"availability"`
	BodyStyle      enums.ProductcatalogvehiclesBodyStyleEnumParam      `facebook:"body_style"`
	Condition      *enums.ProductcatalogvehiclesConditionEnumParam     `facebook:"condition"`
	Currency       string                                              `facebook:"currency"`
	DateFirstOnLot *string                                             `facebook:"date_first_on_lot"`
	DealerID       *core.ID                                            `facebook:"dealer_id"`
	DealerName     *string                                             `facebook:"dealer_name"`
	DealerPhone    *string                                             `facebook:"dealer_phone"`
	Description    string                                              `facebook:"description"`
	Drivetrain     *enums.ProductcatalogvehiclesDrivetrainEnumParam    `facebook:"drivetrain"`
	ExteriorColor  string                                              `facebook:"exterior_color"`
	FbPageID       *core.ID                                            `facebook:"fb_page_id"`
	FuelType       *enums.ProductcatalogvehiclesFuelTypeEnumParam      `facebook:"fuel_type"`
	Images         []map[string]interface{}                            `facebook:"images"`
	InteriorColor  *string                                             `facebook:"interior_color"`
	Make           string                                              `facebook:"make"`
	Mileage        map[string]interface{}                              `facebook:"mileage"`
	Model          string                                              `facebook:"model"`
	Price          uint64                                              `facebook:"price"`
	StateOfVehicle enums.ProductcatalogvehiclesStateOfVehicleEnumParam `facebook:"state_of_vehicle"`
	Title          string                                              `facebook:"title"`
	Transmission   *enums.ProductcatalogvehiclesTransmissionEnumParam  `facebook:"transmission"`
	Trim           *string                                             `facebook:"trim"`
	URL            string                                              `facebook:"url"`
	VehicleID      core.ID                                             `facebook:"vehicle_id"`
	VehicleType    *enums.ProductcatalogvehiclesVehicleTypeEnumParam   `facebook:"vehicle_type"`
	Vin            string                                              `facebook:"vin"`
	Year           uint64                                              `facebook:"year"`
	Extra          core.Params                                         `facebook:"-"`
}

func (params CreateProductCatalogVehiclesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["address"] = params.Address
	if params.Applinks != nil {
		out["applinks"] = *params.Applinks
	}
	if params.Availability != nil {
		out["availability"] = *params.Availability
	}
	out["body_style"] = params.BodyStyle
	if params.Condition != nil {
		out["condition"] = *params.Condition
	}
	out["currency"] = params.Currency
	if params.DateFirstOnLot != nil {
		out["date_first_on_lot"] = *params.DateFirstOnLot
	}
	if params.DealerID != nil {
		out["dealer_id"] = *params.DealerID
	}
	if params.DealerName != nil {
		out["dealer_name"] = *params.DealerName
	}
	if params.DealerPhone != nil {
		out["dealer_phone"] = *params.DealerPhone
	}
	out["description"] = params.Description
	if params.Drivetrain != nil {
		out["drivetrain"] = *params.Drivetrain
	}
	out["exterior_color"] = params.ExteriorColor
	if params.FbPageID != nil {
		out["fb_page_id"] = *params.FbPageID
	}
	if params.FuelType != nil {
		out["fuel_type"] = *params.FuelType
	}
	out["images"] = params.Images
	if params.InteriorColor != nil {
		out["interior_color"] = *params.InteriorColor
	}
	out["make"] = params.Make
	out["mileage"] = params.Mileage
	out["model"] = params.Model
	out["price"] = params.Price
	out["state_of_vehicle"] = params.StateOfVehicle
	out["title"] = params.Title
	if params.Transmission != nil {
		out["transmission"] = *params.Transmission
	}
	if params.Trim != nil {
		out["trim"] = *params.Trim
	}
	out["url"] = params.URL
	out["vehicle_id"] = params.VehicleID
	if params.VehicleType != nil {
		out["vehicle_type"] = *params.VehicleType
	}
	out["vin"] = params.Vin
	out["year"] = params.Year
	return out
}

func CreateProductCatalogVehicles(ctx context.Context, client *core.Client, id string, params CreateProductCatalogVehiclesParams) (*objects.Vehicle, error) {
	var out objects.Vehicle
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "vehicles"), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogVersionConfigsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogVersionConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogVersionConfigs(ctx context.Context, client *core.Client, id string, params GetProductCatalogVersionConfigsParams) (*core.Cursor[objects.CatalogContentVersionConfig], error) {
	var out core.Cursor[objects.CatalogContentVersionConfig]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "version_configs"), params.ToParams(), &out)
	return &out, err
}

type CreateProductCatalogVersionItemsBatchParams struct {
	AllowUpsert *bool                  `facebook:"allow_upsert"`
	ItemType    string                 `facebook:"item_type"`
	ItemVersion string                 `facebook:"item_version"`
	Requests    map[string]interface{} `facebook:"requests"`
	Version     *uint64                `facebook:"version"`
	Extra       core.Params            `facebook:"-"`
}

func (params CreateProductCatalogVersionItemsBatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowUpsert != nil {
		out["allow_upsert"] = *params.AllowUpsert
	}
	out["item_type"] = params.ItemType
	out["item_version"] = params.ItemVersion
	out["requests"] = params.Requests
	if params.Version != nil {
		out["version"] = *params.Version
	}
	return out
}

func CreateProductCatalogVersionItemsBatch(ctx context.Context, client *core.Client, id string, params CreateProductCatalogVersionItemsBatchParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "version_items_batch"), params.ToParams(), &out)
	return &out, err
}

type DeleteProductCatalogParams struct {
	AllowDeleteCatalogWithLiveProductSet *bool       `facebook:"allow_delete_catalog_with_live_product_set"`
	Extra                                core.Params `facebook:"-"`
}

func (params DeleteProductCatalogParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowDeleteCatalogWithLiveProductSet != nil {
		out["allow_delete_catalog_with_live_product_set"] = *params.AllowDeleteCatalogWithLiveProductSet
	}
	return out
}

func DeleteProductCatalog(ctx context.Context, client *core.Client, id string, params DeleteProductCatalogParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetProductCatalogParams struct {
	SegmentUseCases *[]enums.ProductcatalogSegmentUseCases `facebook:"segment_use_cases"`
	Extra           core.Params                            `facebook:"-"`
}

func (params GetProductCatalogParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.SegmentUseCases != nil {
		out["segment_use_cases"] = *params.SegmentUseCases
	}
	return out
}

func GetProductCatalog(ctx context.Context, client *core.Client, id string, params GetProductCatalogParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateProductCatalogParams struct {
	AdditionalVerticalOption   *enums.ProductcatalogAdditionalVerticalOption `facebook:"additional_vertical_option"`
	DaDisplaySettings          *map[string]interface{}                       `facebook:"da_display_settings"`
	DefaultImageURL            *string                                       `facebook:"default_image_url"`
	DestinationCatalogSettings *map[string]interface{}                       `facebook:"destination_catalog_settings"`
	FallbackImageURL           *string                                       `facebook:"fallback_image_url"`
	FlightCatalogSettings      *map[string]interface{}                       `facebook:"flight_catalog_settings"`
	Name                       *string                                       `facebook:"name"`
	PartnerIntegration         *map[string]interface{}                       `facebook:"partner_integration"`
	StoreCatalogSettings       *map[string]interface{}                       `facebook:"store_catalog_settings"`
	Extra                      core.Params                                   `facebook:"-"`
}

func (params UpdateProductCatalogParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdditionalVerticalOption != nil {
		out["additional_vertical_option"] = *params.AdditionalVerticalOption
	}
	if params.DaDisplaySettings != nil {
		out["da_display_settings"] = *params.DaDisplaySettings
	}
	if params.DefaultImageURL != nil {
		out["default_image_url"] = *params.DefaultImageURL
	}
	if params.DestinationCatalogSettings != nil {
		out["destination_catalog_settings"] = *params.DestinationCatalogSettings
	}
	if params.FallbackImageURL != nil {
		out["fallback_image_url"] = *params.FallbackImageURL
	}
	if params.FlightCatalogSettings != nil {
		out["flight_catalog_settings"] = *params.FlightCatalogSettings
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.PartnerIntegration != nil {
		out["partner_integration"] = *params.PartnerIntegration
	}
	if params.StoreCatalogSettings != nil {
		out["store_catalog_settings"] = *params.StoreCatalogSettings
	}
	return out
}

func UpdateProductCatalog(ctx context.Context, client *core.Client, id string, params UpdateProductCatalogParams) (*objects.ProductCatalog, error) {
	var out objects.ProductCatalog
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
