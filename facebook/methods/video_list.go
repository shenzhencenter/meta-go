package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoListVideosParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoListVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoListVideosBatchCall(id string, params GetVideoListVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewGetVideoListVideosBatchRequest(id string, params GetVideoListVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetVideoListVideosBatchCall(id, params, options...))
}

func DecodeGetVideoListVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoListVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetVideoListVideosParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetVideoListVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetVideoListVideos(ctx context.Context, client *core.Client, id string, params GetVideoListVideosParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetVideoListVideosWithResponse(ctx, client, id, params)
	return out, err
}

type GetVideoListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoListBatchCall(id string, params GetVideoListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoListBatchRequest(id string, params GetVideoListParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoList] {
	return core.NewBatchRequest[objects.VideoList](GetVideoListBatchCall(id, params, options...))
}

func DecodeGetVideoListBatchResponse(response *core.BatchResponse) (*objects.VideoList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoListWithResponse(ctx context.Context, client *core.Client, id string, params GetVideoListParams) (*objects.VideoList, *core.Response, error) {
	var out objects.VideoList
	call := GetVideoListBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetVideoList(ctx context.Context, client *core.Client, id string, params GetVideoListParams) (*objects.VideoList, error) {
	out, _, err := GetVideoListWithResponse(ctx, client, id, params)
	return out, err
}
