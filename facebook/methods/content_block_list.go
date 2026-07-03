package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetContentBlockListAppliedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListAppliedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListAppliedAdAccountsBatchCall(id string, params GetContentBlockListAppliedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "applied_ad_accounts"), params.ToParams(), options...)
}

func NewGetContentBlockListAppliedAdAccountsBatchRequest(id string, params GetContentBlockListAppliedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetContentBlockListAppliedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetContentBlockListAppliedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetContentBlockListAppliedAdAccounts(ctx context.Context, client *core.Client, id string, params GetContentBlockListAppliedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetContentBlockListAppliedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetContentBlockListFacebookContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListFacebookContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListFacebookContentBatchCall(id string, params GetContentBlockListFacebookContentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "facebook_content"), params.ToParams(), options...)
}

func NewGetContentBlockListFacebookContentBatchRequest(id string, params GetContentBlockListFacebookContentParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Postish]] {
	return core.NewBatchRequest[core.Cursor[objects.Postish]](GetContentBlockListFacebookContentBatchCall(id, params, options...))
}

func DecodeGetContentBlockListFacebookContentBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Postish], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Postish]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetContentBlockListFacebookContent(ctx context.Context, client *core.Client, id string, params GetContentBlockListFacebookContentParams) (*core.Cursor[objects.Postish], error) {
	var out core.Cursor[objects.Postish]
	call := GetContentBlockListFacebookContentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetContentBlockListInstagramContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListInstagramContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListInstagramContentBatchCall(id string, params GetContentBlockListInstagramContentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_content"), params.ToParams(), options...)
}

func NewGetContentBlockListInstagramContentBatchRequest(id string, params GetContentBlockListInstagramContentParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetContentBlockListInstagramContentBatchCall(id, params, options...))
}

func DecodeGetContentBlockListInstagramContentBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetContentBlockListInstagramContent(ctx context.Context, client *core.Client, id string, params GetContentBlockListInstagramContentParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	call := GetContentBlockListInstagramContentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetContentBlockListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetContentBlockListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetContentBlockListBatchCall(id string, params GetContentBlockListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetContentBlockListBatchRequest(id string, params GetContentBlockListParams, options ...core.BatchOption) *core.BatchRequest[objects.ContentBlockList] {
	return core.NewBatchRequest[objects.ContentBlockList](GetContentBlockListBatchCall(id, params, options...))
}

func DecodeGetContentBlockListBatchResponse(response *core.BatchResponse) (*objects.ContentBlockList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ContentBlockList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetContentBlockList(ctx context.Context, client *core.Client, id string, params GetContentBlockListParams) (*objects.ContentBlockList, error) {
	var out objects.ContentBlockList
	call := GetContentBlockListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
