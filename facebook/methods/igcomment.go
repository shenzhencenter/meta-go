package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGCommentRepliesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGCommentRepliesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGCommentRepliesBatchCall(id string, params GetIGCommentRepliesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "replies"), params.ToParams(), options...)
}

func NewGetIGCommentRepliesBatchRequest(id string, params GetIGCommentRepliesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGComment]] {
	return core.NewBatchRequest[core.Cursor[objects.IGComment]](GetIGCommentRepliesBatchCall(id, params, options...))
}

func DecodeGetIGCommentRepliesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGComment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGComment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGCommentReplies(ctx context.Context, client *core.Client, id string, params GetIGCommentRepliesParams) (*core.Cursor[objects.IGComment], error) {
	var out core.Cursor[objects.IGComment]
	call := GetIGCommentRepliesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateIGCommentRepliesParams struct {
	Message *string     `facebook:"message"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGCommentRepliesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateIGCommentRepliesBatchCall(id string, params CreateIGCommentRepliesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "replies"), params.ToParams(), options...)
}

func NewCreateIGCommentRepliesBatchRequest(id string, params CreateIGCommentRepliesParams, options ...core.BatchOption) *core.BatchRequest[objects.IGComment] {
	return core.NewBatchRequest[objects.IGComment](CreateIGCommentRepliesBatchCall(id, params, options...))
}

func DecodeCreateIGCommentRepliesBatchResponse(response *core.BatchResponse) (*objects.IGComment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGComment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGCommentReplies(ctx context.Context, client *core.Client, id string, params CreateIGCommentRepliesParams) (*objects.IGComment, error) {
	var out objects.IGComment
	call := CreateIGCommentRepliesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteIGCommentParams struct {
	AdID  *core.ID    `facebook:"ad_id"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	return out
}

func DeleteIGCommentBatchCall(id string, params DeleteIGCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteIGCommentBatchRequest(id string, params DeleteIGCommentParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGCommentBatchCall(id, params, options...))
}

func DecodeDeleteIGCommentBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteIGComment(ctx context.Context, client *core.Client, id string, params DeleteIGCommentParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteIGCommentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGCommentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGCommentBatchCall(id string, params GetIGCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGCommentBatchRequest(id string, params GetIGCommentParams, options ...core.BatchOption) *core.BatchRequest[objects.IGComment] {
	return core.NewBatchRequest[objects.IGComment](GetIGCommentBatchCall(id, params, options...))
}

func DecodeGetIGCommentBatchResponse(response *core.BatchResponse) (*objects.IGComment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGComment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGComment(ctx context.Context, client *core.Client, id string, params GetIGCommentParams) (*objects.IGComment, error) {
	var out objects.IGComment
	call := GetIGCommentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateIGCommentParams struct {
	AdID  *core.ID    `facebook:"ad_id"`
	Hide  bool        `facebook:"hide"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	out["hide"] = params.Hide
	return out
}

func UpdateIGCommentBatchCall(id string, params UpdateIGCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateIGCommentBatchRequest(id string, params UpdateIGCommentParams, options ...core.BatchOption) *core.BatchRequest[objects.IGComment] {
	return core.NewBatchRequest[objects.IGComment](UpdateIGCommentBatchCall(id, params, options...))
}

func DecodeUpdateIGCommentBatchResponse(response *core.BatchResponse) (*objects.IGComment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGComment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateIGComment(ctx context.Context, client *core.Client, id string, params UpdateIGCommentParams) (*objects.IGComment, error) {
	var out objects.IGComment
	call := UpdateIGCommentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
