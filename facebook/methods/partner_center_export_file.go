package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPartnerCenterExportFileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerCenterExportFileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerCenterExportFileBatchCall(id string, params GetPartnerCenterExportFileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPartnerCenterExportFileBatchRequest(id string, params GetPartnerCenterExportFileParams, options ...core.BatchOption) *core.BatchRequest[objects.PartnerCenterExportFile] {
	return core.NewBatchRequest[objects.PartnerCenterExportFile](GetPartnerCenterExportFileBatchCall(id, params, options...))
}

func DecodeGetPartnerCenterExportFileBatchResponse(response *core.BatchResponse) (*objects.PartnerCenterExportFile, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PartnerCenterExportFile
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPartnerCenterExportFile(ctx context.Context, client *core.Client, id string, params GetPartnerCenterExportFileParams) (*objects.PartnerCenterExportFile, error) {
	var out objects.PartnerCenterExportFile
	call := GetPartnerCenterExportFileBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
