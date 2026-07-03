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

func GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatus(ctx context.Context, client *core.Client, id string, params GetProductCatalogCheckBatchRequestStatusCheckBatchRequestStatusParams) (*core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet], error) {
	var out core.Cursor[objects.ProductCatalogCheckBatchRequestStatusGet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "check_batch_request_status"), params.ToParams(), &out)
	return &out, err
}
