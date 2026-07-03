package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetRightsManagerDataExportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRightsManagerDataExportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRightsManagerDataExportBatchCall(id string, params GetRightsManagerDataExportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetRightsManagerDataExportBatchRequest(id string, params GetRightsManagerDataExportParams, options ...core.BatchOption) *core.BatchRequest[objects.RightsManagerDataExport] {
	return core.NewBatchRequest[objects.RightsManagerDataExport](GetRightsManagerDataExportBatchCall(id, params, options...))
}

func DecodeGetRightsManagerDataExportBatchResponse(response *core.BatchResponse) (*objects.RightsManagerDataExport, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.RightsManagerDataExport
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetRightsManagerDataExport(ctx context.Context, client *core.Client, id string, params GetRightsManagerDataExportParams) (*objects.RightsManagerDataExport, error) {
	var out objects.RightsManagerDataExport
	call := GetRightsManagerDataExportBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
