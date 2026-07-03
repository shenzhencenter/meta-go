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

func GetProductFeedAutomotiveModelsBatchCall(id string, params GetProductFeedAutomotiveModelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "automotive_models"), params.ToParams(), options...)
}

func NewGetProductFeedAutomotiveModelsBatchRequest(id string, params GetProductFeedAutomotiveModelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AutomotiveModel]] {
	return core.NewBatchRequest[core.Cursor[objects.AutomotiveModel]](GetProductFeedAutomotiveModelsBatchCall(id, params, options...))
}

func DecodeGetProductFeedAutomotiveModelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AutomotiveModel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AutomotiveModel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedAutomotiveModels(ctx context.Context, client *core.Client, id string, params GetProductFeedAutomotiveModelsParams) (*core.Cursor[objects.AutomotiveModel], error) {
	var out core.Cursor[objects.AutomotiveModel]
	call := GetProductFeedAutomotiveModelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedDestinationsBatchCall(id string, params GetProductFeedDestinationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "destinations"), params.ToParams(), options...)
}

func NewGetProductFeedDestinationsBatchRequest(id string, params GetProductFeedDestinationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Destination]] {
	return core.NewBatchRequest[core.Cursor[objects.Destination]](GetProductFeedDestinationsBatchCall(id, params, options...))
}

func DecodeGetProductFeedDestinationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Destination], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Destination]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedDestinations(ctx context.Context, client *core.Client, id string, params GetProductFeedDestinationsParams) (*core.Cursor[objects.Destination], error) {
	var out core.Cursor[objects.Destination]
	call := GetProductFeedDestinationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedFlightsBatchCall(id string, params GetProductFeedFlightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "flights"), params.ToParams(), options...)
}

func NewGetProductFeedFlightsBatchRequest(id string, params GetProductFeedFlightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Flight]] {
	return core.NewBatchRequest[core.Cursor[objects.Flight]](GetProductFeedFlightsBatchCall(id, params, options...))
}

func DecodeGetProductFeedFlightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Flight], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Flight]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedFlights(ctx context.Context, client *core.Client, id string, params GetProductFeedFlightsParams) (*core.Cursor[objects.Flight], error) {
	var out core.Cursor[objects.Flight]
	call := GetProductFeedFlightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedHomeListingsBatchCall(id string, params GetProductFeedHomeListingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "home_listings"), params.ToParams(), options...)
}

func NewGetProductFeedHomeListingsBatchRequest(id string, params GetProductFeedHomeListingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.HomeListing]] {
	return core.NewBatchRequest[core.Cursor[objects.HomeListing]](GetProductFeedHomeListingsBatchCall(id, params, options...))
}

func DecodeGetProductFeedHomeListingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.HomeListing], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.HomeListing]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedHomeListings(ctx context.Context, client *core.Client, id string, params GetProductFeedHomeListingsParams) (*core.Cursor[objects.HomeListing], error) {
	var out core.Cursor[objects.HomeListing]
	call := GetProductFeedHomeListingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedHotelsBatchCall(id string, params GetProductFeedHotelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "hotels"), params.ToParams(), options...)
}

func NewGetProductFeedHotelsBatchRequest(id string, params GetProductFeedHotelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Hotel]] {
	return core.NewBatchRequest[core.Cursor[objects.Hotel]](GetProductFeedHotelsBatchCall(id, params, options...))
}

func DecodeGetProductFeedHotelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Hotel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Hotel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedHotels(ctx context.Context, client *core.Client, id string, params GetProductFeedHotelsParams) (*core.Cursor[objects.Hotel], error) {
	var out core.Cursor[objects.Hotel]
	call := GetProductFeedHotelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedProductsBatchCall(id string, params GetProductFeedProductsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), options...)
}

func NewGetProductFeedProductsBatchRequest(id string, params GetProductFeedProductsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductItem]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductItem]](GetProductFeedProductsBatchCall(id, params, options...))
}

func DecodeGetProductFeedProductsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductItem], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductItem]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedProducts(ctx context.Context, client *core.Client, id string, params GetProductFeedProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	call := GetProductFeedProductsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedRulesBatchCall(id string, params GetProductFeedRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), options...)
}

func NewGetProductFeedRulesBatchRequest(id string, params GetProductFeedRulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedRule]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedRule]](GetProductFeedRulesBatchCall(id, params, options...))
}

func DecodeGetProductFeedRulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedRule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedRule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedRules(ctx context.Context, client *core.Client, id string, params GetProductFeedRulesParams) (*core.Cursor[objects.ProductFeedRule], error) {
	var out core.Cursor[objects.ProductFeedRule]
	call := GetProductFeedRulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateProductFeedRulesBatchCall(id string, params CreateProductFeedRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "rules"), params.ToParams(), options...)
}

func NewCreateProductFeedRulesBatchRequest(id string, params CreateProductFeedRulesParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedRule] {
	return core.NewBatchRequest[objects.ProductFeedRule](CreateProductFeedRulesBatchCall(id, params, options...))
}

func DecodeCreateProductFeedRulesBatchResponse(response *core.BatchResponse) (*objects.ProductFeedRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateProductFeedRules(ctx context.Context, client *core.Client, id string, params CreateProductFeedRulesParams) (*objects.ProductFeedRule, error) {
	var out objects.ProductFeedRule
	call := CreateProductFeedRulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateProductFeedSupplementaryFeedAssocsBatchCall(id string, params CreateProductFeedSupplementaryFeedAssocsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "supplementary_feed_assocs"), params.ToParams(), options...)
}

func NewCreateProductFeedSupplementaryFeedAssocsBatchRequest(id string, params CreateProductFeedSupplementaryFeedAssocsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateProductFeedSupplementaryFeedAssocsBatchCall(id, params, options...))
}

