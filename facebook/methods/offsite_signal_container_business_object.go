package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObjectLinkedApplicationBatchCall(id string, params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "linked_application"), params.ToParams(), options...)
}

func NewGetOffsiteSignalContainerBusinessObjectLinkedApplicationBatchRequest(id string, params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsDataset]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsDataset]](GetOffsiteSignalContainerBusinessObjectLinkedApplicationBatchCall(id, params, options...))
}

func DecodeGetOffsiteSignalContainerBusinessObjectLinkedApplicationBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsDataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsDataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOffsiteSignalContainerBusinessObjectLinkedApplication(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectLinkedApplicationParams) (*core.Cursor[objects.AdsDataset], error) {
	var out core.Cursor[objects.AdsDataset]
	call := GetOffsiteSignalContainerBusinessObjectLinkedApplicationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetOffsiteSignalContainerBusinessObjectLinkedPageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectLinkedPageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObjectLinkedPageBatchCall(id string, params GetOffsiteSignalContainerBusinessObjectLinkedPageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "linked_page"), params.ToParams(), options...)
}

func NewGetOffsiteSignalContainerBusinessObjectLinkedPageBatchRequest(id string, params GetOffsiteSignalContainerBusinessObjectLinkedPageParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetOffsiteSignalContainerBusinessObjectLinkedPageBatchCall(id, params, options...))
}

func DecodeGetOffsiteSignalContainerBusinessObjectLinkedPageBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOffsiteSignalContainerBusinessObjectLinkedPage(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectLinkedPageParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetOffsiteSignalContainerBusinessObjectLinkedPageBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetOffsiteSignalContainerBusinessObjectParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOffsiteSignalContainerBusinessObjectParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOffsiteSignalContainerBusinessObjectBatchCall(id string, params GetOffsiteSignalContainerBusinessObjectParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOffsiteSignalContainerBusinessObjectBatchRequest(id string, params GetOffsiteSignalContainerBusinessObjectParams, options ...core.BatchOption) *core.BatchRequest[objects.OffsiteSignalContainerBusinessObject] {
	return core.NewBatchRequest[objects.OffsiteSignalContainerBusinessObject](GetOffsiteSignalContainerBusinessObjectBatchCall(id, params, options...))
}

func DecodeGetOffsiteSignalContainerBusinessObjectBatchResponse(response *core.BatchResponse) (*objects.OffsiteSignalContainerBusinessObject, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OffsiteSignalContainerBusinessObject
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOffsiteSignalContainerBusinessObject(ctx context.Context, client *core.Client, id string, params GetOffsiteSignalContainerBusinessObjectParams) (*objects.OffsiteSignalContainerBusinessObject, error) {
	var out objects.OffsiteSignalContainerBusinessObject
	call := GetOffsiteSignalContainerBusinessObjectBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
