package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateProductFeedUploadErrorReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateProductFeedUploadErrorReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateProductFeedUploadErrorReportBatchCall(id string, params CreateProductFeedUploadErrorReportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "error_report"), params.ToParams(), options...)
}

func NewCreateProductFeedUploadErrorReportBatchRequest(id string, params CreateProductFeedUploadErrorReportParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedUpload] {
	return core.NewBatchRequest[objects.ProductFeedUpload](CreateProductFeedUploadErrorReportBatchCall(id, params, options...))
}

func DecodeCreateProductFeedUploadErrorReportBatchResponse(response *core.BatchResponse) (*objects.ProductFeedUpload, error) {
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

func CreateProductFeedUploadErrorReport(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadErrorReportParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	call := CreateProductFeedUploadErrorReportBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductFeedUploadErrorsParams struct {
	ErrorPriority *enums.ProductfeeduploaderrorsErrorPriorityEnumParam `facebook:"error_priority"`
	Extra         core.Params                                          `facebook:"-"`
}

func (params GetProductFeedUploadErrorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ErrorPriority != nil {
		out["error_priority"] = *params.ErrorPriority
	}
	return out
}

func GetProductFeedUploadErrorsBatchCall(id string, params GetProductFeedUploadErrorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "errors"), params.ToParams(), options...)
}

func NewGetProductFeedUploadErrorsBatchRequest(id string, params GetProductFeedUploadErrorsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedUploadError]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedUploadError]](GetProductFeedUploadErrorsBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadErrorsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedUploadError], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedUploadError]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadErrors(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorsParams) (*core.Cursor[objects.ProductFeedUploadError], error) {
	var out core.Cursor[objects.ProductFeedUploadError]
	call := GetProductFeedUploadErrorsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductFeedUploadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadBatchCall(id string, params GetProductFeedUploadParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedUploadBatchRequest(id string, params GetProductFeedUploadParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedUpload] {
	return core.NewBatchRequest[objects.ProductFeedUpload](GetProductFeedUploadBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadBatchResponse(response *core.BatchResponse) (*objects.ProductFeedUpload, error) {
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

func GetProductFeedUpload(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	call := GetProductFeedUploadBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
