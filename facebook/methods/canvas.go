package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCanvasPreviewParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasPreviewParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasPreviewBatchCall(id string, params GetCanvasPreviewParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "preview"), params.ToParams(), options...)
}

func NewGetCanvasPreviewBatchRequest(id string, params GetCanvasPreviewParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CanvasPreview]] {
	return core.NewBatchRequest[core.Cursor[objects.CanvasPreview]](GetCanvasPreviewBatchCall(id, params, options...))
}

func DecodeGetCanvasPreviewBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CanvasPreview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CanvasPreview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCanvasPreviewWithResponse(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewParams) (*core.Cursor[objects.CanvasPreview], *core.Response, error) {
	var out core.Cursor[objects.CanvasPreview]
	call := GetCanvasPreviewBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCanvasPreview(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewParams) (*core.Cursor[objects.CanvasPreview], error) {
	out, _, err := GetCanvasPreviewWithResponse(ctx, client, id, params)
	return out, err
}

type GetCanvasPreviewsParams struct {
	UserIds *[]core.ID  `facebook:"user_ids"`
	Extra   core.Params `facebook:"-"`
}

func (params GetCanvasPreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UserIds != nil {
		out["user_ids"] = *params.UserIds
	}
	return out
}

func GetCanvasPreviewsBatchCall(id string, params GetCanvasPreviewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "previews"), params.ToParams(), options...)
}

func NewGetCanvasPreviewsBatchRequest(id string, params GetCanvasPreviewsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.TextWithEntities]] {
	return core.NewBatchRequest[core.Cursor[objects.TextWithEntities]](GetCanvasPreviewsBatchCall(id, params, options...))
}

func DecodeGetCanvasPreviewsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.TextWithEntities], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.TextWithEntities]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCanvasPreviewsWithResponse(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewsParams) (*core.Cursor[objects.TextWithEntities], *core.Response, error) {
	var out core.Cursor[objects.TextWithEntities]
	call := GetCanvasPreviewsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCanvasPreviews(ctx context.Context, client *core.Client, id string, params GetCanvasPreviewsParams) (*core.Cursor[objects.TextWithEntities], error) {
	out, _, err := GetCanvasPreviewsWithResponse(ctx, client, id, params)
	return out, err
}

type GetCanvasParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCanvasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCanvasBatchCall(id string, params GetCanvasParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCanvasBatchRequest(id string, params GetCanvasParams, options ...core.BatchOption) *core.BatchRequest[objects.Canvas] {
	return core.NewBatchRequest[objects.Canvas](GetCanvasBatchCall(id, params, options...))
}

func DecodeGetCanvasBatchResponse(response *core.BatchResponse) (*objects.Canvas, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Canvas
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCanvasWithResponse(ctx context.Context, client *core.Client, id string, params GetCanvasParams) (*objects.Canvas, *core.Response, error) {
	var out objects.Canvas
	call := GetCanvasBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCanvas(ctx context.Context, client *core.Client, id string, params GetCanvasParams) (*objects.Canvas, error) {
	out, _, err := GetCanvasWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateCanvasParams struct {
	BackgroundColor           *string     `facebook:"background_color"`
	BodyElementIds            *[]core.ID  `facebook:"body_element_ids"`
	EnableSwipeToOpen         *bool       `facebook:"enable_swipe_to_open"`
	HeroAssetFacebookPostID   *core.ID    `facebook:"hero_asset_facebook_post_id"`
	HeroAssetInstagramMediaID *core.ID    `facebook:"hero_asset_instagram_media_id"`
	IsHidden                  *bool       `facebook:"is_hidden"`
	IsPublished               *bool       `facebook:"is_published"`
	Name                      *string     `facebook:"name"`
	SourceTemplateID          *core.ID    `facebook:"source_template_id"`
	Extra                     core.Params `facebook:"-"`
}

func (params UpdateCanvasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BackgroundColor != nil {
		out["background_color"] = *params.BackgroundColor
	}
	if params.BodyElementIds != nil {
		out["body_element_ids"] = *params.BodyElementIds
	}
	if params.EnableSwipeToOpen != nil {
		out["enable_swipe_to_open"] = *params.EnableSwipeToOpen
	}
	if params.HeroAssetFacebookPostID != nil {
		out["hero_asset_facebook_post_id"] = *params.HeroAssetFacebookPostID
	}
	if params.HeroAssetInstagramMediaID != nil {
		out["hero_asset_instagram_media_id"] = *params.HeroAssetInstagramMediaID
	}
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.SourceTemplateID != nil {
		out["source_template_id"] = *params.SourceTemplateID
	}
	return out
}

func UpdateCanvasBatchCall(id string, params UpdateCanvasParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCanvasBatchRequest(id string, params UpdateCanvasParams, options ...core.BatchOption) *core.BatchRequest[objects.Canvas] {
	return core.NewBatchRequest[objects.Canvas](UpdateCanvasBatchCall(id, params, options...))
}

func DecodeUpdateCanvasBatchResponse(response *core.BatchResponse) (*objects.Canvas, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Canvas
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCanvasWithResponse(ctx context.Context, client *core.Client, id string, params UpdateCanvasParams) (*objects.Canvas, *core.Response, error) {
	var out objects.Canvas
	call := UpdateCanvasBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateCanvas(ctx context.Context, client *core.Client, id string, params UpdateCanvasParams) (*objects.Canvas, error) {
	out, _, err := UpdateCanvasWithResponse(ctx, client, id, params)
	return out, err
}