func DecodeCreateProductFeedSupplementaryFeedAssocsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateProductFeedSupplementaryFeedAssocs(ctx context.Context, client *core.Client, id string, params CreateProductFeedSupplementaryFeedAssocsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateProductFeedSupplementaryFeedAssocsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedUploadSchedulesBatchCall(id string, params GetProductFeedUploadSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "upload_schedules"), params.ToParams(), options...)
}

func NewGetProductFeedUploadSchedulesBatchRequest(id string, params GetProductFeedUploadSchedulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedSchedule]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedSchedule]](GetProductFeedUploadSchedulesBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadSchedulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedSchedule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedSchedule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadSchedules(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadSchedulesParams) (*core.Cursor[objects.ProductFeedSchedule], error) {
	var out core.Cursor[objects.ProductFeedSchedule]
	call := GetProductFeedUploadSchedulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateProductFeedUploadSchedulesBatchCall(id string, params CreateProductFeedUploadSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "upload_schedules"), params.ToParams(), options...)
}

func NewCreateProductFeedUploadSchedulesBatchRequest(id string, params CreateProductFeedUploadSchedulesParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeed] {
	return core.NewBatchRequest[objects.ProductFeed](CreateProductFeedUploadSchedulesBatchCall(id, params, options...))
}

func DecodeCreateProductFeedUploadSchedulesBatchResponse(response *core.BatchResponse) (*objects.ProductFeed, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeed
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateProductFeedUploadSchedules(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadSchedulesParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	call := CreateProductFeedUploadSchedulesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedUploadsBatchCall(id string, params GetProductFeedUploadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "uploads"), params.ToParams(), options...)
}

func NewGetProductFeedUploadsBatchRequest(id string, params GetProductFeedUploadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedUpload]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedUpload]](GetProductFeedUploadsBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedUpload], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedUpload]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploads(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadsParams) (*core.Cursor[objects.ProductFeedUpload], error) {
	var out core.Cursor[objects.ProductFeedUpload]
	call := GetProductFeedUploadsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateProductFeedUploadsBatchCall(id string, params CreateProductFeedUploadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "uploads"), params.ToParams(), options...)
}

func NewCreateProductFeedUploadsBatchRequest(id string, params CreateProductFeedUploadsParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedUpload] {
	return core.NewBatchRequest[objects.ProductFeedUpload](CreateProductFeedUploadsBatchCall(id, params, options...))
}

func DecodeCreateProductFeedUploadsBatchResponse(response *core.BatchResponse) (*objects.ProductFeedUpload, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedUpload
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateProductFeedUploads(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadsParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	call := CreateProductFeedUploadsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedVehicleOffersBatchCall(id string, params GetProductFeedVehicleOffersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "vehicle_offers"), params.ToParams(), options...)
}

func NewGetProductFeedVehicleOffersBatchRequest(id string, params GetProductFeedVehicleOffersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VehicleOffer]] {
	return core.NewBatchRequest[core.Cursor[objects.VehicleOffer]](GetProductFeedVehicleOffersBatchCall(id, params, options...))
}

func DecodeGetProductFeedVehicleOffersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VehicleOffer], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VehicleOffer]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedVehicleOffers(ctx context.Context, client *core.Client, id string, params GetProductFeedVehicleOffersParams) (*core.Cursor[objects.VehicleOffer], error) {
	var out core.Cursor[objects.VehicleOffer]
	call := GetProductFeedVehicleOffersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedVehiclesBatchCall(id string, params GetProductFeedVehiclesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "vehicles"), params.ToParams(), options...)
}

func NewGetProductFeedVehiclesBatchRequest(id string, params GetProductFeedVehiclesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Vehicle]] {
	return core.NewBatchRequest[core.Cursor[objects.Vehicle]](GetProductFeedVehiclesBatchCall(id, params, options...))
}

func DecodeGetProductFeedVehiclesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Vehicle], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Vehicle]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedVehicles(ctx context.Context, client *core.Client, id string, params GetProductFeedVehiclesParams) (*core.Cursor[objects.Vehicle], error) {
	var out core.Cursor[objects.Vehicle]
	call := GetProductFeedVehiclesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func DeleteProductFeedBatchCall(id string, params DeleteProductFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteProductFeedBatchRequest(id string, params DeleteProductFeedParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteProductFeedBatchCall(id, params, options...))
}

func DecodeDeleteProductFeedBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteProductFeed(ctx context.Context, client *core.Client, id string, params DeleteProductFeedParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteProductFeedBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetProductFeedBatchCall(id string, params GetProductFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedBatchRequest(id string, params GetProductFeedParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeed] {
	return core.NewBatchRequest[objects.ProductFeed](GetProductFeedBatchCall(id, params, options...))
}

func DecodeGetProductFeedBatchResponse(response *core.BatchResponse) (*objects.ProductFeed, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeed
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeed(ctx context.Context, client *core.Client, id string, params GetProductFeedParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	call := GetProductFeedBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateProductFeedBatchCall(id string, params UpdateProductFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateProductFeedBatchRequest(id string, params UpdateProductFeedParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeed] {
	return core.NewBatchRequest[objects.ProductFeed](UpdateProductFeedBatchCall(id, params, options...))
}

func DecodeUpdateProductFeedBatchResponse(response *core.BatchResponse) (*objects.ProductFeed, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeed
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateProductFeed(ctx context.Context, client *core.Client, id string, params UpdateProductFeedParams) (*objects.ProductFeed, error) {
	var out objects.ProductFeed
	call := UpdateProductFeedBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
