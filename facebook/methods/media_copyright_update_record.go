package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMediaCopyrightUpdateRecordParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMediaCopyrightUpdateRecordParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMediaCopyrightUpdateRecordBatchCall(id string, params GetMediaCopyrightUpdateRecordParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMediaCopyrightUpdateRecordBatchRequest(id string, params GetMediaCopyrightUpdateRecordParams, options ...core.BatchOption) *core.BatchRequest[objects.MediaCopyrightUpdateRecord] {
	return core.NewBatchRequest[objects.MediaCopyrightUpdateRecord](GetMediaCopyrightUpdateRecordBatchCall(id, params, options...))
}

func DecodeGetMediaCopyrightUpdateRecordBatchResponse(response *core.BatchResponse) (*objects.MediaCopyrightUpdateRecord, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MediaCopyrightUpdateRecord
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMediaCopyrightUpdateRecord(ctx context.Context, client *core.Client, id string, params GetMediaCopyrightUpdateRecordParams) (*objects.MediaCopyrightUpdateRecord, error) {
	var out objects.MediaCopyrightUpdateRecord
	call := GetMediaCopyrightUpdateRecordBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
