package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func CreateProductFeedUploadErrorReport(ctx context.Context, client *core.Client, id string, params CreateProductFeedUploadErrorReportParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "error_report"), params.ToParams(), &out)
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

func GetProductFeedUploadErrors(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorsParams) (*core.Cursor[objects.ProductFeedUploadError], error) {
	var out core.Cursor[objects.ProductFeedUploadError]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "errors"), params.ToParams(), &out)
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

func GetProductFeedUpload(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadParams) (*objects.ProductFeedUpload, error) {
	var out objects.ProductFeedUpload
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
