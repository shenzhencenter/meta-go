package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreatePublisherBlockListAppendPublisherUrlsParams struct {
	PublisherUrls []string    `facebook:"publisher_urls"`
	Extra         core.Params `facebook:"-"`
}

func (params CreatePublisherBlockListAppendPublisherUrlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["publisher_urls"] = params.PublisherUrls
	return out
}

func CreatePublisherBlockListAppendPublisherUrlsBatchCall(id string, params CreatePublisherBlockListAppendPublisherUrlsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "append_publisher_urls"), params.ToParams(), options...)
}

func NewCreatePublisherBlockListAppendPublisherUrlsBatchRequest(id string, params CreatePublisherBlockListAppendPublisherUrlsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreatePublisherBlockListAppendPublisherUrlsBatchCall(id, params, options...))
}

func DecodeCreatePublisherBlockListAppendPublisherUrlsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreatePublisherBlockListAppendPublisherUrls(ctx context.Context, client *core.Client, id string, params CreatePublisherBlockListAppendPublisherUrlsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreatePublisherBlockListAppendPublisherUrlsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPublisherBlockListPagedWebPublishersParams struct {
	DraftID *core.ID    `facebook:"draft_id"`
	Extra   core.Params `facebook:"-"`
}

func (params GetPublisherBlockListPagedWebPublishersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DraftID != nil {
		out["draft_id"] = *params.DraftID
	}
	return out
}

func GetPublisherBlockListPagedWebPublishersBatchCall(id string, params GetPublisherBlockListPagedWebPublishersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "paged_web_publishers"), params.ToParams(), options...)
}

func NewGetPublisherBlockListPagedWebPublishersBatchRequest(id string, params GetPublisherBlockListPagedWebPublishersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WebPublisher]] {
	return core.NewBatchRequest[core.Cursor[objects.WebPublisher]](GetPublisherBlockListPagedWebPublishersBatchCall(id, params, options...))
}

func DecodeGetPublisherBlockListPagedWebPublishersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WebPublisher], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WebPublisher]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPublisherBlockListPagedWebPublishers(ctx context.Context, client *core.Client, id string, params GetPublisherBlockListPagedWebPublishersParams) (*core.Cursor[objects.WebPublisher], error) {
	var out core.Cursor[objects.WebPublisher]
	call := GetPublisherBlockListPagedWebPublishersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeletePublisherBlockListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePublisherBlockListBatchCall(id string, params DeletePublisherBlockListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePublisherBlockListBatchRequest(id string, params DeletePublisherBlockListParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePublisherBlockListBatchCall(id, params, options...))
}

func DecodeDeletePublisherBlockListBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePublisherBlockList(ctx context.Context, client *core.Client, id string, params DeletePublisherBlockListParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeletePublisherBlockListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPublisherBlockListParams struct {
	AccountID  *core.ID    `facebook:"account_id"`
	BusinessID *core.ID    `facebook:"business_id"`
	DraftID    *core.ID    `facebook:"draft_id"`
	Extra      core.Params `facebook:"-"`
}

func (params GetPublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.DraftID != nil {
		out["draft_id"] = *params.DraftID
	}
	return out
}

func GetPublisherBlockListBatchCall(id string, params GetPublisherBlockListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPublisherBlockListBatchRequest(id string, params GetPublisherBlockListParams, options ...core.BatchOption) *core.BatchRequest[objects.PublisherBlockList] {
	return core.NewBatchRequest[objects.PublisherBlockList](GetPublisherBlockListBatchCall(id, params, options...))
}

func DecodeGetPublisherBlockListBatchResponse(response *core.BatchResponse) (*objects.PublisherBlockList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PublisherBlockList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPublisherBlockList(ctx context.Context, client *core.Client, id string, params GetPublisherBlockListParams) (*objects.PublisherBlockList, error) {
	var out objects.PublisherBlockList
	call := GetPublisherBlockListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdatePublisherBlockListParams struct {
	Spec  map[string]interface{} `facebook:"spec"`
	Extra core.Params            `facebook:"-"`
}

func (params UpdatePublisherBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["spec"] = params.Spec
	return out
}

func UpdatePublisherBlockListBatchCall(id string, params UpdatePublisherBlockListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdatePublisherBlockListBatchRequest(id string, params UpdatePublisherBlockListParams, options ...core.BatchOption) *core.BatchRequest[objects.PublisherBlockList] {
	return core.NewBatchRequest[objects.PublisherBlockList](UpdatePublisherBlockListBatchCall(id, params, options...))
}

func DecodeUpdatePublisherBlockListBatchResponse(response *core.BatchResponse) (*objects.PublisherBlockList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PublisherBlockList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdatePublisherBlockList(ctx context.Context, client *core.Client, id string, params UpdatePublisherBlockListParams) (*objects.PublisherBlockList, error) {
	var out objects.PublisherBlockList
	call := UpdatePublisherBlockListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
