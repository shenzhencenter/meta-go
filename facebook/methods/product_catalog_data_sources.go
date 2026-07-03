package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductCatalogDataSourcesDataSourcesParams struct {
	After               *string                                                         `facebook:"after"`
	Before              *string                                                         `facebook:"before"`
	Fields              *string                                                         `facebook:"fields"`
	IngestionSourceType *enums.ProductcatalogdatasourcesgetIngestionSourceTypeEnumParam `facebook:"ingestion_source_type"`
	Limit               *int                                                            `facebook:"limit"`
	Extra               core.Params                                                     `facebook:"-"`
}

func (params GetProductCatalogDataSourcesDataSourcesParams) ToParams() core.Params {
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
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.IngestionSourceType != nil {
		out["ingestion_source_type"] = *params.IngestionSourceType
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	return out
}

func GetProductCatalogDataSourcesDataSourcesBatchCall(id string, params GetProductCatalogDataSourcesDataSourcesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "data_sources"), params.ToParams(), options...)
}

func NewGetProductCatalogDataSourcesDataSourcesBatchRequest(id string, params GetProductCatalogDataSourcesDataSourcesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalogDataSourcesGet]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalogDataSourcesGet]](GetProductCatalogDataSourcesDataSourcesBatchCall(id, params, options...))
}

func DecodeGetProductCatalogDataSourcesDataSourcesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalogDataSourcesGet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalogDataSourcesGet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductCatalogDataSourcesDataSources(ctx context.Context, client *core.Client, id string, params GetProductCatalogDataSourcesDataSourcesParams) (*core.Cursor[objects.ProductCatalogDataSourcesGet], error) {
	var out core.Cursor[objects.ProductCatalogDataSourcesGet]
	call := GetProductCatalogDataSourcesDataSourcesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
