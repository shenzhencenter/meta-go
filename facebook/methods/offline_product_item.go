package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOfflineProductItemChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineProductItemChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineProductItemChannelsToIntegrityStatusBatchCall(id string, params GetOfflineProductItemChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetOfflineProductItemChannelsToIntegrityStatusBatchRequest(id string, params GetOfflineProductItemChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetOfflineProductItemChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetOfflineProductItemChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineProductItemChannelsToIntegrityStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], *core.Response, error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetOfflineProductItemChannelsToIntegrityStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineProductItemChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	out, _, err := GetOfflineProductItemChannelsToIntegrityStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineProductItemOverrideDetailsParams struct {
	Keys  *[]string                                             `facebook:"keys"`
	Type  *enums.OfflineproductitemoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                           `facebook:"-"`
}

func (params GetOfflineProductItemOverrideDetailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Keys != nil {
		out["keys"] = *params.Keys
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetOfflineProductItemOverrideDetailsBatchCall(id string, params GetOfflineProductItemOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetOfflineProductItemOverrideDetailsBatchRequest(id string, params GetOfflineProductItemOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetOfflineProductItemOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetOfflineProductItemOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OverrideDetails]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineProductItemOverrideDetailsWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], *core.Response, error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetOfflineProductItemOverrideDetailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineProductItemOverrideDetails(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	out, _, err := GetOfflineProductItemOverrideDetailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetOfflineProductItemParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineProductItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineProductItemBatchCall(id string, params GetOfflineProductItemParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOfflineProductItemBatchRequest(id string, params GetOfflineProductItemParams, options ...core.BatchOption) *core.BatchRequest[objects.OfflineProductItem] {
	return core.NewBatchRequest[objects.OfflineProductItem](GetOfflineProductItemBatchCall(id, params, options...))
}

func DecodeGetOfflineProductItemBatchResponse(response *core.BatchResponse) (*objects.OfflineProductItem, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OfflineProductItem
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOfflineProductItemWithResponse(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemParams) (*objects.OfflineProductItem, *core.Response, error) {
	var out objects.OfflineProductItem
	call := GetOfflineProductItemBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOfflineProductItem(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemParams) (*objects.OfflineProductItem, error) {
	out, _, err := GetOfflineProductItemWithResponse(ctx, client, id, params)
	return out, err
}
