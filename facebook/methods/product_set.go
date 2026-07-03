package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductSetAutomotiveModelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetAutomotiveModelsParams) ToParams() core.Params {
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

func GetProductSetAutomotiveModelsBatchCall(id string, params GetProductSetAutomotiveModelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "automotive_models"), params.ToParams(), options...)
}

func NewGetProductSetAutomotiveModelsBatchRequest(id string, params GetProductSetAutomotiveModelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AutomotiveModel]] {
	return core.NewBatchRequest[core.Cursor[objects.AutomotiveModel]](GetProductSetAutomotiveModelsBatchCall(id, params, options...))
}

func DecodeGetProductSetAutomotiveModelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AutomotiveModel], error) {
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

func GetProductSetAutomotiveModels(ctx context.Context, client *core.Client, id string, params GetProductSetAutomotiveModelsParams) (*core.Cursor[objects.AutomotiveModel], error) {
	var out core.Cursor[objects.AutomotiveModel]
	call := GetProductSetAutomotiveModelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetDestinationsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetDestinationsParams) ToParams() core.Params {
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

func GetProductSetDestinationsBatchCall(id string, params GetProductSetDestinationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "destinations"), params.ToParams(), options...)
}

func NewGetProductSetDestinationsBatchRequest(id string, params GetProductSetDestinationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Destination]] {
	return core.NewBatchRequest[core.Cursor[objects.Destination]](GetProductSetDestinationsBatchCall(id, params, options...))
}

func DecodeGetProductSetDestinationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Destination], error) {
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

func GetProductSetDestinations(ctx context.Context, client *core.Client, id string, params GetProductSetDestinationsParams) (*core.Cursor[objects.Destination], error) {
	var out core.Cursor[objects.Destination]
	call := GetProductSetDestinationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetFlightsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetFlightsParams) ToParams() core.Params {
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

func GetProductSetFlightsBatchCall(id string, params GetProductSetFlightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "flights"), params.ToParams(), options...)
}

func NewGetProductSetFlightsBatchRequest(id string, params GetProductSetFlightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Flight]] {
	return core.NewBatchRequest[core.Cursor[objects.Flight]](GetProductSetFlightsBatchCall(id, params, options...))
}

func DecodeGetProductSetFlightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Flight], error) {
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

func GetProductSetFlights(ctx context.Context, client *core.Client, id string, params GetProductSetFlightsParams) (*core.Cursor[objects.Flight], error) {
	var out core.Cursor[objects.Flight]
	call := GetProductSetFlightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetHomeListingsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetHomeListingsParams) ToParams() core.Params {
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

func GetProductSetHomeListingsBatchCall(id string, params GetProductSetHomeListingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "home_listings"), params.ToParams(), options...)
}

func NewGetProductSetHomeListingsBatchRequest(id string, params GetProductSetHomeListingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.HomeListing]] {
	return core.NewBatchRequest[core.Cursor[objects.HomeListing]](GetProductSetHomeListingsBatchCall(id, params, options...))
}

func DecodeGetProductSetHomeListingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.HomeListing], error) {
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

func GetProductSetHomeListings(ctx context.Context, client *core.Client, id string, params GetProductSetHomeListingsParams) (*core.Cursor[objects.HomeListing], error) {
	var out core.Cursor[objects.HomeListing]
	call := GetProductSetHomeListingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetHotelsParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetHotelsParams) ToParams() core.Params {
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

func GetProductSetHotelsBatchCall(id string, params GetProductSetHotelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "hotels"), params.ToParams(), options...)
}

func NewGetProductSetHotelsBatchRequest(id string, params GetProductSetHotelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Hotel]] {
	return core.NewBatchRequest[core.Cursor[objects.Hotel]](GetProductSetHotelsBatchCall(id, params, options...))
}

func DecodeGetProductSetHotelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Hotel], error) {
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

func GetProductSetHotels(ctx context.Context, client *core.Client, id string, params GetProductSetHotelsParams) (*core.Cursor[objects.Hotel], error) {
	var out core.Cursor[objects.Hotel]
	call := GetProductSetHotelsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetProductsParams struct {
	BulkPagination *bool                                           `facebook:"bulk_pagination"`
	ErrorPriority  *enums.ProductsetproductsErrorPriorityEnumParam `facebook:"error_priority"`
	ErrorType      *enums.ProductsetproductsErrorTypeEnumParam     `facebook:"error_type"`
	Filter         *map[string]interface{}                         `facebook:"filter"`
	Extra          core.Params                                     `facebook:"-"`
}

func (params GetProductSetProductsParams) ToParams() core.Params {
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

func GetProductSetProductsBatchCall(id string, params GetProductSetProductsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), options...)
}

func NewGetProductSetProductsBatchRequest(id string, params GetProductSetProductsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductItem]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductItem]](GetProductSetProductsBatchCall(id, params, options...))
}

func DecodeGetProductSetProductsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductItem], error) {
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

func GetProductSetProducts(ctx context.Context, client *core.Client, id string, params GetProductSetProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	call := GetProductSetProductsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetVehicleOffersParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetVehicleOffersParams) ToParams() core.Params {
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

func GetProductSetVehicleOffersBatchCall(id string, params GetProductSetVehicleOffersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "vehicle_offers"), params.ToParams(), options...)
}

func NewGetProductSetVehicleOffersBatchRequest(id string, params GetProductSetVehicleOffersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VehicleOffer]] {
	return core.NewBatchRequest[core.Cursor[objects.VehicleOffer]](GetProductSetVehicleOffersBatchCall(id, params, options...))
}

func DecodeGetProductSetVehicleOffersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VehicleOffer], error) {
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

func GetProductSetVehicleOffers(ctx context.Context, client *core.Client, id string, params GetProductSetVehicleOffersParams) (*core.Cursor[objects.VehicleOffer], error) {
	var out core.Cursor[objects.VehicleOffer]
	call := GetProductSetVehicleOffersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetVehiclesParams struct {
	BulkPagination *bool                   `facebook:"bulk_pagination"`
	Filter         *map[string]interface{} `facebook:"filter"`
	Extra          core.Params             `facebook:"-"`
}

func (params GetProductSetVehiclesParams) ToParams() core.Params {
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

func GetProductSetVehiclesBatchCall(id string, params GetProductSetVehiclesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "vehicles"), params.ToParams(), options...)
}

func NewGetProductSetVehiclesBatchRequest(id string, params GetProductSetVehiclesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Vehicle]] {
	return core.NewBatchRequest[core.Cursor[objects.Vehicle]](GetProductSetVehiclesBatchCall(id, params, options...))
}

func DecodeGetProductSetVehiclesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Vehicle], error) {
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

func GetProductSetVehicles(ctx context.Context, client *core.Client, id string, params GetProductSetVehiclesParams) (*core.Cursor[objects.Vehicle], error) {
	var out core.Cursor[objects.Vehicle]
	call := GetProductSetVehiclesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteProductSetParams struct {
	AllowLiveProductSetDeletion *bool       `facebook:"allow_live_product_set_deletion"`
	Extra                       core.Params `facebook:"-"`
}

func (params DeleteProductSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AllowLiveProductSetDeletion != nil {
		out["allow_live_product_set_deletion"] = *params.AllowLiveProductSetDeletion
	}
	return out
}

func DeleteProductSetBatchCall(id string, params DeleteProductSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteProductSetBatchRequest(id string, params DeleteProductSetParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteProductSetBatchCall(id, params, options...))
}

func DecodeDeleteProductSetBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteProductSet(ctx context.Context, client *core.Client, id string, params DeleteProductSetParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteProductSetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductSetBatchCall(id string, params GetProductSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductSetBatchRequest(id string, params GetProductSetParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductSet] {
	return core.NewBatchRequest[objects.ProductSet](GetProductSetBatchCall(id, params, options...))
}

func DecodeGetProductSetBatchResponse(response *core.BatchResponse) (*objects.ProductSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductSet(ctx context.Context, client *core.Client, id string, params GetProductSetParams) (*objects.ProductSet, error) {
	var out objects.ProductSet
	call := GetProductSetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateProductSetParams struct {
	Filter         *map[string]interface{}   `facebook:"filter"`
	Metadata       *map[string]interface{}   `facebook:"metadata"`
	Name           *string                   `facebook:"name"`
	OrderingInfo   *[]uint64                 `facebook:"ordering_info"`
	PublishToShops *[]map[string]interface{} `facebook:"publish_to_shops"`
	RetailerID     *core.ID                  `facebook:"retailer_id"`
	Extra          core.Params               `facebook:"-"`
}

func (params UpdateProductSetParams) ToParams() core.Params {
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
	if params.Name != nil {
		out["name"] = *params.Name
	}
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

func UpdateProductSetBatchCall(id string, params UpdateProductSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateProductSetBatchRequest(id string, params UpdateProductSetParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductSet] {
	return core.NewBatchRequest[objects.ProductSet](UpdateProductSetBatchCall(id, params, options...))
}

func DecodeUpdateProductSetBatchResponse(response *core.BatchResponse) (*objects.ProductSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateProductSet(ctx context.Context, client *core.Client, id string, params UpdateProductSetParams) (*objects.ProductSet, error) {
	var out objects.ProductSet
	call := UpdateProductSetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
