package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetAdsPixelAdaccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdsPixelAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdsPixelAgenciesParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func DeleteAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelAgenciesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params GetAdsPixelAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelAgenciesParams struct {
	Business       string                                          `facebook:"business"`
	PermittedTasks []enums.AdspixelagenciesPermittedTasksEnumParam `facebook:"permitted_tasks"`
	Extra          core.Params                                     `facebook:"-"`
}

func (params CreateAdsPixelAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	out["permitted_tasks"] = params.PermittedTasks
	return out
}

func CreateAdsPixelAgencies(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAgenciesParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelAhpConfigsParams struct {
	ApplinkAutosetup bool        `facebook:"applink_autosetup"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateAdsPixelAhpConfigsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["applink_autosetup"] = params.ApplinkAutosetup
	return out
}

func CreateAdsPixelAhpConfigs(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAhpConfigsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "ahp_configs"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelAssignedUsers(ctx context.Context, client *core.Client, id string, params GetAdsPixelAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelAssignedUsersParams struct {
	Tasks []enums.AdspixelassignedUsersTasksEnumParam `facebook:"tasks"`
	User  int                                         `facebook:"user"`
	Extra core.Params                                 `facebook:"-"`
}

func (params CreateAdsPixelAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["tasks"] = params.Tasks
	out["user"] = params.User
	return out
}

func CreateAdsPixelAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateAdsPixelAssignedUsersParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelDaChecksParams struct {
	Checks           *[]string                                        `facebook:"checks"`
	ConnectionMethod *enums.AdspixeldaChecksConnectionMethodEnumParam `facebook:"connection_method"`
	Extra            core.Params                                      `facebook:"-"`
}

func (params GetAdsPixelDaChecksParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Checks != nil {
		out["checks"] = *params.Checks
	}
	if params.ConnectionMethod != nil {
		out["connection_method"] = *params.ConnectionMethod
	}
	return out
}

func GetAdsPixelDaChecks(ctx context.Context, client *core.Client, id string, params GetAdsPixelDaChecksParams) (*core.Cursor[objects.DACheck], error) {
	var out core.Cursor[objects.DACheck]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "da_checks"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelEventsParams struct {
	Data          []string                  `facebook:"data"`
	NamespaceID   *core.ID                  `facebook:"namespace_id"`
	PartnerAgent  *string                   `facebook:"partner_agent"`
	Platforms     *[]map[string]interface{} `facebook:"platforms"`
	Progress      *map[string]interface{}   `facebook:"progress"`
	TestEventCode *string                   `facebook:"test_event_code"`
	Trace         *uint64                   `facebook:"trace"`
	UploadID      *core.ID                  `facebook:"upload_id"`
	UploadSource  *string                   `facebook:"upload_source"`
	UploadTag     *string                   `facebook:"upload_tag"`
	Extra         core.Params               `facebook:"-"`
}

func (params CreateAdsPixelEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["data"] = params.Data
	if params.NamespaceID != nil {
		out["namespace_id"] = *params.NamespaceID
	}
	if params.PartnerAgent != nil {
		out["partner_agent"] = *params.PartnerAgent
	}
	if params.Platforms != nil {
		out["platforms"] = *params.Platforms
	}
	if params.Progress != nil {
		out["progress"] = *params.Progress
	}
	if params.TestEventCode != nil {
		out["test_event_code"] = *params.TestEventCode
	}
	if params.Trace != nil {
		out["trace"] = *params.Trace
	}
	if params.UploadID != nil {
		out["upload_id"] = *params.UploadID
	}
	if params.UploadSource != nil {
		out["upload_source"] = *params.UploadSource
	}
	if params.UploadTag != nil {
		out["upload_tag"] = *params.UploadTag
	}
	return out
}

func CreateAdsPixelEvents(ctx context.Context, client *core.Client, id string, params CreateAdsPixelEventsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "events"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelOfflineEventUploadsParams struct {
	EndTime   *time.Time                                        `facebook:"end_time"`
	Order     *enums.AdspixelofflineEventUploadsOrderEnumParam  `facebook:"order"`
	SortBy    *enums.AdspixelofflineEventUploadsSortByEnumParam `facebook:"sort_by"`
	StartTime *time.Time                                        `facebook:"start_time"`
	UploadTag *string                                           `facebook:"upload_tag"`
	Extra     core.Params                                       `facebook:"-"`
}

func (params GetAdsPixelOfflineEventUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.SortBy != nil {
		out["sort_by"] = *params.SortBy
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.UploadTag != nil {
		out["upload_tag"] = *params.UploadTag
	}
	return out
}

func GetAdsPixelOfflineEventUploads(ctx context.Context, client *core.Client, id string, params GetAdsPixelOfflineEventUploadsParams) (*core.Cursor[objects.OfflineConversionDataSetUpload], error) {
	var out core.Cursor[objects.OfflineConversionDataSetUpload]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "offline_event_uploads"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelOpenbridgeConfigurationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelOpenbridgeConfigurationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelOpenbridgeConfigurations(ctx context.Context, client *core.Client, id string, params GetAdsPixelOpenbridgeConfigurationsParams) (*core.Cursor[objects.OpenBridgeConfiguration], error) {
	var out core.Cursor[objects.OpenBridgeConfiguration]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "openbridge_configurations"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelShadowtraffichelperParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdsPixelShadowtraffichelperParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdsPixelShadowtraffichelper(ctx context.Context, client *core.Client, id string, params CreateAdsPixelShadowtraffichelperParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "shadowtraffichelper"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdsPixelSharedAccountsParams struct {
	AccountID core.ID     `facebook:"account_id"`
	Business  string      `facebook:"business"`
	Extra     core.Params `facebook:"-"`
}

func (params DeleteAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["business"] = params.Business
	return out
}

func DeleteAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params DeleteAdsPixelSharedAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelSharedAccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateAdsPixelSharedAccountsParams struct {
	AccountID core.ID     `facebook:"account_id"`
	Business  string      `facebook:"business"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateAdsPixelSharedAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["business"] = params.Business
	return out
}

func CreateAdsPixelSharedAccounts(ctx context.Context, client *core.Client, id string, params CreateAdsPixelSharedAccountsParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "shared_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelSharedAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelSharedAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelSharedAgencies(ctx context.Context, client *core.Client, id string, params GetAdsPixelSharedAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shared_agencies"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelStatsParams struct {
	Agent       *string                                  `facebook:"agent"`
	Aggregation *enums.AdspixelstatsAggregationEnumParam `facebook:"aggregation"`
	EndTime     *time.Time                               `facebook:"end_time"`
	Event       *string                                  `facebook:"event"`
	EventSource *string                                  `facebook:"event_source"`
	StartTime   *time.Time                               `facebook:"start_time"`
	Extra       core.Params                              `facebook:"-"`
}

func (params GetAdsPixelStatsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Agent != nil {
		out["agent"] = *params.Agent
	}
	if params.Aggregation != nil {
		out["aggregation"] = *params.Aggregation
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Event != nil {
		out["event"] = *params.Event
	}
	if params.EventSource != nil {
		out["event_source"] = *params.EventSource
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetAdsPixelStats(ctx context.Context, client *core.Client, id string, params GetAdsPixelStatsParams) (*core.Cursor[objects.AdsPixelStatsResult], error) {
	var out core.Cursor[objects.AdsPixelStatsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "stats"), params.ToParams(), &out)
	return &out, err
}

type GetAdsPixelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixel(ctx context.Context, client *core.Client, id string, params GetAdsPixelParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdsPixelParams struct {
	AutomaticMatchingFields       *[]enums.AdspixelAutomaticMatchingFields `facebook:"automatic_matching_fields"`
	DataUseSetting                *enums.AdspixelDataUseSetting            `facebook:"data_use_setting"`
	EnableAutomaticMatching       *bool                                    `facebook:"enable_automatic_matching"`
	FirstPartyCookieStatus        *enums.AdspixelFirstPartyCookieStatus    `facebook:"first_party_cookie_status"`
	Name                          *string                                  `facebook:"name"`
	ServerEventsBusinessIds       *[]core.ID                               `facebook:"server_events_business_ids"`
	ServerEventsBusinessIdsAdd    *[]string                                `facebook:"server_events_business_ids_add"`
	ServerEventsBusinessIdsRemove *[]string                                `facebook:"server_events_business_ids_remove"`
	Extra                         core.Params                              `facebook:"-"`
}

func (params UpdateAdsPixelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AutomaticMatchingFields != nil {
		out["automatic_matching_fields"] = *params.AutomaticMatchingFields
	}
	if params.DataUseSetting != nil {
		out["data_use_setting"] = *params.DataUseSetting
	}
	if params.EnableAutomaticMatching != nil {
		out["enable_automatic_matching"] = *params.EnableAutomaticMatching
	}
	if params.FirstPartyCookieStatus != nil {
		out["first_party_cookie_status"] = *params.FirstPartyCookieStatus
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.ServerEventsBusinessIds != nil {
		out["server_events_business_ids"] = *params.ServerEventsBusinessIds
	}
	if params.ServerEventsBusinessIdsAdd != nil {
		out["server_events_business_ids_add"] = *params.ServerEventsBusinessIdsAdd
	}
	if params.ServerEventsBusinessIdsRemove != nil {
		out["server_events_business_ids_remove"] = *params.ServerEventsBusinessIdsRemove
	}
	return out
}

func UpdateAdsPixel(ctx context.Context, client *core.Client, id string, params UpdateAdsPixelParams) (*objects.AdsPixel, error) {
	var out objects.AdsPixel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
