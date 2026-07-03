package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetUnifiedThreadMessagesParams struct {
	Source *enums.UnifiedthreadmessagesSourceEnumParam `facebook:"source"`
	Extra  core.Params                                 `facebook:"-"`
}

func (params GetUnifiedThreadMessagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetUnifiedThreadMessagesBatchCall(id string, params GetUnifiedThreadMessagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messages"), params.ToParams(), options...)
}

func NewGetUnifiedThreadMessagesBatchRequest(id string, params GetUnifiedThreadMessagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UnifiedMessage]] {
	return core.NewBatchRequest[core.Cursor[objects.UnifiedMessage]](GetUnifiedThreadMessagesBatchCall(id, params, options...))
}

func DecodeGetUnifiedThreadMessagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UnifiedMessage], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UnifiedMessage]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUnifiedThreadMessages(ctx context.Context, client *core.Client, id string, params GetUnifiedThreadMessagesParams) (*core.Cursor[objects.UnifiedMessage], error) {
	var out core.Cursor[objects.UnifiedMessage]
	call := GetUnifiedThreadMessagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetUnifiedThreadParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUnifiedThreadParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUnifiedThreadBatchCall(id string, params GetUnifiedThreadParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetUnifiedThreadBatchRequest(id string, params GetUnifiedThreadParams, options ...core.BatchOption) *core.BatchRequest[objects.UnifiedThread] {
	return core.NewBatchRequest[objects.UnifiedThread](GetUnifiedThreadBatchCall(id, params, options...))
}

func DecodeGetUnifiedThreadBatchResponse(response *core.BatchResponse) (*objects.UnifiedThread, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.UnifiedThread
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUnifiedThread(ctx context.Context, client *core.Client, id string, params GetUnifiedThreadParams) (*objects.UnifiedThread, error) {
	var out objects.UnifiedThread
	call := GetUnifiedThreadBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
