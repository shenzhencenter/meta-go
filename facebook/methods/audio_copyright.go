package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAudioCopyrightUpdateRecordsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioCopyrightUpdateRecordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioCopyrightUpdateRecordsBatchCall(id string, params GetAudioCopyrightUpdateRecordsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "update_records"), params.ToParams(), options...)
}

func NewGetAudioCopyrightUpdateRecordsBatchRequest(id string, params GetAudioCopyrightUpdateRecordsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MediaCopyrightUpdateRecord]] {
	return core.NewBatchRequest[core.Cursor[objects.MediaCopyrightUpdateRecord]](GetAudioCopyrightUpdateRecordsBatchCall(id, params, options...))
}

func DecodeGetAudioCopyrightUpdateRecordsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAudioCopyrightUpdateRecordsWithResponse(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightUpdateRecordsParams) (*core.Cursor[objects.MediaCopyrightUpdateRecord], *core.Response, error) {
	var out core.Cursor[objects.MediaCopyrightUpdateRecord]
	call := GetAudioCopyrightUpdateRecordsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAudioCopyrightUpdateRecords(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightUpdateRecordsParams) (*core.Cursor[objects.MediaCopyrightUpdateRecord], error) {
	out, _, err := GetAudioCopyrightUpdateRecordsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAudioCopyrightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAudioCopyrightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAudioCopyrightBatchCall(id string, params GetAudioCopyrightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAudioCopyrightBatchRequest(id string, params GetAudioCopyrightParams, options ...core.BatchOption) *core.BatchRequest[objects.AudioCopyright] {
	return core.NewBatchRequest[objects.AudioCopyright](GetAudioCopyrightBatchCall(id, params, options...))
}

func DecodeGetAudioCopyrightBatchResponse(response *core.BatchResponse) (*objects.AudioCopyright, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AudioCopyright
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAudioCopyrightWithResponse(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightParams) (*objects.AudioCopyright, *core.Response, error) {
	var out objects.AudioCopyright
	call := GetAudioCopyrightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAudioCopyright(ctx context.Context, client *core.Client, id string, params GetAudioCopyrightParams) (*objects.AudioCopyright, error) {
	out, _, err := GetAudioCopyrightWithResponse(ctx, client, id, params)
	return out, err
}
