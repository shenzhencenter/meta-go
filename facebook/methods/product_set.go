package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetProductSetAutomotiveModels(ctx context.Context, client *core.Client, id string, params GetProductSetAutomotiveModelsParams) (*core.Cursor[objects.AutomotiveModel], error) {
	var out core.Cursor[objects.AutomotiveModel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "automotive_models"), params.ToParams(), &out)
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

func GetProductSetDestinations(ctx context.Context, client *core.Client, id string, params GetProductSetDestinationsParams) (*core.Cursor[objects.Destination], error) {
	var out core.Cursor[objects.Destination]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "destinations"), params.ToParams(), &out)
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

func GetProductSetFlights(ctx context.Context, client *core.Client, id string, params GetProductSetFlightsParams) (*core.Cursor[objects.Flight], error) {
	var out core.Cursor[objects.Flight]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "flights"), params.ToParams(), &out)
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

func GetProductSetHomeListings(ctx context.Context, client *core.Client, id string, params GetProductSetHomeListingsParams) (*core.Cursor[objects.HomeListing], error) {
	var out core.Cursor[objects.HomeListing]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "home_listings"), params.ToParams(), &out)
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

func GetProductSetHotels(ctx context.Context, client *core.Client, id string, params GetProductSetHotelsParams) (*core.Cursor[objects.Hotel], error) {
	var out core.Cursor[objects.Hotel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "hotels"), params.ToParams(), &out)
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

func GetProductSetProducts(ctx context.Context, client *core.Client, id string, params GetProductSetProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), &out)
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

func GetProductSetVehicleOffers(ctx context.Context, client *core.Client, id string, params GetProductSetVehicleOffersParams) (*core.Cursor[objects.VehicleOffer], error) {
	var out core.Cursor[objects.VehicleOffer]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicle_offers"), params.ToParams(), &out)
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

func GetProductSetVehicles(ctx context.Context, client *core.Client, id string, params GetProductSetVehiclesParams) (*core.Cursor[objects.Vehicle], error) {
	var out core.Cursor[objects.Vehicle]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "vehicles"), params.ToParams(), &out)
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

func DeleteProductSet(ctx context.Context, client *core.Client, id string, params DeleteProductSetParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetProductSet(ctx context.Context, client *core.Client, id string, params GetProductSetParams) (*objects.ProductSet, error) {
	var out objects.ProductSet
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateProductSet(ctx context.Context, client *core.Client, id string, params UpdateProductSetParams) (*objects.ProductSet, error) {
	var out objects.ProductSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
