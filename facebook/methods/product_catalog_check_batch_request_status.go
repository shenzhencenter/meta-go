package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams struct {
	After                    *string                                                               `facebook:"after"`
	Before                   *string                                                               `facebook:"before"`
	ErrorPriority            *enums.ProductcatalogcheckbatchrequeststatusgetErrorPriorityEnumParam `facebook:"error_priority"`
	Fields                   *string                                                               `facebook:"fields"`
	Handle                   string                                                                `facebook:"handle"`
	Limit                    *int                                                                  `facebook:"limit"`
	LoadIdsOfInvalidRequests *bool                                                                 `facebook:"load_ids_of_invalid_requests"`
	Extra                    core.Params                                                           `facebook:"-"`
}

func (params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.After != nil {
		out["after"] = *params.After
	}
	if params.Before != nil {
		out["before"] = *params.Before
	}
	if params.ErrorPriority != nil {
		out["error_priority"] = *params.ErrorPriority
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	out["handle"] = params.Handle
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.LoadIdsOfInvalidRequests != nil {
		out["load_ids_of_invalid_requests"] = *params.LoadIdsOfInvalidRequests
	}
	return out
}

func GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusBatchCall(id string, params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "check_batch_request_status"), params.ToParams(), options...)
}

func NewGetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusBatchRequest(id string, params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet]](GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusBatchCall(id, params, options...))
}

func DecodeGetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams) (*core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet], *core.Response, error) {
	var out core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet]
	call := GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatus(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams) (*core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet], error) {
	out, _, err := GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusWithResponse(ctx, client, id, params)
	return out, err
}
