package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLocalServiceBusinessChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLocalServiceBusinessChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLocalServiceBusinessChannelsToIntegrityStatusBatchCall(id string, params GetLocalServiceBusinessChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetLocalServiceBusinessChannelsToIntegrityStatusBatchRequest(id string, params GetLocalServiceBusinessChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetLocalServiceBusinessChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetLocalServiceBusinessChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetLocalServiceBusinessChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetLocalServiceBusinessChannelsToIntegrityStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLocalServiceBusinessOverrideDetailsParams struct {
	Keys  *[]string                                               `facebook:"keys"`
	Type  *enums.LocalservicebusinessoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                             `facebook:"-"`
}

func (params GetLocalServiceBusinessOverrideDetailsParams) ToParams() core.Params {
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

func GetLocalServiceBusinessOverrideDetailsBatchCall(id string, params GetLocalServiceBusinessOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetLocalServiceBusinessOverrideDetailsBatchRequest(id string, params GetLocalServiceBusinessOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetLocalServiceBusinessOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetLocalServiceBusinessOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetLocalServiceBusinessOverrideDetails(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetLocalServiceBusinessOverrideDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLocalServiceBusinessParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLocalServiceBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLocalServiceBusinessBatchCall(id string, params GetLocalServiceBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLocalServiceBusinessBatchRequest(id string, params GetLocalServiceBusinessParams, options ...core.BatchOption) *core.BatchRequest[objects.LocalServiceBusiness] {
	return core.NewBatchRequest[objects.LocalServiceBusiness](GetLocalServiceBusinessBatchCall(id, params, options...))
}

func DecodeGetLocalServiceBusinessBatchResponse(response *core.BatchResponse) (*objects.LocalServiceBusiness, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LocalServiceBusiness
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLocalServiceBusiness(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessParams) (*objects.LocalServiceBusiness, error) {
	var out objects.LocalServiceBusiness
	call := GetLocalServiceBusinessBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
