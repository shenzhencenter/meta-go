package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetTransactableItemChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetTransactableItemChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetTransactableItemChannelsToIntegrityStatusBatchCall(id string, params GetTransactableItemChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetTransactableItemChannelsToIntegrityStatusBatchRequest(id string, params GetTransactableItemChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetTransactableItemChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetTransactableItemChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetTransactableItemChannelsToIntegrityStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetTransactableItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], *core.Response, error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetTransactableItemChannelsToIntegrityStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetTransactableItemChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetTransactableItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	out, _, err := GetTransactableItemChannelsToIntegrityStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetTransactableItemOverrideDetailsParams struct {
	Keys  *[]string                                           `facebook:"keys"`
	Type  *enums.TransactableitemoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                         `facebook:"-"`
}

func (params GetTransactableItemOverrideDetailsParams) ToParams() core.Params {
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

func GetTransactableItemOverrideDetailsBatchCall(id string, params GetTransactableItemOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetTransactableItemOverrideDetailsBatchRequest(id string, params GetTransactableItemOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetTransactableItemOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetTransactableItemOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetTransactableItemOverrideDetailsWithResponse(ctx context.Context, client *core.Client, id string, params GetTransactableItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], *core.Response, error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetTransactableItemOverrideDetailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetTransactableItemOverrideDetails(ctx context.Context, client *core.Client, id string, params GetTransactableItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	out, _, err := GetTransactableItemOverrideDetailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetTransactableItemParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetTransactableItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetTransactableItemBatchCall(id string, params GetTransactableItemParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetTransactableItemBatchRequest(id string, params GetTransactableItemParams, options ...core.BatchOption) *core.BatchRequest[objects.TransactableItem] {
	return core.NewBatchRequest[objects.TransactableItem](GetTransactableItemBatchCall(id, params, options...))
}

func DecodeGetTransactableItemBatchResponse(response *core.BatchResponse) (*objects.TransactableItem, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.TransactableItem
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetTransactableItemWithResponse(ctx context.Context, client *core.Client, id string, params GetTransactableItemParams) (*objects.TransactableItem, *core.Response, error) {
	var out objects.TransactableItem
	call := GetTransactableItemBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetTransactableItem(ctx context.Context, client *core.Client, id string, params GetTransactableItemParams) (*objects.TransactableItem, error) {
	out, _, err := GetTransactableItemWithResponse(ctx, client, id, params)
	return out, err
}
