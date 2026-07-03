package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedAutomotiveModelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedAutomotiveModelsParams) ToParams() core.Params {
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

func GetProductFeedAutomotiveModels(ctx context.Context, client *core.Client, id string, params GetProductFeedAutomotiveModelsParams) (*core.Cursor[objects.AutomotiveModel], error) {
	var out core.Cursor[objects.AutomotiveModel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "automotive_models"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedDestinationsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedDestinationsParams) ToParams() core.Params {
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

func GetProductFeedDestinations(ctx context.Context, client *core.Client, id string, params GetProductFeedDestinationsParams) (*core.Cursor[objects.Destination], error) {
	var out core.Cursor[objects.Destination]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "destinations"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedFlightsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedFlightsParams) ToParams() core.Params {
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

func GetProductFeedFlights(ctx context.Context, client *core.Client, id string, params GetProductFeedFlightsParams) (*core.Cursor[objects.Flight], error) {
	var out core.Cursor[objects.Flight]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "flights"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedHomeListingsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedHomeListingsParams) ToParams() core.Params {
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

func GetProductFeedHomeListings(ctx context.Context, client *core.Client, id string, params GetProductFeedHomeListingsParams) (*core.Cursor[objects.HomeListing], error) {
	var out core.Cursor[objects.HomeListing]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "home_listings"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedHotelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedHotelsParams) ToParams() core.Params {
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

func GetProductFeedHotels(ctx context.Context, client *core.Client, id string, params GetProductFeedHotelsParams) (*core.Cursor[objects.Hotel], error) {
	var out core.Cursor[objects.Hotel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "hotels"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedProductsParams struct {
	BulkPagination *bool                                            `facebook:"bulk_pagination"`
	ErrorPriority  *enums.ProductfeedproductsErrorPriorityEnumParam `facebook:"error_priority"`
	ErrorType      *enums.ProductfeedproductsErrorTypeEnumParam     `facebook:"error_type"`
	Filter         *map[string]interface{}                          `facebook:"filter"`
	Extra          core.Params                                      `facebook:"-"`
}

func (params GetProductFeedProductsParams) ToParams() core.Params {
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
	return out
}

func GetProductFeedProducts(ctx context.Context, client *core.Client, id string, params GetProductFeedProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedRules(ctx context.Context, client *core.Client, id string, params GetProductFeedRulesParams) (*core.Cursor[objects.ProductFeedRule], error) {
	var out core.Cursor[objects.ProductFeedRule]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), &out)
	return &out, err
}

type CreateProductFeedRulesParams struct {
	Attribute string                                  `facebook:"attribute"`
	Params    *map[string]interface{}                 `facebook:"params"`
	RuleType  enums.ProductfeedrulesRuleTypeEnumParam `facebook:"rule_type"`
	Extra     core.Params                             `facebook:"-"`
}

func (params CreateProductFeedRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["attribute"] = params.Attribute
	if params.Params != nil {
		out["params"] = *params.Params
	}
	out["rule_type"] = params.RuleType
	return out
}

func CreateProductFeedRules(ctx context.Context, client *core.Client, id string, params CreateProductFeedRulesParams) (*objects.ProductFeedRule, error) {
	var out objects.ProductFeedRule
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "rules"), params.ToParams(), &out)
	return &out, err
}

type CreateProductFeedSupplementaryFeedAssocsParams struct {
	AssocData []map[string]interface{} `facebook:"assoc_data"`
	Extra     core.Params              `facebook:"-"`
}

func (params CreateProductFeedSupplementaryFeedAssocsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["assoc_data"] = params.AssocData
	return out
}

func CreateProductFeedSupplementaryFeedAssocs(ctx context.Context, client *core.Client, id string, params CreateProductFeedSupplementaryFeedAssocsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "supplementary_feed_assocs"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedUploadSchedulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadSchedules(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadSchedulesParams) (*core.Cursor[objects.ProductFeedSchedule], error) {
	var out core.Cursor[objects.ProductFeedSchedule]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "upload_schedules"), params.ToParams(), &out)
	return &out, err
}

type CreateProductFeedUploadSchedulesParams struct {
	UploadSchedule *string     `facebook:"upload_schedule"`
	Extra          core.Params `facebook:"-"`
}

func (params CreateProductFeedUploadSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UploadSchedule != nil {
		out["upload_schedule"] = *params.UploadSchedule
	}
	return out
}

func CreateProductFeedUploadSchedules(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadSchedulesParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "upload_schedules"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedUploadsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploads(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadsParams) (*core.Cursor[objects.ProductFeedUpload], error) {
	var out core.Cursor[objects.ProductFeedUpload]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "uploads"), params.ToParams(), &out)
	return &out, err
}

type CreateProductFeedUploadsParams struct {
	FbeExternalBusinessID *core.ID        `facebook:"fbe_external_business_id"`
	File                  *core.FileParam `facebook:"file"`
	Password              *string         `facebook:"password"`
	UpdateOnly            *bool           `facebook:"update_only"`
	URL                   *string         `facebook:"url"`
	Username              *string         `facebook:"username"`
	Extra                 core.Params     `facebook:"-"`
}

func (params CreateProductFeedUploadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.FbeExternalBusinessID != nil {
		out["fbe_external_business_id"] = *params.FbeExternalBusinessID
	}
	if params.File != nil {
		out["file"] = *params.File
	}
	if params.Password != nil {
		out["password"] = *params.Password
	}
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

func CreateProductFeedUploads(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadsParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "uploads"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedVehicleOffersParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedVehicleOffersParams) ToParams() core.Params {
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

func GetProductFeedVehicleOffers(ctx context.Context, client *core.Client, id string, params GetProductFeedVehicleOffersParams) (*core.Cursor[objects.VehicleOffer], error) {
	var out core.Cursor[objects.VehicleOffer]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicle_offers"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedVehiclesParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductFeedVehiclesParams) ToParams() core.Params {
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

func GetProductFeedVehicles(ctx context.Context, client *core.Client, id string, params GetProductFeedVehiclesParams) (*core.Cursor[objects.Vehicle], error) {
	var out core.Cursor[objects.Vehicle]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicles"), params.ToParams(), &out)
	return &out, err
}

type DeleteProductFeedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteProductFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteProductFeed(ctx context.Context, client *core.Client, id string, params DeleteProductFeedParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeed(ctx context.Context, client *core.Client, id string, params GetProductFeedParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateProductFeedParams struct {
	DefaultCurrency    *string                            `facebook:"default_currency"`
	DeletionEnabled    *bool                              `facebook:"deletion_enabled"`
	Delimiter          *enums.ProductfeedDelimiter        `facebook:"delimiter"`
	Encoding           *enums.ProductfeedEncoding         `facebook:"encoding"`
	MigratedFromFeedID *core.ID                           `facebook:"migrated_from_feed_id"`
	Name               *string                            `facebook:"name"`
	QuotedFieldsMode   *enums.ProductfeedQuotedFieldsMode `facebook:"quoted_fields_mode"`
	Schedule           *string                            `facebook:"schedule"`
	UpdateSchedule     *string                            `facebook:"update_schedule"`
	Extra              core.Params                        `facebook:"-"`
}

func (params UpdateProductFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
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
	if params.MigratedFromFeedID != nil {
		out["migrated_from_feed_id"] = *params.MigratedFromFeedID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.QuotedFieldsMode != nil {
		out["quoted_fields_mode"] = *params.QuotedFieldsMode
	}
	if params.Schedule != nil {
		out["schedule"] = *params.Schedule
	}
	if params.UpdateSchedule != nil {
		out["update_schedule"] = *params.UpdateSchedule
	}
	return out
}

func UpdateProductFeed(ctx context.Context, client *core.Client, id string, params UpdateProductFeedParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
